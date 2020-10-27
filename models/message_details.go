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

// MessageDetails message details
//
// swagger:model MessageDetails
type MessageDetails struct {

	// The media (images, files, etc) associated with this message, if any
	Media []*MessageMedia `json:"media"`

	// UUID identifying the message media.
	MessageID string `json:"messageId,omitempty"`

	// The message segment count, greater than 1 if the message content was split into multiple parts for this message type, e.g. SMS character limits.
	MessageSegmentCount int32 `json:"messageSegmentCount,omitempty"`

	// Indicates the delivery status of the message.
	// Enum: [queued sent failed received delivery-success delivery-failed read]
	MessageStatus string `json:"messageStatus,omitempty"`

	// The time when the message was sent or received. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	MessageTime strfmt.DateTime `json:"messageTime,omitempty"`

	// A URI for this message entity.
	// Format: uri
	MessageURI strfmt.URI `json:"messageURI,omitempty"`

	// One or more stickers associated with this message, if any
	Stickers []*MessageSticker `json:"stickers"`
}

// Validate validates this message details
func (m *MessageDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMedia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStickers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageDetails) validateMedia(formats strfmt.Registry) error {

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

var messageDetailsTypeMessageStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["queued","sent","failed","received","delivery-success","delivery-failed","read"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageDetailsTypeMessageStatusPropEnum = append(messageDetailsTypeMessageStatusPropEnum, v)
	}
}

const (

	// MessageDetailsMessageStatusQueued captures enum value "queued"
	MessageDetailsMessageStatusQueued string = "queued"

	// MessageDetailsMessageStatusSent captures enum value "sent"
	MessageDetailsMessageStatusSent string = "sent"

	// MessageDetailsMessageStatusFailed captures enum value "failed"
	MessageDetailsMessageStatusFailed string = "failed"

	// MessageDetailsMessageStatusReceived captures enum value "received"
	MessageDetailsMessageStatusReceived string = "received"

	// MessageDetailsMessageStatusDeliverySuccess captures enum value "delivery-success"
	MessageDetailsMessageStatusDeliverySuccess string = "delivery-success"

	// MessageDetailsMessageStatusDeliveryFailed captures enum value "delivery-failed"
	MessageDetailsMessageStatusDeliveryFailed string = "delivery-failed"

	// MessageDetailsMessageStatusRead captures enum value "read"
	MessageDetailsMessageStatusRead string = "read"
)

// prop value enum
func (m *MessageDetails) validateMessageStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageDetailsTypeMessageStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageDetails) validateMessageStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageStatusEnum("messageStatus", "body", m.MessageStatus); err != nil {
		return err
	}

	return nil
}

func (m *MessageDetails) validateMessageTime(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageTime) { // not required
		return nil
	}

	if err := validate.FormatOf("messageTime", "body", "date-time", m.MessageTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessageDetails) validateMessageURI(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageURI) { // not required
		return nil
	}

	if err := validate.FormatOf("messageURI", "body", "uri", m.MessageURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessageDetails) validateStickers(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *MessageDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageDetails) UnmarshalBinary(b []byte) error {
	var res MessageDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
