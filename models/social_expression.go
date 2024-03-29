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

// SocialExpression social expression
//
// swagger:model SocialExpression
type SocialExpression struct {

	// After-call work for the communication.
	AfterCallWork *AfterCallWork `json:"afterCallWork,omitempty"`

	// Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

	// The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	// Enum: [endpoint client system timeout transfer transfer.conference transfer.consult transfer.forward transfer.noanswer transfer.notavailable transport.failure error peer other spam uncallable]
	DisconnectType string `json:"disconnectType,omitempty"`

	// The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DisconnectedTime strfmt.DateTime `json:"disconnectedTime,omitempty"`

	// True if this call is held and the person on this side hears silence.
	Held bool `json:"held"`

	// A globally unique identifier for this communication.
	ID string `json:"id,omitempty"`

	// The id of the peer communication corresponding to a matching leg for this communication.
	PeerID string `json:"peerId,omitempty"`

	// The text preview of the communication contents
	PreviewText string `json:"previewText,omitempty"`

	// The source provider for the social expression.
	Provider string `json:"provider,omitempty"`

	// A globally unique identifier for the recording associated with this chat.
	RecordingID string `json:"recordingId,omitempty"`

	// The UUID of the script to use.
	ScriptID string `json:"scriptId,omitempty"`

	// The time line of the participant's chat, divided into activity segments.
	Segments []*Segment `json:"segments"`

	// The social network of the communication
	SocialMediaHub string `json:"socialMediaHub,omitempty"`

	// A globally unique identifier for the social media.
	SocialMediaID string `json:"socialMediaId,omitempty"`

	// The user name for the communication.
	SocialUserName string `json:"socialUserName,omitempty"`

	// The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartAlertingTime strfmt.DateTime `json:"startAlertingTime,omitempty"`

	// The timestamp the chat was placed on hold in the cloud clock if the chat is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartHoldTime strfmt.DateTime `json:"startHoldTime,omitempty"`

	// The connection state of this communication.
	// Enum: [alerting dialing contacting offering connected disconnected terminated none]
	State string `json:"state,omitempty"`

	// Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`
}

// Validate validates this social expression
func (m *SocialExpression) Validate(formats strfmt.Registry) error {
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

func (m *SocialExpression) validateAfterCallWork(formats strfmt.Registry) error {
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

func (m *SocialExpression) validateConnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var socialExpressionTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["endpoint","client","system","timeout","transfer","transfer.conference","transfer.consult","transfer.forward","transfer.noanswer","transfer.notavailable","transport.failure","error","peer","other","spam","uncallable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		socialExpressionTypeDisconnectTypePropEnum = append(socialExpressionTypeDisconnectTypePropEnum, v)
	}
}

const (

	// SocialExpressionDisconnectTypeEndpoint captures enum value "endpoint"
	SocialExpressionDisconnectTypeEndpoint string = "endpoint"

	// SocialExpressionDisconnectTypeClient captures enum value "client"
	SocialExpressionDisconnectTypeClient string = "client"

	// SocialExpressionDisconnectTypeSystem captures enum value "system"
	SocialExpressionDisconnectTypeSystem string = "system"

	// SocialExpressionDisconnectTypeTimeout captures enum value "timeout"
	SocialExpressionDisconnectTypeTimeout string = "timeout"

	// SocialExpressionDisconnectTypeTransfer captures enum value "transfer"
	SocialExpressionDisconnectTypeTransfer string = "transfer"

	// SocialExpressionDisconnectTypeTransferDotConference captures enum value "transfer.conference"
	SocialExpressionDisconnectTypeTransferDotConference string = "transfer.conference"

	// SocialExpressionDisconnectTypeTransferDotConsult captures enum value "transfer.consult"
	SocialExpressionDisconnectTypeTransferDotConsult string = "transfer.consult"

	// SocialExpressionDisconnectTypeTransferDotForward captures enum value "transfer.forward"
	SocialExpressionDisconnectTypeTransferDotForward string = "transfer.forward"

	// SocialExpressionDisconnectTypeTransferDotNoanswer captures enum value "transfer.noanswer"
	SocialExpressionDisconnectTypeTransferDotNoanswer string = "transfer.noanswer"

	// SocialExpressionDisconnectTypeTransferDotNotavailable captures enum value "transfer.notavailable"
	SocialExpressionDisconnectTypeTransferDotNotavailable string = "transfer.notavailable"

	// SocialExpressionDisconnectTypeTransportDotFailure captures enum value "transport.failure"
	SocialExpressionDisconnectTypeTransportDotFailure string = "transport.failure"

	// SocialExpressionDisconnectTypeError captures enum value "error"
	SocialExpressionDisconnectTypeError string = "error"

	// SocialExpressionDisconnectTypePeer captures enum value "peer"
	SocialExpressionDisconnectTypePeer string = "peer"

	// SocialExpressionDisconnectTypeOther captures enum value "other"
	SocialExpressionDisconnectTypeOther string = "other"

	// SocialExpressionDisconnectTypeSpam captures enum value "spam"
	SocialExpressionDisconnectTypeSpam string = "spam"

	// SocialExpressionDisconnectTypeUncallable captures enum value "uncallable"
	SocialExpressionDisconnectTypeUncallable string = "uncallable"
)

// prop value enum
func (m *SocialExpression) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, socialExpressionTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SocialExpression) validateDisconnectType(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *SocialExpression) validateDisconnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("disconnectedTime", "body", "date-time", m.DisconnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SocialExpression) validateSegments(formats strfmt.Registry) error {
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

func (m *SocialExpression) validateStartAlertingTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAlertingTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAlertingTime", "body", "date-time", m.StartAlertingTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SocialExpression) validateStartHoldTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartHoldTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startHoldTime", "body", "date-time", m.StartHoldTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var socialExpressionTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		socialExpressionTypeStatePropEnum = append(socialExpressionTypeStatePropEnum, v)
	}
}

const (

	// SocialExpressionStateAlerting captures enum value "alerting"
	SocialExpressionStateAlerting string = "alerting"

	// SocialExpressionStateDialing captures enum value "dialing"
	SocialExpressionStateDialing string = "dialing"

	// SocialExpressionStateContacting captures enum value "contacting"
	SocialExpressionStateContacting string = "contacting"

	// SocialExpressionStateOffering captures enum value "offering"
	SocialExpressionStateOffering string = "offering"

	// SocialExpressionStateConnected captures enum value "connected"
	SocialExpressionStateConnected string = "connected"

	// SocialExpressionStateDisconnected captures enum value "disconnected"
	SocialExpressionStateDisconnected string = "disconnected"

	// SocialExpressionStateTerminated captures enum value "terminated"
	SocialExpressionStateTerminated string = "terminated"

	// SocialExpressionStateNone captures enum value "none"
	SocialExpressionStateNone string = "none"
)

// prop value enum
func (m *SocialExpression) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, socialExpressionTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SocialExpression) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *SocialExpression) validateWrapup(formats strfmt.Registry) error {
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

// ContextValidate validate this social expression based on the context it is used
func (m *SocialExpression) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAfterCallWork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSegments(ctx, formats); err != nil {
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

func (m *SocialExpression) contextValidateAfterCallWork(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SocialExpression) contextValidateSegments(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SocialExpression) contextValidateWrapup(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SocialExpression) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SocialExpression) UnmarshalBinary(b []byte) error {
	var res SocialExpression
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
