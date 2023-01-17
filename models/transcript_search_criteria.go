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

// TranscriptSearchCriteria transcript search criteria
//
// swagger:model TranscriptSearchCriteria
type TranscriptSearchCriteria struct {

	// Set date format for criteria values when using date range search type.  Supports Java date format syntax, example yyyy-MM-dd'T'HH:mm:ss.SSSX.
	DateFormat string `json:"dateFormat,omitempty"`

	// The end value of the range. This field is used for range search types.
	EndValue string `json:"endValue,omitempty"`

	// Field names to search against
	Fields []string `json:"fields"`

	// Groups multiple conditions
	Group []*TranscriptSearchCriteria `json:"group"`

	// How to apply this search criteria against other criteria
	// Enum: [AND OR NOT]
	Operator string `json:"operator,omitempty"`

	// The start value of the range. This field is used for range search types.
	StartValue string `json:"startValue,omitempty"`

	// type
	// Enum: [EXACT EXACT_PHRASE PHRASE DATE_RANGE RANGE GREATER_THAN LESS_THAN]
	Type string `json:"type,omitempty"`

	// A value for the search to match against
	Value string `json:"value,omitempty"`

	// A list of values for the search to match against
	Values []string `json:"values"`
}

// Validate validates this transcript search criteria
func (m *TranscriptSearchCriteria) Validate(formats strfmt.Registry) error {
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

func (m *TranscriptSearchCriteria) validateGroup(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var transcriptSearchCriteriaTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR","NOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptSearchCriteriaTypeOperatorPropEnum = append(transcriptSearchCriteriaTypeOperatorPropEnum, v)
	}
}

const (

	// TranscriptSearchCriteriaOperatorAND captures enum value "AND"
	TranscriptSearchCriteriaOperatorAND string = "AND"

	// TranscriptSearchCriteriaOperatorOR captures enum value "OR"
	TranscriptSearchCriteriaOperatorOR string = "OR"

	// TranscriptSearchCriteriaOperatorNOT captures enum value "NOT"
	TranscriptSearchCriteriaOperatorNOT string = "NOT"
)

// prop value enum
func (m *TranscriptSearchCriteria) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptSearchCriteriaTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptSearchCriteria) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var transcriptSearchCriteriaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXACT","EXACT_PHRASE","PHRASE","DATE_RANGE","RANGE","GREATER_THAN","LESS_THAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptSearchCriteriaTypeTypePropEnum = append(transcriptSearchCriteriaTypeTypePropEnum, v)
	}
}

const (

	// TranscriptSearchCriteriaTypeEXACT captures enum value "EXACT"
	TranscriptSearchCriteriaTypeEXACT string = "EXACT"

	// TranscriptSearchCriteriaTypeEXACTPHRASE captures enum value "EXACT_PHRASE"
	TranscriptSearchCriteriaTypeEXACTPHRASE string = "EXACT_PHRASE"

	// TranscriptSearchCriteriaTypePHRASE captures enum value "PHRASE"
	TranscriptSearchCriteriaTypePHRASE string = "PHRASE"

	// TranscriptSearchCriteriaTypeDATERANGE captures enum value "DATE_RANGE"
	TranscriptSearchCriteriaTypeDATERANGE string = "DATE_RANGE"

	// TranscriptSearchCriteriaTypeRANGE captures enum value "RANGE"
	TranscriptSearchCriteriaTypeRANGE string = "RANGE"

	// TranscriptSearchCriteriaTypeGREATERTHAN captures enum value "GREATER_THAN"
	TranscriptSearchCriteriaTypeGREATERTHAN string = "GREATER_THAN"

	// TranscriptSearchCriteriaTypeLESSTHAN captures enum value "LESS_THAN"
	TranscriptSearchCriteriaTypeLESSTHAN string = "LESS_THAN"
)

// prop value enum
func (m *TranscriptSearchCriteria) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptSearchCriteriaTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptSearchCriteria) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transcript search criteria based on the context it is used
func (m *TranscriptSearchCriteria) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TranscriptSearchCriteria) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Group); i++ {

		if m.Group[i] != nil {
			if err := m.Group[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("group" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TranscriptSearchCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranscriptSearchCriteria) UnmarshalBinary(b []byte) error {
	var res TranscriptSearchCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
