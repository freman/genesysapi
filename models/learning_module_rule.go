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

// LearningModuleRule learning module rule
//
// swagger:model LearningModuleRule
type LearningModuleRule struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// If true, rule is active
	// Required: true
	IsActive *bool `json:"isActive"`

	// The parts of a learning module rule
	// Required: true
	Parts []*LearningModuleRuleParts `json:"parts"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this learning module rule
func (m *LearningModuleRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParts(formats); err != nil {
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

func (m *LearningModuleRule) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("isActive", "body", m.IsActive); err != nil {
		return err
	}

	return nil
}

func (m *LearningModuleRule) validateParts(formats strfmt.Registry) error {

	if err := validate.Required("parts", "body", m.Parts); err != nil {
		return err
	}

	for i := 0; i < len(m.Parts); i++ {
		if swag.IsZero(m.Parts[i]) { // not required
			continue
		}

		if m.Parts[i] != nil {
			if err := m.Parts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningModuleRule) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningModuleRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningModuleRule) UnmarshalBinary(b []byte) error {
	var res LearningModuleRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
