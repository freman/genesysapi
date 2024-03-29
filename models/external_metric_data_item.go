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

// ExternalMetricDataItem external metric data item
//
// swagger:model ExternalMetricDataItem
type ExternalMetricDataItem struct {

	// The number of data points. The default value is 0 when type is Cumulative and the metric data already exists, otherwise 1. When total count reaches 0, the metric data will be deleted.
	Count int32 `json:"count,omitempty"`

	// The date of the metric data. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Required: true
	// Format: date
	DateOccurred *strfmt.Date `json:"dateOccurred"`

	// The ID of the external metric definition
	// Required: true
	MetricID *string `json:"metricId"`

	// The type of the metric data. The default value is Total.
	// Enum: [Total Cumulative]
	Type string `json:"type,omitempty"`

	// The user main email used in user's GenesysCloud account. Must provide either userId or userEmail, but not both.
	UserEmail string `json:"userEmail,omitempty"`

	// The user ID. Must provide either userId or userEmail, but not both.
	UserID string `json:"userId,omitempty"`

	// The value of the metric data. When value is null, the metric data will be deleted.
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this external metric data item
func (m *ExternalMetricDataItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateOccurred(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalMetricDataItem) validateDateOccurred(formats strfmt.Registry) error {

	if err := validate.Required("dateOccurred", "body", m.DateOccurred); err != nil {
		return err
	}

	if err := validate.FormatOf("dateOccurred", "body", "date", m.DateOccurred.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExternalMetricDataItem) validateMetricID(formats strfmt.Registry) error {

	if err := validate.Required("metricId", "body", m.MetricID); err != nil {
		return err
	}

	return nil
}

var externalMetricDataItemTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Total","Cumulative"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalMetricDataItemTypeTypePropEnum = append(externalMetricDataItemTypeTypePropEnum, v)
	}
}

const (

	// ExternalMetricDataItemTypeTotal captures enum value "Total"
	ExternalMetricDataItemTypeTotal string = "Total"

	// ExternalMetricDataItemTypeCumulative captures enum value "Cumulative"
	ExternalMetricDataItemTypeCumulative string = "Cumulative"
)

// prop value enum
func (m *ExternalMetricDataItem) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externalMetricDataItemTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternalMetricDataItem) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ExternalMetricDataItem) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this external metric data item based on context it is used
func (m *ExternalMetricDataItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalMetricDataItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalMetricDataItem) UnmarshalBinary(b []byte) error {
	var res ExternalMetricDataItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
