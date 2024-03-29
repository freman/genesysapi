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

// SmsConfig sms config
//
// swagger:model SmsConfig
type SmsConfig struct {

	// The content template used to formulate the message to send to the contact.
	ContentTemplate *DomainEntityRef `json:"contentTemplate,omitempty"`

	// The Contact List column specifying the message to send to the contact.
	// Required: true
	MessageColumn *string `json:"messageColumn"`

	// The Contact List column specifying the phone number to send a message to.
	// Required: true
	PhoneColumn *string `json:"phoneColumn"`

	// A reference to the SMS Phone Number that will be used as the sender of a message.
	// Required: true
	SenderSmsPhoneNumber *SmsPhoneNumberRef `json:"senderSmsPhoneNumber"`
}

// Validate validates this sms config
func (m *SmsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageColumn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneColumn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderSmsPhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmsConfig) validateContentTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentTemplate) { // not required
		return nil
	}

	if m.ContentTemplate != nil {
		if err := m.ContentTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *SmsConfig) validateMessageColumn(formats strfmt.Registry) error {

	if err := validate.Required("messageColumn", "body", m.MessageColumn); err != nil {
		return err
	}

	return nil
}

func (m *SmsConfig) validatePhoneColumn(formats strfmt.Registry) error {

	if err := validate.Required("phoneColumn", "body", m.PhoneColumn); err != nil {
		return err
	}

	return nil
}

func (m *SmsConfig) validateSenderSmsPhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("senderSmsPhoneNumber", "body", m.SenderSmsPhoneNumber); err != nil {
		return err
	}

	if m.SenderSmsPhoneNumber != nil {
		if err := m.SenderSmsPhoneNumber.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("senderSmsPhoneNumber")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("senderSmsPhoneNumber")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sms config based on the context it is used
func (m *SmsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContentTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSenderSmsPhoneNumber(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmsConfig) contextValidateContentTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ContentTemplate != nil {
		if err := m.ContentTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *SmsConfig) contextValidateSenderSmsPhoneNumber(ctx context.Context, formats strfmt.Registry) error {

	if m.SenderSmsPhoneNumber != nil {
		if err := m.SenderSmsPhoneNumber.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("senderSmsPhoneNumber")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("senderSmsPhoneNumber")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmsConfig) UnmarshalBinary(b []byte) error {
	var res SmsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
