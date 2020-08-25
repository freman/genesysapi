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

// ShiftTradeActivityRule shift trade activity rule
//
// swagger:model ShiftTradeActivityRule
type ShiftTradeActivityRule struct {

	// The action this rule invokes
	// Required: true
	// Enum: [Replace DoNotAllowTrade KeepWithSchedule]
	Action *string `json:"action"`

	// The activity category to which to apply this rule
	// Required: true
	// Enum: [OnQueueWork Break Meal Meeting OffQueueWork TimeOff Training Unavailable Unscheduled]
	ActivityCategory *string `json:"activityCategory"`

	// The activity code ID with which to replace activities belonging to the original category if applicable (required if action == Replace, must be a default activity code ID)
	ActivityCodeIDReplacement string `json:"activityCodeIdReplacement,omitempty"`
}

// Validate validates this shift trade activity rule
func (m *ShiftTradeActivityRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivityCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var shiftTradeActivityRuleTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Replace","DoNotAllowTrade","KeepWithSchedule"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shiftTradeActivityRuleTypeActionPropEnum = append(shiftTradeActivityRuleTypeActionPropEnum, v)
	}
}

const (

	// ShiftTradeActivityRuleActionReplace captures enum value "Replace"
	ShiftTradeActivityRuleActionReplace string = "Replace"

	// ShiftTradeActivityRuleActionDoNotAllowTrade captures enum value "DoNotAllowTrade"
	ShiftTradeActivityRuleActionDoNotAllowTrade string = "DoNotAllowTrade"

	// ShiftTradeActivityRuleActionKeepWithSchedule captures enum value "KeepWithSchedule"
	ShiftTradeActivityRuleActionKeepWithSchedule string = "KeepWithSchedule"
)

// prop value enum
func (m *ShiftTradeActivityRule) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, shiftTradeActivityRuleTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ShiftTradeActivityRule) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

var shiftTradeActivityRuleTypeActivityCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OnQueueWork","Break","Meal","Meeting","OffQueueWork","TimeOff","Training","Unavailable","Unscheduled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shiftTradeActivityRuleTypeActivityCategoryPropEnum = append(shiftTradeActivityRuleTypeActivityCategoryPropEnum, v)
	}
}

const (

	// ShiftTradeActivityRuleActivityCategoryOnQueueWork captures enum value "OnQueueWork"
	ShiftTradeActivityRuleActivityCategoryOnQueueWork string = "OnQueueWork"

	// ShiftTradeActivityRuleActivityCategoryBreak captures enum value "Break"
	ShiftTradeActivityRuleActivityCategoryBreak string = "Break"

	// ShiftTradeActivityRuleActivityCategoryMeal captures enum value "Meal"
	ShiftTradeActivityRuleActivityCategoryMeal string = "Meal"

	// ShiftTradeActivityRuleActivityCategoryMeeting captures enum value "Meeting"
	ShiftTradeActivityRuleActivityCategoryMeeting string = "Meeting"

	// ShiftTradeActivityRuleActivityCategoryOffQueueWork captures enum value "OffQueueWork"
	ShiftTradeActivityRuleActivityCategoryOffQueueWork string = "OffQueueWork"

	// ShiftTradeActivityRuleActivityCategoryTimeOff captures enum value "TimeOff"
	ShiftTradeActivityRuleActivityCategoryTimeOff string = "TimeOff"

	// ShiftTradeActivityRuleActivityCategoryTraining captures enum value "Training"
	ShiftTradeActivityRuleActivityCategoryTraining string = "Training"

	// ShiftTradeActivityRuleActivityCategoryUnavailable captures enum value "Unavailable"
	ShiftTradeActivityRuleActivityCategoryUnavailable string = "Unavailable"

	// ShiftTradeActivityRuleActivityCategoryUnscheduled captures enum value "Unscheduled"
	ShiftTradeActivityRuleActivityCategoryUnscheduled string = "Unscheduled"
)

// prop value enum
func (m *ShiftTradeActivityRule) validateActivityCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, shiftTradeActivityRuleTypeActivityCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ShiftTradeActivityRule) validateActivityCategory(formats strfmt.Registry) error {

	if err := validate.Required("activityCategory", "body", m.ActivityCategory); err != nil {
		return err
	}

	// value enum
	if err := m.validateActivityCategoryEnum("activityCategory", "body", *m.ActivityCategory); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShiftTradeActivityRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShiftTradeActivityRule) UnmarshalBinary(b []byte) error {
	var res ShiftTradeActivityRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}