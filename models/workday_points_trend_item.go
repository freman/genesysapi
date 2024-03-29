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

// WorkdayPointsTrendItem workday points trend item
//
// swagger:model WorkdayPointsTrendItem
type WorkdayPointsTrendItem struct {

	// workday date for the points trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Read Only: true
	// Format: date
	DateWorkday strfmt.Date `json:"dateWorkday,omitempty"`

	// workday points for the date
	// Read Only: true
	Points float64 `json:"points,omitempty"`
}

// Validate validates this workday points trend item
func (m *WorkdayPointsTrendItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateWorkday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkdayPointsTrendItem) validateDateWorkday(formats strfmt.Registry) error {
	if swag.IsZero(m.DateWorkday) { // not required
		return nil
	}

	if err := validate.FormatOf("dateWorkday", "body", "date", m.DateWorkday.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this workday points trend item based on the context it is used
func (m *WorkdayPointsTrendItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateWorkday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkdayPointsTrendItem) contextValidateDateWorkday(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateWorkday", "body", strfmt.Date(m.DateWorkday)); err != nil {
		return err
	}

	return nil
}

func (m *WorkdayPointsTrendItem) contextValidatePoints(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "points", "body", float64(m.Points)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkdayPointsTrendItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkdayPointsTrendItem) UnmarshalBinary(b []byte) error {
	var res WorkdayPointsTrendItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
