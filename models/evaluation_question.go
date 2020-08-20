// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EvaluationQuestion evaluation question
//
// swagger:model EvaluationQuestion
type EvaluationQuestion struct {

	// Options from which to choose an answer for this question. Only used by Multiple Choice type questions.
	AnswerOptions []*AnswerOption `json:"answerOptions"`

	// comments required
	CommentsRequired bool `json:"commentsRequired"`

	// help text
	HelpText string `json:"helpText,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is critical
	IsCritical bool `json:"isCritical"`

	// is kill
	IsKill bool `json:"isKill"`

	// na enabled
	NaEnabled bool `json:"naEnabled"`

	// text
	Text string `json:"text,omitempty"`

	// type
	// Enum: [multipleChoiceQuestion freeTextQuestion npsQuestion readOnlyTextBlockQuestion]
	Type string `json:"type,omitempty"`

	// visibility condition
	VisibilityCondition *VisibilityCondition `json:"visibilityCondition,omitempty"`
}

// Validate validates this evaluation question
func (m *EvaluationQuestion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswerOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibilityCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvaluationQuestion) validateAnswerOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.AnswerOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.AnswerOptions); i++ {
		if swag.IsZero(m.AnswerOptions[i]) { // not required
			continue
		}

		if m.AnswerOptions[i] != nil {
			if err := m.AnswerOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("answerOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var evaluationQuestionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["multipleChoiceQuestion","freeTextQuestion","npsQuestion","readOnlyTextBlockQuestion"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationQuestionTypeTypePropEnum = append(evaluationQuestionTypeTypePropEnum, v)
	}
}

const (

	// EvaluationQuestionTypeMultipleChoiceQuestion captures enum value "multipleChoiceQuestion"
	EvaluationQuestionTypeMultipleChoiceQuestion string = "multipleChoiceQuestion"

	// EvaluationQuestionTypeFreeTextQuestion captures enum value "freeTextQuestion"
	EvaluationQuestionTypeFreeTextQuestion string = "freeTextQuestion"

	// EvaluationQuestionTypeNpsQuestion captures enum value "npsQuestion"
	EvaluationQuestionTypeNpsQuestion string = "npsQuestion"

	// EvaluationQuestionTypeReadOnlyTextBlockQuestion captures enum value "readOnlyTextBlockQuestion"
	EvaluationQuestionTypeReadOnlyTextBlockQuestion string = "readOnlyTextBlockQuestion"
)

// prop value enum
func (m *EvaluationQuestion) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationQuestionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationQuestion) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *EvaluationQuestion) validateVisibilityCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.VisibilityCondition) { // not required
		return nil
	}

	if m.VisibilityCondition != nil {
		if err := m.VisibilityCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibilityCondition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationQuestion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationQuestion) UnmarshalBinary(b []byte) error {
	var res EvaluationQuestion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
