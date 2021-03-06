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

// ActionContractInput Contract definition.
//
// swagger:model ActionContractInput
type ActionContractInput struct {

	// Execution input contract
	// Required: true
	Input *PostInputContract `json:"input"`

	// Execution output contract
	// Required: true
	Output *PostOutputContract `json:"output"`
}

// Validate validates this action contract input
func (m *ActionContractInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionContractInput) validateInput(formats strfmt.Registry) error {

	if err := validate.Required("input", "body", m.Input); err != nil {
		return err
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

func (m *ActionContractInput) validateOutput(formats strfmt.Registry) error {

	if err := validate.Required("output", "body", m.Output); err != nil {
		return err
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionContractInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionContractInput) UnmarshalBinary(b []byte) error {
	var res ActionContractInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
