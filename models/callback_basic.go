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

// CallbackBasic callback basic
//
// swagger:model CallbackBasic
type CallbackBasic struct {

	// After-call work for the communication.
	AfterCallWork *AfterCallWork `json:"afterCallWork,omitempty"`

	// Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

	// The id of the config for automatically placing the callback (and handling the disposition). If null, the callback will not be placed automatically but routed to an agent as per normal.
	AutomatedCallbackConfigID string `json:"automatedCallbackConfigId,omitempty"`

	// The phone number(s) to use to place the callback.
	CallbackNumbers []string `json:"callbackNumbers"`

	// The timestamp when this communication is scheduled in the provider clock. If this value is missing it indicates the callback will be placed immediately. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CallbackScheduledTime strfmt.DateTime `json:"callbackScheduledTime,omitempty"`

	// The name of the user requesting a callback.
	CallbackUserName string `json:"callbackUserName,omitempty"`

	// The phone number displayed to recipients of the phone call. The value should conform to the E164 format.
	CallerID string `json:"callerId,omitempty"`

	// The name displayed to recipients of the phone call.
	CallerIDName string `json:"callerIdName,omitempty"`

	// The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// The preview data to be used when this callback is a Preview.
	DialerPreview *DialerPreview `json:"dialerPreview,omitempty"`

	// The direction of the call
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	// Enum: [endpoint client system timeout transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam uncallable]
	DisconnectType string `json:"disconnectType,omitempty"`

	// The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DisconnectedTime strfmt.DateTime `json:"disconnectedTime,omitempty"`

	// True if the call for the callback uses external dialing.
	ExternalCampaign bool `json:"externalCampaign"`

	// True if this call is held and the person on this side hears silence.
	Held bool `json:"held"`

	// A globally unique identifier for this communication.
	ID string `json:"id,omitempty"`

	// The initial connection state of this communication.
	// Enum: [alerting dialing contacting offering connected disconnected terminated scheduled none]
	InitialState string `json:"initialState,omitempty"`

	// The id of the peer communication corresponding to a matching leg for this communication.
	PeerID string `json:"peerId,omitempty"`

	// The source provider for the callback.
	Provider string `json:"provider,omitempty"`

	// The UUID of the script to use.
	ScriptID string `json:"scriptId,omitempty"`

	// The time line of the participant's callback, divided into activity segments.
	Segments []*Segment `json:"segments"`

	// True if the ability to skip a callback should be enabled.
	SkipEnabled bool `json:"skipEnabled"`

	// The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartAlertingTime strfmt.DateTime `json:"startAlertingTime,omitempty"`

	// The timestamp the callback was placed on hold in the cloud clock if the callback is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartHoldTime strfmt.DateTime `json:"startHoldTime,omitempty"`

	// The connection state of this communication.
	// Enum: [alerting dialing contacting offering connected disconnected terminated scheduled none]
	State string `json:"state,omitempty"`

	// The number of seconds before the system automatically places a call for a callback.  0 means the automatic placement is disabled.
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`

	// The voicemail data to be used when this callback is an ACD voicemail.
	Voicemail *Voicemail `json:"voicemail,omitempty"`

	// Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`
}

// Validate validates this callback basic
func (m *CallbackBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAfterCallWork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallbackScheduledTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDialerPreview(formats); err != nil {
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

	if err := m.validateInitialState(formats); err != nil {
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

	if err := m.validateVoicemail(formats); err != nil {
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

func (m *CallbackBasic) validateAfterCallWork(formats strfmt.Registry) error {
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

func (m *CallbackBasic) validateCallbackScheduledTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CallbackScheduledTime) { // not required
		return nil
	}

	if err := validate.FormatOf("callbackScheduledTime", "body", "date-time", m.CallbackScheduledTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallbackBasic) validateConnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallbackBasic) validateDialerPreview(formats strfmt.Registry) error {
	if swag.IsZero(m.DialerPreview) { // not required
		return nil
	}

	if m.DialerPreview != nil {
		if err := m.DialerPreview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dialerPreview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dialerPreview")
			}
			return err
		}
	}

	return nil
}

var callbackBasicTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callbackBasicTypeDirectionPropEnum = append(callbackBasicTypeDirectionPropEnum, v)
	}
}

const (

	// CallbackBasicDirectionInbound captures enum value "inbound"
	CallbackBasicDirectionInbound string = "inbound"

	// CallbackBasicDirectionOutbound captures enum value "outbound"
	CallbackBasicDirectionOutbound string = "outbound"
)

// prop value enum
func (m *CallbackBasic) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callbackBasicTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallbackBasic) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var callbackBasicTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","timeout","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam","uncallable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callbackBasicTypeDisconnectTypePropEnum = append(callbackBasicTypeDisconnectTypePropEnum, v)
	}
}

const (

	// CallbackBasicDisconnectTypeEndpoint captures enum value "endpoint"
	CallbackBasicDisconnectTypeEndpoint string = "endpoint"

	// CallbackBasicDisconnectTypeClient captures enum value "client"
	CallbackBasicDisconnectTypeClient string = "client"

	// CallbackBasicDisconnectTypeSystem captures enum value "system"
	CallbackBasicDisconnectTypeSystem string = "system"

	// CallbackBasicDisconnectTypeTimeout captures enum value "timeout"
	CallbackBasicDisconnectTypeTimeout string = "timeout"

	// CallbackBasicDisconnectTypeTransfer captures enum value "transfer"
	CallbackBasicDisconnectTypeTransfer string = "transfer"

	// CallbackBasicDisconnectTypeTransferDotConference captures enum value "transfer.conference"
	CallbackBasicDisconnectTypeTransferDotConference string = "transfer.conference"

	// CallbackBasicDisconnectTypeTransferDotConsult captures enum value "transfer.consult"
	CallbackBasicDisconnectTypeTransferDotConsult string = "transfer.consult"

	// CallbackBasicDisconnectTypeTransferDotForward captures enum value "transfer.forward"
	CallbackBasicDisconnectTypeTransferDotForward string = "transfer.forward"

	// CallbackBasicDisconnectTypeTransferDotNoanswer captures enum value "transfer.noanswer"
	CallbackBasicDisconnectTypeTransferDotNoanswer string = "transfer.noanswer"

	// CallbackBasicDisconnectTypeTransferDotNotavailable captures enum value "transfer.notavailable"
	CallbackBasicDisconnectTypeTransferDotNotavailable string = "transfer.notavailable"

	// CallbackBasicDisconnectTypeTransportDotFailure captures enum value "transport.failure"
	CallbackBasicDisconnectTypeTransportDotFailure string = "transport.failure"

	// CallbackBasicDisconnectTypeError captures enum value "error"
	CallbackBasicDisconnectTypeError string = "error"

	// CallbackBasicDisconnectTypePeer captures enum value "peer"
	CallbackBasicDisconnectTypePeer string = "peer"

	// CallbackBasicDisconnectTypeOther captures enum value "other"
	CallbackBasicDisconnectTypeOther string = "other"

	// CallbackBasicDisconnectTypeSpam captures enum value "spam"
	CallbackBasicDisconnectTypeSpam string = "spam"

	// CallbackBasicDisconnectTypeUncallable captures enum value "uncallable"
	CallbackBasicDisconnectTypeUncallable string = "uncallable"
)

// prop value enum
func (m *CallbackBasic) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callbackBasicTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallbackBasic) validateDisconnectType(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *CallbackBasic) validateDisconnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("disconnectedTime", "body", "date-time", m.DisconnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var callbackBasicTypeInitialStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","scheduled","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callbackBasicTypeInitialStatePropEnum = append(callbackBasicTypeInitialStatePropEnum, v)
	}
}

const (

	// CallbackBasicInitialStateAlerting captures enum value "alerting"
	CallbackBasicInitialStateAlerting string = "alerting"

	// CallbackBasicInitialStateDialing captures enum value "dialing"
	CallbackBasicInitialStateDialing string = "dialing"

	// CallbackBasicInitialStateContacting captures enum value "contacting"
	CallbackBasicInitialStateContacting string = "contacting"

	// CallbackBasicInitialStateOffering captures enum value "offering"
	CallbackBasicInitialStateOffering string = "offering"

	// CallbackBasicInitialStateConnected captures enum value "connected"
	CallbackBasicInitialStateConnected string = "connected"

	// CallbackBasicInitialStateDisconnected captures enum value "disconnected"
	CallbackBasicInitialStateDisconnected string = "disconnected"

	// CallbackBasicInitialStateTerminated captures enum value "terminated"
	CallbackBasicInitialStateTerminated string = "terminated"

	// CallbackBasicInitialStateScheduled captures enum value "scheduled"
	CallbackBasicInitialStateScheduled string = "scheduled"

	// CallbackBasicInitialStateNone captures enum value "none"
	CallbackBasicInitialStateNone string = "none"
)

// prop value enum
func (m *CallbackBasic) validateInitialStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callbackBasicTypeInitialStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallbackBasic) validateInitialState(formats strfmt.Registry) error {
	if swag.IsZero(m.InitialState) { // not required
		return nil
	}

	// value enum
	if err := m.validateInitialStateEnum("initialState", "body", m.InitialState); err != nil {
		return err
	}

	return nil
}

func (m *CallbackBasic) validateSegments(formats strfmt.Registry) error {
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

func (m *CallbackBasic) validateStartAlertingTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAlertingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAlertingTime", "body", "date-time", m.StartAlertingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CallbackBasic) validateStartHoldTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var callbackBasicTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","scheduled","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callbackBasicTypeStatePropEnum = append(callbackBasicTypeStatePropEnum, v)
	}
}

const (

	// CallbackBasicStateAlerting captures enum value "alerting"
	CallbackBasicStateAlerting string = "alerting"

	// CallbackBasicStateDialing captures enum value "dialing"
	CallbackBasicStateDialing string = "dialing"

	// CallbackBasicStateContacting captures enum value "contacting"
	CallbackBasicStateContacting string = "contacting"

	// CallbackBasicStateOffering captures enum value "offering"
	CallbackBasicStateOffering string = "offering"

	// CallbackBasicStateConnected captures enum value "connected"
	CallbackBasicStateConnected string = "connected"

	// CallbackBasicStateDisconnected captures enum value "disconnected"
	CallbackBasicStateDisconnected string = "disconnected"

	// CallbackBasicStateTerminated captures enum value "terminated"
	CallbackBasicStateTerminated string = "terminated"

	// CallbackBasicStateScheduled captures enum value "scheduled"
	CallbackBasicStateScheduled string = "scheduled"

	// CallbackBasicStateNone captures enum value "none"
	CallbackBasicStateNone string = "none"
)

// prop value enum
func (m *CallbackBasic) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callbackBasicTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallbackBasic) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *CallbackBasic) validateVoicemail(formats strfmt.Registry) error {
	if swag.IsZero(m.Voicemail) { // not required
		return nil
	}

	if m.Voicemail != nil {
		if err := m.Voicemail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voicemail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("voicemail")
			}
			return err
		}
	}

	return nil
}

func (m *CallbackBasic) validateWrapup(formats strfmt.Registry) error {
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

// ContextValidate validate this callback basic based on the context it is used
func (m *CallbackBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAfterCallWork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDialerPreview(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSegments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVoicemail(ctx, formats); err != nil {
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

func (m *CallbackBasic) contextValidateAfterCallWork(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CallbackBasic) contextValidateDialerPreview(ctx context.Context, formats strfmt.Registry) error {

	if m.DialerPreview != nil {
		if err := m.DialerPreview.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dialerPreview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dialerPreview")
			}
			return err
		}
	}

	return nil
}

func (m *CallbackBasic) contextValidateSegments(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CallbackBasic) contextValidateVoicemail(ctx context.Context, formats strfmt.Registry) error {

	if m.Voicemail != nil {
		if err := m.Voicemail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voicemail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("voicemail")
			}
			return err
		}
	}

	return nil
}

func (m *CallbackBasic) contextValidateWrapup(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CallbackBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallbackBasic) UnmarshalBinary(b []byte) error {
	var res CallbackBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
