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

// AvailableTime available time
//
// swagger:model AvailableTime
type AvailableTime struct {

	// Workforce Management activity category for this availability period
	// Read Only: true
	// Enum: [OnQueueWork Break Meal Meeting OffQueueWork TimeOff Training Unavailable]
	ActivityCategory string `json:"activityCategory,omitempty"`

	// Start of the availability period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateStart strfmt.DateTime `json:"dateStart,omitempty"`

	// Indicates if this availability period is paid in Workforce Management schedule
	// Read Only: true
	IsPaid *bool `json:"isPaid"`

	// Length of availability period in minutes
	// Read Only: true
	LengthInMinutes int32 `json:"lengthInMinutes,omitempty"`

	// Workforce Management schedule information associated with the available time
	// Read Only: true
	WfmSchedule *WfmScheduleReference `json:"wfmSchedule,omitempty"`
}

// Validate validates this available time
func (m *AvailableTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStart(formats); err != nil {
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

var availableTimeTypeActivityCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OnQueueWork","Break","Meal","Meeting","OffQueueWork","TimeOff","Training","Unavailable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		availableTimeTypeActivityCategoryPropEnum = append(availableTimeTypeActivityCategoryPropEnum, v)
	}
}

const (

	// AvailableTimeActivityCategoryOnQueueWork captures enum value "OnQueueWork"
	AvailableTimeActivityCategoryOnQueueWork string = "OnQueueWork"

	// AvailableTimeActivityCategoryBreak captures enum value "Break"
	AvailableTimeActivityCategoryBreak string = "Break"

	// AvailableTimeActivityCategoryMeal captures enum value "Meal"
	AvailableTimeActivityCategoryMeal string = "Meal"

	// AvailableTimeActivityCategoryMeeting captures enum value "Meeting"
	AvailableTimeActivityCategoryMeeting string = "Meeting"

	// AvailableTimeActivityCategoryOffQueueWork captures enum value "OffQueueWork"
	AvailableTimeActivityCategoryOffQueueWork string = "OffQueueWork"

	// AvailableTimeActivityCategoryTimeOff captures enum value "TimeOff"
	AvailableTimeActivityCategoryTimeOff string = "TimeOff"

	// AvailableTimeActivityCategoryTraining captures enum value "Training"
	AvailableTimeActivityCategoryTraining string = "Training"

	// AvailableTimeActivityCategoryUnavailable captures enum value "Unavailable"
	AvailableTimeActivityCategoryUnavailable string = "Unavailable"
)

// prop value enum
func (m *AvailableTime) validateActivityCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, availableTimeTypeActivityCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AvailableTime) validateActivityCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateActivityCategoryEnum("activityCategory", "body", m.ActivityCategory); err != nil {
		return err
	}

	return nil
}

func (m *AvailableTime) validateDateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStart) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStart", "body", "date-time", m.DateStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AvailableTime) validateWfmSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.WfmSchedule) { // not required
		return nil
	}

	if m.WfmSchedule != nil {
		if err := m.WfmSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wfmSchedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailableTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableTime) UnmarshalBinary(b []byte) error {
	var res AvailableTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
