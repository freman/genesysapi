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

// CampaignRuleAction campaign rule action
//
// swagger:model CampaignRuleAction
type CampaignRuleAction struct {

	// The action to take on the campaignRuleActionEntities.
	// Required: true
	// Enum: [turnOnCampaign turnOffCampaign turnOnSequence turnOffSequence setCampaignPriority recycleCampaign setCampaignDialingMode]
	ActionType *string `json:"actionType"`

	// The list of entities that this action will apply to.
	// Required: true
	CampaignRuleActionEntities *CampaignRuleActionEntities `json:"campaignRuleActionEntities"`

	// id
	ID string `json:"id,omitempty"`

	// The parameters for the CampaignRuleAction. Required for certain actionTypes.
	Parameters *CampaignRuleParameters `json:"parameters,omitempty"`
}

// Validate validates this campaign rule action
func (m *CampaignRuleAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCampaignRuleActionEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var campaignRuleActionTypeActionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["turnOnCampaign","turnOffCampaign","turnOnSequence","turnOffSequence","setCampaignPriority","recycleCampaign","setCampaignDialingMode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		campaignRuleActionTypeActionTypePropEnum = append(campaignRuleActionTypeActionTypePropEnum, v)
	}
}

const (

	// CampaignRuleActionActionTypeTurnOnCampaign captures enum value "turnOnCampaign"
	CampaignRuleActionActionTypeTurnOnCampaign string = "turnOnCampaign"

	// CampaignRuleActionActionTypeTurnOffCampaign captures enum value "turnOffCampaign"
	CampaignRuleActionActionTypeTurnOffCampaign string = "turnOffCampaign"

	// CampaignRuleActionActionTypeTurnOnSequence captures enum value "turnOnSequence"
	CampaignRuleActionActionTypeTurnOnSequence string = "turnOnSequence"

	// CampaignRuleActionActionTypeTurnOffSequence captures enum value "turnOffSequence"
	CampaignRuleActionActionTypeTurnOffSequence string = "turnOffSequence"

	// CampaignRuleActionActionTypeSetCampaignPriority captures enum value "setCampaignPriority"
	CampaignRuleActionActionTypeSetCampaignPriority string = "setCampaignPriority"

	// CampaignRuleActionActionTypeRecycleCampaign captures enum value "recycleCampaign"
	CampaignRuleActionActionTypeRecycleCampaign string = "recycleCampaign"

	// CampaignRuleActionActionTypeSetCampaignDialingMode captures enum value "setCampaignDialingMode"
	CampaignRuleActionActionTypeSetCampaignDialingMode string = "setCampaignDialingMode"
)

// prop value enum
func (m *CampaignRuleAction) validateActionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, campaignRuleActionTypeActionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CampaignRuleAction) validateActionType(formats strfmt.Registry) error {

	if err := validate.Required("actionType", "body", m.ActionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionTypeEnum("actionType", "body", *m.ActionType); err != nil {
		return err
	}

	return nil
}

func (m *CampaignRuleAction) validateCampaignRuleActionEntities(formats strfmt.Registry) error {

	if err := validate.Required("campaignRuleActionEntities", "body", m.CampaignRuleActionEntities); err != nil {
		return err
	}

	if m.CampaignRuleActionEntities != nil {
		if err := m.CampaignRuleActionEntities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("campaignRuleActionEntities")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignRuleAction) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CampaignRuleAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CampaignRuleAction) UnmarshalBinary(b []byte) error {
	var res CampaignRuleAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}