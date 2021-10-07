// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Leaderboard leaderboard
//
// swagger:model Leaderboard
type Leaderboard struct {

	// End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Read Only: true
	// Format: date
	DateEndWorkday strfmt.Date `json:"dateEndWorkday,omitempty"`

	// Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Read Only: true
	// Format: date
	DateStartWorkday strfmt.Date `json:"dateStartWorkday,omitempty"`

	// The targeted division for this leaderboard
	// Read Only: true
	Division *Division `json:"division,omitempty"`

	// The list of leaders generated.
	// Read Only: true
	Leaders []*LeaderboardItem `json:"leaders"`

	// The metric id if the leaderboard is about a specific metric
	// Read Only: true
	Metric *AddressableEntityRef `json:"metric,omitempty"`

	// The requesting user's rank
	// Read Only: true
	UserRank *LeaderboardItem `json:"userRank,omitempty"`
}

// Validate validates this leaderboard
func (m *Leaderboard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateEndWorkday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStartWorkday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserRank(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Leaderboard) validateDateEndWorkday(formats strfmt.Registry) error {

	if swag.IsZero(m.DateEndWorkday) { // not required
		return nil
	}

	if err := validate.FormatOf("dateEndWorkday", "body", "date", m.DateEndWorkday.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Leaderboard) validateDateStartWorkday(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStartWorkday) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStartWorkday", "body", "date", m.DateStartWorkday.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Leaderboard) validateDivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *Leaderboard) validateLeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.Leaders) { // not required
		return nil
	}

	for i := 0; i < len(m.Leaders); i++ {
		if swag.IsZero(m.Leaders[i]) { // not required
			continue
		}

		if m.Leaders[i] != nil {
			if err := m.Leaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("leaders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Leaderboard) validateMetric(formats strfmt.Registry) error {

	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *Leaderboard) validateUserRank(formats strfmt.Registry) error {

	if swag.IsZero(m.UserRank) { // not required
		return nil
	}

	if m.UserRank != nil {
		if err := m.UserRank.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userRank")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Leaderboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Leaderboard) UnmarshalBinary(b []byte) error {
	var res Leaderboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
