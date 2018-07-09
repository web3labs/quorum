// Package server contains the core server components.
package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/blk-io/crux/api"
	"github.com/blk-io/crux/utils"
	log "github.com/sirupsen/logrus"
	"encoding/hex"
	"net/textproto"
	"net/http/httputil"
	"os"
	"github.com/kevinburke/nacl"
	"time"
	"crypto/tls"
	"golang.org/x/net/context"
	"golang.org/x/crypto/acme/autocert"
)

// Enclave is the interface used by the transaction enclaves.
type Enclave interface {
	Store(message *[]byte, sender []byte, recipients [][]byte) ([]byte, error)
	StorePayload(encoded []byte) ([]byte, error)
	StorePayloadGrpc(epl api.EncryptedPayload, encoded []byte) ([]byte, error)
	Retrieve(digestHash *[]byte, to *[]byte) ([]byte, error)
	RetrieveDefault(digestHash *[]byte) ([]byte, error)
	RetrieveFor(digestHash *[]byte, reqRecipient *[]byte) (*[]byte, error)
	RetrieveAllFor(reqRecipient *[]byte) error
	Delete(digestHash *[]byte) error
	UpdatePartyInfo(encoded []byte)
	UpdatePartyInfoGrpc(url string, recipients map[[nacl.KeySize]byte]string, parties map[string]bool)
	GetEncodedPartyInfo() []byte
	GetEncodedPartyInfoGrpc() []byte
	GetPartyInfo() (url string, recipients map[[nacl.KeySize]byte]string, parties map[string]bool)
}

// TransactionManager is responsible for handling all transaction requests.
type TransactionManager struct {
	Enclave Enclave
}

const upCheckResponse = "I'm up!"
const apiVersion = "0.3.2"

const version = "/version"
const upCheck = "/upcheck"
const push = "/push"
const resend = "/resend"
const partyInfo = "/partyinfo"
const send = "/send"
const sendRaw = "/sendraw"
const receive = "/receive"
const receiveRaw = "/receiveraw"
const delete = "/delete"

const hFrom = "c11n-from"
const hTo = "c11n-to"
const hKey = "c11n-key"

func requestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if log.GetLevel() == log.DebugLevel {
			dump, err := httputil.DumpRequest(r, true)
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}

			log.Debugf("%q", dump)
		}

		handler.ServeHTTP(w, r)
	})
}

// Init initializes a new TransactionManager instance.
func Init(enc Enclave, port int, ipcPath string, grpc bool, tls bool, certFile, keyFile string) (TransactionManager, error) {
	tm := TransactionManager{Enclave : enc}
	var err error

	if grpc == true {
		err = tm.startRpcServer(port, ipcPath, tls, certFile, keyFile)

	} else {
		err = tm.startHttpserver(port, ipcPath, tls, certFile, keyFile)
	}

	return tm, err
}

func (tm *TransactionManager) makeHTTPServer() *http.Server {
	httpServer := http.NewServeMux()
	httpServer.HandleFunc(upCheck, tm.upcheck)
	httpServer.HandleFunc(version, tm.version)
	httpServer.HandleFunc(push, tm.push)
	httpServer.HandleFunc(resend, tm.resend)
	httpServer.HandleFunc(partyInfo, tm.partyInfo)

	// set timeouts so that a slow or malicious client doesn't
	// hold resources forever
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      httpServer,
	}
}

func (tm *TransactionManager) startHttpserver(port int, ipcPath string, tls1 bool, certFile, keyFile string) error {
	httpServer := tm.makeHTTPServer()

	serverUrl := "localhost:" + strconv.Itoa(port)
	httpServer.Addr = fmt.Sprintf(":%d", port)
	if tls1 {
		var m *autocert.Manager
		hostPolicy := func(ctx context.Context, host string) error {
			// Note: change to your real host
			allowedHost := httpServer.Addr
			if host == allowedHost {
				return nil
			}
			return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
		}
		dataDir := "."
		m = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: hostPolicy,
			Cache:      autocert.DirCache(dataDir),
		}

		httpServer.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

		go func() {
			fmt.Printf("Starting HTTPS server on %s\n", httpServer.Addr)
			err := httpServer.ListenAndServeTLS("", "")
			if err != nil {
				log.Fatalf("httpsSrv.ListendAndServeTLS() failed with %s", err)
			}
		}()
	} else {
		httpServer := http.NewServeMux()
		go func() {
			log.Fatal(http.ListenAndServe(serverUrl, requestLogger(httpServer)))
		}()
		log.Infof("HTTP server is running at: %s", serverUrl)
	}

	// Restricted to IPC
	ipcServer := http.NewServeMux()
	ipcServer.HandleFunc(upCheck, tm.upcheck)
	ipcServer.HandleFunc(version, tm.version)
	ipcServer.HandleFunc(send, tm.send)
	ipcServer.HandleFunc(sendRaw, tm.sendRaw)
	ipcServer.HandleFunc(receive, tm.receive)
	ipcServer.HandleFunc(receiveRaw, tm.receiveRaw)
	ipcServer.HandleFunc(delete, tm.delete)

	ipc, err := utils.CreateIpcSocket(ipcPath)
	if err != nil {
		log.Fatalf("Failed to start IPC Server at %s", ipcPath)
	}
	go func() {
		log.Fatal(http.Serve(ipc, requestLogger(ipcServer)))
	}()
	log.Infof("IPC server is running at: %s", ipcPath)

	return err
}

func CheckCertFiles(certFile, keyFile string) error {
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		return err
	} else if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		return err
	}
	return nil
}

func (s *TransactionManager) upcheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, upCheckResponse)
}

func (s *TransactionManager) version(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, apiVersion)
}

func (s *TransactionManager) send(w http.ResponseWriter, req *http.Request) {
	var sendReq api.SendRequest
	err := json.NewDecoder(req.Body).Decode(&sendReq)
	req.Body.Close()
	if err != nil {
		invalidBody(w, req, err)
		return
	}

	payload, err := base64.StdEncoding.DecodeString(sendReq.Payload)
	if err != nil {
		decodeError(w, req, "payload", sendReq.Payload, err)
		return
	}

	var key []byte
	key, err = s.processSend(w, req, sendReq.From, sendReq.To, &payload)

	if err != nil {
		log.Error(err)
		badRequest(w,
			fmt.Sprintf("Unable to store key: %s, with payload: %s, error: %s\n",
				key, payload, err))
	} else {
		encodedKey := base64.StdEncoding.EncodeToString(key)
		sendResp := api.SendResponse{Key: encodedKey}
		json.NewEncoder(w).Encode(sendResp)
		w.Header().Set("Content-Type", "application/json")
	}
}

func (s *TransactionManager) sendRaw(w http.ResponseWriter, req *http.Request) {

	from := req.Header.Get(hFrom)

	to, ok := req.Header[hTo]
	if !ok {
		to, ok = req.Header[textproto.CanonicalMIMEHeaderKey(hTo)]
		if !ok {
			to = []string{}
		}
	}

	payload, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		invalidBody(w, req, err)
		return
	}

	var key []byte
	key, err = s.processSend(w, req, from, to, &payload)
	if err != nil {
		internalServerError(w, "Unable to process request")
		return
	}

	// Uncomment the below for Quorum v2.0.2 onwards
	// see https://github.com/jpmorganchase/quorum/commit/ee498061b5a74bf1f3290139a53840345fa038cb#diff-63fbbd6b2c0487b8cd4445e881822cdd
	// encodedKey := base64.StdEncoding.EncodeToString(key)
	// fmt.Fprint(w, encodedKey)
	// Then delete this line
	w.Write(key)
}

func (s *TransactionManager) processSend(
	w http.ResponseWriter, req *http.Request,
	b64from string,
	b64recipients []string,
	payload *[]byte) ([]byte, error) {

	log.WithFields(log.Fields{
		"b64From" : b64from,
		"b64Recipients": b64recipients,
		"payload": hex.EncodeToString(*payload),}).Debugf(
		"Processing send request")

	sender, err := base64.StdEncoding.DecodeString(b64from)
	if err != nil {
		decodeError(w, req, "sender", b64from, err)
		return nil, err
	}

	recipients := make([][]byte, len(b64recipients))
	for i, value := range b64recipients {
		recipient, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			decodeError(w, req, "recipient", value, err)
			return nil, err
		} else {
			recipients[i] = recipient
		}
	}

	return s.Enclave.Store(payload, sender, recipients)
}

func (s *TransactionManager) receive(w http.ResponseWriter, req *http.Request) {
	var receiveReq api.ReceiveRequest
	err := json.NewDecoder(req.Body).Decode(&receiveReq)
	req.Body.Close()
	if err != nil {
		invalidBody(w, req, err)
		return
	}

	var payload []byte
	payload, err = s.processReceive(w, req, receiveReq.Key, receiveReq.To)

	if err != nil {
		badRequest(w,
			fmt.Sprintf("Unable to retrieve payload for key: %s, error: %s\n",
				receiveReq.Key, err))
	} else {
		encodedPayload := base64.StdEncoding.EncodeToString(payload)
		sendResp := api.ReceiveResponse{Payload: encodedPayload}
		json.NewEncoder(w).Encode(sendResp)
		w.Header().Set("Content-Type", "application/json")
	}
}

func (s *TransactionManager) receiveRaw(w http.ResponseWriter, req *http.Request) {

	key := req.Header.Get(hKey)
	if key == "" {
		badRequest(w, "key not specified")
		return
	}

	to := req.Header.Get(hTo)

	payload, err := s.processReceive(w, req, key, to)

	if err != nil {
		badRequest(w, fmt.Sprintln(err))
		return
	}

	w.Write(payload)
}

func (s *TransactionManager) processReceive(
	w http.ResponseWriter, req *http.Request, b64Key, b64To string) ([]byte, error) {

	key, err := base64.StdEncoding.DecodeString(b64Key)
	if err != nil {
		return nil, fmt.Errorf("unable to decode key: %s", b64Key)
	}

	if b64To != "" {
		to, err := base64.StdEncoding.DecodeString(b64To)
		if err != nil {
			return nil, fmt.Errorf("unable to decode to: %s", b64Key)
		}

		return s.Enclave.Retrieve(&key, &to)
	} else {
		return s.Enclave.RetrieveDefault(&key)
	}
}

func (s *TransactionManager) delete(w http.ResponseWriter, req *http.Request) {
	var deleteReq api.DeleteRequest
	err := json.NewDecoder(req.Body).Decode(&deleteReq)
	req.Body.Close()
	if err != nil {
		invalidBody(w, req, err)
		return
	}
	key, err := base64.StdEncoding.DecodeString(deleteReq.Key)
	if err != nil {
		decodeError(w, req, "key", deleteReq.Key, err)
	} else {
		err = s.Enclave.Delete(&key)
		if err != nil {
			badRequest(w, fmt.Sprintf("Unable to delete key: %s, error: %s\n", key, err))
		}
	}
}

func (s *TransactionManager) push(w http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		internalServerError(w, fmt.Sprintf("Unable to read request body, error: %s\n", err))
		return
	}

	digestHash, err := s.Enclave.StorePayload(payload)
	if err != nil {
		badRequest(w, fmt.Sprintf("Unable to store payload, error: %s\n", err))
		return
	}

	w.Write(digestHash)
}

func (s *TransactionManager) resend(w http.ResponseWriter, req *http.Request) {
	var resendReq api.ResendRequest
	err := json.NewDecoder(req.Body).Decode(&resendReq)
	req.Body.Close()
	if err != nil {
		invalidBody(w, req, err)
		return
	}

	var publicKey []byte
	publicKey, err = base64.StdEncoding.DecodeString(resendReq.PublicKey)
	if err != nil {
		decodeError(w, req, "publicKey", resendReq.PublicKey, err)
		return
	}

	if resendReq.Type == "all" {
		err = s.Enclave.RetrieveAllFor(&publicKey)
		if err != nil {
			invalidBody(w, req, err)
		}
	} else if resendReq.Type == "individual" {
		var key []byte
		key, err = base64.StdEncoding.DecodeString(resendReq.Key)
		if err != nil {
			decodeError(w, req, "key", resendReq.Key, err)
			return
		}

		var encodedPl *[]byte
		encodedPl, err = s.Enclave.RetrieveFor(&key, &publicKey)
		if err != nil {
			invalidBody(w, req, err)
			return
		}
		w.Write(*encodedPl)
	}
}

func (s *TransactionManager) partyInfo(w http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		internalServerError(w, fmt.Sprintf("Unable to read request body, error: %s\n", err))
		return
	} else {
		s.Enclave.UpdatePartyInfo(payload)
		w.Write(s.Enclave.GetEncodedPartyInfo())
	}
}

func invalidBody(w http.ResponseWriter, req *http.Request, err error) {
	badRequest(w, fmt.Sprintf("Invalid request: %s, error: %s\n", req.URL, err))
}

func decodeError(w http.ResponseWriter, req *http.Request, name string, value string, err error) {
	badRequest(w,
		fmt.Sprintf("Invalid request: %s, unable to decode %s: %s, error: %s\n",
			req.URL, name, value, err))
}

func badRequest(w http.ResponseWriter, message string) {
	log.Error(message)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, message)
}

func internalServerError(w http.ResponseWriter, message string) {
	log.Error(message)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, message)
}
