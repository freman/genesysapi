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

// ConversationMessagingChannel Channel-specific information that describes the message and the message channel/provider.
//
// swagger:model ConversationMessagingChannel
type ConversationMessagingChannel struct {

	// Time the message was deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateDeleted strfmt.DateTime `json:"dateDeleted,omitempty"`

	// Time the message was edited. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Information about the recipient the message is received from.
	// Read Only: true
	From *ConversationMessagingFromRecipient `json:"from,omitempty"`

	// The integration ID.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Unique provider ID of the message such as a Facebook message ID.
	// Read Only: true
	MessageID string `json:"messageId,omitempty"`

	// The provider type.
	// Read Only: true
	// Enum: [Twitter Facebook Instagram Line Whatsapp WebMessaging Open Sms]
	Platform string `json:"platform,omitempty"`

	// Original time of the event. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// Information about the recipient the message is sent to.
	// Read Only: true
	To *ConversationMessagingToRecipient `json:"to,omitempty"`
}

// Validate validates this conversation messaging channel
func (m *ConversationMessagingChannel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationMessagingChannel) validateDateDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.DateDeleted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateDeleted", "body", "date-time", m.DateDeleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConversationMessagingChannel) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConversationMessagingChannel) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

var conversationMessagingChannelTypePlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Twitter","Facebook","Instagram","Line","Whatsapp","WebMessaging","Open","Sms"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationMessagingChannelTypePlatformPropEnum = append(conversationMessagingChannelTypePlatformPropEnum, v)
	}
}

const (

	// ConversationMessagingChannelPlatformTwitter captures enum value "Twitter"
	ConversationMessagingChannelPlatformTwitter string = "Twitter"

	// ConversationMessagingChannelPlatformFacebook captures enum value "Facebook"
	ConversationMessagingChannelPlatformFacebook string = "Facebook"

	// ConversationMessagingChannelPlatformInstagram captures enum value "Instagram"
	ConversationMessagingChannelPlatformInstagram string = "Instagram"

	// ConversationMessagingChannelPlatformLine captures enum value "Line"
	ConversationMessagingChannelPlatformLine string = "Line"

	// ConversationMessagingChannelPlatformWhatsapp captures enum value "Whatsapp"
	ConversationMessagingChannelPlatformWhatsapp string = "Whatsapp"

	// ConversationMessagingChannelPlatformWebMessaging captures enum value "WebMessaging"
	ConversationMessagingChannelPlatformWebMessaging string = "WebMessaging"

	// ConversationMessagingChannelPlatformOpen captures enum value "Open"
	ConversationMessagingChannelPlatformOpen string = "Open"

	// ConversationMessagingChannelPlatformSms captures enum value "Sms"
	ConversationMessagingChannelPlatformSms string = "Sms"
)

// prop value enum
func (m *ConversationMessagingChannel) validatePlatformEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationMessagingChannelTypePlatformPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationMessagingChannel) validatePlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformEnum("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *ConversationMessagingChannel) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConversationMessagingChannel) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if m.To != nil {
		if err := m.To.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationMessagingChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationMessagingChannel) UnmarshalBinary(b []byte) error {
	var res ConversationMessagingChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
