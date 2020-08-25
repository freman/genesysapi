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

// CallableTime callable time
//
// swagger:model CallableTime
type CallableTime struct {

	// The time intervals for which it is acceptable to place outbound calls.
	// Required: true
	TimeSlots []*CampaignTimeSlot `json:"timeSlots"`

	// The time zone for the time slots; for example, Africa/Abidjan
	// Required: true
	TimeZoneID *string `json:"timeZoneId"`
}

// Validate validates this callable time
func (m *CallableTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeSlots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeZoneID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallableTime) validateTimeSlots(formats strfmt.Registry) error {

	if err := validate.Required("timeSlots", "body", m.TimeSlots); err != nil {
		return err
	}

	for i := 0; i < len(m.TimeSlots); i++ {
		if swag.IsZero(m.TimeSlots[i]) { // not required
			continue
		}

		if m.TimeSlots[i] != nil {
			if err := m.TimeSlots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("timeSlots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CallableTime) validateTimeZoneID(formats strfmt.Registry) error {

	if err := validate.Required("timeZoneId", "body", m.TimeZoneID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallableTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallableTime) UnmarshalBinary(b []byte) error {
	var res CallableTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}