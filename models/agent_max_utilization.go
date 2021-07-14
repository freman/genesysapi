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

// AgentMaxUtilization agent max utilization
//
// swagger:model AgentMaxUtilization
type AgentMaxUtilization struct {

	// level
	// Enum: [Agent Organization]
	Level string `json:"level,omitempty"`

	// Map of media type to utilization settings.  Valid media types include call, callback, chat, email, and message.
	Utilization map[string]MediaUtilization `json:"utilization,omitempty"`
}

// Validate validates this agent max utilization
func (m *AgentMaxUtilization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUtilization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var agentMaxUtilizationTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Agent","Organization"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		agentMaxUtilizationTypeLevelPropEnum = append(agentMaxUtilizationTypeLevelPropEnum, v)
	}
}

const (

	// AgentMaxUtilizationLevelAgent captures enum value "Agent"
	AgentMaxUtilizationLevelAgent string = "Agent"

	// AgentMaxUtilizationLevelOrganization captures enum value "Organization"
	AgentMaxUtilizationLevelOrganization string = "Organization"
)

// prop value enum
func (m *AgentMaxUtilization) validateLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, agentMaxUtilizationTypeLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AgentMaxUtilization) validateLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *AgentMaxUtilization) validateUtilization(formats strfmt.Registry) error {

	if swag.IsZero(m.Utilization) { // not required
		return nil
	}

	for k := range m.Utilization {

		if err := validate.Required("utilization"+"."+k, "body", m.Utilization[k]); err != nil {
			return err
		}
		if val, ok := m.Utilization[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentMaxUtilization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentMaxUtilization) UnmarshalBinary(b []byte) error {
	var res AgentMaxUtilization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}