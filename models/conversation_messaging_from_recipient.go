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

// ConversationMessagingFromRecipient Information about the recipient the message is received from.
//
// swagger:model ConversationMessagingFromRecipient
type ConversationMessagingFromRecipient struct {

	// E-mail address of the recipient.
	// Read Only: true
	Email string `json:"email,omitempty"`

	// First name of the recipient.
	FirstName string `json:"firstName,omitempty"`

	// The recipient ID specific to the provider.
	// Required: true
	ID *string `json:"id"`

	// The recipient ID type. This is used to indicate the format used for the ID.
	// Read Only: true
	// Enum: [Email Phone Opaque]
	IDType string `json:"idType,omitempty"`

	// URL of an image that represents the recipient.
	Image string `json:"image,omitempty"`

	// Last name of the recipient.
	LastName string `json:"lastName,omitempty"`

	// Nickname or display name of the recipient.
	// Read Only: true
	Nickname string `json:"nickname,omitempty"`
}

// Validate validates this conversation messaging from recipient
func (m *ConversationMessagingFromRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationMessagingFromRecipient) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var conversationMessagingFromRecipientTypeIDTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Email","Phone","Opaque"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationMessagingFromRecipientTypeIDTypePropEnum = append(conversationMessagingFromRecipientTypeIDTypePropEnum, v)
	}
}

const (

	// ConversationMessagingFromRecipientIDTypeEmail captures enum value "Email"
	ConversationMessagingFromRecipientIDTypeEmail string = "Email"

	// ConversationMessagingFromRecipientIDTypePhone captures enum value "Phone"
	ConversationMessagingFromRecipientIDTypePhone string = "Phone"

	// ConversationMessagingFromRecipientIDTypeOpaque captures enum value "Opaque"
	ConversationMessagingFromRecipientIDTypeOpaque string = "Opaque"
)

// prop value enum
func (m *ConversationMessagingFromRecipient) validateIDTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationMessagingFromRecipientTypeIDTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationMessagingFromRecipient) validateIDType(formats strfmt.Registry) error {

	if swag.IsZero(m.IDType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDTypeEnum("idType", "body", m.IDType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationMessagingFromRecipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationMessagingFromRecipient) UnmarshalBinary(b []byte) error {
	var res ConversationMessagingFromRecipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
