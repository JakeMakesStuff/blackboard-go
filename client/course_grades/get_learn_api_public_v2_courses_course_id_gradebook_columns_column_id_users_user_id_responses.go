// Code generated by go-swagger; DO NOT EDIT.

package course_grades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDReader is a Reader for the GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserID structure.
type GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK creates a GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK() *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK {
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK struct {
	Payload *models.GradeV2
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/gradebook/columns/{columnId}/users/{userId}][%d] getLearnApiPublicV2CoursesCourseIdGradebookColumnsColumnIdUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK) GetPayload() *models.GradeV2 {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GradeV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest creates a GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest() *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest {
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/gradebook/columns/{columnId}/users/{userId}][%d] getLearnApiPublicV2CoursesCourseIdGradebookColumnsColumnIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden creates a GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden() *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden {
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/gradebook/columns/{columnId}/users/{userId}][%d] getLearnApiPublicV2CoursesCourseIdGradebookColumnsColumnIdUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound creates a GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound() *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound {
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound handles this case with default header values.

Grade detail values have not yet been populated
*/
type GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/gradebook/columns/{columnId}/users/{userId}][%d] getLearnApiPublicV2CoursesCourseIdGradebookColumnsColumnIdUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
