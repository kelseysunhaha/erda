// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erda-project/erda-proto-go/core/dicehub/release/pb (interfaces: ReleaseServiceServer)

// Package mock_pb is a generated GoMock package.
package deployment_order

import (
	context "context"
	reflect "reflect"

	pb "github.com/erda-project/erda-proto-go/core/dicehub/release/pb"
	gomock "github.com/golang/mock/gomock"
)

// MockReleaseServiceServer is a mock of ReleaseServiceServer interface.
type MockReleaseServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockReleaseServiceServerMockRecorder
}

// MockReleaseServiceServerMockRecorder is the mock recorder for MockReleaseServiceServer.
type MockReleaseServiceServerMockRecorder struct {
	mock *MockReleaseServiceServer
}

// NewMockReleaseServiceServer creates a new mock instance.
func NewMockReleaseServiceServer(ctrl *gomock.Controller) *MockReleaseServiceServer {
	mock := &MockReleaseServiceServer{ctrl: ctrl}
	mock.recorder = &MockReleaseServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReleaseServiceServer) EXPECT() *MockReleaseServiceServerMockRecorder {
	return m.recorder
}

// CheckVersion mocks base method.
func (m *MockReleaseServiceServer) CheckVersion(arg0 context.Context, arg1 *pb.CheckVersionRequest) (*pb.CheckVersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVersion", arg0, arg1)
	ret0, _ := ret[0].(*pb.CheckVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVersion indicates an expected call of CheckVersion.
func (mr *MockReleaseServiceServerMockRecorder) CheckVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVersion", reflect.TypeOf((*MockReleaseServiceServer)(nil).CheckVersion), arg0, arg1)
}

// CreateRelease mocks base method.
func (m *MockReleaseServiceServer) CreateRelease(arg0 context.Context, arg1 *pb.ReleaseCreateRequest) (*pb.ReleaseCreateResponseData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseCreateResponseData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRelease indicates an expected call of CreateRelease.
func (mr *MockReleaseServiceServerMockRecorder) CreateRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).CreateRelease), arg0, arg1)
}

// DeleteRelease mocks base method.
func (m *MockReleaseServiceServer) DeleteRelease(arg0 context.Context, arg1 *pb.ReleaseDeleteRequest) (*pb.ReleaseDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRelease indicates an expected call of DeleteRelease.
func (mr *MockReleaseServiceServerMockRecorder) DeleteRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).DeleteRelease), arg0, arg1)
}

// DeleteReleases mocks base method.
func (m *MockReleaseServiceServer) DeleteReleases(arg0 context.Context, arg1 *pb.ReleasesDeleteRequest) (*pb.ReleasesDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReleases", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleasesDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReleases indicates an expected call of DeleteReleases.
func (mr *MockReleaseServiceServerMockRecorder) DeleteReleases(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReleases", reflect.TypeOf((*MockReleaseServiceServer)(nil).DeleteReleases), arg0, arg1)
}

// GetIosPlist mocks base method.
func (m *MockReleaseServiceServer) GetIosPlist(arg0 context.Context, arg1 *pb.GetIosPlistRequest) (*pb.GetIosPlistResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIosPlist", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetIosPlistResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIosPlist indicates an expected call of GetIosPlist.
func (mr *MockReleaseServiceServerMockRecorder) GetIosPlist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIosPlist", reflect.TypeOf((*MockReleaseServiceServer)(nil).GetIosPlist), arg0, arg1)
}

// GetLatestReleases mocks base method.
func (m *MockReleaseServiceServer) GetLatestReleases(arg0 context.Context, arg1 *pb.GetLatestReleasesRequest) (*pb.GetLatestReleasesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestReleases", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetLatestReleasesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestReleases indicates an expected call of GetLatestReleases.
func (mr *MockReleaseServiceServerMockRecorder) GetLatestReleases(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestReleases", reflect.TypeOf((*MockReleaseServiceServer)(nil).GetLatestReleases), arg0, arg1)
}

// GetRelease mocks base method.
func (m *MockReleaseServiceServer) GetRelease(arg0 context.Context, arg1 *pb.ReleaseGetRequest) (*pb.ReleaseGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelease indicates an expected call of GetRelease.
func (mr *MockReleaseServiceServerMockRecorder) GetRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).GetRelease), arg0, arg1)
}

// ListRelease mocks base method.
func (m *MockReleaseServiceServer) ListRelease(arg0 context.Context, arg1 *pb.ReleaseListRequest) (*pb.ReleaseListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRelease indicates an expected call of ListRelease.
func (mr *MockReleaseServiceServerMockRecorder) ListRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).ListRelease), arg0, arg1)
}

// ListReleaseName mocks base method.
func (m *MockReleaseServiceServer) ListReleaseName(arg0 context.Context, arg1 *pb.ListReleaseNameRequest) (*pb.ListReleaseNameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleaseName", arg0, arg1)
	ret0, _ := ret[0].(*pb.ListReleaseNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleaseName indicates an expected call of ListReleaseName.
func (mr *MockReleaseServiceServerMockRecorder) ListReleaseName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleaseName", reflect.TypeOf((*MockReleaseServiceServer)(nil).ListReleaseName), arg0, arg1)
}

// ParseReleaseFile mocks base method.
func (m *MockReleaseServiceServer) ParseReleaseFile(arg0 context.Context, arg1 *pb.ParseReleaseFileRequest) (*pb.ParseReleaseFileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseReleaseFile", arg0, arg1)
	ret0, _ := ret[0].(*pb.ParseReleaseFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseReleaseFile indicates an expected call of ParseReleaseFile.
func (mr *MockReleaseServiceServerMockRecorder) ParseReleaseFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseReleaseFile", reflect.TypeOf((*MockReleaseServiceServer)(nil).ParseReleaseFile), arg0, arg1)
}

// PutOffRelease mocks base method.
func (m *MockReleaseServiceServer) PutOffRelease(arg0 context.Context, arg1 *pb.ReleasePutOffRequest) (*pb.ReleasePutOffResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutOffRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleasePutOffResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutOffRelease indicates an expected call of PutOffRelease.
func (mr *MockReleaseServiceServerMockRecorder) PutOffRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutOffRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).PutOffRelease), arg0, arg1)
}

// PutOnRelease mocks base method.
func (m *MockReleaseServiceServer) PutOnRelease(arg0 context.Context, arg1 *pb.ReleasePutOnRequest) (*pb.ReleasePutOnResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutOnRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleasePutOnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutOnRelease indicates an expected call of PutOnRelease.
func (mr *MockReleaseServiceServerMockRecorder) PutOnRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutOnRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).PutOnRelease), arg0, arg1)
}

// ReleaseGC mocks base method.
func (m *MockReleaseServiceServer) ReleaseGC(arg0 context.Context, arg1 *pb.ReleaseGCRequest) (*pb.ReleaseDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseGC", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseGC indicates an expected call of ReleaseGC.
func (mr *MockReleaseServiceServerMockRecorder) ReleaseGC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseGC", reflect.TypeOf((*MockReleaseServiceServer)(nil).ReleaseGC), arg0, arg1)
}

// ToFormalRelease mocks base method.
func (m *MockReleaseServiceServer) ToFormalRelease(arg0 context.Context, arg1 *pb.FormalReleaseRequest) (*pb.FormalReleaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToFormalRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.FormalReleaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToFormalRelease indicates an expected call of ToFormalRelease.
func (mr *MockReleaseServiceServerMockRecorder) ToFormalRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToFormalRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).ToFormalRelease), arg0, arg1)
}

// ToFormalReleases mocks base method.
func (m *MockReleaseServiceServer) ToFormalReleases(arg0 context.Context, arg1 *pb.FormalReleasesRequest) (*pb.FormalReleasesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToFormalReleases", arg0, arg1)
	ret0, _ := ret[0].(*pb.FormalReleasesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToFormalReleases indicates an expected call of ToFormalReleases.
func (mr *MockReleaseServiceServerMockRecorder) ToFormalReleases(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToFormalReleases", reflect.TypeOf((*MockReleaseServiceServer)(nil).ToFormalReleases), arg0, arg1)
}

// UpdateRelease mocks base method.
func (m *MockReleaseServiceServer) UpdateRelease(arg0 context.Context, arg1 *pb.ReleaseUpdateRequest) (*pb.ReleaseUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRelease indicates an expected call of UpdateRelease.
func (mr *MockReleaseServiceServerMockRecorder) UpdateRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).UpdateRelease), arg0, arg1)
}

// UpdateReleaseReference mocks base method.
func (m *MockReleaseServiceServer) UpdateReleaseReference(arg0 context.Context, arg1 *pb.ReleaseReferenceUpdateRequest) (*pb.ReleaseDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReleaseReference", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReleaseReference indicates an expected call of UpdateReleaseReference.
func (mr *MockReleaseServiceServerMockRecorder) UpdateReleaseReference(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReleaseReference", reflect.TypeOf((*MockReleaseServiceServer)(nil).UpdateReleaseReference), arg0, arg1)
}

// UploadRelease mocks base method.
func (m *MockReleaseServiceServer) UploadRelease(arg0 context.Context, arg1 *pb.ReleaseUploadRequest) (*pb.ReleaseUploadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadRelease", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReleaseUploadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadRelease indicates an expected call of UploadRelease.
func (mr *MockReleaseServiceServerMockRecorder) UploadRelease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadRelease", reflect.TypeOf((*MockReleaseServiceServer)(nil).UploadRelease), arg0, arg1)
}