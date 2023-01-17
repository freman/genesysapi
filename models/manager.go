// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Manager Defines a SCIM manager.
//
// swagger:model Manager
type Manager struct {

	// The reference URI of the manager's user record.
	// Read Only: true
	// Format: uri
	DollarRef strfmt.URI `json:"$ref,omitempty"`

	// The ID of the manager.
	Value string `json:"value,omitempty"`
}

// Validate validates this manager
func (m *Manager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDollarRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manager) validateDollarRef(formats strfmt.Registry) error {
	if swag.IsZero(m.DollarRef) { // not required
		return nil
	}

	if err := validate.FormatOf("$ref", "body", "uri", m.DollarRef.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this manager based on the context it is used
func (m *Manager) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDollarRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Manager) contextValidateDollarRef(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "$ref", "body", strfmt.URI(m.DollarRef)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Manager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Manager) UnmarshalBinary(b []byte) error {
	var res Manager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
