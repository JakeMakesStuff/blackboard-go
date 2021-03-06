// Code generated by go-swagger; DO NOT EDIT.

package terms

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

// NewGetLearnAPIPublicV1TermsParams creates a new GetLearnAPIPublicV1TermsParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1TermsParams() *GetLearnAPIPublicV1TermsParams {
	var ()
	return &GetLearnAPIPublicV1TermsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1TermsParamsWithTimeout creates a new GetLearnAPIPublicV1TermsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1TermsParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1TermsParams {
	var ()
	return &GetLearnAPIPublicV1TermsParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1TermsParamsWithContext creates a new GetLearnAPIPublicV1TermsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1TermsParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1TermsParams {
	var ()
	return &GetLearnAPIPublicV1TermsParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1TermsParamsWithHTTPClient creates a new GetLearnAPIPublicV1TermsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1TermsParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1TermsParams {
	var ()
	return &GetLearnAPIPublicV1TermsParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1TermsParams contains all the parameters to send to the API endpoint
for the get learn API public v1 terms operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1TermsParams struct {

	/*AvailabilityAvailable
	  Search for users with availability.available properties that contain this value.

	**Since**: 3100.6.0


	| Type      | Description
	 | --------- | --------- |
	| Yes | Students may access the term and the courses it contains. |
	| No | Students may not access the term or the courses it contains. |


	*/
	AvailabilityAvailable *string
	/*DataSourceID
	  Search for term with this dataSourceId.

	**Since**: 3100.6.0

	*/
	DataSourceID *string
	/*ExternalID
	  Search for term with externalId properties that contain this value.

	**Since**: 3100.6.0

	*/
	ExternalID *string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1TermsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1TermsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1TermsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityAvailable adds the availabilityAvailable to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithAvailabilityAvailable(availabilityAvailable *string) *GetLearnAPIPublicV1TermsParams {
	o.SetAvailabilityAvailable(availabilityAvailable)
	return o
}

// SetAvailabilityAvailable adds the availabilityAvailable to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetAvailabilityAvailable(availabilityAvailable *string) {
	o.AvailabilityAvailable = availabilityAvailable
}

// WithDataSourceID adds the dataSourceID to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithDataSourceID(dataSourceID *string) *GetLearnAPIPublicV1TermsParams {
	o.SetDataSourceID(dataSourceID)
	return o
}

// SetDataSourceID adds the dataSourceId to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetDataSourceID(dataSourceID *string) {
	o.DataSourceID = dataSourceID
}

// WithExternalID adds the externalID to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithExternalID(externalID *string) *GetLearnAPIPublicV1TermsParams {
	o.SetExternalID(externalID)
	return o
}

// SetExternalID adds the externalId to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetExternalID(externalID *string) {
	o.ExternalID = externalID
}

// WithFields adds the fields to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithFields(fields *string) *GetLearnAPIPublicV1TermsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithLimit(limit *int32) *GetLearnAPIPublicV1TermsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) WithOffset(offset *int32) *GetLearnAPIPublicV1TermsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v1 terms params
func (o *GetLearnAPIPublicV1TermsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1TermsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvailabilityAvailable != nil {

		// query param availability.available
		var qrAvailabilityAvailable string
		if o.AvailabilityAvailable != nil {
			qrAvailabilityAvailable = *o.AvailabilityAvailable
		}
		qAvailabilityAvailable := qrAvailabilityAvailable
		if qAvailabilityAvailable != "" {
			if err := r.SetQueryParam("availability.available", qAvailabilityAvailable); err != nil {
				return err
			}
		}

	}

	if o.DataSourceID != nil {

		// query param dataSourceId
		var qrDataSourceID string
		if o.DataSourceID != nil {
			qrDataSourceID = *o.DataSourceID
		}
		qDataSourceID := qrDataSourceID
		if qDataSourceID != "" {
			if err := r.SetQueryParam("dataSourceId", qDataSourceID); err != nil {
				return err
			}
		}

	}

	if o.ExternalID != nil {

		// query param externalId
		var qrExternalID string
		if o.ExternalID != nil {
			qrExternalID = *o.ExternalID
		}
		qExternalID := qrExternalID
		if qExternalID != "" {
			if err := r.SetQueryParam("externalId", qExternalID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
