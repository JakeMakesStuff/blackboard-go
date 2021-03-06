// Code generated by go-swagger; DO NOT EDIT.

package deprecated_course_grades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDReader is a Reader for the DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnID structure.
type DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK creates a DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK() *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK handles this case with default header values.

OK
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK struct {
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] deleteLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdOK ", 200)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest creates a DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest() *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] deleteLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden creates a DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden() *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden handles this case with default header values.

Insufficient Permission to delete the requested resource
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] deleteLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound creates a DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound() *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound handles this case with default header values.

Requested resource could not be found
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] deleteLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict creates a DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict with default headers values
func NewDeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict() *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict {
	return &DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict{}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict handles this case with default header values.

Conflict in deleting this grade column due to associated grades
*/
type DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}][%d] deleteLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
