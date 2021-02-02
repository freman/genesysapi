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

// MessagingRecipient This is used to identify who the message is sent to, as well as who it was sent from. This information is channel specific - depends on capabilities to describe party by the platform
//
// swagger:model MessagingRecipient
type MessagingRecipient struct {

	// Sender's email address
	// Read Only: true
	Email string `json:"email,omitempty"`

	// Sender's first name
	// Read Only: true
	FirstName string `json:"firstName,omitempty"`

	// The recipient identifier specific for particular channel/integration. This is required when sending a message.
	// Required: true
	ID *string `json:"id"`

	// Avatar image
	// Read Only: true
	Image string `json:"image,omitempty"`

	// Sender's last name
	// Read Only: true
	LastName string `json:"lastName,omitempty"`

	// Nickname/user name
	// Read Only: true
	Nickname string `json:"nickname,omitempty"`
}

// Validate validates this messaging recipient
func (m *MessagingRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessagingRecipient) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessagingRecipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessagingRecipient) UnmarshalBinary(b []byte) error {
	var res MessagingRecipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}