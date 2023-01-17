// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchPredictorRequest patch predictor request
//
// swagger:model PatchPredictorRequest
type PatchPredictorRequest struct {

	// Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
	RoutingTimeoutSeconds int32 `json:"routingTimeoutSeconds,omitempty"`

	// The predictor schedule that determines when the predictor is used for routing interactions.
	Schedule *PredictorSchedule `json:"schedule,omitempty"`

	// The predictor balancing configuration to enable workload balancing
	WorkloadBalancingConfig *PredictorWorkloadBalancing `json:"workloadBalancingConfig,omitempty"`
}

// Validate validates this patch predictor request
func (m *PatchPredictorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadBalancingConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchPredictorRequest) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *PatchPredictorRequest) validateWorkloadBalancingConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadBalancingConfig) { // not required
		return nil
	}

	if m.WorkloadBalancingConfig != nil {
		if err := m.WorkloadBalancingConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadBalancingConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadBalancingConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch predictor request based on the context it is used
func (m *PatchPredictorRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkloadBalancingConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchPredictorRequest) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *PatchPredictorRequest) contextValidateWorkloadBalancingConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkloadBalancingConfig != nil {
		if err := m.WorkloadBalancingConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadBalancingConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadBalancingConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchPredictorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchPredictorRequest) UnmarshalBinary(b []byte) error {
	var res PatchPredictorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
