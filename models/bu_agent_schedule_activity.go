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

// BuAgentScheduleActivity bu agent schedule activity
//
// swagger:model BuAgentScheduleActivity
type BuAgentScheduleActivity struct {

	// The ID of the activity code associated with this activity
	ActivityCodeID string `json:"activityCodeId,omitempty"`

	// The description of this activity
	Description string `json:"description,omitempty"`

	// The ID of the external activity associated with this activity, if applicable
	ExternalActivityID string `json:"externalActivityId,omitempty"`

	// The type of the external activity associated with this activity, if applicable
	// Enum: [Coaching]
	ExternalActivityType string `json:"externalActivityType,omitempty"`

	// The length of this activity in minutes
	LengthMinutes int32 `json:"lengthMinutes,omitempty"`

	// Whether this activity is paid
	Paid bool `json:"paid,omitempty"`

	// The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// The ID of the time off request associated with this activity, if applicable
	TimeOffRequestID string `json:"timeOffRequestId,omitempty"`
}

// Validate validates this bu agent schedule activity
func (m *BuAgentScheduleActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalActivityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var buAgentScheduleActivityTypeExternalActivityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Coaching"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buAgentScheduleActivityTypeExternalActivityTypePropEnum = append(buAgentScheduleActivityTypeExternalActivityTypePropEnum, v)
	}
}

const (

	// BuAgentScheduleActivityExternalActivityTypeCoaching captures enum value "Coaching"
	BuAgentScheduleActivityExternalActivityTypeCoaching string = "Coaching"
)

// prop value enum
func (m *BuAgentScheduleActivity) validateExternalActivityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buAgentScheduleActivityTypeExternalActivityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuAgentScheduleActivity) validateExternalActivityType(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalActivityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExternalActivityTypeEnum("externalActivityType", "body", m.ExternalActivityType); err != nil {
		return err
	}

	return nil
}

func (m *BuAgentScheduleActivity) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuAgentScheduleActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuAgentScheduleActivity) UnmarshalBinary(b []byte) error {
	var res BuAgentScheduleActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
