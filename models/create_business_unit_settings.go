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

// CreateBusinessUnitSettings create business unit settings
//
// swagger:model CreateBusinessUnitSettings
type CreateBusinessUnitSettings struct {

	// Short term forecasting settings
	ShortTermForecasting *BuShortTermForecastingSettings `json:"shortTermForecasting,omitempty"`

	// The start day of week for this business unit
	// Required: true
	// Enum: [Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
	StartDayOfWeek *string `json:"startDayOfWeek"`

	// The time zone for this business unit, using the Olsen tz database format
	// Required: true
	TimeZone *string `json:"timeZone"`
}

// Validate validates this create business unit settings
func (m *CreateBusinessUnitSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShortTermForecasting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDayOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBusinessUnitSettings) validateShortTermForecasting(formats strfmt.Registry) error {

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

var createBusinessUnitSettingsTypeStartDayOfWeekPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createBusinessUnitSettingsTypeStartDayOfWeekPropEnum = append(createBusinessUnitSettingsTypeStartDayOfWeekPropEnum, v)
	}
}

const (

	// CreateBusinessUnitSettingsStartDayOfWeekSunday captures enum value "Sunday"
	CreateBusinessUnitSettingsStartDayOfWeekSunday string = "Sunday"

	// CreateBusinessUnitSettingsStartDayOfWeekMonday captures enum value "Monday"
	CreateBusinessUnitSettingsStartDayOfWeekMonday string = "Monday"

	// CreateBusinessUnitSettingsStartDayOfWeekTuesday captures enum value "Tuesday"
	CreateBusinessUnitSettingsStartDayOfWeekTuesday string = "Tuesday"

	// CreateBusinessUnitSettingsStartDayOfWeekWednesday captures enum value "Wednesday"
	CreateBusinessUnitSettingsStartDayOfWeekWednesday string = "Wednesday"

	// CreateBusinessUnitSettingsStartDayOfWeekThursday captures enum value "Thursday"
	CreateBusinessUnitSettingsStartDayOfWeekThursday string = "Thursday"

	// CreateBusinessUnitSettingsStartDayOfWeekFriday captures enum value "Friday"
	CreateBusinessUnitSettingsStartDayOfWeekFriday string = "Friday"

	// CreateBusinessUnitSettingsStartDayOfWeekSaturday captures enum value "Saturday"
	CreateBusinessUnitSettingsStartDayOfWeekSaturday string = "Saturday"
)

// prop value enum
func (m *CreateBusinessUnitSettings) validateStartDayOfWeekEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createBusinessUnitSettingsTypeStartDayOfWeekPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateBusinessUnitSettings) validateStartDayOfWeek(formats strfmt.Registry) error {

	if err := validate.Required("startDayOfWeek", "body", m.StartDayOfWeek); err != nil {
		return err
	}

	// value enum
	if err := m.validateStartDayOfWeekEnum("startDayOfWeek", "body", *m.StartDayOfWeek); err != nil {
		return err
	}

	return nil
}

func (m *CreateBusinessUnitSettings) validateTimeZone(formats strfmt.Registry) error {

	if err := validate.Required("timeZone", "body", m.TimeZone); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateBusinessUnitSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBusinessUnitSettings) UnmarshalBinary(b []byte) error {
	var res CreateBusinessUnitSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
