// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EvaluationScoringSet evaluation scoring set
//
// swagger:model EvaluationScoringSet
type EvaluationScoringSet struct {

	// agent comments
	AgentComments string `json:"agentComments,omitempty"`

	// any failed kill questions
	AnyFailedKillQuestions bool `json:"anyFailedKillQuestions"`

	// comments
	Comments string `json:"comments,omitempty"`

	// question group scores
	QuestionGroupScores []*EvaluationQuestionGroupScore `json:"questionGroupScores"`

	// total critical score
	TotalCriticalScore float32 `json:"totalCriticalScore,omitempty"`

	// total non critical score
	TotalNonCriticalScore float32 `json:"totalNonCriticalScore,omitempty"`

	// total score
	TotalScore float32 `json:"totalScore,omitempty"`
}

// Validate validates this evaluation scoring set
func (m *EvaluationScoringSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuestionGroupScores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvaluationScoringSet) validateQuestionGroupScores(formats strfmt.Registry) error {

	if swag.IsZero(m.QuestionGroupScores) { // not required
		return nil
	}

	for i := 0; i < len(m.QuestionGroupScores); i++ {
		if swag.IsZero(m.QuestionGroupScores[i]) { // not required
			continue
		}

		if m.QuestionGroupScores[i] != nil {
			if err := m.QuestionGroupScores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("questionGroupScores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationScoringSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationScoringSet) UnmarshalBinary(b []byte) error {
	var res EvaluationScoringSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}