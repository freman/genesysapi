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

// ConversationChannel conversation channel
//
// swagger:model ConversationChannel
type ConversationChannel struct {

	// Message type for messaging conversations.
	// Enum: [Unknown Sms Twitter Facebook Line WhatsApp WebMessaging Open Instagram]
	MessageType string `json:"messageType,omitempty"`

	// The source provider for the conversation (e.g. Edge, PureCloud Messaging, PureCloud Email).
	Platform string `json:"platform,omitempty"`

	// The type or category of this channel.
	// Enum: [Unknown Call Callback Email GenericObject Messaging Social Webchat Voice Chat Cobrowse Video Screenshare Message]
	Type string `json:"type,omitempty"`
}

// Validate validates this conversation channel
func (m *ConversationChannel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversationChannelTypeMessageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Sms","Twitter","Facebook","Line","WhatsApp","WebMessaging","Open","Instagram"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationChannelTypeMessageTypePropEnum = append(conversationChannelTypeMessageTypePropEnum, v)
	}
}

const (

	// ConversationChannelMessageTypeUnknown captures enum value "Unknown"
	ConversationChannelMessageTypeUnknown string = "Unknown"

	// ConversationChannelMessageTypeSms captures enum value "Sms"
	ConversationChannelMessageTypeSms string = "Sms"

	// ConversationChannelMessageTypeTwitter captures enum value "Twitter"
	ConversationChannelMessageTypeTwitter string = "Twitter"

	// ConversationChannelMessageTypeFacebook captures enum value "Facebook"
	ConversationChannelMessageTypeFacebook string = "Facebook"

	// ConversationChannelMessageTypeLine captures enum value "Line"
	ConversationChannelMessageTypeLine string = "Line"

	// ConversationChannelMessageTypeWhatsApp captures enum value "WhatsApp"
	ConversationChannelMessageTypeWhatsApp string = "WhatsApp"

	// ConversationChannelMessageTypeWebMessaging captures enum value "WebMessaging"
	ConversationChannelMessageTypeWebMessaging string = "WebMessaging"

	// ConversationChannelMessageTypeOpen captures enum value "Open"
	ConversationChannelMessageTypeOpen string = "Open"

	// ConversationChannelMessageTypeInstagram captures enum value "Instagram"
	ConversationChannelMessageTypeInstagram string = "Instagram"
)

// prop value enum
func (m *ConversationChannel) validateMessageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationChannelTypeMessageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationChannel) validateMessageType(formats strfmt.Registry) error {
	if swag.IsZero(m.MessageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageTypeEnum("messageType", "body", m.MessageType); err != nil {
		return err
	}

	return nil
}

var conversationChannelTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Call","Callback","Email","GenericObject","Messaging","Social","Webchat","Voice","Chat","Cobrowse","Video","Screenshare","Message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationChannelTypeTypePropEnum = append(conversationChannelTypeTypePropEnum, v)
	}
}

const (

	// ConversationChannelTypeUnknown captures enum value "Unknown"
	ConversationChannelTypeUnknown string = "Unknown"

	// ConversationChannelTypeCall captures enum value "Call"
	ConversationChannelTypeCall string = "Call"

	// ConversationChannelTypeCallback captures enum value "Callback"
	ConversationChannelTypeCallback string = "Callback"

	// ConversationChannelTypeEmail captures enum value "Email"
	ConversationChannelTypeEmail string = "Email"

	// ConversationChannelTypeGenericObject captures enum value "GenericObject"
	ConversationChannelTypeGenericObject string = "GenericObject"

	// ConversationChannelTypeMessaging captures enum value "Messaging"
	ConversationChannelTypeMessaging string = "Messaging"

	// ConversationChannelTypeSocial captures enum value "Social"
	ConversationChannelTypeSocial string = "Social"

	// ConversationChannelTypeWebchat captures enum value "Webchat"
	ConversationChannelTypeWebchat string = "Webchat"

	// ConversationChannelTypeVoice captures enum value "Voice"
	ConversationChannelTypeVoice string = "Voice"

	// ConversationChannelTypeChat captures enum value "Chat"
	ConversationChannelTypeChat string = "Chat"

	// ConversationChannelTypeCobrowse captures enum value "Cobrowse"
	ConversationChannelTypeCobrowse string = "Cobrowse"

	// ConversationChannelTypeVideo captures enum value "Video"
	ConversationChannelTypeVideo string = "Video"

	// ConversationChannelTypeScreenshare captures enum value "Screenshare"
	ConversationChannelTypeScreenshare string = "Screenshare"

	// ConversationChannelTypeMessage captures enum value "Message"
	ConversationChannelTypeMessage string = "Message"
)

// prop value enum
func (m *ConversationChannel) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationChannelTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationChannel) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conversation channel based on context it is used
func (m *ConversationChannel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConversationChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationChannel) UnmarshalBinary(b []byte) error {
	var res ConversationChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
