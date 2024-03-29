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

// WhatsAppID User information for a WhatsApp account
//
// swagger:model WhatsAppId
type WhatsAppID struct {

	// The displayName of this person's account in WhatsApp
	DisplayName string `json:"displayName,omitempty"`

	// The phone number associated with this WhatsApp account
	PhoneNumber *PhoneNumber `json:"phoneNumber,omitempty"`
}

// Validate validates this whats app Id
func (m *WhatsAppID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WhatsAppID) validatePhoneNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.PhoneNumber) { // not required
		return nil
	}

	if m.PhoneNumber != nil {
		if err := m.PhoneNumber.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneNumber")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phoneNumber")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this whats app Id based on the context it is used
func (m *WhatsAppID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePhoneNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WhatsAppID) contextValidatePhoneNumber(ctx context.Context, formats strfmt.Registry) error {

	if m.PhoneNumber != nil {
		if err := m.PhoneNumber.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phoneNumber")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phoneNumber")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WhatsAppID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WhatsAppID) UnmarshalBinary(b []byte) error {
	var res WhatsAppID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
