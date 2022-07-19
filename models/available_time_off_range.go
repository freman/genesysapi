// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AvailableTimeOffRange available time off range
//
// swagger:model AvailableTimeOffRange
type AvailableTimeOffRange struct {

	// The list of available time off values in minutes per granularity interval
	AvailableMinutesPerInterval []int32 `json:"availableMinutesPerInterval"`

	// Granularity choice for time off limit
	// Enum: [Daily]
	Granularity string `json:"granularity,omitempty"`

	// Start date of the requested date range. The end date is determined by the size of interval list. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	StartDate strfmt.Date `json:"startDate,omitempty"`

	// The time off limit
	TimeOffLimit *TimeOffLimitReference `json:"timeOffLimit,omitempty"`

	// Whether the time off request can be waitlisted
	WaitlistEnabled bool `json:"waitlistEnabled"`

	// The current number of waitlisted time off requests for every interval per granularity
	WaitlistedRequestsPerInterval []int32 `json:"waitlistedRequestsPerInterval"`
}

// Validate validates this available time off range
func (m *AvailableTimeOffRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGranularity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOffLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var availableTimeOffRangeTypeGranularityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Daily"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		availableTimeOffRangeTypeGranularityPropEnum = append(availableTimeOffRangeTypeGranularityPropEnum, v)
	}
}

const (

	// AvailableTimeOffRangeGranularityDaily captures enum value "Daily"
	AvailableTimeOffRangeGranularityDaily string = "Daily"
)

// prop value enum
func (m *AvailableTimeOffRange) validateGranularityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, availableTimeOffRangeTypeGranularityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AvailableTimeOffRange) validateGranularity(formats strfmt.Registry) error {

	if swag.IsZero(m.Granularity) { // not required
		return nil
	}

	// value enum
	if err := m.validateGranularityEnum("granularity", "body", m.Granularity); err != nil {
		return err
	}

	return nil
}

func (m *AvailableTimeOffRange) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AvailableTimeOffRange) validateTimeOffLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeOffLimit) { // not required
		return nil
	}

	if m.TimeOffLimit != nil {
		if err := m.TimeOffLimit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeOffLimit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailableTimeOffRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableTimeOffRange) UnmarshalBinary(b []byte) error {
	var res AvailableTimeOffRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
