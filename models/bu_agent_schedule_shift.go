// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuAgentScheduleShift bu agent schedule shift
//
// swagger:model BuAgentScheduleShift
type BuAgentScheduleShift struct {

	// The activities associated with this shift
	Activities []*BuAgentScheduleActivity `json:"activities"`

	// The ID of the shift
	ID string `json:"id,omitempty"`

	// The length of this shift in minutes
	// Read Only: true
	LengthMinutes int32 `json:"lengthMinutes,omitempty"`

	// Whether this shift was manually edited. This is only set by clients and is used for rescheduling
	ManuallyEdited bool `json:"manuallyEdited"`

	// The schedule to which this shift belongs
	// Read Only: true
	Schedule *BuScheduleReference `json:"schedule,omitempty"`

	// The start date of this shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`
}

// Validate validates this bu agent schedule shift
func (m *BuAgentScheduleShift) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAgentScheduleShift) validateActivities(formats strfmt.Registry) error {
	if swag.IsZero(m.Activities) { // not required
		return nil
	}

	for i := 0; i < len(m.Activities); i++ {
		if swag.IsZero(m.Activities[i]) { // not required
			continue
		}

		if m.Activities[i] != nil {
			if err := m.Activities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuAgentScheduleShift) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *BuAgentScheduleShift) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bu agent schedule shift based on the context it is used
func (m *BuAgentScheduleShift) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActivities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLengthMinutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAgentScheduleShift) contextValidateActivities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Activities); i++ {

		if m.Activities[i] != nil {
			if err := m.Activities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuAgentScheduleShift) contextValidateLengthMinutes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lengthMinutes", "body", int32(m.LengthMinutes)); err != nil {
		return err
	}

	return nil
}

func (m *BuAgentScheduleShift) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *BuAgentScheduleShift) contextValidateStartDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "startDate", "body", strfmt.DateTime(m.StartDate)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuAgentScheduleShift) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuAgentScheduleShift) UnmarshalBinary(b []byte) error {
	var res BuAgentScheduleShift
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
