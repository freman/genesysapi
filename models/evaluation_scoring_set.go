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

	// Comments from the agent while reviewing evaluation results
	AgentComments string `json:"agentComments,omitempty"`

	// Indicates that at least one fatal question was answered without having the highest score available for the question
	AnyFailedKillQuestions bool `json:"anyFailedKillQuestions"`

	// Overall comments from the evaluator
	Comments string `json:"comments,omitempty"`

	// question group scores
	QuestionGroupScores []*EvaluationQuestionGroupScore `json:"questionGroupScores"`

	// Score of only the critical questions
	TotalCriticalScore float32 `json:"totalCriticalScore,omitempty"`

	// Score of only the non-critical questions
	TotalNonCriticalScore float32 `json:"totalNonCriticalScore,omitempty"`

	// Score of all questions
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
