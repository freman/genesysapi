// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceConditionNode resource condition node
//
// swagger:model ResourceConditionNode
type ResourceConditionNode struct {

	// conjunction
	// Enum: [AND OR]
	Conjunction string `json:"conjunction,omitempty"`

	// operands
	Operands []*ResourceConditionValue `json:"operands"`

	// operator
	// Enum: [EQ IN GE GT LE LT]
	Operator string `json:"operator,omitempty"`

	// terms
	Terms []*ResourceConditionNode `json:"terms"`

	// variable name
	VariableName string `json:"variableName,omitempty"`
}

// Validate validates this resource condition node
func (m *ResourceConditionNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConjunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperands(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var resourceConditionNodeTypeConjunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceConditionNodeTypeConjunctionPropEnum = append(resourceConditionNodeTypeConjunctionPropEnum, v)
	}
}

const (

	// ResourceConditionNodeConjunctionAND captures enum value "AND"
	ResourceConditionNodeConjunctionAND string = "AND"

	// ResourceConditionNodeConjunctionOR captures enum value "OR"
	ResourceConditionNodeConjunctionOR string = "OR"
)

// prop value enum
func (m *ResourceConditionNode) validateConjunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resourceConditionNodeTypeConjunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResourceConditionNode) validateConjunction(formats strfmt.Registry) error {

	if swag.IsZero(m.Conjunction) { // not required
		return nil
	}

	// value enum
	if err := m.validateConjunctionEnum("conjunction", "body", m.Conjunction); err != nil {
		return err
	}

	return nil
}

func (m *ResourceConditionNode) validateOperands(formats strfmt.Registry) error {

	if swag.IsZero(m.Operands) { // not required
		return nil
	}

	for i := 0; i < len(m.Operands); i++ {
		if swag.IsZero(m.Operands[i]) { // not required
			continue
		}

		if m.Operands[i] != nil {
			if err := m.Operands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var resourceConditionNodeTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQ","IN","GE","GT","LE","LT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceConditionNodeTypeOperatorPropEnum = append(resourceConditionNodeTypeOperatorPropEnum, v)
	}
}

const (

	// ResourceConditionNodeOperatorEQ captures enum value "EQ"
	ResourceConditionNodeOperatorEQ string = "EQ"

	// ResourceConditionNodeOperatorIN captures enum value "IN"
	ResourceConditionNodeOperatorIN string = "IN"

	// ResourceConditionNodeOperatorGE captures enum value "GE"
	ResourceConditionNodeOperatorGE string = "GE"

	// ResourceConditionNodeOperatorGT captures enum value "GT"
	ResourceConditionNodeOperatorGT string = "GT"

	// ResourceConditionNodeOperatorLE captures enum value "LE"
	ResourceConditionNodeOperatorLE string = "LE"

	// ResourceConditionNodeOperatorLT captures enum value "LT"
	ResourceConditionNodeOperatorLT string = "LT"
)

// prop value enum
func (m *ResourceConditionNode) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resourceConditionNodeTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResourceConditionNode) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *ResourceConditionNode) validateTerms(formats strfmt.Registry) error {

	if swag.IsZero(m.Terms) { // not required
		return nil
	}

	for i := 0; i < len(m.Terms); i++ {
		if swag.IsZero(m.Terms[i]) { // not required
			continue
		}

		if m.Terms[i] != nil {
			if err := m.Terms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("terms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceConditionNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceConditionNode) UnmarshalBinary(b []byte) error {
	var res ResourceConditionNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
