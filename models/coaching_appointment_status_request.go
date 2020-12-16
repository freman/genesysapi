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

// CoachingAppointmentStatusRequest coaching appointment status request
//
// swagger:model CoachingAppointmentStatusRequest
type CoachingAppointmentStatusRequest struct {

	// The status of the coaching appointment
	// Required: true
	// Enum: [Scheduled InProgress Completed]
	Status *string `json:"status"`
}

// Validate validates this coaching appointment status request
func (m *CoachingAppointmentStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var coachingAppointmentStatusRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Scheduled","InProgress","Completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		coachingAppointmentStatusRequestTypeStatusPropEnum = append(coachingAppointmentStatusRequestTypeStatusPropEnum, v)
	}
}

const (

	// CoachingAppointmentStatusRequestStatusScheduled captures enum value "Scheduled"
	CoachingAppointmentStatusRequestStatusScheduled string = "Scheduled"

	// CoachingAppointmentStatusRequestStatusInProgress captures enum value "InProgress"
	CoachingAppointmentStatusRequestStatusInProgress string = "InProgress"

	// CoachingAppointmentStatusRequestStatusCompleted captures enum value "Completed"
	CoachingAppointmentStatusRequestStatusCompleted string = "Completed"
)

// prop value enum
func (m *CoachingAppointmentStatusRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, coachingAppointmentStatusRequestTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CoachingAppointmentStatusRequest) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CoachingAppointmentStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoachingAppointmentStatusRequest) UnmarshalBinary(b []byte) error {
	var res CoachingAppointmentStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
