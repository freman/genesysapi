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

// AssignUsers assign users
//
// swagger:model AssignUsers
type AssignUsers struct {

	// List of user ids to assign to a performance profile
	// Required: true
	MembersToAssign []string `json:"membersToAssign"`

	// List of user ids to remove from a performance profile
	// Required: true
	MembersToRemove []string `json:"membersToRemove"`
}

// Validate validates this assign users
func (m *AssignUsers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembersToAssign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembersToRemove(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssignUsers) validateMembersToAssign(formats strfmt.Registry) error {

	if err := validate.Required("membersToAssign", "body", m.MembersToAssign); err != nil {
		return err
	}

	return nil
}

func (m *AssignUsers) validateMembersToRemove(formats strfmt.Registry) error {

	if err := validate.Required("membersToRemove", "body", m.MembersToRemove); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssignUsers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignUsers) UnmarshalBinary(b []byte) error {
	var res AssignUsers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}