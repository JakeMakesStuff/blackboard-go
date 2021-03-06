// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateAssignmentResult create assignment result
//
// swagger:model CreateAssignmentResult
type CreateAssignmentResult struct {

	// Id of the assessment created during the assignment creation.  This is only created for Ultra course assignments.
	AssessmentID string `json:"assessmentId,omitempty"`

	// Ids of the file attachments created during assignment creation.  This is only created for Classic course assignments
	AttachmentIds []string `json:"attachmentIds"`

	// Id of the content created during the assignment creation.
	ContentID string `json:"contentId,omitempty"`

	// Id of the grade column created during the assignment creation.
	GradeColumnID string `json:"gradeColumnId,omitempty"`

	// Ids of the assessment questions created during the assignment creation.  This is only created for Ultra course assignments.
	QuestionIds []string `json:"questionIds"`
}

// Validate validates this create assignment result
func (m *CreateAssignmentResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAssignmentResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAssignmentResult) UnmarshalBinary(b []byte) error {
	var res CreateAssignmentResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
