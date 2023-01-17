// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AttributeFilterItem attribute filter item
//
// swagger:model AttributeFilterItem
type AttributeFilterItem struct {

	// id
	ID string `json:"id,omitempty"`

	// operator
	// Enum: [IN RANGE EQUALS NOTEQUALS LESSTHAN LESSTHANEQUALS GREATERTHAN GREATERTHANEQUALS CONTAINS]
	Operator string `json:"operator,omitempty"`

	// values
	Values []string `json:"values"`
}

// Validate validates this attribute filter item
func (m *AttributeFilterItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var attributeFilterItemTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN","RANGE","EQUALS","NOTEQUALS","LESSTHAN","LESSTHANEQUALS","GREATERTHAN","GREATERTHANEQUALS","CONTAINS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attributeFilterItemTypeOperatorPropEnum = append(attributeFilterItemTypeOperatorPropEnum, v)
	}
}

const (

	// AttributeFilterItemOperatorIN captures enum value "IN"
	AttributeFilterItemOperatorIN string = "IN"

	// AttributeFilterItemOperatorRANGE captures enum value "RANGE"
	AttributeFilterItemOperatorRANGE string = "RANGE"

	// AttributeFilterItemOperatorEQUALS captures enum value "EQUALS"
	AttributeFilterItemOperatorEQUALS string = "EQUALS"

	// AttributeFilterItemOperatorNOTEQUALS captures enum value "NOTEQUALS"
	AttributeFilterItemOperatorNOTEQUALS string = "NOTEQUALS"

	// AttributeFilterItemOperatorLESSTHAN captures enum value "LESSTHAN"
	AttributeFilterItemOperatorLESSTHAN string = "LESSTHAN"

	// AttributeFilterItemOperatorLESSTHANEQUALS captures enum value "LESSTHANEQUALS"
	AttributeFilterItemOperatorLESSTHANEQUALS string = "LESSTHANEQUALS"

	// AttributeFilterItemOperatorGREATERTHAN captures enum value "GREATERTHAN"
	AttributeFilterItemOperatorGREATERTHAN string = "GREATERTHAN"

	// AttributeFilterItemOperatorGREATERTHANEQUALS captures enum value "GREATERTHANEQUALS"
	AttributeFilterItemOperatorGREATERTHANEQUALS string = "GREATERTHANEQUALS"

	// AttributeFilterItemOperatorCONTAINS captures enum value "CONTAINS"
	AttributeFilterItemOperatorCONTAINS string = "CONTAINS"
)

// prop value enum
func (m *AttributeFilterItem) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, attributeFilterItemTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttributeFilterItem) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this attribute filter item based on context it is used
func (m *AttributeFilterItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttributeFilterItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttributeFilterItem) UnmarshalBinary(b []byte) error {
	var res AttributeFilterItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
