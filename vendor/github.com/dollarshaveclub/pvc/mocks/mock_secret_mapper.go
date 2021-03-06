// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/dollarshaveclub/pvc (interfaces: SecretMapper)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of SecretMapper interface
type MockSecretMapper struct {
	ctrl     *gomock.Controller
	recorder *_MockSecretMapperRecorder
}

// Recorder for MockSecretMapper (not exported)
type _MockSecretMapperRecorder struct {
	mock *MockSecretMapper
}

func NewMockSecretMapper(ctrl *gomock.Controller) *MockSecretMapper {
	mock := &MockSecretMapper{ctrl: ctrl}
	mock.recorder = &_MockSecretMapperRecorder{mock}
	return mock
}

func (_m *MockSecretMapper) EXPECT() *_MockSecretMapperRecorder {
	return _m.recorder
}

func (_m *MockSecretMapper) MapSecret(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "MapSecret", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSecretMapperRecorder) MapSecret(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MapSecret", arg0)
}
