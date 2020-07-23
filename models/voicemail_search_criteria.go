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

// VoicemailSearchCriteria voicemail search criteria
//
// swagger:model VoicemailSearchCriteria
type VoicemailSearchCriteria struct {

	// The end value of the range. This field is used for range search types.
	EndValue string `json:"endValue,omitempty"`

	// Field names to search against
	Fields []string `json:"fields"`

	// Groups multiple conditions
	Group []*VoicemailSearchCriteria `json:"group"`

	// How to apply this search criteria against other criteria
	// Enum: [AND OR NOT]
	Operator string `json:"operator,omitempty"`

	// The start value of the range. This field is used for range search types.
	StartValue string `json:"startValue,omitempty"`

	// Search Type
	// Required: true
	// Enum: [EXACT STARTS_WITH CONTAINS REGEX TERM TERMS REQUIRED_FIELDS MATCH_ALL]
	Type *string `json:"type"`

	// A value for the search to match against
	Value string `json:"value,omitempty"`

	// A list of values for the search to match against
	Values []string `json:"values"`
}

// Validate validates this voicemail search criteria
func (m *VoicemailSearchCriteria) Validate(formats strfmt.Registry) error {
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

func (m *VoicemailSearchCriteria) validateGroup(formats strfmt.Registry) error {

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

var voicemailSearchCriteriaTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR","NOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		voicemailSearchCriteriaTypeOperatorPropEnum = append(voicemailSearchCriteriaTypeOperatorPropEnum, v)
	}
}

const (

	// VoicemailSearchCriteriaOperatorAND captures enum value "AND"
	VoicemailSearchCriteriaOperatorAND string = "AND"

	// VoicemailSearchCriteriaOperatorOR captures enum value "OR"
	VoicemailSearchCriteriaOperatorOR string = "OR"

	// VoicemailSearchCriteriaOperatorNOT captures enum value "NOT"
	VoicemailSearchCriteriaOperatorNOT string = "NOT"
)

// prop value enum
func (m *VoicemailSearchCriteria) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, voicemailSearchCriteriaTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VoicemailSearchCriteria) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var voicemailSearchCriteriaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXACT","STARTS_WITH","CONTAINS","REGEX","TERM","TERMS","REQUIRED_FIELDS","MATCH_ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		voicemailSearchCriteriaTypeTypePropEnum = append(voicemailSearchCriteriaTypeTypePropEnum, v)
	}
}

const (

	// VoicemailSearchCriteriaTypeEXACT captures enum value "EXACT"
	VoicemailSearchCriteriaTypeEXACT string = "EXACT"

	// VoicemailSearchCriteriaTypeSTARTSWITH captures enum value "STARTS_WITH"
	VoicemailSearchCriteriaTypeSTARTSWITH string = "STARTS_WITH"

	// VoicemailSearchCriteriaTypeCONTAINS captures enum value "CONTAINS"
	VoicemailSearchCriteriaTypeCONTAINS string = "CONTAINS"

	// VoicemailSearchCriteriaTypeREGEX captures enum value "REGEX"
	VoicemailSearchCriteriaTypeREGEX string = "REGEX"

	// VoicemailSearchCriteriaTypeTERM captures enum value "TERM"
	VoicemailSearchCriteriaTypeTERM string = "TERM"

	// VoicemailSearchCriteriaTypeTERMS captures enum value "TERMS"
	VoicemailSearchCriteriaTypeTERMS string = "TERMS"

	// VoicemailSearchCriteriaTypeREQUIREDFIELDS captures enum value "REQUIRED_FIELDS"
	VoicemailSearchCriteriaTypeREQUIREDFIELDS string = "REQUIRED_FIELDS"

	// VoicemailSearchCriteriaTypeMATCHALL captures enum value "MATCH_ALL"
	VoicemailSearchCriteriaTypeMATCHALL string = "MATCH_ALL"
)

// prop value enum
func (m *VoicemailSearchCriteria) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, voicemailSearchCriteriaTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VoicemailSearchCriteria) validateType(formats strfmt.Registry) error {

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
func (m *VoicemailSearchCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoicemailSearchCriteria) UnmarshalBinary(b []byte) error {
	var res VoicemailSearchCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
