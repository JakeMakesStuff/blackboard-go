// Code generated by go-swagger; DO NOT EDIT.

package course_grading_periods

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

// GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDGradebookPeriods structure.
type GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK creates a GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK() *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/periods][%d] getLearnApiPublicV1CoursesCourseIdGradebookPeriodsOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest() *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/periods][%d] getLearnApiPublicV1CoursesCourseIdGradebookPeriodsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden creates a GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden() *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/periods][%d] getLearnApiPublicV1CoursesCourseIdGradebookPeriodsForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound creates a GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound() *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/periods][%d] getLearnApiPublicV1CoursesCourseIdGradebookPeriodsNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody get learn API public v1 courses course ID gradebook periods o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.GradingPeriod `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID gradebook periods o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdGradebookPeriodsOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdGradebookPeriodsOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdGradebookPeriodsOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDGradebookPeriodsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}