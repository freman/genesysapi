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

// AnalyticsConversationSegment analytics conversation segment
//
// swagger:model AnalyticsConversationSegment
type AnalyticsConversationSegment struct {

	// Flag indicating if audio is muted or not (true/false)
	AudioMuted bool `json:"audioMuted"`

	// Indicates whether the segment was a conference
	Conference bool `json:"conference"`

	// The unique identifier of a new conversation when a conversation is ended for a conference
	DestinationConversationID string `json:"destinationConversationId,omitempty"`

	// The unique identifier of a new session when a session is ended for a conference
	DestinationSessionID string `json:"destinationSessionId,omitempty"`

	// The session disconnect type
	// Enum: [client conferenceTransfer consultTransfer endpoint error forwardTransfer noAnswerTransfer notAvailableTransfer other peer spam system timeout transfer transportFailure uncallable]
	DisconnectType string `json:"disconnectType,omitempty"`

	// A code corresponding to the error that occurred
	ErrorCode string `json:"errorCode,omitempty"`

	// Unique identifier for a PureCloud group
	GroupID string `json:"groupId,omitempty"`

	// Additional segment properties
	Properties []*AnalyticsProperty `json:"properties"`

	// Q.850 response code(s)
	Q850ResponseCodes []int64 `json:"q850ResponseCodes"`

	// Queue identifier
	QueueID string `json:"queueId,omitempty"`

	// Unique identifier for the language requested for an interaction
	RequestedLanguageID string `json:"requestedLanguageId,omitempty"`

	// Unique identifier(s) for skill(s) requested for an interaction
	RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`

	// Unique identifier(s) for agent(s) requested for an interaction
	RequestedRoutingUserIds []string `json:"requestedRoutingUserIds"`

	// Scored agents
	ScoredAgents []*AnalyticsScoredAgent `json:"scoredAgents"`

	// The end time of a segment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	SegmentEnd strfmt.DateTime `json:"segmentEnd,omitempty"`

	// The start time of a segment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	SegmentStart strfmt.DateTime `json:"segmentStart,omitempty"`

	// The activity that takes place in the segment, such as hold or interact
	// Enum: [alert callback coaching contacting converting delay dialing hold interact ivr monitoring scheduled sharing system transmitting unknown uploading voicemail wrapup]
	SegmentType string `json:"segmentType,omitempty"`

	// SIP response code(s)
	SipResponseCodes []int64 `json:"sipResponseCodes"`

	// The unique identifier of the previous conversation when a new conversation is created for a conference
	SourceConversationID string `json:"sourceConversationId,omitempty"`

	// The unique identifier of the previous session when a new session is created for a conference
	SourceSessionID string `json:"sourceSessionId,omitempty"`

	// The subject for the initial email that started this conversation
	Subject string `json:"subject,omitempty"`

	// Flag indicating if video is muted/paused or not (true/false)
	VideoMuted bool `json:"videoMuted"`

	// Wrap up code
	WrapUpCode string `json:"wrapUpCode,omitempty"`

	// Note entered by an agent during after-call work
	WrapUpNote string `json:"wrapUpNote,omitempty"`

	// Tag(s) assigned during after-call work
	WrapUpTags []string `json:"wrapUpTags"`
}

// Validate validates this analytics conversation segment
func (m *AnalyticsConversationSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisconnectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoredAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var analyticsConversationSegmentTypeDisconnectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["client","conferenceTransfer","consultTransfer","endpoint","error","forwardTransfer","noAnswerTransfer","notAvailableTransfer","other","peer","spam","system","timeout","transfer","transportFailure","uncallable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsConversationSegmentTypeDisconnectTypePropEnum = append(analyticsConversationSegmentTypeDisconnectTypePropEnum, v)
	}
}

const (

	// AnalyticsConversationSegmentDisconnectTypeClient captures enum value "client"
	AnalyticsConversationSegmentDisconnectTypeClient string = "client"

	// AnalyticsConversationSegmentDisconnectTypeConferenceTransfer captures enum value "conferenceTransfer"
	AnalyticsConversationSegmentDisconnectTypeConferenceTransfer string = "conferenceTransfer"

	// AnalyticsConversationSegmentDisconnectTypeConsultTransfer captures enum value "consultTransfer"
	AnalyticsConversationSegmentDisconnectTypeConsultTransfer string = "consultTransfer"

	// AnalyticsConversationSegmentDisconnectTypeEndpoint captures enum value "endpoint"
	AnalyticsConversationSegmentDisconnectTypeEndpoint string = "endpoint"

	// AnalyticsConversationSegmentDisconnectTypeError captures enum value "error"
	AnalyticsConversationSegmentDisconnectTypeError string = "error"

	// AnalyticsConversationSegmentDisconnectTypeForwardTransfer captures enum value "forwardTransfer"
	AnalyticsConversationSegmentDisconnectTypeForwardTransfer string = "forwardTransfer"

	// AnalyticsConversationSegmentDisconnectTypeNoAnswerTransfer captures enum value "noAnswerTransfer"
	AnalyticsConversationSegmentDisconnectTypeNoAnswerTransfer string = "noAnswerTransfer"

	// AnalyticsConversationSegmentDisconnectTypeNotAvailableTransfer captures enum value "notAvailableTransfer"
	AnalyticsConversationSegmentDisconnectTypeNotAvailableTransfer string = "notAvailableTransfer"

	// AnalyticsConversationSegmentDisconnectTypeOther captures enum value "other"
	AnalyticsConversationSegmentDisconnectTypeOther string = "other"

	// AnalyticsConversationSegmentDisconnectTypePeer captures enum value "peer"
	AnalyticsConversationSegmentDisconnectTypePeer string = "peer"

	// AnalyticsConversationSegmentDisconnectTypeSpam captures enum value "spam"
	AnalyticsConversationSegmentDisconnectTypeSpam string = "spam"

	// AnalyticsConversationSegmentDisconnectTypeSystem captures enum value "system"
	AnalyticsConversationSegmentDisconnectTypeSystem string = "system"

	// AnalyticsConversationSegmentDisconnectTypeTimeout captures enum value "timeout"
	AnalyticsConversationSegmentDisconnectTypeTimeout string = "timeout"

	// AnalyticsConversationSegmentDisconnectTypeTransfer captures enum value "transfer"
	AnalyticsConversationSegmentDisconnectTypeTransfer string = "transfer"

	// AnalyticsConversationSegmentDisconnectTypeTransportFailure captures enum value "transportFailure"
	AnalyticsConversationSegmentDisconnectTypeTransportFailure string = "transportFailure"

	// AnalyticsConversationSegmentDisconnectTypeUncallable captures enum value "uncallable"
	AnalyticsConversationSegmentDisconnectTypeUncallable string = "uncallable"
)

// prop value enum
func (m *AnalyticsConversationSegment) validateDisconnectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsConversationSegmentTypeDisconnectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsConversationSegment) validateDisconnectType(formats strfmt.Registry) error {

	if swag.IsZero(m.DisconnectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisconnectTypeEnum("disconnectType", "body", m.DisconnectType); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsConversationSegment) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for i := 0; i < len(m.Properties); i++ {
		if swag.IsZero(m.Properties[i]) { // not required
			continue
		}

		if m.Properties[i] != nil {
			if err := m.Properties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnalyticsConversationSegment) validateScoredAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.ScoredAgents) { // not required
		return nil
	}

	for i := 0; i < len(m.ScoredAgents); i++ {
		if swag.IsZero(m.ScoredAgents[i]) { // not required
			continue
		}

		if m.ScoredAgents[i] != nil {
			if err := m.ScoredAgents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnalyticsConversationSegment) validateSegmentEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.SegmentEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("segmentEnd", "body", "date-time", m.SegmentEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsConversationSegment) validateSegmentStart(formats strfmt.Registry) error {

	if swag.IsZero(m.SegmentStart) { // not required
		return nil
	}

	if err := validate.FormatOf("segmentStart", "body", "date-time", m.SegmentStart.String(), formats); err != nil {
		return err
	}

	return nil
}

var analyticsConversationSegmentTypeSegmentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alert","callback","coaching","contacting","converting","delay","dialing","hold","interact","ivr","monitoring","scheduled","sharing","system","transmitting","unknown","uploading","voicemail","wrapup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsConversationSegmentTypeSegmentTypePropEnum = append(analyticsConversationSegmentTypeSegmentTypePropEnum, v)
	}
}

const (

	// AnalyticsConversationSegmentSegmentTypeAlert captures enum value "alert"
	AnalyticsConversationSegmentSegmentTypeAlert string = "alert"

	// AnalyticsConversationSegmentSegmentTypeCallback captures enum value "callback"
	AnalyticsConversationSegmentSegmentTypeCallback string = "callback"

	// AnalyticsConversationSegmentSegmentTypeCoaching captures enum value "coaching"
	AnalyticsConversationSegmentSegmentTypeCoaching string = "coaching"

	// AnalyticsConversationSegmentSegmentTypeContacting captures enum value "contacting"
	AnalyticsConversationSegmentSegmentTypeContacting string = "contacting"

	// AnalyticsConversationSegmentSegmentTypeConverting captures enum value "converting"
	AnalyticsConversationSegmentSegmentTypeConverting string = "converting"

	// AnalyticsConversationSegmentSegmentTypeDelay captures enum value "delay"
	AnalyticsConversationSegmentSegmentTypeDelay string = "delay"

	// AnalyticsConversationSegmentSegmentTypeDialing captures enum value "dialing"
	AnalyticsConversationSegmentSegmentTypeDialing string = "dialing"

	// AnalyticsConversationSegmentSegmentTypeHold captures enum value "hold"
	AnalyticsConversationSegmentSegmentTypeHold string = "hold"

	// AnalyticsConversationSegmentSegmentTypeInteract captures enum value "interact"
	AnalyticsConversationSegmentSegmentTypeInteract string = "interact"

	// AnalyticsConversationSegmentSegmentTypeIvr captures enum value "ivr"
	AnalyticsConversationSegmentSegmentTypeIvr string = "ivr"

	// AnalyticsConversationSegmentSegmentTypeMonitoring captures enum value "monitoring"
	AnalyticsConversationSegmentSegmentTypeMonitoring string = "monitoring"

	// AnalyticsConversationSegmentSegmentTypeScheduled captures enum value "scheduled"
	AnalyticsConversationSegmentSegmentTypeScheduled string = "scheduled"

	// AnalyticsConversationSegmentSegmentTypeSharing captures enum value "sharing"
	AnalyticsConversationSegmentSegmentTypeSharing string = "sharing"

	// AnalyticsConversationSegmentSegmentTypeSystem captures enum value "system"
	AnalyticsConversationSegmentSegmentTypeSystem string = "system"

	// AnalyticsConversationSegmentSegmentTypeTransmitting captures enum value "transmitting"
	AnalyticsConversationSegmentSegmentTypeTransmitting string = "transmitting"

	// AnalyticsConversationSegmentSegmentTypeUnknown captures enum value "unknown"
	AnalyticsConversationSegmentSegmentTypeUnknown string = "unknown"

	// AnalyticsConversationSegmentSegmentTypeUploading captures enum value "uploading"
	AnalyticsConversationSegmentSegmentTypeUploading string = "uploading"

	// AnalyticsConversationSegmentSegmentTypeVoicemail captures enum value "voicemail"
	AnalyticsConversationSegmentSegmentTypeVoicemail string = "voicemail"

	// AnalyticsConversationSegmentSegmentTypeWrapup captures enum value "wrapup"
	AnalyticsConversationSegmentSegmentTypeWrapup string = "wrapup"
)

// prop value enum
func (m *AnalyticsConversationSegment) validateSegmentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsConversationSegmentTypeSegmentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsConversationSegment) validateSegmentType(formats strfmt.Registry) error {

	if swag.IsZero(m.SegmentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSegmentTypeEnum("segmentType", "body", m.SegmentType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsConversationSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsConversationSegment) UnmarshalBinary(b []byte) error {
	var res AnalyticsConversationSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
