// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReschedulingManagementUnitResponse rescheduling management unit response
//
// swagger:model ReschedulingManagementUnitResponse
type ReschedulingManagementUnitResponse struct {

	// Whether the rescheduling run is applied for the given management unit
	Applied bool `json:"applied,omitempty"`

	// The management unit
	ManagementUnit *ManagementUnitReference `json:"managementUnit,omitempty"`
}

// Validate validates this rescheduling management unit response
func (m *ReschedulingManagementUnitResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagementUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReschedulingManagementUnitResponse) validateManagementUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagementUnit) { // not required
		return nil
	}

	if m.ManagementUnit != nil {
		if err := m.ManagementUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managementUnit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReschedulingManagementUnitResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReschedulingManagementUnitResponse) UnmarshalBinary(b []byte) error {
	var res ReschedulingManagementUnitResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
