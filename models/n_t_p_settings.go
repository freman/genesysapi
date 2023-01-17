// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NTPSettings n t p settings
//
// swagger:model NTPSettings
type NTPSettings struct {

	// List of NTP servers, in priority order
	Servers []string `json:"servers"`
}

// Validate validates this n t p settings
func (m *NTPSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this n t p settings based on context it is used
func (m *NTPSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NTPSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NTPSettings) UnmarshalBinary(b []byte) error {
	var res NTPSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
