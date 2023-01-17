// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CrossPlatformPolicyUpdate cross platform policy update
//
// swagger:model CrossPlatformPolicyUpdate
type CrossPlatformPolicyUpdate struct {

	// enabled
	Enabled bool `json:"enabled"`
}

// Validate validates this cross platform policy update
func (m *CrossPlatformPolicyUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cross platform policy update based on context it is used
func (m *CrossPlatformPolicyUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CrossPlatformPolicyUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrossPlatformPolicyUpdate) UnmarshalBinary(b []byte) error {
	var res CrossPlatformPolicyUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
