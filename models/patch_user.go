// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchUser patch user
//
// swagger:model PatchUser
type PatchUser struct {

	// The value that denotes if acdAutoAnswer is set on the user
	AcdAutoAnswer bool `json:"acdAutoAnswer,omitempty"`

	// biography
	Biography *Biography `json:"biography,omitempty"`

	// certifications
	Certifications []string `json:"certifications"`

	// employer info
	EmployerInfo *EmployerInfo `json:"employerInfo,omitempty"`

	// The globally unique identifier for the object.
	ID string `json:"id,omitempty"`
}

// Validate validates this patch user
func (m *PatchUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBiography(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployerInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchUser) validateBiography(formats strfmt.Registry) error {

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

func (m *PatchUser) validateEmployerInfo(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PatchUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchUser) UnmarshalBinary(b []byte) error {
	var res PatchUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
