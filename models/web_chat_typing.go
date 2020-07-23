// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebChatTyping web chat typing
//
// swagger:model WebChatTyping
type WebChatTyping struct {

	// The identifier of the conversation
	// Required: true
	Conversation *WebChatConversation `json:"conversation"`

	// The event identifier of this typing indicator event (useful to guard against event re-delivery
	// Required: true
	ID *string `json:"id"`

	// The member who sent the message
	// Required: true
	Sender *WebChatMemberInfo `json:"sender"`

	// The timestamp of the message, in ISO-8601 format
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this web chat typing
func (m *WebChatTyping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
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

func (m *WebChatTyping) validateConversation(formats strfmt.Registry) error {

	if err := validate.Required("conversation", "body", m.Conversation); err != nil {
		return err
	}

	if m.Conversation != nil {
		if err := m.Conversation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversation")
			}
			return err
		}
	}

	return nil
}

func (m *WebChatTyping) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *WebChatTyping) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	if m.Sender != nil {
		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

func (m *WebChatTyping) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebChatTyping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebChatTyping) UnmarshalBinary(b []byte) error {
	var res WebChatTyping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
