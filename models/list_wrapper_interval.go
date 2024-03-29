// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListWrapperInterval list wrapper interval
//
// swagger:model ListWrapperInterval
type ListWrapperInterval struct {

	// values
	Values []string `json:"values"`
}

// Validate validates this list wrapper interval
func (m *ListWrapperInterval) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list wrapper interval based on context it is used
func (m *ListWrapperInterval) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListWrapperInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListWrapperInterval) UnmarshalBinary(b []byte) error {
	var res ListWrapperInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
