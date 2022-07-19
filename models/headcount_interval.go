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

// HeadcountInterval headcount interval
//
// swagger:model HeadcountInterval
type HeadcountInterval struct {

	// The start date-time for this headcount interval in ISO-8601 format, must be within the 8 day schedule
	// Required: true
	// Format: date-time
	Interval *strfmt.DateTime `json:"interval"`

	// Headcount value for this interval
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this headcount interval
func (m *HeadcountInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeadcountInterval) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	if err := validate.FormatOf("interval", "body", "date-time", m.Interval.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HeadcountInterval) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HeadcountInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeadcountInterval) UnmarshalBinary(b []byte) error {
	var res HeadcountInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
