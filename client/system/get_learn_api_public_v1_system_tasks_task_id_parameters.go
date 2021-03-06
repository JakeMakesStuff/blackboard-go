// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetLearnAPIPublicV1SystemTasksTaskIDParams creates a new GetLearnAPIPublicV1SystemTasksTaskIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1SystemTasksTaskIDParams() *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	var ()
	return &GetLearnAPIPublicV1SystemTasksTaskIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1SystemTasksTaskIDParamsWithTimeout creates a new GetLearnAPIPublicV1SystemTasksTaskIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1SystemTasksTaskIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	var ()
	return &GetLearnAPIPublicV1SystemTasksTaskIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1SystemTasksTaskIDParamsWithContext creates a new GetLearnAPIPublicV1SystemTasksTaskIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1SystemTasksTaskIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	var ()
	return &GetLearnAPIPublicV1SystemTasksTaskIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1SystemTasksTaskIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1SystemTasksTaskIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1SystemTasksTaskIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	var ()
	return &GetLearnAPIPublicV1SystemTasksTaskIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1SystemTasksTaskIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 system tasks task ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1SystemTasksTaskIDParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*TaskID*/
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) WithFields(fields *string) *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithTaskID adds the taskID to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) WithTaskID(taskID string) *GetLearnAPIPublicV1SystemTasksTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get learn API public v1 system tasks task ID params
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1SystemTasksTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
