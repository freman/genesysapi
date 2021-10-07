// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResultCounters result counters
//
// swagger:model ResultCounters
type ResultCounters struct {

	// failure
	Failure int32 `json:"failure,omitempty"`

	// success
	Success int32 `json:"success,omitempty"`
}

// Validate validates this result counters
func (m *ResultCounters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResultCounters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultCounters) UnmarshalBinary(b []byte) error {
	var res ResultCounters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
