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

// VoicemailUserPolicy voicemail user policy
//
// swagger:model VoicemailUserPolicy
type VoicemailUserPolicy struct {

	// The number of seconds to ring the user's phone before a call is transfered to voicemail
	AlertTimeoutSeconds int32 `json:"alertTimeoutSeconds,omitempty"`

	// Whether the user has voicemail enabled
	// Read Only: true
	Enabled *bool `json:"enabled,omitempty"`

	// The date the policy was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Read Only: true
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The user's PIN to access their voicemail. This property is only used for updates and never provided otherwise to ensure security
	Pin string `json:"pin,omitempty"`
}

// Validate validates this voicemail user policy
func (m *VoicemailUserPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VoicemailUserPolicy) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VoicemailUserPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoicemailUserPolicy) UnmarshalBinary(b []byte) error {
	var res VoicemailUserPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
