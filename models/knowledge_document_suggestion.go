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

// KnowledgeDocumentSuggestion knowledge document suggestion
//
// swagger:model KnowledgeDocumentSuggestion
type KnowledgeDocumentSuggestion struct {

	// Page size of the returned results.
	PageSize int32 `json:"pageSize,omitempty"`

	// Query to get autocomplete suggestions for the matching knowledge documents.
	// Required: true
	Query *string `json:"query"`

	// Documents matching to the autocomplete suggestions query.
	// Read Only: true
	Results []*KnowledgeDocumentSuggestionResult `json:"results"`
}

// Validate validates this knowledge document suggestion
func (m *KnowledgeDocumentSuggestion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeDocumentSuggestion) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocumentSuggestion) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeDocumentSuggestion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeDocumentSuggestion) UnmarshalBinary(b []byte) error {
	var res KnowledgeDocumentSuggestion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
