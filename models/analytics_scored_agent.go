// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnalyticsScoredAgent analytics scored agent
//
// swagger:model AnalyticsScoredAgent
type AnalyticsScoredAgent struct {

	// Assigned agent score for this conversation (0 - 100, higher being better)
	AgentScore int32 `json:"agentScore,omitempty"`

	// Unique identifier for the agent that was scored for this conversation
	ScoredAgentID string `json:"scoredAgentId,omitempty"`
}

// Validate validates this analytics scored agent
func (m *AnalyticsScoredAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this analytics scored agent based on context it is used
func (m *AnalyticsScoredAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsScoredAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsScoredAgent) UnmarshalBinary(b []byte) error {
	var res AnalyticsScoredAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
