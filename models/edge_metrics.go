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

// EdgeMetrics edge metrics
//
// swagger:model EdgeMetrics
type EdgeMetrics struct {

	// disks
	Disks []*EdgeMetricsDisk `json:"disks"`

	// edge
	Edge *DomainEntityRef `json:"edge,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EventTime strfmt.DateTime `json:"eventTime,omitempty"`

	// memory
	Memory []*EdgeMetricsMemory `json:"memory"`

	// networks
	Networks []*EdgeMetricsNetwork `json:"networks"`

	// processors
	Processors []*EdgeMetricsProcessor `json:"processors"`

	// subsystems
	Subsystems []*EdgeMetricsSubsystem `json:"subsystems"`

	// up time msec
	UpTimeMsec int64 `json:"upTimeMsec,omitempty"`
}

// Validate validates this edge metrics
func (m *EdgeMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeMetrics) validateDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeMetrics) validateEdge(formats strfmt.Registry) error {

	if swag.IsZero(m.Edge) { // not required
		return nil
	}

	if m.Edge != nil {
		if err := m.Edge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeMetrics) validateEventTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeMetrics) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	for i := 0; i < len(m.Memory); i++ {
		if swag.IsZero(m.Memory[i]) { // not required
			continue
		}

		if m.Memory[i] != nil {
			if err := m.Memory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeMetrics) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeMetrics) validateProcessors(formats strfmt.Registry) error {

	if swag.IsZero(m.Processors) { // not required
		return nil
	}

	for i := 0; i < len(m.Processors); i++ {
		if swag.IsZero(m.Processors[i]) { // not required
			continue
		}

		if m.Processors[i] != nil {
			if err := m.Processors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeMetrics) validateSubsystems(formats strfmt.Registry) error {

	if swag.IsZero(m.Subsystems) { // not required
		return nil
	}

	for i := 0; i < len(m.Subsystems); i++ {
		if swag.IsZero(m.Subsystems[i]) { // not required
			continue
		}

		if m.Subsystems[i] != nil {
			if err := m.Subsystems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subsystems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeMetrics) UnmarshalBinary(b []byte) error {
	var res EdgeMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
