// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CampaignRuleParameters campaign rule parameters
//
// swagger:model CampaignRuleParameters
type CampaignRuleParameters struct {

	// The dialing mode to set a campaign to. Required for the 'setCampaignDialingMode' action.
	// Enum: [agentless preview power predictive progressive external]
	DialingMode string `json:"dialingMode,omitempty"`

	// The operator for comparison. Required for a CampaignRuleCondition.
	// Enum: [equals greaterThan greaterThanEqualTo lessThan lessThanEqualTo]
	Operator string `json:"operator,omitempty"`

	// The priority to set a campaign to. Required for the 'setCampaignPriority' action.
	// Enum: [1 2 3 4 5]
	Priority string `json:"priority,omitempty"`

	// The value for comparison. Required for a CampaignRuleCondition.
	Value string `json:"value,omitempty"`
}

// Validate validates this campaign rule parameters
func (m *CampaignRuleParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDialingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var campaignRuleParametersTypeDialingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["agentless","preview","power","predictive","progressive","external"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		campaignRuleParametersTypeDialingModePropEnum = append(campaignRuleParametersTypeDialingModePropEnum, v)
	}
}

const (

	// CampaignRuleParametersDialingModeAgentless captures enum value "agentless"
	CampaignRuleParametersDialingModeAgentless string = "agentless"

	// CampaignRuleParametersDialingModePreview captures enum value "preview"
	CampaignRuleParametersDialingModePreview string = "preview"

	// CampaignRuleParametersDialingModePower captures enum value "power"
	CampaignRuleParametersDialingModePower string = "power"

	// CampaignRuleParametersDialingModePredictive captures enum value "predictive"
	CampaignRuleParametersDialingModePredictive string = "predictive"

	// CampaignRuleParametersDialingModeProgressive captures enum value "progressive"
	CampaignRuleParametersDialingModeProgressive string = "progressive"

	// CampaignRuleParametersDialingModeExternal captures enum value "external"
	CampaignRuleParametersDialingModeExternal string = "external"
)

// prop value enum
func (m *CampaignRuleParameters) validateDialingModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, campaignRuleParametersTypeDialingModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CampaignRuleParameters) validateDialingMode(formats strfmt.Registry) error {
	if swag.IsZero(m.DialingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateDialingModeEnum("dialingMode", "body", m.DialingMode); err != nil {
		return err
	}

	return nil
}

var campaignRuleParametersTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["equals","greaterThan","greaterThanEqualTo","lessThan","lessThanEqualTo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		campaignRuleParametersTypeOperatorPropEnum = append(campaignRuleParametersTypeOperatorPropEnum, v)
	}
}

const (

	// CampaignRuleParametersOperatorEquals captures enum value "equals"
	CampaignRuleParametersOperatorEquals string = "equals"

	// CampaignRuleParametersOperatorGreaterThan captures enum value "greaterThan"
	CampaignRuleParametersOperatorGreaterThan string = "greaterThan"

	// CampaignRuleParametersOperatorGreaterThanEqualTo captures enum value "greaterThanEqualTo"
	CampaignRuleParametersOperatorGreaterThanEqualTo string = "greaterThanEqualTo"

	// CampaignRuleParametersOperatorLessThan captures enum value "lessThan"
	CampaignRuleParametersOperatorLessThan string = "lessThan"

	// CampaignRuleParametersOperatorLessThanEqualTo captures enum value "lessThanEqualTo"
	CampaignRuleParametersOperatorLessThanEqualTo string = "lessThanEqualTo"
)

// prop value enum
func (m *CampaignRuleParameters) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, campaignRuleParametersTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CampaignRuleParameters) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var campaignRuleParametersTypePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","2","3","4","5"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		campaignRuleParametersTypePriorityPropEnum = append(campaignRuleParametersTypePriorityPropEnum, v)
	}
}

const (

	// CampaignRuleParametersPriorityNr1 captures enum value "1"
	CampaignRuleParametersPriorityNr1 string = "1"

	// CampaignRuleParametersPriorityNr2 captures enum value "2"
	CampaignRuleParametersPriorityNr2 string = "2"

	// CampaignRuleParametersPriorityNr3 captures enum value "3"
	CampaignRuleParametersPriorityNr3 string = "3"

	// CampaignRuleParametersPriorityNr4 captures enum value "4"
	CampaignRuleParametersPriorityNr4 string = "4"

	// CampaignRuleParametersPriorityNr5 captures enum value "5"
	CampaignRuleParametersPriorityNr5 string = "5"
)

// prop value enum
func (m *CampaignRuleParameters) validatePriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, campaignRuleParametersTypePriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CampaignRuleParameters) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriorityEnum("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this campaign rule parameters based on context it is used
func (m *CampaignRuleParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CampaignRuleParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CampaignRuleParameters) UnmarshalBinary(b []byte) error {
	var res CampaignRuleParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
