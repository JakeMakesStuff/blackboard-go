// Code generated by go-swagger; DO NOT EDIT.

package course_grade_notations

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

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams creates a new PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams object
// with the default values initialized.
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams() *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	var ()
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParamsWithTimeout creates a new PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParamsWithTimeout(timeout time.Duration) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	var ()
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams{

		timeout: timeout,
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParamsWithContext creates a new PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParamsWithContext(ctx context.Context) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	var ()
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams{

		Context: ctx,
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParamsWithHTTPClient creates a new PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParamsWithHTTPClient(client *http.Client) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	var ()
	return &PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams{
		HTTPClient: client,
	}
}

/*PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams contains all the parameters to send to the API endpoint
for the post learn API public v1 courses course ID gradebook grade notations operation typically these are written to a http.Request
*/
type PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams struct {

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
	/*GradeNotationInput*/
	GradeNotationInput PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) WithTimeout(timeout time.Duration) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) WithContext(ctx context.Context) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) WithHTTPClient(client *http.Client) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) WithCourseID(courseID string) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) WithFields(fields *string) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGradeNotationInput adds the gradeNotationInput to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) WithGradeNotationInput(gradeNotationInput PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBody) *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams {
	o.SetGradeNotationInput(gradeNotationInput)
	return o
}

// SetGradeNotationInput adds the gradeNotationInput to the post learn API public v1 courses course ID gradebook grade notations params
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) SetGradeNotationInput(gradeNotationInput PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsBody) {
	o.GradeNotationInput = gradeNotationInput
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearnAPIPublicV1CoursesCourseIDGradebookGradeNotationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.GradeNotationInput); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}