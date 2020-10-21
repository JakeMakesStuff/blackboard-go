// Code generated by go-swagger; DO NOT EDIT.

package lti

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

// PostLearnAPIPublicV1LtiPlacementsReader is a Reader for the PostLearnAPIPublicV1LtiPlacements structure.
type PostLearnAPIPublicV1LtiPlacementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnAPIPublicV1LtiPlacementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnAPIPublicV1LtiPlacementsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnAPIPublicV1LtiPlacementsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnAPIPublicV1LtiPlacementsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostLearnAPIPublicV1LtiPlacementsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnAPIPublicV1LtiPlacementsCreated creates a PostLearnAPIPublicV1LtiPlacementsCreated with default headers values
func NewPostLearnAPIPublicV1LtiPlacementsCreated() *PostLearnAPIPublicV1LtiPlacementsCreated {
	return &PostLearnAPIPublicV1LtiPlacementsCreated{}
}

/*PostLearnAPIPublicV1LtiPlacementsCreated handles this case with default header values.

Created
*/
type PostLearnAPIPublicV1LtiPlacementsCreated struct {
	Payload *models.LTIPlacement
}

func (o *PostLearnAPIPublicV1LtiPlacementsCreated) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/lti/placements][%d] postLearnApiPublicV1LtiPlacementsCreated  %+v", 201, o.Payload)
}

func (o *PostLearnAPIPublicV1LtiPlacementsCreated) GetPayload() *models.LTIPlacement {
	return o.Payload
}

func (o *PostLearnAPIPublicV1LtiPlacementsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LTIPlacement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1LtiPlacementsBadRequest creates a PostLearnAPIPublicV1LtiPlacementsBadRequest with default headers values
func NewPostLearnAPIPublicV1LtiPlacementsBadRequest() *PostLearnAPIPublicV1LtiPlacementsBadRequest {
	return &PostLearnAPIPublicV1LtiPlacementsBadRequest{}
}

/*PostLearnAPIPublicV1LtiPlacementsBadRequest handles this case with default header values.

Bad Request
*/
type PostLearnAPIPublicV1LtiPlacementsBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1LtiPlacementsBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/lti/placements][%d] postLearnApiPublicV1LtiPlacementsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnAPIPublicV1LtiPlacementsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1LtiPlacementsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1LtiPlacementsForbidden creates a PostLearnAPIPublicV1LtiPlacementsForbidden with default headers values
func NewPostLearnAPIPublicV1LtiPlacementsForbidden() *PostLearnAPIPublicV1LtiPlacementsForbidden {
	return &PostLearnAPIPublicV1LtiPlacementsForbidden{}
}

/*PostLearnAPIPublicV1LtiPlacementsForbidden handles this case with default header values.

Forbidden
*/
type PostLearnAPIPublicV1LtiPlacementsForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1LtiPlacementsForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/lti/placements][%d] postLearnApiPublicV1LtiPlacementsForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnAPIPublicV1LtiPlacementsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1LtiPlacementsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnAPIPublicV1LtiPlacementsConflict creates a PostLearnAPIPublicV1LtiPlacementsConflict with default headers values
func NewPostLearnAPIPublicV1LtiPlacementsConflict() *PostLearnAPIPublicV1LtiPlacementsConflict {
	return &PostLearnAPIPublicV1LtiPlacementsConflict{}
}

/*PostLearnAPIPublicV1LtiPlacementsConflict handles this case with default header values.

Conflict
*/
type PostLearnAPIPublicV1LtiPlacementsConflict struct {
	Payload *models.RestException
}

func (o *PostLearnAPIPublicV1LtiPlacementsConflict) Error() string {
	return fmt.Sprintf("[POST /learn/api/public/v1/lti/placements][%d] postLearnApiPublicV1LtiPlacementsConflict  %+v", 409, o.Payload)
}

func (o *PostLearnAPIPublicV1LtiPlacementsConflict) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnAPIPublicV1LtiPlacementsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnAPIPublicV1LtiPlacementsBody post learn API public v1 lti placements body
swagger:model PostLearnAPIPublicV1LtiPlacementsBody
*/
type PostLearnAPIPublicV1LtiPlacementsBody struct {

	// Whether this placement can accept grades from the tool provider and a grade column can be created for it. This only applies to Enum[Type]#`ContentHandler` types. All others don't support grading and will be set to false.
	AllowGrading bool `json:"allowGrading,omitempty"`

	// Whether the course tool is visible by students, or only to non-students (e.g. instructors). Defaults to true, allowing students to see the tool.
	AllowStudents bool `json:"allowStudents,omitempty"`

	// Id of the creator of the placement
	AuthorID string `json:"authorId,omitempty"`

	// availability
	Availability *PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability `json:"availability,omitempty"`

	// Custom launch parameters for the tool.
	CustomParameters map[string]string `json:"customParameters,omitempty"`

	// The description of the placement. Not required to be unique. Maximum length is 1000 characters, BAD_REQUEST error with message is returned if greater than 1000 characters.
	// Max Length: 1000
	Description string `json:"description,omitempty"`

	// The handle that uniquely identifies this placement. Required to be unique. Maximum length is 32 characters, BAD_REQUEST error with message is returned if greater than 32 characters.
	// Required: true
	// Max Length: 32
	Handle *string `json:"handle"`

	// The URL of the icon for this placement, if any. Not required to be unique, must be a complete and valid URL. Maximum length is 255 characters, BAD_REQUEST error with message is returned if greater than 255 characters or incomplete URL.
	// Max Length: 255
	IconURL string `json:"iconUrl,omitempty"`

	// Whether an instructor created the placement or not (otherwise admin)
	InstructorCreated bool `json:"instructorCreated,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// Whether this placement link should be opened in a new window or not.
	LaunchInNewWindow bool `json:"launchInNewWindow,omitempty"`

	// The name of the placement. Not required to be unique. Maximum length of 50 characters, BAD_REQUEST error with message is returned if greater than 50 characters.
	// Required: true
	// Max Length: 50
	Name *string `json:"name"`

	// secret
	Secret string `json:"secret,omitempty"`

	// The type of placement.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Application | Application or Student Tool Placement |
	// | ContentHandler | Content Type placement |
	// | ContentItemMessage | Content-item Message placement (see IMSGlobal spec)  **Since**: 3300.5.0 |
	// | System | System-level Tools |
	// | Administrator | Administrator-level Tools  **Since**: 3400.1.0 |
	// | UltraUI | Ultra-UI Extensions  **Since**: 3700.6.0 |
	// | BaseNavigation | Base Navigation |
	// | CourseNavigation | Course Navigation |
	//
	// Required: true
	// Enum: [Application ContentHandler ContentItemMessage System Administrator UltraUI BaseNavigation CourseNavigation]
	Type *string `json:"type"`

	// The URL of the tool provider. Not required to be unique, must be a complete and valid URL. Maximum length is 1024 characters, BAD_REQUEST error with message is returned if greater than 1024 characters or incomplete URL.
	// Required: true
	// Max Length: 1024
	URL *string `json:"url"`
}

// Validate validates this post learn API public v1 lti placements body
func (o *PostLearnAPIPublicV1LtiPlacementsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIconURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostLearnAPIPublicV1LtiPlacementsBody) validateAvailability(formats strfmt.Registry) error {

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

func (o *PostLearnAPIPublicV1LtiPlacementsBody) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"description", "body", string(o.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1LtiPlacementsBody) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"handle", "body", o.Handle); err != nil {
		return err
	}

	if err := validate.MaxLength("input"+"."+"handle", "body", string(*o.Handle), 32); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1LtiPlacementsBody) validateIconURL(formats strfmt.Registry) error {

	if swag.IsZero(o.IconURL) { // not required
		return nil
	}

	if err := validate.MaxLength("input"+"."+"iconUrl", "body", string(o.IconURL), 255); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1LtiPlacementsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("input"+"."+"name", "body", string(*o.Name), 50); err != nil {
		return err
	}

	return nil
}

var postLearnApiPublicV1LtiPlacementsBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Application","ContentHandler","ContentItemMessage","System","Administrator","UltraUI","BaseNavigation","CourseNavigation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1LtiPlacementsBodyTypeTypePropEnum = append(postLearnApiPublicV1LtiPlacementsBodyTypeTypePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1LtiPlacementsBodyTypeApplication captures enum value "Application"
	PostLearnAPIPublicV1LtiPlacementsBodyTypeApplication string = "Application"

	// PostLearnAPIPublicV1LtiPlacementsBodyTypeContentHandler captures enum value "ContentHandler"
	PostLearnAPIPublicV1LtiPlacementsBodyTypeContentHandler string = "ContentHandler"

	// PostLearnAPIPublicV1LtiPlacementsBodyTypeContentItemMessage captures enum value "ContentItemMessage"
	PostLearnAPIPublicV1LtiPlacementsBodyTypeContentItemMessage string = "ContentItemMessage"

	// PostLearnAPIPublicV1LtiPlacementsBodyTypeSystem captures enum value "System"
	PostLearnAPIPublicV1LtiPlacementsBodyTypeSystem string = "System"

	// PostLearnAPIPublicV1LtiPlacementsBodyTypeAdministrator captures enum value "Administrator"
	PostLearnAPIPublicV1LtiPlacementsBodyTypeAdministrator string = "Administrator"

	// PostLearnAPIPublicV1LtiPlacementsBodyTypeUltraUI captures enum value "UltraUI"
	PostLearnAPIPublicV1LtiPlacementsBodyTypeUltraUI string = "UltraUI"

	// PostLearnAPIPublicV1LtiPlacementsBodyTypeBaseNavigation captures enum value "BaseNavigation"
	PostLearnAPIPublicV1LtiPlacementsBodyTypeBaseNavigation string = "BaseNavigation"

	// PostLearnAPIPublicV1LtiPlacementsBodyTypeCourseNavigation captures enum value "CourseNavigation"
	PostLearnAPIPublicV1LtiPlacementsBodyTypeCourseNavigation string = "CourseNavigation"
)

// prop value enum
func (o *PostLearnAPIPublicV1LtiPlacementsBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1LtiPlacementsBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1LtiPlacementsBody) validateType(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("input"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

func (o *PostLearnAPIPublicV1LtiPlacementsBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("input"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	if err := validate.MaxLength("input"+"."+"url", "body", string(*o.URL), 1024); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnAPIPublicV1LtiPlacementsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1LtiPlacementsBody) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1LtiPlacementsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability Availability
//
// Settings controlling availability of the placement.
swagger:model PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability
*/
type PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability struct {

	// Whether the placement is available within the system.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Yes |  |
	// | No |  |
	//
	// Enum: [Yes No]
	Available string `json:"available,omitempty"`
}

// Validate validates this post learn API public v1 lti placements params body availability
func (o *PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postLearnApiPublicV1LtiPlacementsParamsBodyAvailabilityTypeAvailablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yes","No"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApiPublicV1LtiPlacementsParamsBodyAvailabilityTypeAvailablePropEnum = append(postLearnApiPublicV1LtiPlacementsParamsBodyAvailabilityTypeAvailablePropEnum, v)
	}
}

const (

	// PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailabilityAvailableYes captures enum value "Yes"
	PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailabilityAvailableYes string = "Yes"

	// PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailabilityAvailableNo captures enum value "No"
	PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailabilityAvailableNo string = "No"
)

// prop value enum
func (o *PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability) validateAvailableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApiPublicV1LtiPlacementsParamsBodyAvailabilityTypeAvailablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability) validateAvailable(formats strfmt.Registry) error {

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
func (o *PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability) UnmarshalBinary(b []byte) error {
	var res PostLearnAPIPublicV1LtiPlacementsParamsBodyAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}