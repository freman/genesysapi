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

// Video video
//
// swagger:model Video
type Video struct {

	// After-call work for the communication.
	AfterCallWork *AfterCallWork `json:"afterCallWork,omitempty"`

	// Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

	// Indicates whether this participant has muted their outgoing audio.
	AudioMuted bool `json:"audioMuted"`

	// The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// The room id context (xmpp jid) for the conference session.
	Context string `json:"context,omitempty"`

	// System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	// Enum: [endpoint client system timeout transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam uncallable]
	DisconnectType string `json:"disconnectType,omitempty"`

	// The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DisconnectedTime strfmt.DateTime `json:"disconnectedTime,omitempty"`

	// A globally unique identifier for this communication.
	ID string `json:"id,omitempty"`

	// List of media stream ids
	Msids []string `json:"msids"`

	// The number of peer participants from the perspective of the participant in the conference.
	PeerCount int32 `json:"peerCount,omitempty"`

	// The id of the peer communication corresponding to a matching leg for this communication.
	PeerID string `json:"peerId,omitempty"`

	// The source provider for the video.
	Provider string `json:"provider,omitempty"`

	// Address and name data for a call endpoint.
	Self *Address `json:"self,omitempty"`

	// Indicates whether this participant is sharing their screen to the session.
	SharingScreen bool `json:"sharingScreen"`

	// The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	StartAlertingTime strfmt.DateTime `json:"startAlertingTime,omitempty"`

	// The connection state of this communication.
	// Enum: [alerting dialing contacting offering connected disconnected terminated none]
	State string `json:"state,omitempty"`

	// Indicates whether this participant has muted/paused their outgoing video.
	VideoMuted bool `json:"videoMuted"`

	// Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`
}

// Validate validates this video
func (m *Video) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAfterCallWork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisconnectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisconnectedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAlertingTime(formats); err != nil {
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

func (m *Video) validateAfterCallWork(formats strfmt.Registry) error {

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

func (m *Video) validateConnectedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var videoTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","timeout","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam","uncallable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		videoTypeDisconnectTypePropEnum = append(videoTypeDisconnectTypePropEnum, v)
	}
}

const (

	// VideoDisconnectTypeEndpoint captures enum value "endpoint"
	VideoDisconnectTypeEndpoint string = "endpoint"

	// VideoDisconnectTypeClient captures enum value "client"
	VideoDisconnectTypeClient string = "client"

	// VideoDisconnectTypeSystem captures enum value "system"
	VideoDisconnectTypeSystem string = "system"

	// VideoDisconnectTypeTimeout captures enum value "timeout"
	VideoDisconnectTypeTimeout string = "timeout"

	// VideoDisconnectTypeTransfer captures enum value "transfer"
	VideoDisconnectTypeTransfer string = "transfer"

	// VideoDisconnectTypeTransferConference captures enum value "transfer.conference"
	VideoDisconnectTypeTransferConference string = "transfer.conference"

	// VideoDisconnectTypeTransferConsult captures enum value "transfer.consult"
	VideoDisconnectTypeTransferConsult string = "transfer.consult"

	// VideoDisconnectTypeTransferForward captures enum value "transfer.forward"
	VideoDisconnectTypeTransferForward string = "transfer.forward"

	// VideoDisconnectTypeTransferNoanswer captures enum value "transfer.noanswer"
	VideoDisconnectTypeTransferNoanswer string = "transfer.noanswer"

	// VideoDisconnectTypeTransferNotavailable captures enum value "transfer.notavailable"
	VideoDisconnectTypeTransferNotavailable string = "transfer.notavailable"

	// VideoDisconnectTypeTransportFailure captures enum value "transport.failure"
	VideoDisconnectTypeTransportFailure string = "transport.failure"

	// VideoDisconnectTypeError captures enum value "error"
	VideoDisconnectTypeError string = "error"

	// VideoDisconnectTypePeer captures enum value "peer"
	VideoDisconnectTypePeer string = "peer"

	// VideoDisconnectTypeOther captures enum value "other"
	VideoDisconnectTypeOther string = "other"

	// VideoDisconnectTypeSpam captures enum value "spam"
	VideoDisconnectTypeSpam string = "spam"

	// VideoDisconnectTypeUncallable captures enum value "uncallable"
	VideoDisconnectTypeUncallable string = "uncallable"
)

// prop value enum
func (m *Video) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, videoTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Video) validateDisconnectType(formats strfmt.Registry) error {

	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *Video) validateDisconnectedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.DisconnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("disconnectedTime", "body", "date-time", m.DisconnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Video) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

func (m *Video) validateStartAlertingTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartAlertingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAlertingTime", "body", "date-time", m.StartAlertingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var videoTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		videoTypeStatePropEnum = append(videoTypeStatePropEnum, v)
	}
}

const (

	// VideoStateAlerting captures enum value "alerting"
	VideoStateAlerting string = "alerting"

	// VideoStateDialing captures enum value "dialing"
	VideoStateDialing string = "dialing"

	// VideoStateContacting captures enum value "contacting"
	VideoStateContacting string = "contacting"

	// VideoStateOffering captures enum value "offering"
	VideoStateOffering string = "offering"

	// VideoStateConnected captures enum value "connected"
	VideoStateConnected string = "connected"

	// VideoStateDisconnected captures enum value "disconnected"
	VideoStateDisconnected string = "disconnected"

	// VideoStateTerminated captures enum value "terminated"
	VideoStateTerminated string = "terminated"

	// VideoStateNone captures enum value "none"
	VideoStateNone string = "none"
)

// prop value enum
func (m *Video) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, videoTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Video) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Video) validateWrapup(formats strfmt.Registry) error {

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
func (m *Video) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Video) UnmarshalBinary(b []byte) error {
	var res Video
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
