// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrunkErrorInfoDetails trunk error info details
//
// swagger:model TrunkErrorInfoDetails
type TrunkErrorInfoDetails struct {

	// code
	Code string `json:"code,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this trunk error info details
func (m *TrunkErrorInfoDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this trunk error info details based on context it is used
func (m *TrunkErrorInfoDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrunkErrorInfoDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrunkErrorInfoDetails) UnmarshalBinary(b []byte) error {
	var res TrunkErrorInfoDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
