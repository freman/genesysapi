// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditQueryServiceMapping audit query service mapping
//
// swagger:model AuditQueryServiceMapping
type AuditQueryServiceMapping struct {

	// List of Services
	Services []*AuditQueryService `json:"services"`
}

// Validate validates this audit query service mapping
func (m *AuditQueryServiceMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditQueryServiceMapping) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditQueryServiceMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditQueryServiceMapping) UnmarshalBinary(b []byte) error {
	var res AuditQueryServiceMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
