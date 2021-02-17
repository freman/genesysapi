// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SpeechTextAnalyticsSettingsRequest speech text analytics settings request
//
// swagger:model SpeechTextAnalyticsSettingsRequest
type SpeechTextAnalyticsSettingsRequest struct {

	// Setting to choose name for the default program for topic detection
	DefaultProgramID string `json:"defaultProgramId,omitempty"`

	// Setting to choose expected dialects
	ExpectedDialects []string `json:"expectedDialects"`
}

// Validate validates this speech text analytics settings request
func (m *SpeechTextAnalyticsSettingsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpeechTextAnalyticsSettingsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpeechTextAnalyticsSettingsRequest) UnmarshalBinary(b []byte) error {
	var res SpeechTextAnalyticsSettingsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
