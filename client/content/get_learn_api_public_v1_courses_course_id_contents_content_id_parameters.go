// Code generated by go-swagger; DO NOT EDIT.

package content

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

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams() *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDContentsContentIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID contents content ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams struct {

	/*ContentID
	 The Content ID.  This may be the primary ID, or any of the following keywords: interactive, indirect, root.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| keyword    | root                                  |


	*/
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentID adds the contentID to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) WithContentID(contentID string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) WithFields(fields *string) *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 courses course ID contents content ID params
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDContentsContentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}