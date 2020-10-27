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

// CallRecord call record
//
// swagger:model CallRecord
type CallRecord struct {

	// Timestamp of the last attempt to reach this number. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	LastAttempt strfmt.DateTime `json:"lastAttempt,omitempty"`

	// Result of the last attempt to reach this number
	// Read Only: true
	LastResult string `json:"lastResult,omitempty"`
}

// Validate validates this call record
func (m *CallRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastAttempt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallRecord) validateLastAttempt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastAttempt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastAttempt", "body", "date-time", m.LastAttempt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallRecord) UnmarshalBinary(b []byte) error {
	var res CallRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
