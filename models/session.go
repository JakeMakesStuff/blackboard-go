// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Session session
//
// swagger:model Session
type Session struct {

	// Time when the session started.
	// Required: true
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created"`

	// Unique identifier of the session.
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// Time when the session was last accessed.
	// Required: true
	// Read Only: true
	// Format: date-time
	LastAccess strfmt.DateTime `json:"lastAccess"`

	// Indicates whether this is a mobile session.
	// Required: true
	// Read Only: true
	Mobile bool `json:"mobile"`

	// Full information of the logged user. This is only set if the caller requests to expand the user information.
	// Required: true
	User *User `json:"user"`

	// Id of the logged in user.
	// Required: true
	// Read Only: true
	UserID string `json:"userId"`
}

// Validate validates this session
func (m *Session) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Session) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateLastAccess(formats strfmt.Registry) error {

	if err := validate.Required("lastAccess", "body", strfmt.DateTime(m.LastAccess)); err != nil {
		return err
	}

	if err := validate.FormatOf("lastAccess", "body", "date-time", m.LastAccess.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateMobile(formats strfmt.Registry) error {

	if err := validate.Required("mobile", "body", bool(m.Mobile)); err != nil {
		return err
	}

	return nil
}

func (m *Session) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *Session) validateUserID(formats strfmt.Registry) error {

	if err := validate.RequiredString("userId", "body", string(m.UserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Session) UnmarshalBinary(b []byte) error {
	var res Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}