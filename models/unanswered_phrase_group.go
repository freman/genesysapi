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

// UnansweredPhraseGroup unanswered phrase group
//
// swagger:model UnansweredPhraseGroup
type UnansweredPhraseGroup struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Knowledge base phrase group label
	Label string `json:"label,omitempty"`

	// List of unanswered phrases in a phrase group
	Phrases []*UnansweredPhrase `json:"phrases"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Unique phrase count of the unlinked phrase group
	UnlinkedPhraseCount int32 `json:"unlinkedPhraseCount,omitempty"`

	// Hit count of the unlinked phrase group
	UnlinkedPhraseHitCount int32 `json:"unlinkedPhraseHitCount,omitempty"`
}

// Validate validates this unanswered phrase group
func (m *UnansweredPhraseGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhrases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnansweredPhraseGroup) validatePhrases(formats strfmt.Registry) error {

	if swag.IsZero(m.Phrases) { // not required
		return nil
	}

	for i := 0; i < len(m.Phrases); i++ {
		if swag.IsZero(m.Phrases[i]) { // not required
			continue
		}

		if m.Phrases[i] != nil {
			if err := m.Phrases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phrases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UnansweredPhraseGroup) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnansweredPhraseGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnansweredPhraseGroup) UnmarshalBinary(b []byte) error {
	var res UnansweredPhraseGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}