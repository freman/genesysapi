// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScimUserRole Defines a user role.
//
// swagger:model ScimUserRole
type ScimUserRole struct {

	// The role of the Genesys Cloud user.
	Value string `json:"value,omitempty"`
}

// Validate validates this scim user role
func (m *ScimUserRole) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scim user role based on context it is used
func (m *ScimUserRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScimUserRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimUserRole) UnmarshalBinary(b []byte) error {
	var res ScimUserRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
