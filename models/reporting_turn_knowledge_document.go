// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReportingTurnKnowledgeDocument reporting turn knowledge document
//
// swagger:model ReportingTurnKnowledgeDocument
type ReportingTurnKnowledgeDocument struct {

	// The corresponding answer to the question.
	Answer string `json:"answer,omitempty"`

	// The confidence score of how well the question matched the search query.
	Confidence float64 `json:"confidence,omitempty"`

	// The ID of the knowledge document.
	ID string `json:"id,omitempty"`

	// The the question that was used to match against the search query.
	Question string `json:"question,omitempty"`
}

// Validate validates this reporting turn knowledge document
func (m *ReportingTurnKnowledgeDocument) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportingTurnKnowledgeDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingTurnKnowledgeDocument) UnmarshalBinary(b []byte) error {
	var res ReportingTurnKnowledgeDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
