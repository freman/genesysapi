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

// OAuthClientListing o auth client listing
//
// swagger:model OAuthClientListing
type OAuthClientListing struct {

	// The number of seconds, between 5mins and 48hrs, until tokens created with this client expire. If this field is omitted, a default of 24 hours will be applied.
	AccessTokenValiditySeconds int64 `json:"accessTokenValiditySeconds,omitempty"`

	// User that created this client
	CreatedBy *DomainEntityRef `json:"createdBy,omitempty"`

	// Date this client was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date this client was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The time at which this client will be deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateToDelete strfmt.DateTime `json:"dateToDelete,omitempty"`

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

	// The state of the OAuth client.
	// Active: The OAuth client can be used to create access tokens. This is the default state.
	// Disabled: Access tokens created by the client are invalid and new ones cannot be created.
	// Inactive: Access tokens cannot be created with this OAuth client and it will be deleted.
	// Enum: [active disabled inactive]
	State string `json:"state,omitempty"`
}

// Validate validates this o auth client listing
func (m *OAuthClientListing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateToDelete(formats); err != nil {
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

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAuthClientListing) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *OAuthClientListing) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientListing) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientListing) validateDateToDelete(formats strfmt.Registry) error {
	if swag.IsZero(m.DateToDelete) { // not required
		return nil
	}

	if err := validate.FormatOf("dateToDelete", "body", "date-time", m.DateToDelete.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientListing) validateModifiedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *OAuthClientListing) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientListing) validateRegisteredRedirectURI(formats strfmt.Registry) error {
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

func (m *OAuthClientListing) validateRoleDivisions(formats strfmt.Registry) error {
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

func (m *OAuthClientListing) validateRoleIds(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("roleIds", "body", m.RoleIds); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientListing) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var oAuthClientListingTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","disabled","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oAuthClientListingTypeStatePropEnum = append(oAuthClientListingTypeStatePropEnum, v)
	}
}

const (

	// OAuthClientListingStateActive captures enum value "active"
	OAuthClientListingStateActive string = "active"

	// OAuthClientListingStateDisabled captures enum value "disabled"
	OAuthClientListingStateDisabled string = "disabled"

	// OAuthClientListingStateInactive captures enum value "inactive"
	OAuthClientListingStateInactive string = "inactive"
)

// prop value enum
func (m *OAuthClientListing) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oAuthClientListingTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OAuthClientListing) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this o auth client listing based on the context it is used
func (m *OAuthClientListing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleDivisions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAuthClientListing) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {
		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *OAuthClientListing) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *OAuthClientListing) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *OAuthClientListing) contextValidateRoleDivisions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *OAuthClientListing) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OAuthClientListing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuthClientListing) UnmarshalBinary(b []byte) error {
	var res OAuthClientListing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
