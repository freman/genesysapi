// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmbeddedIntegration embedded integration
//
// swagger:model EmbeddedIntegration
type EmbeddedIntegration struct {

	// domain whitelist
	DomainWhitelist []string `json:"domainWhitelist"`

	// enable whitelist
	EnableWhitelist bool `json:"enableWhitelist,omitempty"`
}

// Validate validates this embedded integration
func (m *EmbeddedIntegration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmbeddedIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmbeddedIntegration) UnmarshalBinary(b []byte) error {
	var res EmbeddedIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
