// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/amineadminterraform/go-app/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/amineadminterraform/go-app/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateArgoWorkflow mocks base method.
func (m *MockStore) CreateArgoWorkflow(arg0 context.Context, arg1 db.CreateArgoWorkflowParams) (db.ArgoWorkflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArgoWorkflow", arg0, arg1)
	ret0, _ := ret[0].(db.ArgoWorkflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArgoWorkflow indicates an expected call of CreateArgoWorkflow.
func (mr *MockStoreMockRecorder) CreateArgoWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArgoWorkflow", reflect.TypeOf((*MockStore)(nil).CreateArgoWorkflow), arg0, arg1)
}

// CreateEnvLayer mocks base method.
func (m *MockStore) CreateEnvLayer(arg0 context.Context, arg1 db.CreateEnvLayerParams) (db.EnvLayer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnvLayer", arg0, arg1)
	ret0, _ := ret[0].(db.EnvLayer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEnvLayer indicates an expected call of CreateEnvLayer.
func (mr *MockStoreMockRecorder) CreateEnvLayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvLayer", reflect.TypeOf((*MockStore)(nil).CreateEnvLayer), arg0, arg1)
}

// CreateProcess mocks base method.
func (m *MockStore) CreateProcess(arg0 context.Context, arg1 db.CreateProcessParams) (db.Process, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProcess", arg0, arg1)
	ret0, _ := ret[0].(db.Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProcess indicates an expected call of CreateProcess.
func (mr *MockStoreMockRecorder) CreateProcess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProcess", reflect.TypeOf((*MockStore)(nil).CreateProcess), arg0, arg1)
}

// CreateProject mocks base method.
func (m *MockStore) CreateProject(arg0 context.Context, arg1 db.CreateProjectParams) (db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockStoreMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockStore)(nil).CreateProject), arg0, arg1)
}

// CreateProjectEnvironment mocks base method.
func (m *MockStore) CreateProjectEnvironment(arg0 context.Context, arg1 db.CreateProjectEnvironmentParams) (db.ProjectEnvironment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectEnvironment", arg0, arg1)
	ret0, _ := ret[0].(db.ProjectEnvironment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProjectEnvironment indicates an expected call of CreateProjectEnvironment.
func (mr *MockStoreMockRecorder) CreateProjectEnvironment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectEnvironment", reflect.TypeOf((*MockStore)(nil).CreateProjectEnvironment), arg0, arg1)
}

// CreateRequest mocks base method.
func (m *MockStore) CreateRequest(arg0 context.Context, arg1 db.CreateRequestParams) (db.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRequest", arg0, arg1)
	ret0, _ := ret[0].(db.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRequest indicates an expected call of CreateRequest.
func (mr *MockStoreMockRecorder) CreateRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRequest", reflect.TypeOf((*MockStore)(nil).CreateRequest), arg0, arg1)
}

// CreateTemplate mocks base method.
func (m *MockStore) CreateTemplate(arg0 context.Context, arg1 db.CreateTemplateParams) (db.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemplate", arg0, arg1)
	ret0, _ := ret[0].(db.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTemplate indicates an expected call of CreateTemplate.
func (mr *MockStoreMockRecorder) CreateTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemplate", reflect.TypeOf((*MockStore)(nil).CreateTemplate), arg0, arg1)
}

// DeleteArgoWorkflow mocks base method.
func (m *MockStore) DeleteArgoWorkflow(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArgoWorkflow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArgoWorkflow indicates an expected call of DeleteArgoWorkflow.
func (mr *MockStoreMockRecorder) DeleteArgoWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArgoWorkflow", reflect.TypeOf((*MockStore)(nil).DeleteArgoWorkflow), arg0, arg1)
}

// DeleteEnvLayer mocks base method.
func (m *MockStore) DeleteEnvLayer(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnvLayer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvLayer indicates an expected call of DeleteEnvLayer.
func (mr *MockStoreMockRecorder) DeleteEnvLayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvLayer", reflect.TypeOf((*MockStore)(nil).DeleteEnvLayer), arg0, arg1)
}

// DeleteProcess mocks base method.
func (m *MockStore) DeleteProcess(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProcess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProcess indicates an expected call of DeleteProcess.
func (mr *MockStoreMockRecorder) DeleteProcess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProcess", reflect.TypeOf((*MockStore)(nil).DeleteProcess), arg0, arg1)
}

// DeleteProject mocks base method.
func (m *MockStore) DeleteProject(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockStoreMockRecorder) DeleteProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockStore)(nil).DeleteProject), arg0, arg1)
}

// DeleteProjectEnvironment mocks base method.
func (m *MockStore) DeleteProjectEnvironment(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectEnvironment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProjectEnvironment indicates an expected call of DeleteProjectEnvironment.
func (mr *MockStoreMockRecorder) DeleteProjectEnvironment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectEnvironment", reflect.TypeOf((*MockStore)(nil).DeleteProjectEnvironment), arg0, arg1)
}

// DeleteRequest mocks base method.
func (m *MockStore) DeleteRequest(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRequest indicates an expected call of DeleteRequest.
func (mr *MockStoreMockRecorder) DeleteRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRequest", reflect.TypeOf((*MockStore)(nil).DeleteRequest), arg0, arg1)
}

// DeleteTemplate mocks base method.
func (m *MockStore) DeleteTemplate(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTemplate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTemplate indicates an expected call of DeleteTemplate.
func (mr *MockStoreMockRecorder) DeleteTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTemplate", reflect.TypeOf((*MockStore)(nil).DeleteTemplate), arg0, arg1)
}

// GetArgoWorkflow mocks base method.
func (m *MockStore) GetArgoWorkflow(arg0 context.Context, arg1 int64) (db.ArgoWorkflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArgoWorkflow", arg0, arg1)
	ret0, _ := ret[0].(db.ArgoWorkflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArgoWorkflow indicates an expected call of GetArgoWorkflow.
func (mr *MockStoreMockRecorder) GetArgoWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArgoWorkflow", reflect.TypeOf((*MockStore)(nil).GetArgoWorkflow), arg0, arg1)
}

// GetEnvLayer mocks base method.
func (m *MockStore) GetEnvLayer(arg0 context.Context, arg1 int64) (db.EnvLayer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvLayer", arg0, arg1)
	ret0, _ := ret[0].(db.EnvLayer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvLayer indicates an expected call of GetEnvLayer.
func (mr *MockStoreMockRecorder) GetEnvLayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvLayer", reflect.TypeOf((*MockStore)(nil).GetEnvLayer), arg0, arg1)
}

// GetProcess mocks base method.
func (m *MockStore) GetProcess(arg0 context.Context, arg1 int64) (db.Process, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcess", arg0, arg1)
	ret0, _ := ret[0].(db.Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcess indicates an expected call of GetProcess.
func (mr *MockStoreMockRecorder) GetProcess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcess", reflect.TypeOf((*MockStore)(nil).GetProcess), arg0, arg1)
}

// GetProject mocks base method.
func (m *MockStore) GetProject(arg0 context.Context, arg1 int64) (db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0, arg1)
	ret0, _ := ret[0].(db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockStoreMockRecorder) GetProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockStore)(nil).GetProject), arg0, arg1)
}

// GetProjectEnvironment mocks base method.
func (m *MockStore) GetProjectEnvironment(arg0 context.Context, arg1 int64) (db.ProjectEnvironment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectEnvironment", arg0, arg1)
	ret0, _ := ret[0].(db.ProjectEnvironment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectEnvironment indicates an expected call of GetProjectEnvironment.
func (mr *MockStoreMockRecorder) GetProjectEnvironment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectEnvironment", reflect.TypeOf((*MockStore)(nil).GetProjectEnvironment), arg0, arg1)
}

// GetRequest mocks base method.
func (m *MockStore) GetRequest(arg0 context.Context, arg1 int64) (db.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest", arg0, arg1)
	ret0, _ := ret[0].(db.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockStoreMockRecorder) GetRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockStore)(nil).GetRequest), arg0, arg1)
}

// GetTemplate mocks base method.
func (m *MockStore) GetTemplate(arg0 context.Context, arg1 int64) (db.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplate", arg0, arg1)
	ret0, _ := ret[0].(db.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplate indicates an expected call of GetTemplate.
func (mr *MockStoreMockRecorder) GetTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplate", reflect.TypeOf((*MockStore)(nil).GetTemplate), arg0, arg1)
}

// ListArgoWorkflows mocks base method.
func (m *MockStore) ListArgoWorkflows(arg0 context.Context, arg1 db.ListArgoWorkflowsParams) ([]db.ArgoWorkflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListArgoWorkflows", arg0, arg1)
	ret0, _ := ret[0].([]db.ArgoWorkflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArgoWorkflows indicates an expected call of ListArgoWorkflows.
func (mr *MockStoreMockRecorder) ListArgoWorkflows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArgoWorkflows", reflect.TypeOf((*MockStore)(nil).ListArgoWorkflows), arg0, arg1)
}

// ListEnvLayers mocks base method.
func (m *MockStore) ListEnvLayers(arg0 context.Context, arg1 db.ListEnvLayersParams) ([]db.EnvLayer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnvLayers", arg0, arg1)
	ret0, _ := ret[0].([]db.EnvLayer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvLayers indicates an expected call of ListEnvLayers.
func (mr *MockStoreMockRecorder) ListEnvLayers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvLayers", reflect.TypeOf((*MockStore)(nil).ListEnvLayers), arg0, arg1)
}

// ListProcesss mocks base method.
func (m *MockStore) ListProcesss(arg0 context.Context, arg1 db.ListProcesssParams) ([]db.Process, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProcesss", arg0, arg1)
	ret0, _ := ret[0].([]db.Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProcesss indicates an expected call of ListProcesss.
func (mr *MockStoreMockRecorder) ListProcesss(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProcesss", reflect.TypeOf((*MockStore)(nil).ListProcesss), arg0, arg1)
}

// ListProjectEnvironments mocks base method.
func (m *MockStore) ListProjectEnvironments(arg0 context.Context, arg1 db.ListProjectEnvironmentsParams) ([]db.ProjectEnvironment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjectEnvironments", arg0, arg1)
	ret0, _ := ret[0].([]db.ProjectEnvironment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectEnvironments indicates an expected call of ListProjectEnvironments.
func (mr *MockStoreMockRecorder) ListProjectEnvironments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectEnvironments", reflect.TypeOf((*MockStore)(nil).ListProjectEnvironments), arg0, arg1)
}

// ListProjects mocks base method.
func (m *MockStore) ListProjects(arg0 context.Context, arg1 db.ListProjectsParams) ([]db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1)
	ret0, _ := ret[0].([]db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockStoreMockRecorder) ListProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockStore)(nil).ListProjects), arg0, arg1)
}

// ListRequests mocks base method.
func (m *MockStore) ListRequests(arg0 context.Context, arg1 db.ListRequestsParams) ([]db.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRequests", arg0, arg1)
	ret0, _ := ret[0].([]db.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRequests indicates an expected call of ListRequests.
func (mr *MockStoreMockRecorder) ListRequests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRequests", reflect.TypeOf((*MockStore)(nil).ListRequests), arg0, arg1)
}

// ListTemplates mocks base method.
func (m *MockStore) ListTemplates(arg0 context.Context, arg1 db.ListTemplatesParams) ([]db.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTemplates", arg0, arg1)
	ret0, _ := ret[0].([]db.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTemplates indicates an expected call of ListTemplates.
func (mr *MockStoreMockRecorder) ListTemplates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTemplates", reflect.TypeOf((*MockStore)(nil).ListTemplates), arg0, arg1)
}

// UpdateArgoWorkflow mocks base method.
func (m *MockStore) UpdateArgoWorkflow(arg0 context.Context, arg1 db.UpdateArgoWorkflowParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArgoWorkflow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateArgoWorkflow indicates an expected call of UpdateArgoWorkflow.
func (mr *MockStoreMockRecorder) UpdateArgoWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArgoWorkflow", reflect.TypeOf((*MockStore)(nil).UpdateArgoWorkflow), arg0, arg1)
}

// UpdateEnvLayer mocks base method.
func (m *MockStore) UpdateEnvLayer(arg0 context.Context, arg1 db.UpdateEnvLayerParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEnvLayer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEnvLayer indicates an expected call of UpdateEnvLayer.
func (mr *MockStoreMockRecorder) UpdateEnvLayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvLayer", reflect.TypeOf((*MockStore)(nil).UpdateEnvLayer), arg0, arg1)
}

// UpdateProcess mocks base method.
func (m *MockStore) UpdateProcess(arg0 context.Context, arg1 db.UpdateProcessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProcess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProcess indicates an expected call of UpdateProcess.
func (mr *MockStoreMockRecorder) UpdateProcess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProcess", reflect.TypeOf((*MockStore)(nil).UpdateProcess), arg0, arg1)
}

// UpdateProject mocks base method.
func (m *MockStore) UpdateProject(arg0 context.Context, arg1 db.UpdateProjectParams) (db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1)
	ret0, _ := ret[0].(db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockStoreMockRecorder) UpdateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockStore)(nil).UpdateProject), arg0, arg1)
}

// UpdateProjectEnvironment mocks base method.
func (m *MockStore) UpdateProjectEnvironment(arg0 context.Context, arg1 db.UpdateProjectEnvironmentParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectEnvironment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectEnvironment indicates an expected call of UpdateProjectEnvironment.
func (mr *MockStoreMockRecorder) UpdateProjectEnvironment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectEnvironment", reflect.TypeOf((*MockStore)(nil).UpdateProjectEnvironment), arg0, arg1)
}

// UpdateRequest mocks base method.
func (m *MockStore) UpdateRequest(arg0 context.Context, arg1 db.UpdateRequestParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockStoreMockRecorder) UpdateRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockStore)(nil).UpdateRequest), arg0, arg1)
}

// UpdateTemplate mocks base method.
func (m *MockStore) UpdateTemplate(arg0 context.Context, arg1 db.UpdateTemplateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTemplate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTemplate indicates an expected call of UpdateTemplate.
func (mr *MockStoreMockRecorder) UpdateTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTemplate", reflect.TypeOf((*MockStore)(nil).UpdateTemplate), arg0, arg1)
}
