// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokenInfoClonedUser token info cloned user
//
// swagger:model TokenInfoClonedUser
type TokenInfoClonedUser struct {

	// User id of the original native user
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Organization of the original native user
	// Read Only: true
	Organization *Entity `json:"organization,omitempty"`
}

// Validate validates this token info cloned user
func (m *TokenInfoClonedUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenInfoClonedUser) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenInfoClonedUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenInfoClonedUser) UnmarshalBinary(b []byte) error {
	var res TokenInfoClonedUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
