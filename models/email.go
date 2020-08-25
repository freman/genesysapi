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

// Email email
//
// swagger:model Email
type Email struct {

	// After-call work for the communication.
	AfterCallWork *AfterCallWork `json:"afterCallWork,omitempty"`

	// Indicates that the email was auto-generated like an Out of Office reply.
	AutoGenerated bool `json:"autoGenerated"`

	// The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// The direction of the email
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	// Enum: [endpoint client system timeout transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam uncallable]
	DisconnectType string `json:"disconnectType,omitempty"`

	// The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DisconnectedTime strfmt.DateTime `json:"disconnectedTime,omitempty"`

	// A list of uploaded attachments on the email draft.
	DraftAttachments []*Attachment `json:"draftAttachments"`

	// error info
	ErrorInfo *ErrorBody `json:"errorInfo,omitempty"`

	// True if this call is held and the person on this side hears silence.
	Held bool `json:"held"`

	// A globally unique identifier for this communication.
	ID string `json:"id,omitempty"`

	// A globally unique identifier for the stored content of this communication.
	MessageID string `json:"messageId,omitempty"`

	// The number of email messages sent by this participant.
	MessagesSent int32 `json:"messagesSent,omitempty"`

	// The id of the peer communication corresponding to a matching leg for this communication.
	PeerID string `json:"peerId,omitempty"`

	// The source provider for the email.
	Provider string `json:"provider,omitempty"`

	// A globally unique identifier for the recording associated with this call.
	RecordingID string `json:"recordingId,omitempty"`

	// The UUID of the script to use.
	ScriptID string `json:"scriptId,omitempty"`

	// The time line of the participant's email, divided into activity segments.
	Segments []*Segment `json:"segments"`

	// Indicates if the inbound email was marked as spam.
	Spam bool `json:"spam"`

	// The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	StartAlertingTime strfmt.DateTime `json:"startAlertingTime,omitempty"`

	// The timestamp the email was placed on hold in the cloud clock if the email is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	StartHoldTime strfmt.DateTime `json:"startHoldTime,omitempty"`

	// The connection state of this communication.
	// Enum: [alerting connected disconnected none transmitting]
	State string `json:"state,omitempty"`

	// The subject for the initial email that started this conversation.
	Subject string `json:"subject,omitempty"`

	// Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`
}

// Validate validates this email
func (m *Email) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDraftAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorInfo(formats); err != nil {
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

	if err := m.validateWrapup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Email) validateAfterCallWork(formats strfmt.Registry) error {

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

func (m *Email) validateConnectedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var emailTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emailTypeDirectionPropEnum = append(emailTypeDirectionPropEnum, v)
	}
}

const (

	// EmailDirectionInbound captures enum value "inbound"
	EmailDirectionInbound string = "inbound"

	// EmailDirectionOutbound captures enum value "outbound"
	EmailDirectionOutbound string = "outbound"
)

// prop value enum
func (m *Email) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emailTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Email) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var emailTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","timeout","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam","uncallable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emailTypeDisconnectTypePropEnum = append(emailTypeDisconnectTypePropEnum, v)
	}
}

const (

	// EmailDisconnectTypeEndpoint captures enum value "endpoint"
	EmailDisconnectTypeEndpoint string = "endpoint"

	// EmailDisconnectTypeClient captures enum value "client"
	EmailDisconnectTypeClient string = "client"

	// EmailDisconnectTypeSystem captures enum value "system"
	EmailDisconnectTypeSystem string = "system"

	// EmailDisconnectTypeTimeout captures enum value "timeout"
	EmailDisconnectTypeTimeout string = "timeout"

	// EmailDisconnectTypeTransfer captures enum value "transfer"
	EmailDisconnectTypeTransfer string = "transfer"

	// EmailDisconnectTypeTransferConference captures enum value "transfer.conference"
	EmailDisconnectTypeTransferConference string = "transfer.conference"

	// EmailDisconnectTypeTransferConsult captures enum value "transfer.consult"
	EmailDisconnectTypeTransferConsult string = "transfer.consult"

	// EmailDisconnectTypeTransferForward captures enum value "transfer.forward"
	EmailDisconnectTypeTransferForward string = "transfer.forward"

	// EmailDisconnectTypeTransferNoanswer captures enum value "transfer.noanswer"
	EmailDisconnectTypeTransferNoanswer string = "transfer.noanswer"

	// EmailDisconnectTypeTransferNotavailable captures enum value "transfer.notavailable"
	EmailDisconnectTypeTransferNotavailable string = "transfer.notavailable"

	// EmailDisconnectTypeTransportFailure captures enum value "transport.failure"
	EmailDisconnectTypeTransportFailure string = "transport.failure"

	// EmailDisconnectTypeError captures enum value "error"
	EmailDisconnectTypeError string = "error"

	// EmailDisconnectTypePeer captures enum value "peer"
	EmailDisconnectTypePeer string = "peer"

	// EmailDisconnectTypeOther captures enum value "other"
	EmailDisconnectTypeOther string = "other"

	// EmailDisconnectTypeSpam captures enum value "spam"
	EmailDisconnectTypeSpam string = "spam"

	// EmailDisconnectTypeUncallable captures enum value "uncallable"
	EmailDisconnectTypeUncallable string = "uncallable"
)

// prop value enum
func (m *Email) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emailTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Email) validateDisconnectType(formats strfmt.Registry) error {

	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *Email) validateDisconnectedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.DisconnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("disconnectedTime", "body", "date-time", m.DisconnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Email) validateDraftAttachments(formats strfmt.Registry) error {

	if swag.IsZero(m.DraftAttachments) { // not required
		return nil
	}

	for i := 0; i < len(m.DraftAttachments); i++ {
		if swag.IsZero(m.DraftAttachments[i]) { // not required
			continue
		}

		if m.DraftAttachments[i] != nil {
			if err := m.DraftAttachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("draftAttachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Email) validateErrorInfo(formats strfmt.Registry) error {

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

func (m *Email) validateSegments(formats strfmt.Registry) error {

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

func (m *Email) validateStartAlertingTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartAlertingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAlertingTime", "body", "date-time", m.StartAlertingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Email) validateStartHoldTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var emailTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","connected","disconnected","none","transmitting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emailTypeStatePropEnum = append(emailTypeStatePropEnum, v)
	}
}

const (

	// EmailStateAlerting captures enum value "alerting"
	EmailStateAlerting string = "alerting"

	// EmailStateConnected captures enum value "connected"
	EmailStateConnected string = "connected"

	// EmailStateDisconnected captures enum value "disconnected"
	EmailStateDisconnected string = "disconnected"

	// EmailStateNone captures enum value "none"
	EmailStateNone string = "none"

	// EmailStateTransmitting captures enum value "transmitting"
	EmailStateTransmitting string = "transmitting"
)

// prop value enum
func (m *Email) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emailTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Email) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Email) validateWrapup(formats strfmt.Registry) error {

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
func (m *Email) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Email) UnmarshalBinary(b []byte) error {
	var res Email
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}