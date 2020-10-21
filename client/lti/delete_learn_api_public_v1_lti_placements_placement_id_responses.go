// Code generated by go-swagger; DO NOT EDIT.

package lti

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// DeleteLearnAPIPublicV1LtiPlacementsPlacementIDReader is a Reader for the DeleteLearnAPIPublicV1LtiPlacementsPlacementID structure.
type DeleteLearnAPIPublicV1LtiPlacementsPlacementIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent creates a DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent with default headers values
func NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent() *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent {
	return &DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent{}
}

/*DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent handles this case with default header values.

No Content
*/
type DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent struct {
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/lti/placements/{placementId}][%d] deleteLearnApiPublicV1LtiPlacementsPlacementIdNoContent ", 204)
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest creates a DeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest with default headers values
func NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest() *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest {
	return &DeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest{}
}

/*DeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/lti/placements/{placementId}][%d] deleteLearnApiPublicV1LtiPlacementsPlacementIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden creates a DeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden with default headers values
func NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden() *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden {
	return &DeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden{}
}

/*DeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/lti/placements/{placementId}][%d] deleteLearnApiPublicV1LtiPlacementsPlacementIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound creates a DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound with default headers values
func NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound() *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound {
	return &DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound{}
}

/*DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound handles this case with default header values.

Not Found
*/
type DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound struct {
	Payload *models.RestException
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /learn/api/public/v1/lti/placements/{placementId}][%d] deleteLearnApiPublicV1LtiPlacementsPlacementIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}