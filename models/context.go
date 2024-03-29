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

// Context context
//
// swagger:model Context
type Context struct {

	// A list of one or more patterns to match.
	// Required: true
	Patterns []*ContextPattern `json:"patterns"`
}

// Validate validates this context
func (m *Context) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatterns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Context) validatePatterns(formats strfmt.Registry) error {

	if err := validate.Required("patterns", "body", m.Patterns); err != nil {
		return err
	}

	for i := 0; i < len(m.Patterns); i++ {
		if swag.IsZero(m.Patterns[i]) { // not required
			continue
		}

		if m.Patterns[i] != nil {
			if err := m.Patterns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patterns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this context based on the context it is used
func (m *Context) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePatterns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Context) contextValidatePatterns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Patterns); i++ {

		if m.Patterns[i] != nil {
			if err := m.Patterns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patterns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Context) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Context) UnmarshalBinary(b []byte) error {
	var res Context
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
