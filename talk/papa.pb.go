// Code generated by protoc-gen-go. DO NOT EDIT.
// source: papa.proto

package talk

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

type PapaTask struct {
	Id                   uint64   `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PapaTask) Reset()         { *m = PapaTask{} }
func (m *PapaTask) String() string { return proto.CompactTextString(m) }
func (*PapaTask) ProtoMessage()    {}
func (*PapaTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_papa_2d821163f4bfd65f, []int{0}
}
func (m *PapaTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PapaTask.Unmarshal(m, b)
}
func (m *PapaTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PapaTask.Marshal(b, m, deterministic)
}
func (dst *PapaTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PapaTask.Merge(dst, src)
}
func (m *PapaTask) XXX_Size() int {
	return xxx_messageInfo_PapaTask.Size(m)
}
func (m *PapaTask) XXX_DiscardUnknown() {
	xxx_messageInfo_PapaTask.DiscardUnknown(m)
}

var xxx_messageInfo_PapaTask proto.InternalMessageInfo

func (m *PapaTask) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PapaTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*PapaTask)(nil), "talk.PapaTask")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PapaClient is the client API for Papa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PapaClient interface {
	Pull(ctx context.Context, opts ...grpc.CallOption) (Papa_PullClient, error)
}

type papaClient struct {
	cc *grpc.ClientConn
}

func NewPapaClient(cc *grpc.ClientConn) PapaClient {
	return &papaClient{cc}
}

func (c *papaClient) Pull(ctx context.Context, opts ...grpc.CallOption) (Papa_PullClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Papa_serviceDesc.Streams[0], "/talk.Papa/Pull", opts...)
	if err != nil {
		return nil, err
	}
	x := &papaPullClient{stream}
	return x, nil
}

type Papa_PullClient interface {
	Send(*PapaTask) error
	Recv() (*PapaTask, error)
	grpc.ClientStream
}

type papaPullClient struct {
	grpc.ClientStream
}

func (x *papaPullClient) Send(m *PapaTask) error {
	return x.ClientStream.SendMsg(m)
}

func (x *papaPullClient) Recv() (*PapaTask, error) {
	m := new(PapaTask)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PapaServer is the server API for Papa service.
type PapaServer interface {
	Pull(Papa_PullServer) error
}

func RegisterPapaServer(s *grpc.Server, srv PapaServer) {
	s.RegisterService(&_Papa_serviceDesc, srv)
}

func _Papa_Pull_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PapaServer).Pull(&papaPullServer{stream})
}

type Papa_PullServer interface {
	Send(*PapaTask) error
	Recv() (*PapaTask, error)
	grpc.ServerStream
}

type papaPullServer struct {
	grpc.ServerStream
}

func (x *papaPullServer) Send(m *PapaTask) error {
	return x.ServerStream.SendMsg(m)
}

func (x *papaPullServer) Recv() (*PapaTask, error) {
	m := new(PapaTask)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Papa_serviceDesc = grpc.ServiceDesc{
	ServiceName: "talk.Papa",
	HandlerType: (*PapaServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pull",
			Handler:       _Papa_Pull_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "papa.proto",
}

func init() { proto.RegisterFile("papa.proto", fileDescriptor_papa_2d821163f4bfd65f) }

var fileDescriptor_papa_2d821163f4bfd65f = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0x2c, 0x48,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0xcc, 0xc9, 0x56, 0xd2, 0xe3, 0xe2,
	0x08, 0x48, 0x2c, 0x48, 0x0c, 0x49, 0x2c, 0xce, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0b, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x8d, 0x8c, 0xb8, 0x58, 0x40, 0xea, 0x85,
	0xb4, 0xb8, 0x58, 0x02, 0x4a, 0x73, 0x72, 0x84, 0xf8, 0xf4, 0x40, 0xc6, 0xe8, 0xc1, 0xcc, 0x90,
	0x42, 0xe3, 0x6b, 0x30, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0x2d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x10, 0x49, 0xe4, 0xf3, 0x7e, 0x00, 0x00, 0x00,
}
