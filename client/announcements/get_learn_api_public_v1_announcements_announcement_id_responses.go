// Code generated by go-swagger; DO NOT EDIT.

package announcements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// GetLearnAPIPublicV1AnnouncementsAnnouncementIDReader is a Reader for the GetLearnAPIPublicV1AnnouncementsAnnouncementID structure.
type GetLearnAPIPublicV1AnnouncementsAnnouncementIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDOK creates a GetLearnAPIPublicV1AnnouncementsAnnouncementIDOK with default headers values
func NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDOK() *GetLearnAPIPublicV1AnnouncementsAnnouncementIDOK {
	return &GetLearnAPIPublicV1AnnouncementsAnnouncementIDOK{}
}

/*GetLearnAPIPublicV1AnnouncementsAnnouncementIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1AnnouncementsAnnouncementIDOK struct {
	Payload *models.Announcement
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/announcements/{announcementId}][%d] getLearnApiPublicV1AnnouncementsAnnouncementIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDOK) GetPayload() *models.Announcement {
	return o.Payload
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Announcement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden creates a GetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden with default headers values
func NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden() *GetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden {
	return &GetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden{}
}

/*GetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden handles this case with default header values.

The currently authenticated user has insufficient privileges to update an announcement
*/
type GetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/announcements/{announcementId}][%d] getLearnApiPublicV1AnnouncementsAnnouncementIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound creates a GetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound with default headers values
func NewGetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound() *GetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound {
	return &GetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound{}
}

/*GetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound handles this case with default header values.

Announcement cannot be found
*/
type GetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/announcements/{announcementId}][%d] getLearnApiPublicV1AnnouncementsAnnouncementIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1AnnouncementsAnnouncementIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}