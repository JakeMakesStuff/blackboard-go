// Code generated by go-swagger; DO NOT EDIT.

package deprecated_course_grades

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

// GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsers structure.
type GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK creates a GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK() *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}/users][%d] getLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdUsersOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden creates a GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden() *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden handles this case with default header values.

Access Denied
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}/users][%d] getLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound creates a GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound() *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound handles this case with default header values.

Grade detail values have not yet been populated
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/columns/{columnId}/users][%d] getLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody get learn API public v1 courses course ID gradebook columns column ID users o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.Grade `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID gradebook columns column ID users o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdUsersOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdUsersOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdGradebookColumnsColumnIdUsersOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsColumnIDUsersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
