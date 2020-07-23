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

// GroupUpdate group update
//
// swagger:model GroupUpdate
type GroupUpdate struct {

	// addresses
	Addresses []*GroupContact `json:"addresses"`

	// description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// images
	Images []*UserImage `json:"images"`

	// The group name.
	Name string `json:"name,omitempty"`

	// Owners of the group
	OwnerIds []string `json:"ownerIds"`

	// Are membership rules visible to the person requesting to view the group
	RulesVisible bool `json:"rulesVisible,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// State of the group.
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// Current version for this resource.
	// Required: true
	Version *int32 `json:"version"`

	// Who can view this group
	// Enum: [public ownerIds members]
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this group update
func (m *GroupUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
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

func (m *GroupUpdate) validateAddresses(formats strfmt.Registry) error {

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

func (m *GroupUpdate) validateImages(formats strfmt.Registry) error {

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

func (m *GroupUpdate) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var groupUpdateTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupUpdateTypeStatePropEnum = append(groupUpdateTypeStatePropEnum, v)
	}
}

const (

	// GroupUpdateStateActive captures enum value "active"
	GroupUpdateStateActive string = "active"

	// GroupUpdateStateInactive captures enum value "inactive"
	GroupUpdateStateInactive string = "inactive"

	// GroupUpdateStateDeleted captures enum value "deleted"
	GroupUpdateStateDeleted string = "deleted"
)

// prop value enum
func (m *GroupUpdate) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupUpdateTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupUpdate) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *GroupUpdate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

var groupUpdateTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","ownerIds","members"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupUpdateTypeVisibilityPropEnum = append(groupUpdateTypeVisibilityPropEnum, v)
	}
}

const (

	// GroupUpdateVisibilityPublic captures enum value "public"
	GroupUpdateVisibilityPublic string = "public"

	// GroupUpdateVisibilityOwnerIds captures enum value "ownerIds"
	GroupUpdateVisibilityOwnerIds string = "ownerIds"

	// GroupUpdateVisibilityMembers captures enum value "members"
	GroupUpdateVisibilityMembers string = "members"
)

// prop value enum
func (m *GroupUpdate) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, groupUpdateTypeVisibilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GroupUpdate) validateVisibility(formats strfmt.Registry) error {

	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupUpdate) UnmarshalBinary(b []byte) error {
	var res GroupUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
