// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConversationAppSettings Conversation settings that handles chats within the messenger
//
// swagger:model ConversationAppSettings
type ConversationAppSettings struct {

	// The toggle to enable or disable typing indicator for messenger
	ShowAgentTypingIndicator bool `json:"showAgentTypingIndicator"`

	// The toggle to enable or disable typing indicator for messenger
	ShowUserTypingIndicator bool `json:"showUserTypingIndicator"`
}

// Validate validates this conversation app settings
func (m *ConversationAppSettings) Validate(formats strfmt.Registry) error {
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
