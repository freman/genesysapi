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

// CreateTimeOffPlanRequest create time off plan request
//
// swagger:model CreateTimeOffPlanRequest
type CreateTimeOffPlanRequest struct {

	// Whether this time off plan should be used by agents.
	// Required: true
	Active *bool `json:"active"`

	// The set of activity code IDs to associate with this time off plan.
	// Unique: true
	ActivityCodeIds []string `json:"activityCodeIds"`

	// Auto approval rule for the time off plan.
	// Required: true
	// Enum: [Never Always CheckLimits]
	AutoApprovalRule *string `json:"autoApprovalRule"`

	// The number of days before the time off request start date for when the request will be expired from the waitlist.
	DaysBeforeStartToExpireFromWaitlist int32 `json:"daysBeforeStartToExpireFromWaitlist,omitempty"`

	// The name of this time off plan.
	// Required: true
	Name *string `json:"name"`

	// The set of time off limit IDs to associate with this time off plan.
	// Unique: true
	TimeOffLimitIds []string `json:"timeOffLimitIds"`
}

// Validate validates this create time off plan request
func (m *CreateTimeOffPlanRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivityCodeIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoApprovalRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOffLimitIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTimeOffPlanRequest) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *CreateTimeOffPlanRequest) validateActivityCodeIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityCodeIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("activityCodeIds", "body", m.ActivityCodeIds); err != nil {
		return err
	}

	return nil
}

var createTimeOffPlanRequestTypeAutoApprovalRulePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Never","Always","CheckLimits"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createTimeOffPlanRequestTypeAutoApprovalRulePropEnum = append(createTimeOffPlanRequestTypeAutoApprovalRulePropEnum, v)
	}
}

const (

	// CreateTimeOffPlanRequestAutoApprovalRuleNever captures enum value "Never"
	CreateTimeOffPlanRequestAutoApprovalRuleNever string = "Never"

	// CreateTimeOffPlanRequestAutoApprovalRuleAlways captures enum value "Always"
	CreateTimeOffPlanRequestAutoApprovalRuleAlways string = "Always"

	// CreateTimeOffPlanRequestAutoApprovalRuleCheckLimits captures enum value "CheckLimits"
	CreateTimeOffPlanRequestAutoApprovalRuleCheckLimits string = "CheckLimits"
)

// prop value enum
func (m *CreateTimeOffPlanRequest) validateAutoApprovalRuleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createTimeOffPlanRequestTypeAutoApprovalRulePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateTimeOffPlanRequest) validateAutoApprovalRule(formats strfmt.Registry) error {

	if err := validate.Required("autoApprovalRule", "body", m.AutoApprovalRule); err != nil {
		return err
	}

	// value enum
	if err := m.validateAutoApprovalRuleEnum("autoApprovalRule", "body", *m.AutoApprovalRule); err != nil {
		return err
	}

	return nil
}

func (m *CreateTimeOffPlanRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateTimeOffPlanRequest) validateTimeOffLimitIds(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeOffLimitIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("timeOffLimitIds", "body", m.TimeOffLimitIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTimeOffPlanRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTimeOffPlanRequest) UnmarshalBinary(b []byte) error {
	var res CreateTimeOffPlanRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}