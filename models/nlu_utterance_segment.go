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

// NluUtteranceSegment nlu utterance segment
//
// swagger:model NluUtteranceSegment
type NluUtteranceSegment struct {

	// The entity annotation of the segment.
	Entity *NamedEntityAnnotation `json:"entity,omitempty"`

	// The text of the segment.
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this nlu utterance segment
func (m *NluUtteranceSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NluUtteranceSegment) validateEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *NluUtteranceSegment) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nlu utterance segment based on the context it is used
func (m *NluUtteranceSegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NluUtteranceSegment) contextValidateEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.Entity != nil {
		if err := m.Entity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NluUtteranceSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NluUtteranceSegment) UnmarshalBinary(b []byte) error {
	var res NluUtteranceSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
