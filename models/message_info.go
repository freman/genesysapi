// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessageInfo message info
//
// swagger:model MessageInfo
type MessageInfo struct {

	// Key that can be used to localize the message.
	LocalizableMessageCode string `json:"localizableMessageCode,omitempty"`

	// Description of the message.
	Message string `json:"message,omitempty"`

	// Map with fields for variable replacement.
	MessageParams map[string]string `json:"messageParams,omitempty"`

	// Message with template fields for variable replacement.
	MessageWithParams string `json:"messageWithParams,omitempty"`
}

// Validate validates this message info
func (m *MessageInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this message info based on context it is used
func (m *MessageInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MessageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageInfo) UnmarshalBinary(b []byte) error {
	var res MessageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
