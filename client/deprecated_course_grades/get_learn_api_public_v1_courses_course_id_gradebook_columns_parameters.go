// Code generated by go-swagger; DO NOT EDIT.

package deprecated_course_grades

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

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams creates a new GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams() *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID gradebook columns operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams struct {

	/*ContentID
	  Search for grade columns associated with this content item.

	**Since**: 3000.11.0

	*/
	ContentID *string
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

// WithTimeout adds the timeout to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentID adds the contentID to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WithContentID(contentID *string) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) SetContentID(contentID *string) {
	o.ContentID = contentID
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WithFields(fields *string) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WithLimit(limit *int32) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WithOffset(offset *int32) *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v1 courses course ID gradebook columns params
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDGradebookColumnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentID != nil {

		// query param contentId
		var qrContentID string
		if o.ContentID != nil {
			qrContentID = *o.ContentID
		}
		qContentID := qrContentID
		if qContentID != "" {
			if err := r.SetQueryParam("contentId", qContentID); err != nil {
				return err
			}
		}

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
