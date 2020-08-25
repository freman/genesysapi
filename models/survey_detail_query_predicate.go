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

// SurveyDetailQueryPredicate survey detail query predicate
//
// swagger:model SurveyDetailQueryPredicate
type SurveyDetailQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [eventTime queueId surveyCompletedDate surveyFormContextId surveyFormId surveyFormName surveyId surveyPromoterScore surveyStatus userId]
	Dimension string `json:"dimension,omitempty"`

	// Left hand side for metric predicates
	// Enum: [oSurveyTotalScore]
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

// Validate validates this survey detail query predicate
func (m *SurveyDetailQueryPredicate) Validate(formats strfmt.Registry) error {
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

var surveyDetailQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eventTime","queueId","surveyCompletedDate","surveyFormContextId","surveyFormId","surveyFormName","surveyId","surveyPromoterScore","surveyStatus","userId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		surveyDetailQueryPredicateTypeDimensionPropEnum = append(surveyDetailQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// SurveyDetailQueryPredicateDimensionEventTime captures enum value "eventTime"
	SurveyDetailQueryPredicateDimensionEventTime string = "eventTime"

	// SurveyDetailQueryPredicateDimensionQueueID captures enum value "queueId"
	SurveyDetailQueryPredicateDimensionQueueID string = "queueId"

	// SurveyDetailQueryPredicateDimensionSurveyCompletedDate captures enum value "surveyCompletedDate"
	SurveyDetailQueryPredicateDimensionSurveyCompletedDate string = "surveyCompletedDate"

	// SurveyDetailQueryPredicateDimensionSurveyFormContextID captures enum value "surveyFormContextId"
	SurveyDetailQueryPredicateDimensionSurveyFormContextID string = "surveyFormContextId"

	// SurveyDetailQueryPredicateDimensionSurveyFormID captures enum value "surveyFormId"
	SurveyDetailQueryPredicateDimensionSurveyFormID string = "surveyFormId"

	// SurveyDetailQueryPredicateDimensionSurveyFormName captures enum value "surveyFormName"
	SurveyDetailQueryPredicateDimensionSurveyFormName string = "surveyFormName"

	// SurveyDetailQueryPredicateDimensionSurveyID captures enum value "surveyId"
	SurveyDetailQueryPredicateDimensionSurveyID string = "surveyId"

	// SurveyDetailQueryPredicateDimensionSurveyPromoterScore captures enum value "surveyPromoterScore"
	SurveyDetailQueryPredicateDimensionSurveyPromoterScore string = "surveyPromoterScore"

	// SurveyDetailQueryPredicateDimensionSurveyStatus captures enum value "surveyStatus"
	SurveyDetailQueryPredicateDimensionSurveyStatus string = "surveyStatus"

	// SurveyDetailQueryPredicateDimensionUserID captures enum value "userId"
	SurveyDetailQueryPredicateDimensionUserID string = "userId"
)

// prop value enum
func (m *SurveyDetailQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, surveyDetailQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SurveyDetailQueryPredicate) validateDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var surveyDetailQueryPredicateTypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["oSurveyTotalScore"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		surveyDetailQueryPredicateTypeMetricPropEnum = append(surveyDetailQueryPredicateTypeMetricPropEnum, v)
	}
}

const (

	// SurveyDetailQueryPredicateMetricOSurveyTotalScore captures enum value "oSurveyTotalScore"
	SurveyDetailQueryPredicateMetricOSurveyTotalScore string = "oSurveyTotalScore"
)

// prop value enum
func (m *SurveyDetailQueryPredicate) validateMetricEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, surveyDetailQueryPredicateTypeMetricPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SurveyDetailQueryPredicate) validateMetric(formats strfmt.Registry) error {

	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	// value enum
	if err := m.validateMetricEnum("metric", "body", m.Metric); err != nil {
		return err
	}

	return nil
}

var surveyDetailQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		surveyDetailQueryPredicateTypeOperatorPropEnum = append(surveyDetailQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// SurveyDetailQueryPredicateOperatorMatches captures enum value "matches"
	SurveyDetailQueryPredicateOperatorMatches string = "matches"

	// SurveyDetailQueryPredicateOperatorExists captures enum value "exists"
	SurveyDetailQueryPredicateOperatorExists string = "exists"

	// SurveyDetailQueryPredicateOperatorNotExists captures enum value "notExists"
	SurveyDetailQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *SurveyDetailQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, surveyDetailQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SurveyDetailQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *SurveyDetailQueryPredicate) validateRange(formats strfmt.Registry) error {

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

var surveyDetailQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		surveyDetailQueryPredicateTypeTypePropEnum = append(surveyDetailQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// SurveyDetailQueryPredicateTypeDimension captures enum value "dimension"
	SurveyDetailQueryPredicateTypeDimension string = "dimension"

	// SurveyDetailQueryPredicateTypeProperty captures enum value "property"
	SurveyDetailQueryPredicateTypeProperty string = "property"

	// SurveyDetailQueryPredicateTypeMetric captures enum value "metric"
	SurveyDetailQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *SurveyDetailQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, surveyDetailQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SurveyDetailQueryPredicate) validateType(formats strfmt.Registry) error {

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
func (m *SurveyDetailQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SurveyDetailQueryPredicate) UnmarshalBinary(b []byte) error {
	var res SurveyDetailQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}