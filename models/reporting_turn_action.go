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

// ReportingTurnAction reporting turn action
//
// swagger:model ReportingTurnAction
type ReportingTurnAction struct {

	// The ID of the action in the bot flow.
	ActionID string `json:"actionId,omitempty"`

	// The name of the action in the bot flow.
	ActionName string `json:"actionName,omitempty"`

	// The number of the action in the bot flow.
	ActionNumber int32 `json:"actionNumber,omitempty"`

	// action type
	// Enum: [AskForNLUIntentAction AskForBooleanAction AskForSlotAction AskForNLUNextIntentAction BotState DisconnectAction ExitBotFlowAction CallTaskAction EndTaskAction UpdateVariableAction CommunicateAction DecisionAction SwitchAction DataAction DataTableLookupAction GetExternalContactAction GetExternalOrganizationAction LoopAction ExitLoopAction NextLoopAction LoopUntilAction SetActiveIntentAction ClearSlotAction SetFlowOutcomeAction InitializeFlowOutcomeAction AddFlowMilestoneAction AskForStringAction SendResponseAction GetResponseAction ExtractSecureDataAction SecureAction TransferTaskAction DigitalMenuAction WaitForInputAction Unknown]
	ActionType string `json:"actionType,omitempty"`
}

// Validate validates this reporting turn action
func (m *ReportingTurnAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reportingTurnActionTypeActionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AskForNLUIntentAction","AskForBooleanAction","AskForSlotAction","AskForNLUNextIntentAction","BotState","DisconnectAction","ExitBotFlowAction","CallTaskAction","EndTaskAction","UpdateVariableAction","CommunicateAction","DecisionAction","SwitchAction","DataAction","DataTableLookupAction","GetExternalContactAction","GetExternalOrganizationAction","LoopAction","ExitLoopAction","NextLoopAction","LoopUntilAction","SetActiveIntentAction","ClearSlotAction","SetFlowOutcomeAction","InitializeFlowOutcomeAction","AddFlowMilestoneAction","AskForStringAction","SendResponseAction","GetResponseAction","ExtractSecureDataAction","SecureAction","TransferTaskAction","DigitalMenuAction","WaitForInputAction","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingTurnActionTypeActionTypePropEnum = append(reportingTurnActionTypeActionTypePropEnum, v)
	}
}

const (

	// ReportingTurnActionActionTypeAskForNLUIntentAction captures enum value "AskForNLUIntentAction"
	ReportingTurnActionActionTypeAskForNLUIntentAction string = "AskForNLUIntentAction"

	// ReportingTurnActionActionTypeAskForBooleanAction captures enum value "AskForBooleanAction"
	ReportingTurnActionActionTypeAskForBooleanAction string = "AskForBooleanAction"

	// ReportingTurnActionActionTypeAskForSlotAction captures enum value "AskForSlotAction"
	ReportingTurnActionActionTypeAskForSlotAction string = "AskForSlotAction"

	// ReportingTurnActionActionTypeAskForNLUNextIntentAction captures enum value "AskForNLUNextIntentAction"
	ReportingTurnActionActionTypeAskForNLUNextIntentAction string = "AskForNLUNextIntentAction"

	// ReportingTurnActionActionTypeBotState captures enum value "BotState"
	ReportingTurnActionActionTypeBotState string = "BotState"

	// ReportingTurnActionActionTypeDisconnectAction captures enum value "DisconnectAction"
	ReportingTurnActionActionTypeDisconnectAction string = "DisconnectAction"

	// ReportingTurnActionActionTypeExitBotFlowAction captures enum value "ExitBotFlowAction"
	ReportingTurnActionActionTypeExitBotFlowAction string = "ExitBotFlowAction"

	// ReportingTurnActionActionTypeCallTaskAction captures enum value "CallTaskAction"
	ReportingTurnActionActionTypeCallTaskAction string = "CallTaskAction"

	// ReportingTurnActionActionTypeEndTaskAction captures enum value "EndTaskAction"
	ReportingTurnActionActionTypeEndTaskAction string = "EndTaskAction"

	// ReportingTurnActionActionTypeUpdateVariableAction captures enum value "UpdateVariableAction"
	ReportingTurnActionActionTypeUpdateVariableAction string = "UpdateVariableAction"

	// ReportingTurnActionActionTypeCommunicateAction captures enum value "CommunicateAction"
	ReportingTurnActionActionTypeCommunicateAction string = "CommunicateAction"

	// ReportingTurnActionActionTypeDecisionAction captures enum value "DecisionAction"
	ReportingTurnActionActionTypeDecisionAction string = "DecisionAction"

	// ReportingTurnActionActionTypeSwitchAction captures enum value "SwitchAction"
	ReportingTurnActionActionTypeSwitchAction string = "SwitchAction"

	// ReportingTurnActionActionTypeDataAction captures enum value "DataAction"
	ReportingTurnActionActionTypeDataAction string = "DataAction"

	// ReportingTurnActionActionTypeDataTableLookupAction captures enum value "DataTableLookupAction"
	ReportingTurnActionActionTypeDataTableLookupAction string = "DataTableLookupAction"

	// ReportingTurnActionActionTypeGetExternalContactAction captures enum value "GetExternalContactAction"
	ReportingTurnActionActionTypeGetExternalContactAction string = "GetExternalContactAction"

	// ReportingTurnActionActionTypeGetExternalOrganizationAction captures enum value "GetExternalOrganizationAction"
	ReportingTurnActionActionTypeGetExternalOrganizationAction string = "GetExternalOrganizationAction"

	// ReportingTurnActionActionTypeLoopAction captures enum value "LoopAction"
	ReportingTurnActionActionTypeLoopAction string = "LoopAction"

	// ReportingTurnActionActionTypeExitLoopAction captures enum value "ExitLoopAction"
	ReportingTurnActionActionTypeExitLoopAction string = "ExitLoopAction"

	// ReportingTurnActionActionTypeNextLoopAction captures enum value "NextLoopAction"
	ReportingTurnActionActionTypeNextLoopAction string = "NextLoopAction"

	// ReportingTurnActionActionTypeLoopUntilAction captures enum value "LoopUntilAction"
	ReportingTurnActionActionTypeLoopUntilAction string = "LoopUntilAction"

	// ReportingTurnActionActionTypeSetActiveIntentAction captures enum value "SetActiveIntentAction"
	ReportingTurnActionActionTypeSetActiveIntentAction string = "SetActiveIntentAction"

	// ReportingTurnActionActionTypeClearSlotAction captures enum value "ClearSlotAction"
	ReportingTurnActionActionTypeClearSlotAction string = "ClearSlotAction"

	// ReportingTurnActionActionTypeSetFlowOutcomeAction captures enum value "SetFlowOutcomeAction"
	ReportingTurnActionActionTypeSetFlowOutcomeAction string = "SetFlowOutcomeAction"

	// ReportingTurnActionActionTypeInitializeFlowOutcomeAction captures enum value "InitializeFlowOutcomeAction"
	ReportingTurnActionActionTypeInitializeFlowOutcomeAction string = "InitializeFlowOutcomeAction"

	// ReportingTurnActionActionTypeAddFlowMilestoneAction captures enum value "AddFlowMilestoneAction"
	ReportingTurnActionActionTypeAddFlowMilestoneAction string = "AddFlowMilestoneAction"

	// ReportingTurnActionActionTypeAskForStringAction captures enum value "AskForStringAction"
	ReportingTurnActionActionTypeAskForStringAction string = "AskForStringAction"

	// ReportingTurnActionActionTypeSendResponseAction captures enum value "SendResponseAction"
	ReportingTurnActionActionTypeSendResponseAction string = "SendResponseAction"

	// ReportingTurnActionActionTypeGetResponseAction captures enum value "GetResponseAction"
	ReportingTurnActionActionTypeGetResponseAction string = "GetResponseAction"

	// ReportingTurnActionActionTypeExtractSecureDataAction captures enum value "ExtractSecureDataAction"
	ReportingTurnActionActionTypeExtractSecureDataAction string = "ExtractSecureDataAction"

	// ReportingTurnActionActionTypeSecureAction captures enum value "SecureAction"
	ReportingTurnActionActionTypeSecureAction string = "SecureAction"

	// ReportingTurnActionActionTypeTransferTaskAction captures enum value "TransferTaskAction"
	ReportingTurnActionActionTypeTransferTaskAction string = "TransferTaskAction"

	// ReportingTurnActionActionTypeDigitalMenuAction captures enum value "DigitalMenuAction"
	ReportingTurnActionActionTypeDigitalMenuAction string = "DigitalMenuAction"

	// ReportingTurnActionActionTypeWaitForInputAction captures enum value "WaitForInputAction"
	ReportingTurnActionActionTypeWaitForInputAction string = "WaitForInputAction"

	// ReportingTurnActionActionTypeUnknown captures enum value "Unknown"
	ReportingTurnActionActionTypeUnknown string = "Unknown"
)

// prop value enum
func (m *ReportingTurnAction) validateActionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingTurnActionTypeActionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingTurnAction) validateActionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionTypeEnum("actionType", "body", m.ActionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportingTurnAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingTurnAction) UnmarshalBinary(b []byte) error {
	var res ReportingTurnAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
