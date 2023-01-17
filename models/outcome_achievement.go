// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OutcomeAchievement outcome achievement
//
// swagger:model OutcomeAchievement
type OutcomeAchievement struct {

	// Timestamp indicating when the outcome was achieved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	AchievedDate strfmt.DateTime `json:"achievedDate,omitempty"`

	// The outcome that was achieved.
	Outcome *AchievedOutcome `json:"outcome,omitempty"`
}

// Validate validates this outcome achievement
func (m *OutcomeAchievement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAchievedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutcome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutcomeAchievement) validateAchievedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.AchievedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("achievedDate", "body", "date-time", m.AchievedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OutcomeAchievement) validateOutcome(formats strfmt.Registry) error {
	if swag.IsZero(m.Outcome) { // not required
		return nil
	}

	if m.Outcome != nil {
		if err := m.Outcome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outcome")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outcome")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this outcome achievement based on the context it is used
func (m *OutcomeAchievement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOutcome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutcomeAchievement) contextValidateOutcome(ctx context.Context, formats strfmt.Registry) error {

	if m.Outcome != nil {
		if err := m.Outcome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outcome")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outcome")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutcomeAchievement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutcomeAchievement) UnmarshalBinary(b []byte) error {
	var res OutcomeAchievement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
