// Code generated by go-swagger; DO NOT EDIT.

package deprecated_course_group_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDReader is a Reader for the PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserID structure.
type PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated creates a PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated with default headers values
func NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated() *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated {
	return &PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated{}
}

/*PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated handles this case with default header values.

Created
*/
type PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated struct {
	Payload *models.GroupMembership
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV1CoursesCourseIdGroupsGroupIdUsersUserIdCreated  %+v", 201, o.Payload)
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated) GetPayload() *models.GroupMembership {
	return o.Payload
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest creates a PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest with default headers values
func NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest() *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest {
	return &PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest{}
}

/*PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest handles this case with default header values.

Bad Request
*/
type PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV1CoursesCourseIdGroupsGroupIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden creates a PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden with default headers values
func NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden() *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden {
	return &PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden{}
}

/*PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden handles this case with default header values.

Forbidden
*/
type PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV1CoursesCourseIdGroupsGroupIdUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict creates a PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict with default headers values
func NewPutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict() *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict {
	return &PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict{}
}

/*PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict handles this case with default header values.

Conflict
*/
type PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict struct {
	Payload *models.RestException
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict) Error() string {
	return fmt.Sprintf("[PUT /learn/api/public/v1/courses/{courseId}/groups/{groupId}/users/{userId}][%d] putLearnApiPublicV1CoursesCourseIdGroupsGroupIdUsersUserIdConflict  %+v", 409, o.Payload)
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PutLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
