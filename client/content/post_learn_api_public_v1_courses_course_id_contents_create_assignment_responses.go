// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentReader is a Reader for the PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignment structure.
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated creates a PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated() *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated {
	return &PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated{}
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated struct {
	Payload *models.CreateAssignmentResult
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/contents/createAssignment][%d] postLearnApiPublicV1CoursesCourseIdContentsCreateAssignmentCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated) GetPayload() *models.CreateAssignmentResult {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateAssignmentResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest creates a PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest() *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest {
	return &PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest{}
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest handles this case with default header values.

Bad Request
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/contents/createAssignment][%d] postLearnApiPublicV1CoursesCourseIdContentsCreateAssignmentBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden creates a PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden() *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden {
	return &PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden{}
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden handles this case with default header values.

Logged-on User has insufficient privileges
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/contents/createAssignment][%d] postLearnApiPublicV1CoursesCourseIdContentsCreateAssignmentForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody post learn API public v1 courses course ID contents create assignment body
swagger:model PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody struct {

	// availability
	Availability *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability `json:"availability,omitempty"`

	// The description to use when creating the assignment content.
	Description string `json:"description,omitempty"`

	// file upload ids
	FileUploadIds []string `json:"fileUploadIds"`

	// grading
	Grading *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading `json:"grading,omitempty"`

	// The text instructions to use when creating the assignment content. This field supports BbML; see <a target='_blank' href='https://docs.blackboard.com/learn/REST/Blackboard%20Markup%20Language%20-%20BbML.html'>here</a> for more information.
	Instructions string `json:"instructions,omitempty"`

	// The Originality Reporting Tool options to be used for the assignment content item.
	//
	// **Since**: 3800.16.0
	OriginalityReportingTool *models.OriginalityReportingTool `json:"originalityReportingTool,omitempty"`

	// The id of the parent content for the created assignment.
	ParentID string `json:"parentId,omitempty"`

	// The position of the created assignment within the other other content of its parent. Position values are zero-based (the first element has a position value of zero, not one). Default position is last in the list of child contents under the parent.
	Position int32 `json:"position,omitempty"`

	// score
	Score *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyScore `json:"score,omitempty"`

	// The title used for the created assignment content and gradebook column. Typically shown as the main text to click in the course outline when accessing the assignment.
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this post learn API public v1 courses course ID contents create assignment body
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGrading(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOriginalityReportingTool(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody) validateAvailability(formats strfmt.Registry) error {

	if swag.IsZero(o.Availability) { // not required
		return nil
	}

	if o.Availability != nil {
		if err := o.Availability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "availability")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody) validateGrading(formats strfmt.Registry) error {

	if swag.IsZero(o.Grading) { // not required
		return nil
	}

	if o.Grading != nil {
		if err := o.Grading.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "grading")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody) validateOriginalityReportingTool(formats strfmt.Registry) error {

	if swag.IsZero(o.OriginalityReportingTool) { // not required
		return nil
	}

	if o.OriginalityReportingTool != nil {
		if err := o.OriginalityReportingTool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "originalityReportingTool")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody) validateScore(formats strfmt.Registry) error {

	if swag.IsZero(o.Score) { // not required
		return nil
	}

	if o.Score != nil {
		if err := o.Score.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "score")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability Availability
swagger:model PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability struct {

	// adaptive release
	AdaptiveRelease *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease `json:"adaptiveRelease,omitempty"`

	// Whether the created assignment is available to users with the 'guest' role. Defaults to true.
	AllowGuests bool `json:"allowGuests,omitempty"`

	// Whether the created assignment is available to students. Instructors can always access the created assignment. If set to 'PartiallyVisible', the title will be available to students but the body will not. Defaults to Yes.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes |  |
	// | No |  |
	// | PartiallyVisible |  |
	//
	// Enum: [Yes No PartiallyVisible]
	Available string `json:"available,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID contents create assignment params body availability
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdaptiveRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability) validateAdaptiveRelease(formats strfmt.Registry) error {

	if swag.IsZero(o.AdaptiveRelease) { // not required
		return nil
	}

	if o.AdaptiveRelease != nil {
		if err := o.AdaptiveRelease.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "availability" + "." + "adaptiveRelease")
			}
			return err
		}
	}

	return nil
}

var postLearnApiPublicV1CoursesCourseIdContentsCreateAssignmentParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No","PartiallyVisible"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdContentsCreateAssignmentParamsBodyAvailabilityTypeAvailablePropEnum = append(postLearnApiPublicV1CoursesCourseIdContentsCreateAssignmentParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAvailableYes string = "Yes"

	// PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAvailableNo captures enum value "No"
	PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAvailableNo string = "No"

	// PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAvailablePartiallyVisible captures enum value "PartiallyVisible"
	PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAvailablePartiallyVisible string = "PartiallyVisible"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdContentsCreateAssignmentParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(o.Available) { // not required
		return nil
	}

	// value enum
	if err := o.validateAvailableEnum("input"+"."+"availability"+"."+"available", "body", o.Available); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease AdaptiveRelease
//
// Settings controlling adaptive release of created assignment to students.
swagger:model PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease struct {

	// The date when the created assignment will no longer be available to students.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The date when the created assignment will become available to students.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID contents create assignment params body availability adaptive release
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(o.End) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"adaptiveRelease"+"."+"end", "body", "date-time", o.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(o.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"adaptiveRelease"+"."+"start", "body", "date-time", o.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyAvailabilityAdaptiveRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading Grading
swagger:model PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading struct {

	// The number of attempts allowed on the created assignment. Defaults to 1. Maximum allowed is 10 for an Ultra Assignment. Value will be ignored if isUnlimitedAttemptsAllowed is set to true.
	AttemptsAllowed int32 `json:"attemptsAllowed,omitempty"`

	// Date and time that the created assignment will be due. If not specified, this will default to the specified "availability.adaptiveRelease.end" date. If that is also not specified, due date defaults to null.
	// Format: date-time
	Due strfmt.DateTime `json:"due,omitempty"`

	// The grading schema to use for the created assignment. Defaults to Score.
	GradeSchemaID string `json:"gradeSchemaId,omitempty"`

	// Determines if the assignment has unlimited number of attempts.
	//
	// **Since**: 3400.8.0
	IsUnlimitedAttemptsAllowed bool `json:"isUnlimitedAttemptsAllowed,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID contents create assignment params body grading
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading) validateDue(formats strfmt.Registry) error {

	if swag.IsZero(o.Due) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"grading"+"."+"due", "body", "date-time", o.Due.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyGrading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyScore Score
swagger:model PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyScore
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyScore struct {

	// The number of points possible for the created assignment. Defaults to 100.
	Possible float64 `json:"possible,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID contents create assignment params body score
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyScore) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyScore) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyScore) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDContentsCreateAssignmentParamsBodyScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
