// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	ListBeer(ctx context.Context, in *ListBeerReq, opts ...grpc.CallOption) (*ListBeerReply, error)
	CreateBeer(ctx context.Context, in *CreateBeerReq, opts ...grpc.CallOption) (*CreateBeerReply, error)
	GetBeer(ctx context.Context, in *GetBeerReq, opts ...grpc.CallOption) (*GetBeerReply, error)
	DeleteBeer(ctx context.Context, in *DeleteBeerReq, opts ...grpc.CallOption) (*DeleteBeerReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) ListBeer(ctx context.Context, in *ListBeerReq, opts ...grpc.CallOption) (*ListBeerReply, error) {
	out := new(ListBeerReply)
	err := c.cc.Invoke(ctx, "/cart.service.v1.Order/ListBeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateBeer(ctx context.Context, in *CreateBeerReq, opts ...grpc.CallOption) (*CreateBeerReply, error) {
	out := new(CreateBeerReply)
	err := c.cc.Invoke(ctx, "/cart.service.v1.Order/CreateBeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetBeer(ctx context.Context, in *GetBeerReq, opts ...grpc.CallOption) (*GetBeerReply, error) {
	out := new(GetBeerReply)
	err := c.cc.Invoke(ctx, "/cart.service.v1.Order/GetBeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteBeer(ctx context.Context, in *DeleteBeerReq, opts ...grpc.CallOption) (*DeleteBeerReply, error) {
	out := new(DeleteBeerReply)
	err := c.cc.Invoke(ctx, "/cart.service.v1.Order/DeleteBeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	ListBeer(context.Context, *ListBeerReq) (*ListBeerReply, error)
	CreateBeer(context.Context, *CreateBeerReq) (*CreateBeerReply, error)
	GetBeer(context.Context, *GetBeerReq) (*GetBeerReply, error)
	DeleteBeer(context.Context, *DeleteBeerReq) (*DeleteBeerReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) ListBeer(context.Context, *ListBeerReq) (*ListBeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBeer not implemented")
}
func (UnimplementedOrderServer) CreateBeer(context.Context, *CreateBeerReq) (*CreateBeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBeer not implemented")
}
func (UnimplementedOrderServer) GetBeer(context.Context, *GetBeerReq) (*GetBeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeer not implemented")
}
func (UnimplementedOrderServer) DeleteBeer(context.Context, *DeleteBeerReq) (*DeleteBeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBeer not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_ListBeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ListBeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.service.v1.Order/ListBeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ListBeer(ctx, req.(*ListBeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateBeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateBeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.service.v1.Order/CreateBeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateBeer(ctx, req.(*CreateBeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetBeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetBeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.service.v1.Order/GetBeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetBeer(ctx, req.(*GetBeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteBeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBeerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteBeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.service.v1.Order/DeleteBeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteBeer(ctx, req.(*DeleteBeerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.service.v1.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBeer",
			Handler:    _Order_ListBeer_Handler,
		},
		{
			MethodName: "CreateBeer",
			Handler:    _Order_CreateBeer_Handler,
		},
		{
			MethodName: "GetBeer",
			Handler:    _Order_GetBeer_Handler,
		},
		{
			MethodName: "DeleteBeer",
			Handler:    _Order_DeleteBeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/order.proto",
}
