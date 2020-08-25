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

// EffectiveConfiguration Effective Configuration for an ClientApp. This is comprised of the integration specific configuration along with overrides specified in the integration type.
//
// swagger:model EffectiveConfiguration
type EffectiveConfiguration struct {

	// Advanced configuration described by the schema in the advancedSchemaUri field.
	// Required: true
	Advanced map[string]interface{} `json:"advanced"`

	// Credentials required by the integration. The required keys are indicated in the credentials property of the Integration Type
	// Required: true
	Credentials map[string]CredentialInfo `json:"credentials"`

	// The name of the integration, used to distinguish this integration from others of the same type.
	// Required: true
	Name *string `json:"name"`

	// Notes about the integration.
	// Required: true
	Notes *string `json:"notes"`

	// Key-value configuration settings described by the schema in the propertiesSchemaUri field.
	// Required: true
	Properties map[string]interface{} `json:"properties"`
}

// Validate validates this effective configuration
func (m *EffectiveConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvanced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EffectiveConfiguration) validateAdvanced(formats strfmt.Registry) error {

	for k := range m.Advanced {

		if err := validate.Required("advanced"+"."+k, "body", m.Advanced[k]); err != nil {
			return err
		}

		if err := validate.Required("advanced"+"."+k, "body", m.Advanced[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *EffectiveConfiguration) validateCredentials(formats strfmt.Registry) error {

	for k := range m.Credentials {

		if err := validate.Required("credentials"+"."+k, "body", m.Credentials[k]); err != nil {
			return err
		}
		if val, ok := m.Credentials[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *EffectiveConfiguration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EffectiveConfiguration) validateNotes(formats strfmt.Registry) error {

	if err := validate.Required("notes", "body", m.Notes); err != nil {
		return err
	}

	return nil
}

func (m *EffectiveConfiguration) validateProperties(formats strfmt.Registry) error {

	for k := range m.Properties {

		if err := validate.Required("properties"+"."+k, "body", m.Properties[k]); err != nil {
			return err
		}

		if err := validate.Required("properties"+"."+k, "body", m.Properties[k]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EffectiveConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EffectiveConfiguration) UnmarshalBinary(b []byte) error {
	var res EffectiveConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}