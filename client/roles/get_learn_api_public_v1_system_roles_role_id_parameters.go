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

// NewGetLearnAPIPublicV1SystemRolesRoleIDParams creates a new GetLearnAPIPublicV1SystemRolesRoleIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1SystemRolesRoleIDParams() *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	var ()
	return &GetLearnAPIPublicV1SystemRolesRoleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1SystemRolesRoleIDParamsWithTimeout creates a new GetLearnAPIPublicV1SystemRolesRoleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1SystemRolesRoleIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	var ()
	return &GetLearnAPIPublicV1SystemRolesRoleIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1SystemRolesRoleIDParamsWithContext creates a new GetLearnAPIPublicV1SystemRolesRoleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1SystemRolesRoleIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	var ()
	return &GetLearnAPIPublicV1SystemRolesRoleIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1SystemRolesRoleIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1SystemRolesRoleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1SystemRolesRoleIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	var ()
	return &GetLearnAPIPublicV1SystemRolesRoleIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1SystemRolesRoleIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 system roles role ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1SystemRolesRoleIDParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*RoleID
	  The System Role ID.  This may be the primary ID, or the secondary roleId type. The suffix ":custom" will be appended to the roleId of a custom system role if that roleId conflicts with the roleId of a system generated role.  For example, if a custom role roleId is specified as "Guest" then the roleId will actually be "Guest:custom" since there is already a system generated role with the roleId of "Guest".

	 | ID type    | Example                               |
	 |------------|---------------------------------------|
	 | primary    | _123_1                                |
	 | roleId     | roleId:column1                        |


	**Since**: 3300.5.0

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) WithFields(fields *string) *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithRoleID adds the roleID to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) WithRoleID(roleID string) *GetLearnAPIPublicV1SystemRolesRoleIDParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get learn API public v1 system roles role ID params
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1SystemRolesRoleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
