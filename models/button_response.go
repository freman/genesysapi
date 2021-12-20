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

// ButtonResponse button response
//
// swagger:model ButtonResponse
type ButtonResponse struct {

	// Button response message type that captures QuickReply , Cards and Carousel .This is used  as label for Card selection
	// Enum: [QuickReply Card Carousel]
	MessageType string `json:"messageType,omitempty"`

	// Content of the textback payload after clicking a button
	Payload string `json:"payload,omitempty"`

	// Text to show inside the Button reply. This is also used as the response text after clicking on the Button.
	Text string `json:"text,omitempty"`

	// Button response type that captures Button and QuickReply type responses
	// Enum: [Button QuickReply]
	Type string `json:"type,omitempty"`
}

// Validate validates this button response
func (m *ButtonResponse) Validate(formats strfmt.Registry) error {
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

var buttonResponseTypeMessageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["QuickReply","Card","Carousel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buttonResponseTypeMessageTypePropEnum = append(buttonResponseTypeMessageTypePropEnum, v)
	}
}

const (

	// ButtonResponseMessageTypeQuickReply captures enum value "QuickReply"
	ButtonResponseMessageTypeQuickReply string = "QuickReply"

	// ButtonResponseMessageTypeCard captures enum value "Card"
	ButtonResponseMessageTypeCard string = "Card"

	// ButtonResponseMessageTypeCarousel captures enum value "Carousel"
	ButtonResponseMessageTypeCarousel string = "Carousel"
)

// prop value enum
func (m *ButtonResponse) validateMessageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buttonResponseTypeMessageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ButtonResponse) validateMessageType(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageTypeEnum("messageType", "body", m.MessageType); err != nil {
		return err
	}

	return nil
}

var buttonResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Button","QuickReply"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buttonResponseTypeTypePropEnum = append(buttonResponseTypeTypePropEnum, v)
	}
}

const (

	// ButtonResponseTypeButton captures enum value "Button"
	ButtonResponseTypeButton string = "Button"

	// ButtonResponseTypeQuickReply captures enum value "QuickReply"
	ButtonResponseTypeQuickReply string = "QuickReply"
)

// prop value enum
func (m *ButtonResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buttonResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ButtonResponse) validateType(formats strfmt.Registry) error {

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
func (m *ButtonResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ButtonResponse) UnmarshalBinary(b []byte) error {
	var res ButtonResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
