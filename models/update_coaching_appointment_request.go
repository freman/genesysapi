// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateCoachingAppointmentRequest Update coaching appointment request
//
// swagger:model UpdateCoachingAppointmentRequest
type UpdateCoachingAppointmentRequest struct {

	// IDs of conversations associated with this coaching appointment.
	// Unique: true
	ConversationIds []string `json:"conversationIds"`

	// The date/time the coaching appointment starts. Times will be rounded down to the minute. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateStart strfmt.DateTime `json:"dateStart,omitempty"`

	// The description of coaching appointment.
	Description string `json:"description,omitempty"`

	// IDs of documents associated with this coaching appointment.
	// Unique: true
	DocumentIds []string `json:"documentIds"`

	// The list of external links related to the appointment
	ExternalLinks []string `json:"externalLinks"`

	// The duration of coaching appointment in minutes.
	LengthInMinutes int32 `json:"lengthInMinutes,omitempty"`

	// The name of coaching appointment.
	Name string `json:"name,omitempty"`

	// The status of the coaching appointment.
	// Enum: [Scheduled InProgress Completed]
	Status string `json:"status,omitempty"`

	// The Workforce Management schedule the appointment is associated with.
	WfmSchedule *WfmScheduleReference `json:"wfmSchedule,omitempty"`
}

// Validate validates this update coaching appointment request
func (m *UpdateCoachingAppointmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversationIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocumentIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWfmSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCoachingAppointmentRequest) validateConversationIds(formats strfmt.Registry) error {
	if swag.IsZero(m.ConversationIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("conversationIds", "body", m.ConversationIds); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCoachingAppointmentRequest) validateDateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.DateStart) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStart", "body", "date-time", m.DateStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCoachingAppointmentRequest) validateDocumentIds(formats strfmt.Registry) error {
	if swag.IsZero(m.DocumentIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("documentIds", "body", m.DocumentIds); err != nil {
		return err
	}

	return nil
}

var updateCoachingAppointmentRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Scheduled","InProgress","Completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateCoachingAppointmentRequestTypeStatusPropEnum = append(updateCoachingAppointmentRequestTypeStatusPropEnum, v)
	}
}

const (

	// UpdateCoachingAppointmentRequestStatusScheduled captures enum value "Scheduled"
	UpdateCoachingAppointmentRequestStatusScheduled string = "Scheduled"

	// UpdateCoachingAppointmentRequestStatusInProgress captures enum value "InProgress"
	UpdateCoachingAppointmentRequestStatusInProgress string = "InProgress"

	// UpdateCoachingAppointmentRequestStatusCompleted captures enum value "Completed"
	UpdateCoachingAppointmentRequestStatusCompleted string = "Completed"
)

// prop value enum
func (m *UpdateCoachingAppointmentRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateCoachingAppointmentRequestTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateCoachingAppointmentRequest) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCoachingAppointmentRequest) validateWfmSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.WfmSchedule) { // not required
		return nil
	}

	if m.WfmSchedule != nil {
		if err := m.WfmSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wfmSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wfmSchedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update coaching appointment request based on the context it is used
func (m *UpdateCoachingAppointmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWfmSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCoachingAppointmentRequest) contextValidateWfmSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.WfmSchedule != nil {
		if err := m.WfmSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wfmSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wfmSchedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCoachingAppointmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCoachingAppointmentRequest) UnmarshalBinary(b []byte) error {
	var res UpdateCoachingAppointmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
