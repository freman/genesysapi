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

// BuQueryAgentSchedulesRequest bu query agent schedules request
//
// swagger:model BuQueryAgentSchedulesRequest
type BuQueryAgentSchedulesRequest struct {

	// The ID of the management unit to query
	// Required: true
	ManagementUnitID *string `json:"managementUnitId"`

	// The teamIds to report on. If null or not set, results will be queried for requested users if applicable or otherwise all users in the management unit
	// Unique: true
	TeamIds []string `json:"teamIds"`

	// The IDs of the users to query.  Omit to query all user schedules in the management unit. Note: If teamIds is also specified, only schedules for users in the requested teams will be returned
	// Unique: true
	UserIds []string `json:"userIds"`
}

// Validate validates this bu query agent schedules request
func (m *BuQueryAgentSchedulesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagementUnitID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuQueryAgentSchedulesRequest) validateManagementUnitID(formats strfmt.Registry) error {

	if err := validate.Required("managementUnitId", "body", m.ManagementUnitID); err != nil {
		return err
	}

	return nil
}

func (m *BuQueryAgentSchedulesRequest) validateTeamIds(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("teamIds", "body", m.TeamIds); err != nil {
		return err
	}

	return nil
}

func (m *BuQueryAgentSchedulesRequest) validateUserIds(formats strfmt.Registry) error {

	if swag.IsZero(m.UserIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("userIds", "body", m.UserIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuQueryAgentSchedulesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuQueryAgentSchedulesRequest) UnmarshalBinary(b []byte) error {
	var res BuQueryAgentSchedulesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
