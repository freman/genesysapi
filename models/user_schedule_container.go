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

// UserScheduleContainer Container object to hold a map of user schedules
//
// swagger:model UserScheduleContainer
type UserScheduleContainer struct {

	// The reference time zone used for the management unit
	ManagementUnitTimeZone string `json:"managementUnitTimeZone,omitempty"`

	// References to all published week schedules overlapping the start/end date query parameters
	PublishedSchedules []*WeekScheduleReference `json:"publishedSchedules"`

	// Map of user id to user schedule
	UserSchedules map[string]UserSchedule `json:"userSchedules,omitempty"`
}

// Validate validates this user schedule container
func (m *UserScheduleContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublishedSchedules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserScheduleContainer) validatePublishedSchedules(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishedSchedules) { // not required
		return nil
	}

	for i := 0; i < len(m.PublishedSchedules); i++ {
		if swag.IsZero(m.PublishedSchedules[i]) { // not required
			continue
		}

		if m.PublishedSchedules[i] != nil {
			if err := m.PublishedSchedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("publishedSchedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserScheduleContainer) validateUserSchedules(formats strfmt.Registry) error {

	if swag.IsZero(m.UserSchedules) { // not required
		return nil
	}

	for k := range m.UserSchedules {

		if err := validate.Required("userSchedules"+"."+k, "body", m.UserSchedules[k]); err != nil {
			return err
		}
		if val, ok := m.UserSchedules[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserScheduleContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserScheduleContainer) UnmarshalBinary(b []byte) error {
	var res UserScheduleContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
