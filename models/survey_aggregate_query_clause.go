// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SurveyAggregateQueryClause survey aggregate query clause
//
// swagger:model SurveyAggregateQueryClause
type SurveyAggregateQueryClause struct {

	// Like a three-word sentence: (attribute-name) (operator) (target-value).
	// Required: true
	Predicates []*SurveyAggregateQueryPredicate `json:"predicates"`

	// Boolean operation to apply to the provided predicates
	// Required: true
	// Enum: [and or]
	Type *string `json:"type"`
}

// Validate validates this survey aggregate query clause
func (m *SurveyAggregateQueryClause) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePredicates(formats); err != nil {
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

func (m *SurveyAggregateQueryClause) validatePredicates(formats strfmt.Registry) error {

	if err := validate.Required("predicates", "body", m.Predicates); err != nil {
		return err
	}

	for i := 0; i < len(m.Predicates); i++ {
		if swag.IsZero(m.Predicates[i]) { // not required
			continue
		}

		if m.Predicates[i] != nil {
			if err := m.Predicates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("predicates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("predicates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var surveyAggregateQueryClauseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["and","or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		surveyAggregateQueryClauseTypeTypePropEnum = append(surveyAggregateQueryClauseTypeTypePropEnum, v)
	}
}

const (

	// SurveyAggregateQueryClauseTypeAnd captures enum value "and"
	SurveyAggregateQueryClauseTypeAnd string = "and"

	// SurveyAggregateQueryClauseTypeOr captures enum value "or"
	SurveyAggregateQueryClauseTypeOr string = "or"
)

// prop value enum
func (m *SurveyAggregateQueryClause) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, surveyAggregateQueryClauseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SurveyAggregateQueryClause) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this survey aggregate query clause based on the context it is used
func (m *SurveyAggregateQueryClause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePredicates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SurveyAggregateQueryClause) contextValidatePredicates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Predicates); i++ {

		if m.Predicates[i] != nil {
			if err := m.Predicates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("predicates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("predicates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SurveyAggregateQueryClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SurveyAggregateQueryClause) UnmarshalBinary(b []byte) error {
	var res SurveyAggregateQueryClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
