// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObservationMetricData observation metric data
//
// swagger:model ObservationMetricData
type ObservationMetricData struct {

	// metric
	Metric string `json:"metric,omitempty"`

	// List of observations sorted by timestamp in ascending order. This list may be truncated.
	Observations []*ObservationValue `json:"observations"`

	// qualifier
	Qualifier string `json:"qualifier,omitempty"`

	// stats
	Stats *StatisticalSummary `json:"stats,omitempty"`

	// Flag for a truncated list of observations. If truncated, the first half of the list of observations will contain the oldest observations and the second half the newest observations.
	Truncated bool `json:"truncated"`
}

// Validate validates this observation metric data
func (m *ObservationMetricData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObservations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObservationMetricData) validateObservations(formats strfmt.Registry) error {
	if swag.IsZero(m.Observations) { // not required
		return nil
	}

	for i := 0; i < len(m.Observations); i++ {
		if swag.IsZero(m.Observations[i]) { // not required
			continue
		}

		if m.Observations[i] != nil {
			if err := m.Observations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("observations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("observations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObservationMetricData) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this observation metric data based on the context it is used
func (m *ObservationMetricData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObservations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObservationMetricData) contextValidateObservations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Observations); i++ {

		if m.Observations[i] != nil {
			if err := m.Observations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("observations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("observations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObservationMetricData) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {
		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObservationMetricData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObservationMetricData) UnmarshalBinary(b []byte) error {
	var res ObservationMetricData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
