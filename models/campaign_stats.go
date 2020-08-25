// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CampaignStats campaign stats
//
// swagger:model CampaignStats
type CampaignStats struct {

	// Calls per agent adjusted by pace
	// Read Only: true
	AdjustedCallsPerAgent float64 `json:"adjustedCallsPerAgent,omitempty"`

	// Information regarding the campaign's connect rate
	// Read Only: true
	ContactRate *ConnectRate `json:"contactRate,omitempty"`

	// Number of effective available agents not currently being utilized
	// Read Only: true
	EffectiveIdleAgents float64 `json:"effectiveIdleAgents,omitempty"`

	// Number of available agents not currently being utilized
	// Read Only: true
	IdleAgents int32 `json:"idleAgents,omitempty"`

	// Number of campaign calls currently ongoing
	// Read Only: true
	OutstandingCalls int32 `json:"outstandingCalls,omitempty"`

	// Number of campaign calls currently scheduled
	// Read Only: true
	ScheduledCalls int32 `json:"scheduledCalls,omitempty"`
}

// Validate validates this campaign stats
func (m *CampaignStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CampaignStats) validateContactRate(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactRate) { // not required
		return nil
	}

	if m.ContactRate != nil {
		if err := m.ContactRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactRate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CampaignStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CampaignStats) UnmarshalBinary(b []byte) error {
	var res CampaignStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}