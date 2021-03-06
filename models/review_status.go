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

// ReviewStatus review status
//
// swagger:model ReviewStatus
type ReviewStatus struct {

	// The ID of the content.
	//
	// **Since**: 3700.16.0
	// Required: true
	// Read Only: true
	ContentID string `json:"contentId"`

	// The date that the user marked the content as reviewed
	//
	// **Since**: 3700.16.0
	// Required: true
	// Read Only: true
	// Format: date-time
	ReviewDate strfmt.DateTime `json:"reviewDate"`

	// The current status of the content's 'reviewed' attribute.
	//
	// **Since**: 3700.16.0
	Reviewed bool `json:"reviewed,omitempty"`

	// The ID of the user.
	//
	// **Since**: 3700.16.0
	// Required: true
	// Read Only: true
	UserID string `json:"userId"`
}

// Validate validates this review status
func (m *ReviewStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewDate(formats); err != nil {
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

func (m *ReviewStatus) validateContentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("contentId", "body", string(m.ContentID)); err != nil {
		return err
	}

	return nil
}

func (m *ReviewStatus) validateReviewDate(formats strfmt.Registry) error {

	if err := validate.Required("reviewDate", "body", strfmt.DateTime(m.ReviewDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("reviewDate", "body", "date-time", m.ReviewDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReviewStatus) validateUserID(formats strfmt.Registry) error {

	if err := validate.RequiredString("userId", "body", string(m.UserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReviewStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReviewStatus) UnmarshalBinary(b []byte) error {
	var res ReviewStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
