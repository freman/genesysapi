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

// SurveyScoringSet survey scoring set
//
// swagger:model SurveyScoringSet
type SurveyScoringSet struct {

	// nps score
	NpsScore int32 `json:"npsScore,omitempty"`

	// question group scores
	QuestionGroupScores []*SurveyQuestionGroupScore `json:"questionGroupScores"`

	// total score
	TotalScore float32 `json:"totalScore,omitempty"`
}

// Validate validates this survey scoring set
func (m *SurveyScoringSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuestionGroupScores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SurveyScoringSet) validateQuestionGroupScores(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("questionGroupScores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this survey scoring set based on the context it is used
func (m *SurveyScoringSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuestionGroupScores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SurveyScoringSet) contextValidateQuestionGroupScores(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QuestionGroupScores); i++ {

		if m.QuestionGroupScores[i] != nil {
			if err := m.QuestionGroupScores[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("questionGroupScores" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("questionGroupScores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SurveyScoringSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SurveyScoringSet) UnmarshalBinary(b []byte) error {
	var res SurveyScoringSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
