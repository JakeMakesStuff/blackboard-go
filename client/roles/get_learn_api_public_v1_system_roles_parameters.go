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
	"github.com/go-openapi/swag"
)

// NewGetLearnAPIPublicV1SystemRolesParams creates a new GetLearnAPIPublicV1SystemRolesParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1SystemRolesParams() *GetLearnAPIPublicV1SystemRolesParams {
	var ()
	return &GetLearnAPIPublicV1SystemRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1SystemRolesParamsWithTimeout creates a new GetLearnAPIPublicV1SystemRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1SystemRolesParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1SystemRolesParams {
	var ()
	return &GetLearnAPIPublicV1SystemRolesParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1SystemRolesParamsWithContext creates a new GetLearnAPIPublicV1SystemRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1SystemRolesParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1SystemRolesParams {
	var ()
	return &GetLearnAPIPublicV1SystemRolesParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1SystemRolesParamsWithHTTPClient creates a new GetLearnAPIPublicV1SystemRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1SystemRolesParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1SystemRolesParams {
	var ()
	return &GetLearnAPIPublicV1SystemRolesParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1SystemRolesParams contains all the parameters to send to the API endpoint
for the get learn API public v1 system roles operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1SystemRolesParams struct {

	/*Custom
	  Search for system roles by custom flag.  A value of 'true' indicates that search results should be limited to only custom roles.  A value of 'false' indicates results should be limited to built-in roles.  Not setting this field indicates that all system roles should be returned.

	**Since**: 3300.5.0

	*/
	Custom *bool
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Limit
	  The maximum number of results to be returned. There may be less if the query returned less than the maximum.

	*/
	Limit *int32
	/*Offset
	  The number of rows to skip before beginning to return rows. An offset of 0 is the same as omitting the offset parameter.

	*/
	Offset *int32
	/*RoleID
	  Search for system roles with roleId

	**Since**: 3300.5.0

	*/
	RoleID *string
	/*Sort
	  Properties to sort by. This is a comma-delimited list of JSON properties, with an optional "(desc)" suffix to request a descending sort for that property. e.g. "roleId(desc)"

	Supported fields are:

	<ul - roleId
	- custom

	**Since**: 3300.5.0

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustom adds the custom to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithCustom(custom *bool) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetCustom(custom)
	return o
}

// SetCustom adds the custom to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetCustom(custom *bool) {
	o.Custom = custom
}

// WithFields adds the fields to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithFields(fields *string) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithLimit(limit *int32) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithOffset(offset *int32) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithRoleID adds the roleID to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithRoleID(roleID *string) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithSort adds the sort to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) WithSort(sort *string) *GetLearnAPIPublicV1SystemRolesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get learn API public v1 system roles params
func (o *GetLearnAPIPublicV1SystemRolesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1SystemRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Custom != nil {

		// query param custom
		var qrCustom bool
		if o.Custom != nil {
			qrCustom = *o.Custom
		}
		qCustom := swag.FormatBool(qrCustom)
		if qCustom != "" {
			if err := r.SetQueryParam("custom", qCustom); err != nil {
				return err
			}
		}

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

	if o.RoleID != nil {

		// query param roleId
		var qrRoleID string
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {
			if err := r.SetQueryParam("roleId", qRoleID); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}