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

// VisibilityCondition visibility condition
//
// swagger:model VisibilityCondition
type VisibilityCondition struct {

	// combining operation
	// Enum: [AND OR]
	CombiningOperation string `json:"combiningOperation,omitempty"`

	// A list of strings, each representing the location in the form of the Answer Option to depend on. In the format of "/form/questionGroup/{questionGroupIndex}/question/{questionIndex}/answer/{answerIndex}" or, to assume the current question group, "../question/{questionIndex}/answer/{answerIndex}". Note: Indexes are zero-based
	Predicates []interface{} `json:"predicates"`
}

// Validate validates this visibility condition
func (m *VisibilityCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCombiningOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var visibilityConditionTypeCombiningOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		visibilityConditionTypeCombiningOperationPropEnum = append(visibilityConditionTypeCombiningOperationPropEnum, v)
	}
}

const (

	// VisibilityConditionCombiningOperationAND captures enum value "AND"
	VisibilityConditionCombiningOperationAND string = "AND"

	// VisibilityConditionCombiningOperationOR captures enum value "OR"
	VisibilityConditionCombiningOperationOR string = "OR"
)

// prop value enum
func (m *VisibilityCondition) validateCombiningOperationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, visibilityConditionTypeCombiningOperationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VisibilityCondition) validateCombiningOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.CombiningOperation) { // not required
		return nil
	}

	// value enum
	if err := m.validateCombiningOperationEnum("combiningOperation", "body", m.CombiningOperation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VisibilityCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VisibilityCondition) UnmarshalBinary(b []byte) error {
	var res VisibilityCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
