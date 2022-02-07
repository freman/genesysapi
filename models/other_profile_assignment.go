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

// OtherProfileAssignment other profile assignment
//
// swagger:model OtherProfileAssignment
type OtherProfileAssignment struct {

	// The current performance profile that this user belongs to
	CurrentProfile *DomainEntityRef `json:"currentProfile,omitempty"`

	// The globally unique identifier for the object.
	ID string `json:"id,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this other profile assignment
func (m *OtherProfileAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentProfile(formats); err != nil {
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

func (m *OtherProfileAssignment) validateCurrentProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentProfile) { // not required
		return nil
	}

	if m.CurrentProfile != nil {
		if err := m.CurrentProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentProfile")
			}
			return err
		}
	}

	return nil
}

func (m *OtherProfileAssignment) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OtherProfileAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OtherProfileAssignment) UnmarshalBinary(b []byte) error {
	var res OtherProfileAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}