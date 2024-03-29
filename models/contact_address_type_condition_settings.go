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

// ContactAddressTypeConditionSettings contact address type condition settings
//
// swagger:model ContactAddressTypeConditionSettings
type ContactAddressTypeConditionSettings struct {

	// The operator to use when comparing the address types.
	// Required: true
	// Enum: [Equals Contains BeginsWith EndsWith]
	Operator *string `json:"operator"`

	// The type value to compare against the contact column type.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this contact address type condition settings
func (m *ContactAddressTypeConditionSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var contactAddressTypeConditionSettingsTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Equals","Contains","BeginsWith","EndsWith"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactAddressTypeConditionSettingsTypeOperatorPropEnum = append(contactAddressTypeConditionSettingsTypeOperatorPropEnum, v)
	}
}

const (

	// ContactAddressTypeConditionSettingsOperatorEquals captures enum value "Equals"
	ContactAddressTypeConditionSettingsOperatorEquals string = "Equals"

	// ContactAddressTypeConditionSettingsOperatorContains captures enum value "Contains"
	ContactAddressTypeConditionSettingsOperatorContains string = "Contains"

	// ContactAddressTypeConditionSettingsOperatorBeginsWith captures enum value "BeginsWith"
	ContactAddressTypeConditionSettingsOperatorBeginsWith string = "BeginsWith"

	// ContactAddressTypeConditionSettingsOperatorEndsWith captures enum value "EndsWith"
	ContactAddressTypeConditionSettingsOperatorEndsWith string = "EndsWith"
)

// prop value enum
func (m *ContactAddressTypeConditionSettings) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contactAddressTypeConditionSettingsTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContactAddressTypeConditionSettings) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", *m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *ContactAddressTypeConditionSettings) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this contact address type condition settings based on context it is used
func (m *ContactAddressTypeConditionSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactAddressTypeConditionSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAddressTypeConditionSettings) UnmarshalBinary(b []byte) error {
	var res ContactAddressTypeConditionSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
