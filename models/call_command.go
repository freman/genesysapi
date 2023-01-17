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

// CallCommand call command
//
// swagger:model CallCommand
type CallCommand struct {

	// The phone number to dial for this call.
	// Required: true
	CallNumber *string `json:"callNumber"`

	// For a dialer preview or scheduled callback, the phone column associated with the phone number
	PhoneColumn string `json:"phoneColumn,omitempty"`
}

// Validate validates this call command
func (m *CallCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallCommand) validateCallNumber(formats strfmt.Registry) error {

	if err := validate.Required("callNumber", "body", m.CallNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this call command based on context it is used
func (m *CallCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CallCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallCommand) UnmarshalBinary(b []byte) error {
	var res CallCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
