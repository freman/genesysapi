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

// DigitalAction digital action
//
// swagger:model DigitalAction
type DigitalAction struct {

	// The settings for an 'Append to DNC' action.
	AppendToDncActionSettings *AppendToDncActionSettings `json:"appendToDncActionSettings,omitempty"`

	// The settings for a 'do not send' action.
	DoNotSendActionSettings DoNotSendActionSettings `json:"doNotSendActionSettings,omitempty"`

	// The settings for an 'mark contact address uncontactable' action.
	MarkContactAddressUncontactableActionSettings MarkContactAddressUncontactableActionSettings `json:"markContactAddressUncontactableActionSettings,omitempty"`

	// The settings for a 'mark contact uncontactable' action.
	MarkContactUncontactableActionSettings *MarkContactUncontactableActionSettings `json:"markContactUncontactableActionSettings,omitempty"`

	// The settings for a 'Set content template' action.
	SetContentTemplateActionSettings *SetContentTemplateActionSettings `json:"setContentTemplateActionSettings,omitempty"`

	// The settings for an 'update contact column' action.
	UpdateContactColumnActionSettings *UpdateContactColumnActionSettings `json:"updateContactColumnActionSettings,omitempty"`
}

// Validate validates this digital action
func (m *DigitalAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppendToDncActionSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkContactUncontactableActionSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSetContentTemplateActionSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateContactColumnActionSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DigitalAction) validateAppendToDncActionSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.AppendToDncActionSettings) { // not required
		return nil
	}

	if m.AppendToDncActionSettings != nil {
		if err := m.AppendToDncActionSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appendToDncActionSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appendToDncActionSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DigitalAction) validateMarkContactUncontactableActionSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.MarkContactUncontactableActionSettings) { // not required
		return nil
	}

	if m.MarkContactUncontactableActionSettings != nil {
		if err := m.MarkContactUncontactableActionSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("markContactUncontactableActionSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("markContactUncontactableActionSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DigitalAction) validateSetContentTemplateActionSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SetContentTemplateActionSettings) { // not required
		return nil
	}

	if m.SetContentTemplateActionSettings != nil {
		if err := m.SetContentTemplateActionSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setContentTemplateActionSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("setContentTemplateActionSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DigitalAction) validateUpdateContactColumnActionSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateContactColumnActionSettings) { // not required
		return nil
	}

	if m.UpdateContactColumnActionSettings != nil {
		if err := m.UpdateContactColumnActionSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateContactColumnActionSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateContactColumnActionSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this digital action based on the context it is used
func (m *DigitalAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppendToDncActionSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarkContactUncontactableActionSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSetContentTemplateActionSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateContactColumnActionSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DigitalAction) contextValidateAppendToDncActionSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.AppendToDncActionSettings != nil {
		if err := m.AppendToDncActionSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appendToDncActionSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appendToDncActionSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DigitalAction) contextValidateMarkContactUncontactableActionSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.MarkContactUncontactableActionSettings != nil {
		if err := m.MarkContactUncontactableActionSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("markContactUncontactableActionSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("markContactUncontactableActionSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DigitalAction) contextValidateSetContentTemplateActionSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.SetContentTemplateActionSettings != nil {
		if err := m.SetContentTemplateActionSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setContentTemplateActionSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("setContentTemplateActionSettings")
			}
			return err
		}
	}

	return nil
}

func (m *DigitalAction) contextValidateUpdateContactColumnActionSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdateContactColumnActionSettings != nil {
		if err := m.UpdateContactColumnActionSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateContactColumnActionSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateContactColumnActionSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DigitalAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DigitalAction) UnmarshalBinary(b []byte) error {
	var res DigitalAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}