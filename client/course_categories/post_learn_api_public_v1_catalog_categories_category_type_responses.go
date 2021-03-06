// Code generated by go-swagger; DO NOT EDIT.

package course_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PostLearnAPIPublicV1CatalogCategoriesCategoryTypeReader is a Reader for the PostLearnAPIPublicV1CatalogCategoriesCategoryType structure.
type PostLearnAPIPublicV1CatalogCategoriesCategoryTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated creates a PostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated with default headers values
func NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated() *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated {
	return &PostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated{}
}

/*PostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated struct {
	Payload *models.Category
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/catalog/categories/{categoryType}][%d] postLearnApiPublicV1CatalogCategoriesCategoryTypeCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated) GetPayload() *models.Category {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Category)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest creates a PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest with default headers values
func NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest() *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest {
	return &PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest{}
}

/*PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest handles this case with default header values.

The request did not specify valid data
*/
type PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/catalog/categories/{categoryType}][%d] postLearnApiPublicV1CatalogCategoriesCategoryTypeBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden creates a PostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden with default headers values
func NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden() *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden {
	return &PostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden{}
}

/*PostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden handles this case with default header values.

The user does not have entitlements to create categories of the provided type
*/
type PostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/catalog/categories/{categoryType}][%d] postLearnApiPublicV1CatalogCategoriesCategoryTypeForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict creates a PostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict with default headers values
func NewPostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict() *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict {
	return &PostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict{}
}

/*PostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict handles this case with default header values.

The request specified data which is in conflict with existing data
*/
type PostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/catalog/categories/{categoryType}][%d] postLearnApiPublicV1CatalogCategoriesCategoryTypeConflict  %+v", 409, o.Payload)
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody post learn API public v1 catalog categories category type body
swagger:model PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody
*/
type PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody struct {

	// Boolean indicating whether or not the category is available
	// Required: true
	Available *bool `json:"available"`

	// The human-readable id of the category
	// Required: true
	CategoryID *string `json:"categoryId"`

	// The description of the category
	Description string `json:"description,omitempty"`

	// Boolean indicating whether or not the category should appear on the catalog front page
	FrontPage bool `json:"frontPage,omitempty"`

	// The roles for which this category is available, if category is set to restricted. Not populated for lists of categories, only for individual category
	InstitutionRoleIds []string `json:"institutionRoleIds"`

	// The ID of this category's parent category
	ParentID string `json:"parentId,omitempty"`

	// Boolean indicating whether or not category is available to all roles, or restricted to a specific set of roles.
	Restricted bool `json:"restricted,omitempty"`

	// The title of category
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this post learn API public v1 catalog categories category type body
func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCategoryID(formats); err != nil {
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

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody) validateAvailable(formats strfmt.Registry) error {

	if err := validate.Required("category"+"."+"available", "body", o.Available); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody) validateCategoryID(formats strfmt.Registry) error {

	if err := validate.Required("category"+"."+"categoryId", "body", o.CategoryID); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("category"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CatalogCategoriesCategoryTypeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
