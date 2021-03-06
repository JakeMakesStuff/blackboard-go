// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV1UsersUserIDReader is a Reader for the DeleteLearnAPIPublicV1UsersUserID structure.
type DeleteLearnAPIPublicV1UsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1UsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearnAPIPublicV1UsersUserIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearnAPIPublicV1UsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearnAPIPublicV1UsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1UsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1UsersUserIDNoContent creates a DeleteLearnAPIPublicV1UsersUserIDNoContent with default headers values
func NewDeleteLearnAPIPublicV1UsersUserIDNoContent() *DeleteLearnAPIPublicV1UsersUserIDNoContent {
	return &DeleteLearnAPIPublicV1UsersUserIDNoContent{}
}

/*DeleteLearnAPIPublicV1UsersUserIDNoContent handles this case with default header values.

No Content
*/
type DeleteLearnAPIPublicV1UsersUserIDNoContent struct {
}

func (o *DeleteLearnAPIPublicV1UsersUserIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/users/{userId}][%d] deleteLearnApiPublicV1UsersUserIdNoContent ", 204)
}

func (o *DeleteLearnAPIPublicV1UsersUserIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1UsersUserIDBadRequest creates a DeleteLearnAPIPublicV1UsersUserIDBadRequest with default headers values
func NewDeleteLearnAPIPublicV1UsersUserIDBadRequest() *DeleteLearnAPIPublicV1UsersUserIDBadRequest {
	return &DeleteLearnAPIPublicV1UsersUserIDBadRequest{}
}

/*DeleteLearnAPIPublicV1UsersUserIDBadRequest handles this case with default header values.

Invalid userid provided
*/
type DeleteLearnAPIPublicV1UsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1UsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/users/{userId}][%d] deleteLearnApiPublicV1UsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearnAPIPublicV1UsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1UsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1UsersUserIDForbidden creates a DeleteLearnAPIPublicV1UsersUserIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1UsersUserIDForbidden() *DeleteLearnAPIPublicV1UsersUserIDForbidden {
	return &DeleteLearnAPIPublicV1UsersUserIDForbidden{}
}

/*DeleteLearnAPIPublicV1UsersUserIDForbidden handles this case with default header values.

The user is not authorized to delete the specified User object
*/
type DeleteLearnAPIPublicV1UsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1UsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/users/{userId}][%d] deleteLearnApiPublicV1UsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1UsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1UsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1UsersUserIDNotFound creates a DeleteLearnAPIPublicV1UsersUserIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1UsersUserIDNotFound() *DeleteLearnAPIPublicV1UsersUserIDNotFound {
	return &DeleteLearnAPIPublicV1UsersUserIDNotFound{}
}

/*DeleteLearnAPIPublicV1UsersUserIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV1UsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1UsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/users/{userId}][%d] deleteLearnApiPublicV1UsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1UsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1UsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
