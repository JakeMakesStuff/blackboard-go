// Code generated by go-swagger; DO NOT EDIT.

package content_resources

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

// GetLearnAPIPublicV1CoursesCourseIDResourcesReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDResources structure.
type GetLearnAPIPublicV1CoursesCourseIDResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesOK creates a GetLearnAPIPublicV1CoursesCourseIDResourcesOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesOK() *GetLearnAPIPublicV1CoursesCourseIDResourcesOK {
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/resources][%d] getLearnApiPublicV1CoursesCourseIdResourcesOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesNotFound creates a GetLearnAPIPublicV1CoursesCourseIDResourcesNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesNotFound() *GetLearnAPIPublicV1CoursesCourseIDResourcesNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesNotFound handles this case with default header values.

The specified course could not be found
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/resources][%d] getLearnApiPublicV1CoursesCourseIdResourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests creates a GetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests() *GetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests {
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests{}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests handles this case with default header values.

The endpoint is being overloaded with requests
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/resources][%d] getLearnApiPublicV1CoursesCourseIdResourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody get learn API public v1 courses course ID resources o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.Resource `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID resources o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdResourcesOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdResourcesOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdResourcesOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
