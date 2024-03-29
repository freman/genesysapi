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

// Call call
//
// swagger:model Call
type Call struct {

	// After-call work for the communication.
	AfterCallWork *AfterCallWork `json:"afterCallWork,omitempty"`

	// Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

	// UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
	AgentAssistantID string `json:"agentAssistantId,omitempty"`

	// True if this call is held and the person on this side hears hold music.
	Confined bool `json:"confined"`

	// The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// The direction of the call
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// List of reasons that this call was disconnected. This will be set once the call disconnects.
	DisconnectReasons []*DisconnectReason `json:"disconnectReasons"`

	// System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	// Enum: [endpoint client system timeout transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam uncallable]
	DisconnectType string `json:"disconnectType,omitempty"`

	// The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DisconnectedTime strfmt.DateTime `json:"disconnectedTime,omitempty"`

	// If call is an outbound fax of a document from content management, then this is the id in content management.
	DocumentID string `json:"documentId,omitempty"`

	// error info
	ErrorInfo *ErrorInfo `json:"errorInfo,omitempty"`

	// Extra information on fax transmission.
	FaxStatus *FaxStatus `json:"faxStatus,omitempty"`

	// True if this call is held and the person on this side hears silence.
	Held bool `json:"held"`

	// A globally unique identifier for this communication.
	ID string `json:"id,omitempty"`

	// The initial connection state of this communication.
	// Enum: [alerting dialing contacting offering connected disconnected terminated converting uploading transmitting none]
	InitialState string `json:"initialState,omitempty"`

	// True if this call is muted so that remote participants can't hear any audio from this end.
	Muted bool `json:"muted"`

	// Address and name data for a call endpoint.
	Other *Address `json:"other,omitempty"`

	// The id of the peer communication corresponding to a matching leg for this communication.
	PeerID string `json:"peerId,omitempty"`

	// The source provider for the call.
	Provider string `json:"provider,omitempty"`

	// True if this call is being recorded.
	Recording bool `json:"recording"`

	// A globally unique identifier for the recording associated with this call.
	RecordingID string `json:"recordingId,omitempty"`

	// State of recording on this call.
	// Enum: [none active paused]
	RecordingState string `json:"recordingState,omitempty"`

	// The UUID of the script to use.
	ScriptID string `json:"scriptId,omitempty"`

	// The time line of the participant's call, divided into activity segments.
	Segments []*Segment `json:"segments"`

	// Address and name data for a call endpoint.
	Self *Address `json:"self,omitempty"`

	// The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartAlertingTime strfmt.DateTime `json:"startAlertingTime,omitempty"`

	// The timestamp the call was placed on hold in the cloud clock if the call is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartHoldTime strfmt.DateTime `json:"startHoldTime,omitempty"`

	// The connection state of this communication.
	// Enum: [alerting dialing contacting offering connected disconnected terminated converting uploading transmitting none]
	State string `json:"state,omitempty"`

	// User to User Information (UUI) data managed by SIP session application.
	UuiData string `json:"uuiData,omitempty"`

	// Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`
}

// Validate validates this call
func (m *Call) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDisconnectReasons(formats); err != nil {
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

	if err := m.validateFaxStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOther(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordingState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
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

func (m *Call) validateAfterCallWork(formats strfmt.Registry) error {
	if swag.IsZero(m.AfterCallWork) { // not required
		return nil
	}

	if m.AfterCallWork != nil {
		if err := m.AfterCallWork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("afterCallWork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("afterCallWork")
			}
			return err
		}
	}

	return nil
}

func (m *Call) validateConnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var callTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callTypeDirectionPropEnum = append(callTypeDirectionPropEnum, v)
	}
}

const (

	// CallDirectionInbound captures enum value "inbound"
	CallDirectionInbound string = "inbound"

	// CallDirectionOutbound captures enum value "outbound"
	CallDirectionOutbound string = "outbound"
)

// prop value enum
func (m *Call) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Call) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDisconnectReasons(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectReasons) { // not required
		return nil
	}

	for i := 0; i < len(m.DisconnectReasons); i++ {
		if swag.IsZero(m.DisconnectReasons[i]) { // not required
			continue
		}

		if m.DisconnectReasons[i] != nil {
			if err := m.DisconnectReasons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disconnectReasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disconnectReasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var callTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","timeout","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam","uncallable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callTypeDisconnectTypePropEnum = append(callTypeDisconnectTypePropEnum, v)
	}
}

const (

	// CallDisconnectTypeEndpoint captures enum value "endpoint"
	CallDisconnectTypeEndpoint string = "endpoint"

	// CallDisconnectTypeClient captures enum value "client"
	CallDisconnectTypeClient string = "client"

	// CallDisconnectTypeSystem captures enum value "system"
	CallDisconnectTypeSystem string = "system"

	// CallDisconnectTypeTimeout captures enum value "timeout"
	CallDisconnectTypeTimeout string = "timeout"

	// CallDisconnectTypeTransfer captures enum value "transfer"
	CallDisconnectTypeTransfer string = "transfer"

	// CallDisconnectTypeTransferDotConference captures enum value "transfer.conference"
	CallDisconnectTypeTransferDotConference string = "transfer.conference"

	// CallDisconnectTypeTransferDotConsult captures enum value "transfer.consult"
	CallDisconnectTypeTransferDotConsult string = "transfer.consult"

	// CallDisconnectTypeTransferDotForward captures enum value "transfer.forward"
	CallDisconnectTypeTransferDotForward string = "transfer.forward"

	// CallDisconnectTypeTransferDotNoanswer captures enum value "transfer.noanswer"
	CallDisconnectTypeTransferDotNoanswer string = "transfer.noanswer"

	// CallDisconnectTypeTransferDotNotavailable captures enum value "transfer.notavailable"
	CallDisconnectTypeTransferDotNotavailable string = "transfer.notavailable"

	// CallDisconnectTypeTransportDotFailure captures enum value "transport.failure"
	CallDisconnectTypeTransportDotFailure string = "transport.failure"

	// CallDisconnectTypeError captures enum value "error"
	CallDisconnectTypeError string = "error"

	// CallDisconnectTypePeer captures enum value "peer"
	CallDisconnectTypePeer string = "peer"

	// CallDisconnectTypeOther captures enum value "other"
	CallDisconnectTypeOther string = "other"

	// CallDisconnectTypeSpam captures enum value "spam"
	CallDisconnectTypeSpam string = "spam"

	// CallDisconnectTypeUncallable captures enum value "uncallable"
	CallDisconnectTypeUncallable string = "uncallable"
)

// prop value enum
func (m *Call) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Call) validateDisconnectType(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDisconnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("disconnectedTime", "body", "date-time", m.DisconnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateErrorInfo(formats strfmt.Registry) error {
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

func (m *Call) validateFaxStatus(formats strfmt.Registry) error {
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

var callTypeInitialStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","converting","uploading","transmitting","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callTypeInitialStatePropEnum = append(callTypeInitialStatePropEnum, v)
	}
}

const (

	// CallInitialStateAlerting captures enum value "alerting"
	CallInitialStateAlerting string = "alerting"

	// CallInitialStateDialing captures enum value "dialing"
	CallInitialStateDialing string = "dialing"

	// CallInitialStateContacting captures enum value "contacting"
	CallInitialStateContacting string = "contacting"

	// CallInitialStateOffering captures enum value "offering"
	CallInitialStateOffering string = "offering"

	// CallInitialStateConnected captures enum value "connected"
	CallInitialStateConnected string = "connected"

	// CallInitialStateDisconnected captures enum value "disconnected"
	CallInitialStateDisconnected string = "disconnected"

	// CallInitialStateTerminated captures enum value "terminated"
	CallInitialStateTerminated string = "terminated"

	// CallInitialStateConverting captures enum value "converting"
	CallInitialStateConverting string = "converting"

	// CallInitialStateUploading captures enum value "uploading"
	CallInitialStateUploading string = "uploading"

	// CallInitialStateTransmitting captures enum value "transmitting"
	CallInitialStateTransmitting string = "transmitting"

	// CallInitialStateNone captures enum value "none"
	CallInitialStateNone string = "none"
)

// prop value enum
func (m *Call) validateInitialStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callTypeInitialStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Call) validateInitialState(formats strfmt.Registry) error {
	if swag.IsZero(m.InitialState) { // not required
		return nil
	}

	// value enum
	if err := m.validateInitialStateEnum("initialState", "body", m.InitialState); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateOther(formats strfmt.Registry) error {
	if swag.IsZero(m.Other) { // not required
		return nil
	}

	if m.Other != nil {
		if err := m.Other.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("other")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("other")
			}
			return err
		}
	}

	return nil
}

var callTypeRecordingStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","active","paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callTypeRecordingStatePropEnum = append(callTypeRecordingStatePropEnum, v)
	}
}

const (

	// CallRecordingStateNone captures enum value "none"
	CallRecordingStateNone string = "none"

	// CallRecordingStateActive captures enum value "active"
	CallRecordingStateActive string = "active"

	// CallRecordingStatePaused captures enum value "paused"
	CallRecordingStatePaused string = "paused"
)

// prop value enum
func (m *Call) validateRecordingStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callTypeRecordingStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Call) validateRecordingState(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordingState) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecordingStateEnum("recordingState", "body", m.RecordingState); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSegments(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Call) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

func (m *Call) validateStartAlertingTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAlertingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAlertingTime", "body", "date-time", m.StartAlertingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateStartHoldTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var callTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","converting","uploading","transmitting","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callTypeStatePropEnum = append(callTypeStatePropEnum, v)
	}
}

const (

	// CallStateAlerting captures enum value "alerting"
	CallStateAlerting string = "alerting"

	// CallStateDialing captures enum value "dialing"
	CallStateDialing string = "dialing"

	// CallStateContacting captures enum value "contacting"
	CallStateContacting string = "contacting"

	// CallStateOffering captures enum value "offering"
	CallStateOffering string = "offering"

	// CallStateConnected captures enum value "connected"
	CallStateConnected string = "connected"

	// CallStateDisconnected captures enum value "disconnected"
	CallStateDisconnected string = "disconnected"

	// CallStateTerminated captures enum value "terminated"
	CallStateTerminated string = "terminated"

	// CallStateConverting captures enum value "converting"
	CallStateConverting string = "converting"

	// CallStateUploading captures enum value "uploading"
	CallStateUploading string = "uploading"

	// CallStateTransmitting captures enum value "transmitting"
	CallStateTransmitting string = "transmitting"

	// CallStateNone captures enum value "none"
	CallStateNone string = "none"
)

// prop value enum
func (m *Call) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Call) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateWrapup(formats strfmt.Registry) error {
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

// ContextValidate validate this call based on the context it is used
func (m *Call) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAfterCallWork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisconnectReasons(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFaxStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOther(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSegments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
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

func (m *Call) contextValidateAfterCallWork(ctx context.Context, formats strfmt.Registry) error {

	if m.AfterCallWork != nil {
		if err := m.AfterCallWork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("afterCallWork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("afterCallWork")
			}
			return err
		}
	}

	return nil
}

func (m *Call) contextValidateDisconnectReasons(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisconnectReasons); i++ {

		if m.DisconnectReasons[i] != nil {
			if err := m.DisconnectReasons[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disconnectReasons" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disconnectReasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Call) contextValidateErrorInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Call) contextValidateFaxStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Call) contextValidateOther(ctx context.Context, formats strfmt.Registry) error {

	if m.Other != nil {
		if err := m.Other.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("other")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("other")
			}
			return err
		}
	}

	return nil
}

func (m *Call) contextValidateSegments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Segments); i++ {

		if m.Segments[i] != nil {
			if err := m.Segments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Call) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("self")
			}
			return err
		}
	}

	return nil
}

func (m *Call) contextValidateWrapup(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Call) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Call) UnmarshalBinary(b []byte) error {
	var res Call
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
