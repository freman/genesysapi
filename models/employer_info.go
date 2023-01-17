// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmployerInfo employer info
//
// swagger:model EmployerInfo
type EmployerInfo struct {

	// date hire
	DateHire string `json:"dateHire,omitempty"`

	// employee Id
	EmployeeID string `json:"employeeId,omitempty"`

	// employee type
	EmployeeType string `json:"employeeType,omitempty"`

	// official name
	OfficialName string `json:"officialName,omitempty"`
}

// Validate validates this employer info
func (m *EmployerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this employer info based on context it is used
func (m *EmployerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmployerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmployerInfo) UnmarshalBinary(b []byte) error {
	var res EmployerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
