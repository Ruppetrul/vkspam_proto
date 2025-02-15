// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: parser.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Parser_ParsePublic_FullMethodName = "/parser.Parser/ParsePublic"
)

// ParserClient is the client API for Parser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParserClient interface {
	ParsePublic(ctx context.Context, in *ParsePublicRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ParseProgress], error)
}

type parserClient struct {
	cc grpc.ClientConnInterface
}

func NewParserClient(cc grpc.ClientConnInterface) ParserClient {
	return &parserClient{cc}
}

func (c *parserClient) ParsePublic(ctx context.Context, in *ParsePublicRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ParseProgress], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Parser_ServiceDesc.Streams[0], Parser_ParsePublic_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ParsePublicRequest, ParseProgress]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Parser_ParsePublicClient = grpc.ServerStreamingClient[ParseProgress]

// ParserServer is the server API for Parser service.
// All implementations must embed UnimplementedParserServer
// for forward compatibility.
type ParserServer interface {
	ParsePublic(*ParsePublicRequest, grpc.ServerStreamingServer[ParseProgress]) error
	mustEmbedUnimplementedParserServer()
}

// UnimplementedParserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParserServer struct{}

func (UnimplementedParserServer) ParsePublic(*ParsePublicRequest, grpc.ServerStreamingServer[ParseProgress]) error {
	return status.Errorf(codes.Unimplemented, "method ParsePublic not implemented")
}
func (UnimplementedParserServer) mustEmbedUnimplementedParserServer() {}
func (UnimplementedParserServer) testEmbeddedByValue()                {}

// UnsafeParserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParserServer will
// result in compilation errors.
type UnsafeParserServer interface {
	mustEmbedUnimplementedParserServer()
}

func RegisterParserServer(s grpc.ServiceRegistrar, srv ParserServer) {
	// If the following call pancis, it indicates UnimplementedParserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Parser_ServiceDesc, srv)
}

func _Parser_ParsePublic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ParsePublicRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ParserServer).ParsePublic(m, &grpc.GenericServerStream[ParsePublicRequest, ParseProgress]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Parser_ParsePublicServer = grpc.ServerStreamingServer[ParseProgress]

// Parser_ServiceDesc is the grpc.ServiceDesc for Parser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parser.Parser",
	HandlerType: (*ParserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ParsePublic",
			Handler:       _Parser_ParsePublic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "parser.proto",
}
