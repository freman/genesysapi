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

// DirectRoutingCallMediaSettings direct routing call media settings
//
// swagger:model DirectRoutingCallMediaSettings
type DirectRoutingCallMediaSettings struct {

	// Toggle that enables Direct Routing for this media type.
	Enabled bool `json:"enabled"`

	// The Direct Routing inbound flow id for this media type.
	InboundFlow *AddressableEntityRef `json:"inboundFlow,omitempty"`

	// ID of the in-queue flow that handles voicemails for Direct Routing and to transfer calls to ACD voicemail.
	VoicemailFlow *AddressableEntityRef `json:"voicemailFlow,omitempty"`
}

// Validate validates this direct routing call media settings
func (m *DirectRoutingCallMediaSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInboundFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoicemailFlow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectRoutingCallMediaSettings) validateInboundFlow(formats strfmt.Registry) error {
	if swag.IsZero(m.InboundFlow) { // not required
		return nil
	}

	if m.InboundFlow != nil {
		if err := m.InboundFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inboundFlow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inboundFlow")
			}
			return err
		}
	}

	return nil
}

func (m *DirectRoutingCallMediaSettings) validateVoicemailFlow(formats strfmt.Registry) error {
	if swag.IsZero(m.VoicemailFlow) { // not required
		return nil
	}

	if m.VoicemailFlow != nil {
		if err := m.VoicemailFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voicemailFlow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("voicemailFlow")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this direct routing call media settings based on the context it is used
func (m *DirectRoutingCallMediaSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInboundFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVoicemailFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectRoutingCallMediaSettings) contextValidateInboundFlow(ctx context.Context, formats strfmt.Registry) error {

	if m.InboundFlow != nil {
		if err := m.InboundFlow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inboundFlow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inboundFlow")
			}
			return err
		}
	}

	return nil
}

func (m *DirectRoutingCallMediaSettings) contextValidateVoicemailFlow(ctx context.Context, formats strfmt.Registry) error {

	if m.VoicemailFlow != nil {
		if err := m.VoicemailFlow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voicemailFlow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("voicemailFlow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectRoutingCallMediaSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectRoutingCallMediaSettings) UnmarshalBinary(b []byte) error {
	var res DirectRoutingCallMediaSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
