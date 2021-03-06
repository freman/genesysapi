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

// DialogflowIntent dialogflow intent
//
// swagger:model DialogflowIntent
type DialogflowIntent struct {

	// The intent name
	// Required: true
	Name *string `json:"name"`

	// An object mapping parameter names to Parameter objects
	// Required: true
	Parameters map[string]DialogflowParameter `json:"parameters"`
}

// Validate validates this dialogflow intent
func (m *DialogflowIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DialogflowIntent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DialogflowIntent) validateParameters(formats strfmt.Registry) error {

	for k := range m.Parameters {

		if err := validate.Required("parameters"+"."+k, "body", m.Parameters[k]); err != nil {
			return err
		}
		if val, ok := m.Parameters[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DialogflowIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DialogflowIntent) UnmarshalBinary(b []byte) error {
	var res DialogflowIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
