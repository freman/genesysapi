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

// CobrowseMediaParticipant cobrowse media participant
//
// swagger:model CobrowseMediaParticipant
type CobrowseMediaParticipant struct {

	// The participant address.
	Address string `json:"address,omitempty"`

	// Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs int32 `json:"alertingTimeoutMs,omitempty"`

	// A list of ad-hoc attributes for the participant.
	Attributes map[string]string `json:"attributes,omitempty"`

	// This value identifies the role of the co-browse client within the co-browse session (a client is a sharer or a viewer).
	CobrowseRole string `json:"cobrowseRole,omitempty"`

	// The co-browse session ID.
	CobrowseSessionID string `json:"cobrowseSessionId,omitempty"`

	// The time when this participant went connected for this media (eg: video connected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// ID of co-browse participants for which this client has been granted control (list is empty if this client cannot control any shared pages).
	Controlling []string `json:"controlling"`

	// Information on how a communication should be routed to an agent.
	ConversationRoutingData *ConversationRoutingData `json:"conversationRoutingData,omitempty"`

	// The participant's direction.  Values can be: 'inbound' or 'outbound'
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// The reason the participant was disconnected from the conversation.
	// Enum: [endpoint client system transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam]
	DisconnectType string `json:"disconnectType,omitempty"`

	// The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	EndAcwTime strfmt.DateTime `json:"endAcwTime,omitempty"`

	// The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// If the conversation ends in error, contains additional error details.
	ErrorInfo *ErrorInfo `json:"errorInfo,omitempty"`

	// If this participant represents an external contact, then this will be the reference for the external contact.
	ExternalContact *DomainEntityRef `json:"externalContact,omitempty"`

	// If this participant represents an external org, then this will be the reference for the external org.
	ExternalOrganization *DomainEntityRef `json:"externalOrganization,omitempty"`

	// The reason specifying why participant flagged the conversation.
	// Enum: [general]
	FlaggedReason string `json:"flaggedReason,omitempty"`

	// Value is true when the participant is on hold.
	Held bool `json:"held"`

	// The unique participant ID.
	ID string `json:"id,omitempty"`

	// Journey System data/context that is applicable to this communication.  When used for historical purposes, the context should be immutable.  When null, there is no applicable Journey System context.
	JourneyContext *JourneyContext `json:"journeyContext,omitempty"`

	// The display friendly name of the participant.
	Name string `json:"name,omitempty"`

	// The peer communication corresponding to a matching leg for this communication.
	Peer string `json:"peer,omitempty"`

	// The source provider for the communication.
	Provider string `json:"provider,omitempty"`

	// The time when the provider event which triggered this conversation update happened in the corrected provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ProviderEventTime strfmt.DateTime `json:"providerEventTime,omitempty"`

	// The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
	Purpose string `json:"purpose,omitempty"`

	// The PureCloud queue for this participant.
	Queue *DomainEntityRef `json:"queue,omitempty"`

	// The Engage script that should be used by this participant.
	Script *DomainEntityRef `json:"script,omitempty"`

	// The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	StartAcwTime strfmt.DateTime `json:"startAcwTime,omitempty"`

	// The time when this participant's hold started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	StartHoldTime strfmt.DateTime `json:"startHoldTime,omitempty"`

	// The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// The participant's state.  Values can be: 'alerting', 'connected', 'disconnected', 'dialing', 'contacting
	// Enum: [alerting dialing contacting offering connected disconnected terminated converting uploading transmitting none]
	State string `json:"state,omitempty"`

	// The PureCloud team for this participant.
	Team *DomainEntityRef `json:"team,omitempty"`

	// The PureCloud user for this participant.
	User *DomainEntityRef `json:"user,omitempty"`

	// The URL that can be used to open co-browse session in web browser.
	ViewerURL string `json:"viewerUrl,omitempty"`

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

// Validate validates this cobrowse media participant
func (m *CobrowseMediaParticipant) Validate(formats strfmt.Registry) error {
	var res []error

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

	if err := m.validateFlaggedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJourneyContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
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

func (m *CobrowseMediaParticipant) validateConnectedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateConversationRoutingData(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationRoutingData) { // not required
		return nil
	}

	if m.ConversationRoutingData != nil {
		if err := m.ConversationRoutingData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversationRoutingData")
			}
			return err
		}
	}

	return nil
}

var cobrowseMediaParticipantTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cobrowseMediaParticipantTypeDirectionPropEnum = append(cobrowseMediaParticipantTypeDirectionPropEnum, v)
	}
}

const (

	// CobrowseMediaParticipantDirectionInbound captures enum value "inbound"
	CobrowseMediaParticipantDirectionInbound string = "inbound"

	// CobrowseMediaParticipantDirectionOutbound captures enum value "outbound"
	CobrowseMediaParticipantDirectionOutbound string = "outbound"
)

// prop value enum
func (m *CobrowseMediaParticipant) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cobrowseMediaParticipantTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CobrowseMediaParticipant) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var cobrowseMediaParticipantTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cobrowseMediaParticipantTypeDisconnectTypePropEnum = append(cobrowseMediaParticipantTypeDisconnectTypePropEnum, v)
	}
}

const (

	// CobrowseMediaParticipantDisconnectTypeEndpoint captures enum value "endpoint"
	CobrowseMediaParticipantDisconnectTypeEndpoint string = "endpoint"

	// CobrowseMediaParticipantDisconnectTypeClient captures enum value "client"
	CobrowseMediaParticipantDisconnectTypeClient string = "client"

	// CobrowseMediaParticipantDisconnectTypeSystem captures enum value "system"
	CobrowseMediaParticipantDisconnectTypeSystem string = "system"

	// CobrowseMediaParticipantDisconnectTypeTransfer captures enum value "transfer"
	CobrowseMediaParticipantDisconnectTypeTransfer string = "transfer"

	// CobrowseMediaParticipantDisconnectTypeTransferConference captures enum value "transfer.conference"
	CobrowseMediaParticipantDisconnectTypeTransferConference string = "transfer.conference"

	// CobrowseMediaParticipantDisconnectTypeTransferConsult captures enum value "transfer.consult"
	CobrowseMediaParticipantDisconnectTypeTransferConsult string = "transfer.consult"

	// CobrowseMediaParticipantDisconnectTypeTransferForward captures enum value "transfer.forward"
	CobrowseMediaParticipantDisconnectTypeTransferForward string = "transfer.forward"

	// CobrowseMediaParticipantDisconnectTypeTransferNoanswer captures enum value "transfer.noanswer"
	CobrowseMediaParticipantDisconnectTypeTransferNoanswer string = "transfer.noanswer"

	// CobrowseMediaParticipantDisconnectTypeTransferNotavailable captures enum value "transfer.notavailable"
	CobrowseMediaParticipantDisconnectTypeTransferNotavailable string = "transfer.notavailable"

	// CobrowseMediaParticipantDisconnectTypeTransportFailure captures enum value "transport.failure"
	CobrowseMediaParticipantDisconnectTypeTransportFailure string = "transport.failure"

	// CobrowseMediaParticipantDisconnectTypeError captures enum value "error"
	CobrowseMediaParticipantDisconnectTypeError string = "error"

	// CobrowseMediaParticipantDisconnectTypePeer captures enum value "peer"
	CobrowseMediaParticipantDisconnectTypePeer string = "peer"

	// CobrowseMediaParticipantDisconnectTypeOther captures enum value "other"
	CobrowseMediaParticipantDisconnectTypeOther string = "other"

	// CobrowseMediaParticipantDisconnectTypeSpam captures enum value "spam"
	CobrowseMediaParticipantDisconnectTypeSpam string = "spam"
)

// prop value enum
func (m *CobrowseMediaParticipant) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cobrowseMediaParticipantTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CobrowseMediaParticipant) validateDisconnectType(formats strfmt.Registry) error {

	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateEndAcwTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endAcwTime", "body", "date-time", m.EndAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateErrorInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorInfo) { // not required
		return nil
	}

	if m.ErrorInfo != nil {
		if err := m.ErrorInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateExternalContact(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalContact) { // not required
		return nil
	}

	if m.ExternalContact != nil {
		if err := m.ExternalContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalContact")
			}
			return err
		}
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateExternalOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalOrganization) { // not required
		return nil
	}

	if m.ExternalOrganization != nil {
		if err := m.ExternalOrganization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalOrganization")
			}
			return err
		}
	}

	return nil
}

var cobrowseMediaParticipantTypeFlaggedReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cobrowseMediaParticipantTypeFlaggedReasonPropEnum = append(cobrowseMediaParticipantTypeFlaggedReasonPropEnum, v)
	}
}

const (

	// CobrowseMediaParticipantFlaggedReasonGeneral captures enum value "general"
	CobrowseMediaParticipantFlaggedReasonGeneral string = "general"
)

// prop value enum
func (m *CobrowseMediaParticipant) validateFlaggedReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cobrowseMediaParticipantTypeFlaggedReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CobrowseMediaParticipant) validateFlaggedReason(formats strfmt.Registry) error {

	if swag.IsZero(m.FlaggedReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlaggedReasonEnum("flaggedReason", "body", m.FlaggedReason); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateJourneyContext(formats strfmt.Registry) error {

	if swag.IsZero(m.JourneyContext) { // not required
		return nil
	}

	if m.JourneyContext != nil {
		if err := m.JourneyContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("journeyContext")
			}
			return err
		}
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateProviderEventTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderEventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("providerEventTime", "body", "date-time", m.ProviderEventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateQueue(formats strfmt.Registry) error {

	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateScript(formats strfmt.Registry) error {

	if swag.IsZero(m.Script) { // not required
		return nil
	}

	if m.Script != nil {
		if err := m.Script.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateStartAcwTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAcwTime", "body", "date-time", m.StartAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateStartHoldTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var cobrowseMediaParticipantTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","converting","uploading","transmitting","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cobrowseMediaParticipantTypeStatePropEnum = append(cobrowseMediaParticipantTypeStatePropEnum, v)
	}
}

const (

	// CobrowseMediaParticipantStateAlerting captures enum value "alerting"
	CobrowseMediaParticipantStateAlerting string = "alerting"

	// CobrowseMediaParticipantStateDialing captures enum value "dialing"
	CobrowseMediaParticipantStateDialing string = "dialing"

	// CobrowseMediaParticipantStateContacting captures enum value "contacting"
	CobrowseMediaParticipantStateContacting string = "contacting"

	// CobrowseMediaParticipantStateOffering captures enum value "offering"
	CobrowseMediaParticipantStateOffering string = "offering"

	// CobrowseMediaParticipantStateConnected captures enum value "connected"
	CobrowseMediaParticipantStateConnected string = "connected"

	// CobrowseMediaParticipantStateDisconnected captures enum value "disconnected"
	CobrowseMediaParticipantStateDisconnected string = "disconnected"

	// CobrowseMediaParticipantStateTerminated captures enum value "terminated"
	CobrowseMediaParticipantStateTerminated string = "terminated"

	// CobrowseMediaParticipantStateConverting captures enum value "converting"
	CobrowseMediaParticipantStateConverting string = "converting"

	// CobrowseMediaParticipantStateUploading captures enum value "uploading"
	CobrowseMediaParticipantStateUploading string = "uploading"

	// CobrowseMediaParticipantStateTransmitting captures enum value "transmitting"
	CobrowseMediaParticipantStateTransmitting string = "transmitting"

	// CobrowseMediaParticipantStateNone captures enum value "none"
	CobrowseMediaParticipantStateNone string = "none"
)

// prop value enum
func (m *CobrowseMediaParticipant) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cobrowseMediaParticipantTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CobrowseMediaParticipant) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateTeam(formats strfmt.Registry) error {

	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *CobrowseMediaParticipant) validateWrapup(formats strfmt.Registry) error {

	if swag.IsZero(m.Wrapup) { // not required
		return nil
	}

	if m.Wrapup != nil {
		if err := m.Wrapup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wrapup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CobrowseMediaParticipant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CobrowseMediaParticipant) UnmarshalBinary(b []byte) error {
	var res CobrowseMediaParticipant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}