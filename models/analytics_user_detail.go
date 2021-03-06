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

// AnalyticsUserDetail analytics user detail
//
// swagger:model AnalyticsUserDetail
type AnalyticsUserDetail struct {

	// The presence records for the user
	PrimaryPresence []*AnalyticsUserPresenceRecord `json:"primaryPresence"`

	// The ACD routing status records for the user
	RoutingStatus []*AnalyticsRoutingStatusRecord `json:"routingStatus"`

	// The identifier for the user
	UserID string `json:"userId,omitempty"`
}

// Validate validates this analytics user detail
func (m *AnalyticsUserDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrimaryPresence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsUserDetail) validatePrimaryPresence(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryPresence) { // not required
		return nil
	}

	for i := 0; i < len(m.PrimaryPresence); i++ {
		if swag.IsZero(m.PrimaryPresence[i]) { // not required
			continue
		}

		if m.PrimaryPresence[i] != nil {
			if err := m.PrimaryPresence[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("primaryPresence" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnalyticsUserDetail) validateRoutingStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.RoutingStatus); i++ {
		if swag.IsZero(m.RoutingStatus[i]) { // not required
			continue
		}

		if m.RoutingStatus[i] != nil {
			if err := m.RoutingStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routingStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsUserDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsUserDetail) UnmarshalBinary(b []byte) error {
	var res AnalyticsUserDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
