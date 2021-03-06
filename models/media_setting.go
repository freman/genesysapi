// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MediaSetting media setting
//
// swagger:model MediaSetting
type MediaSetting struct {

	// alerting timeout seconds
	AlertingTimeoutSeconds int32 `json:"alertingTimeoutSeconds,omitempty"`

	// service level
	ServiceLevel *ServiceLevel `json:"serviceLevel,omitempty"`
}

// Validate validates this media setting
func (m *MediaSetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaSetting) validateServiceLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceLevel) { // not required
		return nil
	}

	if m.ServiceLevel != nil {
		if err := m.ServiceLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceLevel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediaSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaSetting) UnmarshalBinary(b []byte) error {
	var res MediaSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
