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

// OrgOAuthClient org o auth client
//
// swagger:model OrgOAuthClient
type OrgOAuthClient struct {

	// The OAuth Grant/Client type supported by this client.
	// Code Authorization Grant/Client type - Preferred client type where the Client ID and Secret are required to create tokens. Used where the secret can be secured.
	// Implicit grant type - Client ID only is required to create tokens. Used in browser and mobile apps where the secret can not be secured.
	// SAML2-Bearer extension grant type - SAML2 assertion provider for user authentication at the token endpoint.
	// Client Credential grant type - Used to created access tokens that are tied only to the client.
	//
	// Required: true
	// Enum: [CODE TOKEN SAML2BEARER PASSWORD CLIENT_CREDENTIALS]
	AuthorizedGrantType *string `json:"authorizedGrantType"`

	// User that created this client
	CreatedBy *DomainEntityRef `json:"createdBy,omitempty"`

	// Date this client was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date this client was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// User that last modified this client
	ModifiedBy *DomainEntityRef `json:"modifiedBy,omitempty"`

	// The name of the OAuth client.
	// Required: true
	Name *string `json:"name"`

	// The  oauth client's organization.
	// Read Only: true
	Organization *NamedEntity `json:"organization,omitempty"`

	// Set of roles and their corresponding divisions associated with this client. Roles and divisions only apply to clients using the client_credential grant
	// Unique: true
	RoleDivisions []*RoleDivision `json:"roleDivisions"`

	// The scope requested by this client. Scopes only apply to clients not using the client_credential grant
	Scope []string `json:"scope"`

	// The state of the OAuth client.
	// Active: The OAuth client can be used to create access tokens. This is the default state.
	// Disabled: Access tokens created by the client are invalid and new ones cannot be created.
	// Inactive: Access tokens cannot be created with this OAuth client and it will be deleted.
	// Enum: [active disabled inactive]
	State string `json:"state,omitempty"`
}

// Validate validates this org o auth client
func (m *OrgOAuthClient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizedGrantType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleDivisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var orgOAuthClientTypeAuthorizedGrantTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CODE","TOKEN","SAML2BEARER","PASSWORD","CLIENT_CREDENTIALS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orgOAuthClientTypeAuthorizedGrantTypePropEnum = append(orgOAuthClientTypeAuthorizedGrantTypePropEnum, v)
	}
}

const (

	// OrgOAuthClientAuthorizedGrantTypeCODE captures enum value "CODE"
	OrgOAuthClientAuthorizedGrantTypeCODE string = "CODE"

	// OrgOAuthClientAuthorizedGrantTypeTOKEN captures enum value "TOKEN"
	OrgOAuthClientAuthorizedGrantTypeTOKEN string = "TOKEN"

	// OrgOAuthClientAuthorizedGrantTypeSAML2BEARER captures enum value "SAML2BEARER"
	OrgOAuthClientAuthorizedGrantTypeSAML2BEARER string = "SAML2BEARER"

	// OrgOAuthClientAuthorizedGrantTypePASSWORD captures enum value "PASSWORD"
	OrgOAuthClientAuthorizedGrantTypePASSWORD string = "PASSWORD"

	// OrgOAuthClientAuthorizedGrantTypeCLIENTCREDENTIALS captures enum value "CLIENT_CREDENTIALS"
	OrgOAuthClientAuthorizedGrantTypeCLIENTCREDENTIALS string = "CLIENT_CREDENTIALS"
)

// prop value enum
func (m *OrgOAuthClient) validateAuthorizedGrantTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orgOAuthClientTypeAuthorizedGrantTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrgOAuthClient) validateAuthorizedGrantType(formats strfmt.Registry) error {

	if err := validate.Required("authorizedGrantType", "body", m.AuthorizedGrantType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAuthorizedGrantTypeEnum("authorizedGrantType", "body", *m.AuthorizedGrantType); err != nil {
		return err
	}

	return nil
}

func (m *OrgOAuthClient) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *OrgOAuthClient) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrgOAuthClient) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrgOAuthClient) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *OrgOAuthClient) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OrgOAuthClient) validateOrganization(formats strfmt.Registry) error {

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

func (m *OrgOAuthClient) validateRoleDivisions(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleDivisions) { // not required
		return nil
	}

	if err := validate.UniqueItems("roleDivisions", "body", m.RoleDivisions); err != nil {
		return err
	}

	for i := 0; i < len(m.RoleDivisions); i++ {
		if swag.IsZero(m.RoleDivisions[i]) { // not required
			continue
		}

		if m.RoleDivisions[i] != nil {
			if err := m.RoleDivisions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roleDivisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var orgOAuthClientTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","disabled","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orgOAuthClientTypeStatePropEnum = append(orgOAuthClientTypeStatePropEnum, v)
	}
}

const (

	// OrgOAuthClientStateActive captures enum value "active"
	OrgOAuthClientStateActive string = "active"

	// OrgOAuthClientStateDisabled captures enum value "disabled"
	OrgOAuthClientStateDisabled string = "disabled"

	// OrgOAuthClientStateInactive captures enum value "inactive"
	OrgOAuthClientStateInactive string = "inactive"
)

// prop value enum
func (m *OrgOAuthClient) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orgOAuthClientTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrgOAuthClient) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrgOAuthClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrgOAuthClient) UnmarshalBinary(b []byte) error {
	var res OrgOAuthClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
