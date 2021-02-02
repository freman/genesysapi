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

// DateRangeWithOptionalEnd date range with optional end
//
// swagger:model DateRangeWithOptionalEnd
type DateRangeWithOptionalEnd struct {

	// The end date for work plan rotation or an agent, interpreted in the business unit's time zone. Null denotes open ended date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	EndBusinessUnitDate strfmt.Date `json:"endBusinessUnitDate,omitempty"`

	// The start date for work plan rotation or an agent, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Required: true
	// Format: date
	StartBusinessUnitDate *strfmt.Date `json:"startBusinessUnitDate"`
}

// Validate validates this date range with optional end
func (m *DateRangeWithOptionalEnd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndBusinessUnitDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartBusinessUnitDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DateRangeWithOptionalEnd) validateEndBusinessUnitDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndBusinessUnitDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endBusinessUnitDate", "body", "date", m.EndBusinessUnitDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DateRangeWithOptionalEnd) validateStartBusinessUnitDate(formats strfmt.Registry) error {

	if err := validate.Required("startBusinessUnitDate", "body", m.StartBusinessUnitDate); err != nil {
		return err
	}

	if err := validate.FormatOf("startBusinessUnitDate", "body", "date", m.StartBusinessUnitDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DateRangeWithOptionalEnd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DateRangeWithOptionalEnd) UnmarshalBinary(b []byte) error {
	var res DateRangeWithOptionalEnd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}