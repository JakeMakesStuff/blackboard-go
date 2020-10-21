// Code generated by go-swagger; DO NOT EDIT.

package course_groups

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

// GetLearnAPIPublicV2CoursesCourseIDGroupsSetsReader is a Reader for the GetLearnAPIPublicV2CoursesCourseIDGroupsSets structure.
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK creates a GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK() *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK {
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK struct {
	Payload *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/groups/sets][%d] getLearnApiPublicV2CoursesCourseIdGroupsSetsOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK) GetPayload() *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest creates a GetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest() *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest {
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/groups/sets][%d] getLearnApiPublicV2CoursesCourseIdGroupsSetsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden creates a GetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden with default headers values
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden() *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden {
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden{}
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v2/courses/{courseId}/groups/sets][%d] getLearnApiPublicV2CoursesCourseIdGroupsSetsForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody get learn API public v2 courses course ID groups sets o k body
swagger:model GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.GroupV2 `json:"results"`
}

// Validate validates this get learn API public v2 courses course ID groups sets o k body
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV2CoursesCourseIdGroupsSetsOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV2CoursesCourseIdGroupsSetsOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV2CoursesCourseIdGroupsSetsOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV2CoursesCourseIDGroupsSetsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}