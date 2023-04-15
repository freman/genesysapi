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

// OutcomePredictorRequest outcome predictor request
//
// swagger:model OutcomePredictorRequest
type OutcomePredictorRequest struct {

	// The outcome for which this predictor will provide predictions.
	// Required: true
	Outcome *OutcomeRefRequest `json:"outcome"`
}

// Validate validates this outcome predictor request
func (m *OutcomePredictorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutcome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutcomePredictorRequest) validateOutcome(formats strfmt.Registry) error {

	if err := validate.Required("outcome", "body", m.Outcome); err != nil {
		return err
	}

	if m.Outcome != nil {
		if err := m.Outcome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outcome")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outcome")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this outcome predictor request based on the context it is used
func (m *OutcomePredictorRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOutcome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutcomePredictorRequest) contextValidateOutcome(ctx context.Context, formats strfmt.Registry) error {

	if m.Outcome != nil {
		if err := m.Outcome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outcome")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outcome")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutcomePredictorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutcomePredictorRequest) UnmarshalBinary(b []byte) error {
	var res OutcomePredictorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
