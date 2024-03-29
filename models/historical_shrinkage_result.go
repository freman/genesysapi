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

// HistoricalShrinkageResult historical shrinkage result
//
// swagger:model HistoricalShrinkageResult
type HistoricalShrinkageResult struct {

	// Aggregated shrinkage data for all the activity categories
	AggregatedShrinkage *HistoricalShrinkageAggregateResponse `json:"aggregatedShrinkage,omitempty"`

	// List of all business units of all the agents in response
	BusinessUnitIds []string `json:"businessUnitIds"`

	// End of the date range that was queried, in ISO-8601 format. If it was not set, end date will be set to the queried time
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// Shrinkage for activity categories
	ShrinkageForActivityCategories []*HistoricalShrinkageActivityCategoryResponse `json:"shrinkageForActivityCategories"`

	// Beginning of the date range that was queried, in ISO-8601 format
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Total duration in seconds for which agents in the management unit are actually logged-in
	TotalLoggedInDurationSeconds int32 `json:"totalLoggedInDurationSeconds,omitempty"`

	// Total duration in seconds for which agents in the management unit are scheduled
	TotalScheduledDurationSeconds int32 `json:"totalScheduledDurationSeconds,omitempty"`
}

// Validate validates this historical shrinkage result
func (m *HistoricalShrinkageResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregatedShrinkage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShrinkageForActivityCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoricalShrinkageResult) validateAggregatedShrinkage(formats strfmt.Registry) error {
	if swag.IsZero(m.AggregatedShrinkage) { // not required
		return nil
	}

	if m.AggregatedShrinkage != nil {
		if err := m.AggregatedShrinkage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregatedShrinkage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregatedShrinkage")
			}
			return err
		}
	}

	return nil
}

func (m *HistoricalShrinkageResult) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalShrinkageResult) validateShrinkageForActivityCategories(formats strfmt.Registry) error {
	if swag.IsZero(m.ShrinkageForActivityCategories) { // not required
		return nil
	}

	for i := 0; i < len(m.ShrinkageForActivityCategories); i++ {
		if swag.IsZero(m.ShrinkageForActivityCategories[i]) { // not required
			continue
		}

		if m.ShrinkageForActivityCategories[i] != nil {
			if err := m.ShrinkageForActivityCategories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shrinkageForActivityCategories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shrinkageForActivityCategories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HistoricalShrinkageResult) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this historical shrinkage result based on the context it is used
func (m *HistoricalShrinkageResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregatedShrinkage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShrinkageForActivityCategories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoricalShrinkageResult) contextValidateAggregatedShrinkage(ctx context.Context, formats strfmt.Registry) error {

	if m.AggregatedShrinkage != nil {
		if err := m.AggregatedShrinkage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregatedShrinkage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregatedShrinkage")
			}
			return err
		}
	}

	return nil
}

func (m *HistoricalShrinkageResult) contextValidateShrinkageForActivityCategories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ShrinkageForActivityCategories); i++ {

		if m.ShrinkageForActivityCategories[i] != nil {
			if err := m.ShrinkageForActivityCategories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shrinkageForActivityCategories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shrinkageForActivityCategories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoricalShrinkageResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoricalShrinkageResult) UnmarshalBinary(b []byte) error {
	var res HistoricalShrinkageResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
