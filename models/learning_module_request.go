// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LearningModuleRequest Learning module request
//
// swagger:model LearningModuleRequest
type LearningModuleRequest struct {

	// The assessment form for learning module
	AssessmentForm *AssessmentForm `json:"assessmentForm,omitempty"`

	// The completion time of learning module in days
	// Required: true
	CompletionTimeInDays *int32 `json:"completionTimeInDays"`

	// The cover art for the learning module
	CoverArt *LearningModuleCoverArtRequest `json:"coverArt,omitempty"`

	// The description of learning module
	Description string `json:"description,omitempty"`

	// The list of inform steps in a learning module
	InformSteps []*LearningModuleInformStepRequest `json:"informSteps"`

	// The name of learning module
	// Required: true
	Name *string `json:"name"`

	// The type for the learning module
	// Enum: [Informational AssessedContent Assessment External]
	Type string `json:"type,omitempty"`
}

// Validate validates this learning module request
func (m *LearningModuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssessmentForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletionTimeInDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoverArt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInformSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningModuleRequest) validateAssessmentForm(formats strfmt.Registry) error {
	if swag.IsZero(m.AssessmentForm) { // not required
		return nil
	}

	if m.AssessmentForm != nil {
		if err := m.AssessmentForm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessmentForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assessmentForm")
			}
			return err
		}
	}

	return nil
}

func (m *LearningModuleRequest) validateCompletionTimeInDays(formats strfmt.Registry) error {

	if err := validate.Required("completionTimeInDays", "body", m.CompletionTimeInDays); err != nil {
		return err
	}

	return nil
}

func (m *LearningModuleRequest) validateCoverArt(formats strfmt.Registry) error {
	if swag.IsZero(m.CoverArt) { // not required
		return nil
	}

	if m.CoverArt != nil {
		if err := m.CoverArt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coverArt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("coverArt")
			}
			return err
		}
	}

	return nil
}

func (m *LearningModuleRequest) validateInformSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.InformSteps) { // not required
		return nil
	}

	for i := 0; i < len(m.InformSteps); i++ {
		if swag.IsZero(m.InformSteps[i]) { // not required
			continue
		}

		if m.InformSteps[i] != nil {
			if err := m.InformSteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("informSteps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("informSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningModuleRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var learningModuleRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Informational","AssessedContent","Assessment","External"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningModuleRequestTypeTypePropEnum = append(learningModuleRequestTypeTypePropEnum, v)
	}
}

const (

	// LearningModuleRequestTypeInformational captures enum value "Informational"
	LearningModuleRequestTypeInformational string = "Informational"

	// LearningModuleRequestTypeAssessedContent captures enum value "AssessedContent"
	LearningModuleRequestTypeAssessedContent string = "AssessedContent"

	// LearningModuleRequestTypeAssessment captures enum value "Assessment"
	LearningModuleRequestTypeAssessment string = "Assessment"

	// LearningModuleRequestTypeExternal captures enum value "External"
	LearningModuleRequestTypeExternal string = "External"
)

// prop value enum
func (m *LearningModuleRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningModuleRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningModuleRequest) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this learning module request based on the context it is used
func (m *LearningModuleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssessmentForm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCoverArt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInformSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningModuleRequest) contextValidateAssessmentForm(ctx context.Context, formats strfmt.Registry) error {

	if m.AssessmentForm != nil {
		if err := m.AssessmentForm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessmentForm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assessmentForm")
			}
			return err
		}
	}

	return nil
}

func (m *LearningModuleRequest) contextValidateCoverArt(ctx context.Context, formats strfmt.Registry) error {

	if m.CoverArt != nil {
		if err := m.CoverArt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coverArt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("coverArt")
			}
			return err
		}
	}

	return nil
}

func (m *LearningModuleRequest) contextValidateInformSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InformSteps); i++ {

		if m.InformSteps[i] != nil {
			if err := m.InformSteps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("informSteps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("informSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningModuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningModuleRequest) UnmarshalBinary(b []byte) error {
	var res LearningModuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
