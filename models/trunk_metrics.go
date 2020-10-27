// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrunkMetrics trunk metrics
//
// swagger:model TrunkMetrics
type TrunkMetrics struct {

	// calls
	Calls *TrunkMetricsCalls `json:"calls,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EventTime strfmt.DateTime `json:"eventTime,omitempty"`

	// logical interface
	LogicalInterface *DomainEntityRef `json:"logicalInterface,omitempty"`

	// qos
	Qos *TrunkMetricsQoS `json:"qos,omitempty"`

	// trunk
	Trunk *DomainEntityRef `json:"trunk,omitempty"`
}

// Validate validates this trunk metrics
func (m *TrunkMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicalInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrunk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrunkMetrics) validateCalls(formats strfmt.Registry) error {

	if swag.IsZero(m.Calls) { // not required
		return nil
	}

	if m.Calls != nil {
		if err := m.Calls.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calls")
			}
			return err
		}
	}

	return nil
}

func (m *TrunkMetrics) validateEventTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrunkMetrics) validateLogicalInterface(formats strfmt.Registry) error {

	if swag.IsZero(m.LogicalInterface) { // not required
		return nil
	}

	if m.LogicalInterface != nil {
		if err := m.LogicalInterface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logicalInterface")
			}
			return err
		}
	}

	return nil
}

func (m *TrunkMetrics) validateQos(formats strfmt.Registry) error {

	if swag.IsZero(m.Qos) { // not required
		return nil
	}

	if m.Qos != nil {
		if err := m.Qos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos")
			}
			return err
		}
	}

	return nil
}

func (m *TrunkMetrics) validateTrunk(formats strfmt.Registry) error {

	if swag.IsZero(m.Trunk) { // not required
		return nil
	}

	if m.Trunk != nil {
		if err := m.Trunk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrunkMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrunkMetrics) UnmarshalBinary(b []byte) error {
	var res TrunkMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
