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

// Adjacents adjacents
//
// swagger:model Adjacents
type Adjacents struct {

	// direct reports
	DirectReports []*User `json:"directReports"`

	// siblings
	Siblings []*User `json:"siblings"`

	// superiors
	Superiors []*User `json:"superiors"`
}

// Validate validates this adjacents
func (m *Adjacents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiblings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuperiors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Adjacents) validateDirectReports(formats strfmt.Registry) error {
	if swag.IsZero(m.DirectReports) { // not required
		return nil
	}

	for i := 0; i < len(m.DirectReports); i++ {
		if swag.IsZero(m.DirectReports[i]) { // not required
			continue
		}

		if m.DirectReports[i] != nil {
			if err := m.DirectReports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("directReports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("directReports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Adjacents) validateSiblings(formats strfmt.Registry) error {
	if swag.IsZero(m.Siblings) { // not required
		return nil
	}

	for i := 0; i < len(m.Siblings); i++ {
		if swag.IsZero(m.Siblings[i]) { // not required
			continue
		}

		if m.Siblings[i] != nil {
			if err := m.Siblings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("siblings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("siblings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Adjacents) validateSuperiors(formats strfmt.Registry) error {
	if swag.IsZero(m.Superiors) { // not required
		return nil
	}

	for i := 0; i < len(m.Superiors); i++ {
		if swag.IsZero(m.Superiors[i]) { // not required
			continue
		}

		if m.Superiors[i] != nil {
			if err := m.Superiors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("superiors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("superiors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this adjacents based on the context it is used
func (m *Adjacents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirectReports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSiblings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuperiors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Adjacents) contextValidateDirectReports(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DirectReports); i++ {

		if m.DirectReports[i] != nil {
			if err := m.DirectReports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("directReports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("directReports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Adjacents) contextValidateSiblings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Siblings); i++ {

		if m.Siblings[i] != nil {
			if err := m.Siblings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("siblings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("siblings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Adjacents) contextValidateSuperiors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Superiors); i++ {

		if m.Superiors[i] != nil {
			if err := m.Superiors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("superiors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("superiors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Adjacents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Adjacents) UnmarshalBinary(b []byte) error {
	var res Adjacents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
