// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OutcomeEventScore outcome event score
//
// swagger:model OutcomeEventScore
type OutcomeEventScore struct {

	// The outcome that the score was calculated for.
	Outcome *AddressableEntityRef `json:"outcome,omitempty"`

	// Represents the likelihood of a customer reaching or achieving a given outcome.
	Probability float32 `json:"probability,omitempty"`

	// Represents the max probability reached in the session.
	SessionMaxProbability float32 `json:"sessionMaxProbability,omitempty"`
}

// Validate validates this outcome event score
func (m *OutcomeEventScore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutcome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutcomeEventScore) validateOutcome(formats strfmt.Registry) error {

	if swag.IsZero(m.Outcome) { // not required
		return nil
	}

	if m.Outcome != nil {
		if err := m.Outcome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outcome")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutcomeEventScore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutcomeEventScore) UnmarshalBinary(b []byte) error {
	var res OutcomeEventScore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
