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

// AvailableTimeOffRequest available time off request
//
// swagger:model AvailableTimeOffRequest
type AvailableTimeOffRequest struct {

	// The ID for activity code to query available time off minutes
	// Required: true
	ActivityCodeID *string `json:"activityCodeId"`

	// A list of date ranges of available time off minutes. A maximum number of date ranges is 30. The maximum total number of days in all ranges is 366. If no ranges are specified, then only the presence of the associated time off limit object will be checked. In such case, if the association exists, then the response will contain a list with of a single element filled with timeOffLimitId only.
	DateRanges []*LocalDateRange `json:"dateRanges"`
}

// Validate validates this available time off request
func (m *AvailableTimeOffRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityCodeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailableTimeOffRequest) validateActivityCodeID(formats strfmt.Registry) error {

	if err := validate.Required("activityCodeId", "body", m.ActivityCodeID); err != nil {
		return err
	}

	return nil
}

func (m *AvailableTimeOffRequest) validateDateRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.DateRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.DateRanges); i++ {
		if swag.IsZero(m.DateRanges[i]) { // not required
			continue
		}

		if m.DateRanges[i] != nil {
			if err := m.DateRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dateRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dateRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this available time off request based on the context it is used
func (m *AvailableTimeOffRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailableTimeOffRequest) contextValidateDateRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DateRanges); i++ {

		if m.DateRanges[i] != nil {
			if err := m.DateRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dateRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dateRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailableTimeOffRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableTimeOffRequest) UnmarshalBinary(b []byte) error {
	var res AvailableTimeOffRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
