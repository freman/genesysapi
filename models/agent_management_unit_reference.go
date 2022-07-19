// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AgentManagementUnitReference agent management unit reference
//
// swagger:model AgentManagementUnitReference
type AgentManagementUnitReference struct {

	// The business unit to which the user (agent) belongs. Populate with expand=businessUnit
	BusinessUnit *BusinessUnitReference `json:"businessUnit,omitempty"`

	// The management to which the user (agent) belongs
	ManagementUnit *ManagementUnitReference `json:"managementUnit,omitempty"`

	// The user (agent) for whom the management unit was requested
	User *UserReference `json:"user,omitempty"`
}

// Validate validates this agent management unit reference
func (m *AgentManagementUnitReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentManagementUnitReference) validateBusinessUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.BusinessUnit) { // not required
		return nil
	}

	if m.BusinessUnit != nil {
		if err := m.BusinessUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessUnit")
			}
			return err
		}
	}

	return nil
}

func (m *AgentManagementUnitReference) validateManagementUnit(formats strfmt.Registry) error {

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

func (m *AgentManagementUnitReference) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentManagementUnitReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentManagementUnitReference) UnmarshalBinary(b []byte) error {
	var res AgentManagementUnitReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
