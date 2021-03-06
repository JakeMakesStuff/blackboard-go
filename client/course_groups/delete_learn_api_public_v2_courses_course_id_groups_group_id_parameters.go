// Code generated by go-swagger; DO NOT EDIT.

package course_groups

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

// NewDeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams creates a new DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams object
// with the default values initialized.
func NewDeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams() *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithTimeout creates a new DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams{

		timeout: timeout,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithContext creates a new DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithContext(ctx context.Context) *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams{

		Context: ctx,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithHTTPClient creates a new DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams{
		HTTPClient: client,
	}
}

/*DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams contains all the parameters to send to the API endpoint
for the delete learn API public v2 courses course ID groups group ID operation typically these are written to a http.Request
*/
type DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams struct {

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
	/*GroupID
	 The group ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:breakfastClub              |


	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithContext(ctx context.Context) *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithCourseID(courseID string) *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithGroupID adds the groupID to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithGroupID(groupID string) *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the delete learn API public v2 courses course ID groups group ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
		return err
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
