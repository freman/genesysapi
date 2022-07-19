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

// UserSchedule user schedule
//
// swagger:model UserSchedule
type UserSchedule struct {

	// If marked true for updating an existing user schedule, it will be deleted
	Delete bool `json:"delete"`

	// Markers to indicate a full day time off request, relative to the management unit time zone
	FullDayTimeOffMarkers []*UserScheduleFullDayTimeOffMarker `json:"fullDayTimeOffMarkers"`

	// Version metadata for this schedule
	// Required: true
	Metadata *WfmVersionedEntityMetadata `json:"metadata"`

	// The shifts that belong to this schedule
	Shifts []*UserScheduleShift `json:"shifts"`

	// ID of the work plan associated with the user during schedule creation
	// Read Only: true
	WorkPlanID string `json:"workPlanId,omitempty"`
}

// Validate validates this user schedule
func (m *UserSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullDayTimeOffMarkers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShifts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSchedule) validateFullDayTimeOffMarkers(formats strfmt.Registry) error {

	if swag.IsZero(m.FullDayTimeOffMarkers) { // not required
		return nil
	}

	for i := 0; i < len(m.FullDayTimeOffMarkers); i++ {
		if swag.IsZero(m.FullDayTimeOffMarkers[i]) { // not required
			continue
		}

		if m.FullDayTimeOffMarkers[i] != nil {
			if err := m.FullDayTimeOffMarkers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fullDayTimeOffMarkers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserSchedule) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *UserSchedule) validateShifts(formats strfmt.Registry) error {

	if swag.IsZero(m.Shifts) { // not required
		return nil
	}

	for i := 0; i < len(m.Shifts); i++ {
		if swag.IsZero(m.Shifts[i]) { // not required
			continue
		}

		if m.Shifts[i] != nil {
			if err := m.Shifts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shifts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSchedule) UnmarshalBinary(b []byte) error {
	var res UserSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
