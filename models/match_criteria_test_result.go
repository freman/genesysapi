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

// MatchCriteriaTestResult Results of a matching expression
//
// swagger:model MatchCriteriaTestResult
type MatchCriteriaTestResult struct {

	// The generated json path condition
	GeneratedJSONPathCondition string `json:"generatedJsonPathCondition,omitempty"`

	// The Goessner json path of the field to match
	JSONPath string `json:"jsonPath,omitempty"`

	// The json paths and their values that were compared
	JSONPathExtraction []*MatchTestResult `json:"jsonPathExtraction"`

	// Did the generated json path condition match
	Match bool `json:"match"`

	// The type of operation to perform for matching check
	// Enum: [GreaterThanOrEqual LessThanOrEqual Equal NotEqual LessThan GreaterThan NotIn In Contains All Exists Size]
	Operator string `json:"operator,omitempty"`

	// The value to match on. Only one of value and values can be included
	Value JSONNode `json:"value,omitempty"`

	// The list of values to match on. Only one of value and values can be included
	Values []JSONNode `json:"values"`
}

// Validate validates this match criteria test result
func (m *MatchCriteriaTestResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJSONPathExtraction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MatchCriteriaTestResult) validateJSONPathExtraction(formats strfmt.Registry) error {
	if swag.IsZero(m.JSONPathExtraction) { // not required
		return nil
	}

	for i := 0; i < len(m.JSONPathExtraction); i++ {
		if swag.IsZero(m.JSONPathExtraction[i]) { // not required
			continue
		}

		if m.JSONPathExtraction[i] != nil {
			if err := m.JSONPathExtraction[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jsonPathExtraction" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jsonPathExtraction" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var matchCriteriaTestResultTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GreaterThanOrEqual","LessThanOrEqual","Equal","NotEqual","LessThan","GreaterThan","NotIn","In","Contains","All","Exists","Size"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		matchCriteriaTestResultTypeOperatorPropEnum = append(matchCriteriaTestResultTypeOperatorPropEnum, v)
	}
}

const (

	// MatchCriteriaTestResultOperatorGreaterThanOrEqual captures enum value "GreaterThanOrEqual"
	MatchCriteriaTestResultOperatorGreaterThanOrEqual string = "GreaterThanOrEqual"

	// MatchCriteriaTestResultOperatorLessThanOrEqual captures enum value "LessThanOrEqual"
	MatchCriteriaTestResultOperatorLessThanOrEqual string = "LessThanOrEqual"

	// MatchCriteriaTestResultOperatorEqual captures enum value "Equal"
	MatchCriteriaTestResultOperatorEqual string = "Equal"

	// MatchCriteriaTestResultOperatorNotEqual captures enum value "NotEqual"
	MatchCriteriaTestResultOperatorNotEqual string = "NotEqual"

	// MatchCriteriaTestResultOperatorLessThan captures enum value "LessThan"
	MatchCriteriaTestResultOperatorLessThan string = "LessThan"

	// MatchCriteriaTestResultOperatorGreaterThan captures enum value "GreaterThan"
	MatchCriteriaTestResultOperatorGreaterThan string = "GreaterThan"

	// MatchCriteriaTestResultOperatorNotIn captures enum value "NotIn"
	MatchCriteriaTestResultOperatorNotIn string = "NotIn"

	// MatchCriteriaTestResultOperatorIn captures enum value "In"
	MatchCriteriaTestResultOperatorIn string = "In"

	// MatchCriteriaTestResultOperatorContains captures enum value "Contains"
	MatchCriteriaTestResultOperatorContains string = "Contains"

	// MatchCriteriaTestResultOperatorAll captures enum value "All"
	MatchCriteriaTestResultOperatorAll string = "All"

	// MatchCriteriaTestResultOperatorExists captures enum value "Exists"
	MatchCriteriaTestResultOperatorExists string = "Exists"

	// MatchCriteriaTestResultOperatorSize captures enum value "Size"
	MatchCriteriaTestResultOperatorSize string = "Size"
)

// prop value enum
func (m *MatchCriteriaTestResult) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, matchCriteriaTestResultTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MatchCriteriaTestResult) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this match criteria test result based on the context it is used
func (m *MatchCriteriaTestResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJSONPathExtraction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MatchCriteriaTestResult) contextValidateJSONPathExtraction(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.JSONPathExtraction); i++ {

		if m.JSONPathExtraction[i] != nil {
			if err := m.JSONPathExtraction[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jsonPathExtraction" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jsonPathExtraction" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MatchCriteriaTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchCriteriaTestResult) UnmarshalBinary(b []byte) error {
	var res MatchCriteriaTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
