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

// InboundDomainPatchRequest inbound domain patch request
//
// swagger:model InboundDomainPatchRequest
type InboundDomainPatchRequest struct {

	// The custom SMTP server integration to use when sending outbound emails from this domain.
	CustomSMTPServer *DomainEntityRef `json:"customSMTPServer,omitempty"`

	// The DNS settings if the inbound domain is using a custom Mail From. These settings can only be used on InboundDomains where subDomain is false.
	MailFromSettings *MailFromResult `json:"mailFromSettings,omitempty"`
}

// Validate validates this inbound domain patch request
func (m *InboundDomainPatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomSMTPServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMailFromSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundDomainPatchRequest) validateCustomSMTPServer(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomSMTPServer) { // not required
		return nil
	}

	if m.CustomSMTPServer != nil {
		if err := m.CustomSMTPServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customSMTPServer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customSMTPServer")
			}
			return err
		}
	}

	return nil
}

func (m *InboundDomainPatchRequest) validateMailFromSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.MailFromSettings) { // not required
		return nil
	}

	if m.MailFromSettings != nil {
		if err := m.MailFromSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mailFromSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mailFromSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this inbound domain patch request based on the context it is used
func (m *InboundDomainPatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomSMTPServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMailFromSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundDomainPatchRequest) contextValidateCustomSMTPServer(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomSMTPServer != nil {
		if err := m.CustomSMTPServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customSMTPServer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customSMTPServer")
			}
			return err
		}
	}

	return nil
}

func (m *InboundDomainPatchRequest) contextValidateMailFromSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.MailFromSettings != nil {
		if err := m.MailFromSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mailFromSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mailFromSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InboundDomainPatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InboundDomainPatchRequest) UnmarshalBinary(b []byte) error {
	var res InboundDomainPatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
