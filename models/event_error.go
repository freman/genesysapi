// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventError event error
//
// swagger:model EventError
type EventError struct {

	// The eventId (V4 UUID) for the event that encountered an error.
	EventID string `json:"eventId,omitempty"`

	// A message describing the error.
	Message string `json:"message,omitempty"`

	// The event for this eventId can be resubmitted if this value is true.
	Retryable bool `json:"retryable"`
}

// Validate validates this event error
func (m *EventError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this event error based on context it is used
func (m *EventError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventError) UnmarshalBinary(b []byte) error {
	var res EventError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}