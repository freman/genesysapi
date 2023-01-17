// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UnansweredPhrase unanswered phrase
//
// swagger:model UnansweredPhrase
type UnansweredPhrase struct {

	// Id of an unanswered phrase
	ID string `json:"id,omitempty"`

	// Phrase text of an unanswered phrase
	Text string `json:"text,omitempty"`

	// Hit count of an unlinked phrase
	UnlinkedPhraseHitCount int32 `json:"unlinkedPhraseHitCount,omitempty"`
}

// Validate validates this unanswered phrase
func (m *UnansweredPhrase) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unanswered phrase based on context it is used
func (m *UnansweredPhrase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UnansweredPhrase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnansweredPhrase) UnmarshalBinary(b []byte) error {
	var res UnansweredPhrase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
