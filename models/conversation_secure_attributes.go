// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConversationSecureAttributes conversation secure attributes
//
// swagger:model ConversationSecureAttributes
type ConversationSecureAttributes struct {

	// The map of attribute keys to values.
	Attributes map[string]string `json:"attributes,omitempty"`

	// The version used to detect conflicting updates when using PUT. Not used for PATCH.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this conversation secure attributes
func (m *ConversationSecureAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConversationSecureAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationSecureAttributes) UnmarshalBinary(b []byte) error {
	var res ConversationSecureAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}