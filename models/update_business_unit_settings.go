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

// UpdateBusinessUnitSettings update business unit settings
//
// swagger:model UpdateBusinessUnitSettings
type UpdateBusinessUnitSettings struct {

	// Version metadata for this business unit
	// Required: true
	Metadata *WfmVersionedEntityMetadata `json:"metadata"`

	// Scheduling settings
	Scheduling *BuSchedulingSettings `json:"scheduling,omitempty"`

	// Short term forecasting settings
	ShortTermForecasting *BuShortTermForecastingSettings `json:"shortTermForecasting,omitempty"`

	// The start day of week for this business unit
	// Read Only: true
	// Enum: [Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
	StartDayOfWeek string `json:"startDayOfWeek,omitempty"`

	// The time zone for this business unit, using the Olsen tz database format
	// Read Only: true
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this update business unit settings
func (m *UpdateBusinessUnitSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortTermForecasting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDayOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateBusinessUnitSettings) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateBusinessUnitSettings) validateScheduling(formats strfmt.Registry) error {

	if swag.IsZero(m.Scheduling) { // not required
		return nil
	}

	if m.Scheduling != nil {
		if err := m.Scheduling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduling")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateBusinessUnitSettings) validateShortTermForecasting(formats strfmt.Registry) error {

	if swag.IsZero(m.ShortTermForecasting) { // not required
		return nil
	}

	if m.ShortTermForecasting != nil {
		if err := m.ShortTermForecasting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortTermForecasting")
			}
			return err
		}
	}

	return nil
}

var updateBusinessUnitSettingsTypeStartDayOfWeekPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateBusinessUnitSettingsTypeStartDayOfWeekPropEnum = append(updateBusinessUnitSettingsTypeStartDayOfWeekPropEnum, v)
	}
}

const (

	// UpdateBusinessUnitSettingsStartDayOfWeekSunday captures enum value "Sunday"
	UpdateBusinessUnitSettingsStartDayOfWeekSunday string = "Sunday"

	// UpdateBusinessUnitSettingsStartDayOfWeekMonday captures enum value "Monday"
	UpdateBusinessUnitSettingsStartDayOfWeekMonday string = "Monday"

	// UpdateBusinessUnitSettingsStartDayOfWeekTuesday captures enum value "Tuesday"
	UpdateBusinessUnitSettingsStartDayOfWeekTuesday string = "Tuesday"

	// UpdateBusinessUnitSettingsStartDayOfWeekWednesday captures enum value "Wednesday"
	UpdateBusinessUnitSettingsStartDayOfWeekWednesday string = "Wednesday"

	// UpdateBusinessUnitSettingsStartDayOfWeekThursday captures enum value "Thursday"
	UpdateBusinessUnitSettingsStartDayOfWeekThursday string = "Thursday"

	// UpdateBusinessUnitSettingsStartDayOfWeekFriday captures enum value "Friday"
	UpdateBusinessUnitSettingsStartDayOfWeekFriday string = "Friday"

	// UpdateBusinessUnitSettingsStartDayOfWeekSaturday captures enum value "Saturday"
	UpdateBusinessUnitSettingsStartDayOfWeekSaturday string = "Saturday"
)

// prop value enum
func (m *UpdateBusinessUnitSettings) validateStartDayOfWeekEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateBusinessUnitSettingsTypeStartDayOfWeekPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateBusinessUnitSettings) validateStartDayOfWeek(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDayOfWeek) { // not required
		return nil
	}

	// value enum
	if err := m.validateStartDayOfWeekEnum("startDayOfWeek", "body", m.StartDayOfWeek); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateBusinessUnitSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateBusinessUnitSettings) UnmarshalBinary(b []byte) error {
	var res UpdateBusinessUnitSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
