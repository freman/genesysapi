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

	// The homescreen settings for messenger
	HomeScreen *MessengerHomeScreen `json:"homeScreen,omitempty"`

	// The launcher button settings for messenger
	LauncherButton *LauncherButtonSettings `json:"launcherButton,omitempty"`

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

	if err := m.validateHomeScreen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLauncherButton(formats); err != nil {
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apps")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileUpload")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) validateHomeScreen(formats strfmt.Registry) error {
	if swag.IsZero(m.HomeScreen) { // not required
		return nil
	}

	if m.HomeScreen != nil {
		if err := m.HomeScreen.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("homeScreen")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("homeScreen")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("launcherButton")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("styles")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this messenger settings based on the context it is used
func (m *MessengerSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileUpload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHomeScreen(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLauncherButton(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStyles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessengerSettings) contextValidateApps(ctx context.Context, formats strfmt.Registry) error {

	if m.Apps != nil {
		if err := m.Apps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apps")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) contextValidateFileUpload(ctx context.Context, formats strfmt.Registry) error {

	if m.FileUpload != nil {
		if err := m.FileUpload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileUpload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileUpload")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) contextValidateHomeScreen(ctx context.Context, formats strfmt.Registry) error {

	if m.HomeScreen != nil {
		if err := m.HomeScreen.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("homeScreen")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("homeScreen")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) contextValidateLauncherButton(ctx context.Context, formats strfmt.Registry) error {

	if m.LauncherButton != nil {
		if err := m.LauncherButton.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("launcherButton")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("launcherButton")
			}
			return err
		}
	}

	return nil
}

func (m *MessengerSettings) contextValidateStyles(ctx context.Context, formats strfmt.Registry) error {

	if m.Styles != nil {
		if err := m.Styles.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("styles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("styles")
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
