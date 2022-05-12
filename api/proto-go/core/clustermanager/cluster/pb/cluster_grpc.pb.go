// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: cluster.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	ListCluster(ctx context.Context, in *ListClusterRequest, opts ...grpc.CallOption) (*ListClusterResponse, error)
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*GetClusterResponse, error)
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*CreateClusterResponse, error)
	UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*UpdateClusterResponse, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*DeleteClusterResponse, error)
	PatchCluster(ctx context.Context, in *PatchClusterRequest, opts ...grpc.CallOption) (*PatchClusterResponse, error)
}

type clusterServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewClusterServiceClient(cc grpc1.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) ListCluster(ctx context.Context, in *ListClusterRequest, opts ...grpc.CallOption) (*ListClusterResponse, error) {
	out := new(ListClusterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.clustermanager.cluster.ClusterService/ListCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*GetClusterResponse, error) {
	out := new(GetClusterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.clustermanager.cluster.ClusterService/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*CreateClusterResponse, error) {
	out := new(CreateClusterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.clustermanager.cluster.ClusterService/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*UpdateClusterResponse, error) {
	out := new(UpdateClusterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.clustermanager.cluster.ClusterService/UpdateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*DeleteClusterResponse, error) {
	out := new(DeleteClusterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.clustermanager.cluster.ClusterService/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) PatchCluster(ctx context.Context, in *PatchClusterRequest, opts ...grpc.CallOption) (*PatchClusterResponse, error) {
	out := new(PatchClusterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.clustermanager.cluster.ClusterService/PatchCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	ListCluster(context.Context, *ListClusterRequest) (*ListClusterResponse, error)
	GetCluster(context.Context, *GetClusterRequest) (*GetClusterResponse, error)
	CreateCluster(context.Context, *CreateClusterRequest) (*CreateClusterResponse, error)
	UpdateCluster(context.Context, *UpdateClusterRequest) (*UpdateClusterResponse, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*DeleteClusterResponse, error)
	PatchCluster(context.Context, *PatchClusterRequest) (*PatchClusterResponse, error)
}

// UnimplementedClusterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (*UnimplementedClusterServiceServer) ListCluster(context.Context, *ListClusterRequest) (*ListClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCluster not implemented")
}
func (*UnimplementedClusterServiceServer) GetCluster(context.Context, *GetClusterRequest) (*GetClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (*UnimplementedClusterServiceServer) CreateCluster(context.Context, *CreateClusterRequest) (*CreateClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (*UnimplementedClusterServiceServer) UpdateCluster(context.Context, *UpdateClusterRequest) (*UpdateClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCluster not implemented")
}
func (*UnimplementedClusterServiceServer) DeleteCluster(context.Context, *DeleteClusterRequest) (*DeleteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (*UnimplementedClusterServiceServer) PatchCluster(context.Context, *PatchClusterRequest) (*PatchClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchCluster not implemented")
}

func RegisterClusterServiceServer(s grpc1.ServiceRegistrar, srv ClusterServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_ClusterService_serviceDesc(srv, opts...), srv)
}

var _ClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.clustermanager.cluster.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cluster.proto",
}

func _get_ClusterService_serviceDesc(srv ClusterServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_ClusterService_ListCluster_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListCluster(ctx, req.(*ListClusterRequest))
	}
	var _ClusterService_ListCluster_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ClusterService_ListCluster_info = transport.NewServiceInfo("erda.core.clustermanager.cluster.ClusterService", "ListCluster", srv)
		_ClusterService_ListCluster_Handler = h.Interceptor(_ClusterService_ListCluster_Handler)
	}

	_ClusterService_GetCluster_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetCluster(ctx, req.(*GetClusterRequest))
	}
	var _ClusterService_GetCluster_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ClusterService_GetCluster_info = transport.NewServiceInfo("erda.core.clustermanager.cluster.ClusterService", "GetCluster", srv)
		_ClusterService_GetCluster_Handler = h.Interceptor(_ClusterService_GetCluster_Handler)
	}

	_ClusterService_CreateCluster_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	var _ClusterService_CreateCluster_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ClusterService_CreateCluster_info = transport.NewServiceInfo("erda.core.clustermanager.cluster.ClusterService", "CreateCluster", srv)
		_ClusterService_CreateCluster_Handler = h.Interceptor(_ClusterService_CreateCluster_Handler)
	}

	_ClusterService_UpdateCluster_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateCluster(ctx, req.(*UpdateClusterRequest))
	}
	var _ClusterService_UpdateCluster_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ClusterService_UpdateCluster_info = transport.NewServiceInfo("erda.core.clustermanager.cluster.ClusterService", "UpdateCluster", srv)
		_ClusterService_UpdateCluster_Handler = h.Interceptor(_ClusterService_UpdateCluster_Handler)
	}

	_ClusterService_DeleteCluster_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	var _ClusterService_DeleteCluster_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ClusterService_DeleteCluster_info = transport.NewServiceInfo("erda.core.clustermanager.cluster.ClusterService", "DeleteCluster", srv)
		_ClusterService_DeleteCluster_Handler = h.Interceptor(_ClusterService_DeleteCluster_Handler)
	}

	_ClusterService_PatchCluster_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.PatchCluster(ctx, req.(*PatchClusterRequest))
	}
	var _ClusterService_PatchCluster_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ClusterService_PatchCluster_info = transport.NewServiceInfo("erda.core.clustermanager.cluster.ClusterService", "PatchCluster", srv)
		_ClusterService_PatchCluster_Handler = h.Interceptor(_ClusterService_PatchCluster_Handler)
	}

	var serviceDesc = _ClusterService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "ListCluster",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListClusterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ClusterServiceServer).ListCluster(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ClusterService_ListCluster_info)
				}
				if interceptor == nil {
					return _ClusterService_ListCluster_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.clustermanager.cluster.ClusterService/ListCluster",
				}
				return interceptor(ctx, in, info, _ClusterService_ListCluster_Handler)
			},
		},
		{
			MethodName: "GetCluster",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetClusterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ClusterServiceServer).GetCluster(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ClusterService_GetCluster_info)
				}
				if interceptor == nil {
					return _ClusterService_GetCluster_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.clustermanager.cluster.ClusterService/GetCluster",
				}
				return interceptor(ctx, in, info, _ClusterService_GetCluster_Handler)
			},
		},
		{
			MethodName: "CreateCluster",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateClusterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ClusterServiceServer).CreateCluster(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ClusterService_CreateCluster_info)
				}
				if interceptor == nil {
					return _ClusterService_CreateCluster_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.clustermanager.cluster.ClusterService/CreateCluster",
				}
				return interceptor(ctx, in, info, _ClusterService_CreateCluster_Handler)
			},
		},
		{
			MethodName: "UpdateCluster",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateClusterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ClusterServiceServer).UpdateCluster(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ClusterService_UpdateCluster_info)
				}
				if interceptor == nil {
					return _ClusterService_UpdateCluster_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.clustermanager.cluster.ClusterService/UpdateCluster",
				}
				return interceptor(ctx, in, info, _ClusterService_UpdateCluster_Handler)
			},
		},
		{
			MethodName: "DeleteCluster",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteClusterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ClusterServiceServer).DeleteCluster(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ClusterService_DeleteCluster_info)
				}
				if interceptor == nil {
					return _ClusterService_DeleteCluster_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.clustermanager.cluster.ClusterService/DeleteCluster",
				}
				return interceptor(ctx, in, info, _ClusterService_DeleteCluster_Handler)
			},
		},
		{
			MethodName: "PatchCluster",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(PatchClusterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ClusterServiceServer).PatchCluster(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ClusterService_PatchCluster_info)
				}
				if interceptor == nil {
					return _ClusterService_PatchCluster_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.clustermanager.cluster.ClusterService/PatchCluster",
				}
				return interceptor(ctx, in, info, _ClusterService_PatchCluster_Handler)
			},
		},
	}
	return &serviceDesc
}