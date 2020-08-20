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

// UpdateDraftInput Definition of an Action Draft to be created or updated.
//
// swagger:model UpdateDraftInput
type UpdateDraftInput struct {

	// Category of action
	Category string `json:"category,omitempty"`

	// Configuration to support request and response processing
	Config *ActionConfig `json:"config,omitempty"`

	// Action contract
	Contract *ActionContractInput `json:"contract,omitempty"`

	// Name of action
	Name string `json:"name,omitempty"`

	// Indication of whether or not the action is designed to accept sensitive data
	Secure bool `json:"secure"`

	// Version of current Draft
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this update draft input
func (m *UpdateDraftInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContract(formats); err != nil {
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

func (m *UpdateDraftInput) validateConfig(formats strfmt.Registry) error {

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

func (m *UpdateDraftInput) validateContract(formats strfmt.Registry) error {

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

func (m *UpdateDraftInput) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDraftInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDraftInput) UnmarshalBinary(b []byte) error {
	var res UpdateDraftInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
