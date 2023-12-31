// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StateScemServiceClient is the client API for StateScemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateScemServiceClient interface {
	DeployWorkflow(ctx context.Context, in *DeployWorkflowRequest, opts ...grpc.CallOption) (*DeployWorkflowlResponse, error)
	CreateWorkflowInstance(ctx context.Context, in *CreateWorkflowInstanceRequest, opts ...grpc.CallOption) (*CreateWorkflowInstanceResponse, error)
}

type stateScemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStateScemServiceClient(cc grpc.ClientConnInterface) StateScemServiceClient {
	return &stateScemServiceClient{cc}
}

func (c *stateScemServiceClient) DeployWorkflow(ctx context.Context, in *DeployWorkflowRequest, opts ...grpc.CallOption) (*DeployWorkflowlResponse, error) {
	out := new(DeployWorkflowlResponse)
	err := c.cc.Invoke(ctx, "/pb.StateScemService/DeployWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateScemServiceClient) CreateWorkflowInstance(ctx context.Context, in *CreateWorkflowInstanceRequest, opts ...grpc.CallOption) (*CreateWorkflowInstanceResponse, error) {
	out := new(CreateWorkflowInstanceResponse)
	err := c.cc.Invoke(ctx, "/pb.StateScemService/CreateWorkflowInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StateScemServiceServer is the server API for StateScemService service.
// All implementations must embed UnimplementedStateScemServiceServer
// for forward compatibility
type StateScemServiceServer interface {
	DeployWorkflow(context.Context, *DeployWorkflowRequest) (*DeployWorkflowlResponse, error)
	CreateWorkflowInstance(context.Context, *CreateWorkflowInstanceRequest) (*CreateWorkflowInstanceResponse, error)
	mustEmbedUnimplementedStateScemServiceServer()
}

// UnimplementedStateScemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStateScemServiceServer struct {
}

func (UnimplementedStateScemServiceServer) DeployWorkflow(context.Context, *DeployWorkflowRequest) (*DeployWorkflowlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployWorkflow not implemented")
}
func (UnimplementedStateScemServiceServer) CreateWorkflowInstance(context.Context, *CreateWorkflowInstanceRequest) (*CreateWorkflowInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflowInstance not implemented")
}
func (UnimplementedStateScemServiceServer) mustEmbedUnimplementedStateScemServiceServer() {}

// UnsafeStateScemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateScemServiceServer will
// result in compilation errors.
type UnsafeStateScemServiceServer interface {
	mustEmbedUnimplementedStateScemServiceServer()
}

func RegisterStateScemServiceServer(s grpc.ServiceRegistrar, srv StateScemServiceServer) {
	s.RegisterService(&_StateScemService_serviceDesc, srv)
}

func _StateScemService_DeployWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateScemServiceServer).DeployWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StateScemService/DeployWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateScemServiceServer).DeployWorkflow(ctx, req.(*DeployWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateScemService_CreateWorkflowInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateScemServiceServer).CreateWorkflowInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StateScemService/CreateWorkflowInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateScemServiceServer).CreateWorkflowInstance(ctx, req.(*CreateWorkflowInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StateScemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StateScemService",
	HandlerType: (*StateScemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployWorkflow",
			Handler:    _StateScemService_DeployWorkflow_Handler,
		},
		{
			MethodName: "CreateWorkflowInstance",
			Handler:    _StateScemService_CreateWorkflowInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto_source/state_scem.proto",
}
