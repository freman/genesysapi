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

// ScimV2EnterpriseUser Defines a SCIM enterprise user.
//
// swagger:model ScimV2EnterpriseUser
type ScimV2EnterpriseUser struct {

	// The department that the user belongs to.
	Department string `json:"department,omitempty"`

	// The division that the user belongs to.
	Division string `json:"division,omitempty"`

	// The user's employee number.
	EmployeeNumber string `json:"employeeNumber,omitempty"`

	// The user's manager.
	Manager *Manager `json:"manager,omitempty"`
}

// Validate validates this scim v2 enterprise user
func (m *ScimV2EnterpriseUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManager(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimV2EnterpriseUser) validateManager(formats strfmt.Registry) error {
	if swag.IsZero(m.Manager) { // not required
		return nil
	}

	if m.Manager != nil {
		if err := m.Manager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manager")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this scim v2 enterprise user based on the context it is used
func (m *ScimV2EnterpriseUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimV2EnterpriseUser) contextValidateManager(ctx context.Context, formats strfmt.Registry) error {

	if m.Manager != nil {
		if err := m.Manager.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manager")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimV2EnterpriseUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimV2EnterpriseUser) UnmarshalBinary(b []byte) error {
	var res ScimV2EnterpriseUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
