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

// WebMessagingQuickReply Quick reply object
//
// swagger:model WebMessagingQuickReply
type WebMessagingQuickReply struct {

	// Specifies the type of action that is triggered upon clicking the quick reply.
	// Enum: [Message]
	Action string `json:"action,omitempty"`

	// URL of an image associated with the quick reply.
	Image string `json:"image,omitempty"`

	// Content of the payload included in the quick reply response. Could be an ID identifying the quick reply response.
	// Required: true
	Payload *string `json:"payload"`

	// Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this web messaging quick reply
func (m *WebMessagingQuickReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webMessagingQuickReplyTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webMessagingQuickReplyTypeActionPropEnum = append(webMessagingQuickReplyTypeActionPropEnum, v)
	}
}

const (

	// WebMessagingQuickReplyActionMessage captures enum value "Message"
	WebMessagingQuickReplyActionMessage string = "Message"
)

// prop value enum
func (m *WebMessagingQuickReply) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webMessagingQuickReplyTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebMessagingQuickReply) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *WebMessagingQuickReply) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

func (m *WebMessagingQuickReply) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this web messaging quick reply based on context it is used
func (m *WebMessagingQuickReply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebMessagingQuickReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebMessagingQuickReply) UnmarshalBinary(b []byte) error {
	var res WebMessagingQuickReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
