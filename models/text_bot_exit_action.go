// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TextBotExitAction Settings for a next-action of exiting the bot flow. Any output variables are available in the details.
//
// swagger:model TextBotExitAction
type TextBotExitAction struct {

	// The active intent at the time of the exit.
	ActiveIntent string `json:"activeIntent,omitempty"`

	// Describes where in the Bot Flow the user was when the exit occurred.
	FlowLocation *TextBotFlowLocation `json:"flowLocation,omitempty"`

	// The list of Flow Outcomes for the bot flow and their details.
	FlowOutcomes []*TextBotFlowOutcome `json:"flowOutcomes"`

	// The output data for the bot flow.
	OutputData *TextBotInputOutputData `json:"outputData,omitempty"`

	// The reason for the exit.
	// Required: true
	// Enum: [TriggeredByUser AgentRequestedByUser TriggeredByFlow Error RecognitionFailure]
	Reason *string `json:"reason"`

	// Extended information related to the reason, if available.
	ReasonExtendedInfo string `json:"reasonExtendedInfo,omitempty"`
}

// Validate validates this text bot exit action
func (m *TextBotExitAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlowLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowOutcomes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextBotExitAction) validateFlowLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowLocation) { // not required
		return nil
	}

	if m.FlowLocation != nil {
		if err := m.FlowLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowLocation")
			}
			return err
		}
	}

	return nil
}

func (m *TextBotExitAction) validateFlowOutcomes(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowOutcomes) { // not required
		return nil
	}

	for i := 0; i < len(m.FlowOutcomes); i++ {
		if swag.IsZero(m.FlowOutcomes[i]) { // not required
			continue
		}

		if m.FlowOutcomes[i] != nil {
			if err := m.FlowOutcomes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flowOutcomes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TextBotExitAction) validateOutputData(formats strfmt.Registry) error {

	if swag.IsZero(m.OutputData) { // not required
		return nil
	}

	if m.OutputData != nil {
		if err := m.OutputData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outputData")
			}
			return err
		}
	}

	return nil
}

var textBotExitActionTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TriggeredByUser","AgentRequestedByUser","TriggeredByFlow","Error","RecognitionFailure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		textBotExitActionTypeReasonPropEnum = append(textBotExitActionTypeReasonPropEnum, v)
	}
}

const (

	// TextBotExitActionReasonTriggeredByUser captures enum value "TriggeredByUser"
	TextBotExitActionReasonTriggeredByUser string = "TriggeredByUser"

	// TextBotExitActionReasonAgentRequestedByUser captures enum value "AgentRequestedByUser"
	TextBotExitActionReasonAgentRequestedByUser string = "AgentRequestedByUser"

	// TextBotExitActionReasonTriggeredByFlow captures enum value "TriggeredByFlow"
	TextBotExitActionReasonTriggeredByFlow string = "TriggeredByFlow"

	// TextBotExitActionReasonError captures enum value "Error"
	TextBotExitActionReasonError string = "Error"

	// TextBotExitActionReasonRecognitionFailure captures enum value "RecognitionFailure"
	TextBotExitActionReasonRecognitionFailure string = "RecognitionFailure"
)

// prop value enum
func (m *TextBotExitAction) validateReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, textBotExitActionTypeReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TextBotExitAction) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	// value enum
	if err := m.validateReasonEnum("reason", "body", *m.Reason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextBotExitAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextBotExitAction) UnmarshalBinary(b []byte) error {
	var res TextBotExitAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}