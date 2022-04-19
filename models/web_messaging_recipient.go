// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebMessagingRecipient Information about the recipient the message is sent to or received from.
//
// swagger:model WebMessagingRecipient
type WebMessagingRecipient struct {

	// List of recipient additional identifiers
	// Read Only: true
	AdditionalIds []*RecipientAdditionalIdentifier `json:"additionalIds"`

	// First name of the recipient.
	// Read Only: true
	FirstName string `json:"firstName,omitempty"`

	// Last name of the recipient.
	// Read Only: true
	LastName string `json:"lastName,omitempty"`

	// Nickname or display name of the recipient.
	// Read Only: true
	Nickname string `json:"nickname,omitempty"`
}

// Validate validates this web messaging recipient
func (m *WebMessagingRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebMessagingRecipient) validateAdditionalIds(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalIds) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalIds); i++ {
		if swag.IsZero(m.AdditionalIds[i]) { // not required
			continue
		}

		if m.AdditionalIds[i] != nil {
			if err := m.AdditionalIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebMessagingRecipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebMessagingRecipient) UnmarshalBinary(b []byte) error {
	var res WebMessagingRecipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
