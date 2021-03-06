// Code generated by go-swagger; DO NOT EDIT.

package attendance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserID structure.
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK() *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK struct {
	Payload *models.AttendanceRecord
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK) GetPayload() *models.AttendanceRecord {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttendanceRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest() *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest handles this case with default header values.

Failed to get attendance record; or The request did not specify a valid User/Meeting Id
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound creates a GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound() *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound handles this case with default header values.

The attendance record is not found
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/meetings/{meetingId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdMeetingsMeetingIdUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
