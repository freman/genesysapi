// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShortTermForecastingSettings Short Term Forecasting Settings
//
// swagger:model ShortTermForecastingSettings
type ShortTermForecastingSettings struct {

	// The number of weeks to consider by default when generating a volume forecast
	DefaultHistoryWeeks int32 `json:"defaultHistoryWeeks,omitempty"`
}

// Validate validates this short term forecasting settings
func (m *ShortTermForecastingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShortTermForecastingSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShortTermForecastingSettings) UnmarshalBinary(b []byte) error {
	var res ShortTermForecastingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}