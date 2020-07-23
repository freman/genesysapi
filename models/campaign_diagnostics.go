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

// CampaignDiagnostics campaign diagnostics
//
// swagger:model CampaignDiagnostics
type CampaignDiagnostics struct {

	// Campaign properties that can impact which contacts are callable
	// Read Only: true
	CallableContacts *CallableContactsDiagnostic `json:"callableContacts,omitempty"`

	// Current number of outstanding interactions on the campaign
	// Read Only: true
	OutstandingInteractionsCount int32 `json:"outstandingInteractionsCount,omitempty"`

	// Information regarding the campaign's queue
	// Read Only: true
	QueueUtilizationDiagnostic *QueueUtilizationDiagnostic `json:"queueUtilizationDiagnostic,omitempty"`

	// Information regarding the campaign's rule sets
	// Read Only: true
	RuleSetDiagnostics []*RuleSetDiagnostic `json:"ruleSetDiagnostics"`

	// Current number of scheduled interactions on the campaign
	// Read Only: true
	ScheduledInteractionsCount int32 `json:"scheduledInteractionsCount,omitempty"`
}

// Validate validates this campaign diagnostics
func (m *CampaignDiagnostics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallableContacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueUtilizationDiagnostic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleSetDiagnostics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CampaignDiagnostics) validateCallableContacts(formats strfmt.Registry) error {

	if swag.IsZero(m.CallableContacts) { // not required
		return nil
	}

	if m.CallableContacts != nil {
		if err := m.CallableContacts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callableContacts")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignDiagnostics) validateQueueUtilizationDiagnostic(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueUtilizationDiagnostic) { // not required
		return nil
	}

	if m.QueueUtilizationDiagnostic != nil {
		if err := m.QueueUtilizationDiagnostic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queueUtilizationDiagnostic")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignDiagnostics) validateRuleSetDiagnostics(formats strfmt.Registry) error {

	if swag.IsZero(m.RuleSetDiagnostics) { // not required
		return nil
	}

	for i := 0; i < len(m.RuleSetDiagnostics); i++ {
		if swag.IsZero(m.RuleSetDiagnostics[i]) { // not required
			continue
		}

		if m.RuleSetDiagnostics[i] != nil {
			if err := m.RuleSetDiagnostics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ruleSetDiagnostics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CampaignDiagnostics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CampaignDiagnostics) UnmarshalBinary(b []byte) error {
	var res CampaignDiagnostics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
