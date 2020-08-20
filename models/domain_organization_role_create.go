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

// DomainOrganizationRoleCreate domain organization role create
//
// swagger:model DomainOrganizationRoleCreate
type DomainOrganizationRoleCreate struct {

	// base
	Base bool `json:"base"`

	// default
	Default bool `json:"default"`

	// default role Id
	DefaultRoleID string `json:"defaultRoleId,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The role name
	// Required: true
	Name *string `json:"name"`

	// permission policies
	// Unique: true
	PermissionPolicies []*DomainPermissionPolicy `json:"permissionPolicies"`

	// permissions
	// Unique: true
	Permissions []string `json:"permissions"`

	// Optional unless patch operation.
	RoleNeedsUpdate bool `json:"roleNeedsUpdate"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// A collection of the permissions the role is not using
	// Read Only: true
	// Unique: true
	UnusedPermissions []string `json:"unusedPermissions"`

	// user count
	UserCount int32 `json:"userCount,omitempty"`
}

// Validate validates this domain organization role create
func (m *DomainOrganizationRoleCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissionPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnusedPermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainOrganizationRoleCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DomainOrganizationRoleCreate) validatePermissionPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.PermissionPolicies) { // not required
		return nil
	}

	if err := validate.UniqueItems("permissionPolicies", "body", m.PermissionPolicies); err != nil {
		return err
	}

	for i := 0; i < len(m.PermissionPolicies); i++ {
		if swag.IsZero(m.PermissionPolicies[i]) { // not required
			continue
		}

		if m.PermissionPolicies[i] != nil {
			if err := m.PermissionPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissionPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainOrganizationRoleCreate) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if err := validate.UniqueItems("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *DomainOrganizationRoleCreate) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainOrganizationRoleCreate) validateUnusedPermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.UnusedPermissions) { // not required
		return nil
	}

	if err := validate.UniqueItems("unusedPermissions", "body", m.UnusedPermissions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainOrganizationRoleCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainOrganizationRoleCreate) UnmarshalBinary(b []byte) error {
	var res DomainOrganizationRoleCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
