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

// GKNDocumentationSearchCriteria g k n documentation search criteria
//
// swagger:model GKNDocumentationSearchCriteria
type GKNDocumentationSearchCriteria struct {

	// Set date format for criteria values when using date range search type.  Supports Java date format syntax, example yyyy-MM-dd'T'HH:mm:ss.SSSX.
	DateFormat string `json:"dateFormat,omitempty"`

	// The end value of the range. This field is used for range search types.
	EndValue string `json:"endValue,omitempty"`

	// Field names to search against
	Fields []string `json:"fields"`

	// Groups multiple conditions
	Group []*GKNDocumentationSearchCriteria `json:"group"`

	// How to apply this search criteria against other criteria
	// Enum: [AND OR NOT]
	Operator string `json:"operator,omitempty"`

	// The start value of the range. This field is used for range search types.
	StartValue string `json:"startValue,omitempty"`

	// Search Type
	// Required: true
	// Enum: [SIMPLE]
	Type *string `json:"type"`

	// A value for the search to match against
	Value string `json:"value,omitempty"`

	// A list of values for the search to match against
	Values []string `json:"values"`
}

// Validate validates this g k n documentation search criteria
func (m *GKNDocumentationSearchCriteria) Validate(formats strfmt.Registry) error {
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

func (m *GKNDocumentationSearchCriteria) validateGroup(formats strfmt.Registry) error {
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

var gKNDocumentationSearchCriteriaTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR","NOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gKNDocumentationSearchCriteriaTypeOperatorPropEnum = append(gKNDocumentationSearchCriteriaTypeOperatorPropEnum, v)
	}
}

const (

	// GKNDocumentationSearchCriteriaOperatorAND captures enum value "AND"
	GKNDocumentationSearchCriteriaOperatorAND string = "AND"

	// GKNDocumentationSearchCriteriaOperatorOR captures enum value "OR"
	GKNDocumentationSearchCriteriaOperatorOR string = "OR"

	// GKNDocumentationSearchCriteriaOperatorNOT captures enum value "NOT"
	GKNDocumentationSearchCriteriaOperatorNOT string = "NOT"
)

// prop value enum
func (m *GKNDocumentationSearchCriteria) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gKNDocumentationSearchCriteriaTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GKNDocumentationSearchCriteria) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var gKNDocumentationSearchCriteriaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SIMPLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gKNDocumentationSearchCriteriaTypeTypePropEnum = append(gKNDocumentationSearchCriteriaTypeTypePropEnum, v)
	}
}

const (

	// GKNDocumentationSearchCriteriaTypeSIMPLE captures enum value "SIMPLE"
	GKNDocumentationSearchCriteriaTypeSIMPLE string = "SIMPLE"
)

// prop value enum
func (m *GKNDocumentationSearchCriteria) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gKNDocumentationSearchCriteriaTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GKNDocumentationSearchCriteria) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this g k n documentation search criteria based on the context it is used
func (m *GKNDocumentationSearchCriteria) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GKNDocumentationSearchCriteria) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GKNDocumentationSearchCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GKNDocumentationSearchCriteria) UnmarshalBinary(b []byte) error {
	var res GKNDocumentationSearchCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
