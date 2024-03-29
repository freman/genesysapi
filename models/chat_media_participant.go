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

// ChatMediaParticipant chat media participant
//
// swagger:model ChatMediaParticipant
type ChatMediaParticipant struct {

	// The participant address.
	Address string `json:"address,omitempty"`

	// Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs int32 `json:"alertingTimeoutMs,omitempty"`

	// A list of ad-hoc attributes for the participant.
	Attributes map[string]string `json:"attributes,omitempty"`

	// If available, the URI to the avatar image of this communication.
	AvatarImageURL string `json:"avatarImageUrl,omitempty"`

	// The time when this participant went connected for this media (eg: video connected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// Information on how a communication should be routed to an agent.
	ConversationRoutingData *ConversationRoutingData `json:"conversationRoutingData,omitempty"`

	// The participant's direction.  Values can be: 'inbound' or 'outbound'
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// The reason the participant was disconnected from the conversation.
	// Enum: [endpoint client system transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam]
	DisconnectType string `json:"disconnectType,omitempty"`

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

	// The reason specifying why participant flagged the conversation.
	// Enum: [general]
	FlaggedReason string `json:"flaggedReason,omitempty"`

	// Value is true when the participant is on hold.
	Held bool `json:"held"`

	// The unique participant ID.
	ID string `json:"id,omitempty"`

	// Journey System data/context that is applicable to this communication.  When used for historical purposes, the context should be immutable.  When null, there is no applicable Journey System context.
	JourneyContext *JourneyContext `json:"journeyContext,omitempty"`

	// List of roles this participant's media has had on the conversation, ie monitor, coach, etc
	MediaRoles []string `json:"mediaRoles"`

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

	// The ID of the chat room.
	RoomID string `json:"roomId,omitempty"`

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

// Validate validates this chat media participant
func (m *ChatMediaParticipant) Validate(formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) validateConnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChatMediaParticipant) validateConversationRoutingData(formats strfmt.Registry) error {
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

var chatMediaParticipantTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chatMediaParticipantTypeDirectionPropEnum = append(chatMediaParticipantTypeDirectionPropEnum, v)
	}
}

const (

	// ChatMediaParticipantDirectionInbound captures enum value "inbound"
	ChatMediaParticipantDirectionInbound string = "inbound"

	// ChatMediaParticipantDirectionOutbound captures enum value "outbound"
	ChatMediaParticipantDirectionOutbound string = "outbound"
)

// prop value enum
func (m *ChatMediaParticipant) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chatMediaParticipantTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChatMediaParticipant) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var chatMediaParticipantTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chatMediaParticipantTypeDisconnectTypePropEnum = append(chatMediaParticipantTypeDisconnectTypePropEnum, v)
	}
}

const (

	// ChatMediaParticipantDisconnectTypeEndpoint captures enum value "endpoint"
	ChatMediaParticipantDisconnectTypeEndpoint string = "endpoint"

	// ChatMediaParticipantDisconnectTypeClient captures enum value "client"
	ChatMediaParticipantDisconnectTypeClient string = "client"

	// ChatMediaParticipantDisconnectTypeSystem captures enum value "system"
	ChatMediaParticipantDisconnectTypeSystem string = "system"

	// ChatMediaParticipantDisconnectTypeTransfer captures enum value "transfer"
	ChatMediaParticipantDisconnectTypeTransfer string = "transfer"

	// ChatMediaParticipantDisconnectTypeTransferDotConference captures enum value "transfer.conference"
	ChatMediaParticipantDisconnectTypeTransferDotConference string = "transfer.conference"

	// ChatMediaParticipantDisconnectTypeTransferDotConsult captures enum value "transfer.consult"
	ChatMediaParticipantDisconnectTypeTransferDotConsult string = "transfer.consult"

	// ChatMediaParticipantDisconnectTypeTransferDotForward captures enum value "transfer.forward"
	ChatMediaParticipantDisconnectTypeTransferDotForward string = "transfer.forward"

	// ChatMediaParticipantDisconnectTypeTransferDotNoanswer captures enum value "transfer.noanswer"
	ChatMediaParticipantDisconnectTypeTransferDotNoanswer string = "transfer.noanswer"

	// ChatMediaParticipantDisconnectTypeTransferDotNotavailable captures enum value "transfer.notavailable"
	ChatMediaParticipantDisconnectTypeTransferDotNotavailable string = "transfer.notavailable"

	// ChatMediaParticipantDisconnectTypeTransportDotFailure captures enum value "transport.failure"
	ChatMediaParticipantDisconnectTypeTransportDotFailure string = "transport.failure"

	// ChatMediaParticipantDisconnectTypeError captures enum value "error"
	ChatMediaParticipantDisconnectTypeError string = "error"

	// ChatMediaParticipantDisconnectTypePeer captures enum value "peer"
	ChatMediaParticipantDisconnectTypePeer string = "peer"

	// ChatMediaParticipantDisconnectTypeOther captures enum value "other"
	ChatMediaParticipantDisconnectTypeOther string = "other"

	// ChatMediaParticipantDisconnectTypeSpam captures enum value "spam"
	ChatMediaParticipantDisconnectTypeSpam string = "spam"
)

// prop value enum
func (m *ChatMediaParticipant) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chatMediaParticipantTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChatMediaParticipant) validateDisconnectType(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *ChatMediaParticipant) validateEndAcwTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endAcwTime", "body", "date-time", m.EndAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChatMediaParticipant) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChatMediaParticipant) validateErrorInfo(formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) validateExternalContact(formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) validateExternalOrganization(formats strfmt.Registry) error {
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

var chatMediaParticipantTypeFlaggedReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chatMediaParticipantTypeFlaggedReasonPropEnum = append(chatMediaParticipantTypeFlaggedReasonPropEnum, v)
	}
}

const (

	// ChatMediaParticipantFlaggedReasonGeneral captures enum value "general"
	ChatMediaParticipantFlaggedReasonGeneral string = "general"
)

// prop value enum
func (m *ChatMediaParticipant) validateFlaggedReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chatMediaParticipantTypeFlaggedReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChatMediaParticipant) validateFlaggedReason(formats strfmt.Registry) error {
	if swag.IsZero(m.FlaggedReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlaggedReasonEnum("flaggedReason", "body", m.FlaggedReason); err != nil {
		return err
	}

	return nil
}

func (m *ChatMediaParticipant) validateJourneyContext(formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) validateQueue(formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) validateScript(formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) validateStartAcwTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAcwTime", "body", "date-time", m.StartAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChatMediaParticipant) validateStartHoldTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChatMediaParticipant) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var chatMediaParticipantTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","converting","uploading","transmitting","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chatMediaParticipantTypeStatePropEnum = append(chatMediaParticipantTypeStatePropEnum, v)
	}
}

const (

	// ChatMediaParticipantStateAlerting captures enum value "alerting"
	ChatMediaParticipantStateAlerting string = "alerting"

	// ChatMediaParticipantStateDialing captures enum value "dialing"
	ChatMediaParticipantStateDialing string = "dialing"

	// ChatMediaParticipantStateContacting captures enum value "contacting"
	ChatMediaParticipantStateContacting string = "contacting"

	// ChatMediaParticipantStateOffering captures enum value "offering"
	ChatMediaParticipantStateOffering string = "offering"

	// ChatMediaParticipantStateConnected captures enum value "connected"
	ChatMediaParticipantStateConnected string = "connected"

	// ChatMediaParticipantStateDisconnected captures enum value "disconnected"
	ChatMediaParticipantStateDisconnected string = "disconnected"

	// ChatMediaParticipantStateTerminated captures enum value "terminated"
	ChatMediaParticipantStateTerminated string = "terminated"

	// ChatMediaParticipantStateConverting captures enum value "converting"
	ChatMediaParticipantStateConverting string = "converting"

	// ChatMediaParticipantStateUploading captures enum value "uploading"
	ChatMediaParticipantStateUploading string = "uploading"

	// ChatMediaParticipantStateTransmitting captures enum value "transmitting"
	ChatMediaParticipantStateTransmitting string = "transmitting"

	// ChatMediaParticipantStateNone captures enum value "none"
	ChatMediaParticipantStateNone string = "none"
)

// prop value enum
func (m *ChatMediaParticipant) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chatMediaParticipantTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChatMediaParticipant) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *ChatMediaParticipant) validateTeam(formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) validateUser(formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) validateWrapup(formats strfmt.Registry) error {
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

// ContextValidate validate this chat media participant based on the context it is used
func (m *ChatMediaParticipant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ChatMediaParticipant) contextValidateConversationRoutingData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateErrorInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateExternalContact(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateExternalOrganization(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateJourneyContext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateScript(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ChatMediaParticipant) contextValidateWrapup(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ChatMediaParticipant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChatMediaParticipant) UnmarshalBinary(b []byte) error {
	var res ChatMediaParticipant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
