// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MaxParticipants max participants
//
// swagger:model MaxParticipants
type MaxParticipants struct {

	// The maximum number of participants that are allowed on a conversation.
	MaxParticipants int32 `json:"maxParticipants,omitempty"`
}

// Validate validates this max participants
func (m *MaxParticipants) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MaxParticipants) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaxParticipants) UnmarshalBinary(b []byte) error {
	var res MaxParticipants
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
