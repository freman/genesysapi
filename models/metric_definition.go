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

// MetricDefinition metric definition
//
// swagger:model MetricDefinition
type MetricDefinition struct {

	// A predefined default objective for this metric
	DefaultObjective *DefaultObjective `json:"defaultObjective,omitempty"`

	// Metric names used as dividend
	DividendMetrics []string `json:"dividendMetrics"`

	// Metric names used as divisor
	DivisorMetrics []string `json:"divisorMetrics"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// An optional field to specify if this metric definition is locked to certain template. e.g. punctuality
	LockTemplateID string `json:"lockTemplateId,omitempty"`

	// Flag to indicate if this metricDefinition allows filter based on media types
	MediaTypeFilteringAllowed bool `json:"mediaTypeFilteringAllowed"`

	// name
	Name string `json:"name,omitempty"`

	// Flag to indicate if this metricDefinition allows filter based on queues
	QueueFilteringAllowed bool `json:"queueFilteringAllowed"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// An alternate name for this metric definition, often abbreviation
	ShortName string `json:"shortName,omitempty"`

	// The type of associated metric unit
	// Enum: [None Percent Currency Seconds Number AttendanceStatus Unit]
	UnitType string `json:"unitType,omitempty"`
}

// Validate validates this metric definition
func (m *MetricDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultObjective(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *MetricDefinition) validateDefaultObjective(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultObjective) { // not required
		return nil
	}

	if m.DefaultObjective != nil {
		if err := m.DefaultObjective.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultObjective")
			}
			return err
		}
	}

	return nil
}

func (m *MetricDefinition) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var metricDefinitionTypeUnitTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Percent","Currency","Seconds","Number","AttendanceStatus","Unit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricDefinitionTypeUnitTypePropEnum = append(metricDefinitionTypeUnitTypePropEnum, v)
	}
}

const (

	// MetricDefinitionUnitTypeNone captures enum value "None"
	MetricDefinitionUnitTypeNone string = "None"

	// MetricDefinitionUnitTypePercent captures enum value "Percent"
	MetricDefinitionUnitTypePercent string = "Percent"

	// MetricDefinitionUnitTypeCurrency captures enum value "Currency"
	MetricDefinitionUnitTypeCurrency string = "Currency"

	// MetricDefinitionUnitTypeSeconds captures enum value "Seconds"
	MetricDefinitionUnitTypeSeconds string = "Seconds"

	// MetricDefinitionUnitTypeNumber captures enum value "Number"
	MetricDefinitionUnitTypeNumber string = "Number"

	// MetricDefinitionUnitTypeAttendanceStatus captures enum value "AttendanceStatus"
	MetricDefinitionUnitTypeAttendanceStatus string = "AttendanceStatus"

	// MetricDefinitionUnitTypeUnit captures enum value "Unit"
	MetricDefinitionUnitTypeUnit string = "Unit"
)

// prop value enum
func (m *MetricDefinition) validateUnitTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metricDefinitionTypeUnitTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetricDefinition) validateUnitType(formats strfmt.Registry) error {

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
func (m *MetricDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricDefinition) UnmarshalBinary(b []byte) error {
	var res MetricDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
