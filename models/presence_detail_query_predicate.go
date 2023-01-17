// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PresenceDetailQueryPredicate presence detail query predicate
//
// swagger:model PresenceDetailQueryPredicate
type PresenceDetailQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [organizationPresenceId systemPresence]
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

// Validate validates this presence detail query predicate
func (m *PresenceDetailQueryPredicate) Validate(formats strfmt.Registry) error {
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

var presenceDetailQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["organizationPresenceId","systemPresence"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		presenceDetailQueryPredicateTypeDimensionPropEnum = append(presenceDetailQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// PresenceDetailQueryPredicateDimensionOrganizationPresenceID captures enum value "organizationPresenceId"
	PresenceDetailQueryPredicateDimensionOrganizationPresenceID string = "organizationPresenceId"

	// PresenceDetailQueryPredicateDimensionSystemPresence captures enum value "systemPresence"
	PresenceDetailQueryPredicateDimensionSystemPresence string = "systemPresence"
)

// prop value enum
func (m *PresenceDetailQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, presenceDetailQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PresenceDetailQueryPredicate) validateDimension(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var presenceDetailQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		presenceDetailQueryPredicateTypeOperatorPropEnum = append(presenceDetailQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// PresenceDetailQueryPredicateOperatorMatches captures enum value "matches"
	PresenceDetailQueryPredicateOperatorMatches string = "matches"

	// PresenceDetailQueryPredicateOperatorExists captures enum value "exists"
	PresenceDetailQueryPredicateOperatorExists string = "exists"

	// PresenceDetailQueryPredicateOperatorNotExists captures enum value "notExists"
	PresenceDetailQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *PresenceDetailQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, presenceDetailQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PresenceDetailQueryPredicate) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *PresenceDetailQueryPredicate) validateRange(formats strfmt.Registry) error {
	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

var presenceDetailQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		presenceDetailQueryPredicateTypeTypePropEnum = append(presenceDetailQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// PresenceDetailQueryPredicateTypeDimension captures enum value "dimension"
	PresenceDetailQueryPredicateTypeDimension string = "dimension"

	// PresenceDetailQueryPredicateTypeProperty captures enum value "property"
	PresenceDetailQueryPredicateTypeProperty string = "property"

	// PresenceDetailQueryPredicateTypeMetric captures enum value "metric"
	PresenceDetailQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *PresenceDetailQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, presenceDetailQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PresenceDetailQueryPredicate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this presence detail query predicate based on the context it is used
func (m *PresenceDetailQueryPredicate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PresenceDetailQueryPredicate) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

	if m.Range != nil {
		if err := m.Range.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PresenceDetailQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PresenceDetailQueryPredicate) UnmarshalBinary(b []byte) error {
	var res PresenceDetailQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
