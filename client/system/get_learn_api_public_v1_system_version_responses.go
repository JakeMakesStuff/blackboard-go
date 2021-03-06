// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1SystemVersionReader is a Reader for the GetLearnAPIPublicV1SystemVersion structure.
type GetLearnAPIPublicV1SystemVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1SystemVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1SystemVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1SystemVersionOK creates a GetLearnAPIPublicV1SystemVersionOK with default headers values
func NewGetLearnAPIPublicV1SystemVersionOK() *GetLearnAPIPublicV1SystemVersionOK {
	return &GetLearnAPIPublicV1SystemVersionOK{}
}

/*GetLearnAPIPublicV1SystemVersionOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1SystemVersionOK struct {
	Payload *models.VersionInfo
}

func (o *GetLearnAPIPublicV1SystemVersionOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/system/version][%d] getLearnApiPublicV1SystemVersionOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1SystemVersionOK) GetPayload() *models.VersionInfo {
	return o.Payload
}

func (o *GetLearnAPIPublicV1SystemVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
