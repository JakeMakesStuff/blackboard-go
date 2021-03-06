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

// RestException rest exception
//
// swagger:model RestException
type RestException struct {

	// The error code specific to a particular REST API. It is usually something that conveys information specific to the problem domain. For cases where the HTTP Status code conveys all the information required (such as a 404-Not Found) then the code may be omitted.
	Code string `json:"code,omitempty"`

	// Represents any technical information that a developer calling REST API might find useful.
	DeveloperMessage string `json:"developerMessage,omitempty"`

	// Indicates a URL that anyone seeing the error message can click in a browser. The target web page should describe the error condition fully, as well as potential solutions to help them resolve the error condition
	ExtraInfo string `json:"extraInfo,omitempty"`

	// Error message that should be easy to understand and convey a concise reason as to why the error occurred
	// Required: true
	Message *string `json:"message"`

	// Represents HTTP Status code in the response header.
	Status string `json:"status,omitempty"`
}

// Validate validates this rest exception
func (m *RestException) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestException) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestException) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestException) UnmarshalBinary(b []byte) error {
	var res RestException
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
