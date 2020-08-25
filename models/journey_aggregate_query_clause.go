// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JourneyAggregateQueryClause journey aggregate query clause
//
// swagger:model JourneyAggregateQueryClause
type JourneyAggregateQueryClause struct {

	// Like a three-word sentence: (attribute-name) (operator) (target-value).
	// Required: true
	Predicates []*JourneyAggregateQueryPredicate `json:"predicates"`

	// Boolean operation to apply to the provided predicates
	// Required: true
	// Enum: [and or]
	Type *string `json:"type"`
}

// Validate validates this journey aggregate query clause
func (m *JourneyAggregateQueryClause) Validate(formats strfmt.Registry) error {
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

func (m *JourneyAggregateQueryClause) validatePredicates(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

var journeyAggregateQueryClauseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["and","or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journeyAggregateQueryClauseTypeTypePropEnum = append(journeyAggregateQueryClauseTypeTypePropEnum, v)
	}
}

const (

	// JourneyAggregateQueryClauseTypeAnd captures enum value "and"
	JourneyAggregateQueryClauseTypeAnd string = "and"

	// JourneyAggregateQueryClauseTypeOr captures enum value "or"
	JourneyAggregateQueryClauseTypeOr string = "or"
)

// prop value enum
func (m *JourneyAggregateQueryClause) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journeyAggregateQueryClauseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JourneyAggregateQueryClause) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JourneyAggregateQueryClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JourneyAggregateQueryClause) UnmarshalBinary(b []byte) error {
	var res JourneyAggregateQueryClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
