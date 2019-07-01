// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/types/status.proto

package types // import "github.com/fever365/kratos/pkg/ecode/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status struct {
	// The error code see ecode.Code
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details              []*any.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_88668d6b2bf80f08, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "bilibili.rpc.Status")
}

func init() { proto.RegisterFile("internal/types/status.proto", fileDescriptor_status_88668d6b2bf80f08) }

var fileDescriptor_status_88668d6b2bf80f08 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0xd9, 0x5b, 0xbd, 0xc3, 0x9c, 0x85, 0x04, 0x8b, 0x55, 0x9b, 0xc5, 0x6a, 0x0b, 0x4d,
	0x40, 0x4b, 0x2b, 0xcf, 0x17, 0x58, 0x22, 0x36, 0x76, 0x49, 0x6e, 0x2e, 0x04, 0x92, 0xcc, 0x92,
	0xe4, 0x8a, 0xbc, 0x8e, 0x4f, 0x2a, 0x9b, 0x65, 0x41, 0x8b, 0x19, 0x66, 0x98, 0xff, 0xe7, 0xfb,
	0x87, 0x3c, 0xd8, 0x90, 0x21, 0x06, 0xe9, 0x78, 0x2e, 0x13, 0x24, 0x9e, 0xb2, 0xcc, 0xe7, 0xc4,
	0xa6, 0x88, 0x19, 0xe9, 0xb5, 0xb2, 0xce, 0xce, 0xc5, 0xe2, 0xa4, 0xef, 0xef, 0x0c, 0xa2, 0x71,
	0xc0, 0xeb, 0x4d, 0x9d, 0x4f, 0x5c, 0x86, 0xb2, 0x08, 0x1f, 0x4f, 0x64, 0xfb, 0x59, 0x8d, 0x94,
	0x92, 0x0b, 0x8d, 0x47, 0xe8, 0x9a, 0xbe, 0x19, 0x2e, 0x45, 0x9d, 0x69, 0x47, 0x76, 0x1e, 0x52,
	0x92, 0x06, 0xba, 0x4d, 0xdf, 0x0c, 0x57, 0x62, 0x5d, 0x29, 0x23, 0xbb, 0x23, 0x64, 0x69, 0x5d,
	0xea, 0xda, 0xbe, 0x1d, 0xf6, 0x2f, 0xb7, 0x6c, 0x81, 0xb0, 0x15, 0xc2, 0xde, 0x43, 0x11, 0xab,
	0xe8, 0xf0, 0x45, 0x6e, 0x34, 0x7a, 0xf6, 0x37, 0xd6, 0x61, 0xbf, 0x90, 0xc7, 0xd9, 0x30, 0x36,
	0xdf, 0x4f, 0x06, 0x9f, 0x35, 0x7a, 0x8f, 0x81, 0x3b, 0xab, 0xa2, 0x8c, 0x85, 0xc3, 0x9c, 0x82,
	0xff, 0x7f, 0xf4, 0xad, 0xf6, 0x9f, 0x4d, 0x2b, 0xc6, 0x0f, 0xb5, 0xad, 0xb4, 0xd7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xa3, 0xc1, 0x82, 0x0d, 0x01, 0x00, 0x00,
}
