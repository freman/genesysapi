// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StationSettings Organization settings for stations
//
// swagger:model StationSettings
type StationSettings struct {

	// Configuration options for free-seating
	FreeSeatingConfiguration *FreeSeatingConfiguration `json:"freeSeatingConfiguration,omitempty"`
}

// Validate validates this station settings
func (m *StationSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFreeSeatingConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StationSettings) validateFreeSeatingConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.FreeSeatingConfiguration) { // not required
		return nil
	}

	if m.FreeSeatingConfiguration != nil {
		if err := m.FreeSeatingConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeSeatingConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("freeSeatingConfiguration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this station settings based on the context it is used
func (m *StationSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFreeSeatingConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StationSettings) contextValidateFreeSeatingConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.FreeSeatingConfiguration != nil {
		if err := m.FreeSeatingConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeSeatingConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("freeSeatingConfiguration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StationSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StationSettings) UnmarshalBinary(b []byte) error {
	var res StationSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
