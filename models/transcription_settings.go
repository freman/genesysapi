// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TranscriptionSettings transcription settings
//
// swagger:model TranscriptionSettings
type TranscriptionSettings struct {

	// Setting to enable/disable content search
	ContentSearchEnabled bool `json:"contentSearchEnabled"`

	// Boolean flag indicating whether low latency transcription via Notification API is enabled
	LowLatencyTranscriptionEnabled bool `json:"lowLatencyTranscriptionEnabled"`

	// Setting to enable/disable transcription capability
	// Required: true
	// Enum: [Disabled EnabledGlobally EnabledQueueFlow]
	Transcription *string `json:"transcription"`

	// Configure confidence threshold. The possible values are from 1 to 100.
	// Required: true
	TranscriptionConfidenceThreshold *int32 `json:"transcriptionConfidenceThreshold"`
}

// Validate validates this transcription settings
func (m *TranscriptionSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTranscription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTranscriptionConfidenceThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var transcriptionSettingsTypeTranscriptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Disabled","EnabledGlobally","EnabledQueueFlow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptionSettingsTypeTranscriptionPropEnum = append(transcriptionSettingsTypeTranscriptionPropEnum, v)
	}
}

const (

	// TranscriptionSettingsTranscriptionDisabled captures enum value "Disabled"
	TranscriptionSettingsTranscriptionDisabled string = "Disabled"

	// TranscriptionSettingsTranscriptionEnabledGlobally captures enum value "EnabledGlobally"
	TranscriptionSettingsTranscriptionEnabledGlobally string = "EnabledGlobally"

	// TranscriptionSettingsTranscriptionEnabledQueueFlow captures enum value "EnabledQueueFlow"
	TranscriptionSettingsTranscriptionEnabledQueueFlow string = "EnabledQueueFlow"
)

// prop value enum
func (m *TranscriptionSettings) validateTranscriptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptionSettingsTypeTranscriptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptionSettings) validateTranscription(formats strfmt.Registry) error {

	if err := validate.Required("transcription", "body", m.Transcription); err != nil {
		return err
	}

	// value enum
	if err := m.validateTranscriptionEnum("transcription", "body", *m.Transcription); err != nil {
		return err
	}

	return nil
}

func (m *TranscriptionSettings) validateTranscriptionConfidenceThreshold(formats strfmt.Registry) error {

	if err := validate.Required("transcriptionConfidenceThreshold", "body", m.TranscriptionConfidenceThreshold); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transcription settings based on context it is used
func (m *TranscriptionSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TranscriptionSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranscriptionSettings) UnmarshalBinary(b []byte) error {
	var res TranscriptionSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
