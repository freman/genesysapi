// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RoutingConversationAttributesResponse routing conversation attributes response
//
// swagger:model RoutingConversationAttributesResponse
type RoutingConversationAttributesResponse struct {

	// Current language on in-queue conversation
	Language *Language `json:"language,omitempty"`

	// Current priority value on in-queue conversation. Range:[-25000000, 25000000]
	Priority int32 `json:"priority,omitempty"`

	// Current scored agents on in-queue conversation
	ScoredAgents []*ScoredAgent `json:"scoredAgents"`

	// Current routing skills on in-queue conversation
	Skills []*RoutingSkill `json:"skills"`
}

// Validate validates this routing conversation attributes response
func (m *RoutingConversationAttributesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoredAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkills(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutingConversationAttributesResponse) validateLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.Language) { // not required
		return nil
	}

	if m.Language != nil {
		if err := m.Language.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language")
			}
			return err
		}
	}

	return nil
}

func (m *RoutingConversationAttributesResponse) validateScoredAgents(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *RoutingConversationAttributesResponse) validateSkills(formats strfmt.Registry) error {

	if swag.IsZero(m.Skills) { // not required
		return nil
	}

	for i := 0; i < len(m.Skills); i++ {
		if swag.IsZero(m.Skills[i]) { // not required
			continue
		}

		if m.Skills[i] != nil {
			if err := m.Skills[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoutingConversationAttributesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutingConversationAttributesResponse) UnmarshalBinary(b []byte) error {
	var res RoutingConversationAttributesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
