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

// EdgeNetworkDiagnosticRequest edge network diagnostic request
//
// swagger:model EdgeNetworkDiagnosticRequest
type EdgeNetworkDiagnosticRequest struct {

	// IPv4/6 address or host to be probed for connectivity. No port allowed.
	// Required: true
	Host *string `json:"host"`
}

// Validate validates this edge network diagnostic request
func (m *EdgeNetworkDiagnosticRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeNetworkDiagnosticRequest) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeNetworkDiagnosticRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeNetworkDiagnosticRequest) UnmarshalBinary(b []byte) error {
	var res EdgeNetworkDiagnosticRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
