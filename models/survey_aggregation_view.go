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

// SurveyAggregationView survey aggregation view
//
// swagger:model SurveyAggregationView
type SurveyAggregationView struct {

	// Type of view you wish to create
	// Required: true
	// Enum: [rangeBound]
	Function *string `json:"function"`

	// A unique name for this view. Must be distinct from other views and built-in metric names.
	// Required: true
	Name *string `json:"name"`

	// Range of numbers for slicing up data
	Range *AggregationRange `json:"range,omitempty"`

	// Target metric name
	// Required: true
	// Enum: [nSurveyErrors nSurveyNpsDetractors nSurveyNpsPromoters nSurveyNpsResponses nSurveyQuestionGroupResponses nSurveyQuestionResponses nSurveyResponses nSurveysAbandoned nSurveysDeleted nSurveysExpired nSurveysSent nSurveysStarted oSurveyQuestionGroupScore oSurveyQuestionScore oSurveyTotalScore]
	Target *string `json:"target"`
}

// Validate validates this survey aggregation view
func (m *SurveyAggregationView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var surveyAggregationViewTypeFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rangeBound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		surveyAggregationViewTypeFunctionPropEnum = append(surveyAggregationViewTypeFunctionPropEnum, v)
	}
}

const (

	// SurveyAggregationViewFunctionRangeBound captures enum value "rangeBound"
	SurveyAggregationViewFunctionRangeBound string = "rangeBound"
)

// prop value enum
func (m *SurveyAggregationView) validateFunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, surveyAggregationViewTypeFunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SurveyAggregationView) validateFunction(formats strfmt.Registry) error {

	if err := validate.Required("function", "body", m.Function); err != nil {
		return err
	}

	// value enum
	if err := m.validateFunctionEnum("function", "body", *m.Function); err != nil {
		return err
	}

	return nil
}

func (m *SurveyAggregationView) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SurveyAggregationView) validateRange(formats strfmt.Registry) error {

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

var surveyAggregationViewTypeTargetPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nSurveyErrors","nSurveyNpsDetractors","nSurveyNpsPromoters","nSurveyNpsResponses","nSurveyQuestionGroupResponses","nSurveyQuestionResponses","nSurveyResponses","nSurveysAbandoned","nSurveysDeleted","nSurveysExpired","nSurveysSent","nSurveysStarted","oSurveyQuestionGroupScore","oSurveyQuestionScore","oSurveyTotalScore"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		surveyAggregationViewTypeTargetPropEnum = append(surveyAggregationViewTypeTargetPropEnum, v)
	}
}

const (

	// SurveyAggregationViewTargetNSurveyErrors captures enum value "nSurveyErrors"
	SurveyAggregationViewTargetNSurveyErrors string = "nSurveyErrors"

	// SurveyAggregationViewTargetNSurveyNpsDetractors captures enum value "nSurveyNpsDetractors"
	SurveyAggregationViewTargetNSurveyNpsDetractors string = "nSurveyNpsDetractors"

	// SurveyAggregationViewTargetNSurveyNpsPromoters captures enum value "nSurveyNpsPromoters"
	SurveyAggregationViewTargetNSurveyNpsPromoters string = "nSurveyNpsPromoters"

	// SurveyAggregationViewTargetNSurveyNpsResponses captures enum value "nSurveyNpsResponses"
	SurveyAggregationViewTargetNSurveyNpsResponses string = "nSurveyNpsResponses"

	// SurveyAggregationViewTargetNSurveyQuestionGroupResponses captures enum value "nSurveyQuestionGroupResponses"
	SurveyAggregationViewTargetNSurveyQuestionGroupResponses string = "nSurveyQuestionGroupResponses"

	// SurveyAggregationViewTargetNSurveyQuestionResponses captures enum value "nSurveyQuestionResponses"
	SurveyAggregationViewTargetNSurveyQuestionResponses string = "nSurveyQuestionResponses"

	// SurveyAggregationViewTargetNSurveyResponses captures enum value "nSurveyResponses"
	SurveyAggregationViewTargetNSurveyResponses string = "nSurveyResponses"

	// SurveyAggregationViewTargetNSurveysAbandoned captures enum value "nSurveysAbandoned"
	SurveyAggregationViewTargetNSurveysAbandoned string = "nSurveysAbandoned"

	// SurveyAggregationViewTargetNSurveysDeleted captures enum value "nSurveysDeleted"
	SurveyAggregationViewTargetNSurveysDeleted string = "nSurveysDeleted"

	// SurveyAggregationViewTargetNSurveysExpired captures enum value "nSurveysExpired"
	SurveyAggregationViewTargetNSurveysExpired string = "nSurveysExpired"

	// SurveyAggregationViewTargetNSurveysSent captures enum value "nSurveysSent"
	SurveyAggregationViewTargetNSurveysSent string = "nSurveysSent"

	// SurveyAggregationViewTargetNSurveysStarted captures enum value "nSurveysStarted"
	SurveyAggregationViewTargetNSurveysStarted string = "nSurveysStarted"

	// SurveyAggregationViewTargetOSurveyQuestionGroupScore captures enum value "oSurveyQuestionGroupScore"
	SurveyAggregationViewTargetOSurveyQuestionGroupScore string = "oSurveyQuestionGroupScore"

	// SurveyAggregationViewTargetOSurveyQuestionScore captures enum value "oSurveyQuestionScore"
	SurveyAggregationViewTargetOSurveyQuestionScore string = "oSurveyQuestionScore"

	// SurveyAggregationViewTargetOSurveyTotalScore captures enum value "oSurveyTotalScore"
	SurveyAggregationViewTargetOSurveyTotalScore string = "oSurveyTotalScore"
)

// prop value enum
func (m *SurveyAggregationView) validateTargetEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, surveyAggregationViewTypeTargetPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SurveyAggregationView) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnum("target", "body", *m.Target); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SurveyAggregationView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SurveyAggregationView) UnmarshalBinary(b []byte) error {
	var res SurveyAggregationView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
