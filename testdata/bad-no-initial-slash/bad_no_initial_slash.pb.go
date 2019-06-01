// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdata/bad-no-initial-slash/bad_no_initial_slash.proto

package validate_proto_http_testdata_bad_no_initial_slash

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BadNoInitialSlashServiceClient is the client API for BadNoInitialSlashService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BadNoInitialSlashServiceClient interface {
	Method(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type badNoInitialSlashServiceClient struct {
	cc *grpc.ClientConn
}

func NewBadNoInitialSlashServiceClient(cc *grpc.ClientConn) BadNoInitialSlashServiceClient {
	return &badNoInitialSlashServiceClient{cc}
}

func (c *badNoInitialSlashServiceClient) Method(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/validate_proto_http.testdata.bad_no_initial_slash.BadNoInitialSlashService/Method", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BadNoInitialSlashServiceServer is the server API for BadNoInitialSlashService service.
type BadNoInitialSlashServiceServer interface {
	Method(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterBadNoInitialSlashServiceServer(s *grpc.Server, srv BadNoInitialSlashServiceServer) {
	s.RegisterService(&_BadNoInitialSlashService_serviceDesc, srv)
}

func _BadNoInitialSlashService_Method_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadNoInitialSlashServiceServer).Method(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/validate_proto_http.testdata.bad_no_initial_slash.BadNoInitialSlashService/Method",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadNoInitialSlashServiceServer).Method(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BadNoInitialSlashService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "validate_proto_http.testdata.bad_no_initial_slash.BadNoInitialSlashService",
	HandlerType: (*BadNoInitialSlashServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method",
			Handler:    _BadNoInitialSlashService_Method_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdata/bad-no-initial-slash/bad_no_initial_slash.proto",
}

func init() {
	proto.RegisterFile("testdata/bad-no-initial-slash/bad_no_initial_slash.proto", fileDescriptor_bad_no_initial_slash_fca09031557555d7)
}

var fileDescriptor_bad_no_initial_slash_fca09031557555d7 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8d, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0xd1, 0xa1, 0x42, 0x71, 0xca, 0x20, 0x52, 0x9d, 0x3a, 0x0a, 0x4d, 0x50, 0x17, 0x71,
	0x14, 0x1c, 0x04, 0x75, 0xe9, 0x03, 0x84, 0x13, 0x13, 0x9b, 0x40, 0xcc, 0x09, 0xcd, 0x69, 0xc1,
	0xd5, 0x57, 0xf0, 0xd1, 0xee, 0x2b, 0xdc, 0x07, 0xb9, 0x34, 0x69, 0xb7, 0x7b, 0xc7, 0xf3, 0xfd,
	0xe7, 0xff, 0xfe, 0xfa, 0x89, 0x4c, 0x22, 0x0d, 0x04, 0x42, 0x81, 0xee, 0x02, 0x76, 0x2e, 0x38,
	0x72, 0xe0, 0xbb, 0xe4, 0x21, 0xd9, 0x05, 0xca, 0x80, 0x72, 0x85, 0x32, 0x43, 0x1e, 0x47, 0x24,
	0x64, 0xf7, 0x33, 0x78, 0xa7, 0x81, 0x8c, 0xcc, 0xb7, 0xb4, 0x44, 0x91, 0x6f, 0x36, 0x7e, 0xac,
	0xd8, 0xdc, 0x0c, 0x88, 0x83, 0x37, 0x22, 0x17, 0xd4, 0xf4, 0x2d, 0xcc, 0x4f, 0xa4, 0xdf, 0xe2,
	0x6b, 0x6e, 0xd7, 0x10, 0xa2, 0x13, 0x10, 0x02, 0x12, 0x90, 0xc3, 0x90, 0x4a, 0xfa, 0x60, 0xeb,
	0xeb, 0x17, 0xd0, 0x9f, 0xf8, 0x56, 0x84, 0xfd, 0xe2, 0xeb, 0xcd, 0x38, 0xbb, 0x2f, 0xc3, 0xde,
	0xeb, 0xea, 0xc3, 0x90, 0x45, 0xcd, 0xae, 0x78, 0x91, 0xf0, 0x6d, 0x81, 0xbf, 0x2e, 0x0b, 0xcd,
	0x09, 0xde, 0xb2, 0xbf, 0xdd, 0xfe, 0xff, 0xfc, 0xb2, 0xbd, 0x50, 0xa0, 0xc5, 0x34, 0xfa, 0xe7,
	0xb3, 0x3b, 0x55, 0xe5, 0x9f, 0xc7, 0x43, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x67, 0xcc, 0x37,
	0x1a, 0x01, 0x00, 0x00,
}
