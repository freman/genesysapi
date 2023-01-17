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

// Section section
//
// swagger:model Section
type Section struct {

	// field list
	FieldList []*FieldList `json:"fieldList"`

	// instruction text
	InstructionText string `json:"instructionText,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this section
func (m *Section) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Section) validateFieldList(formats strfmt.Registry) error {
	if swag.IsZero(m.FieldList) { // not required
		return nil
	}

	for i := 0; i < len(m.FieldList); i++ {
		if swag.IsZero(m.FieldList[i]) { // not required
			continue
		}

		if m.FieldList[i] != nil {
			if err := m.FieldList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fieldList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fieldList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this section based on the context it is used
func (m *Section) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFieldList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Section) contextValidateFieldList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FieldList); i++ {

		if m.FieldList[i] != nil {
			if err := m.FieldList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fieldList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fieldList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Section) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Section) UnmarshalBinary(b []byte) error {
	var res Section
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
