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

// NewPatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams creates a new PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams object
// with the default values initialized.
func NewPatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams() *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	var ()
	return &PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithTimeout creates a new PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithTimeout(timeout time.Duration) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	var ()
	return &PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams{

		timeout: timeout,
	}
}

// NewPatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithContext creates a new PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithContext(ctx context.Context) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	var ()
	return &PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams{

		Context: ctx,
	}
}

// NewPatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithHTTPClient creates a new PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParamsWithHTTPClient(client *http.Client) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	var ()
	return &PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams{
		HTTPClient: client,
	}
}

/*PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams contains all the parameters to send to the API endpoint
for the patch learn API public v2 courses course ID groups group ID operation typically these are written to a http.Request
*/
type PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams struct {

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
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*GroupID
	 The group ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:breakfastClub              |


	*/
	GroupID string
	/*Input
	  JSON input object.

	*/
	Input PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithTimeout(timeout time.Duration) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithContext(ctx context.Context) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithHTTPClient(client *http.Client) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithCourseID(courseID string) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithFields(fields *string) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGroupID adds the groupID to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithGroupID(groupID string) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithInput adds the input to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WithInput(input PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDBody) *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the patch learn API public v2 courses course ID groups group ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) SetInput(input PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLearnAPIPublicV2CoursesCourseIDGroupsGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Input); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
