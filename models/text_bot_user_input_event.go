// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TextBotUserInputEvent Settings for an input event to the bot flow indicating user input is available.
//
// swagger:model TextBotUserInputEvent
type TextBotUserInputEvent struct {

	// The input alternatives.
	// Required: true
	Alternatives []*TextBotUserInputAlternative `json:"alternatives"`

	// The input mode.
	// Required: true
	// Enum: [Text]
	Mode *string `json:"mode"`
}

// Validate validates this text bot user input event
func (m *TextBotUserInputEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternatives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextBotUserInputEvent) validateAlternatives(formats strfmt.Registry) error {

	if err := validate.Required("alternatives", "body", m.Alternatives); err != nil {
		return err
	}

	for i := 0; i < len(m.Alternatives); i++ {
		if swag.IsZero(m.Alternatives[i]) { // not required
			continue
		}

		if m.Alternatives[i] != nil {
			if err := m.Alternatives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alternatives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alternatives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var textBotUserInputEventTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Text"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		textBotUserInputEventTypeModePropEnum = append(textBotUserInputEventTypeModePropEnum, v)
	}
}

const (

	// TextBotUserInputEventModeText captures enum value "Text"
	TextBotUserInputEventModeText string = "Text"
)

// prop value enum
func (m *TextBotUserInputEvent) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, textBotUserInputEventTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TextBotUserInputEvent) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this text bot user input event based on the context it is used
func (m *TextBotUserInputEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlternatives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextBotUserInputEvent) contextValidateAlternatives(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Alternatives); i++ {

		if m.Alternatives[i] != nil {
			if err := m.Alternatives[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alternatives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alternatives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextBotUserInputEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextBotUserInputEvent) UnmarshalBinary(b []byte) error {
	var res TextBotUserInputEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
