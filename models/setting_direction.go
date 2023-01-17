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

// SettingDirection setting direction
//
// swagger:model SettingDirection
type SettingDirection struct {

	// Status for the Inbound Direction
	// Enum: [Enabled Disabled]
	Inbound string `json:"inbound,omitempty"`

	// Status for the Outbound Direction
	// Enum: [Enabled Disabled]
	Outbound string `json:"outbound,omitempty"`
}

// Validate validates this setting direction
func (m *SettingDirection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInbound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutbound(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var settingDirectionTypeInboundPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Enabled","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settingDirectionTypeInboundPropEnum = append(settingDirectionTypeInboundPropEnum, v)
	}
}

const (

	// SettingDirectionInboundEnabled captures enum value "Enabled"
	SettingDirectionInboundEnabled string = "Enabled"

	// SettingDirectionInboundDisabled captures enum value "Disabled"
	SettingDirectionInboundDisabled string = "Disabled"
)

// prop value enum
func (m *SettingDirection) validateInboundEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, settingDirectionTypeInboundPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SettingDirection) validateInbound(formats strfmt.Registry) error {
	if swag.IsZero(m.Inbound) { // not required
		return nil
	}

	// value enum
	if err := m.validateInboundEnum("inbound", "body", m.Inbound); err != nil {
		return err
	}

	return nil
}

var settingDirectionTypeOutboundPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Enabled","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settingDirectionTypeOutboundPropEnum = append(settingDirectionTypeOutboundPropEnum, v)
	}
}

const (

	// SettingDirectionOutboundEnabled captures enum value "Enabled"
	SettingDirectionOutboundEnabled string = "Enabled"

	// SettingDirectionOutboundDisabled captures enum value "Disabled"
	SettingDirectionOutboundDisabled string = "Disabled"
)

// prop value enum
func (m *SettingDirection) validateOutboundEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, settingDirectionTypeOutboundPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SettingDirection) validateOutbound(formats strfmt.Registry) error {
	if swag.IsZero(m.Outbound) { // not required
		return nil
	}

	// value enum
	if err := m.validateOutboundEnum("outbound", "body", m.Outbound); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this setting direction based on context it is used
func (m *SettingDirection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingDirection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingDirection) UnmarshalBinary(b []byte) error {
	var res SettingDirection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
