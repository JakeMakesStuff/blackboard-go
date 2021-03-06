// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PostLearnAPIPublicV1CoursesCourseIDContentsReader is a Reader for the PostLearnAPIPublicV1CoursesCourseIDContents structure.
type PostLearnAPIPublicV1CoursesCourseIDContentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1CoursesCourseIDContentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1CoursesCourseIDContentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV1CoursesCourseIDContentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 507:
		result := NewPostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDContentsCreated creates a PostLearnAPIPublicV1CoursesCourseIDContentsCreated with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDContentsCreated() *PostLearnAPIPublicV1CoursesCourseIDContentsCreated {
	return &PostLearnAPIPublicV1CoursesCourseIDContentsCreated{}
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsCreated struct {
	Payload *models.Content
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/contents][%d] postLearnApiPublicV1CoursesCourseIdContentsCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreated) GetPayload() *models.Content {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Content)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDContentsBadRequest creates a PostLearnAPIPublicV1CoursesCourseIDContentsBadRequest with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDContentsBadRequest() *PostLearnAPIPublicV1CoursesCourseIDContentsBadRequest {
	return &PostLearnAPIPublicV1CoursesCourseIDContentsBadRequest{}
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsBadRequest handles this case with default header values.

Bad Request
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/contents][%d] postLearnApiPublicV1CoursesCourseIdContentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDContentsForbidden creates a PostLearnAPIPublicV1CoursesCourseIDContentsForbidden with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDContentsForbidden() *PostLearnAPIPublicV1CoursesCourseIDContentsForbidden {
	return &PostLearnAPIPublicV1CoursesCourseIDContentsForbidden{}
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsForbidden handles this case with default header values.

Forbidden
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/contents][%d] postLearnApiPublicV1CoursesCourseIdContentsForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage creates a PostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage with default headers values
func NewPostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage() *PostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage {
	return &PostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage{}
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage handles this case with default header values.

Folder quota exceeded
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/courses/{courseId}/contents][%d] postLearnApiPublicV1CoursesCourseIdContentsInsufficientStorage  %+v", 507, o.Payload)
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsInsufficientStorage) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsBody post learn API public v1 courses course ID contents body
swagger:model PostLearnAPIPublicV1CoursesCourseIDContentsBody
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsBody struct {

	// availability
	Availability *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability `json:"availability,omitempty"`

	// The body text associated with this content. This field supports BbML; see <a target='_blank' href='https://docs.blackboard.com/learn/REST/Blackboard%20Markup%20Language%20-%20BbML.html'>here</a> for more information.
	Body string `json:"body,omitempty"`

	// Extended settings specific to this content item's content handler.
	ContentHandler *models.ContentHandler `json:"contentHandler,omitempty"`

	// The short description of this content.
	//
	// This field is not used in Classic courses.  For Ultra courses this is used to show information directly on the course outline.
	Description string `json:"description,omitempty"`

	// Indicates whether the content is going to open in a new window.
	//
	// **Since**: 3800.10.0
	LaunchInNewWindow bool `json:"launchInNewWindow,omitempty"`

	// The ID of the content's parent.  Note that top-level contents do not have parents. The 'parentId' field is a writable field as of the Bb Learn 3200.6.0 release.  Specifying a new value in PATCH requests allows the Content object to be moved from one parent to another.
	ParentID string `json:"parentId,omitempty"`

	// The position of this content within its parent folder. Position values are zero-based (the first element has a position value of zero, not one). Default position is last in the list of child contents under the parent.
	Position int32 `json:"position,omitempty"`

	// Indicates whether Review Status is enabled for this content. Content items with review status enabled can be marked as reviewed by students. This can be used to track performance and in Adaptive Release rules to control the release of other content. Reviewable field is currently being used only in Classic courses.
	//
	// **Since**: 3700.15.0
	Reviewable bool `json:"reviewable,omitempty"`

	// The title or name of this content. Typically shown as the main text to click in the course outline when accessing the content.
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this post learn API public v1 courses course ID contents body
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateContentHandler(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBody) validateAvailability(formats strfmt.Registry) error {

	if swag.IsZero(o.Availability) { // not required
		return nil
	}

	if o.Availability != nil {
		if err := o.Availability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "availability")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBody) validateContentHandler(formats strfmt.Registry) error {

	if swag.IsZero(o.ContentHandler) { // not required
		return nil
	}

	if o.ContentHandler != nil {
		if err := o.ContentHandler.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "contentHandler")
			}
			return err
		}
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDContentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability Availability
//
// Settings controlling availability of the content to students.
swagger:model PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability struct {

	// adaptive release
	AdaptiveRelease *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease `json:"adaptiveRelease,omitempty"`

	// Whether this content is available to users with the 'guest' role. Defaults to true.
	AllowGuests bool `json:"allowGuests,omitempty"`

	// Whether the content is currently available to students.  Instructors can always access the content.  If set to 'PartiallyVisible', the title will be available to students but the body will not.  Defaults to Yes.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes |  |
	// | No |  |
	// | PartiallyVisible |  |
	//
	// Enum: [Yes No PartiallyVisible]
	Available string `json:"available,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID contents params body availability
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdaptiveRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability) validateAdaptiveRelease(formats strfmt.Registry) error {

	if swag.IsZero(o.AdaptiveRelease) { // not required
		return nil
	}

	if o.AdaptiveRelease != nil {
		if err := o.AdaptiveRelease.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "availability" + "." + "adaptiveRelease")
			}
			return err
		}
	}

	return nil
}

var postLearnApiPublicV1CoursesCourseIdContentsParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No","PartiallyVisible"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1CoursesCourseIdContentsParamsBodyAvailabilityTypeAvailablePropEnum = append(postLearnApiPublicV1CoursesCourseIdContentsParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAvailableYes string = "Yes"

	// PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAvailableNo captures enum value "No"
	PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAvailableNo string = "No"

	// PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAvailablePartiallyVisible captures enum value "PartiallyVisible"
	PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAvailablePartiallyVisible string = "PartiallyVisible"
)

// prop value enum
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1CoursesCourseIdContentsParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

	if swag.IsZero(o.Available) { // not required
		return nil
	}

	// value enum
	if err := o.validateAvailableEnum("input"+"."+"availability"+"."+"available", "body", o.Available); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease AdaptiveRelease
//
// Settings controlling adaptive release of the content to students.
swagger:model PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease
*/
type PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease struct {

	// The date when this content will no longer be available to students.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The date when this content will become available to students.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`
}

// Validate validates this post learn API public v1 courses course ID contents params body availability adaptive release
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(o.End) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"adaptiveRelease"+"."+"end", "body", "date-time", o.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(o.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("input"+"."+"availability"+"."+"adaptiveRelease"+"."+"start", "body", "date-time", o.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1CoursesCourseIDContentsParamsBodyAvailabilityAdaptiveRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
