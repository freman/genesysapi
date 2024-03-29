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
	"github.com/go-openapi/validate"
)

// RoutingData routing data
//
// swagger:model RoutingData
type RoutingData struct {

	// The identifier of a language to be considered in routing
	LanguageID string `json:"languageId,omitempty"`

	// A list of agents to be preferred in routing
	PreferredAgentIds []string `json:"preferredAgentIds"`

	// The priority for routing
	Priority int32 `json:"priority,omitempty"`

	// The identifier of the routing queue
	// Required: true
	QueueID *string `json:"queueId"`

	// An array of flags indicating how the conversation should be routed. Use "AGENT_OWNED_CALLBACK" when creating an Agent Owned Callback.
	RoutingFlags []string `json:"routingFlags"`

	// A list of scored agents for routing decisions. For Agent Owned Callbacks use one scored agent with a score of 100.
	ScoredAgents []*ScoredAgent `json:"scoredAgents"`

	// A list of skill identifiers to be considered in routing
	SkillIds []string `json:"skillIds"`
}

// Validate validates this routing data
func (m *RoutingData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueueID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoredAgents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutingData) validateQueueID(formats strfmt.Registry) error {

	if err := validate.Required("queueId", "body", m.QueueID); err != nil {
		return err
	}

	return nil
}

func (m *RoutingData) validateScoredAgents(formats strfmt.Registry) error {
	if swag.IsZero(m.ScoredAgents) { // not required
		return nil
	}

	for i := 0; i < len(m.ScoredAgents); i++ {
		if swag.IsZero(m.ScoredAgents[i]) { // not required
			continue
		}

		if m.ScoredAgents[i] != nil {
			if err := m.ScoredAgents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this routing data based on the context it is used
func (m *RoutingData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScoredAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutingData) contextValidateScoredAgents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ScoredAgents); i++ {

		if m.ScoredAgents[i] != nil {
			if err := m.ScoredAgents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoutingData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutingData) UnmarshalBinary(b []byte) error {
	var res RoutingData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
