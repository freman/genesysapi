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

// ReportingTurnKnowledgeSearch reporting turn knowledge search
//
// swagger:model ReportingTurnKnowledgeSearch
type ReportingTurnKnowledgeSearch struct {

	// The list of search documents captured during this reporting turn.
	Documents []*ReportingTurnKnowledgeDocument `json:"documents"`

	// The search query that was used to search the Knowledge Base documents for a matching question.
	Query string `json:"query,omitempty"`

	// The ID of this knowledge search.
	SearchID string `json:"searchId,omitempty"`
}

// Validate validates this reporting turn knowledge search
func (m *ReportingTurnKnowledgeSearch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTurnKnowledgeSearch) validateDocuments(formats strfmt.Registry) error {
	if swag.IsZero(m.Documents) { // not required
		return nil
	}

	for i := 0; i < len(m.Documents); i++ {
		if swag.IsZero(m.Documents[i]) { // not required
			continue
		}

		if m.Documents[i] != nil {
			if err := m.Documents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("documents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this reporting turn knowledge search based on the context it is used
func (m *ReportingTurnKnowledgeSearch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocuments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTurnKnowledgeSearch) contextValidateDocuments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Documents); i++ {

		if m.Documents[i] != nil {
			if err := m.Documents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("documents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportingTurnKnowledgeSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingTurnKnowledgeSearch) UnmarshalBinary(b []byte) error {
	var res ReportingTurnKnowledgeSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
