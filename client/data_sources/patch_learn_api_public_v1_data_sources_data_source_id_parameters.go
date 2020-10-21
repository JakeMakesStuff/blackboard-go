// Code generated by go-swagger; DO NOT EDIT.

package data_sources

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

// NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParams creates a new PatchLearnAPIPublicV1DataSourcesDataSourceIDParams object
// with the default values initialized.
func NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParams() *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	var ()
	return &PatchLearnAPIPublicV1DataSourcesDataSourceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParamsWithTimeout creates a new PatchLearnAPIPublicV1DataSourcesDataSourceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParamsWithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	var ()
	return &PatchLearnAPIPublicV1DataSourcesDataSourceIDParams{

		timeout: timeout,
	}
}

// NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParamsWithContext creates a new PatchLearnAPIPublicV1DataSourcesDataSourceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParamsWithContext(ctx context.Context) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	var ()
	return &PatchLearnAPIPublicV1DataSourcesDataSourceIDParams{

		Context: ctx,
	}
}

// NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParamsWithHTTPClient creates a new PatchLearnAPIPublicV1DataSourcesDataSourceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLearnAPIPublicV1DataSourcesDataSourceIDParamsWithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	var ()
	return &PatchLearnAPIPublicV1DataSourcesDataSourceIDParams{
		HTTPClient: client,
	}
}

/*PatchLearnAPIPublicV1DataSourcesDataSourceIDParams contains all the parameters to send to the API endpoint
for the patch learn API public v1 data sources data source ID operation typically these are written to a http.Request
*/
type PatchLearnAPIPublicV1DataSourcesDataSourceIDParams struct {

	/*DataSourceID
	 The data source ID.  This may be the primary ID, or the secondary ID prefixed with the ID type.

	| ID type    | Example            |
	|------------|--------------------|
	| primary    | _123_1             |
	| externalId | externalId:math101 |


	*/
	DataSourceID string
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input
	  JSON input object.

	*/
	Input PatchLearnAPIPublicV1DataSourcesDataSourceIDBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) WithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) WithContext(ctx context.Context) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) WithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataSourceID adds the dataSourceID to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) WithDataSourceID(dataSourceID string) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	o.SetDataSourceID(dataSourceID)
	return o
}

// SetDataSourceID adds the dataSourceId to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) SetDataSourceID(dataSourceID string) {
	o.DataSourceID = dataSourceID
}

// WithFields adds the fields to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) WithFields(fields *string) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) WithInput(input PatchLearnAPIPublicV1DataSourcesDataSourceIDBody) *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the patch learn API public v1 data sources data source ID params
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) SetInput(input PatchLearnAPIPublicV1DataSourcesDataSourceIDBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLearnAPIPublicV1DataSourcesDataSourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dataSourceId
	if err := r.SetPathParam("dataSourceId", o.DataSourceID); err != nil {
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

	if err := r.SetBodyParam(o.Input); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}