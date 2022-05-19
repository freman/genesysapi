// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GamificationStatus gamification status
//
// swagger:model GamificationStatus
type GamificationStatus struct {

	// Automatic assignment of users to the default profile
	AutomaticUserAssignment bool `json:"automaticUserAssignment"`

	// Gamification start date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	DateStart strfmt.Date `json:"dateStart,omitempty"`

	// Personal best aggregation starting date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	DateStartPersonalBest strfmt.Date `json:"dateStartPersonalBest,omitempty"`

	// Gamification status of the organization.
	IsActive bool `json:"isActive"`
}

// Validate validates this gamification status
func (m *GamificationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStartPersonalBest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GamificationStatus) validateDateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStart) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStart", "body", "date", m.DateStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GamificationStatus) validateDateStartPersonalBest(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStartPersonalBest) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStartPersonalBest", "body", "date", m.DateStartPersonalBest.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GamificationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GamificationStatus) UnmarshalBinary(b []byte) error {
	var res GamificationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
