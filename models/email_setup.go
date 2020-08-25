// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailSetup email setup
//
// swagger:model EmailSetup
type EmailSetup struct {

	// The root PureCloud domain that all sub-domains are created from.
	RootDomain string `json:"rootDomain,omitempty"`
}

// Validate validates this email setup
func (m *EmailSetup) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailSetup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailSetup) UnmarshalBinary(b []byte) error {
	var res EmailSetup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}