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

// DocumentQueryPredicate document query predicate
//
// swagger:model DocumentQueryPredicate
type DocumentQueryPredicate struct {

	// Specifies the document fields to be matched against.
	// Required: true
	Fields []string `json:"fields"`

	// Specifies the matching criteria between the fields and values.
	// Required: true
	// Enum: [Equals NotEquals Contains MatchAll MatchAny]
	Type *string `json:"type"`

	// Specifies the values of the fields to be matched against.
	// Required: true
	Values []string `json:"values"`
}

// Validate validates this document query predicate
func (m *DocumentQueryPredicate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var documentQueryPredicateFieldsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alternatives","categoryId","categoryName","contextId","contextName","contextValueId","contextValueName","documentId","labelId","labelName","title"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentQueryPredicateFieldsItemsEnum = append(documentQueryPredicateFieldsItemsEnum, v)
	}
}

func (m *DocumentQueryPredicate) validateFieldsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentQueryPredicateFieldsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DocumentQueryPredicate) validateFields(formats strfmt.Registry) error {

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	for i := 0; i < len(m.Fields); i++ {

		// value enum
		if err := m.validateFieldsItemsEnum("fields"+"."+strconv.Itoa(i), "body", m.Fields[i]); err != nil {
			return err
		}

	}

	return nil
}

var documentQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Equals","NotEquals","Contains","MatchAll","MatchAny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentQueryPredicateTypeTypePropEnum = append(documentQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// DocumentQueryPredicateTypeEquals captures enum value "Equals"
	DocumentQueryPredicateTypeEquals string = "Equals"

	// DocumentQueryPredicateTypeNotEquals captures enum value "NotEquals"
	DocumentQueryPredicateTypeNotEquals string = "NotEquals"

	// DocumentQueryPredicateTypeContains captures enum value "Contains"
	DocumentQueryPredicateTypeContains string = "Contains"

	// DocumentQueryPredicateTypeMatchAll captures enum value "MatchAll"
	DocumentQueryPredicateTypeMatchAll string = "MatchAll"

	// DocumentQueryPredicateTypeMatchAny captures enum value "MatchAny"
	DocumentQueryPredicateTypeMatchAny string = "MatchAny"
)

// prop value enum
func (m *DocumentQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DocumentQueryPredicate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DocumentQueryPredicate) validateValues(formats strfmt.Registry) error {

	if err := validate.Required("values", "body", m.Values); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this document query predicate based on context it is used
func (m *DocumentQueryPredicate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DocumentQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentQueryPredicate) UnmarshalBinary(b []byte) error {
	var res DocumentQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
