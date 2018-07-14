// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idol.proto

package idol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetIdolMessage struct {
	TargetIdol           string   `protobuf:"bytes,1,opt,name=target_idol,json=targetIdol,proto3" json:"target_idol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIdolMessage) Reset()         { *m = GetIdolMessage{} }
func (m *GetIdolMessage) String() string { return proto.CompactTextString(m) }
func (*GetIdolMessage) ProtoMessage()    {}
func (*GetIdolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_idol_280fa9b33eb88087, []int{0}
}
func (m *GetIdolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIdolMessage.Unmarshal(m, b)
}
func (m *GetIdolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIdolMessage.Marshal(b, m, deterministic)
}
func (dst *GetIdolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIdolMessage.Merge(dst, src)
}
func (m *GetIdolMessage) XXX_Size() int {
	return xxx_messageInfo_GetIdolMessage.Size(m)
}
func (m *GetIdolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIdolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetIdolMessage proto.InternalMessageInfo

func (m *GetIdolMessage) GetTargetIdol() string {
	if m != nil {
		return m.TargetIdol
	}
	return ""
}

type IdolResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdolResponse) Reset()         { *m = IdolResponse{} }
func (m *IdolResponse) String() string { return proto.CompactTextString(m) }
func (*IdolResponse) ProtoMessage()    {}
func (*IdolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_idol_280fa9b33eb88087, []int{1}
}
func (m *IdolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdolResponse.Unmarshal(m, b)
}
func (m *IdolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdolResponse.Marshal(b, m, deterministic)
}
func (dst *IdolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdolResponse.Merge(dst, src)
}
func (m *IdolResponse) XXX_Size() int {
	return xxx_messageInfo_IdolResponse.Size(m)
}
func (m *IdolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdolResponse proto.InternalMessageInfo

func (m *IdolResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IdolResponse) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *IdolResponse) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*GetIdolMessage)(nil), "GetIdolMessage")
	proto.RegisterType((*IdolResponse)(nil), "IdolResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdolClient is the client API for Idol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdolClient interface {
	GetIdol(ctx context.Context, in *GetIdolMessage, opts ...grpc.CallOption) (*IdolResponse, error)
}

type idolClient struct {
	cc *grpc.ClientConn
}

func NewIdolClient(cc *grpc.ClientConn) IdolClient {
	return &idolClient{cc}
}

func (c *idolClient) GetIdol(ctx context.Context, in *GetIdolMessage, opts ...grpc.CallOption) (*IdolResponse, error) {
	out := new(IdolResponse)
	err := c.cc.Invoke(ctx, "/Idol/GetIdol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdolServer is the server API for Idol service.
type IdolServer interface {
	GetIdol(context.Context, *GetIdolMessage) (*IdolResponse, error)
}

func RegisterIdolServer(s *grpc.Server, srv IdolServer) {
	s.RegisterService(&_Idol_serviceDesc, srv)
}

func _Idol_GetIdol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdolMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdolServer).GetIdol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Idol/GetIdol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdolServer).GetIdol(ctx, req.(*GetIdolMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Idol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Idol",
	HandlerType: (*IdolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIdol",
			Handler:    _Idol_GetIdol_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idol.proto",
}

func init() { proto.RegisterFile("idol.proto", fileDescriptor_idol_280fa9b33eb88087) }

var fileDescriptor_idol_280fa9b33eb88087 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4c, 0xc9, 0xcf,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe4, 0xe2, 0x73, 0x4f, 0x2d, 0xf1, 0x4c, 0xc9,
	0xcf, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe7, 0xe2, 0x2e, 0x49, 0x2c, 0x4a,
	0x4f, 0x2d, 0x89, 0x07, 0x29, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x82, 0x08, 0x81,
	0xd4, 0x29, 0x85, 0x73, 0xf1, 0x80, 0xe8, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0xa8, 0x4a, 0x30, 0x5b, 0x48, 0x9a, 0x8b, 0x33, 0x2f,
	0x33, 0x39, 0x3b, 0x1e, 0x2c, 0xc1, 0x04, 0x96, 0xe0, 0x00, 0x09, 0xf8, 0x81, 0x24, 0xc5, 0xb8,
	0xd8, 0x32, 0x52, 0x33, 0xd3, 0x33, 0x4a, 0x24, 0x98, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c,
	0x23, 0x63, 0x2e, 0x16, 0x90, 0xc1, 0x42, 0xda, 0x5c, 0xec, 0x50, 0x37, 0x09, 0xf1, 0xeb, 0xa1,
	0xba, 0x4e, 0x8a, 0x57, 0x0f, 0xd9, 0x6e, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x3f, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0x58, 0x4c, 0x2b, 0xd5, 0x00, 0x00, 0x00,
}