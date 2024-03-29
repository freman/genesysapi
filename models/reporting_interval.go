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

// ReportingInterval reporting interval
//
// swagger:model ReportingInterval
type ReportingInterval struct {

	// The granularity of the reporting interval period
	// Required: true
	// Enum: [Day Week Month]
	IntervalType *string `json:"intervalType"`

	// The value of the reporting interval period for a given interval type
	// Required: true
	IntervalValue *int32 `json:"intervalValue"`
}

// Validate validates this reporting interval
func (m *ReportingInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntervalType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reportingIntervalTypeIntervalTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Day","Week","Month"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingIntervalTypeIntervalTypePropEnum = append(reportingIntervalTypeIntervalTypePropEnum, v)
	}
}

const (

	// ReportingIntervalIntervalTypeDay captures enum value "Day"
	ReportingIntervalIntervalTypeDay string = "Day"

	// ReportingIntervalIntervalTypeWeek captures enum value "Week"
	ReportingIntervalIntervalTypeWeek string = "Week"

	// ReportingIntervalIntervalTypeMonth captures enum value "Month"
	ReportingIntervalIntervalTypeMonth string = "Month"
)

// prop value enum
func (m *ReportingInterval) validateIntervalTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingIntervalTypeIntervalTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingInterval) validateIntervalType(formats strfmt.Registry) error {

	if err := validate.Required("intervalType", "body", m.IntervalType); err != nil {
		return err
	}

	// value enum
	if err := m.validateIntervalTypeEnum("intervalType", "body", *m.IntervalType); err != nil {
		return err
	}

	return nil
}

func (m *ReportingInterval) validateIntervalValue(formats strfmt.Registry) error {

	if err := validate.Required("intervalValue", "body", m.IntervalValue); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this reporting interval based on context it is used
func (m *ReportingInterval) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportingInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingInterval) UnmarshalBinary(b []byte) error {
	var res ReportingInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
