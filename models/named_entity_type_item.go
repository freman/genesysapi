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

// NamedEntityTypeItem named entity type item
//
// swagger:model NamedEntityTypeItem
type NamedEntityTypeItem struct {

	// Additional Language Synonyms for the given named entity value.
	AdditionalLanguages map[string]AdditionalLanguagesSynonyms `json:"additionalLanguages,omitempty"`

	// Synonyms for the given named entity value.
	Synonyms []string `json:"synonyms"`

	// A value for an named entity type definition.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this named entity type item
func (m *NamedEntityTypeItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalLanguages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedEntityTypeItem) validateAdditionalLanguages(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalLanguages) { // not required
		return nil
	}

	for k := range m.AdditionalLanguages {

		if err := validate.Required("additionalLanguages"+"."+k, "body", m.AdditionalLanguages[k]); err != nil {
			return err
		}
		if val, ok := m.AdditionalLanguages[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalLanguages" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalLanguages" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *NamedEntityTypeItem) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this named entity type item based on the context it is used
func (m *NamedEntityTypeItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalLanguages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedEntityTypeItem) contextValidateAdditionalLanguages(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.AdditionalLanguages {

		if val, ok := m.AdditionalLanguages[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamedEntityTypeItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamedEntityTypeItem) UnmarshalBinary(b []byte) error {
	var res NamedEntityTypeItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
