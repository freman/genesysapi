// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CampaignRule campaign rule
//
// swagger:model CampaignRule
type CampaignRule struct {

	// The list of actions that are executed if the conditions are satisfied.
	// Required: true
	CampaignRuleActions []*CampaignRuleAction `json:"campaignRuleActions"`

	// The list of conditions that are evaluated on the entities.
	// Required: true
	CampaignRuleConditions []*CampaignRuleCondition `json:"campaignRuleConditions"`

	// The list of entities that this CampaignRule monitors.
	// Required: true
	CampaignRuleEntities *CampaignRuleEntities `json:"campaignRuleEntities"`

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Whether or not this CampaignRule is currently enabled. Required on updates.
	Enabled bool `json:"enabled"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// match any conditions
	MatchAnyConditions bool `json:"matchAnyConditions"`

	// The name of the CampaignRule.
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this campaign rule
func (m *CampaignRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaignRuleActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCampaignRuleConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCampaignRuleEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CampaignRule) validateCampaignRuleActions(formats strfmt.Registry) error {

	if err := validate.Required("campaignRuleActions", "body", m.CampaignRuleActions); err != nil {
		return err
	}

	for i := 0; i < len(m.CampaignRuleActions); i++ {
		if swag.IsZero(m.CampaignRuleActions[i]) { // not required
			continue
		}

		if m.CampaignRuleActions[i] != nil {
			if err := m.CampaignRuleActions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("campaignRuleActions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CampaignRule) validateCampaignRuleConditions(formats strfmt.Registry) error {

	if err := validate.Required("campaignRuleConditions", "body", m.CampaignRuleConditions); err != nil {
		return err
	}

	for i := 0; i < len(m.CampaignRuleConditions); i++ {
		if swag.IsZero(m.CampaignRuleConditions[i]) { // not required
			continue
		}

		if m.CampaignRuleConditions[i] != nil {
			if err := m.CampaignRuleConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("campaignRuleConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CampaignRule) validateCampaignRuleEntities(formats strfmt.Registry) error {

	if err := validate.Required("campaignRuleEntities", "body", m.CampaignRuleEntities); err != nil {
		return err
	}

	if m.CampaignRuleEntities != nil {
		if err := m.CampaignRuleEntities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("campaignRuleEntities")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignRule) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignRule) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CampaignRule) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CampaignRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CampaignRule) UnmarshalBinary(b []byte) error {
	var res CampaignRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}