// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyErrors policy errors
//
// swagger:model PolicyErrors
type PolicyErrors struct {

	// policy error messages
	PolicyErrorMessages []*PolicyErrorMessage `json:"policyErrorMessages"`
}

// Validate validates this policy errors
func (m *PolicyErrors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyErrorMessages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyErrors) validatePolicyErrorMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyErrorMessages) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyErrorMessages); i++ {
		if swag.IsZero(m.PolicyErrorMessages[i]) { // not required
			continue
		}

		if m.PolicyErrorMessages[i] != nil {
			if err := m.PolicyErrorMessages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyErrorMessages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policyErrorMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this policy errors based on the context it is used
func (m *PolicyErrors) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicyErrorMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyErrors) contextValidatePolicyErrorMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyErrorMessages); i++ {

		if m.PolicyErrorMessages[i] != nil {
			if err := m.PolicyErrorMessages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyErrorMessages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policyErrorMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyErrors) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyErrors) UnmarshalBinary(b []byte) error {
	var res PolicyErrors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
