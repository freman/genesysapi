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

// EvaluationDetailQueryPredicate evaluation detail query predicate
//
// swagger:model EvaluationDetailQueryPredicate
type EvaluationDetailQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [assigneeId calibrationId contextId deleted evaluationId evaluatorId eventTime formId formName queueId released rescored userId]
	Dimension string `json:"dimension,omitempty"`

	// Left hand side for metric predicates
	// Enum: [oTotalCriticalScore oTotalScore]
	Metric string `json:"metric,omitempty"`

	// Optional operator, default is matches
	// Enum: [matches exists notExists]
	Operator string `json:"operator,omitempty"`

	// Right hand side for dimension or metric predicates
	Range *NumericRange `json:"range,omitempty"`

	// Optional type, can usually be inferred
	// Enum: [dimension property metric]
	Type string `json:"type,omitempty"`

	// Right hand side for dimension or metric predicates
	Value string `json:"value,omitempty"`
}

// Validate validates this evaluation detail query predicate
func (m *EvaluationDetailQueryPredicate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDimension(formats); err != nil {
		res = append(res, err)
	}

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

var evaluationDetailQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["assigneeId","calibrationId","contextId","deleted","evaluationId","evaluatorId","eventTime","formId","formName","queueId","released","rescored","userId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationDetailQueryPredicateTypeDimensionPropEnum = append(evaluationDetailQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// EvaluationDetailQueryPredicateDimensionAssigneeID captures enum value "assigneeId"
	EvaluationDetailQueryPredicateDimensionAssigneeID string = "assigneeId"

	// EvaluationDetailQueryPredicateDimensionCalibrationID captures enum value "calibrationId"
	EvaluationDetailQueryPredicateDimensionCalibrationID string = "calibrationId"

	// EvaluationDetailQueryPredicateDimensionContextID captures enum value "contextId"
	EvaluationDetailQueryPredicateDimensionContextID string = "contextId"

	// EvaluationDetailQueryPredicateDimensionDeleted captures enum value "deleted"
	EvaluationDetailQueryPredicateDimensionDeleted string = "deleted"

	// EvaluationDetailQueryPredicateDimensionEvaluationID captures enum value "evaluationId"
	EvaluationDetailQueryPredicateDimensionEvaluationID string = "evaluationId"

	// EvaluationDetailQueryPredicateDimensionEvaluatorID captures enum value "evaluatorId"
	EvaluationDetailQueryPredicateDimensionEvaluatorID string = "evaluatorId"

	// EvaluationDetailQueryPredicateDimensionEventTime captures enum value "eventTime"
	EvaluationDetailQueryPredicateDimensionEventTime string = "eventTime"

	// EvaluationDetailQueryPredicateDimensionFormID captures enum value "formId"
	EvaluationDetailQueryPredicateDimensionFormID string = "formId"

	// EvaluationDetailQueryPredicateDimensionFormName captures enum value "formName"
	EvaluationDetailQueryPredicateDimensionFormName string = "formName"

	// EvaluationDetailQueryPredicateDimensionQueueID captures enum value "queueId"
	EvaluationDetailQueryPredicateDimensionQueueID string = "queueId"

	// EvaluationDetailQueryPredicateDimensionReleased captures enum value "released"
	EvaluationDetailQueryPredicateDimensionReleased string = "released"

	// EvaluationDetailQueryPredicateDimensionRescored captures enum value "rescored"
	EvaluationDetailQueryPredicateDimensionRescored string = "rescored"

	// EvaluationDetailQueryPredicateDimensionUserID captures enum value "userId"
	EvaluationDetailQueryPredicateDimensionUserID string = "userId"
)

// prop value enum
func (m *EvaluationDetailQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationDetailQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationDetailQueryPredicate) validateDimension(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var evaluationDetailQueryPredicateTypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["oTotalCriticalScore","oTotalScore"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationDetailQueryPredicateTypeMetricPropEnum = append(evaluationDetailQueryPredicateTypeMetricPropEnum, v)
	}
}

const (

	// EvaluationDetailQueryPredicateMetricOTotalCriticalScore captures enum value "oTotalCriticalScore"
	EvaluationDetailQueryPredicateMetricOTotalCriticalScore string = "oTotalCriticalScore"

	// EvaluationDetailQueryPredicateMetricOTotalScore captures enum value "oTotalScore"
	EvaluationDetailQueryPredicateMetricOTotalScore string = "oTotalScore"
)

// prop value enum
func (m *EvaluationDetailQueryPredicate) validateMetricEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationDetailQueryPredicateTypeMetricPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationDetailQueryPredicate) validateMetric(formats strfmt.Registry) error {
	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	// value enum
	if err := m.validateMetricEnum("metric", "body", m.Metric); err != nil {
		return err
	}

	return nil
}

var evaluationDetailQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationDetailQueryPredicateTypeOperatorPropEnum = append(evaluationDetailQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// EvaluationDetailQueryPredicateOperatorMatches captures enum value "matches"
	EvaluationDetailQueryPredicateOperatorMatches string = "matches"

	// EvaluationDetailQueryPredicateOperatorExists captures enum value "exists"
	EvaluationDetailQueryPredicateOperatorExists string = "exists"

	// EvaluationDetailQueryPredicateOperatorNotExists captures enum value "notExists"
	EvaluationDetailQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *EvaluationDetailQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationDetailQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationDetailQueryPredicate) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *EvaluationDetailQueryPredicate) validateRange(formats strfmt.Registry) error {
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

var evaluationDetailQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationDetailQueryPredicateTypeTypePropEnum = append(evaluationDetailQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// EvaluationDetailQueryPredicateTypeDimension captures enum value "dimension"
	EvaluationDetailQueryPredicateTypeDimension string = "dimension"

	// EvaluationDetailQueryPredicateTypeProperty captures enum value "property"
	EvaluationDetailQueryPredicateTypeProperty string = "property"

	// EvaluationDetailQueryPredicateTypeMetric captures enum value "metric"
	EvaluationDetailQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *EvaluationDetailQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationDetailQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationDetailQueryPredicate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this evaluation detail query predicate based on the context it is used
func (m *EvaluationDetailQueryPredicate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvaluationDetailQueryPredicate) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EvaluationDetailQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationDetailQueryPredicate) UnmarshalBinary(b []byte) error {
	var res EvaluationDetailQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
