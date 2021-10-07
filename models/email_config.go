// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmailConfig email config
//
// swagger:model EmailConfig
type EmailConfig struct {

	// The content template used to formulate the email to send to the contact.
	ContentTemplate *DomainEntityRef `json:"contentTemplate,omitempty"`

	// The contact list columns specifying the email address(es) of the contact.
	// Required: true
	EmailColumns []string `json:"emailColumns"`

	// The email address that will be used as the sender of the email.
	// Required: true
	FromAddress *FromEmailAddress `json:"fromAddress"`

	// The email address from which any reply will be sent.
	ReplyToAddress *ReplyToEmailAddress `json:"replyToAddress,omitempty"`
}

// Validate validates this email config
func (m *EmailConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplyToAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailConfig) validateContentTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentTemplate) { // not required
		return nil
	}

	if m.ContentTemplate != nil {
		if err := m.ContentTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *EmailConfig) validateEmailColumns(formats strfmt.Registry) error {

	if err := validate.Required("emailColumns", "body", m.EmailColumns); err != nil {
		return err
	}

	return nil
}

func (m *EmailConfig) validateFromAddress(formats strfmt.Registry) error {

	if err := validate.Required("fromAddress", "body", m.FromAddress); err != nil {
		return err
	}

	if m.FromAddress != nil {
		if err := m.FromAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromAddress")
			}
			return err
		}
	}

	return nil
}

func (m *EmailConfig) validateReplyToAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyToAddress) { // not required
		return nil
	}

	if m.ReplyToAddress != nil {
		if err := m.ReplyToAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replyToAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailConfig) UnmarshalBinary(b []byte) error {
	var res EmailConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
