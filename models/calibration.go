// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Calibration calibration
//
// swagger:model Calibration
type Calibration struct {

	// agent
	Agent *User `json:"agent,omitempty"`

	// average score
	AverageScore int32 `json:"averageScore,omitempty"`

	// calibrator
	Calibrator *User `json:"calibrator,omitempty"`

	// context Id
	ContextID string `json:"contextId,omitempty"`

	// conversation
	Conversation *Conversation `json:"conversation,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// evaluation form
	EvaluationForm *EvaluationForm `json:"evaluationForm,omitempty"`

	// evaluations
	Evaluations []*Evaluation `json:"evaluations"`

	// evaluators
	Evaluators []*User `json:"evaluators"`

	// expert evaluator
	ExpertEvaluator *User `json:"expertEvaluator,omitempty"`

	// high score
	HighScore int32 `json:"highScore,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// low score
	LowScore int32 `json:"lowScore,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// scoring index
	ScoringIndex *Evaluation `json:"scoringIndex,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this calibration
func (m *Calibration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalibrator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluationForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpertEvaluator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoringIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Calibration) validateAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *Calibration) validateCalibrator(formats strfmt.Registry) error {

	if swag.IsZero(m.Calibrator) { // not required
		return nil
	}

	if m.Calibrator != nil {
		if err := m.Calibrator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calibrator")
			}
			return err
		}
	}

	return nil
}

func (m *Calibration) validateConversation(formats strfmt.Registry) error {

	if swag.IsZero(m.Conversation) { // not required
		return nil
	}

	if m.Conversation != nil {
		if err := m.Conversation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversation")
			}
			return err
		}
	}

	return nil
}

func (m *Calibration) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Calibration) validateEvaluationForm(formats strfmt.Registry) error {

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

func (m *Calibration) validateEvaluations(formats strfmt.Registry) error {

	if swag.IsZero(m.Evaluations) { // not required
		return nil
	}

	for i := 0; i < len(m.Evaluations); i++ {
		if swag.IsZero(m.Evaluations[i]) { // not required
			continue
		}

		if m.Evaluations[i] != nil {
			if err := m.Evaluations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evaluations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Calibration) validateEvaluators(formats strfmt.Registry) error {

	if swag.IsZero(m.Evaluators) { // not required
		return nil
	}

	for i := 0; i < len(m.Evaluators); i++ {
		if swag.IsZero(m.Evaluators[i]) { // not required
			continue
		}

		if m.Evaluators[i] != nil {
			if err := m.Evaluators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evaluators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Calibration) validateExpertEvaluator(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpertEvaluator) { // not required
		return nil
	}

	if m.ExpertEvaluator != nil {
		if err := m.ExpertEvaluator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expertEvaluator")
			}
			return err
		}
	}

	return nil
}

func (m *Calibration) validateScoringIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.ScoringIndex) { // not required
		return nil
	}

	if m.ScoringIndex != nil {
		if err := m.ScoringIndex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scoringIndex")
			}
			return err
		}
	}

	return nil
}

func (m *Calibration) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Calibration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Calibration) UnmarshalBinary(b []byte) error {
	var res Calibration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
