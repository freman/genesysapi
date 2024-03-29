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

// CallMediaParticipant call media participant
//
// swagger:model CallMediaParticipant
type CallMediaParticipant struct {

	// The participant address.
	Address string `json:"address,omitempty"`

	// Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs int32 `json:"alertingTimeoutMs,omitempty"`

	// The call ANI.
	Ani string `json:"ani,omitempty"`

	// A list of ad-hoc attributes for the participant.
	Attributes map[string]string `json:"attributes,omitempty"`

	// If this participant barged in a participant's call, then this will be the id of the targeted participant.
	BargedParticipantID string `json:"bargedParticipantId,omitempty"`

	// The timestamp when this participant was connected to the barge conference in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	BargedTime strfmt.DateTime `json:"bargedTime,omitempty"`

	// The ID of the participant being coached when performing a call coach.
	CoachedParticipantID string `json:"coachedParticipantId,omitempty"`

	// Value is true when the call is confined.
	Confined bool `json:"confined"`

	// The time when this participant went connected for this media (eg: video connected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// The ID of the consult transfer target participant when performing a consult transfer.
	ConsultParticipantID string `json:"consultParticipantId,omitempty"`

	// Information on how a communication should be routed to an agent.
	ConversationRoutingData *ConversationRoutingData `json:"conversationRoutingData,omitempty"`

	// The participant's direction.  Values can be: 'inbound' or 'outbound'
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// The reason the participant was disconnected from the conversation.
	// Enum: [endpoint client system transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam]
	DisconnectType string `json:"disconnectType,omitempty"`

	// The call DNIS.
	Dnis string `json:"dnis,omitempty"`

	// The ID of the Content Management document if the call is a fax.
	DocumentID string `json:"documentId,omitempty"`

	// The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndAcwTime strfmt.DateTime `json:"endAcwTime,omitempty"`

	// The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// If the conversation ends in error, contains additional error details.
	ErrorInfo *ErrorInfo `json:"errorInfo,omitempty"`

	// If this participant represents an external contact, then this will be the reference for the external contact.
	ExternalContact *DomainEntityRef `json:"externalContact,omitempty"`

	// If this participant represents an external org, then this will be the reference for the external org.
	ExternalOrganization *DomainEntityRef `json:"externalOrganization,omitempty"`

	// Extra fax information if the call is a fax.
	FaxStatus *FaxStatus `json:"faxStatus,omitempty"`

	// The reason specifying why participant flagged the conversation.
	// Enum: [general]
	FlaggedReason string `json:"flaggedReason,omitempty"`

	// The group involved in the group ring call.
	Group *DomainEntityRef `json:"group,omitempty"`

	// Value is true when the participant is on hold.
	Held bool `json:"held"`

	// The unique participant ID.
	ID string `json:"id,omitempty"`

	// Journey System data/context that is applicable to this communication.  When used for historical purposes, the context should be immutable.  When null, there is no applicable Journey System context.
	JourneyContext *JourneyContext `json:"journeyContext,omitempty"`

	// List of roles this participant's media has had on the conversation, ie monitor, coach, etc
	MediaRoles []string `json:"mediaRoles"`

	// The ID of the participant being monitored when performing a call monitor.
	MonitoredParticipantID string `json:"monitoredParticipantId,omitempty"`

	// Value is true when the call is muted.
	Muted bool `json:"muted"`

	// The display friendly name of the participant.
	Name string `json:"name,omitempty"`

	// The peer communication corresponding to a matching leg for this communication.
	Peer string `json:"peer,omitempty"`

	// The source provider for the communication.
	Provider string `json:"provider,omitempty"`

	// The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
	Purpose string `json:"purpose,omitempty"`

	// The PureCloud queue for this participant.
	Queue *DomainEntityRef `json:"queue,omitempty"`

	// Value is true when the call is being recorded.
	Recording bool `json:"recording"`

	// The state of the call recording.
	// Enum: [none active paused]
	RecordingState string `json:"recordingState,omitempty"`

	// The Engage script that should be used by this participant.
	Script *DomainEntityRef `json:"script,omitempty"`

	// The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartAcwTime strfmt.DateTime `json:"startAcwTime,omitempty"`

	// The time when this participant's hold started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartHoldTime strfmt.DateTime `json:"startHoldTime,omitempty"`

	// The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// The participant's state.  Values can be: 'alerting', 'connected', 'disconnected', 'dialing', 'contacting
	// Enum: [alerting dialing contacting offering connected disconnected terminated converting uploading transmitting none]
	State string `json:"state,omitempty"`

	// The PureCloud team for this participant.
	Team *DomainEntityRef `json:"team,omitempty"`

	// The PureCloud user for this participant.
	User *DomainEntityRef `json:"user,omitempty"`

	// User-to-User information which maps to a SIP header field defined in RFC7433. UUI data is used in the Public Switched Telephone Network (PSTN) for use cases described in RFC6567.
	UuiData string `json:"uuiData,omitempty"`

	// Wrapup for this participant, if it has been applied.
	Wrapup *Wrapup `json:"wrapup,omitempty"`

	// The wrap-up prompt indicating the type of wrap-up to be performed.
	WrapupPrompt string `json:"wrapupPrompt,omitempty"`

	// Value is true when the participant requires wrap-up.
	WrapupRequired bool `json:"wrapupRequired"`

	// Value is true when the participant has skipped wrap-up.
	WrapupSkipped bool `json:"wrapupSkipped"`

	// The amount of time the participant has to complete wrap-up.
	WrapupTimeoutMs int32 `json:"wrapupTimeoutMs,omitempty"`
}

// Validate validates this call media participant
func (m *CallMediaParticipant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBargedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationRoutingData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisconnectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndAcwTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaxStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlaggedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJourneyContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordingState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAcwTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartHoldTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrapup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallMediaParticipant) validateBargedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.BargedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("bargedTime", "body", "date-time", m.BargedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateConnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateConversationRoutingData(formats strfmt.Registry) error {
	if swag.IsZero(m.ConversationRoutingData) { // not required
		return nil
	}

	if m.ConversationRoutingData != nil {
		if err := m.ConversationRoutingData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversationRoutingData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversationRoutingData")
			}
			return err
		}
	}

	return nil
}

var callMediaParticipantTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callMediaParticipantTypeDirectionPropEnum = append(callMediaParticipantTypeDirectionPropEnum, v)
	}
}

const (

	// CallMediaParticipantDirectionInbound captures enum value "inbound"
	CallMediaParticipantDirectionInbound string = "inbound"

	// CallMediaParticipantDirectionOutbound captures enum value "outbound"
	CallMediaParticipantDirectionOutbound string = "outbound"
)

// prop value enum
func (m *CallMediaParticipant) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callMediaParticipantTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallMediaParticipant) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var callMediaParticipantTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callMediaParticipantTypeDisconnectTypePropEnum = append(callMediaParticipantTypeDisconnectTypePropEnum, v)
	}
}

const (

	// CallMediaParticipantDisconnectTypeEndpoint captures enum value "endpoint"
	CallMediaParticipantDisconnectTypeEndpoint string = "endpoint"

	// CallMediaParticipantDisconnectTypeClient captures enum value "client"
	CallMediaParticipantDisconnectTypeClient string = "client"

	// CallMediaParticipantDisconnectTypeSystem captures enum value "system"
	CallMediaParticipantDisconnectTypeSystem string = "system"

	// CallMediaParticipantDisconnectTypeTransfer captures enum value "transfer"
	CallMediaParticipantDisconnectTypeTransfer string = "transfer"

	// CallMediaParticipantDisconnectTypeTransferDotConference captures enum value "transfer.conference"
	CallMediaParticipantDisconnectTypeTransferDotConference string = "transfer.conference"

	// CallMediaParticipantDisconnectTypeTransferDotConsult captures enum value "transfer.consult"
	CallMediaParticipantDisconnectTypeTransferDotConsult string = "transfer.consult"

	// CallMediaParticipantDisconnectTypeTransferDotForward captures enum value "transfer.forward"
	CallMediaParticipantDisconnectTypeTransferDotForward string = "transfer.forward"

	// CallMediaParticipantDisconnectTypeTransferDotNoanswer captures enum value "transfer.noanswer"
	CallMediaParticipantDisconnectTypeTransferDotNoanswer string = "transfer.noanswer"

	// CallMediaParticipantDisconnectTypeTransferDotNotavailable captures enum value "transfer.notavailable"
	CallMediaParticipantDisconnectTypeTransferDotNotavailable string = "transfer.notavailable"

	// CallMediaParticipantDisconnectTypeTransportDotFailure captures enum value "transport.failure"
	CallMediaParticipantDisconnectTypeTransportDotFailure string = "transport.failure"

	// CallMediaParticipantDisconnectTypeError captures enum value "error"
	CallMediaParticipantDisconnectTypeError string = "error"

	// CallMediaParticipantDisconnectTypePeer captures enum value "peer"
	CallMediaParticipantDisconnectTypePeer string = "peer"

	// CallMediaParticipantDisconnectTypeOther captures enum value "other"
	CallMediaParticipantDisconnectTypeOther string = "other"

	// CallMediaParticipantDisconnectTypeSpam captures enum value "spam"
	CallMediaParticipantDisconnectTypeSpam string = "spam"
)

// prop value enum
func (m *CallMediaParticipant) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callMediaParticipantTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallMediaParticipant) validateDisconnectType(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateEndAcwTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endAcwTime", "body", "date-time", m.EndAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateErrorInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorInfo) { // not required
		return nil
	}

	if m.ErrorInfo != nil {
		if err := m.ErrorInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) validateExternalContact(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalContact) { // not required
		return nil
	}

	if m.ExternalContact != nil {
		if err := m.ExternalContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalContact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalContact")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) validateExternalOrganization(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalOrganization) { // not required
		return nil
	}

	if m.ExternalOrganization != nil {
		if err := m.ExternalOrganization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalOrganization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalOrganization")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) validateFaxStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.FaxStatus) { // not required
		return nil
	}

	if m.FaxStatus != nil {
		if err := m.FaxStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("faxStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("faxStatus")
			}
			return err
		}
	}

	return nil
}

var callMediaParticipantTypeFlaggedReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callMediaParticipantTypeFlaggedReasonPropEnum = append(callMediaParticipantTypeFlaggedReasonPropEnum, v)
	}
}

const (

	// CallMediaParticipantFlaggedReasonGeneral captures enum value "general"
	CallMediaParticipantFlaggedReasonGeneral string = "general"
)

// prop value enum
func (m *CallMediaParticipant) validateFlaggedReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callMediaParticipantTypeFlaggedReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallMediaParticipant) validateFlaggedReason(formats strfmt.Registry) error {
	if swag.IsZero(m.FlaggedReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlaggedReasonEnum("flaggedReason", "body", m.FlaggedReason); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) validateJourneyContext(formats strfmt.Registry) error {
	if swag.IsZero(m.JourneyContext) { // not required
		return nil
	}

	if m.JourneyContext != nil {
		if err := m.JourneyContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("journeyContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("journeyContext")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

var callMediaParticipantTypeRecordingStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","active","paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callMediaParticipantTypeRecordingStatePropEnum = append(callMediaParticipantTypeRecordingStatePropEnum, v)
	}
}

const (

	// CallMediaParticipantRecordingStateNone captures enum value "none"
	CallMediaParticipantRecordingStateNone string = "none"

	// CallMediaParticipantRecordingStateActive captures enum value "active"
	CallMediaParticipantRecordingStateActive string = "active"

	// CallMediaParticipantRecordingStatePaused captures enum value "paused"
	CallMediaParticipantRecordingStatePaused string = "paused"
)

// prop value enum
func (m *CallMediaParticipant) validateRecordingStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callMediaParticipantTypeRecordingStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallMediaParticipant) validateRecordingState(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordingState) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecordingStateEnum("recordingState", "body", m.RecordingState); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateScript(formats strfmt.Registry) error {
	if swag.IsZero(m.Script) { // not required
		return nil
	}

	if m.Script != nil {
		if err := m.Script.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) validateStartAcwTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAcwTime", "body", "date-time", m.StartAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateStartHoldTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var callMediaParticipantTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","converting","uploading","transmitting","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callMediaParticipantTypeStatePropEnum = append(callMediaParticipantTypeStatePropEnum, v)
	}
}

const (

	// CallMediaParticipantStateAlerting captures enum value "alerting"
	CallMediaParticipantStateAlerting string = "alerting"

	// CallMediaParticipantStateDialing captures enum value "dialing"
	CallMediaParticipantStateDialing string = "dialing"

	// CallMediaParticipantStateContacting captures enum value "contacting"
	CallMediaParticipantStateContacting string = "contacting"

	// CallMediaParticipantStateOffering captures enum value "offering"
	CallMediaParticipantStateOffering string = "offering"

	// CallMediaParticipantStateConnected captures enum value "connected"
	CallMediaParticipantStateConnected string = "connected"

	// CallMediaParticipantStateDisconnected captures enum value "disconnected"
	CallMediaParticipantStateDisconnected string = "disconnected"

	// CallMediaParticipantStateTerminated captures enum value "terminated"
	CallMediaParticipantStateTerminated string = "terminated"

	// CallMediaParticipantStateConverting captures enum value "converting"
	CallMediaParticipantStateConverting string = "converting"

	// CallMediaParticipantStateUploading captures enum value "uploading"
	CallMediaParticipantStateUploading string = "uploading"

	// CallMediaParticipantStateTransmitting captures enum value "transmitting"
	CallMediaParticipantStateTransmitting string = "transmitting"

	// CallMediaParticipantStateNone captures enum value "none"
	CallMediaParticipantStateNone string = "none"
)

// prop value enum
func (m *CallMediaParticipant) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callMediaParticipantTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallMediaParticipant) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *CallMediaParticipant) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) validateWrapup(formats strfmt.Registry) error {
	if swag.IsZero(m.Wrapup) { // not required
		return nil
	}

	if m.Wrapup != nil {
		if err := m.Wrapup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wrapup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wrapup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this call media participant based on the context it is used
func (m *CallMediaParticipant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConversationRoutingData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalContact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFaxStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJourneyContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrapup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallMediaParticipant) contextValidateConversationRoutingData(ctx context.Context, formats strfmt.Registry) error {

	if m.ConversationRoutingData != nil {
		if err := m.ConversationRoutingData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversationRoutingData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversationRoutingData")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateErrorInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorInfo != nil {
		if err := m.ErrorInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateExternalContact(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalContact != nil {
		if err := m.ExternalContact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalContact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalContact")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateExternalOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalOrganization != nil {
		if err := m.ExternalOrganization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalOrganization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalOrganization")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateFaxStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.FaxStatus != nil {
		if err := m.FaxStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("faxStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("faxStatus")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateJourneyContext(ctx context.Context, formats strfmt.Registry) error {

	if m.JourneyContext != nil {
		if err := m.JourneyContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("journeyContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("journeyContext")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {
		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateScript(ctx context.Context, formats strfmt.Registry) error {

	if m.Script != nil {
		if err := m.Script.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.Team != nil {
		if err := m.Team.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *CallMediaParticipant) contextValidateWrapup(ctx context.Context, formats strfmt.Registry) error {

	if m.Wrapup != nil {
		if err := m.Wrapup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wrapup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wrapup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallMediaParticipant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallMediaParticipant) UnmarshalBinary(b []byte) error {
	var res CallMediaParticipant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
