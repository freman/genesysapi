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

// EvaluatorActivity evaluator activity
//
// swagger:model EvaluatorActivity
type EvaluatorActivity struct {

	// evaluator
	Evaluator *User `json:"evaluator,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// num calibrations assigned
	NumCalibrationsAssigned int32 `json:"numCalibrationsAssigned,omitempty"`

	// num calibrations completed
	NumCalibrationsCompleted int32 `json:"numCalibrationsCompleted,omitempty"`

	// num calibrations started
	NumCalibrationsStarted int32 `json:"numCalibrationsStarted,omitempty"`

	// num evaluations assigned
	NumEvaluationsAssigned int32 `json:"numEvaluationsAssigned,omitempty"`

	// num evaluations completed
	NumEvaluationsCompleted int32 `json:"numEvaluationsCompleted,omitempty"`

	// num evaluations started
	NumEvaluationsStarted int32 `json:"numEvaluationsStarted,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this evaluator activity
func (m *EvaluatorActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvaluator(formats); err != nil {
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

func (m *EvaluatorActivity) validateEvaluator(formats strfmt.Registry) error {

	if swag.IsZero(m.Evaluator) { // not required
		return nil
	}

	if m.Evaluator != nil {
		if err := m.Evaluator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluator")
			}
			return err
		}
	}

	return nil
}

func (m *EvaluatorActivity) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvaluatorActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluatorActivity) UnmarshalBinary(b []byte) error {
	var res EvaluatorActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
