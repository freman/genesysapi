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

// BuForecastResult bu forecast result
//
// swagger:model BuForecastResult
type BuForecastResult struct {

	// The forecast data broken up by planning group
	PlanningGroups []*ForecastPlanningGroupData `json:"planningGroups"`

	// The reference start date for interval-based data for this forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ReferenceStartDate strfmt.DateTime `json:"referenceStartDate,omitempty"`

	// The number of weeks in this forecast
	WeekCount int32 `json:"weekCount,omitempty"`

	// The week number represented by this response
	WeekNumber int32 `json:"weekNumber,omitempty"`
}

// Validate validates this bu forecast result
func (m *BuForecastResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlanningGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuForecastResult) validatePlanningGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.PlanningGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.PlanningGroups); i++ {
		if swag.IsZero(m.PlanningGroups[i]) { // not required
			continue
		}

		if m.PlanningGroups[i] != nil {
			if err := m.PlanningGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("planningGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("planningGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuForecastResult) validateReferenceStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferenceStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("referenceStartDate", "body", "date-time", m.ReferenceStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bu forecast result based on the context it is used
func (m *BuForecastResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlanningGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuForecastResult) contextValidatePlanningGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlanningGroups); i++ {

		if m.PlanningGroups[i] != nil {
			if err := m.PlanningGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("planningGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("planningGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuForecastResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuForecastResult) UnmarshalBinary(b []byte) error {
	var res BuForecastResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
