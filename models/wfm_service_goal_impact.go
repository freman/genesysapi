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

// WfmServiceGoalImpact wfm service goal impact
//
// swagger:model WfmServiceGoalImpact
type WfmServiceGoalImpact struct {

	// The maximum allowed percent decrease from the configured goal
	// Required: true
	DecreaseByPercent *float64 `json:"decreaseByPercent"`

	// The maximum allowed percent increase from the configured goal
	// Required: true
	IncreaseByPercent *float64 `json:"increaseByPercent"`
}

// Validate validates this wfm service goal impact
func (m *WfmServiceGoalImpact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecreaseByPercent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncreaseByPercent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmServiceGoalImpact) validateDecreaseByPercent(formats strfmt.Registry) error {

	if err := validate.Required("decreaseByPercent", "body", m.DecreaseByPercent); err != nil {
		return err
	}

	return nil
}

func (m *WfmServiceGoalImpact) validateIncreaseByPercent(formats strfmt.Registry) error {

	if err := validate.Required("increaseByPercent", "body", m.IncreaseByPercent); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wfm service goal impact based on context it is used
func (m *WfmServiceGoalImpact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WfmServiceGoalImpact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmServiceGoalImpact) UnmarshalBinary(b []byte) error {
	var res WfmServiceGoalImpact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
