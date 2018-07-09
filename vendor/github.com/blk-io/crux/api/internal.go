package api

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/kevinburke/nacl"
	"github.com/blk-io/crux/utils"
	"encoding/hex"
	"encoding/json"
	"net/http/httputil"
	"fmt"
)

// EncryptedPayload is the struct used for storing all data associated with an encrypted
// transaction.
type EncryptedPayload struct {
	Sender         nacl.Key
	CipherText     []byte
	Nonce          nacl.Nonce
	RecipientBoxes [][]byte
	RecipientNonce nacl.Nonce
}

// PartyInfo is a struct that stores details of all enclave nodes (or parties) on the network.
type PartyInfo struct {
	url string  // URL identifying this node
	recipients map[[nacl.KeySize]byte]string  // public key -> URL
	parties    map[string]bool  // Node (or party) URLs
	client     utils.HttpClient
}

// GetRecipient retrieves the URL associated with the provided recipient.
func (s *PartyInfo) GetRecipient(key nacl.Key) (string, bool) {
	value, ok := s.recipients[*key]
	return value, ok
}

func (s *PartyInfo) GetAllValues() (string, map[[nacl.KeySize]byte]string, map[string]bool){
	return s.url, s.recipients, s.parties
}

// InitPartyInfo initializes a new PartyInfo store.
func InitPartyInfo(rawUrl string, otherNodes []string, client utils.HttpClient) PartyInfo {
	parties := make(map[string]bool)
	for _, node := range otherNodes {
		parties[node] = true
	}

	return PartyInfo{
		url:        rawUrl,
		recipients: make(map[[nacl.KeySize]byte]string),
		parties:    parties,
		client:     client,
	}
}

// CreatePartyInfo creates a new PartyInfo struct.
func CreatePartyInfo(
	url string,
	otherNodes []string,
	otherKeys []nacl.Key,
	client utils.HttpClient) PartyInfo {

	recipients := make(map[[nacl.KeySize]byte]string)
	parties := make(map[string]bool)
	for i, node := range otherNodes {
		parties[node] = true
		recipients[*otherKeys[i]] = node
	}

	return PartyInfo{
		url:        url,
		recipients: recipients,
		parties:    parties,
		client:     client,
	}
}

// RegisterPublicKeys associates the provided public keys with this node.
func (s *PartyInfo) RegisterPublicKeys(pubKeys []nacl.Key) {
	for _, pubKey := range pubKeys {
		s.recipients[*pubKey] = s.url
	}
}

// GetPartyInfo requests PartyInfo data from all remote nodes this node is aware of. The data
// provided in each response is applied to this node.
func (s *PartyInfo) GetPartyInfo(grpc bool) {
	encodedPartyInfo := EncodePartyInfo(*s)

	// First copy our endpoints as we update this map in place
	urls := make(map[string]bool)
	for k, v := range s.parties {
		urls[k] = v
	}

	for rawUrl := range urls {
		if rawUrl == s.url {
			continue
		}

		endPoint, err := utils.BuildUrl(rawUrl, "/partyinfo")

		if err != nil {
			log.WithFields(log.Fields{"rawUrl": rawUrl, "endPoint": "/partyinfo"}).Errorf(
				"Invalid endpoint provided")
		}

		var req *http.Request
		encoded := getEncoded(grpc, encodedPartyInfo)
		req, err = http.NewRequest("POST", endPoint, bytes.NewBuffer(encoded))

		if err != nil {
			log.WithField("url", rawUrl).Errorf(
				"Error creating /partyinfo request, %v", err)
			break
		}
		req.Header.Set("Content-Type", "application/octet-stream")

		logRequest(req)
		resp, err := s.client.Do(req)
		if err != nil {
			log.WithField("url", rawUrl).Errorf(
				"Error sending /partyinfo request, %v", err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.WithField("url", rawUrl).Errorf(
				"Error sending /partyinfo request, non-200 status code: %v", resp)
			continue
		}

		if grpc {
			err = s.updatePartyInfoGrpc(resp, rawUrl)
		} else {
			err = s.updatePartyInfo(resp, rawUrl)
		}
		if err != nil {
			break
		}
	}
}

func (s *PartyInfo) updatePartyInfoGrpc(resp *http.Response, rawUrl string) error {
	var partyInfoReq UpdatePartyInfo
	err := json.NewDecoder(resp.Body).Decode(&partyInfoReq)
	resp.Body.Close()
	if err != nil {
		log.WithField("url", rawUrl).Errorf(
			"Unable to read partyInfo response from host, %v", err)
		return err
	}
	pi, err := DecodePartyInfo(partyInfoReq.Payload)
	if err != nil {
		log.WithField("url", rawUrl).Errorf(
			"Unable to decode partyInfo response from host, %v", err)
		return err
	}
	s.UpdatePartyInfoGrpc(pi.url, pi.recipients, pi.parties)
	return nil
}

func (s *PartyInfo) updatePartyInfo(resp *http.Response, rawUrl string) error {
	var encoded []byte
	encoded, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.WithField("url", rawUrl).Errorf(
			"Unable to read partyInfo response from host, %v", err)
		return err
	}
	s.UpdatePartyInfo(encoded)
	return nil
}

func getEncoded(grpc bool, encodedPartyInfo []byte) []byte{
	if grpc{
		e, err := json.Marshal(UpdatePartyInfo{encodedPartyInfo})
		if err != nil{
			log.Errorf("Marshalling failed %v", err)
			return nil
		}
		return e
	}
	return encodedPartyInfo[:]
}
func (s *PartyInfo) PollPartyInfo(grpc bool) {
	time.Sleep(time.Duration(rand.Intn(16)) * time.Second)
	s.GetPartyInfo(grpc)

	ticker := time.NewTicker(2 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				s.GetPartyInfo(grpc)
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// UpdatePartyInfo updates the PartyInfo datastore with the provided encoded data.
// This can happen from the /partyinfo server endpoint being hit, or by a response from us hitting
// another nodes /partyinfo endpoint.
// TODO: Control access via a channel for updates.
func (s *PartyInfo) UpdatePartyInfo(encoded []byte) {
	log.Debugf("Updating party info payload: %s", hex.EncodeToString(encoded))
	pi, err := DecodePartyInfo(encoded)

	if err != nil {
		log.WithField("encoded", encoded).Errorf(
			"Unable to decode party info, error: %v", err)
	}

	for publicKey, url := range pi.recipients {
		// we should ignore messages about ourselves
		// in order to stop people masquerading as you, there
		// should be a digital signature associated with each
		// url -> node broadcast
		if url != s.url {
			s.recipients[publicKey] = url
		}
	}

	for url := range pi.parties {
		// we don't want to broadcast party info to ourselves
		s.parties[url] = true
	}
}

func (s *PartyInfo) UpdatePartyInfoGrpc(url string, recipients map[[nacl.KeySize]byte]string, parties map[string]bool) {
	for publicKey, url := range recipients {
		// we should ignore messages about ourselves
		// in order to stop people masquerading as you, there
		// should be a digital signature associated with each
		// url -> node broadcast
		if url != s.url {
			s.recipients[publicKey] = url
		}
	}

	for url := range parties {
		// we don't want to broadcast party info to ourselves
		s.parties[url] = true
	}
}

// Push is responsible for propagating the encoded payload to the given remote node.
func Push(encoded []byte, url string, client utils.HttpClient) (string, error) {

	endPoint, err := utils.BuildUrl(url, "/push")
	if err != nil {
		return "", err
	}

	var req *http.Request
	req, err = http.NewRequest("POST", endPoint, bytes.NewReader(encoded))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	logRequest(req)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code received: %v", resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func logRequest(r *http.Request) {
	if log.GetLevel() == log.DebugLevel {
		dump, err := httputil.DumpRequestOut(r, true)
		if err != nil {
			log.Fatal(err)
		}

		log.Debugf("%q", dump)
	}
}