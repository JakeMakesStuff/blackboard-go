// Code generated by go-swagger; DO NOT EDIT.

package deprecated_courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1CoursesCourseIDReader is a Reader for the GetLearnAPIPublicV1CoursesCourseID structure.
type GetLearnAPIPublicV1CoursesCourseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDOK creates a GetLearnAPIPublicV1CoursesCourseIDOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDOK() *GetLearnAPIPublicV1CoursesCourseIDOK {
	return &GetLearnAPIPublicV1CoursesCourseIDOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDOK struct {
	Payload *models.Course
}

func (o *GetLearnAPIPublicV1CoursesCourseIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}][%d] getLearnApiPublicV1CoursesCourseIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDOK) GetPayload() *models.Course {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Course)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDBadRequest() *GetLearnAPIPublicV1CoursesCourseIDBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDBadRequest handles this case with default header values.

Failed to retrieve course; or The request did not specify a valid courseId
*/
type GetLearnAPIPublicV1CoursesCourseIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}][%d] getLearnApiPublicV1CoursesCourseIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDNotFound creates a GetLearnAPIPublicV1CoursesCourseIDNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDNotFound() *GetLearnAPIPublicV1CoursesCourseIDNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDNotFound handles this case with default header values.

The course is not found or is unavailable and the user does not have the permission to view unavailable courses
*/
type GetLearnAPIPublicV1CoursesCourseIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}][%d] getLearnApiPublicV1CoursesCourseIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
