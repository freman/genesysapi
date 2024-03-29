// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebDeploymentConfigurationVersionEntityRef web deployment configuration version entity ref
//
// swagger:model WebDeploymentConfigurationVersionEntityRef
type WebDeploymentConfigurationVersionEntityRef struct {

	// The configuration version ID
	// Required: true
	ID *string `json:"id"`

	// The configuration version name
	Name string `json:"name,omitempty"`

	// self Uri
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The version of the configuration
	// Example: DRAFT, 1, 2
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this web deployment configuration version entity ref
func (m *WebDeploymentConfigurationVersionEntityRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebDeploymentConfigurationVersionEntityRef) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *WebDeploymentConfigurationVersionEntityRef) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WebDeploymentConfigurationVersionEntityRef) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this web deployment configuration version entity ref based on context it is used
func (m *WebDeploymentConfigurationVersionEntityRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebDeploymentConfigurationVersionEntityRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebDeploymentConfigurationVersionEntityRef) UnmarshalBinary(b []byte) error {
	var res WebDeploymentConfigurationVersionEntityRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
