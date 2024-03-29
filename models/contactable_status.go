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

// ContactableStatus contactable status
//
// swagger:model ContactableStatus
type ContactableStatus struct {

	// A map of individual contact method columns to whether the individual column is contactable for the associated media type.
	ColumnStatus map[string]ColumnStatus `json:"columnStatus,omitempty"`

	// Indicates whether or not the entire contact is contactable for the associated media type.
	Contactable bool `json:"contactable"`
}

// Validate validates this contactable status
func (m *ContactableStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumnStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactableStatus) validateColumnStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ColumnStatus) { // not required
		return nil
	}

	for k := range m.ColumnStatus {

		if err := validate.Required("columnStatus"+"."+k, "body", m.ColumnStatus[k]); err != nil {
			return err
		}
		if val, ok := m.ColumnStatus[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columnStatus" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columnStatus" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contactable status based on the context it is used
func (m *ContactableStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColumnStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactableStatus) contextValidateColumnStatus(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ColumnStatus {

		if val, ok := m.ColumnStatus[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactableStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactableStatus) UnmarshalBinary(b []byte) error {
	var res ContactableStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
