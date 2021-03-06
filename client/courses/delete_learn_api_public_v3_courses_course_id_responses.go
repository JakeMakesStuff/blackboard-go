// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV3CoursesCourseIDReader is a Reader for the DeleteLearnAPIPublicV3CoursesCourseID structure.
type DeleteLearnAPIPublicV3CoursesCourseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV3CoursesCourseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteLearnAPIPublicV3CoursesCourseIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearnAPIPublicV3CoursesCourseIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearnAPIPublicV3CoursesCourseIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV3CoursesCourseIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV3CoursesCourseIDAccepted creates a DeleteLearnAPIPublicV3CoursesCourseIDAccepted with default headers values
func NewDeleteLearnAPIPublicV3CoursesCourseIDAccepted() *DeleteLearnAPIPublicV3CoursesCourseIDAccepted {
	return &DeleteLearnAPIPublicV3CoursesCourseIDAccepted{}
}

/*DeleteLearnAPIPublicV3CoursesCourseIDAccepted handles this case with default header values.

Accepted
*/
type DeleteLearnAPIPublicV3CoursesCourseIDAccepted struct {
	/*A URI to query the status of the corresponding delete operation
	 */
	Location string
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDAccepted) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v3/courses/{courseId}][%d] deleteLearnApiPublicV3CoursesCourseIdAccepted ", 202)
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewDeleteLearnAPIPublicV3CoursesCourseIDBadRequest creates a DeleteLearnAPIPublicV3CoursesCourseIDBadRequest with default headers values
func NewDeleteLearnAPIPublicV3CoursesCourseIDBadRequest() *DeleteLearnAPIPublicV3CoursesCourseIDBadRequest {
	return &DeleteLearnAPIPublicV3CoursesCourseIDBadRequest{}
}

/*DeleteLearnAPIPublicV3CoursesCourseIDBadRequest handles this case with default header values.

Invalid courseId provided
*/
type DeleteLearnAPIPublicV3CoursesCourseIDBadRequest struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v3/courses/{courseId}][%d] deleteLearnApiPublicV3CoursesCourseIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV3CoursesCourseIDForbidden creates a DeleteLearnAPIPublicV3CoursesCourseIDForbidden with default headers values
func NewDeleteLearnAPIPublicV3CoursesCourseIDForbidden() *DeleteLearnAPIPublicV3CoursesCourseIDForbidden {
	return &DeleteLearnAPIPublicV3CoursesCourseIDForbidden{}
}

/*DeleteLearnAPIPublicV3CoursesCourseIDForbidden handles this case with default header values.

The user is not authorized to delete the specified Course object
*/
type DeleteLearnAPIPublicV3CoursesCourseIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v3/courses/{courseId}][%d] deleteLearnApiPublicV3CoursesCourseIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV3CoursesCourseIDNotFound creates a DeleteLearnAPIPublicV3CoursesCourseIDNotFound with default headers values
func NewDeleteLearnAPIPublicV3CoursesCourseIDNotFound() *DeleteLearnAPIPublicV3CoursesCourseIDNotFound {
	return &DeleteLearnAPIPublicV3CoursesCourseIDNotFound{}
}

/*DeleteLearnAPIPublicV3CoursesCourseIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV3CoursesCourseIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v3/courses/{courseId}][%d] deleteLearnApiPublicV3CoursesCourseIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV3CoursesCourseIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
