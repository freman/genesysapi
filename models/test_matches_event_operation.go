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

// TestMatchesEventOperation Results from evaluating matching criteria against test input
//
// swagger:model TestMatchesEventOperation
type TestMatchesEventOperation struct {

	// Triggers that matched
	MatchedTriggers []*TestModeTrigger `json:"matchedTriggers"`

	// The name of the processing step
	Name string `json:"name,omitempty"`

	// The number of the processing step
	Step int32 `json:"step,omitempty"`

	// Triggers that did not match
	UnmatchedTriggers []*TestModeTrigger `json:"unmatchedTriggers"`
}

// Validate validates this test matches event operation
func (m *TestMatchesEventOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchedTriggers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnmatchedTriggers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestMatchesEventOperation) validateMatchedTriggers(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchedTriggers) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchedTriggers); i++ {
		if swag.IsZero(m.MatchedTriggers[i]) { // not required
			continue
		}

		if m.MatchedTriggers[i] != nil {
			if err := m.MatchedTriggers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchedTriggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchedTriggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestMatchesEventOperation) validateUnmatchedTriggers(formats strfmt.Registry) error {
	if swag.IsZero(m.UnmatchedTriggers) { // not required
		return nil
	}

	for i := 0; i < len(m.UnmatchedTriggers); i++ {
		if swag.IsZero(m.UnmatchedTriggers[i]) { // not required
			continue
		}

		if m.UnmatchedTriggers[i] != nil {
			if err := m.UnmatchedTriggers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unmatchedTriggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("unmatchedTriggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this test matches event operation based on the context it is used
func (m *TestMatchesEventOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatchedTriggers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnmatchedTriggers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestMatchesEventOperation) contextValidateMatchedTriggers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MatchedTriggers); i++ {

		if m.MatchedTriggers[i] != nil {
			if err := m.MatchedTriggers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchedTriggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchedTriggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestMatchesEventOperation) contextValidateUnmatchedTriggers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UnmatchedTriggers); i++ {

		if m.UnmatchedTriggers[i] != nil {
			if err := m.UnmatchedTriggers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unmatchedTriggers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("unmatchedTriggers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestMatchesEventOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestMatchesEventOperation) UnmarshalBinary(b []byte) error {
	var res TestMatchesEventOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
