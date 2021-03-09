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

// SchedulingSettingsResponse Scheduling Settings
//
// swagger:model SchedulingSettingsResponse
type SchedulingSettingsResponse struct {

	// Default shrinkage percent for scheduling
	DefaultShrinkagePercent float64 `json:"defaultShrinkagePercent,omitempty"`

	// Max occupancy percent for deferred work
	MaxOccupancyPercentForDeferredWork int32 `json:"maxOccupancyPercentForDeferredWork,omitempty"`

	// Planning period settings for scheduling
	PlanningPeriod *PlanningPeriodSettings `json:"planningPeriod,omitempty"`

	// Shrinkage overrides for scheduling
	ShrinkageOverrides *ShrinkageOverrides `json:"shrinkageOverrides,omitempty"`

	// Start day of weekend for scheduling
	// Enum: [Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
	StartDayOfWeekend string `json:"startDayOfWeekend,omitempty"`
}

// Validate validates this scheduling settings response
func (m *SchedulingSettingsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlanningPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShrinkageOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDayOfWeekend(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulingSettingsResponse) validatePlanningPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.PlanningPeriod) { // not required
		return nil
	}

	if m.PlanningPeriod != nil {
		if err := m.PlanningPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("planningPeriod")
			}
			return err
		}
	}

	return nil
}

func (m *SchedulingSettingsResponse) validateShrinkageOverrides(formats strfmt.Registry) error {

	if swag.IsZero(m.ShrinkageOverrides) { // not required
		return nil
	}

	if m.ShrinkageOverrides != nil {
		if err := m.ShrinkageOverrides.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shrinkageOverrides")
			}
			return err
		}
	}

	return nil
}

var schedulingSettingsResponseTypeStartDayOfWeekendPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schedulingSettingsResponseTypeStartDayOfWeekendPropEnum = append(schedulingSettingsResponseTypeStartDayOfWeekendPropEnum, v)
	}
}

const (

	// SchedulingSettingsResponseStartDayOfWeekendSunday captures enum value "Sunday"
	SchedulingSettingsResponseStartDayOfWeekendSunday string = "Sunday"

	// SchedulingSettingsResponseStartDayOfWeekendMonday captures enum value "Monday"
	SchedulingSettingsResponseStartDayOfWeekendMonday string = "Monday"

	// SchedulingSettingsResponseStartDayOfWeekendTuesday captures enum value "Tuesday"
	SchedulingSettingsResponseStartDayOfWeekendTuesday string = "Tuesday"

	// SchedulingSettingsResponseStartDayOfWeekendWednesday captures enum value "Wednesday"
	SchedulingSettingsResponseStartDayOfWeekendWednesday string = "Wednesday"

	// SchedulingSettingsResponseStartDayOfWeekendThursday captures enum value "Thursday"
	SchedulingSettingsResponseStartDayOfWeekendThursday string = "Thursday"

	// SchedulingSettingsResponseStartDayOfWeekendFriday captures enum value "Friday"
	SchedulingSettingsResponseStartDayOfWeekendFriday string = "Friday"

	// SchedulingSettingsResponseStartDayOfWeekendSaturday captures enum value "Saturday"
	SchedulingSettingsResponseStartDayOfWeekendSaturday string = "Saturday"
)

// prop value enum
func (m *SchedulingSettingsResponse) validateStartDayOfWeekendEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, schedulingSettingsResponseTypeStartDayOfWeekendPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SchedulingSettingsResponse) validateStartDayOfWeekend(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDayOfWeekend) { // not required
		return nil
	}

	// value enum
	if err := m.validateStartDayOfWeekendEnum("startDayOfWeekend", "body", m.StartDayOfWeekend); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulingSettingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulingSettingsResponse) UnmarshalBinary(b []byte) error {
	var res SchedulingSettingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
