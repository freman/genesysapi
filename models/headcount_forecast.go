// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HeadcountForecast Headcount interval information for schedule
//
// swagger:model HeadcountForecast
type HeadcountForecast struct {

	// Headcount information with shrinkage
	// Required: true
	Required []*HeadcountInterval `json:"required"`

	// Headcount information without shrinkage
	// Required: true
	RequiredWithoutShrinkage []*HeadcountInterval `json:"requiredWithoutShrinkage"`
}

// Validate validates this headcount forecast
func (m *HeadcountForecast) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredWithoutShrinkage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeadcountForecast) validateRequired(formats strfmt.Registry) error {

	if err := validate.Required("required", "body", m.Required); err != nil {
		return err
	}

	for i := 0; i < len(m.Required); i++ {
		if swag.IsZero(m.Required[i]) { // not required
			continue
		}

		if m.Required[i] != nil {
			if err := m.Required[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HeadcountForecast) validateRequiredWithoutShrinkage(formats strfmt.Registry) error {

	if err := validate.Required("requiredWithoutShrinkage", "body", m.RequiredWithoutShrinkage); err != nil {
		return err
	}

	for i := 0; i < len(m.RequiredWithoutShrinkage); i++ {
		if swag.IsZero(m.RequiredWithoutShrinkage[i]) { // not required
			continue
		}

		if m.RequiredWithoutShrinkage[i] != nil {
			if err := m.RequiredWithoutShrinkage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requiredWithoutShrinkage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HeadcountForecast) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeadcountForecast) UnmarshalBinary(b []byte) error {
	var res HeadcountForecast
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
