// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NumericRange numeric range
//
// swagger:model NumericRange
type NumericRange struct {

	// Greater than
	Gt float64 `json:"gt,omitempty"`

	// Greater than or equal to
	Gte float64 `json:"gte,omitempty"`

	// Less than
	Lt float64 `json:"lt,omitempty"`

	// Less than or equal to
	Lte float64 `json:"lte,omitempty"`
}

// Validate validates this numeric range
func (m *NumericRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this numeric range based on context it is used
func (m *NumericRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NumericRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NumericRange) UnmarshalBinary(b []byte) error {
	var res NumericRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
