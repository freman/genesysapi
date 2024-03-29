// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgeMetricsProcessor edge metrics processor
//
// swagger:model EdgeMetricsProcessor
type EdgeMetricsProcessor struct {

	// Percent time processor was active.
	ActiveTimePct float64 `json:"activeTimePct,omitempty"`

	// Machine CPU identifier. 'total' will always be included in the array and is the total of all CPU resources.
	CPUID string `json:"cpuId,omitempty"`

	// Percent time processor was idle.
	IdleTimePct float64 `json:"idleTimePct,omitempty"`

	// Percent time processor spent in privileged mode.
	PrivilegedTimePct float64 `json:"privilegedTimePct,omitempty"`

	// Percent time processor spent in user mode.
	UserTimePct float64 `json:"userTimePct,omitempty"`
}

// Validate validates this edge metrics processor
func (m *EdgeMetricsProcessor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edge metrics processor based on context it is used
func (m *EdgeMetricsProcessor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgeMetricsProcessor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeMetricsProcessor) UnmarshalBinary(b []byte) error {
	var res EdgeMetricsProcessor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
