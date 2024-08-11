// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: stream.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Stream_GetStream_FullMethodName = "/com.databrew.gateway.stream.Stream/GetStream"
	Stream_AckOffset_FullMethodName = "/com.databrew.gateway.stream.Stream/AckOffset"
)

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamClient interface {
	GetStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Stream_GetStreamClient, error)
	AckOffset(ctx context.Context, in *AcceptCursor, opts ...grpc.CallOption) (*AcceptedCursor, error)
}

type streamClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) GetStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Stream_GetStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Stream_ServiceDesc.Streams[0], Stream_GetStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &streamGetStreamClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stream_GetStreamClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamGetStreamClient struct {
	grpc.ClientStream
}

func (x *streamGetStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamClient) AckOffset(ctx context.Context, in *AcceptCursor, opts ...grpc.CallOption) (*AcceptedCursor, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcceptedCursor)
	err := c.cc.Invoke(ctx, Stream_AckOffset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamServer is the server API for Stream service.
// All implementations must embed UnimplementedStreamServer
// for forward compatibility
type StreamServer interface {
	GetStream(*StreamRequest, Stream_GetStreamServer) error
	AckOffset(context.Context, *AcceptCursor) (*AcceptedCursor, error)
	mustEmbedUnimplementedStreamServer()
}

// UnimplementedStreamServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (UnimplementedStreamServer) GetStream(*StreamRequest, Stream_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedStreamServer) AckOffset(context.Context, *AcceptCursor) (*AcceptedCursor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckOffset not implemented")
}
func (UnimplementedStreamServer) mustEmbedUnimplementedStreamServer() {}

// UnsafeStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServer will
// result in compilation errors.
type UnsafeStreamServer interface {
	mustEmbedUnimplementedStreamServer()
}

func RegisterStreamServer(s grpc.ServiceRegistrar, srv StreamServer) {
	s.RegisterService(&Stream_ServiceDesc, srv)
}

func _Stream_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServer).GetStream(m, &streamGetStreamServer{ServerStream: stream})
}

type Stream_GetStreamServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamGetStreamServer struct {
	grpc.ServerStream
}

func (x *streamGetStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Stream_AckOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptCursor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServer).AckOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stream_AckOffset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServer).AckOffset(ctx, req.(*AcceptCursor))
	}
	return interceptor(ctx, in, info, handler)
}

// Stream_ServiceDesc is the grpc.ServiceDesc for Stream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.databrew.gateway.stream.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AckOffset",
			Handler:    _Stream_AckOffset_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Stream_GetStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}
