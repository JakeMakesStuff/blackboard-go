// Code generated by go-swagger; DO NOT EDIT.

package course_grades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDReader is a Reader for the PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaID structure.
type PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK creates a PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK() *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK {
	return &PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK handles this case with default header values.

OK
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK struct {
	Payload *models.GradeSchema
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/gradebook/schemas/{schemaId}][%d] patchLearnApiPublicV1CoursesCourseIdGradebookSchemasSchemaIdOK  %+v", 200, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK) GetPayload() *models.GradeSchema {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GradeSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest creates a PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest() *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest {
	return &PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest handles this case with default header values.

Bad Request
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/gradebook/schemas/{schemaId}][%d] patchLearnApiPublicV1CoursesCourseIdGradebookSchemasSchemaIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden creates a PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden() *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden {
	return &PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden handles this case with default header values.

Forbidden
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/gradebook/schemas/{schemaId}][%d] patchLearnApiPublicV1CoursesCourseIdGradebookSchemasSchemaIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound creates a PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound with default headers values
func NewPatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound() *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound {
	return &PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound{}
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound handles this case with default header values.

Not Found
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound struct {
	Payload *models.RestException
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /learn/api/public/v1/courses/{courseId}/gradebook/schemas/{schemaId}][%d] patchLearnApiPublicV1CoursesCourseIdGradebookSchemasSchemaIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBody patch learn API public v1 courses course ID gradebook schemas schema ID body
swagger:model PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBody
*/
type PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBody struct {

	// The description of this grade schema.
	Description string `json:"description,omitempty"`

	// The externalId associated with this grade schema.
	ExternalID string `json:"externalId,omitempty"`

	// The list of grade symbols for this grade schema. Only returned for Tabular scaleType schemas.
	Symbols []*models.GradeSymbol `json:"symbols"`

	// The title of this grade schema.
	Title string `json:"title,omitempty"`
}

// Validate validates this patch learn API public v1 courses course ID gradebook schemas schema ID body
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSymbols(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBody) validateSymbols(formats strfmt.Registry) error {

	if swag.IsZero(o.Symbols) { // not required
		return nil
	}

	for i := 0; i < len(o.Symbols); i++ {
		if swag.IsZero(o.Symbols[i]) { // not required
			continue
		}

		if o.Symbols[i] != nil {
			if err := o.Symbols[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input" + "." + "symbols" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBody) UnmarshalBinary(b []byte) error {
	var res PatchLearnAPIPublicV1CoursesCourseIDGradebookSchemasSchemaIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
