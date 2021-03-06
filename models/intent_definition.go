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

// IntentDefinition intent definition
//
// swagger:model IntentDefinition
type IntentDefinition struct {

	// The bindings for the named entity types used in this intent.
	// Required: true
	EntityTypeBindings []*NamedEntityTypeBinding `json:"entityTypeBindings"`

	// The name of the intent.
	// Required: true
	Name *string `json:"name"`

	// The utterances that act as training phrases for the intent.
	// Required: true
	Utterances []*NluUtterance `json:"utterances"`
}

// Validate validates this intent definition
func (m *IntentDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityTypeBindings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUtterances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntentDefinition) validateEntityTypeBindings(formats strfmt.Registry) error {

	if err := validate.Required("entityTypeBindings", "body", m.EntityTypeBindings); err != nil {
		return err
	}

	for i := 0; i < len(m.EntityTypeBindings); i++ {
		if swag.IsZero(m.EntityTypeBindings[i]) { // not required
			continue
		}

		if m.EntityTypeBindings[i] != nil {
			if err := m.EntityTypeBindings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entityTypeBindings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IntentDefinition) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IntentDefinition) validateUtterances(formats strfmt.Registry) error {

	if err := validate.Required("utterances", "body", m.Utterances); err != nil {
		return err
	}

	for i := 0; i < len(m.Utterances); i++ {
		if swag.IsZero(m.Utterances[i]) { // not required
			continue
		}

		if m.Utterances[i] != nil {
			if err := m.Utterances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("utterances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntentDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntentDefinition) UnmarshalBinary(b []byte) error {
	var res IntentDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
