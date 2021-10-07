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

// LearningAssessmentScoringRequest learning assessment scoring request
//
// swagger:model LearningAssessmentScoringRequest
type LearningAssessmentScoringRequest struct {

	// The answers to score
	// Required: true
	Answers *AssessmentScoringSet `json:"answers"`

	// The assessment form to score against
	// Required: true
	AssessmentForm *AssessmentForm `json:"assessmentForm"`
}

// Validate validates this learning assessment scoring request
func (m *LearningAssessmentScoringRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssessmentForm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningAssessmentScoringRequest) validateAnswers(formats strfmt.Registry) error {

	if err := validate.Required("answers", "body", m.Answers); err != nil {
		return err
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

func (m *LearningAssessmentScoringRequest) validateAssessmentForm(formats strfmt.Registry) error {

	if err := validate.Required("assessmentForm", "body", m.AssessmentForm); err != nil {
		return err
	}

	if m.AssessmentForm != nil {
		if err := m.AssessmentForm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessmentForm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningAssessmentScoringRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningAssessmentScoringRequest) UnmarshalBinary(b []byte) error {
	var res LearningAssessmentScoringRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
