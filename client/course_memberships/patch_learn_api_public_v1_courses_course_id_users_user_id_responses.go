// Code generated by go-swagger; DO NOT EDIT.

package course_memberships

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

// PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDReader is a Reader for the PatchLearnAPIPublicV1CoursesCourseIDUsersUserID structure.
type PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK creates a PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK() *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK {
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK handles this case with default header values.

OK
*/
type PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK struct {
	Payload *models.CourseMembership
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK) GetPayload() *models.CourseMembership {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CourseMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest creates a PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest() *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest {
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest handles this case with default header values.

The request did not specify valid data
*/
type PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden creates a PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden() *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden {
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden handles this case with default header values.

User has insufficient privileges
*/
type PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound creates a PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound() *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound {
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound handles this case with default header values.

Course not found; or Course-membership not found
*/
type PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict creates a PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict() *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict {
	return &PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict handles this case with default header values.

Conflict
*/
type PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] patchLearnApiPublicV1CoursesCourseIdUsersUserIdConflict  %+v", 409, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody patch learn API public v1 courses course ID users user ID body
swagger:model PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody
*/
type PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody struct {

	// Settings controlling availability of the course membership.
	Availability *models.Availability `json:"availability,omitempty"`

	// The primary ID of the child, cross-listed course, in which the user is directly enrolled. </p> This field is read only in Learn versions 3000.11.0 through 3400.0.0. As of 3400.1.0, this field is mutable.  </p> If this membership's course is a parent course in a cross-listed set, the childCourseId can be updated to move the membership enrollment between child courses and the parent course in  the set.  Patching the childCourseId to "null" will move the membership to the parent course.
	//
	// **Since**: 3000.11.0
	ChildCourseID string `json:"childCourseId,omitempty"`

	// The user's role in the course.
	//
	// These roles are also valid for an organization, although they are named differently in the UI.
	//
	// Custom course roles may also be referenced by their IDs.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Instructor | Has access to all areas in the Control Panel. This role is generally given to those developing, teaching, or facilitating the class. Instructors may access a course that is unavailable to students. This role is customizable and may have different capabilities from what is documented here. |
	// | BbFacilitator | The facilitator is an instructor like role. Facilitators are restricted versions of an instructor, in that they are able to deliver course instruction and administer all aspects of a pre-constructed course, but are not allowed to modify or alter the course. This role is customizable and may have different capabilities from what is documented here. |
	// | TeachingAssistant | The teaching assistant role is that of a co-teacher. Teaching assistants are able to administer all areas of a course. Their only limitations are those imposed by the instructor or Blackboard administrator at your school. This role is customizable and may have different capabilities from what is documented here. |
	// | CourseBuilder | Manages the course without having access to student grades. This role is customizable and may have different capabilities from what is documented here. |
	// | Grader | Assists the instructor in the creation, management, delivery, and grading of items. This role is customizable and may have different capabilities from what is documented here. |
	// | Student |  |
	// | Guest | Has no access to the Control Panel. Areas within the course are made available to guests, but typically they can only view course materials; they do not have access to tests or assessments, and do not have permission to post on discussion boards. This role's behavior is immutable. |
	//
	// Enum: [Instructor BbFacilitator TeachingAssistant CourseBuilder Grader Student Guest]
	CourseRoleID string `json:"courseRoleId,omitempty"`

	// The ID of the data source associated with this course.  This may optionally be the data source's externalId using the syntax "externalId:math101".
	DataSourceID string `json:"dataSourceId,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID users user ID body
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCourseRoleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody) validateAvailability(formats strfmt.Registry) error {

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

var patchLearnApiPublicV1CoursesCourseIdUsersUserIdBodyTypeCourseRoleIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Instructor","BbFacilitator","TeachingAssistant","CourseBuilder","Grader","Student","Guest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchLearnApiPublicV1CoursesCourseIdUsersUserIdBodyTypeCourseRoleIDPropEnum = append(patchLearnApiPublicV1CoursesCourseIdUsersUserIdBodyTypeCourseRoleIDPropEnum, v)
	}
}

const (

	// PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDInstructor captures enum value "Instructor"
	PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDInstructor string = "Instructor"

	// PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDBbFacilitator captures enum value "BbFacilitator"
	PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDBbFacilitator string = "BbFacilitator"

	// PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDTeachingAssistant captures enum value "TeachingAssistant"
	PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDTeachingAssistant string = "TeachingAssistant"

	// PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDCourseBuilder captures enum value "CourseBuilder"
	PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDCourseBuilder string = "CourseBuilder"

	// PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDGrader captures enum value "Grader"
	PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDGrader string = "Grader"

	// PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDStudent captures enum value "Student"
	PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDStudent string = "Student"

	// PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDGuest captures enum value "Guest"
	PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBodyCourseRoleIDGuest string = "Guest"
)

// prop value enum
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody) validateCourseRoleIDEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchLearnApiPublicV1CoursesCourseIdUsersUserIdBodyTypeCourseRoleIDPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody) validateCourseRoleID(formats strfmt.Registry) error {

	if swag.IsZero(o.CourseRoleID) { // not required
		return nil
	}

	// value enum
	if err := o.validateCourseRoleIDEnum("input"+"."+"courseRoleId", "body", o.CourseRoleID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDUsersUserIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}