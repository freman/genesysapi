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

// QueueObservationQueryPredicate queue observation query predicate
//
// swagger:model QueueObservationQueryPredicate
type QueueObservationQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [mediaType queueId]
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

// Validate validates this queue observation query predicate
func (m *QueueObservationQueryPredicate) Validate(formats strfmt.Registry) error {
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

var queueObservationQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mediaType","queueId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queueObservationQueryPredicateTypeDimensionPropEnum = append(queueObservationQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// QueueObservationQueryPredicateDimensionMediaType captures enum value "mediaType"
	QueueObservationQueryPredicateDimensionMediaType string = "mediaType"

	// QueueObservationQueryPredicateDimensionQueueID captures enum value "queueId"
	QueueObservationQueryPredicateDimensionQueueID string = "queueId"
)

// prop value enum
func (m *QueueObservationQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, queueObservationQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QueueObservationQueryPredicate) validateDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var queueObservationQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queueObservationQueryPredicateTypeOperatorPropEnum = append(queueObservationQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// QueueObservationQueryPredicateOperatorMatches captures enum value "matches"
	QueueObservationQueryPredicateOperatorMatches string = "matches"

	// QueueObservationQueryPredicateOperatorExists captures enum value "exists"
	QueueObservationQueryPredicateOperatorExists string = "exists"

	// QueueObservationQueryPredicateOperatorNotExists captures enum value "notExists"
	QueueObservationQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *QueueObservationQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, queueObservationQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QueueObservationQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *QueueObservationQueryPredicate) validateRange(formats strfmt.Registry) error {

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

var queueObservationQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queueObservationQueryPredicateTypeTypePropEnum = append(queueObservationQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// QueueObservationQueryPredicateTypeDimension captures enum value "dimension"
	QueueObservationQueryPredicateTypeDimension string = "dimension"

	// QueueObservationQueryPredicateTypeProperty captures enum value "property"
	QueueObservationQueryPredicateTypeProperty string = "property"

	// QueueObservationQueryPredicateTypeMetric captures enum value "metric"
	QueueObservationQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *QueueObservationQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, queueObservationQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QueueObservationQueryPredicate) validateType(formats strfmt.Registry) error {

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
func (m *QueueObservationQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueueObservationQueryPredicate) UnmarshalBinary(b []byte) error {
	var res QueueObservationQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
