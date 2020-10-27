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

// TranscriptConversationDetailSearchCriteria transcript conversation detail search criteria
//
// swagger:model TranscriptConversationDetailSearchCriteria
type TranscriptConversationDetailSearchCriteria struct {

	// The end value of the range. This field is used for range search types.
	EndValue string `json:"endValue,omitempty"`

	// Field names to search against
	Fields []string `json:"fields"`

	// Groups multiple conditions
	Group []*TranscriptConversationDetailSearchCriteria `json:"group"`

	// How to apply this search criteria against other criteria
	// Enum: [AND OR NOT]
	Operator string `json:"operator,omitempty"`

	// The start value of the range. This field is used for range search types.
	StartValue string `json:"startValue,omitempty"`

	// type
	// Enum: [EXACT EXACT_PHRASE PHRASE DATE_RANGE]
	Type string `json:"type,omitempty"`

	// A value for the search to match against
	Value string `json:"value,omitempty"`

	// A list of values for the search to match against
	Values []string `json:"values"`
}

// Validate validates this transcript conversation detail search criteria
func (m *TranscriptConversationDetailSearchCriteria) Validate(formats strfmt.Registry) error {
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

func (m *TranscriptConversationDetailSearchCriteria) validateGroup(formats strfmt.Registry) error {

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

var transcriptConversationDetailSearchCriteriaTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR","NOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptConversationDetailSearchCriteriaTypeOperatorPropEnum = append(transcriptConversationDetailSearchCriteriaTypeOperatorPropEnum, v)
	}
}

const (

	// TranscriptConversationDetailSearchCriteriaOperatorAND captures enum value "AND"
	TranscriptConversationDetailSearchCriteriaOperatorAND string = "AND"

	// TranscriptConversationDetailSearchCriteriaOperatorOR captures enum value "OR"
	TranscriptConversationDetailSearchCriteriaOperatorOR string = "OR"

	// TranscriptConversationDetailSearchCriteriaOperatorNOT captures enum value "NOT"
	TranscriptConversationDetailSearchCriteriaOperatorNOT string = "NOT"
)

// prop value enum
func (m *TranscriptConversationDetailSearchCriteria) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptConversationDetailSearchCriteriaTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptConversationDetailSearchCriteria) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var transcriptConversationDetailSearchCriteriaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXACT","EXACT_PHRASE","PHRASE","DATE_RANGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptConversationDetailSearchCriteriaTypeTypePropEnum = append(transcriptConversationDetailSearchCriteriaTypeTypePropEnum, v)
	}
}

const (

	// TranscriptConversationDetailSearchCriteriaTypeEXACT captures enum value "EXACT"
	TranscriptConversationDetailSearchCriteriaTypeEXACT string = "EXACT"

	// TranscriptConversationDetailSearchCriteriaTypeEXACTPHRASE captures enum value "EXACT_PHRASE"
	TranscriptConversationDetailSearchCriteriaTypeEXACTPHRASE string = "EXACT_PHRASE"

	// TranscriptConversationDetailSearchCriteriaTypePHRASE captures enum value "PHRASE"
	TranscriptConversationDetailSearchCriteriaTypePHRASE string = "PHRASE"

	// TranscriptConversationDetailSearchCriteriaTypeDATERANGE captures enum value "DATE_RANGE"
	TranscriptConversationDetailSearchCriteriaTypeDATERANGE string = "DATE_RANGE"
)

// prop value enum
func (m *TranscriptConversationDetailSearchCriteria) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptConversationDetailSearchCriteriaTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptConversationDetailSearchCriteria) validateType(formats strfmt.Registry) error {

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
func (m *TranscriptConversationDetailSearchCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranscriptConversationDetailSearchCriteria) UnmarshalBinary(b []byte) error {
	var res TranscriptConversationDetailSearchCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
