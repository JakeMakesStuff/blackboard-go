// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AggregateReviewStatus aggregate review status
//
// swagger:model AggregateReviewStatus
type AggregateReviewStatus struct {

	// The ID of the course.
	//
	// **Since**: 3700.15.0
	CourseID string `json:"courseId,omitempty"`

	// The number of content items that can be reviewed in the given course.
	//
	// **Since**: 3700.15.0
	ReviewableCount int32 `json:"reviewableCount,omitempty"`

	// The number of content items the user has reviewed.
	//
	// **Since**: 3700.15.0
	ReviewedCount int32 `json:"reviewedCount,omitempty"`

	// The ID of the User.
	//
	// **Since**: 3700.15.0
	UserID string `json:"userId,omitempty"`
}

// Validate validates this aggregate review status
func (m *AggregateReviewStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AggregateReviewStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateReviewStatus) UnmarshalBinary(b []byte) error {
	var res AggregateReviewStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
