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

// SearchCriteria search criteria
//
// swagger:model SearchCriteria
type SearchCriteria struct {

	// The end value of the range. This field is used for range search types.
	EndValue string `json:"endValue,omitempty"`

	// Field names to search against
	Fields []string `json:"fields"`

	// Groups multiple conditions
	Group []*SearchCriteria `json:"group"`

	// How to apply this search criteria against other criteria
	// Enum: [AND OR NOT]
	Operator string `json:"operator,omitempty"`

	// The start value of the range. This field is used for range search types.
	StartValue string `json:"startValue,omitempty"`

	// type
	// Enum: [EXACT CONTAINS STARTS_WITH REQUIRED_FIELDS RANGE DATE_RANGE LESS_THAN LESS_THAN_EQUAL_TO GREATER_THAN GREATER_THAN_EQUAL_TO SIMPLE TERM TERMS QUERY_STRING MATCH_ALL]
	Type string `json:"type,omitempty"`

	// A value for the search to match against
	Value string `json:"value,omitempty"`

	// A list of values for the search to match against
	Values []string `json:"values"`
}

// Validate validates this search criteria
func (m *SearchCriteria) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
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

func (m *SearchCriteria) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	for i := 0; i < len(m.Group); i++ {
		if swag.IsZero(m.Group[i]) { // not required
			continue
		}

		if m.Group[i] != nil {
			if err := m.Group[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var searchCriteriaTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR","NOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchCriteriaTypeOperatorPropEnum = append(searchCriteriaTypeOperatorPropEnum, v)
	}
}

const (

	// SearchCriteriaOperatorAND captures enum value "AND"
	SearchCriteriaOperatorAND string = "AND"

	// SearchCriteriaOperatorOR captures enum value "OR"
	SearchCriteriaOperatorOR string = "OR"

	// SearchCriteriaOperatorNOT captures enum value "NOT"
	SearchCriteriaOperatorNOT string = "NOT"
)

// prop value enum
func (m *SearchCriteria) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchCriteriaTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchCriteria) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var searchCriteriaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXACT","CONTAINS","STARTS_WITH","REQUIRED_FIELDS","RANGE","DATE_RANGE","LESS_THAN","LESS_THAN_EQUAL_TO","GREATER_THAN","GREATER_THAN_EQUAL_TO","SIMPLE","TERM","TERMS","QUERY_STRING","MATCH_ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchCriteriaTypeTypePropEnum = append(searchCriteriaTypeTypePropEnum, v)
	}
}

const (

	// SearchCriteriaTypeEXACT captures enum value "EXACT"
	SearchCriteriaTypeEXACT string = "EXACT"

	// SearchCriteriaTypeCONTAINS captures enum value "CONTAINS"
	SearchCriteriaTypeCONTAINS string = "CONTAINS"

	// SearchCriteriaTypeSTARTSWITH captures enum value "STARTS_WITH"
	SearchCriteriaTypeSTARTSWITH string = "STARTS_WITH"

	// SearchCriteriaTypeREQUIREDFIELDS captures enum value "REQUIRED_FIELDS"
	SearchCriteriaTypeREQUIREDFIELDS string = "REQUIRED_FIELDS"

	// SearchCriteriaTypeRANGE captures enum value "RANGE"
	SearchCriteriaTypeRANGE string = "RANGE"

	// SearchCriteriaTypeDATERANGE captures enum value "DATE_RANGE"
	SearchCriteriaTypeDATERANGE string = "DATE_RANGE"

	// SearchCriteriaTypeLESSTHAN captures enum value "LESS_THAN"
	SearchCriteriaTypeLESSTHAN string = "LESS_THAN"

	// SearchCriteriaTypeLESSTHANEQUALTO captures enum value "LESS_THAN_EQUAL_TO"
	SearchCriteriaTypeLESSTHANEQUALTO string = "LESS_THAN_EQUAL_TO"

	// SearchCriteriaTypeGREATERTHAN captures enum value "GREATER_THAN"
	SearchCriteriaTypeGREATERTHAN string = "GREATER_THAN"

	// SearchCriteriaTypeGREATERTHANEQUALTO captures enum value "GREATER_THAN_EQUAL_TO"
	SearchCriteriaTypeGREATERTHANEQUALTO string = "GREATER_THAN_EQUAL_TO"

	// SearchCriteriaTypeSIMPLE captures enum value "SIMPLE"
	SearchCriteriaTypeSIMPLE string = "SIMPLE"

	// SearchCriteriaTypeTERM captures enum value "TERM"
	SearchCriteriaTypeTERM string = "TERM"

	// SearchCriteriaTypeTERMS captures enum value "TERMS"
	SearchCriteriaTypeTERMS string = "TERMS"

	// SearchCriteriaTypeQUERYSTRING captures enum value "QUERY_STRING"
	SearchCriteriaTypeQUERYSTRING string = "QUERY_STRING"

	// SearchCriteriaTypeMATCHALL captures enum value "MATCH_ALL"
	SearchCriteriaTypeMATCHALL string = "MATCH_ALL"
)

// prop value enum
func (m *SearchCriteria) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchCriteriaTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchCriteria) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchCriteria) UnmarshalBinary(b []byte) error {
	var res SearchCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
