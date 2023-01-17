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

// ExternalMetricDefinitionCreateRequest external metric definition create request
//
// swagger:model ExternalMetricDefinitionCreateRequest
type ExternalMetricDefinitionCreateRequest struct {

	// The default objective type of the External Metric Definition
	// Required: true
	// Enum: [HigherIsBetter LowerIsBetter TargetArea]
	DefaultObjectiveType *string `json:"defaultObjectiveType"`

	// True if the External Metric Definition is enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// The name of the External Metric Definition
	// Required: true
	Name *string `json:"name"`

	// The decimal precision of the External Metric Definition. Must be at least 0 and at most 5
	// Required: true
	// Maximum: 5
	// Minimum: 0
	Precision *int32 `json:"precision"`

	// The unit of the External Metric Definition
	// Required: true
	// Enum: [Seconds Percent Number Currency]
	Unit *string `json:"unit"`

	// The unit definition of the External Metric Definition
	UnitDefinition string `json:"unitDefinition,omitempty"`
}

// Validate validates this external metric definition create request
func (m *ExternalMetricDefinitionCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultObjectiveType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrecision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var externalMetricDefinitionCreateRequestTypeDefaultObjectiveTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HigherIsBetter","LowerIsBetter","TargetArea"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalMetricDefinitionCreateRequestTypeDefaultObjectiveTypePropEnum = append(externalMetricDefinitionCreateRequestTypeDefaultObjectiveTypePropEnum, v)
	}
}

const (

	// ExternalMetricDefinitionCreateRequestDefaultObjectiveTypeHigherIsBetter captures enum value "HigherIsBetter"
	ExternalMetricDefinitionCreateRequestDefaultObjectiveTypeHigherIsBetter string = "HigherIsBetter"

	// ExternalMetricDefinitionCreateRequestDefaultObjectiveTypeLowerIsBetter captures enum value "LowerIsBetter"
	ExternalMetricDefinitionCreateRequestDefaultObjectiveTypeLowerIsBetter string = "LowerIsBetter"

	// ExternalMetricDefinitionCreateRequestDefaultObjectiveTypeTargetArea captures enum value "TargetArea"
	ExternalMetricDefinitionCreateRequestDefaultObjectiveTypeTargetArea string = "TargetArea"
)

// prop value enum
func (m *ExternalMetricDefinitionCreateRequest) validateDefaultObjectiveTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externalMetricDefinitionCreateRequestTypeDefaultObjectiveTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternalMetricDefinitionCreateRequest) validateDefaultObjectiveType(formats strfmt.Registry) error {

	if err := validate.Required("defaultObjectiveType", "body", m.DefaultObjectiveType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDefaultObjectiveTypeEnum("defaultObjectiveType", "body", *m.DefaultObjectiveType); err != nil {
		return err
	}

	return nil
}

func (m *ExternalMetricDefinitionCreateRequest) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ExternalMetricDefinitionCreateRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ExternalMetricDefinitionCreateRequest) validatePrecision(formats strfmt.Registry) error {

	if err := validate.Required("precision", "body", m.Precision); err != nil {
		return err
	}

	if err := validate.MinimumInt("precision", "body", int64(*m.Precision), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("precision", "body", int64(*m.Precision), 5, false); err != nil {
		return err
	}

	return nil
}

var externalMetricDefinitionCreateRequestTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Seconds","Percent","Number","Currency"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalMetricDefinitionCreateRequestTypeUnitPropEnum = append(externalMetricDefinitionCreateRequestTypeUnitPropEnum, v)
	}
}

const (

	// ExternalMetricDefinitionCreateRequestUnitSeconds captures enum value "Seconds"
	ExternalMetricDefinitionCreateRequestUnitSeconds string = "Seconds"

	// ExternalMetricDefinitionCreateRequestUnitPercent captures enum value "Percent"
	ExternalMetricDefinitionCreateRequestUnitPercent string = "Percent"

	// ExternalMetricDefinitionCreateRequestUnitNumber captures enum value "Number"
	ExternalMetricDefinitionCreateRequestUnitNumber string = "Number"

	// ExternalMetricDefinitionCreateRequestUnitCurrency captures enum value "Currency"
	ExternalMetricDefinitionCreateRequestUnitCurrency string = "Currency"
)

// prop value enum
func (m *ExternalMetricDefinitionCreateRequest) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externalMetricDefinitionCreateRequestTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternalMetricDefinitionCreateRequest) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this external metric definition create request based on context it is used
func (m *ExternalMetricDefinitionCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalMetricDefinitionCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalMetricDefinitionCreateRequest) UnmarshalBinary(b []byte) error {
	var res ExternalMetricDefinitionCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
