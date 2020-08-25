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

// UserObservationQueryPredicate user observation query predicate
//
// swagger:model UserObservationQueryPredicate
type UserObservationQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [userId]
	Dimension string `json:"dimension,omitempty"`

	// Optional operator, default is matches
	// Enum: [matches exists notExists]
	Operator string `json:"operator,omitempty"`

	// Right hand side for dimension predicates
	Range *NumericRange `json:"range,omitempty"`

	// Optional type, can usually be inferred
	// Enum: [dimension property metric]
	Type string `json:"type,omitempty"`

	// Right hand side for dimension predicates
	Value string `json:"value,omitempty"`
}

// Validate validates this user observation query predicate
func (m *UserObservationQueryPredicate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDimension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userObservationQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["userId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userObservationQueryPredicateTypeDimensionPropEnum = append(userObservationQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// UserObservationQueryPredicateDimensionUserID captures enum value "userId"
	UserObservationQueryPredicateDimensionUserID string = "userId"
)

// prop value enum
func (m *UserObservationQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userObservationQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserObservationQueryPredicate) validateDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var userObservationQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userObservationQueryPredicateTypeOperatorPropEnum = append(userObservationQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// UserObservationQueryPredicateOperatorMatches captures enum value "matches"
	UserObservationQueryPredicateOperatorMatches string = "matches"

	// UserObservationQueryPredicateOperatorExists captures enum value "exists"
	UserObservationQueryPredicateOperatorExists string = "exists"

	// UserObservationQueryPredicateOperatorNotExists captures enum value "notExists"
	UserObservationQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *UserObservationQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userObservationQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserObservationQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *UserObservationQueryPredicate) validateRange(formats strfmt.Registry) error {

	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

var userObservationQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userObservationQueryPredicateTypeTypePropEnum = append(userObservationQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// UserObservationQueryPredicateTypeDimension captures enum value "dimension"
	UserObservationQueryPredicateTypeDimension string = "dimension"

	// UserObservationQueryPredicateTypeProperty captures enum value "property"
	UserObservationQueryPredicateTypeProperty string = "property"

	// UserObservationQueryPredicateTypeMetric captures enum value "metric"
	UserObservationQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *UserObservationQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userObservationQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserObservationQueryPredicate) validateType(formats strfmt.Registry) error {

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
func (m *UserObservationQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserObservationQueryPredicate) UnmarshalBinary(b []byte) error {
	var res UserObservationQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}