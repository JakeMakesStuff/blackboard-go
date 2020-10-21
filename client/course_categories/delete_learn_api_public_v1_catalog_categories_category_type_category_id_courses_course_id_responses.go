// Code generated by go-swagger; DO NOT EDIT.

package course_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDReader is a Reader for the DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseID structure.
type DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent creates a DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent with default headers values
func NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent() *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent {
	return &DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent{}
}

/*DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent handles this case with default header values.

No Content
*/
type DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent struct {
}

func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/catalog/categories/{categoryType}/{categoryId}/courses/{courseId}][%d] deleteLearnApiPublicV1CatalogCategoriesCategoryTypeCategoryIdCoursesCourseIdNoContent ", 204)
}

func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden creates a DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden() *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden {
	return &DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden{}
}

/*DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/catalog/categories/{categoryType}/{categoryId}/courses/{courseId}][%d] deleteLearnApiPublicV1CatalogCategoriesCategoryTypeCategoryIdCoursesCourseIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound creates a DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound() *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound {
	return &DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound{}
}

/*DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/catalog/categories/{categoryType}/{categoryId}/courses/{courseId}][%d] deleteLearnApiPublicV1CatalogCategoriesCategoryTypeCategoryIdCoursesCourseIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1CatalogCategoriesCategoryTypeCategoryIDCoursesCourseIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}