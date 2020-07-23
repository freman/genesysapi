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

// CreateUser create user
//
// swagger:model CreateUser
type CreateUser struct {

	// Email addresses and phone numbers for this user
	Addresses []*Contact `json:"addresses"`

	// department
	Department string `json:"department,omitempty"`

	// The division to which this user will belong
	// Required: true
	DivisionID *string `json:"divisionId"`

	// User's email and username
	// Required: true
	Email *string `json:"email"`

	// User's full name
	// Required: true
	Name *string `json:"name"`

	// User's password
	// Required: true
	Password *string `json:"password"`

	// Optional initialized state of the user. If not specified, state will be Active if invites are sent, otherwise Inactive.
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this create user
func (m *CreateUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivisionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
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

func (m *CreateUser) validateAddresses(formats strfmt.Registry) error {

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

func (m *CreateUser) validateDivisionID(formats strfmt.Registry) error {

	if err := validate.Required("divisionId", "body", m.DivisionID); err != nil {
		return err
	}

	return nil
}

func (m *CreateUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *CreateUser) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateUser) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

var createUserTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createUserTypeStatePropEnum = append(createUserTypeStatePropEnum, v)
	}
}

const (

	// CreateUserStateActive captures enum value "active"
	CreateUserStateActive string = "active"

	// CreateUserStateInactive captures enum value "inactive"
	CreateUserStateInactive string = "inactive"

	// CreateUserStateDeleted captures enum value "deleted"
	CreateUserStateDeleted string = "deleted"
)

// prop value enum
func (m *CreateUser) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createUserTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateUser) validateState(formats strfmt.Registry) error {

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
func (m *CreateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUser) UnmarshalBinary(b []byte) error {
	var res CreateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
