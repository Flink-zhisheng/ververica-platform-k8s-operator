// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	appmanager "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// EventsService is an autogenerated mock type for the EventsService type
type EventsService struct {
	mock.Mock
}

// GetEvents provides a mock function with given fields: ctx, namespaceName, opts
func (_m *EventsService) GetEvents(ctx context.Context, namespaceName string, opts *appmanager.GetEventsOpts) (*appmanagerapi.ResourceListOfEvent, error) {
	ret := _m.Called(ctx, namespaceName, opts)

	var r0 *appmanagerapi.ResourceListOfEvent
	if rf, ok := ret.Get(0).(func(context.Context, string, *appmanager.GetEventsOpts) *appmanagerapi.ResourceListOfEvent); ok {
		r0 = rf(ctx, namespaceName, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appmanagerapi.ResourceListOfEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *appmanager.GetEventsOpts) error); ok {
		r1 = rf(ctx, namespaceName, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
