// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkPlanPatternRequest work plan pattern request
//
// swagger:model WorkPlanPatternRequest
type WorkPlanPatternRequest struct {

	// List of work plan IDs in order of rotation on a weekly basis. Values in the list cannot be null or empty
	// Required: true
	WorkPlanIds []string `json:"workPlanIds"`
}

// Validate validates this work plan pattern request
func (m *WorkPlanPatternRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkPlanIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkPlanPatternRequest) validateWorkPlanIds(formats strfmt.Registry) error {

	if err := validate.Required("workPlanIds", "body", m.WorkPlanIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkPlanPatternRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkPlanPatternRequest) UnmarshalBinary(b []byte) error {
	var res WorkPlanPatternRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
