// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CobrowseSettings Settings concerning cobrowse
//
// swagger:model CobrowseSettings
type CobrowseSettings struct {

	// Whether the viewer should have option to request control
	AllowAgentControl bool `json:"allowAgentControl"`

	// Whether or not cobrowse is enabled
	Enabled bool `json:"enabled"`

	// Mask patterns that will apply to pages being shared
	MaskSelectors []string `json:"maskSelectors"`
}

// Validate validates this cobrowse settings
func (m *CobrowseSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cobrowse settings based on context it is used
func (m *CobrowseSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CobrowseSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CobrowseSettings) UnmarshalBinary(b []byte) error {
	var res CobrowseSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
