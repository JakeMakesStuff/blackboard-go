// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PostLearnAPIPublicV1UploadsReader is a Reader for the PostLearnAPIPublicV1Uploads structure.
type PostLearnAPIPublicV1UploadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1UploadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1UploadsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewPostLearnAPIPublicV1UploadsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1UploadsCreated creates a PostLearnAPIPublicV1UploadsCreated with default headers values
func NewPostLearnAPIPublicV1UploadsCreated() *PostLearnAPIPublicV1UploadsCreated {
	return &PostLearnAPIPublicV1UploadsCreated{}
}

/*PostLearnAPIPublicV1UploadsCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1UploadsCreated struct {
	Payload *models.UploadedFileInfo
}

func (o *PostLearnAPIPublicV1UploadsCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/uploads][%d] postLearnApiPublicV1UploadsCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1UploadsCreated) GetPayload() *models.UploadedFileInfo {
	return o.Payload
}

func (o *PostLearnAPIPublicV1UploadsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadedFileInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1UploadsUnprocessableEntity creates a PostLearnAPIPublicV1UploadsUnprocessableEntity with default headers values
func NewPostLearnAPIPublicV1UploadsUnprocessableEntity() *PostLearnAPIPublicV1UploadsUnprocessableEntity {
	return &PostLearnAPIPublicV1UploadsUnprocessableEntity{}
}

/*PostLearnAPIPublicV1UploadsUnprocessableEntity handles this case with default header values.

File is potentially unsafe as determined from an XSS scan
*/
type PostLearnAPIPublicV1UploadsUnprocessableEntity struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1UploadsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/uploads][%d] postLearnApiPublicV1UploadsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostLearnAPIPublicV1UploadsUnprocessableEntity) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1UploadsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}