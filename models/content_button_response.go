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

// ContentButtonResponse Button response object representing the click of a structured message button, such as a quick reply.
//
// swagger:model ContentButtonResponse
type ContentButtonResponse struct {

	// An ID assigned to the button response (Deprecated).
	ID string `json:"id,omitempty"`

	// The response payload associated with the clicked button.
	// Required: true
	Payload *string `json:"payload"`

	// The response text from the button click.
	// Required: true
	Text *string `json:"text"`

	// Describes the button that resulted in the Button Response.
	// Required: true
	// Enum: [Button QuickReply]
	Type *string `json:"type"`
}

// Validate validates this content button response
func (m *ContentButtonResponse) Validate(formats strfmt.Registry) error {
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

func (m *ContentButtonResponse) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

func (m *ContentButtonResponse) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

var contentButtonResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Button","QuickReply"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentButtonResponseTypeTypePropEnum = append(contentButtonResponseTypeTypePropEnum, v)
	}
}

const (

	// ContentButtonResponseTypeButton captures enum value "Button"
	ContentButtonResponseTypeButton string = "Button"

	// ContentButtonResponseTypeQuickReply captures enum value "QuickReply"
	ContentButtonResponseTypeQuickReply string = "QuickReply"
)

// prop value enum
func (m *ContentButtonResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contentButtonResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContentButtonResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this content button response based on context it is used
func (m *ContentButtonResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContentButtonResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentButtonResponse) UnmarshalBinary(b []byte) error {
	var res ContentButtonResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
