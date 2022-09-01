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

// ConversationAppSettings Conversation settings that handles chats within the messenger
//
// swagger:model ConversationAppSettings
type ConversationAppSettings struct {

	// The auto start for the messenger conversation
	AutoStart *AutoStart `json:"autoStart,omitempty"`

	// Deprecated. The auto start type for the messenger conversation
	// Enum: [Standard Automatic]
	AutoStartType string `json:"autoStartType,omitempty"`

	// The conversation disconnect settings for the messenger app
	ConversationDisconnect *ConversationDisconnectSettings `json:"conversationDisconnect,omitempty"`

	// The markdown for the messenger app
	Markdown *Markdown `json:"markdown,omitempty"`

	// The toggle to enable or disable typing indicator for messenger
	ShowAgentTypingIndicator bool `json:"showAgentTypingIndicator"`

	// The toggle to enable or disable typing indicator for messenger
	ShowUserTypingIndicator bool `json:"showUserTypingIndicator"`
}

// Validate validates this conversation app settings
func (m *ConversationAppSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoStartType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationDisconnect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkdown(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationAppSettings) validateAutoStart(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoStart) { // not required
		return nil
	}

	if m.AutoStart != nil {
		if err := m.AutoStart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoStart")
			}
			return err
		}
	}

	return nil
}

var conversationAppSettingsTypeAutoStartTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Standard","Automatic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationAppSettingsTypeAutoStartTypePropEnum = append(conversationAppSettingsTypeAutoStartTypePropEnum, v)
	}
}

const (

	// ConversationAppSettingsAutoStartTypeStandard captures enum value "Standard"
	ConversationAppSettingsAutoStartTypeStandard string = "Standard"

	// ConversationAppSettingsAutoStartTypeAutomatic captures enum value "Automatic"
	ConversationAppSettingsAutoStartTypeAutomatic string = "Automatic"
)

// prop value enum
func (m *ConversationAppSettings) validateAutoStartTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationAppSettingsTypeAutoStartTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationAppSettings) validateAutoStartType(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoStartType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAutoStartTypeEnum("autoStartType", "body", m.AutoStartType); err != nil {
		return err
	}

	return nil
}

func (m *ConversationAppSettings) validateConversationDisconnect(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationDisconnect) { // not required
		return nil
	}

	if m.ConversationDisconnect != nil {
		if err := m.ConversationDisconnect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversationDisconnect")
			}
			return err
		}
	}

	return nil
}

func (m *ConversationAppSettings) validateMarkdown(formats strfmt.Registry) error {

	if swag.IsZero(m.Markdown) { // not required
		return nil
	}

	if m.Markdown != nil {
		if err := m.Markdown.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("markdown")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationAppSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationAppSettings) UnmarshalBinary(b []byte) error {
	var res ConversationAppSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
