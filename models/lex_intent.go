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

// LexIntent lex intent
//
// swagger:model LexIntent
type LexIntent struct {

	// A description of the intent
	Description string `json:"description,omitempty"`

	// The intent name
	// Required: true
	Name *string `json:"name"`

	// An object mapping slot names to Slot objects
	// Required: true
	Slots map[string]LexSlot `json:"slots"`

	// The intent version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this lex intent
func (m *LexIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LexIntent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *LexIntent) validateSlots(formats strfmt.Registry) error {

	if err := validate.Required("slots", "body", m.Slots); err != nil {
		return err
	}

	for k := range m.Slots {

		if err := validate.Required("slots"+"."+k, "body", m.Slots[k]); err != nil {
			return err
		}
		if val, ok := m.Slots[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slots" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("slots" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *LexIntent) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this lex intent based on the context it is used
func (m *LexIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSlots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LexIntent) contextValidateSlots(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("slots", "body", m.Slots); err != nil {
		return err
	}

	for k := range m.Slots {

		if val, ok := m.Slots[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LexIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LexIntent) UnmarshalBinary(b []byte) error {
	var res LexIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
