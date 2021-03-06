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

// ActivityCodeContainer Container for a map of ActivityCodeId to ActivityCode
//
// swagger:model ActivityCodeContainer
type ActivityCodeContainer struct {

	// Map of activity code id to activity code
	ActivityCodes map[string]ActivityCode `json:"activityCodes,omitempty"`

	// Version metadata for the associated management unit's list of activity codes
	// Required: true
	Metadata *WfmVersionedEntityMetadata `json:"metadata"`
}

// Validate validates this activity code container
func (m *ActivityCodeContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityCodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityCodeContainer) validateActivityCodes(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityCodes) { // not required
		return nil
	}

	for k := range m.ActivityCodes {

		if err := validate.Required("activityCodes"+"."+k, "body", m.ActivityCodes[k]); err != nil {
			return err
		}
		if val, ok := m.ActivityCodes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ActivityCodeContainer) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityCodeContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityCodeContainer) UnmarshalBinary(b []byte) error {
	var res ActivityCodeContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
