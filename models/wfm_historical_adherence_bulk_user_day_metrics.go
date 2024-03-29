// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WfmHistoricalAdherenceBulkUserDayMetrics wfm historical adherence bulk user day metrics
//
// swagger:model WfmHistoricalAdherenceBulkUserDayMetrics
type WfmHistoricalAdherenceBulkUserDayMetrics struct {

	// Total duration in seconds for all actually worked activities
	ActualLengthSeconds int32 `json:"actualLengthSeconds,omitempty"`

	// Total adherence percentage for this user, in the scale of 0 - 100
	AdherencePercentage float64 `json:"adherencePercentage,omitempty"`

	// Duration of schedule in seconds included for adherence percentage calculation
	AdherenceScheduleSeconds int32 `json:"adherenceScheduleSeconds,omitempty"`

	// Total actually worked duration in seconds for OnQueue activities
	ConformanceActualSeconds int32 `json:"conformanceActualSeconds,omitempty"`

	// Total conformance percentage for this user, in the scale of 0 - 100. Conformance percentage can be greater than 100 when the actual on queue time is greater than the scheduled on queue time for the same period.
	ConformancePercentage float64 `json:"conformancePercentage,omitempty"`

	// Total scheduled duration in seconds for OnQueue activities
	ConformanceScheduleSeconds int32 `json:"conformanceScheduleSeconds,omitempty"`

	// Start of day offset in seconds relative to query start time
	DayStartOffsetSeconds int32 `json:"dayStartOffsetSeconds,omitempty"`

	// Total number of adherence exceptions for this user
	ExceptionCount int32 `json:"exceptionCount,omitempty"`

	// Total duration in seconds of adherence exceptions for this user
	ExceptionDurationSeconds int32 `json:"exceptionDurationSeconds,omitempty"`

	// The impact duration in seconds of current adherence state for this user
	ImpactSeconds int32 `json:"impactSeconds,omitempty"`

	// Total duration in seconds for all scheduled activities
	ScheduleLengthSeconds int32 `json:"scheduleLengthSeconds,omitempty"`
}

// Validate validates this wfm historical adherence bulk user day metrics
func (m *WfmHistoricalAdherenceBulkUserDayMetrics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this wfm historical adherence bulk user day metrics based on context it is used
func (m *WfmHistoricalAdherenceBulkUserDayMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WfmHistoricalAdherenceBulkUserDayMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmHistoricalAdherenceBulkUserDayMetrics) UnmarshalBinary(b []byte) error {
	var res WfmHistoricalAdherenceBulkUserDayMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
