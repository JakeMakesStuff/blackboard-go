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

// GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDReader is a Reader for the GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserID structure.
type GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK creates a GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK() *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK {
	return &GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK struct {
	Payload *models.GroupMembership
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/groups/{groupId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdGroupsGroupIdUsersUserIdOK  %+v", 200, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK) GetPayload() *models.GroupMembership {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupMembership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest creates a GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest() *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest {
	return &GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest handles this case with default header values.

Bad Request
*/
type GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/groups/{groupId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdGroupsGroupIdUsersUserIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden creates a GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden() *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden {
	return &GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden handles this case with default header values.

Forbidden
*/
type GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/groups/{groupId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdGroupsGroupIdUsersUserIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound creates a GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound with default headers values
func NewGetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound() *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound {
	return &GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound{}
}

/*GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound handles this case with default header values.

Not Found
*/
type GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound struct {
	Payload *models.RestException
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/courses/{courseId}/groups/{groupId}/users/{userId}][%d] getLearnApiPublicV1CoursesCourseIdGroupsGroupIdUsersUserIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound) GetPayload() *models.RestException {
	return o.Payload
}

func (o *GetLearnAPIPublicV1CoursesCourseIDGroupsGroupIDUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
