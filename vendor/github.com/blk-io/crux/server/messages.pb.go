// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ApiVersion struct {
	Version              string   `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiVersion) Reset()         { *m = ApiVersion{} }
func (m *ApiVersion) String() string { return proto.CompactTextString(m) }
func (*ApiVersion) ProtoMessage()    {}
func (*ApiVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{0}
}
func (m *ApiVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiVersion.Unmarshal(m, b)
}
func (m *ApiVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiVersion.Marshal(b, m, deterministic)
}
func (dst *ApiVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiVersion.Merge(dst, src)
}
func (m *ApiVersion) XXX_Size() int {
	return xxx_messageInfo_ApiVersion.Size(m)
}
func (m *ApiVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ApiVersion proto.InternalMessageInfo

func (m *ApiVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type UpCheckResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpCheckResponse) Reset()         { *m = UpCheckResponse{} }
func (m *UpCheckResponse) String() string { return proto.CompactTextString(m) }
func (*UpCheckResponse) ProtoMessage()    {}
func (*UpCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{1}
}
func (m *UpCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpCheckResponse.Unmarshal(m, b)
}
func (m *UpCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpCheckResponse.Marshal(b, m, deterministic)
}
func (dst *UpCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpCheckResponse.Merge(dst, src)
}
func (m *UpCheckResponse) XXX_Size() int {
	return xxx_messageInfo_UpCheckResponse.Size(m)
}
func (m *UpCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpCheckResponse proto.InternalMessageInfo

func (m *UpCheckResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendRequest struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	To                   []string `protobuf:"bytes,3,rep,name=to" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{2}
}
func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (dst *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(dst, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SendRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendRequest) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

type SendResponse struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{3}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (dst *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(dst, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type ReceiveRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveRequest) Reset()         { *m = ReceiveRequest{} }
func (m *ReceiveRequest) String() string { return proto.CompactTextString(m) }
func (*ReceiveRequest) ProtoMessage()    {}
func (*ReceiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{4}
}
func (m *ReceiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveRequest.Unmarshal(m, b)
}
func (m *ReceiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveRequest.Marshal(b, m, deterministic)
}
func (dst *ReceiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveRequest.Merge(dst, src)
}
func (m *ReceiveRequest) XXX_Size() int {
	return xxx_messageInfo_ReceiveRequest.Size(m)
}
func (m *ReceiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveRequest proto.InternalMessageInfo

func (m *ReceiveRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ReceiveRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type ReceiveResponse struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveResponse) Reset()         { *m = ReceiveResponse{} }
func (m *ReceiveResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiveResponse) ProtoMessage()    {}
func (*ReceiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{5}
}
func (m *ReceiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveResponse.Unmarshal(m, b)
}
func (m *ReceiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveResponse.Marshal(b, m, deterministic)
}
func (dst *ReceiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveResponse.Merge(dst, src)
}
func (m *ReceiveResponse) XXX_Size() int {
	return xxx_messageInfo_ReceiveResponse.Size(m)
}
func (m *ReceiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveResponse proto.InternalMessageInfo

func (m *ReceiveResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PartyInfo struct {
	Url                  string            `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Receipients          map[string][]byte `protobuf:"bytes,2,rep,name=receipients" json:"receipients,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Parties              map[string]bool   `protobuf:"bytes,3,rep,name=parties" json:"parties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PartyInfo) Reset()         { *m = PartyInfo{} }
func (m *PartyInfo) String() string { return proto.CompactTextString(m) }
func (*PartyInfo) ProtoMessage()    {}
func (*PartyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{6}
}
func (m *PartyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartyInfo.Unmarshal(m, b)
}
func (m *PartyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartyInfo.Marshal(b, m, deterministic)
}
func (dst *PartyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartyInfo.Merge(dst, src)
}
func (m *PartyInfo) XXX_Size() int {
	return xxx_messageInfo_PartyInfo.Size(m)
}
func (m *PartyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PartyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PartyInfo proto.InternalMessageInfo

func (m *PartyInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PartyInfo) GetReceipients() map[string][]byte {
	if m != nil {
		return m.Receipients
	}
	return nil
}

func (m *PartyInfo) GetParties() map[string]bool {
	if m != nil {
		return m.Parties
	}
	return nil
}

type PartyInfoResponse struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartyInfoResponse) Reset()         { *m = PartyInfoResponse{} }
func (m *PartyInfoResponse) String() string { return proto.CompactTextString(m) }
func (*PartyInfoResponse) ProtoMessage()    {}
func (*PartyInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{7}
}
func (m *PartyInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartyInfoResponse.Unmarshal(m, b)
}
func (m *PartyInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartyInfoResponse.Marshal(b, m, deterministic)
}
func (dst *PartyInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartyInfoResponse.Merge(dst, src)
}
func (m *PartyInfoResponse) XXX_Size() int {
	return xxx_messageInfo_PartyInfoResponse.Size(m)
}
func (m *PartyInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PartyInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PartyInfoResponse proto.InternalMessageInfo

func (m *PartyInfoResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type EncryptedPayload struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	CipherText           []byte   `protobuf:"bytes,2,opt,name=cipherText,proto3" json:"cipherText,omitempty"`
	Nonce                []byte   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ReciepientBoxes      [][]byte `protobuf:"bytes,4,rep,name=reciepientBoxes,proto3" json:"reciepientBoxes,omitempty"`
	ReciepientNonce      []byte   `protobuf:"bytes,5,opt,name=reciepientNonce,proto3" json:"reciepientNonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptedPayload) Reset()         { *m = EncryptedPayload{} }
func (m *EncryptedPayload) String() string { return proto.CompactTextString(m) }
func (*EncryptedPayload) ProtoMessage()    {}
func (*EncryptedPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{8}
}
func (m *EncryptedPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedPayload.Unmarshal(m, b)
}
func (m *EncryptedPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedPayload.Marshal(b, m, deterministic)
}
func (dst *EncryptedPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedPayload.Merge(dst, src)
}
func (m *EncryptedPayload) XXX_Size() int {
	return xxx_messageInfo_EncryptedPayload.Size(m)
}
func (m *EncryptedPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedPayload.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedPayload proto.InternalMessageInfo

func (m *EncryptedPayload) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *EncryptedPayload) GetCipherText() []byte {
	if m != nil {
		return m.CipherText
	}
	return nil
}

func (m *EncryptedPayload) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EncryptedPayload) GetReciepientBoxes() [][]byte {
	if m != nil {
		return m.ReciepientBoxes
	}
	return nil
}

func (m *EncryptedPayload) GetReciepientNonce() []byte {
	if m != nil {
		return m.ReciepientNonce
	}
	return nil
}

type PushPayload struct {
	Ep                   *EncryptedPayload `protobuf:"bytes,1,opt,name=ep" json:"ep,omitempty"`
	Encoded              []byte            `protobuf:"bytes,2,opt,name=encoded,proto3" json:"encoded,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PushPayload) Reset()         { *m = PushPayload{} }
func (m *PushPayload) String() string { return proto.CompactTextString(m) }
func (*PushPayload) ProtoMessage()    {}
func (*PushPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_68f7ab936725f9e3, []int{9}
}
func (m *PushPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushPayload.Unmarshal(m, b)
}
func (m *PushPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushPayload.Marshal(b, m, deterministic)
}
func (dst *PushPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushPayload.Merge(dst, src)
}
func (m *PushPayload) XXX_Size() int {
	return xxx_messageInfo_PushPayload.Size(m)
}
func (m *PushPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PushPayload.DiscardUnknown(m)
}

var xxx_messageInfo_PushPayload proto.InternalMessageInfo

func (m *PushPayload) GetEp() *EncryptedPayload {
	if m != nil {
		return m.Ep
	}
	return nil
}

func (m *PushPayload) GetEncoded() []byte {
	if m != nil {
		return m.Encoded
	}
	return nil
}

func init() {
	proto.RegisterType((*ApiVersion)(nil), "server.ApiVersion")
	proto.RegisterType((*UpCheckResponse)(nil), "server.UpCheckResponse")
	proto.RegisterType((*SendRequest)(nil), "server.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "server.SendResponse")
	proto.RegisterType((*ReceiveRequest)(nil), "server.ReceiveRequest")
	proto.RegisterType((*ReceiveResponse)(nil), "server.ReceiveResponse")
	proto.RegisterType((*PartyInfo)(nil), "server.PartyInfo")
	proto.RegisterMapType((map[string]bool)(nil), "server.PartyInfo.PartiesEntry")
	proto.RegisterMapType((map[string][]byte)(nil), "server.PartyInfo.ReceipientsEntry")
	proto.RegisterType((*PartyInfoResponse)(nil), "server.PartyInfoResponse")
	proto.RegisterType((*EncryptedPayload)(nil), "server.EncryptedPayload")
	proto.RegisterType((*PushPayload)(nil), "server.PushPayload")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_messages_68f7ab936725f9e3) }

var fileDescriptor_messages_68f7ab936725f9e3 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0x52, 0xe2, 0xd6, 0x23, 0x61, 0xbb, 0xa2, 0x14, 0x61, 0x4a, 0x10, 0x7b, 0x28, 0x82,
	0x50, 0x07, 0xdc, 0x4b, 0xc8, 0xa1, 0xd0, 0x8f, 0x1c, 0x4a, 0xa1, 0xb8, 0xdb, 0x8f, 0xbb, 0x2a,
	0x4d, 0xec, 0x25, 0xce, 0xee, 0x76, 0x77, 0x6d, 0xa2, 0x9f, 0xd2, 0xff, 0xd1, 0x1f, 0x58, 0xf6,
	0x43, 0x8e, 0xea, 0x14, 0x9a, 0xdb, 0xbc, 0xe1, 0xbd, 0x37, 0x6f, 0x66, 0x25, 0x18, 0xdf, 0xa0,
	0xd6, 0xd5, 0x0a, 0xf5, 0x5c, 0x2a, 0x61, 0x44, 0x36, 0xd4, 0xa8, 0x76, 0xa8, 0x66, 0xcf, 0x57,
	0x42, 0xac, 0x36, 0x78, 0x56, 0x49, 0x76, 0x56, 0x71, 0x2e, 0x4c, 0x65, 0x98, 0xe0, 0x81, 0x45,
	0x5e, 0x00, 0xbc, 0x91, 0xec, 0x3b, 0x2a, 0xcd, 0x04, 0xcf, 0x72, 0x78, 0xb4, 0xf3, 0x65, 0x3e,
	0x28, 0x06, 0xe5, 0x88, 0x76, 0x90, 0x9c, 0xc2, 0xe4, 0x9b, 0x7c, 0xb7, 0xc6, 0xfa, 0x9a, 0xa2,
	0x96, 0x82, 0x6b, 0xb4, 0xe4, 0x30, 0xb2, 0x23, 0x07, 0x48, 0x3e, 0x42, 0xf2, 0x05, 0x79, 0x43,
	0xf1, 0xe7, 0x16, 0xb5, 0xb1, 0x44, 0x59, 0xb5, 0x1b, 0x51, 0x35, 0x8e, 0x98, 0xd2, 0x0e, 0x66,
	0x19, 0x1c, 0x5d, 0x29, 0x71, 0x93, 0x47, 0x4e, 0xef, 0xea, 0x6c, 0x0c, 0x91, 0x11, 0x79, 0x5c,
	0xc4, 0xe5, 0x88, 0x46, 0x46, 0x90, 0x02, 0x52, 0x6f, 0x16, 0xc6, 0x4e, 0x21, 0xbe, 0xc6, 0x36,
	0x38, 0xd9, 0x92, 0x2c, 0x60, 0x4c, 0xb1, 0x46, 0xb6, 0xc3, 0x6e, 0xe2, 0x3d, 0x4e, 0x70, 0xf5,
	0x73, 0xac, 0xeb, 0x29, 0x4c, 0xf6, 0x9a, 0xbb, 0x7d, 0xfe, 0x1d, 0x93, 0xfc, 0x8a, 0x60, 0xb4,
	0xac, 0x94, 0x69, 0x3f, 0xf0, 0x2b, 0x61, 0xcd, 0xb7, 0x6a, 0x13, 0x76, 0xb6, 0x65, 0xf6, 0x1e,
	0x12, 0x65, 0xcd, 0x24, 0x43, 0x6e, 0x74, 0x1e, 0x15, 0x71, 0x99, 0x2c, 0xc8, 0xdc, 0x3f, 0xc0,
	0x7c, 0xaf, 0x9c, 0xd3, 0x3b, 0xd2, 0x25, 0x37, 0xaa, 0xa5, 0x7d, 0x59, 0x76, 0x6e, 0xe7, 0x2b,
	0xc3, 0x50, 0xbb, 0xed, 0x93, 0xc5, 0xc9, 0x7d, 0x87, 0xa5, 0x27, 0x78, 0x75, 0x47, 0x9f, 0xbd,
	0x86, 0xe9, 0xa1, 0x75, 0xff, 0x04, 0x23, 0x7f, 0x82, 0xa7, 0x70, 0xbc, 0xab, 0x36, 0x5b, 0x74,
	0x57, 0x48, 0xa9, 0x07, 0x17, 0xd1, 0xf9, 0x60, 0x76, 0x01, 0x69, 0xdf, 0xf8, 0x7f, 0xda, 0xc7,
	0x3d, 0x2d, 0x79, 0x09, 0x4f, 0xf6, 0xf1, 0x1e, 0x70, 0xca, 0xdf, 0x03, 0x98, 0x5e, 0xf2, 0x5a,
	0xb5, 0xd2, 0x60, 0xb3, 0x0c, 0x9f, 0xc1, 0x33, 0x18, 0x6a, 0xe4, 0x0d, 0xaa, 0xc0, 0x0e, 0x28,
	0x3b, 0x01, 0xa8, 0x99, 0x5c, 0xa3, 0xfa, 0x8a, 0xb7, 0x26, 0xc4, 0xee, 0x75, 0x6c, 0x2a, 0x2e,
	0x78, 0x8d, 0x79, 0xec, 0x37, 0x72, 0x20, 0x2b, 0x61, 0xa2, 0xb0, 0x66, 0xe8, 0xae, 0xf1, 0x56,
	0xdc, 0xa2, 0xce, 0x8f, 0x8a, 0xb8, 0x4c, 0xe9, 0x61, 0xfb, 0x6f, 0xe6, 0x27, 0xe7, 0x74, 0xec,
	0x9c, 0x0e, 0xdb, 0xe4, 0x33, 0x24, 0xcb, 0xad, 0x5e, 0x77, 0x81, 0x4b, 0x88, 0x50, 0xba, 0xb0,
	0xc9, 0x22, 0xef, 0x5e, 0xe9, 0x70, 0x2d, 0x1a, 0xa1, 0xb4, 0x97, 0x40, 0x5e, 0x8b, 0x06, 0x9b,
	0x90, 0xbf, 0x83, 0x3f, 0x86, 0xee, 0x07, 0x7c, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x7b,
	0x28, 0x36, 0xb8, 0x03, 0x00, 0x00,
}