// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ValueWrapperPlanningPeriodSettings value wrapper planning period settings
//
// swagger:model ValueWrapperPlanningPeriodSettings
type ValueWrapperPlanningPeriodSettings struct {

	// The value for the associated field
	Value *PlanningPeriodSettings `json:"value,omitempty"`
}

// Validate validates this value wrapper planning period settings
func (m *ValueWrapperPlanningPeriodSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValueWrapperPlanningPeriodSettings) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ValueWrapperPlanningPeriodSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValueWrapperPlanningPeriodSettings) UnmarshalBinary(b []byte) error {
	var res ValueWrapperPlanningPeriodSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
