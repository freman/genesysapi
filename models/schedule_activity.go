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

// ScheduleActivity schedule activity
//
// swagger:model ScheduleActivity
type ScheduleActivity struct {

	// The ID of the activity code associated with this activity
	ActivityCodeID string `json:"activityCodeId,omitempty"`

	// The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateStart strfmt.DateTime `json:"dateStart,omitempty"`

	// The description of this activity
	Description string `json:"description,omitempty"`

	// The ID of the external activity associated with this activity, if applicable
	ExternalActivityID string `json:"externalActivityId,omitempty"`

	// The type of the external activity associated with this activity, if applicable
	// Enum: [Coaching Learning]
	ExternalActivityType string `json:"externalActivityType,omitempty"`

	// The length of this activity in minutes
	LengthMinutes int32 `json:"lengthMinutes,omitempty"`

	// Whether this activity is paid
	Paid bool `json:"paid"`

	// The ID of the time off request associated with this activity, if applicable
	TimeOffRequestID string `json:"timeOffRequestId,omitempty"`
}

// Validate validates this schedule activity
func (m *ScheduleActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalActivityType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleActivity) validateDateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.DateStart) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStart", "body", "date-time", m.DateStart.String(), formats); err != nil {
		return err
	}

	return nil
}

var scheduleActivityTypeExternalActivityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Coaching","Learning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleActivityTypeExternalActivityTypePropEnum = append(scheduleActivityTypeExternalActivityTypePropEnum, v)
	}
}

const (

	// ScheduleActivityExternalActivityTypeCoaching captures enum value "Coaching"
	ScheduleActivityExternalActivityTypeCoaching string = "Coaching"

	// ScheduleActivityExternalActivityTypeLearning captures enum value "Learning"
	ScheduleActivityExternalActivityTypeLearning string = "Learning"
)

// prop value enum
func (m *ScheduleActivity) validateExternalActivityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduleActivityTypeExternalActivityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScheduleActivity) validateExternalActivityType(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalActivityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExternalActivityTypeEnum("externalActivityType", "body", m.ExternalActivityType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this schedule activity based on context it is used
func (m *ScheduleActivity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleActivity) UnmarshalBinary(b []byte) error {
	var res ScheduleActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
