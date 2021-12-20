// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ComparisonPeriod comparison period
//
// swagger:model ComparisonPeriod
type ComparisonPeriod struct {

	// End date of the comparison period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateEnded strfmt.DateTime `json:"dateEnded,omitempty"`

	// Start date of the comparison period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateStarted strfmt.DateTime `json:"dateStarted,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Total interactions not routed by predictive routing (GPR was off)
	// Read Only: true
	InteractionCountOff int64 `json:"interactionCountOff,omitempty"`

	// Total interactions handled by predictive routing (GPR was on)
	// Read Only: true
	InteractionCountOn int64 `json:"interactionCountOn,omitempty"`

	// Key Performance Indicator optimised during the comparison period.
	// Read Only: true
	Kpi string `json:"kpi,omitempty"`

	// KPI results for each metric
	// Read Only: true
	KpiResults []*KpiResult `json:"kpiResults"`

	// Absolute metric (in which the KPI is based) total for the interactions not routed by predictive routing (GPR was off)
	// Read Only: true
	KpiTotalOff int64 `json:"kpiTotalOff,omitempty"`

	// Absolute metric (in which the KPI is based) total for the interactions handled by predictive routing (GPR was on)
	// Read Only: true
	KpiTotalOn int64 `json:"kpiTotalOn,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this comparison period
func (m *ComparisonPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateEnded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKpiResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComparisonPeriod) validateDateEnded(formats strfmt.Registry) error {

	if swag.IsZero(m.DateEnded) { // not required
		return nil
	}

	if err := validate.FormatOf("dateEnded", "body", "date-time", m.DateEnded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ComparisonPeriod) validateDateStarted(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStarted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStarted", "body", "date-time", m.DateStarted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ComparisonPeriod) validateKpiResults(formats strfmt.Registry) error {

	if swag.IsZero(m.KpiResults) { // not required
		return nil
	}

	for i := 0; i < len(m.KpiResults); i++ {
		if swag.IsZero(m.KpiResults[i]) { // not required
			continue
		}

		if m.KpiResults[i] != nil {
			if err := m.KpiResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kpiResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComparisonPeriod) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComparisonPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComparisonPeriod) UnmarshalBinary(b []byte) error {
	var res ComparisonPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
