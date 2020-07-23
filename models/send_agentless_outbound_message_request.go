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

// SendAgentlessOutboundMessageRequest send agentless outbound message request
//
// swagger:model SendAgentlessOutboundMessageRequest
type SendAgentlessOutboundMessageRequest struct {

	// The messaging address of the sender of the message. For an SMS messenger type, this must be a currently provisioned sms phone number.
	// Required: true
	FromAddress *string `json:"fromAddress"`

	// The text of the message to send
	// Required: true
	TextBody *string `json:"textBody"`

	// The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234
	// Required: true
	ToAddress *string `json:"toAddress"`

	// The recipient messaging address messenger type.
	// Required: true
	// Enum: [sms facebook twitter line whatsapp]
	ToAddressMessengerType *string `json:"toAddressMessengerType"`
}

// Validate validates this send agentless outbound message request
func (m *SendAgentlessOutboundMessageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToAddressMessengerType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendAgentlessOutboundMessageRequest) validateFromAddress(formats strfmt.Registry) error {

	if err := validate.Required("fromAddress", "body", m.FromAddress); err != nil {
		return err
	}

	return nil
}

func (m *SendAgentlessOutboundMessageRequest) validateTextBody(formats strfmt.Registry) error {

	if err := validate.Required("textBody", "body", m.TextBody); err != nil {
		return err
	}

	return nil
}

func (m *SendAgentlessOutboundMessageRequest) validateToAddress(formats strfmt.Registry) error {

	if err := validate.Required("toAddress", "body", m.ToAddress); err != nil {
		return err
	}

	return nil
}

var sendAgentlessOutboundMessageRequestTypeToAddressMessengerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms","facebook","twitter","line","whatsapp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sendAgentlessOutboundMessageRequestTypeToAddressMessengerTypePropEnum = append(sendAgentlessOutboundMessageRequestTypeToAddressMessengerTypePropEnum, v)
	}
}

const (

	// SendAgentlessOutboundMessageRequestToAddressMessengerTypeSms captures enum value "sms"
	SendAgentlessOutboundMessageRequestToAddressMessengerTypeSms string = "sms"

	// SendAgentlessOutboundMessageRequestToAddressMessengerTypeFacebook captures enum value "facebook"
	SendAgentlessOutboundMessageRequestToAddressMessengerTypeFacebook string = "facebook"

	// SendAgentlessOutboundMessageRequestToAddressMessengerTypeTwitter captures enum value "twitter"
	SendAgentlessOutboundMessageRequestToAddressMessengerTypeTwitter string = "twitter"

	// SendAgentlessOutboundMessageRequestToAddressMessengerTypeLine captures enum value "line"
	SendAgentlessOutboundMessageRequestToAddressMessengerTypeLine string = "line"

	// SendAgentlessOutboundMessageRequestToAddressMessengerTypeWhatsapp captures enum value "whatsapp"
	SendAgentlessOutboundMessageRequestToAddressMessengerTypeWhatsapp string = "whatsapp"
)

// prop value enum
func (m *SendAgentlessOutboundMessageRequest) validateToAddressMessengerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sendAgentlessOutboundMessageRequestTypeToAddressMessengerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SendAgentlessOutboundMessageRequest) validateToAddressMessengerType(formats strfmt.Registry) error {

	if err := validate.Required("toAddressMessengerType", "body", m.ToAddressMessengerType); err != nil {
		return err
	}

	// value enum
	if err := m.validateToAddressMessengerTypeEnum("toAddressMessengerType", "body", *m.ToAddressMessengerType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendAgentlessOutboundMessageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendAgentlessOutboundMessageRequest) UnmarshalBinary(b []byte) error {
	var res SendAgentlessOutboundMessageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
