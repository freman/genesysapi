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

// ButtonComponent Structured template button object.
//
// swagger:model ButtonComponent
type ButtonComponent struct {

	// The button actions (Deprecated).
	Actions *ContentActions `json:"actions,omitempty"`

	// Text to show inside the button.
	Title string `json:"title,omitempty"`
}

// Validate validates this button component
func (m *ButtonComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ButtonComponent) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {
		if err := m.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this button component based on the context it is used
func (m *ButtonComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ButtonComponent) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	if m.Actions != nil {
		if err := m.Actions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ButtonComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ButtonComponent) UnmarshalBinary(b []byte) error {
	var res ButtonComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
