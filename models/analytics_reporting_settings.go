// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnalyticsReportingSettings analytics reporting settings
//
// swagger:model AnalyticsReportingSettings
type AnalyticsReportingSettings struct {

	// Indication of whether or not personal data is masked in data export and the Analytics/Reporting UI
	PiiMaskingEnabled bool `json:"piiMaskingEnabled"`
}

// Validate validates this analytics reporting settings
func (m *AnalyticsReportingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsReportingSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsReportingSettings) UnmarshalBinary(b []byte) error {
	var res AnalyticsReportingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
