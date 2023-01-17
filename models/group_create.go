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

// GroupCreate group create
//
// swagger:model GroupCreate
type GroupCreate struct {

	// addresses
	Addresses []*GroupContact `json:"addresses"`

	// Last modified date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// images
	Images []*UserImage `json:"images"`

	// Number of members.
	// Read Only: true
	MemberCount int64 `json:"memberCount,omitempty"`

	// The group name.
	// Required: true
	Name *string `json:"name"`

	// Owners of the group
	OwnerIds []string `json:"ownerIds"`

	// Are membership rules visible to the person requesting to view the group
	// Required: true
	RulesVisible *bool `json:"rulesVisible"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Active, inactive, or deleted state.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// Type of group.
	// Required: true
	// Enum: [official social]
	Type *string `json:"type"`

	// Current version for this resource.
	// Read Only: true
	Version int32 `json:"version,omitempty"`

	// Who can view this group
	// Required: true
	// Enum: [public owners members]
	Visibility *string `json:"visibility"`
}

// Validate validates this group create
func (m *GroupCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRulesVisible(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupCreate) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GroupCreate) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreate) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GroupCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreate) validateRulesVisible(formats strfmt.Registry) error {

	if err := validate.Required("rulesVisible", "body", m.RulesVisible); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreate) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var groupCreateTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupCreateTypeStatePropEnum = append(groupCreateTypeStatePropEnum, v)
	}
}

const (

	// GroupCreateStateActive captures enum value "active"
	GroupCreateStateActive string = "active"

	// GroupCreateStateInactive captures enum value "inactive"
	GroupCreateStateInactive string = "inactive"

	// GroupCreateStateDeleted captures enum value "deleted"
	GroupCreateStateDeleted string = "deleted"
)

// prop value enum
func (m *GroupCreate) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupCreateTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupCreate) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var groupCreateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["official","social"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupCreateTypeTypePropEnum = append(groupCreateTypeTypePropEnum, v)
	}
}

const (

	// GroupCreateTypeOfficial captures enum value "official"
	GroupCreateTypeOfficial string = "official"

	// GroupCreateTypeSocial captures enum value "social"
	GroupCreateTypeSocial string = "social"
)

// prop value enum
func (m *GroupCreate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupCreateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupCreate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

var groupCreateTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","owners","members"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupCreateTypeVisibilityPropEnum = append(groupCreateTypeVisibilityPropEnum, v)
	}
}

const (

	// GroupCreateVisibilityPublic captures enum value "public"
	GroupCreateVisibilityPublic string = "public"

	// GroupCreateVisibilityOwners captures enum value "owners"
	GroupCreateVisibilityOwners string = "owners"

	// GroupCreateVisibilityMembers captures enum value "members"
	GroupCreateVisibilityMembers string = "members"
)

// prop value enum
func (m *GroupCreate) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupCreateTypeVisibilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupCreate) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("visibility", "body", m.Visibility); err != nil {
		return err
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", *m.Visibility); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this group create based on the context it is used
func (m *GroupCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemberCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupCreate) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GroupCreate) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreate) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {
			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GroupCreate) contextValidateMemberCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "memberCount", "body", int64(m.MemberCount)); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreate) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreate) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreate) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", int32(m.Version)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupCreate) UnmarshalBinary(b []byte) error {
	var res GroupCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
