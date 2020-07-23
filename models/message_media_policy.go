// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessageMediaPolicy message media policy
//
// swagger:model MessageMediaPolicy
type MessageMediaPolicy struct {

	// Actions applied when specified conditions are met
	Actions *PolicyActions `json:"actions,omitempty"`

	// Conditions for when actions should be applied
	Conditions *MessageMediaPolicyConditions `json:"conditions,omitempty"`
}

// Validate validates this message media policy
func (m *MessageMediaPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageMediaPolicy) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {
		if err := m.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *MessageMediaPolicy) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	if m.Conditions != nil {
		if err := m.Conditions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageMediaPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageMediaPolicy) UnmarshalBinary(b []byte) error {
	var res MessageMediaPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
