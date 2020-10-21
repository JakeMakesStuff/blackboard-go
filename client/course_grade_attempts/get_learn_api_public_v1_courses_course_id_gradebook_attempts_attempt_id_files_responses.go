// Code generated by go-swagger; DO NOT EDIT.

package course_grade_attempts

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

// GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFiles structure.
type GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK creates a GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK() *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK struct {
	Payload *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/attempts/{attemptId}/files][%d] getLearnApiPublicV1CoursesCourseIdGradebookAttemptsAttemptIdFilesOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK) GetPayload() *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest() *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/attempts/{attemptId}/files][%d] getLearnApiPublicV1CoursesCourseIdGradebookAttemptsAttemptIdFilesBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden creates a GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden() *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/attempts/{attemptId}/files][%d] getLearnApiPublicV1CoursesCourseIdGradebookAttemptsAttemptIdFilesForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound creates a GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound() *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/gradebook/attempts/{attemptId}/files][%d] getLearnApiPublicV1CoursesCourseIdGradebookAttemptsAttemptIdFilesNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody get learn API public v1 courses course ID gradebook attempts attempt ID files o k body
swagger:model GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.AttemptFile `json:"results"`
}

// Validate validates this get learn API public v1 courses course ID gradebook attempts attempt ID files o k body
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdGradebookAttemptsAttemptIdFilesOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CoursesCourseIdGradebookAttemptsAttemptIdFilesOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CoursesCourseIdGradebookAttemptsAttemptIdFilesOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CoursesCourseIDGradebookAttemptsAttemptIDFilesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}