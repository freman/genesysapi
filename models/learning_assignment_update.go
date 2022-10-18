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

// LearningAssignmentUpdate learning assignment update
//
// swagger:model LearningAssignmentUpdate
type LearningAssignmentUpdate struct {

	// An updated Assessment
	Assessment *LearningAssessment `json:"assessment,omitempty"`

	// The Learning Assignment state
	// Enum: [Assigned InProgress Completed Deleted NotCompleted InvalidSchedule]
	State string `json:"state,omitempty"`
}

// Validate validates this learning assignment update
func (m *LearningAssignmentUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssessment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningAssignmentUpdate) validateAssessment(formats strfmt.Registry) error {

	if swag.IsZero(m.Assessment) { // not required
		return nil
	}

	if m.Assessment != nil {
		if err := m.Assessment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessment")
			}
			return err
		}
	}

	return nil
}

var learningAssignmentUpdateTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Assigned","InProgress","Completed","Deleted","NotCompleted","InvalidSchedule"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningAssignmentUpdateTypeStatePropEnum = append(learningAssignmentUpdateTypeStatePropEnum, v)
	}
}

const (

	// LearningAssignmentUpdateStateAssigned captures enum value "Assigned"
	LearningAssignmentUpdateStateAssigned string = "Assigned"

	// LearningAssignmentUpdateStateInProgress captures enum value "InProgress"
	LearningAssignmentUpdateStateInProgress string = "InProgress"

	// LearningAssignmentUpdateStateCompleted captures enum value "Completed"
	LearningAssignmentUpdateStateCompleted string = "Completed"

	// LearningAssignmentUpdateStateDeleted captures enum value "Deleted"
	LearningAssignmentUpdateStateDeleted string = "Deleted"

	// LearningAssignmentUpdateStateNotCompleted captures enum value "NotCompleted"
	LearningAssignmentUpdateStateNotCompleted string = "NotCompleted"

	// LearningAssignmentUpdateStateInvalidSchedule captures enum value "InvalidSchedule"
	LearningAssignmentUpdateStateInvalidSchedule string = "InvalidSchedule"
)

// prop value enum
func (m *LearningAssignmentUpdate) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningAssignmentUpdateTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningAssignmentUpdate) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningAssignmentUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningAssignmentUpdate) UnmarshalBinary(b []byte) error {
	var res LearningAssignmentUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
