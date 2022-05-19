// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MessagingRecipient Information about the recipient the message is sent to or received from.
//
// swagger:model MessagingRecipient
type MessagingRecipient struct {

	// List of recipient additional identifiers
	AdditionalIds []*RecipientAdditionalIdentifier `json:"additionalIds"`

	// E-mail address of the recipient.
	// Read Only: true
	Email string `json:"email,omitempty"`

	// First name of the recipient.
	// Read Only: true
	FirstName string `json:"firstName,omitempty"`

	// The recipient ID specific to the provider.
	// Required: true
	ID *string `json:"id"`

	// The recipient ID type. This is used to indicate the format used for the ID.
	// Enum: [Email Phone Opaque]
	IDType string `json:"idType,omitempty"`

	// URL of an image that represents the recipient.
	// Read Only: true
	Image string `json:"image,omitempty"`

	// Last name of the recipient.
	// Read Only: true
	LastName string `json:"lastName,omitempty"`

	// Nickname or display name of the recipient.
	// Read Only: true
	Nickname string `json:"nickname,omitempty"`
}

// Validate validates this messaging recipient
func (m *MessagingRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessagingRecipient) validateAdditionalIds(formats strfmt.Registry) error {

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

func (m *MessagingRecipient) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var messagingRecipientTypeIDTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Email","Phone","Opaque"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messagingRecipientTypeIDTypePropEnum = append(messagingRecipientTypeIDTypePropEnum, v)
	}
}

const (

	// MessagingRecipientIDTypeEmail captures enum value "Email"
	MessagingRecipientIDTypeEmail string = "Email"

	// MessagingRecipientIDTypePhone captures enum value "Phone"
	MessagingRecipientIDTypePhone string = "Phone"

	// MessagingRecipientIDTypeOpaque captures enum value "Opaque"
	MessagingRecipientIDTypeOpaque string = "Opaque"
)

// prop value enum
func (m *MessagingRecipient) validateIDTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messagingRecipientTypeIDTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessagingRecipient) validateIDType(formats strfmt.Registry) error {

	if swag.IsZero(m.IDType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDTypeEnum("idType", "body", m.IDType); err != nil {
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
