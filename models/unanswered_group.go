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

// UnansweredGroup unanswered group
//
// swagger:model UnansweredGroup
type UnansweredGroup struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Knowledge base unanswered group label
	Label string `json:"label,omitempty"`

	// Represents a list of phrase groups inside an unanswered group
	PhraseGroups []*UnansweredPhraseGroup `json:"phraseGroups"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Statistics object containing the various hit counts for an unanswered group
	Statistics *KnowledgeGroupStatistics `json:"statistics,omitempty"`

	// Represents a list of documents that may be linked to an unanswered group
	SuggestedDocuments []*UnansweredGroupSuggestedDocument `json:"suggestedDocuments"`
}

// Validate validates this unanswered group
func (m *UnansweredGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhraseGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuggestedDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnansweredGroup) validatePhraseGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.PhraseGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.PhraseGroups); i++ {
		if swag.IsZero(m.PhraseGroups[i]) { // not required
			continue
		}

		if m.PhraseGroups[i] != nil {
			if err := m.PhraseGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phraseGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UnansweredGroup) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UnansweredGroup) validateStatistics(formats strfmt.Registry) error {

	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	if m.Statistics != nil {
		if err := m.Statistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

func (m *UnansweredGroup) validateSuggestedDocuments(formats strfmt.Registry) error {

	if swag.IsZero(m.SuggestedDocuments) { // not required
		return nil
	}

	for i := 0; i < len(m.SuggestedDocuments); i++ {
		if swag.IsZero(m.SuggestedDocuments[i]) { // not required
			continue
		}

		if m.SuggestedDocuments[i] != nil {
			if err := m.SuggestedDocuments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suggestedDocuments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnansweredGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnansweredGroup) UnmarshalBinary(b []byte) error {
	var res UnansweredGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}