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

// EntityTypeCriteria entity type criteria
//
// swagger:model EntityTypeCriteria
type EntityTypeCriteria struct {

	// The entity to match the pattern against.
	// Enum: [visit]
	EntityType string `json:"entityType,omitempty"`

	// The criteria key.
	// Required: true
	Key *string `json:"key"`

	// The comparison operator.
	// Enum: [containsAll containsAny notContainsAll notContainsAny equal notEqual greaterThan greaterThanOrEqual lessThan lessThanOrEqual startsWith endsWith]
	Operator string `json:"operator,omitempty"`

	// Should criteria be case insensitive.
	// Required: true
	ShouldIgnoreCase *bool `json:"shouldIgnoreCase"`

	// The criteria values.
	// Required: true
	Values []string `json:"values"`
}

// Validate validates this entity type criteria
func (m *EntityTypeCriteria) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShouldIgnoreCase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var entityTypeCriteriaTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["visit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		entityTypeCriteriaTypeEntityTypePropEnum = append(entityTypeCriteriaTypeEntityTypePropEnum, v)
	}
}

const (

	// EntityTypeCriteriaEntityTypeVisit captures enum value "visit"
	EntityTypeCriteriaEntityTypeVisit string = "visit"
)

// prop value enum
func (m *EntityTypeCriteria) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, entityTypeCriteriaTypeEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EntityTypeCriteria) validateEntityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *EntityTypeCriteria) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

var entityTypeCriteriaTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["containsAll","containsAny","notContainsAll","notContainsAny","equal","notEqual","greaterThan","greaterThanOrEqual","lessThan","lessThanOrEqual","startsWith","endsWith"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		entityTypeCriteriaTypeOperatorPropEnum = append(entityTypeCriteriaTypeOperatorPropEnum, v)
	}
}

const (

	// EntityTypeCriteriaOperatorContainsAll captures enum value "containsAll"
	EntityTypeCriteriaOperatorContainsAll string = "containsAll"

	// EntityTypeCriteriaOperatorContainsAny captures enum value "containsAny"
	EntityTypeCriteriaOperatorContainsAny string = "containsAny"

	// EntityTypeCriteriaOperatorNotContainsAll captures enum value "notContainsAll"
	EntityTypeCriteriaOperatorNotContainsAll string = "notContainsAll"

	// EntityTypeCriteriaOperatorNotContainsAny captures enum value "notContainsAny"
	EntityTypeCriteriaOperatorNotContainsAny string = "notContainsAny"

	// EntityTypeCriteriaOperatorEqual captures enum value "equal"
	EntityTypeCriteriaOperatorEqual string = "equal"

	// EntityTypeCriteriaOperatorNotEqual captures enum value "notEqual"
	EntityTypeCriteriaOperatorNotEqual string = "notEqual"

	// EntityTypeCriteriaOperatorGreaterThan captures enum value "greaterThan"
	EntityTypeCriteriaOperatorGreaterThan string = "greaterThan"

	// EntityTypeCriteriaOperatorGreaterThanOrEqual captures enum value "greaterThanOrEqual"
	EntityTypeCriteriaOperatorGreaterThanOrEqual string = "greaterThanOrEqual"

	// EntityTypeCriteriaOperatorLessThan captures enum value "lessThan"
	EntityTypeCriteriaOperatorLessThan string = "lessThan"

	// EntityTypeCriteriaOperatorLessThanOrEqual captures enum value "lessThanOrEqual"
	EntityTypeCriteriaOperatorLessThanOrEqual string = "lessThanOrEqual"

	// EntityTypeCriteriaOperatorStartsWith captures enum value "startsWith"
	EntityTypeCriteriaOperatorStartsWith string = "startsWith"

	// EntityTypeCriteriaOperatorEndsWith captures enum value "endsWith"
	EntityTypeCriteriaOperatorEndsWith string = "endsWith"
)

// prop value enum
func (m *EntityTypeCriteria) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, entityTypeCriteriaTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EntityTypeCriteria) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *EntityTypeCriteria) validateShouldIgnoreCase(formats strfmt.Registry) error {

	if err := validate.Required("shouldIgnoreCase", "body", m.ShouldIgnoreCase); err != nil {
		return err
	}

	return nil
}

func (m *EntityTypeCriteria) validateValues(formats strfmt.Registry) error {

	if err := validate.Required("values", "body", m.Values); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityTypeCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityTypeCriteria) UnmarshalBinary(b []byte) error {
	var res EntityTypeCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
