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

// InboundOnlySetting inbound only setting
//
// swagger:model InboundOnlySetting
type InboundOnlySetting struct {

	// inbound
	// Enum: [Enabled Disabled]
	Inbound string `json:"inbound,omitempty"`
}

// Validate validates this inbound only setting
func (m *InboundOnlySetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInbound(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var inboundOnlySettingTypeInboundPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Enabled","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inboundOnlySettingTypeInboundPropEnum = append(inboundOnlySettingTypeInboundPropEnum, v)
	}
}

const (

	// InboundOnlySettingInboundEnabled captures enum value "Enabled"
	InboundOnlySettingInboundEnabled string = "Enabled"

	// InboundOnlySettingInboundDisabled captures enum value "Disabled"
	InboundOnlySettingInboundDisabled string = "Disabled"
)

// prop value enum
func (m *InboundOnlySetting) validateInboundEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, inboundOnlySettingTypeInboundPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InboundOnlySetting) validateInbound(formats strfmt.Registry) error {
	if swag.IsZero(m.Inbound) { // not required
		return nil
	}

	// value enum
	if err := m.validateInboundEnum("inbound", "body", m.Inbound); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this inbound only setting based on context it is used
func (m *InboundOnlySetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InboundOnlySetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InboundOnlySetting) UnmarshalBinary(b []byte) error {
	var res InboundOnlySetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
