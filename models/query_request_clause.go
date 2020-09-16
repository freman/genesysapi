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

// QueryRequestClause query request clause
//
// swagger:model QueryRequestClause
type QueryRequestClause struct {

	// The list of predicates used to filter the data
	// Required: true
	Predicates []*QueryRequestPredicate `json:"predicates"`

	// The logic used to combine the predicates
	// Required: true
	// Enum: [And Or]
	Type *string `json:"type"`
}

// Validate validates this query request clause
func (m *QueryRequestClause) Validate(formats strfmt.Registry) error {
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

func (m *QueryRequestClause) validatePredicates(formats strfmt.Registry) error {

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

var queryRequestClauseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["And","Or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryRequestClauseTypeTypePropEnum = append(queryRequestClauseTypeTypePropEnum, v)
	}
}

const (

	// QueryRequestClauseTypeAnd captures enum value "And"
	QueryRequestClauseTypeAnd string = "And"

	// QueryRequestClauseTypeOr captures enum value "Or"
	QueryRequestClauseTypeOr string = "Or"
)

// prop value enum
func (m *QueryRequestClause) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, queryRequestClauseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QueryRequestClause) validateType(formats strfmt.Registry) error {

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
func (m *QueryRequestClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryRequestClause) UnmarshalBinary(b []byte) error {
	var res QueryRequestClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
