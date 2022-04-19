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

// WebMessagingEventPresence A Presence event.
//
// swagger:model WebMessagingEventPresence
type WebMessagingEventPresence struct {

	// Describes the type of Presence event.
	// Required: true
	// Enum: [Join]
	Type *string `json:"type"`
}

// Validate validates this web messaging event presence
func (m *WebMessagingEventPresence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webMessagingEventPresenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Join"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webMessagingEventPresenceTypeTypePropEnum = append(webMessagingEventPresenceTypeTypePropEnum, v)
	}
}

const (

	// WebMessagingEventPresenceTypeJoin captures enum value "Join"
	WebMessagingEventPresenceTypeJoin string = "Join"
)

// prop value enum
func (m *WebMessagingEventPresence) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webMessagingEventPresenceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebMessagingEventPresence) validateType(formats strfmt.Registry) error {

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
func (m *WebMessagingEventPresence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebMessagingEventPresence) UnmarshalBinary(b []byte) error {
	var res WebMessagingEventPresence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
