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

// ConversationDetailQueryClause conversation detail query clause
//
// swagger:model ConversationDetailQueryClause
type ConversationDetailQueryClause struct {

	// Like a three-word sentence: (attribute-name) (operator) (target-value).
	// Required: true
	Predicates []*ConversationDetailQueryPredicate `json:"predicates"`

	// Boolean operation to apply to the provided predicates
	// Required: true
	// Enum: [and or]
	Type *string `json:"type"`
}

// Validate validates this conversation detail query clause
func (m *ConversationDetailQueryClause) Validate(formats strfmt.Registry) error {
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

func (m *ConversationDetailQueryClause) validatePredicates(formats strfmt.Registry) error {

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

var conversationDetailQueryClauseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["and","or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationDetailQueryClauseTypeTypePropEnum = append(conversationDetailQueryClauseTypeTypePropEnum, v)
	}
}

const (

	// ConversationDetailQueryClauseTypeAnd captures enum value "and"
	ConversationDetailQueryClauseTypeAnd string = "and"

	// ConversationDetailQueryClauseTypeOr captures enum value "or"
	ConversationDetailQueryClauseTypeOr string = "or"
)

// prop value enum
func (m *ConversationDetailQueryClause) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationDetailQueryClauseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationDetailQueryClause) validateType(formats strfmt.Registry) error {

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
func (m *ConversationDetailQueryClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationDetailQueryClause) UnmarshalBinary(b []byte) error {
	var res ConversationDetailQueryClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}