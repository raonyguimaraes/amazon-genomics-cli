// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/mocks/manager/interfaces.go

// Package managermocks is a generated GoMock package.
package managermocks

import (
	io "io"
	reflect "reflect"

	workflow "github.com/aws/amazon-genomics-cli/internal/pkg/cli/workflow"
	gomock "github.com/golang/mock/gomock"
)

// MockWorkflowManager is a mock of WorkflowManager interface.
type MockWorkflowManager struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowManagerMockRecorder
}

// MockWorkflowManagerMockRecorder is the mock recorder for MockWorkflowManager.
type MockWorkflowManagerMockRecorder struct {
	mock *MockWorkflowManager
}

// NewMockWorkflowManager creates a new mock instance.
func NewMockWorkflowManager(ctrl *gomock.Controller) *MockWorkflowManager {
	mock := &MockWorkflowManager{ctrl: ctrl}
	mock.recorder = &MockWorkflowManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkflowManager) EXPECT() *MockWorkflowManagerMockRecorder {
	return m.recorder
}

// GetRunLog mocks base method.
func (m *MockWorkflowManager) GetRunLog(runId string) (workflow.RunLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunLog", runId)
	ret0, _ := ret[0].(workflow.RunLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunLog indicates an expected call of GetRunLog.
func (mr *MockWorkflowManagerMockRecorder) GetRunLog(runId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunLog", reflect.TypeOf((*MockWorkflowManager)(nil).GetRunLog), runId)
}

// GetRunLogData mocks base method.
func (m *MockWorkflowManager) GetRunLogData(runId, dataUrl string) (*io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunLogData", runId, dataUrl)
	ret0, _ := ret[0].(*io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunLogData indicates an expected call of GetRunLogData.
func (mr *MockWorkflowManagerMockRecorder) GetRunLogData(runId, dataUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunLogData", reflect.TypeOf((*MockWorkflowManager)(nil).GetRunLogData), runId, dataUrl)
}

// GetWorkflowTasks mocks base method.
func (m *MockWorkflowManager) GetWorkflowTasks(runId string) ([]workflow.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowTasks", runId)
	ret0, _ := ret[0].([]workflow.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowTasks indicates an expected call of GetWorkflowTasks.
func (mr *MockWorkflowManagerMockRecorder) GetWorkflowTasks(runId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowTasks", reflect.TypeOf((*MockWorkflowManager)(nil).GetWorkflowTasks), runId)
}

// OutputByInstanceId mocks base method.
func (m *MockWorkflowManager) OutputByInstanceId(instanceId string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutputByInstanceId", instanceId)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OutputByInstanceId indicates an expected call of OutputByInstanceId.
func (mr *MockWorkflowManagerMockRecorder) OutputByInstanceId(instanceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutputByInstanceId", reflect.TypeOf((*MockWorkflowManager)(nil).OutputByInstanceId), instanceId)
}

// StatusWorkflowAll mocks base method.
func (m *MockWorkflowManager) StatusWorkflowAll(numInstances int) ([]workflow.InstanceSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusWorkflowAll", numInstances)
	ret0, _ := ret[0].([]workflow.InstanceSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusWorkflowAll indicates an expected call of StatusWorkflowAll.
func (mr *MockWorkflowManagerMockRecorder) StatusWorkflowAll(numInstances interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusWorkflowAll", reflect.TypeOf((*MockWorkflowManager)(nil).StatusWorkflowAll), numInstances)
}

// StatusWorkflowByContext mocks base method.
func (m *MockWorkflowManager) StatusWorkflowByContext(contextName string, numInstances int) ([]workflow.InstanceSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusWorkflowByContext", contextName, numInstances)
	ret0, _ := ret[0].([]workflow.InstanceSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusWorkflowByContext indicates an expected call of StatusWorkflowByContext.
func (mr *MockWorkflowManagerMockRecorder) StatusWorkflowByContext(contextName, numInstances interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusWorkflowByContext", reflect.TypeOf((*MockWorkflowManager)(nil).StatusWorkflowByContext), contextName, numInstances)
}

// StatusWorkflowByInstanceId mocks base method.
func (m *MockWorkflowManager) StatusWorkflowByInstanceId(instanceId string) ([]workflow.InstanceSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusWorkflowByInstanceId", instanceId)
	ret0, _ := ret[0].([]workflow.InstanceSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusWorkflowByInstanceId indicates an expected call of StatusWorkflowByInstanceId.
func (mr *MockWorkflowManagerMockRecorder) StatusWorkflowByInstanceId(instanceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusWorkflowByInstanceId", reflect.TypeOf((*MockWorkflowManager)(nil).StatusWorkflowByInstanceId), instanceId)
}

// StatusWorkflowByName mocks base method.
func (m *MockWorkflowManager) StatusWorkflowByName(workflowName string, numInstances int) ([]workflow.InstanceSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusWorkflowByName", workflowName, numInstances)
	ret0, _ := ret[0].([]workflow.InstanceSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusWorkflowByName indicates an expected call of StatusWorkflowByName.
func (mr *MockWorkflowManagerMockRecorder) StatusWorkflowByName(workflowName, numInstances interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusWorkflowByName", reflect.TypeOf((*MockWorkflowManager)(nil).StatusWorkflowByName), workflowName, numInstances)
}
