// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AtzmTimeSlotWithTimeZone atzm time slot with time zone
//
// swagger:model AtzmTimeSlotWithTimeZone
type AtzmTimeSlotWithTimeZone struct {

	// The earliest time to dial a contact. Valid format is HH:mm
	EarliestCallableTime string `json:"earliestCallableTime,omitempty"`

	// The latest time to dial a contact. Valid format is HH:mm
	LatestCallableTime string `json:"latestCallableTime,omitempty"`

	// The time zone to use for contacts that cannot be mapped.
	TimeZoneID string `json:"timeZoneId,omitempty"`
}

// Validate validates this atzm time slot with time zone
func (m *AtzmTimeSlotWithTimeZone) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AtzmTimeSlotWithTimeZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AtzmTimeSlotWithTimeZone) UnmarshalBinary(b []byte) error {
	var res AtzmTimeSlotWithTimeZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}