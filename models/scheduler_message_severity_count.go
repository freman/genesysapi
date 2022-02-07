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

// SchedulerMessageSeverityCount scheduler message severity count
//
// swagger:model SchedulerMessageSeverityCount
type SchedulerMessageSeverityCount struct {

	// The number of schedule messages with the given severity
	Count int32 `json:"count,omitempty"`

	// The schedule message severity
	// Enum: [Ignore Information Warning Error]
	Severity string `json:"severity,omitempty"`
}

// Validate validates this scheduler message severity count
func (m *SchedulerMessageSeverityCount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var schedulerMessageSeverityCountTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ignore","Information","Warning","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schedulerMessageSeverityCountTypeSeverityPropEnum = append(schedulerMessageSeverityCountTypeSeverityPropEnum, v)
	}
}

const (

	// SchedulerMessageSeverityCountSeverityIgnore captures enum value "Ignore"
	SchedulerMessageSeverityCountSeverityIgnore string = "Ignore"

	// SchedulerMessageSeverityCountSeverityInformation captures enum value "Information"
	SchedulerMessageSeverityCountSeverityInformation string = "Information"

	// SchedulerMessageSeverityCountSeverityWarning captures enum value "Warning"
	SchedulerMessageSeverityCountSeverityWarning string = "Warning"

	// SchedulerMessageSeverityCountSeverityError captures enum value "Error"
	SchedulerMessageSeverityCountSeverityError string = "Error"
)

// prop value enum
func (m *SchedulerMessageSeverityCount) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, schedulerMessageSeverityCountTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SchedulerMessageSeverityCount) validateSeverity(formats strfmt.Registry) error {

	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerMessageSeverityCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerMessageSeverityCount) UnmarshalBinary(b []byte) error {
	var res SchedulerMessageSeverityCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}