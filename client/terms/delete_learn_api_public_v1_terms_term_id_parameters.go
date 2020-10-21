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

// NewDeleteLearnAPIPublicV1TermsTermIDParams creates a new DeleteLearnAPIPublicV1TermsTermIDParams object
// with the default values initialized.
func NewDeleteLearnAPIPublicV1TermsTermIDParams() *DeleteLearnAPIPublicV1TermsTermIDParams {
	var ()
	return &DeleteLearnAPIPublicV1TermsTermIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearnAPIPublicV1TermsTermIDParamsWithTimeout creates a new DeleteLearnAPIPublicV1TermsTermIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearnAPIPublicV1TermsTermIDParamsWithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1TermsTermIDParams {
	var ()
	return &DeleteLearnAPIPublicV1TermsTermIDParams{

		timeout: timeout,
	}
}

// NewDeleteLearnAPIPublicV1TermsTermIDParamsWithContext creates a new DeleteLearnAPIPublicV1TermsTermIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearnAPIPublicV1TermsTermIDParamsWithContext(ctx context.Context) *DeleteLearnAPIPublicV1TermsTermIDParams {
	var ()
	return &DeleteLearnAPIPublicV1TermsTermIDParams{

		Context: ctx,
	}
}

// NewDeleteLearnAPIPublicV1TermsTermIDParamsWithHTTPClient creates a new DeleteLearnAPIPublicV1TermsTermIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearnAPIPublicV1TermsTermIDParamsWithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1TermsTermIDParams {
	var ()
	return &DeleteLearnAPIPublicV1TermsTermIDParams{
		HTTPClient: client,
	}
}

/*DeleteLearnAPIPublicV1TermsTermIDParams contains all the parameters to send to the API endpoint
for the delete learn API public v1 terms term ID operation typically these are written to a http.Request
*/
type DeleteLearnAPIPublicV1TermsTermIDParams struct {

	/*TermID
	 The term ID.  This may be the primary ID, or the secondary ID prefixed with the ID type.

	| ID type    | Example                |
	|------------|------------------------|
	| primary    | _123_1                 |
	| externalId | externalId:spring.2016 |


	*/
	TermID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete learn API public v1 terms term ID params
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) WithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1TermsTermIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learn API public v1 terms term ID params
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learn API public v1 terms term ID params
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) WithContext(ctx context.Context) *DeleteLearnAPIPublicV1TermsTermIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learn API public v1 terms term ID params
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learn API public v1 terms term ID params
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) WithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1TermsTermIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learn API public v1 terms term ID params
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTermID adds the termID to the delete learn API public v1 terms term ID params
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) WithTermID(termID string) *DeleteLearnAPIPublicV1TermsTermIDParams {
	o.SetTermID(termID)
	return o
}

// SetTermID adds the termId to the delete learn API public v1 terms term ID params
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) SetTermID(termID string) {
	o.TermID = termID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearnAPIPublicV1TermsTermIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param termId
	if err := r.SetPathParam("termId", o.TermID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}