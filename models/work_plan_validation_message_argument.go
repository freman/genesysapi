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

// WorkPlanValidationMessageArgument work plan validation message argument
//
// swagger:model WorkPlanValidationMessageArgument
type WorkPlanValidationMessageArgument struct {

	// The type of the argument associated with violation messages
	// Enum: [Count MaxShiftCount Minutes ShiftId]
	Type string `json:"type,omitempty"`

	// The value of the argument
	Value string `json:"value,omitempty"`
}

// Validate validates this work plan validation message argument
func (m *WorkPlanValidationMessageArgument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var workPlanValidationMessageArgumentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Count","MaxShiftCount","Minutes","ShiftId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workPlanValidationMessageArgumentTypeTypePropEnum = append(workPlanValidationMessageArgumentTypeTypePropEnum, v)
	}
}

const (

	// WorkPlanValidationMessageArgumentTypeCount captures enum value "Count"
	WorkPlanValidationMessageArgumentTypeCount string = "Count"

	// WorkPlanValidationMessageArgumentTypeMaxShiftCount captures enum value "MaxShiftCount"
	WorkPlanValidationMessageArgumentTypeMaxShiftCount string = "MaxShiftCount"

	// WorkPlanValidationMessageArgumentTypeMinutes captures enum value "Minutes"
	WorkPlanValidationMessageArgumentTypeMinutes string = "Minutes"

	// WorkPlanValidationMessageArgumentTypeShiftID captures enum value "ShiftId"
	WorkPlanValidationMessageArgumentTypeShiftID string = "ShiftId"
)

// prop value enum
func (m *WorkPlanValidationMessageArgument) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workPlanValidationMessageArgumentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WorkPlanValidationMessageArgument) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkPlanValidationMessageArgument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkPlanValidationMessageArgument) UnmarshalBinary(b []byte) error {
	var res WorkPlanValidationMessageArgument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}