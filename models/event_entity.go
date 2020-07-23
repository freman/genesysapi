// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventEntity event entity
//
// swagger:model EventEntity
type EventEntity struct {

	// Type of entity the event pertains to. e.g. integration
	EntityType string `json:"entityType,omitempty"`

	// ID of the entity the event pertains to.
	ID string `json:"id,omitempty"`
}

// Validate validates this event entity
func (m *EventEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventEntity) UnmarshalBinary(b []byte) error {
	var res EventEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
