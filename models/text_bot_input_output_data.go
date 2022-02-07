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

// TextBotInputOutputData Input/Output data related to a bot flow which is exiting gracefully.
//
// swagger:model TextBotInputOutputData
type TextBotInputOutputData struct {

	// The input/output variables using the format as appropriate for the variable data type in the flow definition.
	// Required: true
	Variables map[string]interface{} `json:"variables"`
}

// Validate validates this text bot input output data
func (m *TextBotInputOutputData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextBotInputOutputData) validateVariables(formats strfmt.Registry) error {

	for k := range m.Variables {

		if err := validate.Required("variables"+"."+k, "body", m.Variables[k]); err != nil {
			return err
		}

		if err := validate.Required("variables"+"."+k, "body", m.Variables[k]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextBotInputOutputData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextBotInputOutputData) UnmarshalBinary(b []byte) error {
	var res TextBotInputOutputData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}