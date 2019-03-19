// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mergedprofile.proto

package spotify

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MergedProfileRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MergedProfileRequest) Reset()         { *m = MergedProfileRequest{} }
func (m *MergedProfileRequest) String() string { return proto.CompactTextString(m) }
func (*MergedProfileRequest) ProtoMessage()    {}
func (*MergedProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_936d088d10e2a1e0, []int{0}
}

func (m *MergedProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MergedProfileRequest.Unmarshal(m, b)
}
func (m *MergedProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MergedProfileRequest.Marshal(b, m, deterministic)
}
func (m *MergedProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergedProfileRequest.Merge(m, src)
}
func (m *MergedProfileRequest) XXX_Size() int {
	return xxx_messageInfo_MergedProfileRequest.Size(m)
}
func (m *MergedProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MergedProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MergedProfileRequest proto.InternalMessageInfo

type MergedProfileReply struct {
	Username             *string  `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Artistid             *string  `protobuf:"bytes,2,opt,name=artistid" json:"artistid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MergedProfileReply) Reset()         { *m = MergedProfileReply{} }
func (m *MergedProfileReply) String() string { return proto.CompactTextString(m) }
func (*MergedProfileReply) ProtoMessage()    {}
func (*MergedProfileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_936d088d10e2a1e0, []int{1}
}

func (m *MergedProfileReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MergedProfileReply.Unmarshal(m, b)
}
func (m *MergedProfileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MergedProfileReply.Marshal(b, m, deterministic)
}
func (m *MergedProfileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergedProfileReply.Merge(m, src)
}
func (m *MergedProfileReply) XXX_Size() int {
	return xxx_messageInfo_MergedProfileReply.Size(m)
}
func (m *MergedProfileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MergedProfileReply.DiscardUnknown(m)
}

var xxx_messageInfo_MergedProfileReply proto.InternalMessageInfo

func (m *MergedProfileReply) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *MergedProfileReply) GetArtistid() string {
	if m != nil && m.Artistid != nil {
		return *m.Artistid
	}
	return ""
}

func init() {
	proto.RegisterType((*MergedProfileRequest)(nil), "spotify.MergedProfileRequest")
	proto.RegisterType((*MergedProfileReply)(nil), "spotify.MergedProfileReply")
}

func init() { proto.RegisterFile("mergedprofile.proto", fileDescriptor_936d088d10e2a1e0) }

var fileDescriptor_936d088d10e2a1e0 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x4d, 0x2d, 0x4a,
	0x4f, 0x4d, 0x29, 0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2f, 0x2e, 0xc8, 0x2f, 0xc9, 0x4c, 0xab, 0x54, 0x12, 0xe3, 0x12, 0xf1, 0x05, 0xcb, 0x07,
	0x40, 0xe4, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x7c, 0xb8, 0x84, 0xd0, 0xc4, 0x0b,
	0x72, 0x2a, 0x85, 0xa4, 0xb8, 0x38, 0x4a, 0x8b, 0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x90, 0x5c, 0x62, 0x51, 0x49, 0x66, 0x71, 0x49, 0x66,
	0x8a, 0x04, 0x13, 0x44, 0x0e, 0xc6, 0x07, 0x04, 0x00, 0x00, 0xff, 0xff, 0xec, 0xc6, 0xe6, 0x57,
	0x84, 0x00, 0x00, 0x00,
}
