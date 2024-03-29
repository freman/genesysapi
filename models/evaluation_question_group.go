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

// EvaluationQuestionGroup evaluation question group
//
// swagger:model EvaluationQuestionGroup
type EvaluationQuestionGroup struct {

	// default answers to highest
	DefaultAnswersToHighest bool `json:"defaultAnswersToHighest"`

	// default answers to n a
	DefaultAnswersToNA bool `json:"defaultAnswersToNA"`

	// id
	ID string `json:"id,omitempty"`

	// manual weight
	ManualWeight bool `json:"manualWeight"`

	// na enabled
	NaEnabled bool `json:"naEnabled"`

	// name
	Name string `json:"name,omitempty"`

	// questions
	Questions []*EvaluationQuestion `json:"questions"`

	// type
	Type string `json:"type,omitempty"`

	// visibility condition
	VisibilityCondition *VisibilityCondition `json:"visibilityCondition,omitempty"`

	// weight
	Weight float32 `json:"weight,omitempty"`
}

// Validate validates this evaluation question group
func (m *EvaluationQuestionGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuestions(formats); err != nil {
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

func (m *EvaluationQuestionGroup) validateQuestions(formats strfmt.Registry) error {
	if swag.IsZero(m.Questions) { // not required
		return nil
	}

	for i := 0; i < len(m.Questions); i++ {
		if swag.IsZero(m.Questions[i]) { // not required
			continue
		}

		if m.Questions[i] != nil {
			if err := m.Questions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("questions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("questions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EvaluationQuestionGroup) validateVisibilityCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.VisibilityCondition) { // not required
		return nil
	}

	if m.VisibilityCondition != nil {
		if err := m.VisibilityCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibilityCondition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("visibilityCondition")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this evaluation question group based on the context it is used
func (m *EvaluationQuestionGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuestions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVisibilityCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvaluationQuestionGroup) contextValidateQuestions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Questions); i++ {

		if m.Questions[i] != nil {
			if err := m.Questions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("questions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("questions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EvaluationQuestionGroup) contextValidateVisibilityCondition(ctx context.Context, formats strfmt.Registry) error {

	if m.VisibilityCondition != nil {
		if err := m.VisibilityCondition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibilityCondition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("visibilityCondition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationQuestionGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationQuestionGroup) UnmarshalBinary(b []byte) error {
	var res EvaluationQuestionGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
