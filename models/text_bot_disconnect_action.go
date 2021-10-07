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

// TextBotDisconnectAction Settings for a next-action of disconnecting, including the reason code for the disconnect.
//
// swagger:model TextBotDisconnectAction
type TextBotDisconnectAction struct {

	// Describes where in the Bot Flow the user was when the disconnect occurred.
	FlowLocation *TextBotFlowLocation `json:"flowLocation,omitempty"`

	// The list of Flow Outcomes for the bot flow and their details.
	FlowOutcomes []*TextBotFlowOutcome `json:"flowOutcomes"`

	// The reason for the disconnect.
	// Required: true
	// Enum: [TriggeredByUser TriggeredByFlow SessionExpired Error RecognitionFailure]
	Reason *string `json:"reason"`

	// Extended information related to the reason, if available.
	ReasonExtendedInfo string `json:"reasonExtendedInfo,omitempty"`
}

// Validate validates this text bot disconnect action
func (m *TextBotDisconnectAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlowLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowOutcomes(formats); err != nil {
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

func (m *TextBotDisconnectAction) validateFlowLocation(formats strfmt.Registry) error {

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

func (m *TextBotDisconnectAction) validateFlowOutcomes(formats strfmt.Registry) error {

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

var textBotDisconnectActionTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TriggeredByUser","TriggeredByFlow","SessionExpired","Error","RecognitionFailure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		textBotDisconnectActionTypeReasonPropEnum = append(textBotDisconnectActionTypeReasonPropEnum, v)
	}
}

const (

	// TextBotDisconnectActionReasonTriggeredByUser captures enum value "TriggeredByUser"
	TextBotDisconnectActionReasonTriggeredByUser string = "TriggeredByUser"

	// TextBotDisconnectActionReasonTriggeredByFlow captures enum value "TriggeredByFlow"
	TextBotDisconnectActionReasonTriggeredByFlow string = "TriggeredByFlow"

	// TextBotDisconnectActionReasonSessionExpired captures enum value "SessionExpired"
	TextBotDisconnectActionReasonSessionExpired string = "SessionExpired"

	// TextBotDisconnectActionReasonError captures enum value "Error"
	TextBotDisconnectActionReasonError string = "Error"

	// TextBotDisconnectActionReasonRecognitionFailure captures enum value "RecognitionFailure"
	TextBotDisconnectActionReasonRecognitionFailure string = "RecognitionFailure"
)

// prop value enum
func (m *TextBotDisconnectAction) validateReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, textBotDisconnectActionTypeReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TextBotDisconnectAction) validateReason(formats strfmt.Registry) error {

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
func (m *TextBotDisconnectAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextBotDisconnectAction) UnmarshalBinary(b []byte) error {
	var res TextBotDisconnectAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
