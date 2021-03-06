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

// MeteredAssignmentByAgent metered assignment by agent
//
// swagger:model MeteredAssignmentByAgent
type MeteredAssignmentByAgent struct {

	// evaluation context Id
	EvaluationContextID string `json:"evaluationContextId,omitempty"`

	// evaluation form
	EvaluationForm *EvaluationForm `json:"evaluationForm,omitempty"`

	// evaluators
	Evaluators []*User `json:"evaluators"`

	// max number evaluations
	MaxNumberEvaluations int32 `json:"maxNumberEvaluations,omitempty"`

	// time interval
	TimeInterval *TimeInterval `json:"timeInterval,omitempty"`

	// time zone
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this metered assignment by agent
func (m *MeteredAssignmentByAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvaluationForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MeteredAssignmentByAgent) validateEvaluationForm(formats strfmt.Registry) error {

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

func (m *MeteredAssignmentByAgent) validateEvaluators(formats strfmt.Registry) error {

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

func (m *MeteredAssignmentByAgent) validateTimeInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInterval) { // not required
		return nil
	}

	if m.TimeInterval != nil {
		if err := m.TimeInterval.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeInterval")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MeteredAssignmentByAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeteredAssignmentByAgent) UnmarshalBinary(b []byte) error {
	var res MeteredAssignmentByAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
