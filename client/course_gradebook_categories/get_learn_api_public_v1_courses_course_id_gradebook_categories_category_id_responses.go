// Code generated by go-swagger; DO NOT EDIT.

package course_gradebook_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryID structure.
type GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK creates a GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK() *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK struct {
	Payload *models.GradebookCategory
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/categories/{categoryId}][%d] getLearnApiPublicV1CoursesCourseIdGradebookCategoriesCategoryIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK) GetPayload() *models.GradebookCategory {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GradebookCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest() *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/categories/{categoryId}][%d] getLearnApiPublicV1CoursesCourseIdGradebookCategoriesCategoryIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden creates a GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden() *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/categories/{categoryId}][%d] getLearnApiPublicV1CoursesCourseIdGradebookCategoriesCategoryIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound creates a GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound() *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/categories/{categoryId}][%d] getLearnApiPublicV1CoursesCourseIdGradebookCategoriesCategoryIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}