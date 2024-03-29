// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ParticipantAttributes participant attributes
//
// swagger:model ParticipantAttributes
type ParticipantAttributes struct {

	// The map of attribute keys to values.
	Attributes map[string]string `json:"attributes,omitempty"`
}

// Validate validates this participant attributes
func (m *ParticipantAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this participant attributes based on context it is used
func (m *ParticipantAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParticipantAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParticipantAttributes) UnmarshalBinary(b []byte) error {
	var res ParticipantAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
