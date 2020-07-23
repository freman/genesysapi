// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeInterval time interval
//
// swagger:model TimeInterval
type TimeInterval struct {

	// days
	Days int32 `json:"days,omitempty"`

	// hours
	Hours int32 `json:"hours,omitempty"`

	// months
	Months int32 `json:"months,omitempty"`

	// weeks
	Weeks int32 `json:"weeks,omitempty"`
}

// Validate validates this time interval
func (m *TimeInterval) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeInterval) UnmarshalBinary(b []byte) error {
	var res TimeInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
