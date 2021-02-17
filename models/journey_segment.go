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

// JourneySegment journey segment
//
// swagger:model JourneySegment
type JourneySegment struct {

	// Time, in days, from when the segment is assigned until it is automatically unassigned.
	AssignmentExpirationDays int32 `json:"assignmentExpirationDays,omitempty"`

	// The hexadecimal color value of the segment.
	Color string `json:"color,omitempty"`

	// The context of the segment.
	Context *Context `json:"context,omitempty"`

	// Timestamp indicating when the segment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// A description of the segment.
	Description string `json:"description,omitempty"`

	// The display name of the segment.
	// Required: true
	DisplayName *string `json:"displayName"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Whether or not the segment is active.
	IsActive bool `json:"isActive"`

	// The pattern of rules defining the segment.
	Journey *Journey `json:"journey,omitempty"`

	// Timestamp indicating when the the segment was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The target entity that a segment applies to.
	// Enum: [Session Customer]
	Scope string `json:"scope,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Whether or not the segment should be displayed to agent/supervisor users.
	ShouldDisplayToAgent bool `json:"shouldDisplayToAgent"`

	// The version of the segment.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this journey segment
func (m *JourneySegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJourney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JourneySegment) validateContext(formats strfmt.Registry) error {

	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

func (m *JourneySegment) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JourneySegment) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *JourneySegment) validateJourney(formats strfmt.Registry) error {

	if swag.IsZero(m.Journey) { // not required
		return nil
	}

	if m.Journey != nil {
		if err := m.Journey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("journey")
			}
			return err
		}
	}

	return nil
}

func (m *JourneySegment) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var journeySegmentTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Session","Customer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journeySegmentTypeScopePropEnum = append(journeySegmentTypeScopePropEnum, v)
	}
}

const (

	// JourneySegmentScopeSession captures enum value "Session"
	JourneySegmentScopeSession string = "Session"

	// JourneySegmentScopeCustomer captures enum value "Customer"
	JourneySegmentScopeCustomer string = "Customer"
)

// prop value enum
func (m *JourneySegment) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journeySegmentTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JourneySegment) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *JourneySegment) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JourneySegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JourneySegment) UnmarshalBinary(b []byte) error {
	var res JourneySegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
