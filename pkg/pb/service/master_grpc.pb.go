// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterServiceClient interface {
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetCard(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCardResponse, error)
	GetCharacter(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCharacterResponse, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/service.MasterService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) GetCard(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCardResponse, error) {
	out := new(GetCardResponse)
	err := c.cc.Invoke(ctx, "/service.MasterService/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) GetCharacter(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCharacterResponse, error) {
	out := new(GetCharacterResponse)
	err := c.cc.Invoke(ctx, "/service.MasterService/GetCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
// All implementations should embed UnimplementedMasterServiceServer
// for forward compatibility
type MasterServiceServer interface {
	GetAll(context.Context, *Empty) (*GetAllResponse, error)
	GetCard(context.Context, *Empty) (*GetCardResponse, error)
	GetCharacter(context.Context, *Empty) (*GetCharacterResponse, error)
}

// UnimplementedMasterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMasterServiceServer struct {
}

func (UnimplementedMasterServiceServer) GetAll(context.Context, *Empty) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedMasterServiceServer) GetCard(context.Context, *Empty) (*GetCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (UnimplementedMasterServiceServer) GetCharacter(context.Context, *Empty) (*GetCharacterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacter not implemented")
}

// UnsafeMasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServiceServer will
// result in compilation errors.
type UnsafeMasterServiceServer interface {
	mustEmbedUnimplementedMasterServiceServer()
}

func RegisterMasterServiceServer(s grpc.ServiceRegistrar, srv MasterServiceServer) {
	s.RegisterService(&MasterService_ServiceDesc, srv)
}

func _MasterService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MasterService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MasterService/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetCard(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_GetCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MasterService/GetCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetCharacter(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterService_ServiceDesc is the grpc.ServiceDesc for MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _MasterService_GetAll_Handler,
		},
		{
			MethodName: "GetCard",
			Handler:    _MasterService_GetCard_Handler,
		},
		{
			MethodName: "GetCharacter",
			Handler:    _MasterService_GetCharacter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/master.proto",
}