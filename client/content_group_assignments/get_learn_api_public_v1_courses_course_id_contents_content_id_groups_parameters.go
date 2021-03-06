// Code generated by go-swagger; DO NOT EDIT.

package content_group_assignments

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

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams() *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID contents content ID groups operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams struct {

	/*ContentID*/
	ContentID string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentID adds the contentID to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WithContentID(contentID string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WithFields(fields *string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WithLimit(limit *int32) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WithOffset(offset *int32) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v1 courses course ID contents content ID groups params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contentId
	if err := r.SetPathParam("contentId", o.ContentID); err != nil {
		return err
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
