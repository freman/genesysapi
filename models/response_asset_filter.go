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

// ResponseAssetFilter response asset filter
//
// swagger:model ResponseAssetFilter
type ResponseAssetFilter struct {

	// The end value of the range. This field is used for range search types. Accepts numbers and date in ISO8601 format
	EndValue string `json:"endValue,omitempty"`

	// Field name to search against. Allowed Values: divisionId, name, contentLength, contentType, dateCreated
	Fields []string `json:"fields"`

	// The start value of the range. This field is used for range search types. Accepts numbers and date in ISO8601 format
	StartValue string `json:"startValue,omitempty"`

	// How to apply this search criteria against other criteria
	// Enum: [TERM TERMS STARTS_WITH RANGE GREATER_THAN_EQUAL_TO LESS_THAN_EQUAL_TO DATE_RANGE]
	Type string `json:"type,omitempty"`

	// A value for the search to match against
	Value string `json:"value,omitempty"`

	// A list of values for the search to match against
	Values []string `json:"values"`
}

// Validate validates this response asset filter
func (m *ResponseAssetFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var responseAssetFilterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TERM","TERMS","STARTS_WITH","RANGE","GREATER_THAN_EQUAL_TO","LESS_THAN_EQUAL_TO","DATE_RANGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responseAssetFilterTypeTypePropEnum = append(responseAssetFilterTypeTypePropEnum, v)
	}
}

const (

	// ResponseAssetFilterTypeTERM captures enum value "TERM"
	ResponseAssetFilterTypeTERM string = "TERM"

	// ResponseAssetFilterTypeTERMS captures enum value "TERMS"
	ResponseAssetFilterTypeTERMS string = "TERMS"

	// ResponseAssetFilterTypeSTARTSWITH captures enum value "STARTS_WITH"
	ResponseAssetFilterTypeSTARTSWITH string = "STARTS_WITH"

	// ResponseAssetFilterTypeRANGE captures enum value "RANGE"
	ResponseAssetFilterTypeRANGE string = "RANGE"

	// ResponseAssetFilterTypeGREATERTHANEQUALTO captures enum value "GREATER_THAN_EQUAL_TO"
	ResponseAssetFilterTypeGREATERTHANEQUALTO string = "GREATER_THAN_EQUAL_TO"

	// ResponseAssetFilterTypeLESSTHANEQUALTO captures enum value "LESS_THAN_EQUAL_TO"
	ResponseAssetFilterTypeLESSTHANEQUALTO string = "LESS_THAN_EQUAL_TO"

	// ResponseAssetFilterTypeDATERANGE captures enum value "DATE_RANGE"
	ResponseAssetFilterTypeDATERANGE string = "DATE_RANGE"
)

// prop value enum
func (m *ResponseAssetFilter) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responseAssetFilterTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponseAssetFilter) validateType(formats strfmt.Registry) error {

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
func (m *ResponseAssetFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseAssetFilter) UnmarshalBinary(b []byte) error {
	var res ResponseAssetFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}