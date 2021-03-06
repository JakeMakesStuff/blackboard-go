// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CourseRole course role
//
// swagger:model CourseRole
type CourseRole struct {

	// Implies access to unavailable courses, display in the Course catalog, and receiving email enrollment requests
	ActAsInstructor bool `json:"actAsInstructor,omitempty"`

	// availability
	Availability *CourseRoleAvailability `json:"availability,omitempty"`

	// False if the course role exists as a system generated course role. True if the course role was created by an admin user.
	//
	// Both custom and system generated course roles can be modified, but only custom course roles can be deleted.
	Custom bool `json:"custom,omitempty"`

	// Optional description of the course role
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// The primary ID of the course role
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// The role name presented to users when the course role is associated with a course.
	// Required: true
	// Max Length: 64
	NameForCourses *string `json:"nameForCourses"`

	// The role name presented to users when the course role is associated with an organization.
	// Required: true
	// Max Length: 64
	NameForOrganizations *string `json:"nameForOrganizations"`

	// The identifier used to assign the course role to a course enrollment.
	//
	// For system defined course roles, this value will be one of Student, Instructor, TeachingAssistant, CourseBuilder, Grader or Guest
	//
	// For custom course roles, this will be the roleId entered during the creation of the role. Allowed characters for the custom course roleId's are any letter, number, period, underscore and dash.
	// Required: true
	// Read Only: true
	RoleID string `json:"roleId"`
}

// Validate validates this course role
func (m *CourseRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameForCourses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameForOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CourseRole) validateAvailability(formats strfmt.Registry) error {

	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	if m.Availability != nil {
		if err := m.Availability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availability")
			}
			return err
		}
	}

	return nil
}

func (m *CourseRole) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *CourseRole) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CourseRole) validateNameForCourses(formats strfmt.Registry) error {

	if err := validate.Required("nameForCourses", "body", m.NameForCourses); err != nil {
		return err
	}

	if err := validate.MaxLength("nameForCourses", "body", string(*m.NameForCourses), 64); err != nil {
		return err
	}

	return nil
}

func (m *CourseRole) validateNameForOrganizations(formats strfmt.Registry) error {

	if err := validate.Required("nameForOrganizations", "body", m.NameForOrganizations); err != nil {
		return err
	}

	if err := validate.MaxLength("nameForOrganizations", "body", string(*m.NameForOrganizations), 64); err != nil {
		return err
	}

	return nil
}

func (m *CourseRole) validateRoleID(formats strfmt.Registry) error {

	if err := validate.RequiredString("roleId", "body", string(m.RoleID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CourseRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CourseRole) UnmarshalBinary(b []byte) error {
	var res CourseRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CourseRoleAvailability Availability
//
// swagger:model CourseRoleAvailability
type CourseRoleAvailability struct {

	// Whether the course role is currently available to course enrollments, organization enrollments, both or neither.  Valid values are:
	//
	// - Course: Course Role is available to Course Enrollments
	// - Organization: Course Role is available to Organization Enrollments
	// - Both: Course Role is available to both Course and Organization Enrollments
	// - None: Course Role is not available
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Course | Course only |
	// | Organization | Organization only |
	// | CourseAndOrganization | Both Course and Organization |
	// | No | Neither Course nor Organization |
	//
	// Enum: [Course Organization CourseAndOrganization No]
	Available string `json:"available,omitempty"`
}

// Validate validates this course role availability
func (m *CourseRoleAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var courseRoleAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Course","Organization","CourseAndOrganization","No"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		courseRoleAvailabilityTypeAvailablePropEnum = append(courseRoleAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// CourseRoleAvailabilityAvailableCourse captures enum value "Course"
	CourseRoleAvailabilityAvailableCourse string = "Course"

	// CourseRoleAvailabilityAvailableOrganization captures enum value "Organization"
	CourseRoleAvailabilityAvailableOrganization string = "Organization"

	// CourseRoleAvailabilityAvailableCourseAndOrganization captures enum value "CourseAndOrganization"
	CourseRoleAvailabilityAvailableCourseAndOrganization string = "CourseAndOrganization"

	// CourseRoleAvailabilityAvailableNo captures enum value "No"
	CourseRoleAvailabilityAvailableNo string = "No"
)

// prop value enum
func (m *CourseRoleAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, courseRoleAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CourseRoleAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(m.Available) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailableEnum("availability"+"."+"available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CourseRoleAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CourseRoleAvailability) UnmarshalBinary(b []byte) error {
	var res CourseRoleAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
