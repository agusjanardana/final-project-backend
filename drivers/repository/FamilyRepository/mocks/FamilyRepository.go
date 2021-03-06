// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	records "vaccine-app-be/drivers/records"

	mock "github.com/stretchr/testify/mock"
)

// FamilyRepository is an autogenerated mock type for the FamilyRepository type
type FamilyRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, family
func (_m *FamilyRepository) Create(ctx context.Context, family records.FamilyMember) (records.FamilyMember, error) {
	ret := _m.Called(ctx, family)

	var r0 records.FamilyMember
	if rf, ok := ret.Get(0).(func(context.Context, records.FamilyMember) records.FamilyMember); ok {
		r0 = rf(ctx, family)
	} else {
		r0 = ret.Get(0).(records.FamilyMember)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, records.FamilyMember) error); ok {
		r1 = rf(ctx, family)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id, citizenId
func (_m *FamilyRepository) Delete(ctx context.Context, id int, citizenId int) (records.FamilyMember, error) {
	ret := _m.Called(ctx, id, citizenId)

	var r0 records.FamilyMember
	if rf, ok := ret.Get(0).(func(context.Context, int, int) records.FamilyMember); ok {
		r0 = rf(ctx, id, citizenId)
	} else {
		r0 = ret.Get(0).(records.FamilyMember)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, id, citizenId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCitizenOwnFamily provides a mock function with given fields: ctx, citizenId
func (_m *FamilyRepository) GetCitizenOwnFamily(ctx context.Context, citizenId int) ([]records.FamilyMember, error) {
	ret := _m.Called(ctx, citizenId)

	var r0 []records.FamilyMember
	if rf, ok := ret.Get(0).(func(context.Context, int) []records.FamilyMember); ok {
		r0 = rf(ctx, citizenId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]records.FamilyMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, citizenId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFamilyById provides a mock function with given fields: ctx, id
func (_m *FamilyRepository) GetFamilyById(ctx context.Context, id int) (records.FamilyMember, error) {
	ret := _m.Called(ctx, id)

	var r0 records.FamilyMember
	if rf, ok := ret.Get(0).(func(context.Context, int) records.FamilyMember); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(records.FamilyMember)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, family
func (_m *FamilyRepository) Update(ctx context.Context, id int, family records.FamilyMember) (records.FamilyMember, error) {
	ret := _m.Called(ctx, id, family)

	var r0 records.FamilyMember
	if rf, ok := ret.Get(0).(func(context.Context, int, records.FamilyMember) records.FamilyMember); ok {
		r0 = rf(ctx, id, family)
	} else {
		r0 = ret.Get(0).(records.FamilyMember)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, records.FamilyMember) error); ok {
		r1 = rf(ctx, id, family)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
