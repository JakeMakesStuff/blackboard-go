// Code generated by go-swagger; DO NOT EDIT.

package attendance

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

// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersReader is a Reader for the PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsers structure.
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated creates a PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated() *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated {
	return &PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated{}
}

/*PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated struct {
	Payload *models.AttendanceRecord
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users][%d] postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated) GetPayload() *models.AttendanceRecord {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttendanceRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest creates a PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest() *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest {
	return &PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest{}
}

/*PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest handles this case with default header values.

The request did not specify valid data
*/
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users][%d] postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden creates a PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden() *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden {
	return &PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden{}
}

/*PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden handles this case with default header values.

The user does not have entitlements to create attendace records
*/
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users][%d] postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody post learn API public v1 courses course ID meetings meeting ID users body
swagger:model PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody
*/
type PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody struct {

	// The primary id of the meeting.
	// Required: true
	MeetingID *string `json:"meetingId"`

	// The attendance status of the user.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Absent |  |
	// | Late |  |
	// | Present |  |
	// | Excused |  |
	//
	// Required: true
	// Enum: [Absent Late Present Excused]
	Status *string `json:"status"`

	// The learn external id of the user.
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this post learn API public v1 courses course ID meetings meeting ID users body
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMeetingID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody) validateMeetingID(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"meetingId", "body", o.MeetingID); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Absent","Late","Present","Excused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBodyTypeStatusPropEnum = append(postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBodyTypeStatusPropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBodyStatusAbsent captures enum value "Absent"
	PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBodyStatusAbsent string = "Absent"

	// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBodyStatusLate captures enum value "Late"
	PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBodyStatusLate string = "Late"

	// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBodyStatusPresent captures enum value "Present"
	PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBodyStatusPresent string = "Present"

	// PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBodyStatusExcused captures enum value "Excused"
	PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBodyStatusExcused string = "Excused"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	// value enum
	if err := o.validateStatusEnum("input"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"userId", "body", o.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
