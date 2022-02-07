// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessengerSettings Settings concerning messenger
//
// swagger:model MessengerSettings
type MessengerSettings struct {

	// The apps embedded in the messenger
	Apps *MessengerApps `json:"apps,omitempty"`

	// Whether or not messenger is enabled
	Enabled bool `json:"enabled"`

	// The file upload settings for messenger
	FileUpload *FileUploadSettings `json:"fileUpload,omitempty"`

	// The launcher button settings for messenger
	LauncherButton *LauncherButtonSettings `json:"launcherButton,omitempty"`

	// The position settings for messenger
	Position *MessengerPositionSettings `json:"position,omitempty"`

	// The style settings for messenger
	Styles *MessengerStyles `json:"styles,omitempty"`
}

// Validate validates this messenger settings
func (m *MessengerSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileUpload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLauncherButton(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStyles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessengerSettings) validateApps(formats strfmt.Registry) error {

	if swag.IsZero(m.Apps) { // not required
		return nil
	}

	if m.Apps != nil {
		if err := m.Apps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apps")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) validateFileUpload(formats strfmt.Registry) error {

	if swag.IsZero(m.FileUpload) { // not required
		return nil
	}

	if m.FileUpload != nil {
		if err := m.FileUpload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileUpload")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) validateLauncherButton(formats strfmt.Registry) error {

	if swag.IsZero(m.LauncherButton) { // not required
		return nil
	}

	if m.LauncherButton != nil {
		if err := m.LauncherButton.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("launcherButton")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) validateStyles(formats strfmt.Registry) error {

	if swag.IsZero(m.Styles) { // not required
		return nil
	}

	if m.Styles != nil {
		if err := m.Styles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("styles")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessengerSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessengerSettings) UnmarshalBinary(b []byte) error {
	var res MessengerSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}