// Code generated by go-swagger; DO NOT EDIT.

package attendance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams creates a new DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams object
// with the default values initialized.
func NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams() *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	var ()
	return &DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParamsWithTimeout creates a new DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParamsWithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	var ()
	return &DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams{

		timeout: timeout,
	}
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParamsWithContext creates a new DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParamsWithContext(ctx context.Context) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	var ()
	return &DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams{

		Context: ctx,
	}
}

// NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParamsWithHTTPClient creates a new DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParamsWithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	var ()
	return &DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams{
		HTTPClient: client,
	}
}

/*DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams contains all the parameters to send to the API endpoint
for the delete learn API public v1 courses course ID meetings meeting ID users user ID operation typically these are written to a http.Request
*/
type DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams struct {

	/*CourseID
	 The course or organization ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:math101                    |
	| courseId   | courseId:math101                      |
	| uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	*/
	CourseID string
	/*MeetingID*/
	MeetingID string
	/*UserID
	 The user ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:jsmith                     |
	| userName   | userName:jsmith                       |
	| uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) WithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) WithContext(ctx context.Context) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) WithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) WithCourseID(courseID string) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithMeetingID adds the meetingID to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) WithMeetingID(meetingID string) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	o.SetMeetingID(meetingID)
	return o
}

// SetMeetingID adds the meetingId to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) SetMeetingID(meetingID string) {
	o.MeetingID = meetingID
}

// WithUserID adds the userID to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) WithUserID(userID string) *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete learn API public v1 courses course ID meetings meeting ID users user ID params
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
		return err
	}

	// path param meetingId
	if err := r.SetPathParam("meetingId", o.MeetingID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
