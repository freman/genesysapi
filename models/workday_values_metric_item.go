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

// WorkdayValuesMetricItem workday values metric item
//
// swagger:model WorkdayValuesMetricItem
type WorkdayValuesMetricItem struct {

	// The average value of the metric
	// Read Only: true
	Average float64 `json:"average,omitempty"`

	// Gamification metric for the average and the trend
	// Read Only: true
	MetricDefinition *MetricDefinition `json:"metricDefinition,omitempty"`

	// The metric value trend
	// Read Only: true
	Trend []*WorkdayValuesTrendItem `json:"trend"`

	// The unit type of the metric value
	// Read Only: true
	// Enum: [None Percent Seconds Number AttendanceStatus Unit]
	UnitType string `json:"unitType,omitempty"`
}

// Validate validates this workday values metric item
func (m *WorkdayValuesMetricItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetricDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkdayValuesMetricItem) validateMetricDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.MetricDefinition) { // not required
		return nil
	}

	if m.MetricDefinition != nil {
		if err := m.MetricDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metricDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *WorkdayValuesMetricItem) validateTrend(formats strfmt.Registry) error {

	if swag.IsZero(m.Trend) { // not required
		return nil
	}

	for i := 0; i < len(m.Trend); i++ {
		if swag.IsZero(m.Trend[i]) { // not required
			continue
		}

		if m.Trend[i] != nil {
			if err := m.Trend[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trend" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var workdayValuesMetricItemTypeUnitTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Percent","Seconds","Number","AttendanceStatus","Unit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workdayValuesMetricItemTypeUnitTypePropEnum = append(workdayValuesMetricItemTypeUnitTypePropEnum, v)
	}
}

const (

	// WorkdayValuesMetricItemUnitTypeNone captures enum value "None"
	WorkdayValuesMetricItemUnitTypeNone string = "None"

	// WorkdayValuesMetricItemUnitTypePercent captures enum value "Percent"
	WorkdayValuesMetricItemUnitTypePercent string = "Percent"

	// WorkdayValuesMetricItemUnitTypeSeconds captures enum value "Seconds"
	WorkdayValuesMetricItemUnitTypeSeconds string = "Seconds"

	// WorkdayValuesMetricItemUnitTypeNumber captures enum value "Number"
	WorkdayValuesMetricItemUnitTypeNumber string = "Number"

	// WorkdayValuesMetricItemUnitTypeAttendanceStatus captures enum value "AttendanceStatus"
	WorkdayValuesMetricItemUnitTypeAttendanceStatus string = "AttendanceStatus"

	// WorkdayValuesMetricItemUnitTypeUnit captures enum value "Unit"
	WorkdayValuesMetricItemUnitTypeUnit string = "Unit"
)

// prop value enum
func (m *WorkdayValuesMetricItem) validateUnitTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workdayValuesMetricItemTypeUnitTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WorkdayValuesMetricItem) validateUnitType(formats strfmt.Registry) error {

	if swag.IsZero(m.UnitType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitTypeEnum("unitType", "body", m.UnitType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkdayValuesMetricItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkdayValuesMetricItem) UnmarshalBinary(b []byte) error {
	var res WorkdayValuesMetricItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
