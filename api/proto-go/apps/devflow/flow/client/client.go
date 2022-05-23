// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: flow.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/apps/devflow/flow/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// FlowService flow.proto
	FlowService() pb.FlowServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		flowService: pb.NewFlowServiceClient(cc),
	}
}

type serviceClients struct {
	flowService pb.FlowServiceClient
}

func (c *serviceClients) FlowService() pb.FlowServiceClient {
	return c.flowService
}

type flowServiceWrapper struct {
	client pb.FlowServiceClient
	opts   []grpc1.CallOption
}

func (s *flowServiceWrapper) CreateFlowNode(ctx context.Context, req *pb.CreateFlowNodeRequest) (*pb.CreateFlowNodeResponse, error) {
	return s.client.CreateFlowNode(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *flowServiceWrapper) OperationDeploy(ctx context.Context, req *pb.OperationDeployRequest) (*pb.OperationDeployResponse, error) {
	return s.client.OperationDeploy(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *flowServiceWrapper) DeleteFlowNode(ctx context.Context, req *pb.DeleteFlowNodeRequest) (*pb.DeleteFlowNodeResponse, error) {
	return s.client.DeleteFlowNode(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *flowServiceWrapper) Reconstruction(ctx context.Context, req *pb.ReconstructionRequest) (*pb.ReconstructionResponse, error) {
	return s.client.Reconstruction(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *flowServiceWrapper) GetDevFlowInfo(ctx context.Context, req *pb.GetDevFlowInfoRequest) (*pb.GetDevFlowInfoResponse, error) {
	return s.client.GetDevFlowInfo(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}