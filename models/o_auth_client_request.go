// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OAuthClientRequest o auth client request
//
// swagger:model OAuthClientRequest
type OAuthClientRequest struct {

	// The number of seconds, between 5mins and 48hrs, until tokens created with this client expire. If this field is omitted, a default of 24 hours will be applied.
	AccessTokenValiditySeconds int64 `json:"accessTokenValiditySeconds,omitempty"`

	// The OAuth Grant/Client type supported by this client.
	// Code Authorization Grant/Client type - Preferred client type where the Client ID and Secret are required to create tokens. Used where the secret can be secured.
	// PKCE-Enabled Code Authorization grant type - Code grant type which requires PKCE challenge and verifier to create tokens. Used in public clients for increased security.
	// Implicit grant type - Client ID only is required to create tokens. Used in browser and mobile apps where the secret can not be secured.
	// SAML2-Bearer extension grant type - SAML2 assertion provider for user authentication at the token endpoint.
	// Client Credential grant type - Used to created access tokens that are tied only to the client.
	//
	// Required: true
	// Enum: [CODE TOKEN SAML2-BEARER PASSWORD CLIENT-CREDENTIALS]
	AuthorizedGrantType *string `json:"authorizedGrantType"`

	// The time at which this client will be deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateToDelete strfmt.DateTime `json:"dateToDelete,omitempty"`

	// description
	Description string `json:"description,omitempty"`

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

	// The state of the OAuth client.
	// Active: The OAuth client can be used to create access tokens. This is the default state.
	// Disabled: Access tokens created by the client are invalid and new ones cannot be created.
	// Inactive: Access tokens cannot be created with this OAuth client and it will be deleted.
	// Enum: [active disabled inactive]
	State string `json:"state,omitempty"`
}

// Validate validates this o auth client request
func (m *OAuthClientRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizedGrantType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateToDelete(formats); err != nil {
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

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oAuthClientRequestTypeAuthorizedGrantTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CODE","TOKEN","SAML2-BEARER","PASSWORD","CLIENT-CREDENTIALS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oAuthClientRequestTypeAuthorizedGrantTypePropEnum = append(oAuthClientRequestTypeAuthorizedGrantTypePropEnum, v)
	}
}

const (

	// OAuthClientRequestAuthorizedGrantTypeCODE captures enum value "CODE"
	OAuthClientRequestAuthorizedGrantTypeCODE string = "CODE"

	// OAuthClientRequestAuthorizedGrantTypeTOKEN captures enum value "TOKEN"
	OAuthClientRequestAuthorizedGrantTypeTOKEN string = "TOKEN"

	// OAuthClientRequestAuthorizedGrantTypeSAML2DashBEARER captures enum value "SAML2-BEARER"
	OAuthClientRequestAuthorizedGrantTypeSAML2DashBEARER string = "SAML2-BEARER"

	// OAuthClientRequestAuthorizedGrantTypePASSWORD captures enum value "PASSWORD"
	OAuthClientRequestAuthorizedGrantTypePASSWORD string = "PASSWORD"

	// OAuthClientRequestAuthorizedGrantTypeCLIENTDashCREDENTIALS captures enum value "CLIENT-CREDENTIALS"
	OAuthClientRequestAuthorizedGrantTypeCLIENTDashCREDENTIALS string = "CLIENT-CREDENTIALS"
)

// prop value enum
func (m *OAuthClientRequest) validateAuthorizedGrantTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oAuthClientRequestTypeAuthorizedGrantTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OAuthClientRequest) validateAuthorizedGrantType(formats strfmt.Registry) error {

	if err := validate.Required("authorizedGrantType", "body", m.AuthorizedGrantType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAuthorizedGrantTypeEnum("authorizedGrantType", "body", *m.AuthorizedGrantType); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientRequest) validateDateToDelete(formats strfmt.Registry) error {
	if swag.IsZero(m.DateToDelete) { // not required
		return nil
	}

	if err := validate.FormatOf("dateToDelete", "body", "date-time", m.DateToDelete.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientRequest) validateRegisteredRedirectURI(formats strfmt.Registry) error {
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

func (m *OAuthClientRequest) validateRoleDivisions(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roleDivisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OAuthClientRequest) validateRoleIds(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("roleIds", "body", m.RoleIds); err != nil {
		return err
	}

	return nil
}

var oAuthClientRequestTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","disabled","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oAuthClientRequestTypeStatePropEnum = append(oAuthClientRequestTypeStatePropEnum, v)
	}
}

const (

	// OAuthClientRequestStateActive captures enum value "active"
	OAuthClientRequestStateActive string = "active"

	// OAuthClientRequestStateDisabled captures enum value "disabled"
	OAuthClientRequestStateDisabled string = "disabled"

	// OAuthClientRequestStateInactive captures enum value "inactive"
	OAuthClientRequestStateInactive string = "inactive"
)

// prop value enum
func (m *OAuthClientRequest) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oAuthClientRequestTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OAuthClientRequest) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this o auth client request based on the context it is used
func (m *OAuthClientRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleDivisions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAuthClientRequest) contextValidateRoleDivisions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RoleDivisions); i++ {

		if m.RoleDivisions[i] != nil {
			if err := m.RoleDivisions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roleDivisions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roleDivisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OAuthClientRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuthClientRequest) UnmarshalBinary(b []byte) error {
	var res OAuthClientRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
