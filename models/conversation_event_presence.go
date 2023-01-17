// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConversationEventPresence A Presence event.
//
// swagger:model ConversationEventPresence
type ConversationEventPresence struct {

	// Describes the type of Presence event.
	// Required: true
	// Enum: [Join Disconnect]
	Type *string `json:"type"`
}

// Validate validates this conversation event presence
func (m *ConversationEventPresence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversationEventPresenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Join","Disconnect"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationEventPresenceTypeTypePropEnum = append(conversationEventPresenceTypeTypePropEnum, v)
	}
}

const (

	// ConversationEventPresenceTypeJoin captures enum value "Join"
	ConversationEventPresenceTypeJoin string = "Join"

	// ConversationEventPresenceTypeDisconnect captures enum value "Disconnect"
	ConversationEventPresenceTypeDisconnect string = "Disconnect"
)

// prop value enum
func (m *ConversationEventPresence) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationEventPresenceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationEventPresence) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conversation event presence based on context it is used
func (m *ConversationEventPresence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConversationEventPresence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationEventPresence) UnmarshalBinary(b []byte) error {
	var res ConversationEventPresence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
