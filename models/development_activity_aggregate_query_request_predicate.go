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

// DevelopmentActivityAggregateQueryRequestPredicate development activity aggregate query request predicate
//
// swagger:model DevelopmentActivityAggregateQueryRequestPredicate
type DevelopmentActivityAggregateQueryRequestPredicate struct {

	// Each predicates specifies a dimension.
	// Required: true
	// Enum: [attendeeId type moduleId isPassed]
	Dimension *string `json:"dimension"`

	// Corresponding value for dimensions in predicates. If the dimension is type, Valid Values: Informational, AssessedContent, Assessment, Coaching
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this development activity aggregate query request predicate
func (m *DevelopmentActivityAggregateQueryRequestPredicate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDimension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var developmentActivityAggregateQueryRequestPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["attendeeId","type","moduleId","isPassed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		developmentActivityAggregateQueryRequestPredicateTypeDimensionPropEnum = append(developmentActivityAggregateQueryRequestPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// DevelopmentActivityAggregateQueryRequestPredicateDimensionAttendeeID captures enum value "attendeeId"
	DevelopmentActivityAggregateQueryRequestPredicateDimensionAttendeeID string = "attendeeId"

	// DevelopmentActivityAggregateQueryRequestPredicateDimensionType captures enum value "type"
	DevelopmentActivityAggregateQueryRequestPredicateDimensionType string = "type"

	// DevelopmentActivityAggregateQueryRequestPredicateDimensionModuleID captures enum value "moduleId"
	DevelopmentActivityAggregateQueryRequestPredicateDimensionModuleID string = "moduleId"

	// DevelopmentActivityAggregateQueryRequestPredicateDimensionIsPassed captures enum value "isPassed"
	DevelopmentActivityAggregateQueryRequestPredicateDimensionIsPassed string = "isPassed"
)

// prop value enum
func (m *DevelopmentActivityAggregateQueryRequestPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, developmentActivityAggregateQueryRequestPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DevelopmentActivityAggregateQueryRequestPredicate) validateDimension(formats strfmt.Registry) error {

	if err := validate.Required("dimension", "body", m.Dimension); err != nil {
		return err
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", *m.Dimension); err != nil {
		return err
	}

	return nil
}

func (m *DevelopmentActivityAggregateQueryRequestPredicate) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DevelopmentActivityAggregateQueryRequestPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevelopmentActivityAggregateQueryRequestPredicate) UnmarshalBinary(b []byte) error {
	var res DevelopmentActivityAggregateQueryRequestPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
