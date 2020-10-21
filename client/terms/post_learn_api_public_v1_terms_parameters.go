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
)

// NewPostLearnAPIPublicV1TermsParams creates a new PostLearnAPIPublicV1TermsParams object
// with the default values initialized.
func NewPostLearnAPIPublicV1TermsParams() *PostLearnAPIPublicV1TermsParams {
	var ()
	return &PostLearnAPIPublicV1TermsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearnAPIPublicV1TermsParamsWithTimeout creates a new PostLearnAPIPublicV1TermsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearnAPIPublicV1TermsParamsWithTimeout(timeout time.Duration) *PostLearnAPIPublicV1TermsParams {
	var ()
	return &PostLearnAPIPublicV1TermsParams{

		timeout: timeout,
	}
}

// NewPostLearnAPIPublicV1TermsParamsWithContext creates a new PostLearnAPIPublicV1TermsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearnAPIPublicV1TermsParamsWithContext(ctx context.Context) *PostLearnAPIPublicV1TermsParams {
	var ()
	return &PostLearnAPIPublicV1TermsParams{

		Context: ctx,
	}
}

// NewPostLearnAPIPublicV1TermsParamsWithHTTPClient creates a new PostLearnAPIPublicV1TermsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearnAPIPublicV1TermsParamsWithHTTPClient(client *http.Client) *PostLearnAPIPublicV1TermsParams {
	var ()
	return &PostLearnAPIPublicV1TermsParams{
		HTTPClient: client,
	}
}

/*PostLearnAPIPublicV1TermsParams contains all the parameters to send to the API endpoint
for the post learn API public v1 terms operation typically these are written to a http.Request
*/
type PostLearnAPIPublicV1TermsParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input
	  JSON input object.

	*/
	Input PostLearnAPIPublicV1TermsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) WithTimeout(timeout time.Duration) *PostLearnAPIPublicV1TermsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) WithContext(ctx context.Context) *PostLearnAPIPublicV1TermsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) WithHTTPClient(client *http.Client) *PostLearnAPIPublicV1TermsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) WithFields(fields *string) *PostLearnAPIPublicV1TermsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) WithInput(input PostLearnAPIPublicV1TermsBody) *PostLearnAPIPublicV1TermsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the post learn API public v1 terms params
func (o *PostLearnAPIPublicV1TermsParams) SetInput(input PostLearnAPIPublicV1TermsBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearnAPIPublicV1TermsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.Input); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}