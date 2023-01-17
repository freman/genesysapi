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

// SupportCenterCustomMessage support center custom message
//
// swagger:model SupportCenterCustomMessage
type SupportCenterCustomMessage struct {

	// Default value for the message
	DefaultValue string `json:"defaultValue,omitempty"`

	// Type of the message
	// Enum: [Welcome Fallback]
	Type string `json:"type,omitempty"`
}

// Validate validates this support center custom message
func (m *SupportCenterCustomMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var supportCenterCustomMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Welcome","Fallback"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supportCenterCustomMessageTypeTypePropEnum = append(supportCenterCustomMessageTypeTypePropEnum, v)
	}
}

const (

	// SupportCenterCustomMessageTypeWelcome captures enum value "Welcome"
	SupportCenterCustomMessageTypeWelcome string = "Welcome"

	// SupportCenterCustomMessageTypeFallback captures enum value "Fallback"
	SupportCenterCustomMessageTypeFallback string = "Fallback"
)

// prop value enum
func (m *SupportCenterCustomMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, supportCenterCustomMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SupportCenterCustomMessage) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this support center custom message based on context it is used
func (m *SupportCenterCustomMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupportCenterCustomMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportCenterCustomMessage) UnmarshalBinary(b []byte) error {
	var res SupportCenterCustomMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
