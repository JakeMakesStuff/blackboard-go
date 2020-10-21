// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewGetLearnAPIPublicV1CourseRolesRoleIDParams creates a new GetLearnAPIPublicV1CourseRolesRoleIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CourseRolesRoleIDParams() *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	var ()
	return &GetLearnAPIPublicV1CourseRolesRoleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CourseRolesRoleIDParamsWithTimeout creates a new GetLearnAPIPublicV1CourseRolesRoleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CourseRolesRoleIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	var ()
	return &GetLearnAPIPublicV1CourseRolesRoleIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CourseRolesRoleIDParamsWithContext creates a new GetLearnAPIPublicV1CourseRolesRoleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CourseRolesRoleIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	var ()
	return &GetLearnAPIPublicV1CourseRolesRoleIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CourseRolesRoleIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1CourseRolesRoleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CourseRolesRoleIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	var ()
	return &GetLearnAPIPublicV1CourseRolesRoleIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CourseRolesRoleIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 course roles role ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CourseRolesRoleIDParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*RoleID
	  The course role ID.  This may be the primary ID, or the roleId. The suffix ":custom" will be appended to the roleId of a custom course role if that roleId conflicts with the roleId of a system generated course role.  For example, if a custom role roleId is specified as "Student" then the roleId will actually be "Student:custom" since there is already a system generated course role with the roleId of "Student".

	 | ID type    | Examples                                                   |
	 |------------|------------------------------------------------------------|
	 | primary    | _123_1                                                     |
	 | roleId     | roleId:Student, roleId:MyCustomRole, roleId:Student:custom |


	**Since**: 3300.5.0

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) WithFields(fields *string) *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithRoleID adds the roleID to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) WithRoleID(roleID string) *GetLearnAPIPublicV1CourseRolesRoleIDParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get learn API public v1 course roles role ID params
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CourseRolesRoleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}