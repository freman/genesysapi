// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkspaceSummary workspace summary
//
// swagger:model WorkspaceSummary
type WorkspaceSummary struct {

	// total document byte count
	TotalDocumentByteCount int64 `json:"totalDocumentByteCount,omitempty"`

	// total document count
	TotalDocumentCount int64 `json:"totalDocumentCount,omitempty"`
}

// Validate validates this workspace summary
func (m *WorkspaceSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceSummary) UnmarshalBinary(b []byte) error {
	var res WorkspaceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
