// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuGenerateScheduleRequest bu generate schedule request
//
// swagger:model BuGenerateScheduleRequest
type BuGenerateScheduleRequest struct {

	// The description for the schedule
	// Required: true
	Description *string `json:"description"`

	// The forecast to use when generating the schedule.  Note that the forecast must fully encompass the schedule's start week + week count
	// Required: true
	ShortTermForecast *BuShortTermForecastReference `json:"shortTermForecast"`

	// The number of weeks in the schedule. One extra day is added at the end
	// Required: true
	WeekCount *int32 `json:"weekCount"`
}

// Validate validates this bu generate schedule request
func (m *BuGenerateScheduleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortTermForecast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuGenerateScheduleRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *BuGenerateScheduleRequest) validateShortTermForecast(formats strfmt.Registry) error {

	if err := validate.Required("shortTermForecast", "body", m.ShortTermForecast); err != nil {
		return err
	}

	if m.ShortTermForecast != nil {
		if err := m.ShortTermForecast.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortTermForecast")
			}
			return err
		}
	}

	return nil
}

func (m *BuGenerateScheduleRequest) validateWeekCount(formats strfmt.Registry) error {

	if err := validate.Required("weekCount", "body", m.WeekCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuGenerateScheduleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuGenerateScheduleRequest) UnmarshalBinary(b []byte) error {
	var res BuGenerateScheduleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
