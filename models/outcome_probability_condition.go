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

// OutcomeProbabilityCondition outcome probability condition
//
// swagger:model OutcomeProbabilityCondition
type OutcomeProbabilityCondition struct {

	// Probability value for the selected outcome at or above which the action map will trigger.
	// Required: true
	MaximumProbability *float32 `json:"maximumProbability"`

	// The outcome ID.
	// Required: true
	OutcomeID *string `json:"outcomeId"`

	// Additional probability condition, where if set, the action map will trigger if the current outcome probability is lower or equal to the value.
	Probability float32 `json:"probability,omitempty"`
}

// Validate validates this outcome probability condition
func (m *OutcomeProbabilityCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaximumProbability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutcomeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutcomeProbabilityCondition) validateMaximumProbability(formats strfmt.Registry) error {

	if err := validate.Required("maximumProbability", "body", m.MaximumProbability); err != nil {
		return err
	}

	return nil
}

func (m *OutcomeProbabilityCondition) validateOutcomeID(formats strfmt.Registry) error {

	if err := validate.Required("outcomeId", "body", m.OutcomeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this outcome probability condition based on context it is used
func (m *OutcomeProbabilityCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OutcomeProbabilityCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutcomeProbabilityCondition) UnmarshalBinary(b []byte) error {
	var res OutcomeProbabilityCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
