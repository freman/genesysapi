// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FlowExecutionLaunchRequest Parameters for launching a flow.
//
// swagger:model FlowExecutionLaunchRequest
type FlowExecutionLaunchRequest struct {

	// ID of the flow to launch.
	// Required: true
	FlowID *string `json:"flowId"`

	// The version of the flow to launch. Omit this value (or supply null/empty) to use the latest published version.
	FlowVersion string `json:"flowVersion,omitempty"`

	// Input values to the flow. Valid values are defined by a flow's input JSON schema.
	InputData map[string]interface{} `json:"inputData,omitempty"`

	// A displayable name to assign to the new flow execution
	Name string `json:"name,omitempty"`
}

// Validate validates this flow execution launch request
func (m *FlowExecutionLaunchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlowID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowExecutionLaunchRequest) validateFlowID(formats strfmt.Registry) error {

	if err := validate.Required("flowId", "body", m.FlowID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this flow execution launch request based on context it is used
func (m *FlowExecutionLaunchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlowExecutionLaunchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowExecutionLaunchRequest) UnmarshalBinary(b []byte) error {
	var res FlowExecutionLaunchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
