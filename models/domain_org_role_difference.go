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

// DomainOrgRoleDifference domain org role difference
//
// swagger:model DomainOrgRoleDifference
type DomainOrgRoleDifference struct {

	// added permission policies
	AddedPermissionPolicies []*DomainPermissionPolicy `json:"addedPermissionPolicies"`

	// removed permission policies
	RemovedPermissionPolicies []*DomainPermissionPolicy `json:"removedPermissionPolicies"`

	// role from default
	RoleFromDefault *DomainOrganizationRole `json:"roleFromDefault,omitempty"`

	// same permission policies
	SamePermissionPolicies []*DomainPermissionPolicy `json:"samePermissionPolicies"`

	// user org role
	UserOrgRole *DomainOrganizationRole `json:"userOrgRole,omitempty"`
}

// Validate validates this domain org role difference
func (m *DomainOrgRoleDifference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddedPermissionPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemovedPermissionPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleFromDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamePermissionPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserOrgRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainOrgRoleDifference) validateAddedPermissionPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.AddedPermissionPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.AddedPermissionPolicies); i++ {
		if swag.IsZero(m.AddedPermissionPolicies[i]) { // not required
			continue
		}

		if m.AddedPermissionPolicies[i] != nil {
			if err := m.AddedPermissionPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addedPermissionPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainOrgRoleDifference) validateRemovedPermissionPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.RemovedPermissionPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.RemovedPermissionPolicies); i++ {
		if swag.IsZero(m.RemovedPermissionPolicies[i]) { // not required
			continue
		}

		if m.RemovedPermissionPolicies[i] != nil {
			if err := m.RemovedPermissionPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("removedPermissionPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainOrgRoleDifference) validateRoleFromDefault(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleFromDefault) { // not required
		return nil
	}

	if m.RoleFromDefault != nil {
		if err := m.RoleFromDefault.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roleFromDefault")
			}
			return err
		}
	}

	return nil
}

func (m *DomainOrgRoleDifference) validateSamePermissionPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.SamePermissionPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.SamePermissionPolicies); i++ {
		if swag.IsZero(m.SamePermissionPolicies[i]) { // not required
			continue
		}

		if m.SamePermissionPolicies[i] != nil {
			if err := m.SamePermissionPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("samePermissionPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainOrgRoleDifference) validateUserOrgRole(formats strfmt.Registry) error {

	if swag.IsZero(m.UserOrgRole) { // not required
		return nil
	}

	if m.UserOrgRole != nil {
		if err := m.UserOrgRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userOrgRole")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainOrgRoleDifference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainOrgRoleDifference) UnmarshalBinary(b []byte) error {
	var res DomainOrgRoleDifference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}