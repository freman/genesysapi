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

// PhraseAssociations phrase associations
//
// swagger:model PhraseAssociations
type PhraseAssociations struct {

	// Id of the document to be linked
	// Required: true
	DocumentID *string `json:"documentId"`

	// Id of the phrase to be linked
	// Required: true
	PhraseID *string `json:"phraseId"`
}

// Validate validates this phrase associations
func (m *PhraseAssociations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhraseID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhraseAssociations) validateDocumentID(formats strfmt.Registry) error {

	if err := validate.Required("documentId", "body", m.DocumentID); err != nil {
		return err
	}

	return nil
}

func (m *PhraseAssociations) validatePhraseID(formats strfmt.Registry) error {

	if err := validate.Required("phraseId", "body", m.PhraseID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PhraseAssociations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhraseAssociations) UnmarshalBinary(b []byte) error {
	var res PhraseAssociations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
