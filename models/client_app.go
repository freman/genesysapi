// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClientApp Details for a ClientApp
//
// swagger:model ClientApp
type ClientApp struct {

	// Read-only attributes for the integration.
	// Read Only: true
	Attributes map[string]string `json:"attributes,omitempty"`

	// Configuration information for the integration.
	// Read Only: true
	Config *ClientAppConfigurationInfo `json:"config,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Type of the integration
	// Read Only: true
	IntegrationType *IntegrationType `json:"integrationType,omitempty"`

	// Configured state of the integration.
	// Required: true
	// Enum: [ENABLED DISABLED DELETED]
	IntendedState *string `json:"intendedState"`

	// The name of the integration, used to distinguish this integration from others of the same type.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Notes about the integration.
	// Read Only: true
	Notes string `json:"notes,omitempty"`

	// Last reported status of the integration.
	// Read Only: true
	ReportedState *IntegrationStatusInfo `json:"reportedState,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this client app
func (m *ClientApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntendedState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportedState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientApp) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *ClientApp) validateIntegrationType(formats strfmt.Registry) error {

	if swag.IsZero(m.IntegrationType) { // not required
		return nil
	}

	if m.IntegrationType != nil {
		if err := m.IntegrationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrationType")
			}
			return err
		}
	}

	return nil
}

var clientAppTypeIntendedStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientAppTypeIntendedStatePropEnum = append(clientAppTypeIntendedStatePropEnum, v)
	}
}

const (

	// ClientAppIntendedStateENABLED captures enum value "ENABLED"
	ClientAppIntendedStateENABLED string = "ENABLED"

	// ClientAppIntendedStateDISABLED captures enum value "DISABLED"
	ClientAppIntendedStateDISABLED string = "DISABLED"

	// ClientAppIntendedStateDELETED captures enum value "DELETED"
	ClientAppIntendedStateDELETED string = "DELETED"
)

// prop value enum
func (m *ClientApp) validateIntendedStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientAppTypeIntendedStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientApp) validateIntendedState(formats strfmt.Registry) error {

	if err := validate.Required("intendedState", "body", m.IntendedState); err != nil {
		return err
	}

	// value enum
	if err := m.validateIntendedStateEnum("intendedState", "body", *m.IntendedState); err != nil {
		return err
	}

	return nil
}

func (m *ClientApp) validateReportedState(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportedState) { // not required
		return nil
	}

	if m.ReportedState != nil {
		if err := m.ReportedState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reportedState")
			}
			return err
		}
	}

	return nil
}

func (m *ClientApp) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientApp) UnmarshalBinary(b []byte) error {
	var res ClientApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
