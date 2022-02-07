// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Knowledge knowledge
//
// swagger:model Knowledge
type Knowledge struct {

	// whether or not knowledge base is enabled
	Enabled bool `json:"enabled"`

	// The knowledge base for messenger
	KnowledgeBase *AddressableEntityRef `json:"knowledgeBase,omitempty"`
}

// Validate validates this knowledge
func (m *Knowledge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKnowledgeBase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Knowledge) validateKnowledgeBase(formats strfmt.Registry) error {

	if swag.IsZero(m.KnowledgeBase) { // not required
		return nil
	}

	if m.KnowledgeBase != nil {
		if err := m.KnowledgeBase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("knowledgeBase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Knowledge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Knowledge) UnmarshalBinary(b []byte) error {
	var res Knowledge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}