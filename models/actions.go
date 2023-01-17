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

// Actions actions
//
// swagger:model Actions
type Actions struct {

	// skills to remove
	// Unique: true
	SkillsToRemove []*SkillsToRemove `json:"skillsToRemove"`
}

// Validate validates this actions
func (m *Actions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSkillsToRemove(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Actions) validateSkillsToRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.SkillsToRemove) { // not required
		return nil
	}

	if err := validate.UniqueItems("skillsToRemove", "body", m.SkillsToRemove); err != nil {
		return err
	}

	for i := 0; i < len(m.SkillsToRemove); i++ {
		if swag.IsZero(m.SkillsToRemove[i]) { // not required
			continue
		}

		if m.SkillsToRemove[i] != nil {
			if err := m.SkillsToRemove[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skillsToRemove" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("skillsToRemove" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this actions based on the context it is used
func (m *Actions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSkillsToRemove(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Actions) contextValidateSkillsToRemove(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SkillsToRemove); i++ {

		if m.SkillsToRemove[i] != nil {
			if err := m.SkillsToRemove[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skillsToRemove" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("skillsToRemove" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Actions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Actions) UnmarshalBinary(b []byte) error {
	var res Actions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
