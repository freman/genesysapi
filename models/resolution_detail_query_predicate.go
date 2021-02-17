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

// ResolutionDetailQueryPredicate resolution detail query predicate
//
// swagger:model ResolutionDetailQueryPredicate
type ResolutionDetailQueryPredicate struct {

	// Left hand side for metric predicates
	// Enum: [nNextContactAvoided]
	Metric string `json:"metric,omitempty"`

	// Optional operator, default is matches
	// Enum: [matches exists notExists]
	Operator string `json:"operator,omitempty"`

	// Right hand side for metric predicates
	Range *NumericRange `json:"range,omitempty"`

	// Optional type, can usually be inferred
	// Enum: [dimension property metric]
	Type string `json:"type,omitempty"`

	// Right hand side for metric predicates
	Value string `json:"value,omitempty"`
}

// Validate validates this resolution detail query predicate
func (m *ResolutionDetailQueryPredicate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetric(formats); err != nil {
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

var resolutionDetailQueryPredicateTypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nNextContactAvoided"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resolutionDetailQueryPredicateTypeMetricPropEnum = append(resolutionDetailQueryPredicateTypeMetricPropEnum, v)
	}
}

const (

	// ResolutionDetailQueryPredicateMetricNNextContactAvoided captures enum value "nNextContactAvoided"
	ResolutionDetailQueryPredicateMetricNNextContactAvoided string = "nNextContactAvoided"
)

// prop value enum
func (m *ResolutionDetailQueryPredicate) validateMetricEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resolutionDetailQueryPredicateTypeMetricPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResolutionDetailQueryPredicate) validateMetric(formats strfmt.Registry) error {

	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	// value enum
	if err := m.validateMetricEnum("metric", "body", m.Metric); err != nil {
		return err
	}

	return nil
}

var resolutionDetailQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resolutionDetailQueryPredicateTypeOperatorPropEnum = append(resolutionDetailQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// ResolutionDetailQueryPredicateOperatorMatches captures enum value "matches"
	ResolutionDetailQueryPredicateOperatorMatches string = "matches"

	// ResolutionDetailQueryPredicateOperatorExists captures enum value "exists"
	ResolutionDetailQueryPredicateOperatorExists string = "exists"

	// ResolutionDetailQueryPredicateOperatorNotExists captures enum value "notExists"
	ResolutionDetailQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *ResolutionDetailQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resolutionDetailQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResolutionDetailQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *ResolutionDetailQueryPredicate) validateRange(formats strfmt.Registry) error {

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

var resolutionDetailQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resolutionDetailQueryPredicateTypeTypePropEnum = append(resolutionDetailQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// ResolutionDetailQueryPredicateTypeDimension captures enum value "dimension"
	ResolutionDetailQueryPredicateTypeDimension string = "dimension"

	// ResolutionDetailQueryPredicateTypeProperty captures enum value "property"
	ResolutionDetailQueryPredicateTypeProperty string = "property"

	// ResolutionDetailQueryPredicateTypeMetric captures enum value "metric"
	ResolutionDetailQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *ResolutionDetailQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resolutionDetailQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResolutionDetailQueryPredicate) validateType(formats strfmt.Registry) error {

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
func (m *ResolutionDetailQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResolutionDetailQueryPredicate) UnmarshalBinary(b []byte) error {
	var res ResolutionDetailQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
