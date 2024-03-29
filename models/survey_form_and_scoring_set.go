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

// SurveyFormAndScoringSet survey form and scoring set
//
// swagger:model SurveyFormAndScoringSet
type SurveyFormAndScoringSet struct {

	// answers
	Answers *SurveyScoringSet `json:"answers,omitempty"`

	// survey form
	SurveyForm *SurveyForm `json:"surveyForm,omitempty"`
}

// Validate validates this survey form and scoring set
func (m *SurveyFormAndScoringSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveyForm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SurveyFormAndScoringSet) validateAnswers(formats strfmt.Registry) error {
	if swag.IsZero(m.Answers) { // not required
		return nil
	}

	if m.Answers != nil {
		if err := m.Answers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("answers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("answers")
			}
			return err
		}
	}

	return nil
}

func (m *SurveyFormAndScoringSet) validateSurveyForm(formats strfmt.Registry) error {
	if swag.IsZero(m.SurveyForm) { // not required
		return nil
	}

	if m.SurveyForm != nil {
		if err := m.SurveyForm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surveyForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("surveyForm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this survey form and scoring set based on the context it is used
func (m *SurveyFormAndScoringSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnswers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSurveyForm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SurveyFormAndScoringSet) contextValidateAnswers(ctx context.Context, formats strfmt.Registry) error {

	if m.Answers != nil {
		if err := m.Answers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("answers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("answers")
			}
			return err
		}
	}

	return nil
}

func (m *SurveyFormAndScoringSet) contextValidateSurveyForm(ctx context.Context, formats strfmt.Registry) error {

	if m.SurveyForm != nil {
		if err := m.SurveyForm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surveyForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("surveyForm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SurveyFormAndScoringSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SurveyFormAndScoringSet) UnmarshalBinary(b []byte) error {
	var res SurveyFormAndScoringSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
