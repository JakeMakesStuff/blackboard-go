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

// Term term
//
// swagger:model Term
type Term struct {

	// availability
	Availability *TermAvailability `json:"availability,omitempty"`

	// The ID of the data source associated with this term.  This may optionally be the data source's externalId using the syntax "externalId:math101".
	DataSourceID string `json:"dataSourceId,omitempty"`

	// The description of the term. This field supports BbML; see <a target='_blank' href='https://docs.blackboard.com/learn/REST/Blackboard%20Markup%20Language%20-%20BbML.html'>here</a> for more information.
	Description string `json:"description,omitempty"`

	// An externally-defined unique ID for the term.
	//
	// Formerly known as 'sourcedidId'.
	// Required: true
	// Max Length: 256
	ExternalID *string `json:"externalId"`

	// The primary ID of the term.
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// The name of the term.
	// Required: true
	// Max Length: 333
	Name *string `json:"name"`
}

// Validate validates this term
func (m *Term) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Term) validateAvailability(formats strfmt.Registry) error {

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

func (m *Term) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("externalId", "body", m.ExternalID); err != nil {
		return err
	}

	if err := validate.MaxLength("externalId", "body", string(*m.ExternalID), 256); err != nil {
		return err
	}

	return nil
}

func (m *Term) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Term) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 333); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Term) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Term) UnmarshalBinary(b []byte) error {
	var res Term
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TermAvailability Availability
//
// Settings controlling availability of the term to students.
//
// swagger:model TermAvailability
type TermAvailability struct {

	// Whether the term and the courses it contains are available to students.  Instructors can always access their courses.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes | Students may access the term and the courses it contains. |
	// | No | Students may not access the term or the courses it contains. |
	//
	// Enum: [Yes No]
	Available string `json:"available,omitempty"`

	// duration
	Duration *TermAvailabilityDuration `json:"duration,omitempty"`
}

// Validate validates this term availability
func (m *TermAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var termAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		termAvailabilityTypeAvailablePropEnum = append(termAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// TermAvailabilityAvailableYes captures enum value "Yes"
	TermAvailabilityAvailableYes string = "Yes"

	// TermAvailabilityAvailableNo captures enum value "No"
	TermAvailabilityAvailableNo string = "No"
)

// prop value enum
func (m *TermAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, termAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TermAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(m.Available) { // not required
		return nil
	}

	// value enum
	if err := m.validateAvailableEnum("availability"+"."+"available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

func (m *TermAvailability) validateDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if m.Duration != nil {
		if err := m.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availability" + "." + "duration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TermAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TermAvailability) UnmarshalBinary(b []byte) error {
	var res TermAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TermAvailabilityDuration Duration
//
// Settings controlling the length of time the term is available.
//
// swagger:model TermAvailabilityDuration
type TermAvailabilityDuration struct {

	// The number of days courses within this term can be used.  May only be set if availability.duration.type is FixedNumDays.
	DaysOfUse int32 `json:"daysOfUse,omitempty"`

	// The date this term ends.  May only be set if availability.duration.type is DateRange.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The date this term starts.  May only be set if availability.duration.type is DateRange.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// The intended length of the term.  Possible values are:
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Continuous | The term is active on an ongoing basis. This is the default. |
	// | DateRange | The term will only be available between specific date ranges. |
	// | FixedNumDays | The term will only be available for a set number of days. |
	//
	// Enum: [Continuous DateRange FixedNumDays]
	Type string `json:"type,omitempty"`
}

// Validate validates this term availability duration
func (m *TermAvailabilityDuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TermAvailabilityDuration) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("availability"+"."+"duration"+"."+"end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TermAvailabilityDuration) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("availability"+"."+"duration"+"."+"start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var termAvailabilityDurationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Continuous","DateRange","FixedNumDays"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		termAvailabilityDurationTypeTypePropEnum = append(termAvailabilityDurationTypeTypePropEnum, v)
	}
}

const (

	// TermAvailabilityDurationTypeContinuous captures enum value "Continuous"
	TermAvailabilityDurationTypeContinuous string = "Continuous"

	// TermAvailabilityDurationTypeDateRange captures enum value "DateRange"
	TermAvailabilityDurationTypeDateRange string = "DateRange"

	// TermAvailabilityDurationTypeFixedNumDays captures enum value "FixedNumDays"
	TermAvailabilityDurationTypeFixedNumDays string = "FixedNumDays"
)

// prop value enum
func (m *TermAvailabilityDuration) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, termAvailabilityDurationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TermAvailabilityDuration) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("availability"+"."+"duration"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TermAvailabilityDuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TermAvailabilityDuration) UnmarshalBinary(b []byte) error {
	var res TermAvailabilityDuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}