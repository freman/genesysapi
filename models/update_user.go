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

// UpdateUser update user
//
// swagger:model UpdateUser
type UpdateUser struct {

	// The value that denotes if acdAutoAnswer is set on the user
	AcdAutoAnswer bool `json:"acdAutoAnswer"`

	// Email address, phone number, and/or extension for this user. One entry is allowed per media type
	Addresses []*Contact `json:"addresses"`

	// biography
	Biography *Biography `json:"biography,omitempty"`

	// certifications
	Certifications []string `json:"certifications"`

	// chat
	Chat *Chat `json:"chat,omitempty"`

	// department
	Department *string `json:"department,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// employer info
	EmployerInfo *EmployerInfo `json:"employerInfo,omitempty"`

	// The groups the user is a member of
	Groups []*Group `json:"groups"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// images
	Images []*UserImage `json:"images"`

	// The user placement at each site location.
	Locations []*Location `json:"locations"`

	// manager
	Manager string `json:"manager,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The address(s) used for primary contact. Updates to the corresponding address in the addresses list will be reflected here.
	// Read Only: true
	PrimaryContactInfo []*Contact `json:"primaryContactInfo"`

	// Profile skills possessed by the user
	ProfileSkills []string `json:"profileSkills"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The state of the user. This property can be used to restore a deleted user or transition between active and inactive. If specified, it is the only modifiable field.
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// title
	Title *string `json:"title,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// This value should be the current version of the user. The current version can be obtained with a GET on the user before doing a PATCH.
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this update user
func (m *UpdateUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBiography(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryContactInfo(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUser) validateAddresses(formats strfmt.Registry) error {

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

func (m *UpdateUser) validateBiography(formats strfmt.Registry) error {

	if swag.IsZero(m.Biography) { // not required
		return nil
	}

	if m.Biography != nil {
		if err := m.Biography.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("biography")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateUser) validateChat(formats strfmt.Registry) error {

	if swag.IsZero(m.Chat) { // not required
		return nil
	}

	if m.Chat != nil {
		if err := m.Chat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chat")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateUser) validateEmployerInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.EmployerInfo) { // not required
		return nil
	}

	if m.EmployerInfo != nil {
		if err := m.EmployerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateUser) validateGroups(formats strfmt.Registry) error {

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

func (m *UpdateUser) validateImages(formats strfmt.Registry) error {

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

func (m *UpdateUser) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	for i := 0; i < len(m.Locations); i++ {
		if swag.IsZero(m.Locations[i]) { // not required
			continue
		}

		if m.Locations[i] != nil {
			if err := m.Locations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateUser) validatePrimaryContactInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryContactInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.PrimaryContactInfo); i++ {
		if swag.IsZero(m.PrimaryContactInfo[i]) { // not required
			continue
		}

		if m.PrimaryContactInfo[i] != nil {
			if err := m.PrimaryContactInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("primaryContactInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateUser) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var updateUserTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateUserTypeStatePropEnum = append(updateUserTypeStatePropEnum, v)
	}
}

const (

	// UpdateUserStateActive captures enum value "active"
	UpdateUserStateActive string = "active"

	// UpdateUserStateInactive captures enum value "inactive"
	UpdateUserStateInactive string = "inactive"

	// UpdateUserStateDeleted captures enum value "deleted"
	UpdateUserStateDeleted string = "deleted"
)

// prop value enum
func (m *UpdateUser) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateUserTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateUser) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUser) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUser) UnmarshalBinary(b []byte) error {
	var res UpdateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
