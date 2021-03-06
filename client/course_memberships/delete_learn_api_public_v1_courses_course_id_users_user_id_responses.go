// Code generated by go-swagger; DO NOT EDIT.

package course_memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDReader is a Reader for the DeleteLearnAPIPublicV1CoursesCourseIDUsersUserID structure.
type DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent creates a DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent() *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent {
	return &DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent handles this case with default header values.

No Content
*/
type DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent struct {
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] deleteLearnApiPublicV1CoursesCourseIdUsersUserIdNoContent ", 204)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest creates a DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest() *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest {
	return &DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] deleteLearnApiPublicV1CoursesCourseIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden creates a DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden() *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden {
	return &DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] deleteLearnApiPublicV1CoursesCourseIdUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound creates a DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound() *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound {
	return &DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/users/{userId}][%d] deleteLearnApiPublicV1CoursesCourseIdUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
