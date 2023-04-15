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

// MediaSettings media settings
//
// swagger:model MediaSettings
type MediaSettings struct {

	// alerting timeout seconds
	AlertingTimeoutSeconds int32 `json:"alertingTimeoutSeconds,omitempty"`

	// Indicates if auto-answer is enabled for the given media type or subtype (default is false).  Subtype settings take precedence over media type settings.
	EnableAutoAnswer bool `json:"enableAutoAnswer"`

	// service level
	ServiceLevel *ServiceLevel `json:"serviceLevel,omitempty"`

	// Map of media subtype to media subtype specific settings.
	SubTypeSettings map[string]BaseMediaSettings `json:"subTypeSettings,omitempty"`
}

// Validate validates this media settings
func (m *MediaSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubTypeSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaSettings) validateServiceLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceLevel) { // not required
		return nil
	}

	if m.ServiceLevel != nil {
		if err := m.ServiceLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceLevel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceLevel")
			}
			return err
		}
	}

	return nil
}

func (m *MediaSettings) validateSubTypeSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SubTypeSettings) { // not required
		return nil
	}

	for k := range m.SubTypeSettings {

		if err := validate.Required("subTypeSettings"+"."+k, "body", m.SubTypeSettings[k]); err != nil {
			return err
		}
		if val, ok := m.SubTypeSettings[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subTypeSettings" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subTypeSettings" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this media settings based on the context it is used
func (m *MediaSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubTypeSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaSettings) contextValidateServiceLevel(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceLevel != nil {
		if err := m.ServiceLevel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceLevel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceLevel")
			}
			return err
		}
	}

	return nil
}

func (m *MediaSettings) contextValidateSubTypeSettings(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.SubTypeSettings {

		if val, ok := m.SubTypeSettings[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediaSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaSettings) UnmarshalBinary(b []byte) error {
	var res MediaSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}