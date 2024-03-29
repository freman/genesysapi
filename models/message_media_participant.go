// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MessageMediaParticipant message media participant
//
// swagger:model MessageMediaParticipant
type MessageMediaParticipant struct {

	// The participant address.
	Address string `json:"address,omitempty"`

	// Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs int32 `json:"alertingTimeoutMs,omitempty"`

	// A list of ad-hoc attributes for the participant.
	Attributes map[string]string `json:"attributes,omitempty"`

	// If true, the participant member is authenticated.
	Authenticated bool `json:"authenticated"`

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

	// Address for the participant on the sending side of the message conversation. If the address is a phone number, E.164 format is recommended.
	FromAddress *Address `json:"fromAddress,omitempty"`

	// Value is true when the participant is on hold.
	Held bool `json:"held"`

	// The unique participant ID.
	ID string `json:"id,omitempty"`

	// Journey System data/context that is applicable to this communication.  When used for historical purposes, the context should be immutable.  When null, there is no applicable Journey System context.
	JourneyContext *JourneyContext `json:"journeyContext,omitempty"`

	// List of roles this participant's media has had on the conversation, ie monitor, coach, etc
	MediaRoles []string `json:"mediaRoles"`

	// Message instance details on the communication.
	Messages []*MessageDetails `json:"messages"`

	// The ID of the participant being monitored when performing a message monitor.
	MonitoredParticipantID string `json:"monitoredParticipantId,omitempty"`

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

	// Indicates the country where the recipient is associated in ISO 3166-1 alpha-2 format.
	RecipientCountry string `json:"recipientCountry,omitempty"`

	// The type of the recipient. Eg: Provisioned phoneNumber is the recipient for sms message type.
	RecipientType string `json:"recipientType,omitempty"`

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

	// Address for the participant on receiving side of the message conversation. If the address is a phone number, E.164 format is recommended.
	ToAddress *Address `json:"toAddress,omitempty"`

	// Indicates the type of message platform from which the message originated.
	// Enum: [unknown sms twitter facebook line whatsapp telegram kakao webmessaging open instagram]
	Type string `json:"type,omitempty"`

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

// Validate validates this message media participant
func (m *MessageMediaParticipant) Validate(formats strfmt.Registry) error {
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

	if err := m.validateFromAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJourneyContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
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

	if err := m.validateToAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *MessageMediaParticipant) validateConnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateConversationRoutingData(formats strfmt.Registry) error {
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

var messageMediaParticipantTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageMediaParticipantTypeDirectionPropEnum = append(messageMediaParticipantTypeDirectionPropEnum, v)
	}
}

const (

	// MessageMediaParticipantDirectionInbound captures enum value "inbound"
	MessageMediaParticipantDirectionInbound string = "inbound"

	// MessageMediaParticipantDirectionOutbound captures enum value "outbound"
	MessageMediaParticipantDirectionOutbound string = "outbound"
)

// prop value enum
func (m *MessageMediaParticipant) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageMediaParticipantTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageMediaParticipant) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var messageMediaParticipantTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageMediaParticipantTypeDisconnectTypePropEnum = append(messageMediaParticipantTypeDisconnectTypePropEnum, v)
	}
}

const (

	// MessageMediaParticipantDisconnectTypeEndpoint captures enum value "endpoint"
	MessageMediaParticipantDisconnectTypeEndpoint string = "endpoint"

	// MessageMediaParticipantDisconnectTypeClient captures enum value "client"
	MessageMediaParticipantDisconnectTypeClient string = "client"

	// MessageMediaParticipantDisconnectTypeSystem captures enum value "system"
	MessageMediaParticipantDisconnectTypeSystem string = "system"

	// MessageMediaParticipantDisconnectTypeTransfer captures enum value "transfer"
	MessageMediaParticipantDisconnectTypeTransfer string = "transfer"

	// MessageMediaParticipantDisconnectTypeTransferDotConference captures enum value "transfer.conference"
	MessageMediaParticipantDisconnectTypeTransferDotConference string = "transfer.conference"

	// MessageMediaParticipantDisconnectTypeTransferDotConsult captures enum value "transfer.consult"
	MessageMediaParticipantDisconnectTypeTransferDotConsult string = "transfer.consult"

	// MessageMediaParticipantDisconnectTypeTransferDotForward captures enum value "transfer.forward"
	MessageMediaParticipantDisconnectTypeTransferDotForward string = "transfer.forward"

	// MessageMediaParticipantDisconnectTypeTransferDotNoanswer captures enum value "transfer.noanswer"
	MessageMediaParticipantDisconnectTypeTransferDotNoanswer string = "transfer.noanswer"

	// MessageMediaParticipantDisconnectTypeTransferDotNotavailable captures enum value "transfer.notavailable"
	MessageMediaParticipantDisconnectTypeTransferDotNotavailable string = "transfer.notavailable"

	// MessageMediaParticipantDisconnectTypeTransportDotFailure captures enum value "transport.failure"
	MessageMediaParticipantDisconnectTypeTransportDotFailure string = "transport.failure"

	// MessageMediaParticipantDisconnectTypeError captures enum value "error"
	MessageMediaParticipantDisconnectTypeError string = "error"

	// MessageMediaParticipantDisconnectTypePeer captures enum value "peer"
	MessageMediaParticipantDisconnectTypePeer string = "peer"

	// MessageMediaParticipantDisconnectTypeOther captures enum value "other"
	MessageMediaParticipantDisconnectTypeOther string = "other"

	// MessageMediaParticipantDisconnectTypeSpam captures enum value "spam"
	MessageMediaParticipantDisconnectTypeSpam string = "spam"
)

// prop value enum
func (m *MessageMediaParticipant) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageMediaParticipantTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageMediaParticipant) validateDisconnectType(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateEndAcwTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endAcwTime", "body", "date-time", m.EndAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateErrorInfo(formats strfmt.Registry) error {
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

func (m *MessageMediaParticipant) validateExternalContact(formats strfmt.Registry) error {
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

func (m *MessageMediaParticipant) validateExternalOrganization(formats strfmt.Registry) error {
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

var messageMediaParticipantTypeFlaggedReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageMediaParticipantTypeFlaggedReasonPropEnum = append(messageMediaParticipantTypeFlaggedReasonPropEnum, v)
	}
}

const (

	// MessageMediaParticipantFlaggedReasonGeneral captures enum value "general"
	MessageMediaParticipantFlaggedReasonGeneral string = "general"
)

// prop value enum
func (m *MessageMediaParticipant) validateFlaggedReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageMediaParticipantTypeFlaggedReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageMediaParticipant) validateFlaggedReason(formats strfmt.Registry) error {
	if swag.IsZero(m.FlaggedReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlaggedReasonEnum("flaggedReason", "body", m.FlaggedReason); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateFromAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.FromAddress) { // not required
		return nil
	}

	if m.FromAddress != nil {
		if err := m.FromAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fromAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MessageMediaParticipant) validateJourneyContext(formats strfmt.Registry) error {
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

func (m *MessageMediaParticipant) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessageMediaParticipant) validateQueue(formats strfmt.Registry) error {
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

func (m *MessageMediaParticipant) validateScript(formats strfmt.Registry) error {
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

func (m *MessageMediaParticipant) validateStartAcwTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAcwTime", "body", "date-time", m.StartAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateStartHoldTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var messageMediaParticipantTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","converting","uploading","transmitting","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageMediaParticipantTypeStatePropEnum = append(messageMediaParticipantTypeStatePropEnum, v)
	}
}

const (

	// MessageMediaParticipantStateAlerting captures enum value "alerting"
	MessageMediaParticipantStateAlerting string = "alerting"

	// MessageMediaParticipantStateDialing captures enum value "dialing"
	MessageMediaParticipantStateDialing string = "dialing"

	// MessageMediaParticipantStateContacting captures enum value "contacting"
	MessageMediaParticipantStateContacting string = "contacting"

	// MessageMediaParticipantStateOffering captures enum value "offering"
	MessageMediaParticipantStateOffering string = "offering"

	// MessageMediaParticipantStateConnected captures enum value "connected"
	MessageMediaParticipantStateConnected string = "connected"

	// MessageMediaParticipantStateDisconnected captures enum value "disconnected"
	MessageMediaParticipantStateDisconnected string = "disconnected"

	// MessageMediaParticipantStateTerminated captures enum value "terminated"
	MessageMediaParticipantStateTerminated string = "terminated"

	// MessageMediaParticipantStateConverting captures enum value "converting"
	MessageMediaParticipantStateConverting string = "converting"

	// MessageMediaParticipantStateUploading captures enum value "uploading"
	MessageMediaParticipantStateUploading string = "uploading"

	// MessageMediaParticipantStateTransmitting captures enum value "transmitting"
	MessageMediaParticipantStateTransmitting string = "transmitting"

	// MessageMediaParticipantStateNone captures enum value "none"
	MessageMediaParticipantStateNone string = "none"
)

// prop value enum
func (m *MessageMediaParticipant) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageMediaParticipantTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageMediaParticipant) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateTeam(formats strfmt.Registry) error {
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

func (m *MessageMediaParticipant) validateToAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.ToAddress) { // not required
		return nil
	}

	if m.ToAddress != nil {
		if err := m.ToAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("toAddress")
			}
			return err
		}
	}

	return nil
}

var messageMediaParticipantTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","sms","twitter","facebook","line","whatsapp","telegram","kakao","webmessaging","open","instagram"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageMediaParticipantTypeTypePropEnum = append(messageMediaParticipantTypeTypePropEnum, v)
	}
}

const (

	// MessageMediaParticipantTypeUnknown captures enum value "unknown"
	MessageMediaParticipantTypeUnknown string = "unknown"

	// MessageMediaParticipantTypeSms captures enum value "sms"
	MessageMediaParticipantTypeSms string = "sms"

	// MessageMediaParticipantTypeTwitter captures enum value "twitter"
	MessageMediaParticipantTypeTwitter string = "twitter"

	// MessageMediaParticipantTypeFacebook captures enum value "facebook"
	MessageMediaParticipantTypeFacebook string = "facebook"

	// MessageMediaParticipantTypeLine captures enum value "line"
	MessageMediaParticipantTypeLine string = "line"

	// MessageMediaParticipantTypeWhatsapp captures enum value "whatsapp"
	MessageMediaParticipantTypeWhatsapp string = "whatsapp"

	// MessageMediaParticipantTypeTelegram captures enum value "telegram"
	MessageMediaParticipantTypeTelegram string = "telegram"

	// MessageMediaParticipantTypeKakao captures enum value "kakao"
	MessageMediaParticipantTypeKakao string = "kakao"

	// MessageMediaParticipantTypeWebmessaging captures enum value "webmessaging"
	MessageMediaParticipantTypeWebmessaging string = "webmessaging"

	// MessageMediaParticipantTypeOpen captures enum value "open"
	MessageMediaParticipantTypeOpen string = "open"

	// MessageMediaParticipantTypeInstagram captures enum value "instagram"
	MessageMediaParticipantTypeInstagram string = "instagram"
)

// prop value enum
func (m *MessageMediaParticipant) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageMediaParticipantTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageMediaParticipant) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *MessageMediaParticipant) validateUser(formats strfmt.Registry) error {
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

func (m *MessageMediaParticipant) validateWrapup(formats strfmt.Registry) error {
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

// ContextValidate validate this message media participant based on the context it is used
func (m *MessageMediaParticipant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateFromAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJourneyContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
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

	if err := m.contextValidateToAddress(ctx, formats); err != nil {
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

func (m *MessageMediaParticipant) contextValidateConversationRoutingData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateErrorInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateExternalContact(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateExternalOrganization(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateFromAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.FromAddress != nil {
		if err := m.FromAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fromAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MessageMediaParticipant) contextValidateJourneyContext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Messages); i++ {

		if m.Messages[i] != nil {
			if err := m.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessageMediaParticipant) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateScript(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateToAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.ToAddress != nil {
		if err := m.ToAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("toAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MessageMediaParticipant) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MessageMediaParticipant) contextValidateWrapup(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MessageMediaParticipant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageMediaParticipant) UnmarshalBinary(b []byte) error {
	var res MessageMediaParticipant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
