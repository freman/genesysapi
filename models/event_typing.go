// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventTyping A Typing event.
//
// swagger:model EventTyping
type EventTyping struct {

	// The duration of the Typing event in milliseconds.
	// Read Only: true
	Duration int64 `json:"duration,omitempty"`

	// Describes the type of Typing event.
	// Required: true
	// Enum: [On]
	Type *string `json:"type"`
}

// Validate validates this event typing
func (m *EventTyping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var eventTypingTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["On"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventTypingTypeTypePropEnum = append(eventTypingTypeTypePropEnum, v)
	}
}

const (

	// EventTypingTypeOn captures enum value "On"
	EventTypingTypeOn string = "On"
)

// prop value enum
func (m *EventTyping) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventTypingTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EventTyping) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventTyping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventTyping) UnmarshalBinary(b []byte) error {
	var res EventTyping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
