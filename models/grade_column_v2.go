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

// GradeColumnV2 grade column v2
//
// swagger:model GradeColumnV2
type GradeColumnV2 struct {

	// availability
	Availability *GradeColumnV2Availability `json:"availability,omitempty"`

	// For grade columns associated with a content item, the ID of the content item.
	//
	// **Since**: 3000.11.0
	// Required: true
	// Read Only: true
	ContentID string `json:"contentId"`

	// The date this grade column was created.
	// Required: true
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created"`

	// The description of the grade column.
	Description string `json:"description,omitempty"`

	// The display name of the grade column. Only applicable for Classic courses. Ultra courses will simply use the `name` field.
	//
	// **Since**: 3300.2.0
	DisplayName string `json:"displayName,omitempty"`

	// Whether this grade column is an external grade column.
	ExternalGrade bool `json:"externalGrade,omitempty"`

	// The externalId for this grade column
	ExternalID string `json:"externalId,omitempty"`

	// The externalId for this grade column
	//
	// **Since**: 3500.2.0
	ExternalToolID string `json:"externalToolId,omitempty"`

	// The formula used for determining the value for the grade column, if it is a calculated column.
	//
	// **Since**: 3400.5.0
	// Required: true
	Formula *GradingFormulaV2 `json:"formula"`

	// The gradebook category ID for the grade column.
	//
	// **Since**: 3400.2.0
	GradebookCategoryID string `json:"gradebookCategoryId,omitempty"`

	// grading
	// Required: true
	Grading *GradeColumnV2Grading `json:"grading"`

	// The primary ID of the grade column.
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// Indicates whether or not this column is included in gradebook calculations. Cannot be set for Ultra courses. Default: true
	//
	// **Since**: 3800.4.0
	IncludeInCalculations bool `json:"includeInCalculations,omitempty"`

	// The name of the grade column.
	// Required: true
	Name *string `json:"name"`

	// score
	Score *GradeColumnV2Score `json:"score,omitempty"`

	// Indicates whether or not column statistics are shown to students. Read-only for Ultra courses. Default: false
	//
	// **Since**: 3800.4.0
	ShowStatisticsToStudents bool `json:"showStatisticsToStudents,omitempty"`
}

// Validate validates this grade column v2
func (m *GradeColumnV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormula(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GradeColumnV2) validateAvailability(formats strfmt.Registry) error {

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

func (m *GradeColumnV2) validateContentID(formats strfmt.Registry) error {

	if err := validate.RequiredString("contentId", "body", string(m.ContentID)); err != nil {
		return err
	}

	return nil
}

func (m *GradeColumnV2) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GradeColumnV2) validateFormula(formats strfmt.Registry) error {

	if err := validate.Required("formula", "body", m.Formula); err != nil {
		return err
	}

	if m.Formula != nil {
		if err := m.Formula.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formula")
			}
			return err
		}
	}

	return nil
}

func (m *GradeColumnV2) validateGrading(formats strfmt.Registry) error {

	if err := validate.Required("grading", "body", m.Grading); err != nil {
		return err
	}

	if m.Grading != nil {
		if err := m.Grading.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grading")
			}
			return err
		}
	}

	return nil
}

func (m *GradeColumnV2) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *GradeColumnV2) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GradeColumnV2) validateScore(formats strfmt.Registry) error {

	if swag.IsZero(m.Score) { // not required
		return nil
	}

	if m.Score != nil {
		if err := m.Score.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("score")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GradeColumnV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GradeColumnV2) UnmarshalBinary(b []byte) error {
	var res GradeColumnV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GradeColumnV2Availability Availability
//
// Settings controlling the availability/visibility of grade column data.
//
// swagger:model GradeColumnV2Availability
type GradeColumnV2Availability struct {

	// Whether this grade column is available to students
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes | Students may view the grade column. |
	// | No | Students may not view the grade column. |
	//
	// Enum: [Yes No]
	Available string `json:"available,omitempty"`
}

// Validate validates this grade column v2 availability
func (m *GradeColumnV2Availability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gradeColumnV2AvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gradeColumnV2AvailabilityTypeAvailablePropEnum = append(gradeColumnV2AvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// GradeColumnV2AvailabilityAvailableYes captures enum value "Yes"
	GradeColumnV2AvailabilityAvailableYes string = "Yes"

	// GradeColumnV2AvailabilityAvailableNo captures enum value "No"
	GradeColumnV2AvailabilityAvailableNo string = "No"
)

// prop value enum
func (m *GradeColumnV2Availability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gradeColumnV2AvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GradeColumnV2Availability) validateAvailable(formats strfmt.Registry) error {

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
func (m *GradeColumnV2Availability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GradeColumnV2Availability) UnmarshalBinary(b []byte) error {
	var res GradeColumnV2Availability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GradeColumnV2Grading Grading
//
// Settings controlling whether numerical and text grade values for this grade column are calculated, determined based on attempts, or manually entered.
//
// swagger:model GradeColumnV2Grading
type GradeColumnV2Grading struct {

	// anonymous grading
	// Required: true
	AnonymousGrading *GradeColumnV2GradingAnonymousGrading `json:"anonymousGrading"`

	// Number of attempts allowed for the grade column.
	AttemptsAllowed int32 `json:"attemptsAllowed,omitempty"`

	// The date on which attempts are due for the grade column.
	// Format: date-time
	Due strfmt.DateTime `json:"due,omitempty"`

	// The ID of the grade schema associated with this grade column. Mutable since 3400.2.0
	//
	// **Since**: 3200.13.0
	SchemaID string `json:"schemaId,omitempty"`

	// The scoring model for the submitted grade column attempts.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Last |  |
	// | Highest |  |
	// | Lowest |  |
	// | First |  |
	// | Average |  |
	//
	// Enum: [Last Highest Lowest First Average]
	ScoringModel string `json:"scoringModel,omitempty"`

	// The type of Grading settings for this Grade Column.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Attempts | Indicates score and grade values are determined based on user attempts |
	// | Calculated | Indicates score and grade values are determined by applying a calculated formula. |
	// | Manual | Indicates score and grade values are manually entered. |
	//
	// Required: true
	// Read Only: true
	// Enum: [Attempts Calculated Manual]
	Type string `json:"type"`
}

// Validate validates this grade column v2 grading
func (m *GradeColumnV2Grading) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnonymousGrading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoringModel(formats); err != nil {
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

func (m *GradeColumnV2Grading) validateAnonymousGrading(formats strfmt.Registry) error {

	if err := validate.Required("grading"+"."+"anonymousGrading", "body", m.AnonymousGrading); err != nil {
		return err
	}

	if m.AnonymousGrading != nil {
		if err := m.AnonymousGrading.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grading" + "." + "anonymousGrading")
			}
			return err
		}
	}

	return nil
}

func (m *GradeColumnV2Grading) validateDue(formats strfmt.Registry) error {

	if swag.IsZero(m.Due) { // not required
		return nil
	}

	if err := validate.FormatOf("grading"+"."+"due", "body", "date-time", m.Due.String(), formats); err != nil {
		return err
	}

	return nil
}

var gradeColumnV2GradingTypeScoringModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Last","Highest","Lowest","First","Average"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gradeColumnV2GradingTypeScoringModelPropEnum = append(gradeColumnV2GradingTypeScoringModelPropEnum, v)
	}
}

const (

	// GradeColumnV2GradingScoringModelLast captures enum value "Last"
	GradeColumnV2GradingScoringModelLast string = "Last"

	// GradeColumnV2GradingScoringModelHighest captures enum value "Highest"
	GradeColumnV2GradingScoringModelHighest string = "Highest"

	// GradeColumnV2GradingScoringModelLowest captures enum value "Lowest"
	GradeColumnV2GradingScoringModelLowest string = "Lowest"

	// GradeColumnV2GradingScoringModelFirst captures enum value "First"
	GradeColumnV2GradingScoringModelFirst string = "First"

	// GradeColumnV2GradingScoringModelAverage captures enum value "Average"
	GradeColumnV2GradingScoringModelAverage string = "Average"
)

// prop value enum
func (m *GradeColumnV2Grading) validateScoringModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gradeColumnV2GradingTypeScoringModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GradeColumnV2Grading) validateScoringModel(formats strfmt.Registry) error {

	if swag.IsZero(m.ScoringModel) { // not required
		return nil
	}

	// value enum
	if err := m.validateScoringModelEnum("grading"+"."+"scoringModel", "body", m.ScoringModel); err != nil {
		return err
	}

	return nil
}

var gradeColumnV2GradingTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Attempts","Calculated","Manual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gradeColumnV2GradingTypeTypePropEnum = append(gradeColumnV2GradingTypeTypePropEnum, v)
	}
}

const (

	// GradeColumnV2GradingTypeAttempts captures enum value "Attempts"
	GradeColumnV2GradingTypeAttempts string = "Attempts"

	// GradeColumnV2GradingTypeCalculated captures enum value "Calculated"
	GradeColumnV2GradingTypeCalculated string = "Calculated"

	// GradeColumnV2GradingTypeManual captures enum value "Manual"
	GradeColumnV2GradingTypeManual string = "Manual"
)

// prop value enum
func (m *GradeColumnV2Grading) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gradeColumnV2GradingTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GradeColumnV2Grading) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("grading"+"."+"type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("grading"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GradeColumnV2Grading) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GradeColumnV2Grading) UnmarshalBinary(b []byte) error {
	var res GradeColumnV2Grading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GradeColumnV2GradingAnonymousGrading AnonymousGrading
//
// Settings for anonymous grading
//
// swagger:model GradeColumnV2GradingAnonymousGrading
type GradeColumnV2GradingAnonymousGrading struct {

	// Date after which grades are released from being anonymized, if AnonymousGrading type is 'Date'.
	// Format: date-time
	ReleaseAfter strfmt.DateTime `json:"releaseAfter,omitempty"`

	// The type of AnonymousGrading settings for this Attempts based Grade Column.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | None | Indicates anonymous grading is not enabled. |
	// | AfterAllGraded | Indicates anonymized grades are released after all attempts have been graded. |
	// | Date | Indicates anonymized grades are released after a specified release date. |
	//
	// Required: true
	// Enum: [None AfterAllGraded Date]
	Type *string `json:"type"`
}

// Validate validates this grade column v2 grading anonymous grading
func (m *GradeColumnV2GradingAnonymousGrading) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReleaseAfter(formats); err != nil {
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

func (m *GradeColumnV2GradingAnonymousGrading) validateReleaseAfter(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("grading"+"."+"anonymousGrading"+"."+"releaseAfter", "body", "date-time", m.ReleaseAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

var gradeColumnV2GradingAnonymousGradingTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","AfterAllGraded","Date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gradeColumnV2GradingAnonymousGradingTypeTypePropEnum = append(gradeColumnV2GradingAnonymousGradingTypeTypePropEnum, v)
	}
}

const (

	// GradeColumnV2GradingAnonymousGradingTypeNone captures enum value "None"
	GradeColumnV2GradingAnonymousGradingTypeNone string = "None"

	// GradeColumnV2GradingAnonymousGradingTypeAfterAllGraded captures enum value "AfterAllGraded"
	GradeColumnV2GradingAnonymousGradingTypeAfterAllGraded string = "AfterAllGraded"

	// GradeColumnV2GradingAnonymousGradingTypeDate captures enum value "Date"
	GradeColumnV2GradingAnonymousGradingTypeDate string = "Date"
)

// prop value enum
func (m *GradeColumnV2GradingAnonymousGrading) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gradeColumnV2GradingAnonymousGradingTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GradeColumnV2GradingAnonymousGrading) validateType(formats strfmt.Registry) error {

	if err := validate.Required("grading"+"."+"anonymousGrading"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("grading"+"."+"anonymousGrading"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GradeColumnV2GradingAnonymousGrading) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GradeColumnV2GradingAnonymousGrading) UnmarshalBinary(b []byte) error {
	var res GradeColumnV2GradingAnonymousGrading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GradeColumnV2Score Score
//
// Settings controlling the numerical scoring of this grade column.
//
// swagger:model GradeColumnV2Score
type GradeColumnV2Score struct {

	// The points possible for this grade column.
	Possible float64 `json:"possible,omitempty"`
}

// Validate validates this grade column v2 score
func (m *GradeColumnV2Score) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GradeColumnV2Score) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GradeColumnV2Score) UnmarshalBinary(b []byte) error {
	var res GradeColumnV2Score
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
