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

// PresenceDetailQueryFilter presence detail query filter
//
// swagger:model PresenceDetailQueryFilter
type PresenceDetailQueryFilter struct {

	// Boolean 'and/or' logic with up to two-levels of nesting
	Clauses []*PresenceDetailQueryClause `json:"clauses"`

	// Like a three-word sentence: (attribute-name) (operator) (target-value).
	Predicates []*PresenceDetailQueryPredicate `json:"predicates"`

	// Boolean operation to apply to the provided predicates and clauses
	// Required: true
	// Enum: [and or]
	Type *string `json:"type"`
}

// Validate validates this presence detail query filter
func (m *PresenceDetailQueryFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClauses(formats); err != nil {
		res = append(res, err)
	}

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

func (m *PresenceDetailQueryFilter) validateClauses(formats strfmt.Registry) error {
	if swag.IsZero(m.Clauses) { // not required
		return nil
	}

	for i := 0; i < len(m.Clauses); i++ {
		if swag.IsZero(m.Clauses[i]) { // not required
			continue
		}

		if m.Clauses[i] != nil {
			if err := m.Clauses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clauses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clauses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PresenceDetailQueryFilter) validatePredicates(formats strfmt.Registry) error {
	if swag.IsZero(m.Predicates) { // not required
		return nil
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

var presenceDetailQueryFilterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["and","or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		presenceDetailQueryFilterTypeTypePropEnum = append(presenceDetailQueryFilterTypeTypePropEnum, v)
	}
}

const (

	// PresenceDetailQueryFilterTypeAnd captures enum value "and"
	PresenceDetailQueryFilterTypeAnd string = "and"

	// PresenceDetailQueryFilterTypeOr captures enum value "or"
	PresenceDetailQueryFilterTypeOr string = "or"
)

// prop value enum
func (m *PresenceDetailQueryFilter) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, presenceDetailQueryFilterTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PresenceDetailQueryFilter) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this presence detail query filter based on the context it is used
func (m *PresenceDetailQueryFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClauses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePredicates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PresenceDetailQueryFilter) contextValidateClauses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clauses); i++ {

		if m.Clauses[i] != nil {
			if err := m.Clauses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clauses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clauses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PresenceDetailQueryFilter) contextValidatePredicates(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PresenceDetailQueryFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PresenceDetailQueryFilter) UnmarshalBinary(b []byte) error {
	var res PresenceDetailQueryFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
