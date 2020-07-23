// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConnectRate connect rate
//
// swagger:model ConnectRate
type ConnectRate struct {

	// Number of call attempts made
	// Read Only: true
	Attempts int64 `json:"attempts,omitempty"`

	// Ratio of connects to attempts
	// Read Only: true
	ConnectRatio float64 `json:"connectRatio,omitempty"`

	// Number of calls with a live voice detected
	// Read Only: true
	Connects int64 `json:"connects,omitempty"`
}

// Validate validates this connect rate
func (m *ConnectRate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectRate) UnmarshalBinary(b []byte) error {
	var res ConnectRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
