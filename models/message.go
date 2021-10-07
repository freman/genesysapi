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

// Message message
//
// swagger:model Message
type Message struct {

	// After-call work for the communication.
	AfterCallWork *AfterCallWork `json:"afterCallWork,omitempty"`

	// Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

	// UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
	AgentAssistantID string `json:"agentAssistantId,omitempty"`

	// If true, the participant member is authenticated.
	Authenticated bool `json:"authenticated"`

	// The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// The direction of the message.
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	// Enum: [endpoint client system timeout transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam uncallable]
	DisconnectType string `json:"disconnectType,omitempty"`

	// The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DisconnectedTime strfmt.DateTime `json:"disconnectedTime,omitempty"`

	// error info
	ErrorInfo *ErrorBody `json:"errorInfo,omitempty"`

	// Address and name data for a call endpoint.
	FromAddress *Address `json:"fromAddress,omitempty"`

	// True if this call is held and the person on this side hears silence.
	Held bool `json:"held"`

	// A globally unique identifier for this communication.
	ID string `json:"id,omitempty"`

	// A subset of the Journey System's data relevant to a part of a conversation (for external linkage and internal usage/context).
	JourneyContext *JourneyContext `json:"journeyContext,omitempty"`

	// The messages sent on this communication channel.
	Messages []*MessageDetails `json:"messages"`

	// The id of the peer communication corresponding to a matching leg for this communication.
	PeerID string `json:"peerId,omitempty"`

	// The source provider for the message.
	Provider string `json:"provider,omitempty"`

	// Indicates the country where the recipient is associated in ISO 3166-1 alpha-2 format.
	RecipientCountry string `json:"recipientCountry,omitempty"`

	// The type of the recipient. Eg: Provisioned phoneNumber is the recipient for sms message type.
	RecipientType string `json:"recipientType,omitempty"`

	// A globally unique identifier for the recording associated with this message.
	RecordingID string `json:"recordingId,omitempty"`

	// The UUID of the script to use.
	ScriptID string `json:"scriptId,omitempty"`

	// The time line of the participant's message, divided into activity segments.
	Segments []*Segment `json:"segments"`

	// The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartAlertingTime strfmt.DateTime `json:"startAlertingTime,omitempty"`

	// The timestamp the message was placed on hold in the cloud clock if the message is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartHoldTime strfmt.DateTime `json:"startHoldTime,omitempty"`

	// The connection state of this communication.
	// Enum: [alerting connected disconnected]
	State string `json:"state,omitempty"`

	// Address and name data for a call endpoint.
	ToAddress *Address `json:"toAddress,omitempty"`

	// Indicates the type of message platform from which the message originated.
	// Enum: [unknown sms twitter facebook line whatsapp telegram kakao webmessaging open]
	Type string `json:"type,omitempty"`

	// Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`
}

// Validate validates this message
func (m *Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAfterCallWork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisconnectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisconnectedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorInfo(formats); err != nil {
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

	if err := m.validateSegments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAlertingTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartHoldTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *Message) validateAfterCallWork(formats strfmt.Registry) error {

	if swag.IsZero(m.AfterCallWork) { // not required
		return nil
	}

	if m.AfterCallWork != nil {
		if err := m.AfterCallWork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("afterCallWork")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateConnectedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var messageTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageTypeDirectionPropEnum = append(messageTypeDirectionPropEnum, v)
	}
}

const (

	// MessageDirectionInbound captures enum value "inbound"
	MessageDirectionInbound string = "inbound"

	// MessageDirectionOutbound captures enum value "outbound"
	MessageDirectionOutbound string = "outbound"
)

// prop value enum
func (m *Message) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Message) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var messageTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","timeout","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam","uncallable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageTypeDisconnectTypePropEnum = append(messageTypeDisconnectTypePropEnum, v)
	}
}

const (

	// MessageDisconnectTypeEndpoint captures enum value "endpoint"
	MessageDisconnectTypeEndpoint string = "endpoint"

	// MessageDisconnectTypeClient captures enum value "client"
	MessageDisconnectTypeClient string = "client"

	// MessageDisconnectTypeSystem captures enum value "system"
	MessageDisconnectTypeSystem string = "system"

	// MessageDisconnectTypeTimeout captures enum value "timeout"
	MessageDisconnectTypeTimeout string = "timeout"

	// MessageDisconnectTypeTransfer captures enum value "transfer"
	MessageDisconnectTypeTransfer string = "transfer"

	// MessageDisconnectTypeTransferConference captures enum value "transfer.conference"
	MessageDisconnectTypeTransferConference string = "transfer.conference"

	// MessageDisconnectTypeTransferConsult captures enum value "transfer.consult"
	MessageDisconnectTypeTransferConsult string = "transfer.consult"

	// MessageDisconnectTypeTransferForward captures enum value "transfer.forward"
	MessageDisconnectTypeTransferForward string = "transfer.forward"

	// MessageDisconnectTypeTransferNoanswer captures enum value "transfer.noanswer"
	MessageDisconnectTypeTransferNoanswer string = "transfer.noanswer"

	// MessageDisconnectTypeTransferNotavailable captures enum value "transfer.notavailable"
	MessageDisconnectTypeTransferNotavailable string = "transfer.notavailable"

	// MessageDisconnectTypeTransportFailure captures enum value "transport.failure"
	MessageDisconnectTypeTransportFailure string = "transport.failure"

	// MessageDisconnectTypeError captures enum value "error"
	MessageDisconnectTypeError string = "error"

	// MessageDisconnectTypePeer captures enum value "peer"
	MessageDisconnectTypePeer string = "peer"

	// MessageDisconnectTypeOther captures enum value "other"
	MessageDisconnectTypeOther string = "other"

	// MessageDisconnectTypeSpam captures enum value "spam"
	MessageDisconnectTypeSpam string = "spam"

	// MessageDisconnectTypeUncallable captures enum value "uncallable"
	MessageDisconnectTypeUncallable string = "uncallable"
)

// prop value enum
func (m *Message) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Message) validateDisconnectType(formats strfmt.Registry) error {

	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *Message) validateDisconnectedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.DisconnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("disconnectedTime", "body", "date-time", m.DisconnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Message) validateErrorInfo(formats strfmt.Registry) error {

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

func (m *Message) validateFromAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.FromAddress) { // not required
		return nil
	}

	if m.FromAddress != nil {
		if err := m.FromAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromAddress")
			}
			return err
		}
	}

	return nil
}

func (m *Message) validateJourneyContext(formats strfmt.Registry) error {

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

func (m *Message) validateMessages(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *Message) validateSegments(formats strfmt.Registry) error {

	if swag.IsZero(m.Segments) { // not required
		return nil
	}

	for i := 0; i < len(m.Segments); i++ {
		if swag.IsZero(m.Segments[i]) { // not required
			continue
		}

		if m.Segments[i] != nil {
			if err := m.Segments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Message) validateStartAlertingTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartAlertingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAlertingTime", "body", "date-time", m.StartAlertingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Message) validateStartHoldTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var messageTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","connected","disconnected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageTypeStatePropEnum = append(messageTypeStatePropEnum, v)
	}
}

const (

	// MessageStateAlerting captures enum value "alerting"
	MessageStateAlerting string = "alerting"

	// MessageStateConnected captures enum value "connected"
	MessageStateConnected string = "connected"

	// MessageStateDisconnected captures enum value "disconnected"
	MessageStateDisconnected string = "disconnected"
)

// prop value enum
func (m *Message) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Message) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Message) validateToAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.ToAddress) { // not required
		return nil
	}

	if m.ToAddress != nil {
		if err := m.ToAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toAddress")
			}
			return err
		}
	}

	return nil
}

var messageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","sms","twitter","facebook","line","whatsapp","telegram","kakao","webmessaging","open"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageTypeTypePropEnum = append(messageTypeTypePropEnum, v)
	}
}

const (

	// MessageTypeUnknown captures enum value "unknown"
	MessageTypeUnknown string = "unknown"

	// MessageTypeSms captures enum value "sms"
	MessageTypeSms string = "sms"

	// MessageTypeTwitter captures enum value "twitter"
	MessageTypeTwitter string = "twitter"

	// MessageTypeFacebook captures enum value "facebook"
	MessageTypeFacebook string = "facebook"

	// MessageTypeLine captures enum value "line"
	MessageTypeLine string = "line"

	// MessageTypeWhatsapp captures enum value "whatsapp"
	MessageTypeWhatsapp string = "whatsapp"

	// MessageTypeTelegram captures enum value "telegram"
	MessageTypeTelegram string = "telegram"

	// MessageTypeKakao captures enum value "kakao"
	MessageTypeKakao string = "kakao"

	// MessageTypeWebmessaging captures enum value "webmessaging"
	MessageTypeWebmessaging string = "webmessaging"

	// MessageTypeOpen captures enum value "open"
	MessageTypeOpen string = "open"
)

// prop value enum
func (m *Message) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Message) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Message) validateWrapup(formats strfmt.Registry) error {

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
func (m *Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Message) UnmarshalBinary(b []byte) error {
	var res Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
