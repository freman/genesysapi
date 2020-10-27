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

// TrustGroup trust group
//
// swagger:model TrustGroup
type TrustGroup struct {

	// addresses
	Addresses []*GroupContact `json:"addresses"`

	// The user that added trusted group.
	// Read Only: true
	CreatedBy *OrgUser `json:"createdBy,omitempty"`

	// The date on which the trusted group was added. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

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
	Owners []*User `json:"owners"`

	// Are membership rules visible to the person requesting to view the group
	// Required: true
	RulesVisible *bool `json:"rulesVisible"`

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

// Validate validates this trust group
func (m *TrustGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
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

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRulesVisible(formats); err != nil {
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

func (m *TrustGroup) validateAddresses(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *TrustGroup) validateCreatedBy(formats strfmt.Registry) error {

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

func (m *TrustGroup) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrustGroup) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrustGroup) validateImages(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *TrustGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TrustGroup) validateOwners(formats strfmt.Registry) error {

	if swag.IsZero(m.Owners) { // not required
		return nil
	}

	for i := 0; i < len(m.Owners); i++ {
		if swag.IsZero(m.Owners[i]) { // not required
			continue
		}

		if m.Owners[i] != nil {
			if err := m.Owners[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("owners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TrustGroup) validateRulesVisible(formats strfmt.Registry) error {

	if err := validate.Required("rulesVisible", "body", m.RulesVisible); err != nil {
		return err
	}

	return nil
}

var trustGroupTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trustGroupTypeStatePropEnum = append(trustGroupTypeStatePropEnum, v)
	}
}

const (

	// TrustGroupStateActive captures enum value "active"
	TrustGroupStateActive string = "active"

	// TrustGroupStateInactive captures enum value "inactive"
	TrustGroupStateInactive string = "inactive"

	// TrustGroupStateDeleted captures enum value "deleted"
	TrustGroupStateDeleted string = "deleted"
)

// prop value enum
func (m *TrustGroup) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trustGroupTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TrustGroup) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var trustGroupTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["official","social"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trustGroupTypeTypePropEnum = append(trustGroupTypeTypePropEnum, v)
	}
}

const (

	// TrustGroupTypeOfficial captures enum value "official"
	TrustGroupTypeOfficial string = "official"

	// TrustGroupTypeSocial captures enum value "social"
	TrustGroupTypeSocial string = "social"
)

// prop value enum
func (m *TrustGroup) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trustGroupTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TrustGroup) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

var trustGroupTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","owners","members"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trustGroupTypeVisibilityPropEnum = append(trustGroupTypeVisibilityPropEnum, v)
	}
}

const (

	// TrustGroupVisibilityPublic captures enum value "public"
	TrustGroupVisibilityPublic string = "public"

	// TrustGroupVisibilityOwners captures enum value "owners"
	TrustGroupVisibilityOwners string = "owners"

	// TrustGroupVisibilityMembers captures enum value "members"
	TrustGroupVisibilityMembers string = "members"
)

// prop value enum
func (m *TrustGroup) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trustGroupTypeVisibilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TrustGroup) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("visibility", "body", m.Visibility); err != nil {
		return err
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", *m.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrustGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrustGroup) UnmarshalBinary(b []byte) error {
	var res TrustGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
