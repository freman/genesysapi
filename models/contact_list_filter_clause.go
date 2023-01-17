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

// ContactListFilterClause contact list filter clause
//
// swagger:model ContactListFilterClause
type ContactListFilterClause struct {

	// How to join predicates together.
	// Enum: [AND OR]
	FilterType string `json:"filterType,omitempty"`

	// Conditions to filter the contacts by.
	Predicates []*ContactListFilterPredicate `json:"predicates"`
}

// Validate validates this contact list filter clause
func (m *ContactListFilterClause) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterType(formats); err != nil {
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

var contactListFilterClauseTypeFilterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactListFilterClauseTypeFilterTypePropEnum = append(contactListFilterClauseTypeFilterTypePropEnum, v)
	}
}

const (

	// ContactListFilterClauseFilterTypeAND captures enum value "AND"
	ContactListFilterClauseFilterTypeAND string = "AND"

	// ContactListFilterClauseFilterTypeOR captures enum value "OR"
	ContactListFilterClauseFilterTypeOR string = "OR"
)

// prop value enum
func (m *ContactListFilterClause) validateFilterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contactListFilterClauseTypeFilterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContactListFilterClause) validateFilterType(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFilterTypeEnum("filterType", "body", m.FilterType); err != nil {
		return err
	}

	return nil
}

func (m *ContactListFilterClause) validatePredicates(formats strfmt.Registry) error {
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

// ContextValidate validate this contact list filter clause based on the context it is used
func (m *ContactListFilterClause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePredicates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactListFilterClause) contextValidatePredicates(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ContactListFilterClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactListFilterClause) UnmarshalBinary(b []byte) error {
	var res ContactListFilterClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
