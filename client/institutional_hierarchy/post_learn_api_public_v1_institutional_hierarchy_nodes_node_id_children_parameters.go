// Code generated by go-swagger; DO NOT EDIT.

package institutional_hierarchy

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

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams creates a new PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams object
// with the default values initialized.
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams() *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	var ()
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParamsWithTimeout creates a new PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParamsWithTimeout(timeout time.Duration) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	var ()
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams{

		timeout: timeout,
	}
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParamsWithContext creates a new PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParamsWithContext(ctx context.Context) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	var ()
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams{

		Context: ctx,
	}
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParamsWithHTTPClient creates a new PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParamsWithHTTPClient(client *http.Client) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	var ()
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams{
		HTTPClient: client,
	}
}

/*PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams contains all the parameters to send to the API endpoint
for the post learn API public v1 institutional hierarchy nodes node ID children operation typically these are written to a http.Request
*/
type PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input*/
	Input PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenBody
	/*NodeID
	 The node ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:breakfastClub              |



	*/
	NodeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) WithTimeout(timeout time.Duration) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) WithContext(ctx context.Context) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) WithHTTPClient(client *http.Client) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) WithFields(fields *string) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) WithInput(input PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenBody) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) SetInput(input PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenBody) {
	o.Input = input
}

// WithNodeID adds the nodeID to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) WithNodeID(nodeID string) *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the post learn API public v1 institutional hierarchy nodes node ID children params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesNodeIDChildrenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nodeId
	if err := r.SetPathParam("nodeId", o.NodeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
