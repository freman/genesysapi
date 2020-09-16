// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnalyticsProposedAgent analytics proposed agent
//
// swagger:model AnalyticsProposedAgent
type AnalyticsProposedAgent struct {

	// Proposed agent rank for this conversation from predictive routing (lower is better)
	AgentRank int32 `json:"agentRank,omitempty"`

	// Unique identifier of an agent that was proposed by predictive routing
	ProposedAgentID string `json:"proposedAgentId,omitempty"`
}

// Validate validates this analytics proposed agent
func (m *AnalyticsProposedAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsProposedAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsProposedAgent) UnmarshalBinary(b []byte) error {
	var res AnalyticsProposedAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
