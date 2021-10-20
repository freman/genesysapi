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

// CreatePerformanceProfile create performance profile
//
// swagger:model CreatePerformanceProfile
type CreatePerformanceProfile struct {

	// The flag for active profiles
	// Required: true
	Active *bool `json:"active"`

	// Creation date for this performance profile. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// A description about this performance profile
	// Required: true
	Description *string `json:"description"`

	// The associated division for this Performance Profile
	// Required: true
	Division *WritableDivision `json:"division"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The maximum rank size for the leaderboard. This counts the number of ranks can be retrieved in a leaderboard queries
	// Required: true
	MaxLeaderboardRankSize *int32 `json:"maxLeaderboardRankSize"`

	// The number of members in this performance profile
	// Read Only: true
	MemberCount int32 `json:"memberCount,omitempty"`

	// Order of the associated metrics. The list should contain valid ids for metrics
	// Read Only: true
	MetricOrders []string `json:"metricOrders"`

	// A name for this performance profile
	// Required: true
	Name *string `json:"name"`

	// The reporting interval periods for this performance profile
	// Required: true
	ReportingIntervals []*ReportingInterval `json:"reportingIntervals"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this create performance profile
func (m *CreatePerformanceProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxLeaderboardRankSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportingIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePerformanceProfile) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *CreatePerformanceProfile) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreatePerformanceProfile) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CreatePerformanceProfile) validateDivision(formats strfmt.Registry) error {

	if err := validate.Required("division", "body", m.Division); err != nil {
		return err
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

func (m *CreatePerformanceProfile) validateMaxLeaderboardRankSize(formats strfmt.Registry) error {

	if err := validate.Required("maxLeaderboardRankSize", "body", m.MaxLeaderboardRankSize); err != nil {
		return err
	}

	return nil
}

func (m *CreatePerformanceProfile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreatePerformanceProfile) validateReportingIntervals(formats strfmt.Registry) error {

	if err := validate.Required("reportingIntervals", "body", m.ReportingIntervals); err != nil {
		return err
	}

	for i := 0; i < len(m.ReportingIntervals); i++ {
		if swag.IsZero(m.ReportingIntervals[i]) { // not required
			continue
		}

		if m.ReportingIntervals[i] != nil {
			if err := m.ReportingIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportingIntervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreatePerformanceProfile) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePerformanceProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePerformanceProfile) UnmarshalBinary(b []byte) error {
	var res CreatePerformanceProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
