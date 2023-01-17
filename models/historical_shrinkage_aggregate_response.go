// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HistoricalShrinkageAggregateResponse historical shrinkage aggregate response
//
// swagger:model HistoricalShrinkageAggregateResponse
type HistoricalShrinkageAggregateResponse struct {

	// Aggregated actual value in percent from 0.0 to 100.0 for scheduled activities
	ActualShrinkagePercent float64 `json:"actualShrinkagePercent,omitempty"`

	// Aggregated actual value in seconds for scheduled activities
	ActualShrinkageSeconds int32 `json:"actualShrinkageSeconds,omitempty"`

	// Aggregated shrinkage value in seconds for paid activities
	PaidShrinkageSeconds int32 `json:"paidShrinkageSeconds,omitempty"`

	// Aggregated shrinkage value in seconds for planned activities
	PlannedShrinkageSeconds int32 `json:"plannedShrinkageSeconds,omitempty"`

	// Aggregated shrinkage value in percent from 0.0 to 100.0 for scheduled activities
	ScheduledShrinkagePercent float64 `json:"scheduledShrinkagePercent,omitempty"`

	// Aggregated shrinkage value in seconds for scheduled activities
	ScheduledShrinkageSeconds int32 `json:"scheduledShrinkageSeconds,omitempty"`

	// Aggregated shrinkage value in seconds for unpaid activities
	UnpaidShrinkageSeconds int32 `json:"unpaidShrinkageSeconds,omitempty"`

	// Aggregated shrinkage value in seconds for unplanned activities
	UnplannedShrinkageSeconds int32 `json:"unplannedShrinkageSeconds,omitempty"`
}

// Validate validates this historical shrinkage aggregate response
func (m *HistoricalShrinkageAggregateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this historical shrinkage aggregate response based on context it is used
func (m *HistoricalShrinkageAggregateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoricalShrinkageAggregateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoricalShrinkageAggregateResponse) UnmarshalBinary(b []byte) error {
	var res HistoricalShrinkageAggregateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}