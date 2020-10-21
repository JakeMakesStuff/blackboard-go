// Code generated by go-swagger; DO NOT EDIT.

package calendar

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

// GetLearnAPIPublicV1CalendarsReader is a Reader for the GetLearnAPIPublicV1Calendars structure.
type GetLearnAPIPublicV1CalendarsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CalendarsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CalendarsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CalendarsOK creates a GetLearnAPIPublicV1CalendarsOK with default headers values
func NewGetLearnAPIPublicV1CalendarsOK() *GetLearnAPIPublicV1CalendarsOK {
	return &GetLearnAPIPublicV1CalendarsOK{}
}

/*GetLearnAPIPublicV1CalendarsOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CalendarsOK struct {
	Payload *GetLearnAPIPublicV1CalendarsOKBody
}

func (o *GetLearnAPIPublicV1CalendarsOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/calendars][%d] getLearnApiPublicV1CalendarsOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CalendarsOK) GetPayload() *GetLearnAPIPublicV1CalendarsOKBody {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CalendarsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLearnAPIPublicV1CalendarsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetLearnAPIPublicV1CalendarsOKBody get learn API public v1 calendars o k body
swagger:model GetLearnAPIPublicV1CalendarsOKBody
*/
type GetLearnAPIPublicV1CalendarsOKBody struct {

	// paging
	Paging *models.PagingInfo `json:"paging,omitempty"`

	// results
	// Required: true
	Results []*models.Calendar `json:"results"`
}

// Validate validates this get learn API public v1 calendars o k body
func (o *GetLearnAPIPublicV1CalendarsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetLearnAPIPublicV1CalendarsOKBody) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(o.Paging) { // not required
		return nil
	}

	if o.Paging != nil {
		if err := o.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getLearnApiPublicV1CalendarsOK" + "." + "paging")
			}
			return err
		}
	}

	return nil
}

func (o *GetLearnAPIPublicV1CalendarsOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("getLearnApiPublicV1CalendarsOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getLearnApiPublicV1CalendarsOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CalendarsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLearnAPIPublicV1CalendarsOKBody) UnmarshalBinary(b []byte) error {
	var res GetLearnAPIPublicV1CalendarsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}