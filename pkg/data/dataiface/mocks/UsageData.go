// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import usage "github.com/Optum/dce/pkg/usage"

// UsageData is an autogenerated mock type for the UsageData type
type UsageData struct {
	mock.Mock
}

// Write provides a mock function with given fields: usg
func (_m *UsageData) Write(usg *usage.Usage) (*usage.Usage, error) {
	ret := _m.Called(usg)

	var r0 *usage.Usage
	if rf, ok := ret.Get(0).(func(*usage.Usage) *usage.Usage); ok {
		r0 = rf(usg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usage.Usage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usage.Usage) error); ok {
		r1 = rf(usg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}