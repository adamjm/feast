// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/types/FeatureRowExtended.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Error struct {
	Cause                string   `protobuf:"bytes,1,opt,name=cause,proto3" json:"cause,omitempty"`
	Transform            string   `protobuf:"bytes,2,opt,name=transform,proto3" json:"transform,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	StackTrace           string   `protobuf:"bytes,4,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_7823aa2c72575793, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

func (m *Error) GetTransform() string {
	if m != nil {
		return m.Transform
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetStackTrace() string {
	if m != nil {
		return m.StackTrace
	}
	return ""
}

type Attempt struct {
	Attempts             int32    `protobuf:"varint,1,opt,name=attempts,proto3" json:"attempts,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attempt) Reset()         { *m = Attempt{} }
func (m *Attempt) String() string { return proto.CompactTextString(m) }
func (*Attempt) ProtoMessage()    {}
func (*Attempt) Descriptor() ([]byte, []int) {
	return fileDescriptor_7823aa2c72575793, []int{1}
}

func (m *Attempt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attempt.Unmarshal(m, b)
}
func (m *Attempt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attempt.Marshal(b, m, deterministic)
}
func (m *Attempt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attempt.Merge(m, src)
}
func (m *Attempt) XXX_Size() int {
	return xxx_messageInfo_Attempt.Size(m)
}
func (m *Attempt) XXX_DiscardUnknown() {
	xxx_messageInfo_Attempt.DiscardUnknown(m)
}

var xxx_messageInfo_Attempt proto.InternalMessageInfo

func (m *Attempt) GetAttempts() int32 {
	if m != nil {
		return m.Attempts
	}
	return 0
}

func (m *Attempt) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type FeatureRowExtended struct {
	Row                  *FeatureRow          `protobuf:"bytes,1,opt,name=row,proto3" json:"row,omitempty"`
	LastAttempt          *Attempt             `protobuf:"bytes,2,opt,name=last_attempt,json=lastAttempt,proto3" json:"last_attempt,omitempty"`
	FirstSeen            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=first_seen,json=firstSeen,proto3" json:"first_seen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeatureRowExtended) Reset()         { *m = FeatureRowExtended{} }
func (m *FeatureRowExtended) String() string { return proto.CompactTextString(m) }
func (*FeatureRowExtended) ProtoMessage()    {}
func (*FeatureRowExtended) Descriptor() ([]byte, []int) {
	return fileDescriptor_7823aa2c72575793, []int{2}
}

func (m *FeatureRowExtended) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureRowExtended.Unmarshal(m, b)
}
func (m *FeatureRowExtended) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureRowExtended.Marshal(b, m, deterministic)
}
func (m *FeatureRowExtended) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureRowExtended.Merge(m, src)
}
func (m *FeatureRowExtended) XXX_Size() int {
	return xxx_messageInfo_FeatureRowExtended.Size(m)
}
func (m *FeatureRowExtended) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureRowExtended.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureRowExtended proto.InternalMessageInfo

func (m *FeatureRowExtended) GetRow() *FeatureRow {
	if m != nil {
		return m.Row
	}
	return nil
}

func (m *FeatureRowExtended) GetLastAttempt() *Attempt {
	if m != nil {
		return m.LastAttempt
	}
	return nil
}

func (m *FeatureRowExtended) GetFirstSeen() *timestamp.Timestamp {
	if m != nil {
		return m.FirstSeen
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "feast.types.Error")
	proto.RegisterType((*Attempt)(nil), "feast.types.Attempt")
	proto.RegisterType((*FeatureRowExtended)(nil), "feast.types.FeatureRowExtended")
}

func init() {
	proto.RegisterFile("feast/types/FeatureRowExtended.proto", fileDescriptor_7823aa2c72575793)
}

var fileDescriptor_7823aa2c72575793 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6b, 0xe2, 0x40,
	0x14, 0xc7, 0x71, 0xdd, 0xac, 0xeb, 0xcb, 0x9e, 0x06, 0xc1, 0x10, 0x04, 0x17, 0xd9, 0x83, 0x7b,
	0x99, 0x01, 0x0b, 0x2d, 0x3d, 0x56, 0xb0, 0xd7, 0x96, 0xd4, 0x53, 0x0f, 0x95, 0x31, 0xbe, 0xa4,
	0x56, 0x93, 0x09, 0x33, 0x2f, 0xb5, 0xfd, 0xbb, 0xfa, 0x0f, 0x96, 0xbc, 0x98, 0xaa, 0xd8, 0xdb,
	0xbc, 0xf7, 0xfd, 0xbc, 0x5f, 0xf3, 0x85, 0x7f, 0x09, 0x6a, 0x47, 0x8a, 0xde, 0x0b, 0x74, 0xea,
	0x16, 0x35, 0x95, 0x16, 0x23, 0xb3, 0x9b, 0xbd, 0x11, 0xe6, 0x2b, 0x5c, 0xc9, 0xc2, 0x1a, 0x32,
	0xc2, 0x67, 0x4a, 0x32, 0x15, 0x0e, 0x53, 0x63, 0xd2, 0x2d, 0x2a, 0x96, 0x96, 0x65, 0xa2, 0x68,
	0x9d, 0xa1, 0x23, 0x9d, 0x15, 0x35, 0x1d, 0x0e, 0xbe, 0xef, 0x59, 0xab, 0xa3, 0x57, 0xf0, 0x66,
	0xd6, 0x1a, 0x2b, 0x7a, 0xe0, 0xc5, 0xba, 0x74, 0x18, 0xb4, 0xfe, 0xb6, 0xc6, 0xdd, 0xa8, 0x0e,
	0xc4, 0x00, 0xba, 0x64, 0x75, 0xee, 0x12, 0x63, 0xb3, 0xe0, 0x07, 0x2b, 0x87, 0x84, 0x08, 0xa0,
	0x93, 0xa1, 0x73, 0x3a, 0xc5, 0xa0, 0xcd, 0x5a, 0x13, 0x8a, 0x21, 0xf8, 0x8e, 0x74, 0xbc, 0x59,
	0x90, 0xd5, 0x31, 0x06, 0x3f, 0x59, 0x05, 0x4e, 0xcd, 0xab, 0xcc, 0xe8, 0x0e, 0x3a, 0x37, 0x44,
	0x98, 0x15, 0x24, 0x42, 0xf8, 0xad, 0xeb, 0xa7, 0xe3, 0xe1, 0x5e, 0xf4, 0x15, 0x8b, 0x31, 0x78,
	0x58, 0xad, 0xc7, 0xb3, 0xfd, 0x89, 0x90, 0x47, 0xa7, 0x4b, 0x5e, 0x3c, 0xaa, 0x81, 0xd1, 0x47,
	0x0b, 0xc4, 0xf9, 0x8f, 0x89, 0xff, 0xd0, 0xb6, 0x66, 0xc7, 0x7d, 0xfd, 0x49, 0xff, 0xa4, 0xfc,
	0x40, 0x47, 0x15, 0x23, 0xae, 0xe0, 0xcf, 0x56, 0x3b, 0x5a, 0xec, 0x87, 0xef, 0x47, 0xf6, 0x4e,
	0x6a, 0xf6, 0x3b, 0x47, 0x7e, 0x45, 0x36, 0x07, 0x5c, 0x03, 0x24, 0x6b, 0xeb, 0x68, 0xe1, 0x10,
	0x73, 0xfe, 0x09, 0x7f, 0x12, 0xca, 0xda, 0x17, 0xd9, 0xf8, 0x22, 0xe7, 0x8d, 0x2f, 0x51, 0x97,
	0xe9, 0x07, 0xc4, 0x7c, 0xfa, 0x04, 0xc7, 0x66, 0x4e, 0xfb, 0xe7, 0x17, 0xdc, 0x57, 0xf5, 0x8f,
	0x97, 0xe9, 0x9a, 0x9e, 0xcb, 0xa5, 0x8c, 0x4d, 0xa6, 0x52, 0xf3, 0x82, 0x1b, 0x55, 0xbb, 0xca,
	0xdd, 0x9d, 0x4a, 0x31, 0x47, 0xab, 0x09, 0x57, 0x2a, 0x35, 0xea, 0xc8, 0xef, 0xe5, 0x2f, 0x06,
	0x2e, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xe4, 0x2a, 0x2f, 0x59, 0x02, 0x00, 0x00,
}
