// Code generated by go-swagger; DO NOT EDIT.

package roles

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
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1CourseRolesReader is a Reader for the GetLearnAPIPublicV1CourseRoles structure.
type GetLearnAPIPublicV1CourseRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CourseRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CourseRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLearnAPIPublicV1CourseRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CourseRolesOK creates a GetLearnAPIPublicV1CourseRolesOK with default headers values
func NewGetLearnAPIPublicV1CourseRolesOK() *GetLearnAPIPublicV1CourseRolesOK {
	return &GetLearnAPIPublicV1CourseRolesOK{}
}

/*GetLearnAPIPublicV1CourseRolesOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CourseRolesOK struct {
	Payload *GetLearnAPIPublicV1CourseRolesOKBody
}

func (o *GetLearnAPIPublicV1CourseRolesOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courseRoles][%d] getLearnApiPublicV1CourseRolesOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CourseRolesOK) GetPayload() *GetLearnAPIPublicV1CourseRolesOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CourseRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CourseRolesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CourseRolesForbidden creates a GetLearnAPIPublicV1CourseRolesForbidden with default headers values
func NewGetLearnAPIPublicV1CourseRolesForbidden() *GetLearnAPIPublicV1CourseRolesForbidden {
	return &GetLearnAPIPublicV1CourseRolesForbidden{}
}

/*GetLearnAPIPublicV1CourseRolesForbidden handles this case with default header values.

Access Denied
*/
type GetLearnAPIPublicV1CourseRolesForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CourseRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courseRoles][%d] getLearnApiPublicV1CourseRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CourseRolesForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CourseRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CourseRolesOKBody get learn API public v1 course roles o k body
swagger:model GetLearnAPIPublicV1CourseRolesOKBody
*/
type GetLearnAPIPublicV1CourseRolesOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.CourseRole `json:"results"`
}

// Validate validates this get learn API public v1 course roles o k body
func (o *GetLearnAPIPublicV1CourseRolesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLearnAPIPublicV1CourseRolesOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CourseRolesOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CourseRolesOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CourseRolesOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CourseRolesOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CourseRolesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CourseRolesOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CourseRolesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}