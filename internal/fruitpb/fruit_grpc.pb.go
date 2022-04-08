// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fruitpb

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

// FruitServiceClient is the client API for FruitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FruitServiceClient interface {
	Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error)
}

type fruitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFruitServiceClient(cc grpc.ClientConnInterface) FruitServiceClient {
	return &fruitServiceClient{cc}
}

func (c *fruitServiceClient) Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error) {
	out := new(FilterResponse)
	err := c.cc.Invoke(ctx, "/fruit.FruitService/Filter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FruitServiceServer is the server API for FruitService service.
// All implementations must embed UnimplementedFruitServiceServer
// for forward compatibility
type FruitServiceServer interface {
	Filter(context.Context, *FilterRequest) (*FilterResponse, error)
	mustEmbedUnimplementedFruitServiceServer()
}

// UnimplementedFruitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFruitServiceServer struct {
}

func (UnimplementedFruitServiceServer) Filter(context.Context, *FilterRequest) (*FilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Filter not implemented")
}
func (UnimplementedFruitServiceServer) mustEmbedUnimplementedFruitServiceServer() {}

// UnsafeFruitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FruitServiceServer will
// result in compilation errors.
type UnsafeFruitServiceServer interface {
	mustEmbedUnimplementedFruitServiceServer()
}

func RegisterFruitServiceServer(s grpc.ServiceRegistrar, srv FruitServiceServer) {
	s.RegisterService(&FruitService_ServiceDesc, srv)
}

func _FruitService_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FruitServiceServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fruit.FruitService/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FruitServiceServer).Filter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FruitService_ServiceDesc is the grpc.ServiceDesc for FruitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FruitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fruit.FruitService",
	HandlerType: (*FruitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Filter",
			Handler:    _FruitService_Filter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/fruitpb/fruit.proto",
}