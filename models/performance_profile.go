// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PerformanceProfile performance profile
//
// swagger:model PerformanceProfile
type PerformanceProfile struct {

	// The flag for active profiles
	// Read Only: true
	Active *bool `json:"active"`

	// Creation date for this performance profile. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// A description about this performance profile
	// Required: true
	Description *string `json:"description"`

	// The division for this performance profile associate to
	// Read Only: true
	Division *Division `json:"division,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The maximum rank size for the leaderboard. This counts the number of ranks can be retrieved in a leaderboard queries
	MaxLeaderboardRankSize int32 `json:"maxLeaderboardRankSize,omitempty"`

	// The number of members in this performance profile
	// Read Only: true
	MemberCount int32 `json:"memberCount,omitempty"`

	// Order of the associated metrics. The list should contain valid ids for metrics
	// Required: true
	MetricOrders []string `json:"metricOrders"`

	// A name for this performance profile
	// Required: true
	Name *string `json:"name"`

	// The reporting interval periods for this performance profile
	ReportingIntervals []*ReportingInterval `json:"reportingIntervals"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this performance profile
func (m *PerformanceProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricOrders(formats); err != nil {
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

func (m *PerformanceProfile) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceProfile) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceProfile) validateDivision(formats strfmt.Registry) error {
	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceProfile) validateMetricOrders(formats strfmt.Registry) error {

	if err := validate.Required("metricOrders", "body", m.MetricOrders); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceProfile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceProfile) validateReportingIntervals(formats strfmt.Registry) error {
	if swag.IsZero(m.ReportingIntervals) { // not required
		return nil
	}

	for i := 0; i < len(m.ReportingIntervals); i++ {
		if swag.IsZero(m.ReportingIntervals[i]) { // not required
			continue
		}

		if m.ReportingIntervals[i] != nil {
			if err := m.ReportingIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportingIntervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reportingIntervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PerformanceProfile) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance profile based on the context it is used
func (m *PerformanceProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDivision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemberCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReportingIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceProfile) contextValidateActive(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceProfile) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceProfile) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

	if m.Division != nil {
		if err := m.Division.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceProfile) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceProfile) contextValidateMemberCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "memberCount", "body", int32(m.MemberCount)); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceProfile) contextValidateReportingIntervals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReportingIntervals); i++ {

		if m.ReportingIntervals[i] != nil {
			if err := m.ReportingIntervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportingIntervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reportingIntervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PerformanceProfile) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceProfile) UnmarshalBinary(b []byte) error {
	var res PerformanceProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
