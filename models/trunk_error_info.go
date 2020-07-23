// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrunkErrorInfo trunk error info
//
// swagger:model TrunkErrorInfo
type TrunkErrorInfo struct {

	// code
	Code string `json:"code,omitempty"`

	// details
	Details *TrunkErrorInfoDetails `json:"details,omitempty"`

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this trunk error info
func (m *TrunkErrorInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrunkErrorInfo) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrunkErrorInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrunkErrorInfo) UnmarshalBinary(b []byte) error {
	var res TrunkErrorInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
