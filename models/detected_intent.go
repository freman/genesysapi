// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DetectedIntent detected intent
//
// swagger:model DetectedIntent
type DetectedIntent struct {

	// The collection of named entities detected.
	// Read Only: true
	Entities []*DetectedNamedEntity `json:"entities"`

	// The name of the detected intent.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The probability of the detected intent.
	// Read Only: true
	Probability float64 `json:"probability,omitempty"`
}

// Validate validates this detected intent
func (m *DetectedIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetectedIntent) validateEntities(formats strfmt.Registry) error {
	if swag.IsZero(m.Entities) { // not required
		return nil
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this detected intent based on the context it is used
func (m *DetectedIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProbability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetectedIntent) contextValidateEntities(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "entities", "body", []*DetectedNamedEntity(m.Entities)); err != nil {
		return err
	}

	for i := 0; i < len(m.Entities); i++ {

		if m.Entities[i] != nil {
			if err := m.Entities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetectedIntent) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *DetectedIntent) contextValidateProbability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "probability", "body", float64(m.Probability)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetectedIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetectedIntent) UnmarshalBinary(b []byte) error {
	var res DetectedIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
