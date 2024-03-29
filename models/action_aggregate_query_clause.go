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

// ActionAggregateQueryClause action aggregate query clause
//
// swagger:model ActionAggregateQueryClause
type ActionAggregateQueryClause struct {

	// Like a three-word sentence: (attribute-name) (operator) (target-value).
	// Required: true
	Predicates []*ActionAggregateQueryPredicate `json:"predicates"`

	// Boolean operation to apply to the provided predicates
	// Required: true
	// Enum: [and or]
	Type *string `json:"type"`
}

// Validate validates this action aggregate query clause
func (m *ActionAggregateQueryClause) Validate(formats strfmt.Registry) error {
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

func (m *ActionAggregateQueryClause) validatePredicates(formats strfmt.Registry) error {

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

var actionAggregateQueryClauseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["and","or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionAggregateQueryClauseTypeTypePropEnum = append(actionAggregateQueryClauseTypeTypePropEnum, v)
	}
}

const (

	// ActionAggregateQueryClauseTypeAnd captures enum value "and"
	ActionAggregateQueryClauseTypeAnd string = "and"

	// ActionAggregateQueryClauseTypeOr captures enum value "or"
	ActionAggregateQueryClauseTypeOr string = "or"
)

// prop value enum
func (m *ActionAggregateQueryClause) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, actionAggregateQueryClauseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActionAggregateQueryClause) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this action aggregate query clause based on the context it is used
func (m *ActionAggregateQueryClause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePredicates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionAggregateQueryClause) contextValidatePredicates(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ActionAggregateQueryClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionAggregateQueryClause) UnmarshalBinary(b []byte) error {
	var res ActionAggregateQueryClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
