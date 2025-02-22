// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	context "context"

	component "github.com/itsopenmiso/openmiso-plugin-sdk/component"

	mock "github.com/stretchr/testify/mock"
)

// LogViewer is an autogenerated mock type for the LogViewer type
type LogViewer struct {
	mock.Mock
}

// NextLogBatch provides a mock function with given fields: ctx
func (_m *LogViewer) NextLogBatch(ctx context.Context) ([]component.LogEvent, error) {
	ret := _m.Called(ctx)

	var r0 []component.LogEvent
	if rf, ok := ret.Get(0).(func(context.Context) []component.LogEvent); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]component.LogEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
