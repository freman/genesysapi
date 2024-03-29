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

// MessagingTemplate The messaging template identifies a structured message templates supported by a messaging channel.
//
// swagger:model MessagingTemplate
type MessagingTemplate struct {

	// Defines a messaging template for a WhatsApp messaging channel
	WhatsApp *WhatsAppDefinition `json:"whatsApp,omitempty"`
}

// Validate validates this messaging template
func (m *MessagingTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWhatsApp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessagingTemplate) validateWhatsApp(formats strfmt.Registry) error {
	if swag.IsZero(m.WhatsApp) { // not required
		return nil
	}

	if m.WhatsApp != nil {
		if err := m.WhatsApp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("whatsApp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("whatsApp")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this messaging template based on the context it is used
func (m *MessagingTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWhatsApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessagingTemplate) contextValidateWhatsApp(ctx context.Context, formats strfmt.Registry) error {

	if m.WhatsApp != nil {
		if err := m.WhatsApp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("whatsApp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("whatsApp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessagingTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessagingTemplate) UnmarshalBinary(b []byte) error {
	var res MessagingTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
