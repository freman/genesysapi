// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AvailableTranslations available translations
//
// swagger:model AvailableTranslations
type AvailableTranslations struct {

	// builtin
	Builtin []string `json:"builtin"`

	// org specific
	OrgSpecific []string `json:"orgSpecific"`
}

// Validate validates this available translations
func (m *AvailableTranslations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this available translations based on context it is used
func (m *AvailableTranslations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvailableTranslations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableTranslations) UnmarshalBinary(b []byte) error {
	var res AvailableTranslations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
