// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Number number
//
// swagger:model Number
type Number struct {

	// end
	End string `json:"end,omitempty"`

	// start
	Start string `json:"start,omitempty"`
}

// Validate validates this number
func (m *Number) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this number based on context it is used
func (m *Number) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Number) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Number) UnmarshalBinary(b []byte) error {
	var res Number
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
