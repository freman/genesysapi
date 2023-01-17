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

// QueryTimeOffLimitValuesRequest query time off limit values request
//
// swagger:model QueryTimeOffLimitValuesRequest
type QueryTimeOffLimitValuesRequest struct {

	// The activity code id to filter the affected limit objects by. Required if timeOffLimitId is not specified
	ActivityCodeID string `json:"activityCodeId,omitempty"`

	// The list of the date ranges to return time off limit, allocated and waitlisted minutes. The valid number of date ranges is between 1 and 30. Maximum total number of days in all ranges in 366.
	// Required: true
	DateRanges []*LocalDateRange `json:"dateRanges"`

	// The time off limit object id to retrieve values for. Required if activityCodeId is not specified
	TimeOffLimitID string `json:"timeOffLimitId,omitempty"`
}

// Validate validates this query time off limit values request
func (m *QueryTimeOffLimitValuesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryTimeOffLimitValuesRequest) validateDateRanges(formats strfmt.Registry) error {

	if err := validate.Required("dateRanges", "body", m.DateRanges); err != nil {
		return err
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

// ContextValidate validate this query time off limit values request based on the context it is used
func (m *QueryTimeOffLimitValuesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryTimeOffLimitValuesRequest) contextValidateDateRanges(ctx context.Context, formats strfmt.Registry) error {

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
func (m *QueryTimeOffLimitValuesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryTimeOffLimitValuesRequest) UnmarshalBinary(b []byte) error {
	var res QueryTimeOffLimitValuesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
