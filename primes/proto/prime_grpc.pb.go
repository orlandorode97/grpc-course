// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: prime.proto

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

// PrimeServiceClient is the client API for PrimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrimeServiceClient interface {
	DescomposeNumber(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_DescomposeNumberClient, error)
}

type primeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeServiceClient(cc grpc.ClientConnInterface) PrimeServiceClient {
	return &primeServiceClient{cc}
}

func (c *primeServiceClient) DescomposeNumber(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_DescomposeNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrimeService_ServiceDesc.Streams[0], "/primes.PrimeService/DescomposeNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeServiceDescomposeNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeService_DescomposeNumberClient interface {
	Recv() (*PrimeRespose, error)
	grpc.ClientStream
}

type primeServiceDescomposeNumberClient struct {
	grpc.ClientStream
}

func (x *primeServiceDescomposeNumberClient) Recv() (*PrimeRespose, error) {
	m := new(PrimeRespose)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeServiceServer is the server API for PrimeService service.
// All implementations must embed UnimplementedPrimeServiceServer
// for forward compatibility
type PrimeServiceServer interface {
	DescomposeNumber(*PrimeRequest, PrimeService_DescomposeNumberServer) error
	mustEmbedUnimplementedPrimeServiceServer()
}

// UnimplementedPrimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrimeServiceServer struct {
}

func (UnimplementedPrimeServiceServer) DescomposeNumber(*PrimeRequest, PrimeService_DescomposeNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method DescomposeNumber not implemented")
}
func (UnimplementedPrimeServiceServer) mustEmbedUnimplementedPrimeServiceServer() {}

// UnsafePrimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimeServiceServer will
// result in compilation errors.
type UnsafePrimeServiceServer interface {
	mustEmbedUnimplementedPrimeServiceServer()
}

func RegisterPrimeServiceServer(s grpc.ServiceRegistrar, srv PrimeServiceServer) {
	s.RegisterService(&PrimeService_ServiceDesc, srv)
}

func _PrimeService_DescomposeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeServiceServer).DescomposeNumber(m, &primeServiceDescomposeNumberServer{stream})
}

type PrimeService_DescomposeNumberServer interface {
	Send(*PrimeRespose) error
	grpc.ServerStream
}

type primeServiceDescomposeNumberServer struct {
	grpc.ServerStream
}

func (x *primeServiceDescomposeNumberServer) Send(m *PrimeRespose) error {
	return x.ServerStream.SendMsg(m)
}

// PrimeService_ServiceDesc is the grpc.ServiceDesc for PrimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "primes.PrimeService",
	HandlerType: (*PrimeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DescomposeNumber",
			Handler:       _PrimeService_DescomposeNumber_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "prime.proto",
}
