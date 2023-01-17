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

// DocumentQueryClause document query clause
//
// swagger:model DocumentQueryClause
type DocumentQueryClause struct {

	// Specifies how the predicates will be applied together.
	// Required: true
	// Enum: [Or And]
	Operator *string `json:"operator"`

	// To apply multiple conditions. Limit of 10 predicates across all clauses.
	// Required: true
	Predicates []*DocumentQueryPredicate `json:"predicates"`
}

// Validate validates this document query clause
func (m *DocumentQueryClause) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePredicates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var documentQueryClauseTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Or","And"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentQueryClauseTypeOperatorPropEnum = append(documentQueryClauseTypeOperatorPropEnum, v)
	}
}

const (

	// DocumentQueryClauseOperatorOr captures enum value "Or"
	DocumentQueryClauseOperatorOr string = "Or"

	// DocumentQueryClauseOperatorAnd captures enum value "And"
	DocumentQueryClauseOperatorAnd string = "And"
)

// prop value enum
func (m *DocumentQueryClause) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentQueryClauseTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DocumentQueryClause) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", *m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *DocumentQueryClause) validatePredicates(formats strfmt.Registry) error {

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

// ContextValidate validate this document query clause based on the context it is used
func (m *DocumentQueryClause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePredicates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocumentQueryClause) contextValidatePredicates(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DocumentQueryClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentQueryClause) UnmarshalBinary(b []byte) error {
	var res DocumentQueryClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
