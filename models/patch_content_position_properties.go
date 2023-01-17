// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchContentPositionProperties patch content position properties
//
// swagger:model PatchContentPositionProperties
type PatchContentPositionProperties struct {

	// Bottom positioning offset.
	Bottom string `json:"bottom,omitempty"`

	// Left positioning offset.
	Left string `json:"left,omitempty"`

	// Right positioning offset.
	Right string `json:"right,omitempty"`

	// Top positioning offset.
	Top string `json:"top,omitempty"`
}

// Validate validates this patch content position properties
func (m *PatchContentPositionProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch content position properties based on context it is used
func (m *PatchContentPositionProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchContentPositionProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchContentPositionProperties) UnmarshalBinary(b []byte) error {
	var res PatchContentPositionProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
