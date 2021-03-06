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

// BuAgentScheduleShift bu agent schedule shift
//
// swagger:model BuAgentScheduleShift
type BuAgentScheduleShift struct {

	// The activities associated with this shift
	Activities []*BuAgentScheduleActivity `json:"activities"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The length of this shift in minutes
	LengthMinutes int32 `json:"lengthMinutes,omitempty"`

	// Whether this shift was manually edited. This is only set by clients and is used for rescheduling
	ManuallyEdited bool `json:"manuallyEdited"`

	// The schedule to which this shift belongs
	// Read Only: true
	Schedule *BuScheduleReference `json:"schedule,omitempty"`

	// The start date of this shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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
