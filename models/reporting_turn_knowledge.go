// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReportingTurnKnowledge reporting turn knowledge
//
// swagger:model ReportingTurnKnowledge
type ReportingTurnKnowledge struct {

	// The knowledge feedback data that was captured during this reporting turn.
	Feedback *ReportingTurnKnowledgeFeedback `json:"feedback,omitempty"`

	// The Knowledge Base ID that the captured knowledge data relates to.
	KnowledgeBaseID string `json:"knowledgeBaseId,omitempty"`

	// The knowledge search data that was captured during this reporting turn.
	Search *ReportingTurnKnowledgeSearch `json:"search,omitempty"`
}

// Validate validates this reporting turn knowledge
func (m *ReportingTurnKnowledge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeedback(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTurnKnowledge) validateFeedback(formats strfmt.Registry) error {
	if swag.IsZero(m.Feedback) { // not required
		return nil
	}

	if m.Feedback != nil {
		if err := m.Feedback.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feedback")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("feedback")
			}
			return err
		}
	}

	return nil
}

func (m *ReportingTurnKnowledge) validateSearch(formats strfmt.Registry) error {
	if swag.IsZero(m.Search) { // not required
		return nil
	}

	if m.Search != nil {
		if err := m.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this reporting turn knowledge based on the context it is used
func (m *ReportingTurnKnowledge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeedback(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTurnKnowledge) contextValidateFeedback(ctx context.Context, formats strfmt.Registry) error {

	if m.Feedback != nil {
		if err := m.Feedback.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feedback")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("feedback")
			}
			return err
		}
	}

	return nil
}

func (m *ReportingTurnKnowledge) contextValidateSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.Search != nil {
		if err := m.Search.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportingTurnKnowledge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingTurnKnowledge) UnmarshalBinary(b []byte) error {
	var res ReportingTurnKnowledge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
