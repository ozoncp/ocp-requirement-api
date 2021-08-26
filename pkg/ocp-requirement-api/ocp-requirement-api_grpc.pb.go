// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_requirement_api

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

// OcpRequirementApiClient is the client API for OcpRequirementApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpRequirementApiClient interface {
	ListRequirementsV1(ctx context.Context, in *ListRequirementsV1Request, opts ...grpc.CallOption) (*ListRequirementsV1Response, error)
	DescribeRequirementV1(ctx context.Context, in *DescribeRequirementV1Request, opts ...grpc.CallOption) (*DescribeRequirementV1Response, error)
	CreateRequirementV1(ctx context.Context, in *CreateRequirementV1Request, opts ...grpc.CallOption) (*CreateRequirementV1Response, error)
	MultiCreateRequirementV1(ctx context.Context, in *MultiCreateRequirementV1Request, opts ...grpc.CallOption) (*MultiCreateRequirementV1Response, error)
	UpdateRequirementV1(ctx context.Context, in *UpdateRequirementV1Request, opts ...grpc.CallOption) (*UpdateRequirementV1Response, error)
	RemoveRequirementV1(ctx context.Context, in *RemoveRequirementV1Request, opts ...grpc.CallOption) (*RemoveRequirementV1Response, error)
}

type ocpRequirementApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpRequirementApiClient(cc grpc.ClientConnInterface) OcpRequirementApiClient {
	return &ocpRequirementApiClient{cc}
}

func (c *ocpRequirementApiClient) ListRequirementsV1(ctx context.Context, in *ListRequirementsV1Request, opts ...grpc.CallOption) (*ListRequirementsV1Response, error) {
	out := new(ListRequirementsV1Response)
	err := c.cc.Invoke(ctx, "/ocp.requirement.api.OcpRequirementApi/ListRequirementsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpRequirementApiClient) DescribeRequirementV1(ctx context.Context, in *DescribeRequirementV1Request, opts ...grpc.CallOption) (*DescribeRequirementV1Response, error) {
	out := new(DescribeRequirementV1Response)
	err := c.cc.Invoke(ctx, "/ocp.requirement.api.OcpRequirementApi/DescribeRequirementV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpRequirementApiClient) CreateRequirementV1(ctx context.Context, in *CreateRequirementV1Request, opts ...grpc.CallOption) (*CreateRequirementV1Response, error) {
	out := new(CreateRequirementV1Response)
	err := c.cc.Invoke(ctx, "/ocp.requirement.api.OcpRequirementApi/CreateRequirementV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpRequirementApiClient) MultiCreateRequirementV1(ctx context.Context, in *MultiCreateRequirementV1Request, opts ...grpc.CallOption) (*MultiCreateRequirementV1Response, error) {
	out := new(MultiCreateRequirementV1Response)
	err := c.cc.Invoke(ctx, "/ocp.requirement.api.OcpRequirementApi/MultiCreateRequirementV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpRequirementApiClient) UpdateRequirementV1(ctx context.Context, in *UpdateRequirementV1Request, opts ...grpc.CallOption) (*UpdateRequirementV1Response, error) {
	out := new(UpdateRequirementV1Response)
	err := c.cc.Invoke(ctx, "/ocp.requirement.api.OcpRequirementApi/UpdateRequirementV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpRequirementApiClient) RemoveRequirementV1(ctx context.Context, in *RemoveRequirementV1Request, opts ...grpc.CallOption) (*RemoveRequirementV1Response, error) {
	out := new(RemoveRequirementV1Response)
	err := c.cc.Invoke(ctx, "/ocp.requirement.api.OcpRequirementApi/RemoveRequirementV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpRequirementApiServer is the server API for OcpRequirementApi service.
// All implementations must embed UnimplementedOcpRequirementApiServer
// for forward compatibility
type OcpRequirementApiServer interface {
	ListRequirementsV1(context.Context, *ListRequirementsV1Request) (*ListRequirementsV1Response, error)
	DescribeRequirementV1(context.Context, *DescribeRequirementV1Request) (*DescribeRequirementV1Response, error)
	CreateRequirementV1(context.Context, *CreateRequirementV1Request) (*CreateRequirementV1Response, error)
	MultiCreateRequirementV1(context.Context, *MultiCreateRequirementV1Request) (*MultiCreateRequirementV1Response, error)
	UpdateRequirementV1(context.Context, *UpdateRequirementV1Request) (*UpdateRequirementV1Response, error)
	RemoveRequirementV1(context.Context, *RemoveRequirementV1Request) (*RemoveRequirementV1Response, error)
	mustEmbedUnimplementedOcpRequirementApiServer()
}

// UnimplementedOcpRequirementApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpRequirementApiServer struct {
}

func (UnimplementedOcpRequirementApiServer) ListRequirementsV1(context.Context, *ListRequirementsV1Request) (*ListRequirementsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRequirementsV1 not implemented")
}
func (UnimplementedOcpRequirementApiServer) DescribeRequirementV1(context.Context, *DescribeRequirementV1Request) (*DescribeRequirementV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRequirementV1 not implemented")
}
func (UnimplementedOcpRequirementApiServer) CreateRequirementV1(context.Context, *CreateRequirementV1Request) (*CreateRequirementV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequirementV1 not implemented")
}
func (UnimplementedOcpRequirementApiServer) MultiCreateRequirementV1(context.Context, *MultiCreateRequirementV1Request) (*MultiCreateRequirementV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateRequirementV1 not implemented")
}
func (UnimplementedOcpRequirementApiServer) UpdateRequirementV1(context.Context, *UpdateRequirementV1Request) (*UpdateRequirementV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRequirementV1 not implemented")
}
func (UnimplementedOcpRequirementApiServer) RemoveRequirementV1(context.Context, *RemoveRequirementV1Request) (*RemoveRequirementV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRequirementV1 not implemented")
}
func (UnimplementedOcpRequirementApiServer) mustEmbedUnimplementedOcpRequirementApiServer() {}

// UnsafeOcpRequirementApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpRequirementApiServer will
// result in compilation errors.
type UnsafeOcpRequirementApiServer interface {
	mustEmbedUnimplementedOcpRequirementApiServer()
}

func RegisterOcpRequirementApiServer(s grpc.ServiceRegistrar, srv OcpRequirementApiServer) {
	s.RegisterService(&OcpRequirementApi_ServiceDesc, srv)
}

func _OcpRequirementApi_ListRequirementsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequirementsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpRequirementApiServer).ListRequirementsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.requirement.api.OcpRequirementApi/ListRequirementsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpRequirementApiServer).ListRequirementsV1(ctx, req.(*ListRequirementsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpRequirementApi_DescribeRequirementV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequirementV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpRequirementApiServer).DescribeRequirementV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.requirement.api.OcpRequirementApi/DescribeRequirementV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpRequirementApiServer).DescribeRequirementV1(ctx, req.(*DescribeRequirementV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpRequirementApi_CreateRequirementV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequirementV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpRequirementApiServer).CreateRequirementV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.requirement.api.OcpRequirementApi/CreateRequirementV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpRequirementApiServer).CreateRequirementV1(ctx, req.(*CreateRequirementV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpRequirementApi_MultiCreateRequirementV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateRequirementV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpRequirementApiServer).MultiCreateRequirementV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.requirement.api.OcpRequirementApi/MultiCreateRequirementV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpRequirementApiServer).MultiCreateRequirementV1(ctx, req.(*MultiCreateRequirementV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpRequirementApi_UpdateRequirementV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequirementV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpRequirementApiServer).UpdateRequirementV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.requirement.api.OcpRequirementApi/UpdateRequirementV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpRequirementApiServer).UpdateRequirementV1(ctx, req.(*UpdateRequirementV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpRequirementApi_RemoveRequirementV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequirementV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpRequirementApiServer).RemoveRequirementV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.requirement.api.OcpRequirementApi/RemoveRequirementV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpRequirementApiServer).RemoveRequirementV1(ctx, req.(*RemoveRequirementV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpRequirementApi_ServiceDesc is the grpc.ServiceDesc for OcpRequirementApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpRequirementApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.requirement.api.OcpRequirementApi",
	HandlerType: (*OcpRequirementApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRequirementsV1",
			Handler:    _OcpRequirementApi_ListRequirementsV1_Handler,
		},
		{
			MethodName: "DescribeRequirementV1",
			Handler:    _OcpRequirementApi_DescribeRequirementV1_Handler,
		},
		{
			MethodName: "CreateRequirementV1",
			Handler:    _OcpRequirementApi_CreateRequirementV1_Handler,
		},
		{
			MethodName: "MultiCreateRequirementV1",
			Handler:    _OcpRequirementApi_MultiCreateRequirementV1_Handler,
		},
		{
			MethodName: "UpdateRequirementV1",
			Handler:    _OcpRequirementApi_UpdateRequirementV1_Handler,
		},
		{
			MethodName: "RemoveRequirementV1",
			Handler:    _OcpRequirementApi_RemoveRequirementV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-requirement-api/ocp-requirement-api.proto",
}
