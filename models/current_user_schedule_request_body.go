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

// CurrentUserScheduleRequestBody POST request body for fetching the current user's schedule over a given range
//
// swagger:model CurrentUserScheduleRequestBody
type CurrentUserScheduleRequestBody struct {

	// End of the range of schedules to fetch, in ISO-8601 format
	// Required: true
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate"`

	// Whether to load the full week's schedule (for the current user) of any week overlapping the start/end date query parameters, defaults to false
	LoadFullWeeks bool `json:"loadFullWeeks"`

	// Beginning of the range of schedules to fetch, in ISO-8601 format
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate"`
}

// Validate validates this current user schedule request body
func (m *CurrentUserScheduleRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
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

func (m *CurrentUserScheduleRequestBody) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("endDate", "body", m.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CurrentUserScheduleRequestBody) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", m.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrentUserScheduleRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrentUserScheduleRequestBody) UnmarshalBinary(b []byte) error {
	var res CurrentUserScheduleRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}