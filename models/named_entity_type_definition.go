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

// NamedEntityTypeDefinition named entity type definition
//
// swagger:model NamedEntityTypeDefinition
type NamedEntityTypeDefinition struct {

	// Description of the of the named entity type.
	Description string `json:"description,omitempty"`

	// The mechanism enabling detection of the named entity type.
	// Required: true
	Mechanism *NamedEntityTypeMechanism `json:"mechanism"`

	// The name of the entity type.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this named entity type definition
func (m *NamedEntityTypeDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMechanism(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedEntityTypeDefinition) validateMechanism(formats strfmt.Registry) error {

	if err := validate.Required("mechanism", "body", m.Mechanism); err != nil {
		return err
	}

	if m.Mechanism != nil {
		if err := m.Mechanism.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mechanism")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mechanism")
			}
			return err
		}
	}

	return nil
}

func (m *NamedEntityTypeDefinition) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this named entity type definition based on the context it is used
func (m *NamedEntityTypeDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMechanism(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedEntityTypeDefinition) contextValidateMechanism(ctx context.Context, formats strfmt.Registry) error {

	if m.Mechanism != nil {
		if err := m.Mechanism.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mechanism")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mechanism")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamedEntityTypeDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamedEntityTypeDefinition) UnmarshalBinary(b []byte) error {
	var res NamedEntityTypeDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
