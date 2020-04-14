// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/models"
	mock "github.com/stretchr/testify/mock"

	router "github.com/monitoror/monitoror/service/router"
)

// MonitorableRouter is an autogenerated mock type for the MonitorableRouter type
type MonitorableRouter struct {
	mock.Mock
}

// Group provides a mock function with given fields: path, variant
func (_m *MonitorableRouter) Group(path string, variantName models.VariantName) router.MonitorableRouterGroup {
	ret := _m.Called(path, variantName)

	var r0 router.MonitorableRouterGroup
	if rf, ok := ret.Get(0).(func(string, models.VariantName) router.MonitorableRouterGroup); ok {
		r0 = rf(path, variantName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(router.MonitorableRouterGroup)
		}
	}

	return r0
}
