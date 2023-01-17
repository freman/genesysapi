// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserScheduleFullDayTimeOffMarker user schedule full day time off marker
//
// swagger:model UserScheduleFullDayTimeOffMarker
type UserScheduleFullDayTimeOffMarker struct {

	// The id for the activity code.  Look up a map of activity codes with the activities route
	ActivityCodeID string `json:"activityCodeId,omitempty"`

	// If marked true for updating an existing full day time off marker, it will be deleted
	Delete bool `json:"delete"`

	// The description associated with the time off request that this marker corresponds to
	Description string `json:"description,omitempty"`

	// Whether this is paid time off
	IsPaid bool `json:"isPaid"`

	// The length in minutes of this time off marker
	LengthInMinutes int32 `json:"lengthInMinutes,omitempty"`

	// The date associated with the time off request that this marker corresponds to.  Date only, in ISO-8601 format.
	ManagementUnitDate string `json:"managementUnitDate,omitempty"`
}

// Validate validates this user schedule full day time off marker
func (m *UserScheduleFullDayTimeOffMarker) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user schedule full day time off marker based on context it is used
func (m *UserScheduleFullDayTimeOffMarker) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserScheduleFullDayTimeOffMarker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserScheduleFullDayTimeOffMarker) UnmarshalBinary(b []byte) error {
	var res UserScheduleFullDayTimeOffMarker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
