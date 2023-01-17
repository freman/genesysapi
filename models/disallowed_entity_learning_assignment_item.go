// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DisallowedEntityLearningAssignmentItem disallowed entity learning assignment item
//
// swagger:model DisallowedEntityLearningAssignmentItem
type DisallowedEntityLearningAssignmentItem struct {

	// The entity that was disallowed
	Entity *LearningAssignmentItem `json:"entity,omitempty"`

	// The error code associated with this disallowed entity
	ErrorCode string `json:"errorCode,omitempty"`
}

// Validate validates this disallowed entity learning assignment item
func (m *DisallowedEntityLearningAssignmentItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DisallowedEntityLearningAssignmentItem) validateEntity(formats strfmt.Registry) error {
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

// ContextValidate validate this disallowed entity learning assignment item based on the context it is used
func (m *DisallowedEntityLearningAssignmentItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DisallowedEntityLearningAssignmentItem) contextValidateEntity(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DisallowedEntityLearningAssignmentItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisallowedEntityLearningAssignmentItem) UnmarshalBinary(b []byte) error {
	var res DisallowedEntityLearningAssignmentItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
