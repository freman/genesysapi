// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TemplateParameter template parameter
//
// swagger:model TemplateParameter
type TemplateParameter struct {

	// Response substitution identifier
	ID string `json:"id,omitempty"`

	// Response substitution value
	Value string `json:"value,omitempty"`
}

// Validate validates this template parameter
func (m *TemplateParameter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplateParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateParameter) UnmarshalBinary(b []byte) error {
	var res TemplateParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
