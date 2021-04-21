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

// SendAgentlessOutboundMessageResponse send agentless outbound message response
//
// swagger:model SendAgentlessOutboundMessageResponse
type SendAgentlessOutboundMessageResponse struct {

	// The identifier of the conversation.
	ConversationID string `json:"conversationId,omitempty"`

	// The sender of the message.
	FromAddress string `json:"fromAddress,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The messaging template sent
	MessagingTemplate *MessagingTemplateRequest `json:"messagingTemplate,omitempty"`

	// Type of messenger.
	// Enum: [sms facebook twitter line whatsapp webmessaging open]
	MessengerType string `json:"messengerType,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The body of the text message.
	TextBody string `json:"textBody,omitempty"`

	// The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// The recipient of the message.
	ToAddress string `json:"toAddress,omitempty"`

	// Details of the user created the job
	// Read Only: true
	User *AddressableEntityRef `json:"user,omitempty"`
}

// Validate validates this send agentless outbound message response
func (m *SendAgentlessOutboundMessageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessagingTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessengerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendAgentlessOutboundMessageResponse) validateMessagingTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.MessagingTemplate) { // not required
		return nil
	}

	if m.MessagingTemplate != nil {
		if err := m.MessagingTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagingTemplate")
			}
			return err
		}
	}

	return nil
}

var sendAgentlessOutboundMessageResponseTypeMessengerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms","facebook","twitter","line","whatsapp","webmessaging","open"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sendAgentlessOutboundMessageResponseTypeMessengerTypePropEnum = append(sendAgentlessOutboundMessageResponseTypeMessengerTypePropEnum, v)
	}
}

const (

	// SendAgentlessOutboundMessageResponseMessengerTypeSms captures enum value "sms"
	SendAgentlessOutboundMessageResponseMessengerTypeSms string = "sms"

	// SendAgentlessOutboundMessageResponseMessengerTypeFacebook captures enum value "facebook"
	SendAgentlessOutboundMessageResponseMessengerTypeFacebook string = "facebook"

	// SendAgentlessOutboundMessageResponseMessengerTypeTwitter captures enum value "twitter"
	SendAgentlessOutboundMessageResponseMessengerTypeTwitter string = "twitter"

	// SendAgentlessOutboundMessageResponseMessengerTypeLine captures enum value "line"
	SendAgentlessOutboundMessageResponseMessengerTypeLine string = "line"

	// SendAgentlessOutboundMessageResponseMessengerTypeWhatsapp captures enum value "whatsapp"
	SendAgentlessOutboundMessageResponseMessengerTypeWhatsapp string = "whatsapp"

	// SendAgentlessOutboundMessageResponseMessengerTypeWebmessaging captures enum value "webmessaging"
	SendAgentlessOutboundMessageResponseMessengerTypeWebmessaging string = "webmessaging"

	// SendAgentlessOutboundMessageResponseMessengerTypeOpen captures enum value "open"
	SendAgentlessOutboundMessageResponseMessengerTypeOpen string = "open"
)

// prop value enum
func (m *SendAgentlessOutboundMessageResponse) validateMessengerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sendAgentlessOutboundMessageResponseTypeMessengerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SendAgentlessOutboundMessageResponse) validateMessengerType(formats strfmt.Registry) error {

	if swag.IsZero(m.MessengerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessengerTypeEnum("messengerType", "body", m.MessengerType); err != nil {
		return err
	}

	return nil
}

func (m *SendAgentlessOutboundMessageResponse) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SendAgentlessOutboundMessageResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SendAgentlessOutboundMessageResponse) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendAgentlessOutboundMessageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendAgentlessOutboundMessageResponse) UnmarshalBinary(b []byte) error {
	var res SendAgentlessOutboundMessageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
