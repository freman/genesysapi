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
	"github.com/go-openapi/validate"
)

// AssessmentFormQuestionGroup assessment form question group
//
// swagger:model AssessmentFormQuestionGroup
type AssessmentFormQuestionGroup struct {

	// default answers to highest
	DefaultAnswersToHighest bool `json:"defaultAnswersToHighest"`

	// default answers to n a
	DefaultAnswersToNA bool `json:"defaultAnswersToNA"`

	// The ID of the question group,
	ID string `json:"id,omitempty"`

	// manual weight
	ManualWeight bool `json:"manualWeight"`

	// na enabled
	NaEnabled bool `json:"naEnabled"`

	// The question group name
	// Required: true
	Name *string `json:"name"`

	// The list of questions for this question group
	// Required: true
	Questions []*AssessmentFormQuestion `json:"questions"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The question group type
	// Required: true
	Type *string `json:"type"`

	// visibility condition
	VisibilityCondition *VisibilityCondition `json:"visibilityCondition,omitempty"`

	// weight
	Weight float32 `json:"weight,omitempty"`
}

// Validate validates this assessment form question group
func (m *AssessmentFormQuestionGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuestions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *AssessmentFormQuestionGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AssessmentFormQuestionGroup) validateQuestions(formats strfmt.Registry) error {

	if err := validate.Required("questions", "body", m.Questions); err != nil {
		return err
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

func (m *AssessmentFormQuestionGroup) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AssessmentFormQuestionGroup) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AssessmentFormQuestionGroup) validateVisibilityCondition(formats strfmt.Registry) error {
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

// ContextValidate validate this assessment form question group based on the context it is used
func (m *AssessmentFormQuestionGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuestions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
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

func (m *AssessmentFormQuestionGroup) contextValidateQuestions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AssessmentFormQuestionGroup) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *AssessmentFormQuestionGroup) contextValidateVisibilityCondition(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AssessmentFormQuestionGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssessmentFormQuestionGroup) UnmarshalBinary(b []byte) error {
	var res AssessmentFormQuestionGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
