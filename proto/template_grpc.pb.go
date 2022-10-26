// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateClient interface {
	Join(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Lamport, error)
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Template_SubscribeClient, error)
	Send(ctx context.Context, in *Lamport, opts ...grpc.CallOption) (*Empty, error)
}

type templateClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateClient(cc grpc.ClientConnInterface) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) Join(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Lamport, error) {
	out := new(Lamport)
	err := c.cc.Invoke(ctx, "/proto.Template/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Template_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Template_ServiceDesc.Streams[0], "/proto.Template/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &templateSubscribeClient{stream}
	return x, nil
}

type Template_SubscribeClient interface {
	Send(*Lamport) error
	Recv() (*Lamport, error)
	grpc.ClientStream
}

type templateSubscribeClient struct {
	grpc.ClientStream
}

func (x *templateSubscribeClient) Send(m *Lamport) error {
	return x.ClientStream.SendMsg(m)
}

func (x *templateSubscribeClient) Recv() (*Lamport, error) {
	m := new(Lamport)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *templateClient) Send(ctx context.Context, in *Lamport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Template/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServer is the server API for Template service.
// All implementations must embed UnimplementedTemplateServer
// for forward compatibility
type TemplateServer interface {
	Join(context.Context, *Empty) (*Lamport, error)
	Subscribe(Template_SubscribeServer) error
	Send(context.Context, *Lamport) (*Empty, error)
	mustEmbedUnimplementedTemplateServer()
}

// UnimplementedTemplateServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateServer struct {
}

func (UnimplementedTemplateServer) Join(context.Context, *Empty) (*Lamport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedTemplateServer) Subscribe(Template_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedTemplateServer) Send(context.Context, *Lamport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedTemplateServer) mustEmbedUnimplementedTemplateServer() {}

// UnsafeTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServer will
// result in compilation errors.
type UnsafeTemplateServer interface {
	mustEmbedUnimplementedTemplateServer()
}

func RegisterTemplateServer(s grpc.ServiceRegistrar, srv TemplateServer) {
	s.RegisterService(&Template_ServiceDesc, srv)
}

func _Template_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Join(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TemplateServer).Subscribe(&templateSubscribeServer{stream})
}

type Template_SubscribeServer interface {
	Send(*Lamport) error
	Recv() (*Lamport, error)
	grpc.ServerStream
}

type templateSubscribeServer struct {
	grpc.ServerStream
}

func (x *templateSubscribeServer) Send(m *Lamport) error {
	return x.ServerStream.SendMsg(m)
}

func (x *templateSubscribeServer) Recv() (*Lamport, error) {
	m := new(Lamport)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Template_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lamport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Send(ctx, req.(*Lamport))
	}
	return interceptor(ctx, in, info, handler)
}

// Template_ServiceDesc is the grpc.ServiceDesc for Template service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Template_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Template_Join_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Template_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Template_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/template.proto",
}
