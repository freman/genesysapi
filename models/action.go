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

// Action action
//
// swagger:model Action
type Action struct {

	// Category of Action
	Category string `json:"category,omitempty"`

	// Configuration to support request and response processing
	Config *ActionConfig `json:"config,omitempty"`

	// Action contract
	Contract *ActionContract `json:"contract,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The ID of the integration for which this action is associated
	IntegrationID string `json:"integrationId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Indication of whether or not the action is designed to accept sensitive data
	Secure bool `json:"secure,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Version of this action
	Version int32 `json:"version,omitempty"`
}

// Validate validates this action
func (m *Action) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContract(formats); err != nil {
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

func (m *Action) validateConfig(formats strfmt.Registry) error {

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

func (m *Action) validateContract(formats strfmt.Registry) error {

	if swag.IsZero(m.Contract) { // not required
		return nil
	}

	if m.Contract != nil {
		if err := m.Contract.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contract")
			}
			return err
		}
	}

	return nil
}

func (m *Action) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Action) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Action) UnmarshalBinary(b []byte) error {
	var res Action
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
