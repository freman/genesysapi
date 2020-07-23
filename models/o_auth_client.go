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

// OAuthClient o auth client
//
// swagger:model OAuthClient
type OAuthClient struct {

	// The number of seconds, between 5mins and 48hrs, until tokens created with this client expire. If this field is omitted, a default of 24 hours will be applied.
	AccessTokenValiditySeconds int64 `json:"accessTokenValiditySeconds,omitempty"`

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

	// description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// User that last modified this client
	ModifiedBy *DomainEntityRef `json:"modifiedBy,omitempty"`

	// The name of the OAuth client.
	// Required: true
	Name *string `json:"name"`

	// List of allowed callbacks for this client. For example: https://myap.example.com/auth/callback
	RegisteredRedirectURI []strfmt.URI `json:"registeredRedirectUri"`

	// Set of roles and their corresponding divisions associated with this client. Roles and divisions only apply to clients using the client_credential grant
	// Unique: true
	RoleDivisions []*RoleDivision `json:"roleDivisions"`

	// Deprecated. Use roleDivisions instead.
	// Unique: true
	RoleIds []string `json:"roleIds"`

	// The scope requested by this client. Scopes only apply to clients not using the client_credential grant
	Scope []string `json:"scope"`

	// System created secret assigned to this client. Secrets are required for code authorization and client credential grants.
	Secret string `json:"secret,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this o auth client
func (m *OAuthClient) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRegisteredRedirectURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleDivisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleIds(formats); err != nil {
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

var oAuthClientTypeAuthorizedGrantTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CODE","TOKEN","SAML2BEARER","PASSWORD","CLIENT_CREDENTIALS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oAuthClientTypeAuthorizedGrantTypePropEnum = append(oAuthClientTypeAuthorizedGrantTypePropEnum, v)
	}
}

const (

	// OAuthClientAuthorizedGrantTypeCODE captures enum value "CODE"
	OAuthClientAuthorizedGrantTypeCODE string = "CODE"

	// OAuthClientAuthorizedGrantTypeTOKEN captures enum value "TOKEN"
	OAuthClientAuthorizedGrantTypeTOKEN string = "TOKEN"

	// OAuthClientAuthorizedGrantTypeSAML2BEARER captures enum value "SAML2BEARER"
	OAuthClientAuthorizedGrantTypeSAML2BEARER string = "SAML2BEARER"

	// OAuthClientAuthorizedGrantTypePASSWORD captures enum value "PASSWORD"
	OAuthClientAuthorizedGrantTypePASSWORD string = "PASSWORD"

	// OAuthClientAuthorizedGrantTypeCLIENTCREDENTIALS captures enum value "CLIENT_CREDENTIALS"
	OAuthClientAuthorizedGrantTypeCLIENTCREDENTIALS string = "CLIENT_CREDENTIALS"
)

// prop value enum
func (m *OAuthClient) validateAuthorizedGrantTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oAuthClientTypeAuthorizedGrantTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OAuthClient) validateAuthorizedGrantType(formats strfmt.Registry) error {

	if err := validate.Required("authorizedGrantType", "body", m.AuthorizedGrantType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAuthorizedGrantTypeEnum("authorizedGrantType", "body", *m.AuthorizedGrantType); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClient) validateCreatedBy(formats strfmt.Registry) error {

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

func (m *OAuthClient) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClient) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClient) validateModifiedBy(formats strfmt.Registry) error {

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

func (m *OAuthClient) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClient) validateRegisteredRedirectURI(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredRedirectURI) { // not required
		return nil
	}

	for i := 0; i < len(m.RegisteredRedirectURI); i++ {

		if err := validate.FormatOf("registeredRedirectUri"+"."+strconv.Itoa(i), "body", "uri", m.RegisteredRedirectURI[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *OAuthClient) validateRoleDivisions(formats strfmt.Registry) error {

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

func (m *OAuthClient) validateRoleIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("roleIds", "body", m.RoleIds); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClient) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OAuthClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuthClient) UnmarshalBinary(b []byte) error {
	var res OAuthClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
