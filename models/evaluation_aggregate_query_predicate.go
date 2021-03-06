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

// EvaluationAggregateQueryPredicate evaluation aggregate query predicate
//
// swagger:model EvaluationAggregateQueryPredicate
type EvaluationAggregateQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [calibrationId contextId conversationId divisionId evaluationCreatedDate evaluationId evaluatorId formId queueId rescored teamId userId]
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

// Validate validates this evaluation aggregate query predicate
func (m *EvaluationAggregateQueryPredicate) Validate(formats strfmt.Registry) error {
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

var evaluationAggregateQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["calibrationId","contextId","conversationId","divisionId","evaluationCreatedDate","evaluationId","evaluatorId","formId","queueId","rescored","teamId","userId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationAggregateQueryPredicateTypeDimensionPropEnum = append(evaluationAggregateQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// EvaluationAggregateQueryPredicateDimensionCalibrationID captures enum value "calibrationId"
	EvaluationAggregateQueryPredicateDimensionCalibrationID string = "calibrationId"

	// EvaluationAggregateQueryPredicateDimensionContextID captures enum value "contextId"
	EvaluationAggregateQueryPredicateDimensionContextID string = "contextId"

	// EvaluationAggregateQueryPredicateDimensionConversationID captures enum value "conversationId"
	EvaluationAggregateQueryPredicateDimensionConversationID string = "conversationId"

	// EvaluationAggregateQueryPredicateDimensionDivisionID captures enum value "divisionId"
	EvaluationAggregateQueryPredicateDimensionDivisionID string = "divisionId"

	// EvaluationAggregateQueryPredicateDimensionEvaluationCreatedDate captures enum value "evaluationCreatedDate"
	EvaluationAggregateQueryPredicateDimensionEvaluationCreatedDate string = "evaluationCreatedDate"

	// EvaluationAggregateQueryPredicateDimensionEvaluationID captures enum value "evaluationId"
	EvaluationAggregateQueryPredicateDimensionEvaluationID string = "evaluationId"

	// EvaluationAggregateQueryPredicateDimensionEvaluatorID captures enum value "evaluatorId"
	EvaluationAggregateQueryPredicateDimensionEvaluatorID string = "evaluatorId"

	// EvaluationAggregateQueryPredicateDimensionFormID captures enum value "formId"
	EvaluationAggregateQueryPredicateDimensionFormID string = "formId"

	// EvaluationAggregateQueryPredicateDimensionQueueID captures enum value "queueId"
	EvaluationAggregateQueryPredicateDimensionQueueID string = "queueId"

	// EvaluationAggregateQueryPredicateDimensionRescored captures enum value "rescored"
	EvaluationAggregateQueryPredicateDimensionRescored string = "rescored"

	// EvaluationAggregateQueryPredicateDimensionTeamID captures enum value "teamId"
	EvaluationAggregateQueryPredicateDimensionTeamID string = "teamId"

	// EvaluationAggregateQueryPredicateDimensionUserID captures enum value "userId"
	EvaluationAggregateQueryPredicateDimensionUserID string = "userId"
)

// prop value enum
func (m *EvaluationAggregateQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationAggregateQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationAggregateQueryPredicate) validateDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var evaluationAggregateQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationAggregateQueryPredicateTypeOperatorPropEnum = append(evaluationAggregateQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// EvaluationAggregateQueryPredicateOperatorMatches captures enum value "matches"
	EvaluationAggregateQueryPredicateOperatorMatches string = "matches"

	// EvaluationAggregateQueryPredicateOperatorExists captures enum value "exists"
	EvaluationAggregateQueryPredicateOperatorExists string = "exists"

	// EvaluationAggregateQueryPredicateOperatorNotExists captures enum value "notExists"
	EvaluationAggregateQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *EvaluationAggregateQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationAggregateQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationAggregateQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *EvaluationAggregateQueryPredicate) validateRange(formats strfmt.Registry) error {

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

var evaluationAggregateQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationAggregateQueryPredicateTypeTypePropEnum = append(evaluationAggregateQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// EvaluationAggregateQueryPredicateTypeDimension captures enum value "dimension"
	EvaluationAggregateQueryPredicateTypeDimension string = "dimension"

	// EvaluationAggregateQueryPredicateTypeProperty captures enum value "property"
	EvaluationAggregateQueryPredicateTypeProperty string = "property"

	// EvaluationAggregateQueryPredicateTypeMetric captures enum value "metric"
	EvaluationAggregateQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *EvaluationAggregateQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationAggregateQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationAggregateQueryPredicate) validateType(formats strfmt.Registry) error {

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
func (m *EvaluationAggregateQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationAggregateQueryPredicate) UnmarshalBinary(b []byte) error {
	var res EvaluationAggregateQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
