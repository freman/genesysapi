// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuSchedulingSettingsRequest bu scheduling settings request
//
// swagger:model BuSchedulingSettingsRequest
type BuSchedulingSettingsRequest struct {

	// Schedule generation message severity configuration
	MessageSeverities []*SchedulerMessageTypeSeverity `json:"messageSeverities"`
}

// Validate validates this bu scheduling settings request
func (m *BuSchedulingSettingsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageSeverities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuSchedulingSettingsRequest) validateMessageSeverities(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageSeverities) { // not required
		return nil
	}

	for i := 0; i < len(m.MessageSeverities); i++ {
		if swag.IsZero(m.MessageSeverities[i]) { // not required
			continue
		}

		if m.MessageSeverities[i] != nil {
			if err := m.MessageSeverities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messageSeverities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuSchedulingSettingsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuSchedulingSettingsRequest) UnmarshalBinary(b []byte) error {
	var res BuSchedulingSettingsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
