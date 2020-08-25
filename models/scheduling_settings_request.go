// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchedulingSettingsRequest Scheduling Settings
//
// swagger:model SchedulingSettingsRequest
type SchedulingSettingsRequest struct {

	// Default shrinkage percent for scheduling
	DefaultShrinkagePercent float64 `json:"defaultShrinkagePercent,omitempty"`

	// Max occupancy percent for deferred work
	MaxOccupancyPercentForDeferredWork int32 `json:"maxOccupancyPercentForDeferredWork,omitempty"`

	// Shrinkage overrides for scheduling
	ShrinkageOverrides *ShrinkageOverrides `json:"shrinkageOverrides,omitempty"`
}

// Validate validates this scheduling settings request
func (m *SchedulingSettingsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShrinkageOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulingSettingsRequest) validateShrinkageOverrides(formats strfmt.Registry) error {

	if swag.IsZero(m.ShrinkageOverrides) { // not required
		return nil
	}

	if m.ShrinkageOverrides != nil {
		if err := m.ShrinkageOverrides.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shrinkageOverrides")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulingSettingsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulingSettingsRequest) UnmarshalBinary(b []byte) error {
	var res SchedulingSettingsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}