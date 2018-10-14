// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/point.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

func init() { proto.RegisterFile("proto/point.proto", fileDescriptor_a92371c4ee379440) }

var fileDescriptor_a92371c4ee379440 = []byte{
	// 63 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x03, 0xb3, 0x85, 0x98, 0x0a, 0x92, 0x8c, 0xf8,
	0xb8, 0x78, 0x02, 0x40, 0x42, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x49, 0x6c, 0x60, 0x29,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x5a, 0xf0, 0xb5, 0x2f, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PointServiceClient is the client API for PointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PointServiceClient interface {
}

type pointServiceClient struct {
	cc *grpc.ClientConn
}

func NewPointServiceClient(cc *grpc.ClientConn) PointServiceClient {
	return &pointServiceClient{cc}
}

// PointServiceServer is the server API for PointService service.
type PointServiceServer interface {
}

func RegisterPointServiceServer(s *grpc.Server, srv PointServiceServer) {
	s.RegisterService(&_PointService_serviceDesc, srv)
}

var _PointService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PointService",
	HandlerType: (*PointServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "proto/point.proto",
}
