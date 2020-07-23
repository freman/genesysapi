// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgeAutoUpdateConfig edge auto update config
//
// swagger:model EdgeAutoUpdateConfig
type EdgeAutoUpdateConfig struct {

	// Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	End string `json:"end,omitempty"`

	// rrule
	Rrule string `json:"rrule,omitempty"`

	// Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	Start string `json:"start,omitempty"`

	// time zone
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this edge auto update config
func (m *EdgeAutoUpdateConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgeAutoUpdateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeAutoUpdateConfig) UnmarshalBinary(b []byte) error {
	var res EdgeAutoUpdateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
