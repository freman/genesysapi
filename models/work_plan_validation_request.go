// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkPlanValidationRequest Work plan information
//
// swagger:model WorkPlanValidationRequest
type WorkPlanValidationRequest struct {

	// Agents in this work plan
	Agents []*DeletableUserReference `json:"agents"`

	// Whether the minimum time between shifts constraint is enabled for this work plan
	ConstrainMinimumTimeBetweenShifts bool `json:"constrainMinimumTimeBetweenShifts"`

	// Whether paid time granularity is constrained for this workplan
	ConstrainPaidTimeGranularity bool `json:"constrainPaidTimeGranularity"`

	// Whether the weekly paid time constraint is enabled for this work plan
	ConstrainWeeklyPaidTime bool `json:"constrainWeeklyPaidTime"`

	// Whether the work plan is enabled for scheduling
	Enabled bool `json:"enabled"`

	// Whether the weekly paid time constraint is flexible for this work plan
	FlexibleWeeklyPaidTime bool `json:"flexibleWeeklyPaidTime"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Maximum number days in a week allowed to be scheduled for this work plan
	MaximumDays int32 `json:"maximumDays,omitempty"`

	// Minimum time between shifts in minutes defined in this work plan. Used if constrainMinimumTimeBetweenShifts == true
	MinimumTimeBetweenShiftsMinutes int32 `json:"minimumTimeBetweenShiftsMinutes,omitempty"`

	// The minimum number of days that agents assigned to a work plan must work per week
	MinimumWorkingDaysPerWeek int32 `json:"minimumWorkingDaysPerWeek,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Optional days to schedule for this work plan
	OptionalDays *SetWrapperDayOfWeek `json:"optionalDays,omitempty"`

	// Granularity in minutes allowed for shift paid time in this work plan. Used if constrainPaidTimeGranularity == true
	PaidTimeGranularityMinutes int32 `json:"paidTimeGranularityMinutes,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Variance in minutes among start times of shifts in this work plan
	ShiftStartVariances *ListWrapperShiftStartVariance `json:"shiftStartVariances,omitempty"`

	// Shifts in this work plan
	Shifts []*WorkPlanShift `json:"shifts"`

	// Exact weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == false
	WeeklyExactPaidMinutes int32 `json:"weeklyExactPaidMinutes,omitempty"`

	// Maximum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
	WeeklyMaximumPaidMinutes int32 `json:"weeklyMaximumPaidMinutes,omitempty"`

	// Minimum weekly paid time in minutes for this work plan. Used if flexibleWeeklyPaidTime == true
	WeeklyMinimumPaidMinutes int32 `json:"weeklyMinimumPaidMinutes,omitempty"`
}

// Validate validates this work plan validation request
func (m *WorkPlanValidationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionalDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *WorkPlanValidationRequest) validateAgents(formats strfmt.Registry) error {

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

func (m *WorkPlanValidationRequest) validateOptionalDays(formats strfmt.Registry) error {

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

func (m *WorkPlanValidationRequest) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkPlanValidationRequest) validateShiftStartVariances(formats strfmt.Registry) error {

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

func (m *WorkPlanValidationRequest) validateShifts(formats strfmt.Registry) error {

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
func (m *WorkPlanValidationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkPlanValidationRequest) UnmarshalBinary(b []byte) error {
	var res WorkPlanValidationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
