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

// CreateBusinessUnitRequest create business unit request
//
// swagger:model CreateBusinessUnitRequest
type CreateBusinessUnitRequest struct {

	// The ID of the division to which the business unit should be added
	// Required: true
	DivisionID *string `json:"divisionId"`

	// The name of the business unit
	// Required: true
	Name *string `json:"name"`

	// Configuration for the business unit
	// Required: true
	Settings *CreateBusinessUnitSettingsRequest `json:"settings"`
}

// Validate validates this create business unit request
func (m *CreateBusinessUnitRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivisionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBusinessUnitRequest) validateDivisionID(formats strfmt.Registry) error {

	if err := validate.Required("divisionId", "body", m.DivisionID); err != nil {
		return err
	}

	return nil
}

func (m *CreateBusinessUnitRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateBusinessUnitRequest) validateSettings(formats strfmt.Registry) error {

	if err := validate.Required("settings", "body", m.Settings); err != nil {
		return err
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create business unit request based on the context it is used
func (m *CreateBusinessUnitRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBusinessUnitRequest) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {
		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateBusinessUnitRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBusinessUnitRequest) UnmarshalBinary(b []byte) error {
	var res CreateBusinessUnitRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
