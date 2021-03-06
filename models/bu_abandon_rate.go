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

// BuAbandonRate Service goal abandon rate configuration
//
// swagger:model BuAbandonRate
type BuAbandonRate struct {

	// Whether to include abandon rate in the associated configuration
	// Required: true
	Include *bool `json:"include"`

	// Abandon rate percent goal. Required if include == true
	Percent int32 `json:"percent,omitempty"`
}

// Validate validates this bu abandon rate
func (m *BuAbandonRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInclude(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAbandonRate) validateInclude(formats strfmt.Registry) error {

	if err := validate.Required("include", "body", m.Include); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuAbandonRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuAbandonRate) UnmarshalBinary(b []byte) error {
	var res BuAbandonRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
