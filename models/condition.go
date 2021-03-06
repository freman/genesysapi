// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Condition condition
//
// swagger:model Condition
type Condition struct {

	// An attribute name associated with this Condition. Required for a contactAttributeCondition.
	AttributeName string `json:"attributeName,omitempty"`

	// List of wrap-up code identifiers. Required for a wrapupCondition.
	Codes []string `json:"codes"`

	// If true, inverts the result of evaluating this Condition. Default is false.
	Inverted bool `json:"inverted"`

	// An operation with which to evaluate the Condition. Not used for a DataActionCondition.
	// Enum: [EQUALS LESS_THAN LESS_THAN_EQUALS GREATER_THAN GREATER_THAN_EQUALS CONTAINS BEGINS_WITH ENDS_WITH BEFORE AFTER IN]
	Operator string `json:"operator,omitempty"`

	// A value associated with the property type of this Condition. Required for a contactPropertyCondition.
	Property string `json:"property,omitempty"`

	// The type of the property associated with this Condition. Required for a contactPropertyCondition.
	// Enum: [LAST_ATTEMPT_BY_COLUMN LAST_ATTEMPT_OVERALL LAST_WRAPUP_BY_COLUMN LAST_WRAPUP_OVERALL]
	PropertyType string `json:"propertyType,omitempty"`

	// The type of the condition.
	// Enum: [wrapupCondition contactAttributeCondition phoneNumberCondition phoneNumberTypeCondition callAnalysisCondition contactPropertyCondition dataActionCondition]
	Type string `json:"type,omitempty"`

	// A value associated with this Condition. This could be text, a number, or a relative time. Not used for a DataActionCondition.
	Value string `json:"value,omitempty"`

	// The type of the value associated with this Condition. Not used for a DataActionCondition.
	// Enum: [STRING NUMERIC DATETIME PERIOD]
	ValueType string `json:"valueType,omitempty"`
}

// Validate validates this condition
func (m *Condition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conditionTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EQUALS","LESS_THAN","LESS_THAN_EQUALS","GREATER_THAN","GREATER_THAN_EQUALS","CONTAINS","BEGINS_WITH","ENDS_WITH","BEFORE","AFTER","IN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeOperatorPropEnum = append(conditionTypeOperatorPropEnum, v)
	}
}

const (

	// ConditionOperatorEQUALS captures enum value "EQUALS"
	ConditionOperatorEQUALS string = "EQUALS"

	// ConditionOperatorLESSTHAN captures enum value "LESS_THAN"
	ConditionOperatorLESSTHAN string = "LESS_THAN"

	// ConditionOperatorLESSTHANEQUALS captures enum value "LESS_THAN_EQUALS"
	ConditionOperatorLESSTHANEQUALS string = "LESS_THAN_EQUALS"

	// ConditionOperatorGREATERTHAN captures enum value "GREATER_THAN"
	ConditionOperatorGREATERTHAN string = "GREATER_THAN"

	// ConditionOperatorGREATERTHANEQUALS captures enum value "GREATER_THAN_EQUALS"
	ConditionOperatorGREATERTHANEQUALS string = "GREATER_THAN_EQUALS"

	// ConditionOperatorCONTAINS captures enum value "CONTAINS"
	ConditionOperatorCONTAINS string = "CONTAINS"

	// ConditionOperatorBEGINSWITH captures enum value "BEGINS_WITH"
	ConditionOperatorBEGINSWITH string = "BEGINS_WITH"

	// ConditionOperatorENDSWITH captures enum value "ENDS_WITH"
	ConditionOperatorENDSWITH string = "ENDS_WITH"

	// ConditionOperatorBEFORE captures enum value "BEFORE"
	ConditionOperatorBEFORE string = "BEFORE"

	// ConditionOperatorAFTER captures enum value "AFTER"
	ConditionOperatorAFTER string = "AFTER"

	// ConditionOperatorIN captures enum value "IN"
	ConditionOperatorIN string = "IN"
)

// prop value enum
func (m *Condition) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conditionTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Condition) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var conditionTypePropertyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LAST_ATTEMPT_BY_COLUMN","LAST_ATTEMPT_OVERALL","LAST_WRAPUP_BY_COLUMN","LAST_WRAPUP_OVERALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypePropertyTypePropEnum = append(conditionTypePropertyTypePropEnum, v)
	}
}

const (

	// ConditionPropertyTypeLASTATTEMPTBYCOLUMN captures enum value "LAST_ATTEMPT_BY_COLUMN"
	ConditionPropertyTypeLASTATTEMPTBYCOLUMN string = "LAST_ATTEMPT_BY_COLUMN"

	// ConditionPropertyTypeLASTATTEMPTOVERALL captures enum value "LAST_ATTEMPT_OVERALL"
	ConditionPropertyTypeLASTATTEMPTOVERALL string = "LAST_ATTEMPT_OVERALL"

	// ConditionPropertyTypeLASTWRAPUPBYCOLUMN captures enum value "LAST_WRAPUP_BY_COLUMN"
	ConditionPropertyTypeLASTWRAPUPBYCOLUMN string = "LAST_WRAPUP_BY_COLUMN"

	// ConditionPropertyTypeLASTWRAPUPOVERALL captures enum value "LAST_WRAPUP_OVERALL"
	ConditionPropertyTypeLASTWRAPUPOVERALL string = "LAST_WRAPUP_OVERALL"
)

// prop value enum
func (m *Condition) validatePropertyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conditionTypePropertyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Condition) validatePropertyType(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePropertyTypeEnum("propertyType", "body", m.PropertyType); err != nil {
		return err
	}

	return nil
}

var conditionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wrapupCondition","contactAttributeCondition","phoneNumberCondition","phoneNumberTypeCondition","callAnalysisCondition","contactPropertyCondition","dataActionCondition"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeTypePropEnum = append(conditionTypeTypePropEnum, v)
	}
}

const (

	// ConditionTypeWrapupCondition captures enum value "wrapupCondition"
	ConditionTypeWrapupCondition string = "wrapupCondition"

	// ConditionTypeContactAttributeCondition captures enum value "contactAttributeCondition"
	ConditionTypeContactAttributeCondition string = "contactAttributeCondition"

	// ConditionTypePhoneNumberCondition captures enum value "phoneNumberCondition"
	ConditionTypePhoneNumberCondition string = "phoneNumberCondition"

	// ConditionTypePhoneNumberTypeCondition captures enum value "phoneNumberTypeCondition"
	ConditionTypePhoneNumberTypeCondition string = "phoneNumberTypeCondition"

	// ConditionTypeCallAnalysisCondition captures enum value "callAnalysisCondition"
	ConditionTypeCallAnalysisCondition string = "callAnalysisCondition"

	// ConditionTypeContactPropertyCondition captures enum value "contactPropertyCondition"
	ConditionTypeContactPropertyCondition string = "contactPropertyCondition"

	// ConditionTypeDataActionCondition captures enum value "dataActionCondition"
	ConditionTypeDataActionCondition string = "dataActionCondition"
)

// prop value enum
func (m *Condition) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conditionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Condition) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

var conditionTypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STRING","NUMERIC","DATETIME","PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeValueTypePropEnum = append(conditionTypeValueTypePropEnum, v)
	}
}

const (

	// ConditionValueTypeSTRING captures enum value "STRING"
	ConditionValueTypeSTRING string = "STRING"

	// ConditionValueTypeNUMERIC captures enum value "NUMERIC"
	ConditionValueTypeNUMERIC string = "NUMERIC"

	// ConditionValueTypeDATETIME captures enum value "DATETIME"
	ConditionValueTypeDATETIME string = "DATETIME"

	// ConditionValueTypePERIOD captures enum value "PERIOD"
	ConditionValueTypePERIOD string = "PERIOD"
)

// prop value enum
func (m *Condition) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conditionTypeValueTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Condition) validateValueType(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueType) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueTypeEnum("valueType", "body", m.ValueType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Condition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Condition) UnmarshalBinary(b []byte) error {
	var res Condition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
