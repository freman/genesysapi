// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailMediaPolicy email media policy
//
// swagger:model EmailMediaPolicy
type EmailMediaPolicy struct {

	// Actions applied when specified conditions are met
	Actions *PolicyActions `json:"actions,omitempty"`

	// Conditions for when actions should be applied
	Conditions *EmailMediaPolicyConditions `json:"conditions,omitempty"`
}

// Validate validates this email media policy
func (m *EmailMediaPolicy) Validate(formats strfmt.Registry) error {
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

func (m *EmailMediaPolicy) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {
		if err := m.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *EmailMediaPolicy) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	if m.Conditions != nil {
		if err := m.Conditions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this email media policy based on the context it is used
func (m *EmailMediaPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailMediaPolicy) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	if m.Actions != nil {
		if err := m.Actions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *EmailMediaPolicy) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	if m.Conditions != nil {
		if err := m.Conditions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailMediaPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailMediaPolicy) UnmarshalBinary(b []byte) error {
	var res EmailMediaPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
