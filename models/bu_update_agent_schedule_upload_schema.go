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

// BuUpdateAgentScheduleUploadSchema bu update agent schedule upload schema
//
// swagger:model BuUpdateAgentScheduleUploadSchema
type BuUpdateAgentScheduleUploadSchema struct {

	// Whether to delete this agent's schedule. Defaults to false if not set
	Delete bool `json:"delete"`

	// Any full day time off markers that apply to this agent schedule
	FullDayTimeOffMarkers []*BuFullDayTimeOffMarker `json:"fullDayTimeOffMarkers"`

	// Version metadata for this agent schedule. Required if updating or deleting an existing agent schedule, otherwise should be omitted
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// The shift definitions for this agent schedule
	Shifts []*BuAgentScheduleShift `json:"shifts"`

	// The ID of the user to whom this agent schedule applies
	// Required: true
	UserID *string `json:"userId"`

	// The ID of the work plan for this user.  Mutually exclusive with workPlanIdsPerWeek
	WorkPlanID *ValueWrapperString `json:"workPlanId,omitempty"`

	// The IDs of the work plans per week for this user.  Mutually exclusive with workPlanId
	WorkPlanIdsPerWeek *ListWrapperString `json:"workPlanIdsPerWeek,omitempty"`
}

// Validate validates this bu update agent schedule upload schema
func (m *BuUpdateAgentScheduleUploadSchema) Validate(formats strfmt.Registry) error {
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

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkPlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkPlanIdsPerWeek(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuUpdateAgentScheduleUploadSchema) validateFullDayTimeOffMarkers(formats strfmt.Registry) error {

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

func (m *BuUpdateAgentScheduleUploadSchema) validateMetadata(formats strfmt.Registry) error {

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

func (m *BuUpdateAgentScheduleUploadSchema) validateShifts(formats strfmt.Registry) error {

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

func (m *BuUpdateAgentScheduleUploadSchema) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *BuUpdateAgentScheduleUploadSchema) validateWorkPlanID(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkPlanID) { // not required
		return nil
	}

	if m.WorkPlanID != nil {
		if err := m.WorkPlanID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workPlanId")
			}
			return err
		}
	}

	return nil
}

func (m *BuUpdateAgentScheduleUploadSchema) validateWorkPlanIdsPerWeek(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkPlanIdsPerWeek) { // not required
		return nil
	}

	if m.WorkPlanIdsPerWeek != nil {
		if err := m.WorkPlanIdsPerWeek.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workPlanIdsPerWeek")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuUpdateAgentScheduleUploadSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuUpdateAgentScheduleUploadSchema) UnmarshalBinary(b []byte) error {
	var res BuUpdateAgentScheduleUploadSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
