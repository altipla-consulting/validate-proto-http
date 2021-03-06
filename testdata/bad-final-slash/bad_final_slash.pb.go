// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdata/bad-final-slash/bad_final_slash.proto

package validate_proto_http_testdata_bad_final_slash

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

// BadFinalSlashClient is the client API for BadFinalSlash service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BadFinalSlashClient interface {
	Method(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type badFinalSlashClient struct {
	cc *grpc.ClientConn
}

func NewBadFinalSlashClient(cc *grpc.ClientConn) BadFinalSlashClient {
	return &badFinalSlashClient{cc}
}

func (c *badFinalSlashClient) Method(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/validate_proto_http.testdata.bad_final_slash.BadFinalSlash/Method", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BadFinalSlashServer is the server API for BadFinalSlash service.
type BadFinalSlashServer interface {
	Method(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterBadFinalSlashServer(s *grpc.Server, srv BadFinalSlashServer) {
	s.RegisterService(&_BadFinalSlash_serviceDesc, srv)
}

func _BadFinalSlash_Method_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadFinalSlashServer).Method(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/validate_proto_http.testdata.bad_final_slash.BadFinalSlash/Method",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadFinalSlashServer).Method(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BadFinalSlash_serviceDesc = grpc.ServiceDesc{
	ServiceName: "validate_proto_http.testdata.bad_final_slash.BadFinalSlash",
	HandlerType: (*BadFinalSlashServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method",
			Handler:    _BadFinalSlash_Method_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdata/bad-final-slash/bad_final_slash.proto",
}

func init() {
	proto.RegisterFile("testdata/bad-final-slash/bad_final_slash.proto", fileDescriptor_bad_final_slash_8b2b5480c366e74d)
}

var fileDescriptor_bad_final_slash_8b2b5480c366e74d = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x49, 0x2d, 0x2e,
	0x49, 0x49, 0x2c, 0x49, 0xd4, 0x4f, 0x4a, 0x4c, 0xd1, 0x4d, 0xcb, 0xcc, 0x4b, 0xcc, 0xd1, 0x2d,
	0xce, 0x49, 0x2c, 0xce, 0x00, 0xf1, 0xe3, 0xc1, 0xfc, 0x78, 0x30, 0x5f, 0xaf, 0xa0, 0x28, 0xbf,
	0x24, 0x5f, 0x48, 0xa7, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x35, 0x1e, 0xcc, 0x8f, 0xcf,
	0x28, 0x29, 0x29, 0x80, 0x9b, 0xa1, 0x87, 0xa6, 0x47, 0x4a, 0x3a, 0x3d, 0x3f, 0x3f, 0x3d, 0x27,
	0x55, 0x1f, 0xac, 0x36, 0xa9, 0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0x62, 0x94, 0x94,
	0x0c, 0x54, 0x32, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f,
	0xaf, 0x18, 0x22, 0x6b, 0x14, 0xcf, 0xc5, 0xeb, 0x94, 0x98, 0xe2, 0x06, 0x32, 0x2c, 0x18, 0x64,
	0x96, 0x90, 0x1f, 0x17, 0x9b, 0x6f, 0x6a, 0x49, 0x46, 0x7e, 0x8a, 0x90, 0x98, 0x1e, 0x44, 0xa7,
	0x1e, 0xcc, 0x58, 0x3d, 0x57, 0x90, 0xb1, 0x52, 0x38, 0xc4, 0x95, 0x44, 0x9a, 0x2e, 0x3f, 0x99,
	0xcc, 0xc4, 0xa7, 0xc4, 0x09, 0xf2, 0x93, 0x7e, 0x69, 0x51, 0x8e, 0xbe, 0x15, 0xa3, 0x56, 0x12,
	0x1b, 0x58, 0x95, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x5b, 0x57, 0x18, 0x02, 0x01, 0x00,
	0x00,
}
