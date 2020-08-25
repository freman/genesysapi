// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuAgentScheduleQueryResponse bu agent schedule query response
//
// swagger:model BuAgentScheduleQueryResponse
type BuAgentScheduleQueryResponse struct {

	// Full day time off markers which apply to this agent schedule
	FullDayTimeOffMarkers []*BuFullDayTimeOffMarker `json:"fullDayTimeOffMarkers"`

	// Versioned entity metadata for this agent schedule
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// The shift definitions for this agent schedule
	Shifts []*BuAgentScheduleShift `json:"shifts"`

	// The user to whom this agent schedule applies
	User *UserReference `json:"user,omitempty"`

	// The work plan for this user
	WorkPlan *WorkPlanReference `json:"workPlan,omitempty"`
}

// Validate validates this bu agent schedule query response
func (m *BuAgentScheduleQueryResponse) Validate(formats strfmt.Registry) error {
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

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkPlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAgentScheduleQueryResponse) validateFullDayTimeOffMarkers(formats strfmt.Registry) error {

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

func (m *BuAgentScheduleQueryResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
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

func (m *BuAgentScheduleQueryResponse) validateShifts(formats strfmt.Registry) error {

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

func (m *BuAgentScheduleQueryResponse) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *BuAgentScheduleQueryResponse) validateWorkPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkPlan) { // not required
		return nil
	}

	if m.WorkPlan != nil {
		if err := m.WorkPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workPlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuAgentScheduleQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuAgentScheduleQueryResponse) UnmarshalBinary(b []byte) error {
	var res BuAgentScheduleQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}