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

// LauncherButtonSettings The settings for the launcher button
//
// swagger:model LauncherButtonSettings
type LauncherButtonSettings struct {

	// The visibility settings for the button
	// Enum: [On Off OnDemand]
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this launcher button settings
func (m *LauncherButtonSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var launcherButtonSettingsTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["On","Off","OnDemand"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		launcherButtonSettingsTypeVisibilityPropEnum = append(launcherButtonSettingsTypeVisibilityPropEnum, v)
	}
}

const (

	// LauncherButtonSettingsVisibilityOn captures enum value "On"
	LauncherButtonSettingsVisibilityOn string = "On"

	// LauncherButtonSettingsVisibilityOff captures enum value "Off"
	LauncherButtonSettingsVisibilityOff string = "Off"

	// LauncherButtonSettingsVisibilityOnDemand captures enum value "OnDemand"
	LauncherButtonSettingsVisibilityOnDemand string = "OnDemand"
)

// prop value enum
func (m *LauncherButtonSettings) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, launcherButtonSettingsTypeVisibilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LauncherButtonSettings) validateVisibility(formats strfmt.Registry) error {
	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this launcher button settings based on context it is used
func (m *LauncherButtonSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LauncherButtonSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LauncherButtonSettings) UnmarshalBinary(b []byte) error {
	var res LauncherButtonSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
