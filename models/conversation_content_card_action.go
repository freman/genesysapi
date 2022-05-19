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

// ConversationContentCardAction CardAction Object
//
// swagger:model ConversationContentCardAction
type ConversationContentCardAction struct {

	// Text to be returned as the payload from a ButtonResponse when a button is clicked.
	Payload string `json:"payload,omitempty"`

	// The response text from the button click.
	Text string `json:"text,omitempty"`

	// Describes the type of action.
	// Enum: [Link Postback Unknown]
	Type string `json:"type,omitempty"`

	// A URL of a web page to direct the user to.
	URL string `json:"url,omitempty"`
}

// Validate validates this conversation content card action
func (m *ConversationContentCardAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversationContentCardActionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Link","Postback","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationContentCardActionTypeTypePropEnum = append(conversationContentCardActionTypeTypePropEnum, v)
	}
}

const (

	// ConversationContentCardActionTypeLink captures enum value "Link"
	ConversationContentCardActionTypeLink string = "Link"

	// ConversationContentCardActionTypePostback captures enum value "Postback"
	ConversationContentCardActionTypePostback string = "Postback"

	// ConversationContentCardActionTypeUnknown captures enum value "Unknown"
	ConversationContentCardActionTypeUnknown string = "Unknown"
)

// prop value enum
func (m *ConversationContentCardAction) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationContentCardActionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationContentCardAction) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationContentCardAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationContentCardAction) UnmarshalBinary(b []byte) error {
	var res ConversationContentCardAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}