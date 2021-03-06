// Code generated by go-swagger; DO NOT EDIT.

package courses

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

// PatchLearnAPIPublicV3CoursesCourseIDReader is a Reader for the PatchLearnAPIPublicV3CoursesCourseID structure.
type PatchLearnAPIPublicV3CoursesCourseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLearnAPIPublicV3CoursesCourseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLearnAPIPublicV3CoursesCourseIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLearnAPIPublicV3CoursesCourseIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLearnAPIPublicV3CoursesCourseIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLearnAPIPublicV3CoursesCourseIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchLearnAPIPublicV3CoursesCourseIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLearnAPIPublicV3CoursesCourseIDOK creates a PatchLearnAPIPublicV3CoursesCourseIDOK with default headers values
func NewPatchLearnAPIPublicV3CoursesCourseIDOK() *PatchLearnAPIPublicV3CoursesCourseIDOK {
	return &PatchLearnAPIPublicV3CoursesCourseIDOK{}
}

/*PatchLearnAPIPublicV3CoursesCourseIDOK handles this case with default header values.

OK
*/
type PatchLearnAPIPublicV3CoursesCourseIDOK struct {
	Payload *models.CourseV2
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDOK) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v3/courses/{courseId}][%d] patchLearnApiPublicV3CoursesCourseIdOK  %+v", 200, o.Payload)
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDOK) GetPayload() *models.CourseV2 {
	return o.Payload
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CourseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV3CoursesCourseIDBadRequest creates a PatchLearnAPIPublicV3CoursesCourseIDBadRequest with default headers values
func NewPatchLearnAPIPublicV3CoursesCourseIDBadRequest() *PatchLearnAPIPublicV3CoursesCourseIDBadRequest {
	return &PatchLearnAPIPublicV3CoursesCourseIDBadRequest{}
}

/*PatchLearnAPIPublicV3CoursesCourseIDBadRequest handles this case with default header values.

The request did not specify a valid course
*/
type PatchLearnAPIPublicV3CoursesCourseIDBadRequest struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v3/courses/{courseId}][%d] patchLearnApiPublicV3CoursesCourseIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV3CoursesCourseIDForbidden creates a PatchLearnAPIPublicV3CoursesCourseIDForbidden with default headers values
func NewPatchLearnAPIPublicV3CoursesCourseIDForbidden() *PatchLearnAPIPublicV3CoursesCourseIDForbidden {
	return &PatchLearnAPIPublicV3CoursesCourseIDForbidden{}
}

/*PatchLearnAPIPublicV3CoursesCourseIDForbidden handles this case with default header values.

The user does not have entitlements to modify courses
*/
type PatchLearnAPIPublicV3CoursesCourseIDForbidden struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v3/courses/{courseId}][%d] patchLearnApiPublicV3CoursesCourseIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV3CoursesCourseIDNotFound creates a PatchLearnAPIPublicV3CoursesCourseIDNotFound with default headers values
func NewPatchLearnAPIPublicV3CoursesCourseIDNotFound() *PatchLearnAPIPublicV3CoursesCourseIDNotFound {
	return &PatchLearnAPIPublicV3CoursesCourseIDNotFound{}
}

/*PatchLearnAPIPublicV3CoursesCourseIDNotFound handles this case with default header values.

The course was not found or is unavailable and the user does not have the permission to view unavailable courses
*/
type PatchLearnAPIPublicV3CoursesCourseIDNotFound struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v3/courses/{courseId}][%d] patchLearnApiPublicV3CoursesCourseIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV3CoursesCourseIDConflict creates a PatchLearnAPIPublicV3CoursesCourseIDConflict with default headers values
func NewPatchLearnAPIPublicV3CoursesCourseIDConflict() *PatchLearnAPIPublicV3CoursesCourseIDConflict {
	return &PatchLearnAPIPublicV3CoursesCourseIDConflict{}
}

/*PatchLearnAPIPublicV3CoursesCourseIDConflict handles this case with default header values.

A course with the same courseId or externalId already exists
*/
type PatchLearnAPIPublicV3CoursesCourseIDConflict struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v3/courses/{courseId}][%d] patchLearnApiPublicV3CoursesCourseIdConflict  %+v", 409, o.Payload)
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchLearnAPIPublicV3CoursesCourseIDBody patch learn API public v3 courses course ID body
swagger:model PatchLearnAPIPublicV3CoursesCourseIDBody
*/
type PatchLearnAPIPublicV3CoursesCourseIDBody struct {

	// Whether guests (users with the role guest) are allowed access to the course. Defaults to true.
	AllowGuests bool `json:"allowGuests,omitempty"`

	// availability
	Availability *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability `json:"availability,omitempty"`

	// This status does not affect availability of the course for viewing in any way. closedComplete is valid for both Ultra and Classic courses. If an Ultra course is in closedComplete mode, updates are not possible. For a Classic course in closedComplete mode, updates are still possible (through Web UI but not through REST i.e. closed is enforced for original courses when updated through REST the same way Ultra courses are blocked) but new notifications are not generated.
	ClosedComplete bool `json:"closedComplete,omitempty"`

	// The ID of the data source associated with this course. This may optionally be the data source's externalId using the syntax "externalId:math101".
	DataSourceID string `json:"dataSourceId,omitempty"`

	// The description of the course.
	Description string `json:"description,omitempty"`

	// enrollment
	Enrollment *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment `json:"enrollment,omitempty"`

	// An optional externally-defined unique ID for the course. Defaults to the courseId.
	//
	// Formerly known as 'batchUid'.
	// Max Length: 256
	ExternalID string `json:"externalId,omitempty"`

	// locale
	Locale *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale `json:"locale,omitempty"`

	// The name of the course.
	// Max Length: 333
	Name string `json:"name,omitempty"`

	// The ID of the term associated to this course. This may optionally be the term's externalId using the syntax "externalId:spring.2016".
	TermID string `json:"termId,omitempty"`
}

// Validate validates this patch learn API public v3 courses course ID body
func (o *PatchLearnAPIPublicV3CoursesCourseIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDBody) validateAvailability(formats strfmt.Registry) error {

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

func (o *PatchLearnAPIPublicV3CoursesCourseIDBody) validateEnrollment(formats strfmt.Registry) error {

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

func (o *PatchLearnAPIPublicV3CoursesCourseIDBody) validateExternalID(formats strfmt.Registry) error {

	if swag.IsZero(o.ExternalID) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"externalId", "body", string(o.ExternalID), 256); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDBody) validateLocale(formats strfmt.Registry) error {

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

func (o *PatchLearnAPIPublicV3CoursesCourseIDBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"name", "body", string(o.Name), 333); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV3CoursesCourseIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV3CoursesCourseIDBody) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV3CoursesCourseIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability Availability
//
// Settings controlling availability of the course to students.
swagger:model PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability
*/
type PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability struct {

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
	Duration *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration `json:"duration,omitempty"`
}

// Validate validates this patch learn API public v3 courses course ID params body availability
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability) Validate(formats strfmt.Registry) error {
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

var patchLearnApiPublicV3CoursesCourseIdParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No","Disabled","Term"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV3CoursesCourseIdParamsBodyAvailabilityTypeAvailablePropEnum = append(patchLearnApiPublicV3CoursesCourseIdParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityAvailableYes string = "Yes"

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityAvailableNo captures enum value "No"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityAvailableNo string = "No"

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityAvailableDisabled captures enum value "Disabled"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityAvailableDisabled string = "Disabled"

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityAvailableTerm captures enum value "Term"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityAvailableTerm string = "Term"
)

// prop value enum
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV3CoursesCourseIdParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(o.Available) { // not required
		return nil
	}

	// value enum
	if err := o.validateAvailableEnum("input"+"."+"availability"+"."+"available", "body", o.Available); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability) validateDuration(formats strfmt.Registry) error {

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
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration Duration
//
// Settings controlling the length of time the course is available.
swagger:model PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration
*/
type PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration struct {

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

// Validate validates this patch learn API public v3 courses course ID params body availability duration
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration) Validate(formats strfmt.Registry) error {
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

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(o.End) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"duration"+"."+"end", "body", "date-time", o.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(o.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"duration"+"."+"start", "body", "date-time", o.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var patchLearnApiPublicV3CoursesCourseIdParamsBodyAvailabilityDurationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Continuous","DateRange","FixedNumDays","Term"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV3CoursesCourseIdParamsBodyAvailabilityDurationTypeTypePropEnum = append(patchLearnApiPublicV3CoursesCourseIdParamsBodyAvailabilityDurationTypeTypePropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDurationTypeContinuous captures enum value "Continuous"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDurationTypeContinuous string = "Continuous"

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDurationTypeDateRange captures enum value "DateRange"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDurationTypeDateRange string = "DateRange"

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDurationTypeFixedNumDays captures enum value "FixedNumDays"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDurationTypeFixedNumDays string = "FixedNumDays"

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDurationTypeTerm captures enum value "Term"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDurationTypeTerm string = "Term"
)

// prop value enum
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV3CoursesCourseIdParamsBodyAvailabilityDurationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration) validateType(formats strfmt.Registry) error {

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
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV3CoursesCourseIDParamsBodyAvailabilityDuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment Enrollment
//
// Settings controlling how students may enroll in the course.
swagger:model PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment
*/
type PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment struct {

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

// Validate validates this patch learn API public v3 courses course ID params body enrollment
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment) Validate(formats strfmt.Registry) error {
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

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment) validateAccessCode(formats strfmt.Registry) error {

	if swag.IsZero(o.AccessCode) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"enrollment"+"."+"accessCode", "body", string(o.AccessCode), 50); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(o.End) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"enrollment"+"."+"end", "body", "date-time", o.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(o.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"enrollment"+"."+"start", "body", "date-time", o.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var patchLearnApiPublicV3CoursesCourseIdParamsBodyEnrollmentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InstructorLed","SelfEnrollment","EmailEnrollment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV3CoursesCourseIdParamsBodyEnrollmentTypeTypePropEnum = append(patchLearnApiPublicV3CoursesCourseIdParamsBodyEnrollmentTypeTypePropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollmentTypeInstructorLed captures enum value "InstructorLed"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollmentTypeInstructorLed string = "InstructorLed"

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollmentTypeSelfEnrollment captures enum value "SelfEnrollment"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollmentTypeSelfEnrollment string = "SelfEnrollment"

	// PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollmentTypeEmailEnrollment captures enum value "EmailEnrollment"
	PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollmentTypeEmailEnrollment string = "EmailEnrollment"
)

// prop value enum
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV3CoursesCourseIdParamsBodyEnrollmentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment) validateType(formats strfmt.Registry) error {

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
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV3CoursesCourseIDParamsBodyEnrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale Locale
//
// Settings controlling localization within the course.
swagger:model PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale
*/
type PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale struct {

	// Whether students are forced to use the course's specified locale.
	Force bool `json:"force,omitempty"`

	// The locale of this course.
	// Max Length: 20
	ID string `json:"id,omitempty"`
}

// Validate validates this patch learn API public v3 courses course ID params body locale
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"locale"+"."+"id", "body", string(o.ID), 20); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV3CoursesCourseIDParamsBodyLocale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
