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

// SearchAggregation search aggregation
//
// swagger:model SearchAggregation
type SearchAggregation struct {

	// The field used for aggregation
	Field string `json:"field,omitempty"`

	// The name of the aggregation. The response aggregation uses this name.
	Name string `json:"name,omitempty"`

	// The order in which aggregation results are sorted
	Order []string `json:"order"`

	// The number aggregations results to return out of the entire result set
	Size int32 `json:"size,omitempty"`

	// The type of aggregation to perform
	// Enum: [COUNT SUM AVERAGE TERM CONTAINS STARTS_WITH ENDS_WITH]
	Type string `json:"type,omitempty"`

	// A value to use for aggregation
	Value string `json:"value,omitempty"`
}

// Validate validates this search aggregation
func (m *SearchAggregation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrder(formats); err != nil {
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

var searchAggregationOrderItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VALUE_DESC","VALUE_ASC","COUNT_DESC","COUNT_ASC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchAggregationOrderItemsEnum = append(searchAggregationOrderItemsEnum, v)
	}
}

func (m *SearchAggregation) validateOrderItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchAggregationOrderItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchAggregation) validateOrder(formats strfmt.Registry) error {
	if swag.IsZero(m.Order) { // not required
		return nil
	}

	for i := 0; i < len(m.Order); i++ {

		// value enum
		if err := m.validateOrderItemsEnum("order"+"."+strconv.Itoa(i), "body", m.Order[i]); err != nil {
			return err
		}

	}

	return nil
}

var searchAggregationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COUNT","SUM","AVERAGE","TERM","CONTAINS","STARTS_WITH","ENDS_WITH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchAggregationTypeTypePropEnum = append(searchAggregationTypeTypePropEnum, v)
	}
}

const (

	// SearchAggregationTypeCOUNT captures enum value "COUNT"
	SearchAggregationTypeCOUNT string = "COUNT"

	// SearchAggregationTypeSUM captures enum value "SUM"
	SearchAggregationTypeSUM string = "SUM"

	// SearchAggregationTypeAVERAGE captures enum value "AVERAGE"
	SearchAggregationTypeAVERAGE string = "AVERAGE"

	// SearchAggregationTypeTERM captures enum value "TERM"
	SearchAggregationTypeTERM string = "TERM"

	// SearchAggregationTypeCONTAINS captures enum value "CONTAINS"
	SearchAggregationTypeCONTAINS string = "CONTAINS"

	// SearchAggregationTypeSTARTSWITH captures enum value "STARTS_WITH"
	SearchAggregationTypeSTARTSWITH string = "STARTS_WITH"

	// SearchAggregationTypeENDSWITH captures enum value "ENDS_WITH"
	SearchAggregationTypeENDSWITH string = "ENDS_WITH"
)

// prop value enum
func (m *SearchAggregation) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchAggregationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchAggregation) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this search aggregation based on context it is used
func (m *SearchAggregation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchAggregation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchAggregation) UnmarshalBinary(b []byte) error {
	var res SearchAggregation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
