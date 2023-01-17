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

// SearchUpdateRequest search update request
//
// swagger:model SearchUpdateRequest
type SearchUpdateRequest struct {

	// Mark the search as answered/unanswered
	// Required: true
	Answered *bool `json:"answered"`

	// The selected search result chosen as the answer.
	SelectedAnswer *SelectedAnswer `json:"selectedAnswer,omitempty"`

	// The unique identifier of this session
	// Read Only: true
	SessionID string `json:"sessionId,omitempty"`
}

// Validate validates this search update request
func (m *SearchUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectedAnswer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchUpdateRequest) validateAnswered(formats strfmt.Registry) error {

	if err := validate.Required("answered", "body", m.Answered); err != nil {
		return err
	}

	return nil
}

func (m *SearchUpdateRequest) validateSelectedAnswer(formats strfmt.Registry) error {
	if swag.IsZero(m.SelectedAnswer) { // not required
		return nil
	}

	if m.SelectedAnswer != nil {
		if err := m.SelectedAnswer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selectedAnswer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selectedAnswer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this search update request based on the context it is used
func (m *SearchUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelectedAnswer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchUpdateRequest) contextValidateSelectedAnswer(ctx context.Context, formats strfmt.Registry) error {

	if m.SelectedAnswer != nil {
		if err := m.SelectedAnswer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selectedAnswer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selectedAnswer")
			}
			return err
		}
	}

	return nil
}

func (m *SearchUpdateRequest) contextValidateSessionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sessionId", "body", string(m.SessionID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchUpdateRequest) UnmarshalBinary(b []byte) error {
	var res SearchUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
