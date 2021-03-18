// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserBestPoints user best points
//
// swagger:model UserBestPoints
type UserBestPoints struct {

	// List of best point for the requested user
	// Read Only: true
	BestPoints []*UserBestPointsItem `json:"bestPoints"`

	// The requested user for the best points
	// Read Only: true
	User *UserReference `json:"user,omitempty"`
}

// Validate validates this user best points
func (m *UserBestPoints) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBestPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserBestPoints) validateBestPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.BestPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.BestPoints); i++ {
		if swag.IsZero(m.BestPoints[i]) { // not required
			continue
		}

		if m.BestPoints[i] != nil {
			if err := m.BestPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bestPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserBestPoints) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserBestPoints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserBestPoints) UnmarshalBinary(b []byte) error {
	var res UserBestPoints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}