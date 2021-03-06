// Code generated by go-swagger; DO NOT EDIT.

package course_group_users

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
	"github.com/go-openapi/swag"
)

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams creates a new GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams object
// with the default values initialized.
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams() *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParamsWithTimeout creates a new GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParamsWithContext creates a new GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParamsWithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParamsWithHTTPClient creates a new GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams contains all the parameters to send to the API endpoint
for the get learn API public v2 courses course ID groups group ID users operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams struct {

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
	/*Limit
	  The maximum number of results to be returned. There may be less if the query returned less than the maximum.

	*/
	Limit *int32
	/*Offset
	  The number of rows to skip before beginning to return rows. An offset of 0 is the same as omitting the offset parameter.

	*/
	Offset *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WithCourseID(courseID string) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WithFields(fields *string) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGroupID adds the groupID to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WithGroupID(groupID string) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithLimit adds the limit to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WithLimit(limit *int32) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WithOffset(offset *int32) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v2 courses course ID groups group ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
