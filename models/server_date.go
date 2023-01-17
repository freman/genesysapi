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

// ServerDate server date
//
// swagger:model ServerDate
type ServerDate struct {

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CurrentDate strfmt.DateTime `json:"currentDate,omitempty"`
}

// Validate validates this server date
func (m *ServerDate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerDate) validateCurrentDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("currentDate", "body", "date-time", m.CurrentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this server date based on context it is used
func (m *ServerDate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerDate) UnmarshalBinary(b []byte) error {
	var res ServerDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
