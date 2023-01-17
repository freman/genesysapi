// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SpeechTextAnalyticsSettingsResponse speech text analytics settings response
//
// swagger:model SpeechTextAnalyticsSettingsResponse
type SpeechTextAnalyticsSettingsResponse struct {

	// Setting to choose name for the default program for topic detection
	DefaultProgram *AddressableEntityRef `json:"defaultProgram,omitempty"`

	// Setting to choose expected dialects
	ExpectedDialects []string `json:"expectedDialects"`
}

// Validate validates this speech text analytics settings response
func (m *SpeechTextAnalyticsSettingsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultProgram(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpeechTextAnalyticsSettingsResponse) validateDefaultProgram(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultProgram) { // not required
		return nil
	}

	if m.DefaultProgram != nil {
		if err := m.DefaultProgram.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultProgram")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultProgram")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this speech text analytics settings response based on the context it is used
func (m *SpeechTextAnalyticsSettingsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultProgram(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpeechTextAnalyticsSettingsResponse) contextValidateDefaultProgram(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultProgram != nil {
		if err := m.DefaultProgram.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultProgram")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultProgram")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpeechTextAnalyticsSettingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpeechTextAnalyticsSettingsResponse) UnmarshalBinary(b []byte) error {
	var res SpeechTextAnalyticsSettingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
