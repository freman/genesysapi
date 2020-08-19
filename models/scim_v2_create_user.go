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

// ScimV2CreateUser Defines the creation of a SCIM user.
//
// swagger:model ScimV2CreateUser
type ScimV2CreateUser struct {

	// Indicates whether the user's administrative status is active.
	Active bool `json:"active,omitempty"`

	// The display name of the user.
	// Required: true
	DisplayName *string `json:"displayName"`

	// The list of the user's email addresses.
	Emails []*ScimEmail `json:"emails"`

	// The external ID of the user. Set by the provisioning client. "caseExact" is set to "true". "mutability" is set to "readWrite".
	ExternalID string `json:"externalId,omitempty"`

	// The list of groups that the user is a member of.
	Groups []*ScimV2GroupReference `json:"groups"`

	// The new password for the Genesys Cloud user. Does not return an existing password.
	Password string `json:"password,omitempty"`

	// The list of the user's phone numbers.
	PhoneNumbers []*ScimPhoneNumber `json:"phoneNumbers"`

	// The list of roles assigned to the user.
	Roles []*ScimUserRole `json:"roles"`

	// The list of supported schemas.
	// Read Only: true
	Schemas []string `json:"schemas"`

	// The user's title.
	Title string `json:"title,omitempty"`

	// The URI of the schema for the enterprise user.
	UrnIetfParamsScimSchemasExtensionEnterprise20User *ScimV2EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`

	// The URI of the schema for the Genesys Cloud user.
	UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User *ScimUserExtensions `json:"urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User,omitempty"`

	// The user's Genesys Cloud email address. Must be unique.
	// Required: true
	UserName *string `json:"userName"`
}

// Validate validates this scim v2 create user
func (m *ScimV2CreateUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumbers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrnIetfParamsScimSchemasExtensionEnterprise20User(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrnIetfParamsScimSchemasExtensionGenesysPurecloud20User(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimV2CreateUser) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *ScimV2CreateUser) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScimV2CreateUser) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScimV2CreateUser) validatePhoneNumbers(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumbers) { // not required
		return nil
	}

	for i := 0; i < len(m.PhoneNumbers); i++ {
		if swag.IsZero(m.PhoneNumbers[i]) { // not required
			continue
		}

		if m.PhoneNumbers[i] != nil {
			if err := m.PhoneNumbers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneNumbers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScimV2CreateUser) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScimV2CreateUser) validateUrnIetfParamsScimSchemasExtensionEnterprise20User(formats strfmt.Registry) error {

	if swag.IsZero(m.UrnIetfParamsScimSchemasExtensionEnterprise20User) { // not required
		return nil
	}

	if m.UrnIetfParamsScimSchemasExtensionEnterprise20User != nil {
		if err := m.UrnIetfParamsScimSchemasExtensionEnterprise20User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User")
			}
			return err
		}
	}

	return nil
}

func (m *ScimV2CreateUser) validateUrnIetfParamsScimSchemasExtensionGenesysPurecloud20User(formats strfmt.Registry) error {

	if swag.IsZero(m.UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User) { // not required
		return nil
	}

	if m.UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User != nil {
		if err := m.UrnIetfParamsScimSchemasExtensionGenesysPurecloud20User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User")
			}
			return err
		}
	}

	return nil
}

func (m *ScimV2CreateUser) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimV2CreateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimV2CreateUser) UnmarshalBinary(b []byte) error {
	var res ScimV2CreateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
