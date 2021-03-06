// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MessageData message data
//
// swagger:model MessageData
type MessageData struct {

	// User who sent this message.
	CreatedBy *User `json:"createdBy,omitempty"`

	// The direction of the message.
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// The sender of the text message.
	FromAddress string `json:"fromAddress,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The media details associated to a message.
	Media []*MessageMedia `json:"media"`

	// Type of text messenger.
	// Enum: [sms facebook twitter line whatsapp webmessaging]
	MessengerType string `json:"messengerType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The unique identifier of the message from provider
	ProviderMessageID string `json:"providerMessageId,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The status of the message.
	// Required: true
	// Enum: [queued sent failed received delivery-success delivery-failed read]
	Status *string `json:"status"`

	// The sticker details associated to a message.
	Stickers []*MessageSticker `json:"stickers"`

	// The body of the text message.
	// Required: true
	TextBody *string `json:"textBody"`

	// The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`

	// The recipient of the text message.
	ToAddress string `json:"toAddress,omitempty"`
}

// Validate validates this message data
func (m *MessageData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessengerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStickers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageData) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

var messageDataTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageDataTypeDirectionPropEnum = append(messageDataTypeDirectionPropEnum, v)
	}
}

const (

	// MessageDataDirectionInbound captures enum value "inbound"
	MessageDataDirectionInbound string = "inbound"

	// MessageDataDirectionOutbound captures enum value "outbound"
	MessageDataDirectionOutbound string = "outbound"
)

// prop value enum
func (m *MessageData) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageDataTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageData) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *MessageData) validateMedia(formats strfmt.Registry) error {

	if swag.IsZero(m.Media) { // not required
		return nil
	}

	for i := 0; i < len(m.Media); i++ {
		if swag.IsZero(m.Media[i]) { // not required
			continue
		}

		if m.Media[i] != nil {
			if err := m.Media[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("media" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var messageDataTypeMessengerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms","facebook","twitter","line","whatsapp","webmessaging"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageDataTypeMessengerTypePropEnum = append(messageDataTypeMessengerTypePropEnum, v)
	}
}

const (

	// MessageDataMessengerTypeSms captures enum value "sms"
	MessageDataMessengerTypeSms string = "sms"

	// MessageDataMessengerTypeFacebook captures enum value "facebook"
	MessageDataMessengerTypeFacebook string = "facebook"

	// MessageDataMessengerTypeTwitter captures enum value "twitter"
	MessageDataMessengerTypeTwitter string = "twitter"

	// MessageDataMessengerTypeLine captures enum value "line"
	MessageDataMessengerTypeLine string = "line"

	// MessageDataMessengerTypeWhatsapp captures enum value "whatsapp"
	MessageDataMessengerTypeWhatsapp string = "whatsapp"

	// MessageDataMessengerTypeWebmessaging captures enum value "webmessaging"
	MessageDataMessengerTypeWebmessaging string = "webmessaging"
)

// prop value enum
func (m *MessageData) validateMessengerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageDataTypeMessengerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageData) validateMessengerType(formats strfmt.Registry) error {

	if swag.IsZero(m.MessengerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessengerTypeEnum("messengerType", "body", m.MessengerType); err != nil {
		return err
	}

	return nil
}

func (m *MessageData) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var messageDataTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["queued","sent","failed","received","delivery-success","delivery-failed","read"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageDataTypeStatusPropEnum = append(messageDataTypeStatusPropEnum, v)
	}
}

const (

	// MessageDataStatusQueued captures enum value "queued"
	MessageDataStatusQueued string = "queued"

	// MessageDataStatusSent captures enum value "sent"
	MessageDataStatusSent string = "sent"

	// MessageDataStatusFailed captures enum value "failed"
	MessageDataStatusFailed string = "failed"

	// MessageDataStatusReceived captures enum value "received"
	MessageDataStatusReceived string = "received"

	// MessageDataStatusDeliverySuccess captures enum value "delivery-success"
	MessageDataStatusDeliverySuccess string = "delivery-success"

	// MessageDataStatusDeliveryFailed captures enum value "delivery-failed"
	MessageDataStatusDeliveryFailed string = "delivery-failed"

	// MessageDataStatusRead captures enum value "read"
	MessageDataStatusRead string = "read"
)

// prop value enum
func (m *MessageData) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageDataTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageData) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *MessageData) validateStickers(formats strfmt.Registry) error {

	if swag.IsZero(m.Stickers) { // not required
		return nil
	}

	for i := 0; i < len(m.Stickers); i++ {
		if swag.IsZero(m.Stickers[i]) { // not required
			continue
		}

		if m.Stickers[i] != nil {
			if err := m.Stickers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stickers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessageData) validateTextBody(formats strfmt.Registry) error {

	if err := validate.Required("textBody", "body", m.TextBody); err != nil {
		return err
	}

	return nil
}

func (m *MessageData) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageData) UnmarshalBinary(b []byte) error {
	var res MessageData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
