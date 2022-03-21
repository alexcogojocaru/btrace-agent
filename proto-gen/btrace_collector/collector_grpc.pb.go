// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: collector.proto

package btrace_collector

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CollectorServiceClient is the client API for CollectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectorServiceClient interface {
	StreamSpan(ctx context.Context, in *SpanC, opts ...grpc.CallOption) (*ResponseC, error)
}

type collectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorServiceClient(cc grpc.ClientConnInterface) CollectorServiceClient {
	return &collectorServiceClient{cc}
}

func (c *collectorServiceClient) StreamSpan(ctx context.Context, in *SpanC, opts ...grpc.CallOption) (*ResponseC, error) {
	out := new(ResponseC)
	err := c.cc.Invoke(ctx, "/CollectorService/StreamSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorServiceServer is the server API for CollectorService service.
// All implementations must embed UnimplementedCollectorServiceServer
// for forward compatibility
type CollectorServiceServer interface {
	StreamSpan(context.Context, *SpanC) (*ResponseC, error)
	mustEmbedUnimplementedCollectorServiceServer()
}

// UnimplementedCollectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollectorServiceServer struct {
}

func (UnimplementedCollectorServiceServer) StreamSpan(context.Context, *SpanC) (*ResponseC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamSpan not implemented")
}
func (UnimplementedCollectorServiceServer) mustEmbedUnimplementedCollectorServiceServer() {}

// UnsafeCollectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectorServiceServer will
// result in compilation errors.
type UnsafeCollectorServiceServer interface {
	mustEmbedUnimplementedCollectorServiceServer()
}

func RegisterCollectorServiceServer(s grpc.ServiceRegistrar, srv CollectorServiceServer) {
	s.RegisterService(&CollectorService_ServiceDesc, srv)
}

func _CollectorService_StreamSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServiceServer).StreamSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CollectorService/StreamSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServiceServer).StreamSpan(ctx, req.(*SpanC))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectorService_ServiceDesc is the grpc.ServiceDesc for CollectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CollectorService",
	HandlerType: (*CollectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamSpan",
			Handler:    _CollectorService_StreamSpan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector.proto",
}
