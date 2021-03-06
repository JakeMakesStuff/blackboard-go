// Code generated by go-swagger; DO NOT EDIT.

package course_grades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasReader is a Reader for the PostLearnAPIPublicV1CoursesCourseIDGradebookSchemas structure.
type PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated creates a PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated() *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated {
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated struct {
	Payload *models.GradeSchema
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/gradebook/schemas][%d] postLearnApiPublicV1CoursesCourseIdGradebookSchemasCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated) GetPayload() *models.GradeSchema {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GradeSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest creates a PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest() *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest {
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest handles this case with default header values.

Bad Request
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/gradebook/schemas][%d] postLearnApiPublicV1CoursesCourseIdGradebookSchemasBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden creates a PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden() *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden {
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden{}
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden handles this case with default header values.

Forbidden
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/gradebook/schemas][%d] postLearnApiPublicV1CoursesCourseIdGradebookSchemasForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody post learn API public v1 courses course ID gradebook schemas body
swagger:model PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody struct {

	// The description of this grade schema.
	Description string `json:"description,omitempty"`

	// The externalId associated with this grade schema.
	ExternalID string `json:"externalId,omitempty"`

	// The scale type of this grade schema.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Percent |  |
	// | Score |  |
	// | Tabular |  |
	// | Text |  |
	// | CompleteIncomplete |  |
	//
	// Enum: [Percent Score Tabular Text CompleteIncomplete]
	ScaleType string `json:"scaleType,omitempty"`

	// The list of grade symbols for this grade schema. Only returned for Tabular scaleType schemas.
	Symbols []*models.GradeSymbol `json:"symbols"`

	// The title of this grade schema.
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this post learn API public v1 courses course ID gradebook schemas body
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateScaleType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSymbols(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postLearnApiPublicV1CoursesCourseIdGradebookSchemasBodyTypeScaleTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Percent","Score","Tabular","Text","CompleteIncomplete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdGradebookSchemasBodyTypeScaleTypePropEnum = append(postLearnApiPublicV1CoursesCourseIdGradebookSchemasBodyTypeScaleTypePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypePercent captures enum value "Percent"
	PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypePercent string = "Percent"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypeScore captures enum value "Score"
	PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypeScore string = "Score"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypeTabular captures enum value "Tabular"
	PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypeTabular string = "Tabular"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypeText captures enum value "Text"
	PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypeText string = "Text"

	// PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypeCompleteIncomplete captures enum value "CompleteIncomplete"
	PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBodyScaleTypeCompleteIncomplete string = "CompleteIncomplete"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody) validateScaleTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdGradebookSchemasBodyTypeScaleTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody) validateScaleType(formats strfmt.Registry) error {

	if swag.IsZero(o.ScaleType) { // not required
		return nil
	}

	// value enum
	if err := o.validateScaleTypeEnum("gradeSchemaTOPubV1"+"."+"scaleType", "body", o.ScaleType); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody) validateSymbols(formats strfmt.Registry) error {

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
					return ve.ValidateName("gradeSchemaTOPubV1" + "." + "symbols" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("gradeSchemaTOPubV1"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDGradebookSchemasBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
