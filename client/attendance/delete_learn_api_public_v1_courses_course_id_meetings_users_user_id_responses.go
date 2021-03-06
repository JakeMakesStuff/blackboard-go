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

// DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDReader is a Reader for the DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserID structure.
type DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent creates a DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent() *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent {
	return &DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent handles this case with default header values.

No Content
*/
type DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent struct {
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/meetings/users/{userId}][%d] deleteLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdNoContent ", 204)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest creates a DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest() *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest {
	return &DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest handles this case with default header values.

Invalid user id provided
*/
type DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/meetings/users/{userId}][%d] deleteLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden creates a DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden() *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden {
	return &DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden handles this case with default header values.

The user is not authorized to delete the specified records in meeting object
*/
type DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/meetings/users/{userId}][%d] deleteLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound creates a DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound() *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound {
	return &DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/meetings/users/{userId}][%d] deleteLearnApiPublicV1CoursesCourseIdMeetingsUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
