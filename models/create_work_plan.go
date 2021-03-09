// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateWorkPlan Work plan information
//
// swagger:model CreateWorkPlan
type CreateWorkPlan struct {

	// Agents in this work plan
	Agents []*UserReference `json:"agents"`

	// Whether to constrain the maximum consecutive working days
	ConstrainMaximumConsecutiveWorkingDays bool `json:"constrainMaximumConsecutiveWorkingDays"`

	// Whether to constrain the maximum consecutive working weekends
	ConstrainMaximumConsecutiveWorkingWeekends bool `json:"constrainMaximumConsecutiveWorkingWeekends"`

	// Whether the minimum time between shifts constraint is enabled for this work plan
	ConstrainMinimumTimeBetweenShifts bool `json:"constrainMinimumTimeBetweenShifts"`

	// Whether paid time granularity should be constrained for this workplan
	ConstrainPaidTimeGranularity bool `json:"constrainPaidTimeGranularity"`

	// Whether the weekly paid time constraint is enabled for this work plan
	ConstrainWeeklyPaidTime bool `json:"constrainWeeklyPaidTime"`

	// Whether the work plan is enabled for scheduling
	Enabled bool `json:"enabled"`

	// Whether the weekly paid time constraint is flexible for this work plan
	FlexibleWeeklyPaidTime bool `json:"flexibleWeeklyPaidTime"`

	// The maximum number of consecutive days that agents assigned to this work plan are allowed to work. Used if constrainMaximumConsecutiveWorkingDays == true
	MaximumConsecutiveWorkingDays int32 `json:"maximumConsecutiveWorkingDays,omitempty"`

	// The maximum number of consecutive weekends that agents who are assigned to this work plan are allowed to work
	MaximumConsecutiveWorkingWeekends int32 `json:"maximumConsecutiveWorkingWeekends,omitempty"`

	// Maximum number days in a week allowed to be scheduled for this work plan
	MaximumDays int32 `json:"maximumDays,omitempty"`

	// Maximum days off in the planning period
	MaximumDaysOffPerPlanningPeriod int32 `json:"maximumDaysOffPerPlanningPeriod,omitempty"`

	// Maximum paid minutes in the planning period
	MaximumPaidMinutesPerPlanningPeriod int32 `json:"maximumPaidMinutesPerPlanningPeriod,omitempty"`

	// Minimum amount of consecutive non working minutes per week that agents who are assigned this work plan are allowed to have off
	MinimumConsecutiveNonWorkingMinutesPerWeek int32 `json:"minimumConsecutiveNonWorkingMinutesPerWeek,omitempty"`

	// Minimum days off in the planning period
	MinimumDaysOffPerPlanningPeriod int32 `json:"minimumDaysOffPerPlanningPeriod,omitempty"`

	// Minimum paid minutes in the planning period
	MinimumPaidMinutesPerPlanningPeriod int32 `json:"minimumPaidMinutesPerPlanningPeriod,omitempty"`

	// The time period in minutes for the duration between the start times of two consecutive working days
	MinimumShiftStartDistanceMinutes int32 `json:"minimumShiftStartDistanceMinutes,omitempty"`

	// Minimum time between shifts in minutes defined in this work plan. Used if constrainMinimumTimeBetweenShifts == true
	MinimumTimeBetweenShiftsMinutes int32 `json:"minimumTimeBetweenShiftsMinutes,omitempty"`

	// The minimum number of days that agents assigned to a work plan must work per week
	MinimumWorkingDaysPerWeek int32 `json:"minimumWorkingDaysPerWeek,omitempty"`

	// Name of this work plan
	// Required: true
	Name *string `json:"name"`

	// Optional days to schedule for this work plan
	OptionalDays *SetWrapperDayOfWeek `json:"optionalDays,omitempty"`

	// Granularity in minutes allowed for shift paid time in this work plan. Used if constrainPaidTimeGranularity == true
	PaidTimeGranularityMinutes int32 `json:"paidTimeGranularityMinutes,omitempty"`

	// This constraint ensures that an agent starts each workday within a user-defined time threshold
	// Enum: [ShiftStart ShiftStartAndPaidDuration]
	ShiftStartVarianceType string `json:"shiftStartVarianceType,omitempty"`

	// Variance in minutes among start times of shifts in this work plan
	ShiftStartVariances *ListWrapperShiftStartVariance `json:"shiftStartVariances,omitempty"`

	// Shifts in this work plan
	Shifts []*CreateWorkPlanShift `json:"shifts"`

	// Exact weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == false
	WeeklyExactPaidMinutes int32 `json:"weeklyExactPaidMinutes,omitempty"`

	// Maximum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
	WeeklyMaximumPaidMinutes int32 `json:"weeklyMaximumPaidMinutes,omitempty"`

	// Minimum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
	WeeklyMinimumPaidMinutes int32 `json:"weeklyMinimumPaidMinutes,omitempty"`
}

// Validate validates this create work plan
func (m *CreateWorkPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionalDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShiftStartVarianceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShiftStartVariances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShifts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateWorkPlan) validateAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.Agents) { // not required
		return nil
	}

	for i := 0; i < len(m.Agents); i++ {
		if swag.IsZero(m.Agents[i]) { // not required
			continue
		}

		if m.Agents[i] != nil {
			if err := m.Agents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateWorkPlan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateWorkPlan) validateOptionalDays(formats strfmt.Registry) error {

	if swag.IsZero(m.OptionalDays) { // not required
		return nil
	}

	if m.OptionalDays != nil {
		if err := m.OptionalDays.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("optionalDays")
			}
			return err
		}
	}

	return nil
}

var createWorkPlanTypeShiftStartVarianceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ShiftStart","ShiftStartAndPaidDuration"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createWorkPlanTypeShiftStartVarianceTypePropEnum = append(createWorkPlanTypeShiftStartVarianceTypePropEnum, v)
	}
}

const (

	// CreateWorkPlanShiftStartVarianceTypeShiftStart captures enum value "ShiftStart"
	CreateWorkPlanShiftStartVarianceTypeShiftStart string = "ShiftStart"

	// CreateWorkPlanShiftStartVarianceTypeShiftStartAndPaidDuration captures enum value "ShiftStartAndPaidDuration"
	CreateWorkPlanShiftStartVarianceTypeShiftStartAndPaidDuration string = "ShiftStartAndPaidDuration"
)

// prop value enum
func (m *CreateWorkPlan) validateShiftStartVarianceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createWorkPlanTypeShiftStartVarianceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateWorkPlan) validateShiftStartVarianceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ShiftStartVarianceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateShiftStartVarianceTypeEnum("shiftStartVarianceType", "body", m.ShiftStartVarianceType); err != nil {
		return err
	}

	return nil
}

func (m *CreateWorkPlan) validateShiftStartVariances(formats strfmt.Registry) error {

	if swag.IsZero(m.ShiftStartVariances) { // not required
		return nil
	}

	if m.ShiftStartVariances != nil {
		if err := m.ShiftStartVariances.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shiftStartVariances")
			}
			return err
		}
	}

	return nil
}

func (m *CreateWorkPlan) validateShifts(formats strfmt.Registry) error {

	if swag.IsZero(m.Shifts) { // not required
		return nil
	}

	for i := 0; i < len(m.Shifts); i++ {
		if swag.IsZero(m.Shifts[i]) { // not required
			continue
		}

		if m.Shifts[i] != nil {
			if err := m.Shifts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shifts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateWorkPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateWorkPlan) UnmarshalBinary(b []byte) error {
	var res CreateWorkPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
