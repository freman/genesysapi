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

// ResponseSubstitution Contains information about the substitutions associated with a response.
//
// swagger:model ResponseSubstitution
type ResponseSubstitution struct {

	// Response substitution default value.
	DefaultValue string `json:"defaultValue,omitempty"`

	// Response substitution description.
	Description string `json:"description,omitempty"`

	// Response substitution identifier.
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this response substitution
func (m *ResponseSubstitution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseSubstitution) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseSubstitution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseSubstitution) UnmarshalBinary(b []byte) error {
	var res ResponseSubstitution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
