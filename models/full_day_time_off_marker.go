// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FullDayTimeOffMarker full day time off marker
//
// swagger:model FullDayTimeOffMarker
type FullDayTimeOffMarker struct {

	// The ID of the activity code associated with the time off marker
	ActivityCodeID string `json:"activityCodeId,omitempty"`

	// The date of the time off marker, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	BusinessUnitDate strfmt.Date `json:"businessUnitDate,omitempty"`

	// The description of the time off marker
	Description string `json:"description,omitempty"`

	// The length of the time off marker in minutes
	LengthMinutes int32 `json:"lengthMinutes,omitempty"`

	// Whether the time off marker is paid
	Paid bool `json:"paid"`

	// The ID of the time off request
	TimeOffRequestID string `json:"timeOffRequestId,omitempty"`
}

// Validate validates this full day time off marker
func (m *FullDayTimeOffMarker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessUnitDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullDayTimeOffMarker) validateBusinessUnitDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BusinessUnitDate) { // not required
		return nil
	}

	if err := validate.FormatOf("businessUnitDate", "body", "date", m.BusinessUnitDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this full day time off marker based on context it is used
func (m *FullDayTimeOffMarker) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FullDayTimeOffMarker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullDayTimeOffMarker) UnmarshalBinary(b []byte) error {
	var res FullDayTimeOffMarker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
