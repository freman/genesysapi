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

// ExternalMetricDefinitionUpdateRequest external metric definition update request
//
// swagger:model ExternalMetricDefinitionUpdateRequest
type ExternalMetricDefinitionUpdateRequest struct {

	// The default objective type of the External Metric Definition
	// Enum: [HigherIsBetter LowerIsBetter TargetArea]
	DefaultObjectiveType string `json:"defaultObjectiveType,omitempty"`

	// True if the External Metric Definition is enabled
	Enabled bool `json:"enabled"`

	// The name of the External Metric Definition
	Name string `json:"name,omitempty"`

	// The decimal precision of the External Metric Definition. Must be at least 0 and at most 5
	// Maximum: 5
	// Minimum: 0
	Precision *int32 `json:"precision,omitempty"`
}

// Validate validates this external metric definition update request
func (m *ExternalMetricDefinitionUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultObjectiveType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrecision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var externalMetricDefinitionUpdateRequestTypeDefaultObjectiveTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HigherIsBetter","LowerIsBetter","TargetArea"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalMetricDefinitionUpdateRequestTypeDefaultObjectiveTypePropEnum = append(externalMetricDefinitionUpdateRequestTypeDefaultObjectiveTypePropEnum, v)
	}
}

const (

	// ExternalMetricDefinitionUpdateRequestDefaultObjectiveTypeHigherIsBetter captures enum value "HigherIsBetter"
	ExternalMetricDefinitionUpdateRequestDefaultObjectiveTypeHigherIsBetter string = "HigherIsBetter"

	// ExternalMetricDefinitionUpdateRequestDefaultObjectiveTypeLowerIsBetter captures enum value "LowerIsBetter"
	ExternalMetricDefinitionUpdateRequestDefaultObjectiveTypeLowerIsBetter string = "LowerIsBetter"

	// ExternalMetricDefinitionUpdateRequestDefaultObjectiveTypeTargetArea captures enum value "TargetArea"
	ExternalMetricDefinitionUpdateRequestDefaultObjectiveTypeTargetArea string = "TargetArea"
)

// prop value enum
func (m *ExternalMetricDefinitionUpdateRequest) validateDefaultObjectiveTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externalMetricDefinitionUpdateRequestTypeDefaultObjectiveTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternalMetricDefinitionUpdateRequest) validateDefaultObjectiveType(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultObjectiveType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDefaultObjectiveTypeEnum("defaultObjectiveType", "body", m.DefaultObjectiveType); err != nil {
		return err
	}

	return nil
}

func (m *ExternalMetricDefinitionUpdateRequest) validatePrecision(formats strfmt.Registry) error {

	if swag.IsZero(m.Precision) { // not required
		return nil
	}

	if err := validate.MinimumInt("precision", "body", int64(*m.Precision), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("precision", "body", int64(*m.Precision), 5, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalMetricDefinitionUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalMetricDefinitionUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ExternalMetricDefinitionUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}