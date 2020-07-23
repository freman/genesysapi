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

// UserDetailQueryPredicate user detail query predicate
//
// swagger:model UserDetailQueryPredicate
type UserDetailQueryPredicate struct {

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

// Validate validates this user detail query predicate
func (m *UserDetailQueryPredicate) Validate(formats strfmt.Registry) error {
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

var userDetailQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["userId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userDetailQueryPredicateTypeDimensionPropEnum = append(userDetailQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// UserDetailQueryPredicateDimensionUserID captures enum value "userId"
	UserDetailQueryPredicateDimensionUserID string = "userId"
)

// prop value enum
func (m *UserDetailQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userDetailQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserDetailQueryPredicate) validateDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var userDetailQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userDetailQueryPredicateTypeOperatorPropEnum = append(userDetailQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// UserDetailQueryPredicateOperatorMatches captures enum value "matches"
	UserDetailQueryPredicateOperatorMatches string = "matches"

	// UserDetailQueryPredicateOperatorExists captures enum value "exists"
	UserDetailQueryPredicateOperatorExists string = "exists"

	// UserDetailQueryPredicateOperatorNotExists captures enum value "notExists"
	UserDetailQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *UserDetailQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userDetailQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserDetailQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *UserDetailQueryPredicate) validateRange(formats strfmt.Registry) error {

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

var userDetailQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userDetailQueryPredicateTypeTypePropEnum = append(userDetailQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// UserDetailQueryPredicateTypeDimension captures enum value "dimension"
	UserDetailQueryPredicateTypeDimension string = "dimension"

	// UserDetailQueryPredicateTypeProperty captures enum value "property"
	UserDetailQueryPredicateTypeProperty string = "property"

	// UserDetailQueryPredicateTypeMetric captures enum value "metric"
	UserDetailQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *UserDetailQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userDetailQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserDetailQueryPredicate) validateType(formats strfmt.Registry) error {

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
func (m *UserDetailQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDetailQueryPredicate) UnmarshalBinary(b []byte) error {
	var res UserDetailQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
