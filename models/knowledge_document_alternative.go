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

// KnowledgeDocumentAlternative knowledge document alternative
//
// swagger:model KnowledgeDocumentAlternative
type KnowledgeDocumentAlternative struct {

	// Autocomplete enabled for the alternate phrase.
	// Required: true
	Autocomplete *bool `json:"autocomplete"`

	// Alternate phrasing to the document title, having a limit of 500 words.
	// Required: true
	Phrase *string `json:"phrase"`
}

// Validate validates this knowledge document alternative
func (m *KnowledgeDocumentAlternative) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutocomplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhrase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeDocumentAlternative) validateAutocomplete(formats strfmt.Registry) error {

	if err := validate.Required("autocomplete", "body", m.Autocomplete); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocumentAlternative) validatePhrase(formats strfmt.Registry) error {

	if err := validate.Required("phrase", "body", m.Phrase); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this knowledge document alternative based on context it is used
func (m *KnowledgeDocumentAlternative) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeDocumentAlternative) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeDocumentAlternative) UnmarshalBinary(b []byte) error {
	var res KnowledgeDocumentAlternative
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
