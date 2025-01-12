// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package services

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/trento-project/trento/web/models"
)

// MockHostsService is an autogenerated mock type for the HostsService type
type MockHostsService struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: _a0, _a1
func (_m *MockHostsService) GetAll(_a0 *HostsFilter, _a1 *Page) (models.HostList, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.HostList
	if rf, ok := ret.Get(0).(func(*HostsFilter, *Page) models.HostList); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.HostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*HostsFilter, *Page) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBySAPSystemID provides a mock function with given fields: _a0
func (_m *MockHostsService) GetAllBySAPSystemID(_a0 string) (models.HostList, error) {
	ret := _m.Called(_a0)

	var r0 models.HostList
	if rf, ok := ret.Get(0).(func(string) models.HostList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.HostList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSIDs provides a mock function with given fields:
func (_m *MockHostsService) GetAllSIDs() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTags provides a mock function with given fields:
func (_m *MockHostsService) GetAllTags() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *MockHostsService) GetByID(_a0 string) (*models.Host, error) {
	ret := _m.Called(_a0)

	var r0 *models.Host
	if rf, ok := ret.Get(0).(func(string) *models.Host); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Host)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCount provides a mock function with given fields:
func (_m *MockHostsService) GetCount() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Heartbeat provides a mock function with given fields: agentID
func (_m *MockHostsService) Heartbeat(agentID string) error {
	ret := _m.Called(agentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(agentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
