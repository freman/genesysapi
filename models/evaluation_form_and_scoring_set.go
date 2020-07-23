// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EvaluationFormAndScoringSet evaluation form and scoring set
//
// swagger:model EvaluationFormAndScoringSet
type EvaluationFormAndScoringSet struct {

	// answers
	Answers *EvaluationScoringSet `json:"answers,omitempty"`

	// evaluation form
	EvaluationForm *EvaluationForm `json:"evaluationForm,omitempty"`
}

// Validate validates this evaluation form and scoring set
func (m *EvaluationFormAndScoringSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluationForm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvaluationFormAndScoringSet) validateAnswers(formats strfmt.Registry) error {

	if swag.IsZero(m.Answers) { // not required
		return nil
	}

	if m.Answers != nil {
		if err := m.Answers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("answers")
			}
			return err
		}
	}

	return nil
}

func (m *EvaluationFormAndScoringSet) validateEvaluationForm(formats strfmt.Registry) error {

	if swag.IsZero(m.EvaluationForm) { // not required
		return nil
	}

	if m.EvaluationForm != nil {
		if err := m.EvaluationForm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluationForm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationFormAndScoringSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationFormAndScoringSet) UnmarshalBinary(b []byte) error {
	var res EvaluationFormAndScoringSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
