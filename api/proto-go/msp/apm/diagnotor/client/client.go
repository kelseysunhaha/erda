// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: apm_diagnotor.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/diagnotor/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// DiagnotorService apm_diagnotor.proto
	DiagnotorService() pb.DiagnotorServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		diagnotorService: pb.NewDiagnotorServiceClient(cc),
	}
}

type serviceClients struct {
	diagnotorService pb.DiagnotorServiceClient
}

func (c *serviceClients) DiagnotorService() pb.DiagnotorServiceClient {
	return c.diagnotorService
}

type diagnotorServiceWrapper struct {
	client pb.DiagnotorServiceClient
	opts   []grpc1.CallOption
}

func (s *diagnotorServiceWrapper) ListServices(ctx context.Context, req *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	return s.client.ListServices(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *diagnotorServiceWrapper) StartDiagnosis(ctx context.Context, req *pb.StartDiagnosisRequest) (*pb.StartDiagnosisResponse, error) {
	return s.client.StartDiagnosis(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *diagnotorServiceWrapper) QueryDiagnosisStatus(ctx context.Context, req *pb.QueryDiagnosisStatusRequest) (*pb.QueryDiagnosisStatusResponse, error) {
	return s.client.QueryDiagnosisStatus(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *diagnotorServiceWrapper) StopDiagnosis(ctx context.Context, req *pb.StopDiagnosisRequest) (*pb.StopDiagnosisResponse, error) {
	return s.client.StopDiagnosis(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *diagnotorServiceWrapper) ListProcesses(ctx context.Context, req *pb.ListProcessesRequest) (*pb.ListProcessesResponse, error) {
	return s.client.ListProcesses(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}