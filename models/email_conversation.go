// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmailConversation email conversation
//
// swagger:model EmailConversation
type EmailConversation struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The list of other media channels involved in the conversation.
	OtherMediaUris []strfmt.URI `json:"otherMediaUris"`

	// The list of participants involved in the conversation.
	Participants []*EmailMediaParticipant `json:"participants"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this email conversation
func (m *EmailConversation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtherMediaUris(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailConversation) validateOtherMediaUris(formats strfmt.Registry) error {

	if swag.IsZero(m.OtherMediaUris) { // not required
		return nil
	}

	for i := 0; i < len(m.OtherMediaUris); i++ {

		if err := validate.FormatOf("otherMediaUris"+"."+strconv.Itoa(i), "body", "uri", m.OtherMediaUris[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *EmailConversation) validateParticipants(formats strfmt.Registry) error {

	if swag.IsZero(m.Participants) { // not required
		return nil
	}

	for i := 0; i < len(m.Participants); i++ {
		if swag.IsZero(m.Participants[i]) { // not required
			continue
		}

		if m.Participants[i] != nil {
			if err := m.Participants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("participants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailConversation) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailConversation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailConversation) UnmarshalBinary(b []byte) error {
	var res EmailConversation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}