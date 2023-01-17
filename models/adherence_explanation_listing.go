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
)

// AdherenceExplanationListing adherence explanation listing
//
// swagger:model AdherenceExplanationListing
type AdherenceExplanationListing struct {

	// entities
	Entities []*AdherenceExplanationResponse `json:"entities"`
}

// Validate validates this adherence explanation listing
func (m *AdherenceExplanationListing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdherenceExplanationListing) validateEntities(formats strfmt.Registry) error {
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

// ContextValidate validate this adherence explanation listing based on the context it is used
func (m *AdherenceExplanationListing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdherenceExplanationListing) contextValidateEntities(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *AdherenceExplanationListing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdherenceExplanationListing) UnmarshalBinary(b []byte) error {
	var res AdherenceExplanationListing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}