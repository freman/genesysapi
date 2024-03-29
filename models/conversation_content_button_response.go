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

// ConversationContentButtonResponse Button response object representing the click of a structured message button, such as a quick reply.
//
// swagger:model ConversationContentButtonResponse
type ConversationContentButtonResponse struct {

	// The response payload associated with the clicked button.
	// Required: true
	Payload *string `json:"payload"`

	// The response text from the button click.
	// Required: true
	Text *string `json:"text"`

	// Describes the button that resulted in the Button Response.
	// Enum: [Button QuickReply]
	Type string `json:"type,omitempty"`
}

// Validate validates this conversation content button response
func (m *ConversationContentButtonResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
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

func (m *ConversationContentButtonResponse) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

func (m *ConversationContentButtonResponse) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

var conversationContentButtonResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Button","QuickReply"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationContentButtonResponseTypeTypePropEnum = append(conversationContentButtonResponseTypeTypePropEnum, v)
	}
}

const (

	// ConversationContentButtonResponseTypeButton captures enum value "Button"
	ConversationContentButtonResponseTypeButton string = "Button"

	// ConversationContentButtonResponseTypeQuickReply captures enum value "QuickReply"
	ConversationContentButtonResponseTypeQuickReply string = "QuickReply"
)

// prop value enum
func (m *ConversationContentButtonResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationContentButtonResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationContentButtonResponse) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conversation content button response based on context it is used
func (m *ConversationContentButtonResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConversationContentButtonResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationContentButtonResponse) UnmarshalBinary(b []byte) error {
	var res ConversationContentButtonResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
