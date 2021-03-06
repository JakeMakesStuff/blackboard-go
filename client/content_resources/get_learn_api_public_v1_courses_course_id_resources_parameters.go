// Code generated by go-swagger; DO NOT EDIT.

package content_resources

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

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesParams creates a new GetLearnAPIPublicV1CoursesCourseIDResourcesParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesParams() *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDResourcesParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDResourcesParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDResourcesParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDResourcesParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID resources operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDResourcesParams struct {

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
	/*Limit
	  The maximum number of results to be returned. There may be less if the query returned less than the maximum.

	*/
	Limit *int32
	/*Offset
	  The number of rows to skip before beginning to return rows. An offset of 0 is the same as omitting the offset parameter.

	*/
	Offset *int32
	/*Type
	  Search for Resources whose Type matches the specified value.  Valid values are 'File' and 'Folder'.


	| Type      | Description
	 | --------- | --------- |
	| File |  |
	| Folder |  |


	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WithFields(fields *string) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WithLimit(limit *int32) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WithOffset(offset *int32) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithType adds the typeVar to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WithType(typeVar *string) *GetLearnAPIPublicV1CoursesCourseIDResourcesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get learn API public v1 courses course ID resources params
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
