// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.3.0
// source: tunnel.proto

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

// TunnelClient is the client API for Tunnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TunnelClient interface {
	Auto(ctx context.Context, opts ...grpc.CallOption) (Tunnel_AutoClient, error)
	Http(ctx context.Context, opts ...grpc.CallOption) (Tunnel_HttpClient, error)
	Socks5(ctx context.Context, opts ...grpc.CallOption) (Tunnel_Socks5Client, error)
	V2RaySsr(ctx context.Context, opts ...grpc.CallOption) (Tunnel_V2RaySsrClient, error)
	V2RayVmess(ctx context.Context, opts ...grpc.CallOption) (Tunnel_V2RayVmessClient, error)
	V2RayVless(ctx context.Context, opts ...grpc.CallOption) (Tunnel_V2RayVlessClient, error)
	Health(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type tunnelClient struct {
	cc grpc.ClientConnInterface
}

func NewTunnelClient(cc grpc.ClientConnInterface) TunnelClient {
	return &tunnelClient{cc}
}

func (c *tunnelClient) Auto(ctx context.Context, opts ...grpc.CallOption) (Tunnel_AutoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tunnel_ServiceDesc.Streams[0], "/tunnel.Tunnel/Auto", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelAutoClient{stream}
	return x, nil
}

type Tunnel_AutoClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type tunnelAutoClient struct {
	grpc.ClientStream
}

func (x *tunnelAutoClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelAutoClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelClient) Http(ctx context.Context, opts ...grpc.CallOption) (Tunnel_HttpClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tunnel_ServiceDesc.Streams[1], "/tunnel.Tunnel/Http", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelHttpClient{stream}
	return x, nil
}

type Tunnel_HttpClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type tunnelHttpClient struct {
	grpc.ClientStream
}

func (x *tunnelHttpClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelHttpClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelClient) Socks5(ctx context.Context, opts ...grpc.CallOption) (Tunnel_Socks5Client, error) {
	stream, err := c.cc.NewStream(ctx, &Tunnel_ServiceDesc.Streams[2], "/tunnel.Tunnel/Socks5", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelSocks5Client{stream}
	return x, nil
}

type Tunnel_Socks5Client interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type tunnelSocks5Client struct {
	grpc.ClientStream
}

func (x *tunnelSocks5Client) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelSocks5Client) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelClient) V2RaySsr(ctx context.Context, opts ...grpc.CallOption) (Tunnel_V2RaySsrClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tunnel_ServiceDesc.Streams[3], "/tunnel.Tunnel/V2raySsr", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelV2RaySsrClient{stream}
	return x, nil
}

type Tunnel_V2RaySsrClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type tunnelV2RaySsrClient struct {
	grpc.ClientStream
}

func (x *tunnelV2RaySsrClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelV2RaySsrClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelClient) V2RayVmess(ctx context.Context, opts ...grpc.CallOption) (Tunnel_V2RayVmessClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tunnel_ServiceDesc.Streams[4], "/tunnel.Tunnel/V2rayVmess", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelV2RayVmessClient{stream}
	return x, nil
}

type Tunnel_V2RayVmessClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type tunnelV2RayVmessClient struct {
	grpc.ClientStream
}

func (x *tunnelV2RayVmessClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelV2RayVmessClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelClient) V2RayVless(ctx context.Context, opts ...grpc.CallOption) (Tunnel_V2RayVlessClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tunnel_ServiceDesc.Streams[5], "/tunnel.Tunnel/V2rayVless", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelV2RayVlessClient{stream}
	return x, nil
}

type Tunnel_V2RayVlessClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type tunnelV2RayVlessClient struct {
	grpc.ClientStream
}

func (x *tunnelV2RayVlessClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelV2RayVlessClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tunnelClient) Health(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/tunnel.Tunnel/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TunnelServer is the server API for Tunnel service.
// All implementations must embed UnimplementedTunnelServer
// for forward compatibility
type TunnelServer interface {
	Auto(Tunnel_AutoServer) error
	Http(Tunnel_HttpServer) error
	Socks5(Tunnel_Socks5Server) error
	V2RaySsr(Tunnel_V2RaySsrServer) error
	V2RayVmess(Tunnel_V2RayVmessServer) error
	V2RayVless(Tunnel_V2RayVlessServer) error
	Health(context.Context, *Ping) (*Pong, error)
	mustEmbedUnimplementedTunnelServer()
}

// UnimplementedTunnelServer must be embedded to have forward compatible implementations.
type UnimplementedTunnelServer struct {
}

func (UnimplementedTunnelServer) Auto(Tunnel_AutoServer) error {
	return status.Errorf(codes.Unimplemented, "method Auto not implemented")
}
func (UnimplementedTunnelServer) Http(Tunnel_HttpServer) error {
	return status.Errorf(codes.Unimplemented, "method Http not implemented")
}
func (UnimplementedTunnelServer) Socks5(Tunnel_Socks5Server) error {
	return status.Errorf(codes.Unimplemented, "method Socks5 not implemented")
}
func (UnimplementedTunnelServer) V2RaySsr(Tunnel_V2RaySsrServer) error {
	return status.Errorf(codes.Unimplemented, "method V2RaySsr not implemented")
}
func (UnimplementedTunnelServer) V2RayVmess(Tunnel_V2RayVmessServer) error {
	return status.Errorf(codes.Unimplemented, "method V2RayVmess not implemented")
}
func (UnimplementedTunnelServer) V2RayVless(Tunnel_V2RayVlessServer) error {
	return status.Errorf(codes.Unimplemented, "method V2RayVless not implemented")
}
func (UnimplementedTunnelServer) Health(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedTunnelServer) mustEmbedUnimplementedTunnelServer() {}

// UnsafeTunnelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TunnelServer will
// result in compilation errors.
type UnsafeTunnelServer interface {
	mustEmbedUnimplementedTunnelServer()
}

func RegisterTunnelServer(s grpc.ServiceRegistrar, srv TunnelServer) {
	s.RegisterService(&Tunnel_ServiceDesc, srv)
}

func _Tunnel_Auto_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).Auto(&tunnelAutoServer{stream})
}

type Tunnel_AutoServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type tunnelAutoServer struct {
	grpc.ServerStream
}

func (x *tunnelAutoServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelAutoServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunnel_Http_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).Http(&tunnelHttpServer{stream})
}

type Tunnel_HttpServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type tunnelHttpServer struct {
	grpc.ServerStream
}

func (x *tunnelHttpServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelHttpServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunnel_Socks5_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).Socks5(&tunnelSocks5Server{stream})
}

type Tunnel_Socks5Server interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type tunnelSocks5Server struct {
	grpc.ServerStream
}

func (x *tunnelSocks5Server) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelSocks5Server) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunnel_V2RaySsr_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).V2RaySsr(&tunnelV2RaySsrServer{stream})
}

type Tunnel_V2RaySsrServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type tunnelV2RaySsrServer struct {
	grpc.ServerStream
}

func (x *tunnelV2RaySsrServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelV2RaySsrServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunnel_V2RayVmess_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).V2RayVmess(&tunnelV2RayVmessServer{stream})
}

type Tunnel_V2RayVmessServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type tunnelV2RayVmessServer struct {
	grpc.ServerStream
}

func (x *tunnelV2RayVmessServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelV2RayVmessServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunnel_V2RayVless_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).V2RayVless(&tunnelV2RayVlessServer{stream})
}

type Tunnel_V2RayVlessServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type tunnelV2RayVlessServer struct {
	grpc.ServerStream
}

func (x *tunnelV2RayVlessServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelV2RayVlessServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Tunnel_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TunnelServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tunnel.Tunnel/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TunnelServer).Health(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

// Tunnel_ServiceDesc is the grpc.ServiceDesc for Tunnel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tunnel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tunnel.Tunnel",
	HandlerType: (*TunnelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Tunnel_Health_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Auto",
			Handler:       _Tunnel_Auto_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Http",
			Handler:       _Tunnel_Http_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Socks5",
			Handler:       _Tunnel_Socks5_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "V2raySsr",
			Handler:       _Tunnel_V2RaySsr_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "V2rayVmess",
			Handler:       _Tunnel_V2RayVmess_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "V2rayVless",
			Handler:       _Tunnel_V2RayVless_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tunnel.proto",
}
