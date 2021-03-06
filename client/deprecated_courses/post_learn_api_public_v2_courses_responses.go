// Code generated by go-swagger; DO NOT EDIT.

package deprecated_courses

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

// PostLearnAPIPublicV2CoursesReader is a Reader for the PostLearnAPIPublicV2Courses structure.
type PostLearnAPIPublicV2CoursesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV2CoursesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV2CoursesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV2CoursesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV2CoursesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLearnAPIPublicV2CoursesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV2CoursesCreated creates a PostLearnAPIPublicV2CoursesCreated with default headers values
func NewPostLearnAPIPublicV2CoursesCreated() *PostLearnAPIPublicV2CoursesCreated {
	return &PostLearnAPIPublicV2CoursesCreated{}
}

/*PostLearnAPIPublicV2CoursesCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV2CoursesCreated struct {
	Payload *models.CourseV2
}

func (o *PostLearnAPIPublicV2CoursesCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v2/courses][%d] postLearnApiPublicV2CoursesCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV2CoursesCreated) GetPayload() *models.CourseV2 {
	return o.Payload
}

func (o *PostLearnAPIPublicV2CoursesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CourseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV2CoursesBadRequest creates a PostLearnAPIPublicV2CoursesBadRequest with default headers values
func NewPostLearnAPIPublicV2CoursesBadRequest() *PostLearnAPIPublicV2CoursesBadRequest {
	return &PostLearnAPIPublicV2CoursesBadRequest{}
}

/*PostLearnAPIPublicV2CoursesBadRequest handles this case with default header values.

The request did not specify valid data
*/
type PostLearnAPIPublicV2CoursesBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV2CoursesBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v2/courses][%d] postLearnApiPublicV2CoursesBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV2CoursesBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV2CoursesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV2CoursesForbidden creates a PostLearnAPIPublicV2CoursesForbidden with default headers values
func NewPostLearnAPIPublicV2CoursesForbidden() *PostLearnAPIPublicV2CoursesForbidden {
	return &PostLearnAPIPublicV2CoursesForbidden{}
}

/*PostLearnAPIPublicV2CoursesForbidden handles this case with default header values.

The user does not have entitlements to create courses
*/
type PostLearnAPIPublicV2CoursesForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV2CoursesForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v2/courses][%d] postLearnApiPublicV2CoursesForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV2CoursesForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV2CoursesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV2CoursesConflict creates a PostLearnAPIPublicV2CoursesConflict with default headers values
func NewPostLearnAPIPublicV2CoursesConflict() *PostLearnAPIPublicV2CoursesConflict {
	return &PostLearnAPIPublicV2CoursesConflict{}
}

/*PostLearnAPIPublicV2CoursesConflict handles this case with default header values.

A course with the same courseId or externalId already exists
*/
type PostLearnAPIPublicV2CoursesConflict struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV2CoursesConflict) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v2/courses][%d] postLearnApiPublicV2CoursesConflict  %+v", 409, o.Payload)
}

func (o *PostLearnAPIPublicV2CoursesConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV2CoursesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV2CoursesBody post learn API public v2 courses body
swagger:model PostLearnAPIPublicV2CoursesBody
*/
type PostLearnAPIPublicV2CoursesBody struct {

	// Whether guests (users with the role guest) are allowed access to the course. Defaults to true.
	AllowGuests bool `json:"allowGuests,omitempty"`

	// availability
	Availability *PostLearnAPIPublicV2CoursesParamsBodyAvailability `json:"availability,omitempty"`

	// This status does not affect availability of the course for viewing in any way. closedComplete is valid for both Ultra and Classic courses. If an Ultra course is in closedComplete mode, updates are not possible. For a Classic course in closedComplete mode, updates are still possible (through Web UI but not through REST i.e. closed is enforced for original courses when updated through REST the same way Ultra courses are blocked) but new notifications are not generated.
	ClosedComplete bool `json:"closedComplete,omitempty"`

	// The Course ID attribute, shown to users in the UI.
	// Required: true
	// Max Length: 100
	CourseID *string `json:"courseId"`

	// The ID of the data source associated with this course. This may optionally be the data source's externalId using the syntax "externalId:math101".
	DataSourceID string `json:"dataSourceId,omitempty"`

	// The description of the course.
	Description string `json:"description,omitempty"`

	// enrollment
	Enrollment *PostLearnAPIPublicV2CoursesParamsBodyEnrollment `json:"enrollment,omitempty"`

	// An optional externally-defined unique ID for the course. Defaults to the courseId.
	//
	// Formerly known as 'batchUid'.
	// Max Length: 256
	ExternalID string `json:"externalId,omitempty"`

	// locale
	Locale *PostLearnAPIPublicV2CoursesParamsBodyLocale `json:"locale,omitempty"`

	// The name of the course.
	// Required: true
	// Max Length: 333
	Name *string `json:"name"`

	// Whether this object represents an Organization. Defaults to false.
	Organization bool `json:"organization,omitempty"`

	// The ID of the term associated to this course. This may optionally be the term's externalId using the syntax "externalId:spring.2016".
	TermID string `json:"termId,omitempty"`

	// Whether the course is rendered using Classic or Ultra Course View.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Undecided | The ultra status is not decided. |
	// | Classic | The course is decided as classic. |
	// | Ultra | The course is decided as ultra |
	// | UltraPreview | The course is currently in Ultra mode but during the preview state where it may still be reverted via a Restore to the classic state |
	//
	// Enum: [Undecided Classic Ultra UltraPreview]
	UltraStatus string `json:"ultraStatus,omitempty"`
}

// Validate validates this post learn API public v2 courses body
func (o *PostLearnAPIPublicV2CoursesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCourseID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnrollment(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUltraStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV2CoursesBody) validateAvailability(formats strfmt.Registry) error {

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

func (o *PostLearnAPIPublicV2CoursesBody) validateCourseID(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"courseId", "body", o.CourseID); err != nil {
		return err
	}

	if err := validate.MaxLength("input"+"."+"courseId", "body", string(*o.CourseID), 100); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV2CoursesBody) validateEnrollment(formats strfmt.Registry) error {

	if swag.IsZero(o.Enrollment) { // not required
		return nil
	}

	if o.Enrollment != nil {
		if err := o.Enrollment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "enrollment")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV2CoursesBody) validateExternalID(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalID) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"externalId", "body", string(o.ExternalID), 256); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV2CoursesBody) validateLocale(formats strfmt.Registry) error {

	if swag.IsZero(o.Locale) { // not required
		return nil
	}

	if o.Locale != nil {
		if err := o.Locale.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "locale")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV2CoursesBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("input"+"."+"name", "body", string(*o.Name), 333); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV2CoursesBodyTypeUltraStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Undecided","Classic","Ultra","UltraPreview"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV2CoursesBodyTypeUltraStatusPropEnum = append(postLearnApiPublicV2CoursesBodyTypeUltraStatusPropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV2CoursesBodyUltraStatusUndecided captures enum value "Undecided"
	PostLearnAPIPublicV2CoursesBodyUltraStatusUndecided string = "Undecided"

	// PostLearnAPIPublicV2CoursesBodyUltraStatusClassic captures enum value "Classic"
	PostLearnAPIPublicV2CoursesBodyUltraStatusClassic string = "Classic"

	// PostLearnAPIPublicV2CoursesBodyUltraStatusUltra captures enum value "Ultra"
	PostLearnAPIPublicV2CoursesBodyUltraStatusUltra string = "Ultra"

	// PostLearnAPIPublicV2CoursesBodyUltraStatusUltraPreview captures enum value "UltraPreview"
	PostLearnAPIPublicV2CoursesBodyUltraStatusUltraPreview string = "UltraPreview"
)

// prop value enum
func (o *PostLearnAPIPublicV2CoursesBody) validateUltraStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV2CoursesBodyTypeUltraStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV2CoursesBody) validateUltraStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.UltraStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateUltraStatusEnum("input"+"."+"ultraStatus", "body", o.UltraStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV2CoursesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV2CoursesParamsBodyAvailability Availability
//
// Settings controlling availability of the course to students.
swagger:model PostLearnAPIPublicV2CoursesParamsBodyAvailability
*/
type PostLearnAPIPublicV2CoursesParamsBodyAvailability struct {

	// Whether the course is currently available to students. Instructors can always access the course if they have 'Access unavailable course' entitlement. If set to 'Term', the course's parent term availability settings will be used.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes | Students may access the course. |
	// | No | Students may not access the course. |
	// | Disabled | Disabled by the SIS. Students may not access the course. @since 3100.0.0 |
	// | Term | Availability is inherited from the term settings. Requires a termId be set. |
	//
	// Enum: [Yes No Disabled Term]
	Available string `json:"available,omitempty"`

	// duration
	Duration *PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration `json:"duration,omitempty"`
}

// Validate validates this post learn API public v2 courses params body availability
func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postLearnApiPublicV2CoursesParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No","Disabled","Term"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV2CoursesParamsBodyAvailabilityTypeAvailablePropEnum = append(postLearnApiPublicV2CoursesParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV2CoursesParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PostLearnAPIPublicV2CoursesParamsBodyAvailabilityAvailableYes string = "Yes"

	// PostLearnAPIPublicV2CoursesParamsBodyAvailabilityAvailableNo captures enum value "No"
	PostLearnAPIPublicV2CoursesParamsBodyAvailabilityAvailableNo string = "No"

	// PostLearnAPIPublicV2CoursesParamsBodyAvailabilityAvailableDisabled captures enum value "Disabled"
	PostLearnAPIPublicV2CoursesParamsBodyAvailabilityAvailableDisabled string = "Disabled"

	// PostLearnAPIPublicV2CoursesParamsBodyAvailabilityAvailableTerm captures enum value "Term"
	PostLearnAPIPublicV2CoursesParamsBodyAvailabilityAvailableTerm string = "Term"
)

// prop value enum
func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV2CoursesParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(o.Available) { // not required
		return nil
	}

	// value enum
	if err := o.validateAvailableEnum("input"+"."+"availability"+"."+"available", "body", o.Available); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailability) validateDuration(formats strfmt.Registry) error {

	if swag.IsZero(o.Duration) { // not required
		return nil
	}

	if o.Duration != nil {
		if err := o.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "availability" + "." + "duration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV2CoursesParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration Duration
//
// Settings controlling the length of time the course is available.
swagger:model PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration
*/
type PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration struct {

	// The number of days this course can be used. May only be set if availability.duration.type is FixedNumDays.
	DaysOfUse int32 `json:"daysOfUse,omitempty"`

	// The date this course ends. May only be set if availability.duration.type is DateRange.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The date this course starts. May only be set if availability.duration.type is DateRange.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// The intended length of the course. Possible values are:
	//
	// - Continuous: The course is active on an ongoing basis. This is the default.
	// - DateRange: The course will only be available between specific date ranges.
	// - FixedNumDays: The course will only be available for a set number of days.
	// - Term: The course's parent term duration settings will be used.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Continuous | Course is active on an ongoing basis. |
	// | DateRange | Course is only intended to be available between specific date ranges |
	// | FixedNumDays | Course is only available for a set number of days |
	// | Term | Course availablity is dictated by its associated term |
	//
	// Enum: [Continuous DateRange FixedNumDays Term]
	Type string `json:"type,omitempty"`
}

// Validate validates this post learn API public v2 courses params body availability duration
func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(o.End) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"duration"+"."+"end", "body", "date-time", o.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(o.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"duration"+"."+"start", "body", "date-time", o.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV2CoursesParamsBodyAvailabilityDurationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Continuous","DateRange","FixedNumDays","Term"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV2CoursesParamsBodyAvailabilityDurationTypeTypePropEnum = append(postLearnApiPublicV2CoursesParamsBodyAvailabilityDurationTypeTypePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDurationTypeContinuous captures enum value "Continuous"
	PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDurationTypeContinuous string = "Continuous"

	// PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDurationTypeDateRange captures enum value "DateRange"
	PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDurationTypeDateRange string = "DateRange"

	// PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDurationTypeFixedNumDays captures enum value "FixedNumDays"
	PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDurationTypeFixedNumDays string = "FixedNumDays"

	// PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDurationTypeTerm captures enum value "Term"
	PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDurationTypeTerm string = "Term"
)

// prop value enum
func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV2CoursesParamsBodyAvailabilityDurationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("input"+"."+"availability"+"."+"duration"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV2CoursesParamsBodyAvailabilityDuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV2CoursesParamsBodyEnrollment Enrollment
//
// Settings controlling how students may enroll in the course.
swagger:model PostLearnAPIPublicV2CoursesParamsBodyEnrollment
*/
type PostLearnAPIPublicV2CoursesParamsBodyEnrollment struct {

	// The enrollment access code associated with this course. May only be set if enrollment.type is SelfEnrollment.
	// Max Length: 50
	AccessCode string `json:"accessCode,omitempty"`

	// The date on which enrollment in this course ends. May only be set if enrollment.type is SelfEnrollment.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The date on which enrollments are allowed for the course. May only be set if enrollment.type is SelfEnrollment.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// Specifies the enrollment options for the course. Defaults to InstructorLed.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | InstructorLed | Enrollment tasks for the course can only performed by the instructor |
	// | SelfEnrollment | Instructors have the ability to enroll users, and students can also enroll themselves in the course |
	// | EmailEnrollment | Instructors have the ability to enroll users, and students can email requests to the instructor for enrollment |
	//
	// Enum: [InstructorLed SelfEnrollment EmailEnrollment]
	Type string `json:"type,omitempty"`
}

// Validate validates this post learn API public v2 courses params body enrollment
func (o *PostLearnAPIPublicV2CoursesParamsBodyEnrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessCode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyEnrollment) validateAccessCode(formats strfmt.Registry) error {

	if swag.IsZero(o.AccessCode) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"enrollment"+"."+"accessCode", "body", string(o.AccessCode), 50); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyEnrollment) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(o.End) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"enrollment"+"."+"end", "body", "date-time", o.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyEnrollment) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(o.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"enrollment"+"."+"start", "body", "date-time", o.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV2CoursesParamsBodyEnrollmentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InstructorLed","SelfEnrollment","EmailEnrollment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV2CoursesParamsBodyEnrollmentTypeTypePropEnum = append(postLearnApiPublicV2CoursesParamsBodyEnrollmentTypeTypePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV2CoursesParamsBodyEnrollmentTypeInstructorLed captures enum value "InstructorLed"
	PostLearnAPIPublicV2CoursesParamsBodyEnrollmentTypeInstructorLed string = "InstructorLed"

	// PostLearnAPIPublicV2CoursesParamsBodyEnrollmentTypeSelfEnrollment captures enum value "SelfEnrollment"
	PostLearnAPIPublicV2CoursesParamsBodyEnrollmentTypeSelfEnrollment string = "SelfEnrollment"

	// PostLearnAPIPublicV2CoursesParamsBodyEnrollmentTypeEmailEnrollment captures enum value "EmailEnrollment"
	PostLearnAPIPublicV2CoursesParamsBodyEnrollmentTypeEmailEnrollment string = "EmailEnrollment"
)

// prop value enum
func (o *PostLearnAPIPublicV2CoursesParamsBodyEnrollment) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV2CoursesParamsBodyEnrollmentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyEnrollment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("input"+"."+"enrollment"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesParamsBodyEnrollment) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesParamsBodyEnrollment) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV2CoursesParamsBodyEnrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV2CoursesParamsBodyLocale Locale
//
// Settings controlling localization within the course.
swagger:model PostLearnAPIPublicV2CoursesParamsBodyLocale
*/
type PostLearnAPIPublicV2CoursesParamsBodyLocale struct {

	// Whether students are forced to use the course's specified locale.
	Force bool `json:"force,omitempty"`

	// The locale of this course.
	// Max Length: 20
	ID string `json:"id,omitempty"`
}

// Validate validates this post learn API public v2 courses params body locale
func (o *PostLearnAPIPublicV2CoursesParamsBodyLocale) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV2CoursesParamsBodyLocale) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"locale"+"."+"id", "body", string(o.ID), 20); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesParamsBodyLocale) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV2CoursesParamsBodyLocale) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV2CoursesParamsBodyLocale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
