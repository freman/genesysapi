// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CopyBuForecastRequest copy bu forecast request
//
// swagger:model CopyBuForecastRequest
type CopyBuForecastRequest struct {

	// The description for the forecast
	// Required: true
	Description *string `json:"description"`

	// The start date of the new forecast to create from the existing forecast. Must correspond to the start day of week for the business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Required: true
	// Format: date
	WeekDate *strfmt.Date `json:"weekDate"`
}

// Validate validates this copy bu forecast request
func (m *CopyBuForecastRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopyBuForecastRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CopyBuForecastRequest) validateWeekDate(formats strfmt.Registry) error {

	if err := validate.Required("weekDate", "body", m.WeekDate); err != nil {
		return err
	}

	if err := validate.FormatOf("weekDate", "body", "date", m.WeekDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this copy bu forecast request based on context it is used
func (m *CopyBuForecastRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CopyBuForecastRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyBuForecastRequest) UnmarshalBinary(b []byte) error {
	var res CopyBuForecastRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
