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

// NluConfusionMatrixRow nlu confusion matrix row
//
// swagger:model NluConfusionMatrixRow
type NluConfusionMatrixRow struct {

	// The columns of confusion matrix for the intent
	// Required: true
	Columns []*NluConfusionMatrixColumn `json:"columns"`

	// The name of the intent for the row.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this nlu confusion matrix row
func (m *NluConfusionMatrixRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NluConfusionMatrixRow) validateColumns(formats strfmt.Registry) error {

	if err := validate.Required("columns", "body", m.Columns); err != nil {
		return err
	}

	for i := 0; i < len(m.Columns); i++ {
		if swag.IsZero(m.Columns[i]) { // not required
			continue
		}

		if m.Columns[i] != nil {
			if err := m.Columns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NluConfusionMatrixRow) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nlu confusion matrix row based on the context it is used
func (m *NluConfusionMatrixRow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NluConfusionMatrixRow) contextValidateColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Columns); i++ {

		if m.Columns[i] != nil {
			if err := m.Columns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NluConfusionMatrixRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NluConfusionMatrixRow) UnmarshalBinary(b []byte) error {
	var res NluConfusionMatrixRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
