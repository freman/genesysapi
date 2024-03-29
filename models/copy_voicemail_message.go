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

// CopyVoicemailMessage Used to copy a VoicemailMessage to either a User or a Group
//
// swagger:model CopyVoicemailMessage
type CopyVoicemailMessage struct {

	// The id of the Group to copy the VoicemailMessage to
	GroupID string `json:"groupId,omitempty"`

	// The id of the User to copy the VoicemailMessage to
	UserID string `json:"userId,omitempty"`

	// The id of the VoicemailMessage to copy
	// Required: true
	VoicemailMessageID *string `json:"voicemailMessageId"`
}

// Validate validates this copy voicemail message
func (m *CopyVoicemailMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVoicemailMessageID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopyVoicemailMessage) validateVoicemailMessageID(formats strfmt.Registry) error {

	if err := validate.Required("voicemailMessageId", "body", m.VoicemailMessageID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this copy voicemail message based on context it is used
func (m *CopyVoicemailMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CopyVoicemailMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyVoicemailMessage) UnmarshalBinary(b []byte) error {
	var res CopyVoicemailMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
