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

// ForecastPlanningGroupData forecast planning group data
//
// swagger:model ForecastPlanningGroupData
type ForecastPlanningGroupData struct {

	// Forecast average handle time per 15 minute interval in seconds
	// Required: true
	AverageHandleTimeSecondsPerInterval []float64 `json:"averageHandleTimeSecondsPerInterval"`

	// Forecast offered counts per 15 minute interval for this week of the forecast
	// Required: true
	OfferedPerInterval []float64 `json:"offeredPerInterval"`

	// The ID of the planning group to which this data applies. Note this is a snapshot of the planning group at the time of forecast creation and may not correspond to the current configuration
	// Required: true
	PlanningGroupID *string `json:"planningGroupId"`
}

// Validate validates this forecast planning group data
func (m *ForecastPlanningGroupData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAverageHandleTimeSecondsPerInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferedPerInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanningGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ForecastPlanningGroupData) validateAverageHandleTimeSecondsPerInterval(formats strfmt.Registry) error {

	if err := validate.Required("averageHandleTimeSecondsPerInterval", "body", m.AverageHandleTimeSecondsPerInterval); err != nil {
		return err
	}

	return nil
}

func (m *ForecastPlanningGroupData) validateOfferedPerInterval(formats strfmt.Registry) error {

	if err := validate.Required("offeredPerInterval", "body", m.OfferedPerInterval); err != nil {
		return err
	}

	return nil
}

func (m *ForecastPlanningGroupData) validatePlanningGroupID(formats strfmt.Registry) error {

	if err := validate.Required("planningGroupId", "body", m.PlanningGroupID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this forecast planning group data based on context it is used
func (m *ForecastPlanningGroupData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ForecastPlanningGroupData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ForecastPlanningGroupData) UnmarshalBinary(b []byte) error {
	var res ForecastPlanningGroupData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
