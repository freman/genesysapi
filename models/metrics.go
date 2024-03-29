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

// Metrics metrics
//
// swagger:model Metrics
type Metrics struct {

	// The created date of this metric. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The unlinked workday for this metric if this metric was ever unlinked. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Read Only: true
	// Format: date
	DateUnlinked strfmt.Date `json:"dateUnlinked,omitempty"`

	// A flag for whether this metric is enabled for a performance profile
	Enabled bool `json:"enabled"`

	// The id of associated external metric definition
	ExternalMetricDefinitionID string `json:"externalMetricDefinitionId,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The linked metric entity reference
	// Read Only: true
	LinkedMetric *AddressableEntityRef `json:"linkedMetric,omitempty"`

	// Achievable maximum points for this metric
	MaxPoints int32 `json:"maxPoints,omitempty"`

	// The id of associated metric definition
	MetricDefinitionID string `json:"metricDefinitionId,omitempty"`

	// The name of associated metric definition
	MetricDefinitionName string `json:"metricDefinitionName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The order of metric within a performance profile
	Order int32 `json:"order,omitempty"`

	// Performance profile id of this metric
	PerformanceProfileID string `json:"performanceProfileId,omitempty"`

	// Precision of linked external metric
	// Read Only: true
	Precision int32 `json:"precision,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The source performance profile when this metric is linked
	// Read Only: true
	SourcePerformanceProfile *PerformanceProfile `json:"sourcePerformanceProfile,omitempty"`

	// The name of associated objective template
	TemplateName string `json:"templateName,omitempty"`

	// Unit definition of linked external metric
	// Read Only: true
	UnitDefinition string `json:"unitDefinition,omitempty"`

	// Corresponding unit type for this metric
	// Enum: [None Percent Currency Seconds Number AttendanceStatus Unit]
	UnitType string `json:"unitType,omitempty"`
}

// Validate validates this metrics
func (m *Metrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateUnlinked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePerformanceProfile(formats); err != nil {
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

func (m *Metrics) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) validateDateUnlinked(formats strfmt.Registry) error {
	if swag.IsZero(m.DateUnlinked) { // not required
		return nil
	}

	if err := validate.FormatOf("dateUnlinked", "body", "date", m.DateUnlinked.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) validateLinkedMetric(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkedMetric) { // not required
		return nil
	}

	if m.LinkedMetric != nil {
		if err := m.LinkedMetric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linkedMetric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("linkedMetric")
			}
			return err
		}
	}

	return nil
}

func (m *Metrics) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) validateSourcePerformanceProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.SourcePerformanceProfile) { // not required
		return nil
	}

	if m.SourcePerformanceProfile != nil {
		if err := m.SourcePerformanceProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcePerformanceProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourcePerformanceProfile")
			}
			return err
		}
	}

	return nil
}

var metricsTypeUnitTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Percent","Currency","Seconds","Number","AttendanceStatus","Unit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricsTypeUnitTypePropEnum = append(metricsTypeUnitTypePropEnum, v)
	}
}

const (

	// MetricsUnitTypeNone captures enum value "None"
	MetricsUnitTypeNone string = "None"

	// MetricsUnitTypePercent captures enum value "Percent"
	MetricsUnitTypePercent string = "Percent"

	// MetricsUnitTypeCurrency captures enum value "Currency"
	MetricsUnitTypeCurrency string = "Currency"

	// MetricsUnitTypeSeconds captures enum value "Seconds"
	MetricsUnitTypeSeconds string = "Seconds"

	// MetricsUnitTypeNumber captures enum value "Number"
	MetricsUnitTypeNumber string = "Number"

	// MetricsUnitTypeAttendanceStatus captures enum value "AttendanceStatus"
	MetricsUnitTypeAttendanceStatus string = "AttendanceStatus"

	// MetricsUnitTypeUnit captures enum value "Unit"
	MetricsUnitTypeUnit string = "Unit"
)

// prop value enum
func (m *Metrics) validateUnitTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metricsTypeUnitTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Metrics) validateUnitType(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUnitTypeEnum("unitType", "body", m.UnitType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrics based on the context it is used
func (m *Metrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateUnlinked(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkedMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrecision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourcePerformanceProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnitDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Metrics) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) contextValidateDateUnlinked(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateUnlinked", "body", strfmt.Date(m.DateUnlinked)); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) contextValidateLinkedMetric(ctx context.Context, formats strfmt.Registry) error {

	if m.LinkedMetric != nil {
		if err := m.LinkedMetric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linkedMetric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("linkedMetric")
			}
			return err
		}
	}

	return nil
}

func (m *Metrics) contextValidatePrecision(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "precision", "body", int32(m.Precision)); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) contextValidateSourcePerformanceProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.SourcePerformanceProfile != nil {
		if err := m.SourcePerformanceProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcePerformanceProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourcePerformanceProfile")
			}
			return err
		}
	}

	return nil
}

func (m *Metrics) contextValidateUnitDefinition(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "unitDefinition", "body", string(m.UnitDefinition)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Metrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metrics) UnmarshalBinary(b []byte) error {
	var res Metrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
