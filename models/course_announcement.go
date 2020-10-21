// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CourseAnnouncement course announcement
//
// swagger:model CourseAnnouncement
type CourseAnnouncement struct {

	// availability
	// Required: true
	Availability *CourseAnnouncementAvailability `json:"availability"`

	// The message body of the Announcement. This field supports BbML; see <a target='_blank' href='https://docs.blackboard.com/learn/REST/Blackboard%20Markup%20Language%20-%20BbML.html'>here</a> for more information.
	Body string `json:"body,omitempty"`

	// The date that the Announcement was created.
	// Required: true
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created"`

	// The user that created the Announcement.
	// Required: true
	// Read Only: true
	Creator string `json:"creator"`

	// An indication of whether or not the Announcement is in draft status.
	Draft bool `json:"draft,omitempty"`

	// Primary key identifier
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// modified
	// Required: true
	// Read Only: true
	// Format: date-time
	Modified strfmt.DateTime `json:"modified"`

	// The number of users that the Announcement will reach.
	//
	// Shown when adding the query parameter: "expand=readCount".
	// Required: true
	// Read Only: true
	Participants int32 `json:"participants"`

	// The position of the Announcement.
	// Required: true
	// Read Only: true
	Position int32 `json:"position"`

	// The number of reads for the Announcement.
	//
	// Shown when adding the query parameter: "expand=readCount".
	// Required: true
	// Read Only: true
	ReadCount int32 `json:"readCount"`

	// The title of this Announcement.
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this course announcement
func (m *CourseAnnouncement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CourseAnnouncement) validateAvailability(formats strfmt.Registry) error {

	if err := validate.Required("availability", "body", m.Availability); err != nil {
		return err
	}

	if m.Availability != nil {
		if err := m.Availability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availability")
			}
			return err
		}
	}

	return nil
}

func (m *CourseAnnouncement) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CourseAnnouncement) validateCreator(formats strfmt.Registry) error {

	if err := validate.RequiredString("creator", "body", string(m.Creator)); err != nil {
		return err
	}

	return nil
}

func (m *CourseAnnouncement) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CourseAnnouncement) validateModified(formats strfmt.Registry) error {

	if err := validate.Required("modified", "body", strfmt.DateTime(m.Modified)); err != nil {
		return err
	}

	if err := validate.FormatOf("modified", "body", "date-time", m.Modified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CourseAnnouncement) validateParticipants(formats strfmt.Registry) error {

	if err := validate.Required("participants", "body", int32(m.Participants)); err != nil {
		return err
	}

	return nil
}

func (m *CourseAnnouncement) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", int32(m.Position)); err != nil {
		return err
	}

	return nil
}

func (m *CourseAnnouncement) validateReadCount(formats strfmt.Registry) error {

	if err := validate.Required("readCount", "body", int32(m.ReadCount)); err != nil {
		return err
	}

	return nil
}

func (m *CourseAnnouncement) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CourseAnnouncement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CourseAnnouncement) UnmarshalBinary(b []byte) error {
	var res CourseAnnouncement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CourseAnnouncementAvailability Availability
//
// Settings controlling availability of the course to students.
//
// swagger:model CourseAnnouncementAvailability
type CourseAnnouncementAvailability struct {

	// duration
	// Required: true
	Duration *CourseAnnouncementAvailabilityDuration `json:"duration"`
}

// Validate validates this course announcement availability
func (m *CourseAnnouncementAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CourseAnnouncementAvailability) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("availability"+"."+"duration", "body", m.Duration); err != nil {
		return err
	}

	if m.Duration != nil {
		if err := m.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availability" + "." + "duration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CourseAnnouncementAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CourseAnnouncementAvailability) UnmarshalBinary(b []byte) error {
	var res CourseAnnouncementAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CourseAnnouncementAvailabilityDuration Duration
//
// Duration indicates when the System Announcement is Available based on whether it is Permanent or if the date/time of the request falls within its Start & End dates.
//
// swagger:model CourseAnnouncementAvailabilityDuration
type CourseAnnouncementAvailabilityDuration struct {

	// The date this Announcement stops being Available.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The date this Announcement starts being Available.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// Indicates whether this Course Announcement is always displayed (Permanent) or if it is shown only between the Start and End dates (Restricted).
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Permanent | The Announcement will always be displayed. |
	// | Restricted | The Announcement will start being displayed on Duration.Start and stop being displayed on Duration.End |
	//
	// Required: true
	// Read Only: true
	// Enum: [Permanent Restricted]
	Type string `json:"type"`
}

// Validate validates this course announcement availability duration
func (m *CourseAnnouncementAvailabilityDuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CourseAnnouncementAvailabilityDuration) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("availability"+"."+"duration"+"."+"end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CourseAnnouncementAvailabilityDuration) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("availability"+"."+"duration"+"."+"start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

var courseAnnouncementAvailabilityDurationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Permanent","Restricted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		courseAnnouncementAvailabilityDurationTypeTypePropEnum = append(courseAnnouncementAvailabilityDurationTypeTypePropEnum, v)
	}
}

const (

	// CourseAnnouncementAvailabilityDurationTypePermanent captures enum value "Permanent"
	CourseAnnouncementAvailabilityDurationTypePermanent string = "Permanent"

	// CourseAnnouncementAvailabilityDurationTypeRestricted captures enum value "Restricted"
	CourseAnnouncementAvailabilityDurationTypeRestricted string = "Restricted"
)

// prop value enum
func (m *CourseAnnouncementAvailabilityDuration) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, courseAnnouncementAvailabilityDurationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CourseAnnouncementAvailabilityDuration) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("availability"+"."+"duration"+"."+"type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("availability"+"."+"duration"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CourseAnnouncementAvailabilityDuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CourseAnnouncementAvailabilityDuration) UnmarshalBinary(b []byte) error {
	var res CourseAnnouncementAvailabilityDuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}