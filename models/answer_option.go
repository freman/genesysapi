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

// AnswerOption answer option
//
// swagger:model AnswerOption
type AnswerOption struct {

	// List of assistance conditions which are combined together with a logical AND operator. Eg ( assistanceCondtion1 && assistanceCondition2 ) wherein assistanceCondition could be ( EXISTS topic1 || topic2 || ... ) or (NOTEXISTS topic3 || topic4 || ...).
	AssistanceConditions []*AssistanceCondition `json:"assistanceConditions"`

	// id
	ID string `json:"id,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// value
	Value int32 `json:"value,omitempty"`
}

// Validate validates this answer option
func (m *AnswerOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssistanceConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnswerOption) validateAssistanceConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.AssistanceConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.AssistanceConditions); i++ {
		if swag.IsZero(m.AssistanceConditions[i]) { // not required
			continue
		}

		if m.AssistanceConditions[i] != nil {
			if err := m.AssistanceConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assistanceConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assistanceConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this answer option based on the context it is used
func (m *AnswerOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssistanceConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnswerOption) contextValidateAssistanceConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AssistanceConditions); i++ {

		if m.AssistanceConditions[i] != nil {
			if err := m.AssistanceConditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assistanceConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assistanceConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnswerOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnswerOption) UnmarshalBinary(b []byte) error {
	var res AnswerOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
