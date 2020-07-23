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

// RoleDivision role division
//
// swagger:model RoleDivision
type RoleDivision struct {

	// Division associated with the given role which forms a grant
	// Required: true
	DivisionID *string `json:"divisionId"`

	// Role to be associated with the given division which forms a grant
	// Required: true
	RoleID *string `json:"roleId"`
}

// Validate validates this role division
func (m *RoleDivision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivisionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleDivision) validateDivisionID(formats strfmt.Registry) error {

	if err := validate.Required("divisionId", "body", m.DivisionID); err != nil {
		return err
	}

	return nil
}

func (m *RoleDivision) validateRoleID(formats strfmt.Registry) error {

	if err := validate.Required("roleId", "body", m.RoleID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleDivision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleDivision) UnmarshalBinary(b []byte) error {
	var res RoleDivision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
