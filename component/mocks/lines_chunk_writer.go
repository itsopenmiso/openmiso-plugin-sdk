// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LinesChunkWriter is an autogenerated mock type for the LinesChunkWriter type
type LinesChunkWriter struct {
	mock.Mock
}

// OutputLines provides a mock function with given fields: _a0
func (_m *LinesChunkWriter) OutputLines(_a0 []string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}