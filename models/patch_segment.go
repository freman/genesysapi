// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchSegment patch segment
//
// swagger:model PatchSegment
type PatchSegment struct {

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

	// Details of an entity corresponding to this segment in an external system.
	ExternalSegment *PatchExternalSegment `json:"externalSegment,omitempty"`

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

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Whether or not the segment should be displayed to agent/supervisor users.
	ShouldDisplayToAgent bool `json:"shouldDisplayToAgent"`

	// The version of the segment.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this patch segment
func (m *PatchSegment) Validate(formats strfmt.Registry) error {
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

	if err := m.validateExternalSegment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJourney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
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

func (m *PatchSegment) validateContext(formats strfmt.Registry) error {

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

func (m *PatchSegment) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PatchSegment) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *PatchSegment) validateExternalSegment(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalSegment) { // not required
		return nil
	}

	if m.ExternalSegment != nil {
		if err := m.ExternalSegment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalSegment")
			}
			return err
		}
	}

	return nil
}

func (m *PatchSegment) validateJourney(formats strfmt.Registry) error {

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

func (m *PatchSegment) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PatchSegment) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchSegment) UnmarshalBinary(b []byte) error {
	var res PatchSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
