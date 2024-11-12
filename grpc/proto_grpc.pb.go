// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: grpc/proto.proto

package proto

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
	BiddingService_PlaceBid_FullMethodName  = "/BiddingService/PlaceBid"
	BiddingService_GetResult_FullMethodName = "/BiddingService/GetResult"
)

// BiddingServiceClient is the client API for BiddingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiddingServiceClient interface {
	PlaceBid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error)
	GetResult(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error)
}

type biddingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiddingServiceClient(cc grpc.ClientConnInterface) BiddingServiceClient {
	return &biddingServiceClient{cc}
}

func (c *biddingServiceClient) PlaceBid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidResponse)
	err := c.cc.Invoke(ctx, BiddingService_PlaceBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *biddingServiceClient) GetResult(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, BiddingService_GetResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingServiceServer is the server API for BiddingService service.
// All implementations must embed UnimplementedBiddingServiceServer
// for forward compatibility.
type BiddingServiceServer interface {
	PlaceBid(context.Context, *BidRequest) (*BidResponse, error)
	GetResult(context.Context, *ResultRequest) (*ResultResponse, error)
	mustEmbedUnimplementedBiddingServiceServer()
}

// UnimplementedBiddingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBiddingServiceServer struct{}

func (UnimplementedBiddingServiceServer) PlaceBid(context.Context, *BidRequest) (*BidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBid not implemented")
}
func (UnimplementedBiddingServiceServer) GetResult(context.Context, *ResultRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedBiddingServiceServer) mustEmbedUnimplementedBiddingServiceServer() {}
func (UnimplementedBiddingServiceServer) testEmbeddedByValue()                        {}

// UnsafeBiddingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiddingServiceServer will
// result in compilation errors.
type UnsafeBiddingServiceServer interface {
	mustEmbedUnimplementedBiddingServiceServer()
}

func RegisterBiddingServiceServer(s grpc.ServiceRegistrar, srv BiddingServiceServer) {
	// If the following call pancis, it indicates UnimplementedBiddingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BiddingService_ServiceDesc, srv)
}

func _BiddingService_PlaceBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingServiceServer).PlaceBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BiddingService_PlaceBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingServiceServer).PlaceBid(ctx, req.(*BidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BiddingService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BiddingService_GetResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingServiceServer).GetResult(ctx, req.(*ResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BiddingService_ServiceDesc is the grpc.ServiceDesc for BiddingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BiddingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BiddingService",
	HandlerType: (*BiddingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBid",
			Handler:    _BiddingService_PlaceBid_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _BiddingService_GetResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
