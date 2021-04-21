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

// AnalyticsSession analytics session
//
// swagger:model AnalyticsSession
type AnalyticsSession struct {

	// ID(s) of Skill(s) that are active on the conversation
	ActiveSkillIds []string `json:"activeSkillIds"`

	// Marker for an agent that skipped after call work
	AcwSkipped bool `json:"acwSkipped"`

	// The address that initiated an action
	AddressFrom string `json:"addressFrom,omitempty"`

	// The email address for the participant on the other side of the email conversation
	AddressOther string `json:"addressOther,omitempty"`

	// The email address for the participant on this side of the email conversation
	AddressSelf string `json:"addressSelf,omitempty"`

	// The address receiving an action
	AddressTo string `json:"addressTo,omitempty"`

	// Unique identifier of the active virtual agent assistant
	AgentAssistantID string `json:"agentAssistantId,omitempty"`

	// Bullseye ring of the targeted agent
	AgentBullseyeRing int32 `json:"agentBullseyeRing,omitempty"`

	// Automatic Number Identification (caller's number)
	Ani string `json:"ani,omitempty"`

	// ID of the user that manually assigned a conversation
	AssignerID string `json:"assignerId,omitempty"`

	// Flag that indicates that the identity of the customer has been asserted as verified by the provider.
	Authenticated bool `json:"authenticated"`

	// Callback phone number(s)
	CallbackNumbers []string `json:"callbackNumbers"`

	// Scheduled callback date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CallbackScheduledTime strfmt.DateTime `json:"callbackScheduledTime,omitempty"`

	// The name of the user requesting a call back
	CallbackUserName string `json:"callbackUserName,omitempty"`

	// Describes side of the cobrowse (sharer or viewer)
	CobrowseRole string `json:"cobrowseRole,omitempty"`

	// A unique identifier for a PureCloud cobrowse room
	CobrowseRoomID string `json:"cobrowseRoomId,omitempty"`

	// The direction of the communication
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// (Dialer) Analyzer (for example speech.person)
	DispositionAnalyzer string `json:"dispositionAnalyzer,omitempty"`

	// (Dialer) Result of the analysis (for example disposition.classification.callable.machine)
	DispositionName string `json:"dispositionName,omitempty"`

	// Dialed number identification service (number dialed by the calling party)
	Dnis string `json:"dnis,omitempty"`

	// Unique identifier of the edge device
	EdgeID string `json:"edgeId,omitempty"`

	// IVR flow execution associated with this session
	Flow *AnalyticsFlow `json:"flow,omitempty"`

	// Type of flow in that occurred when entering ACD.
	FlowInType string `json:"flowInType,omitempty"`

	// Type of flow out that occurred when emitting tFlowOut.
	FlowOutType string `json:"flowOutType,omitempty"`

	// Identifier of the journey action.
	JourneyActionID string `json:"journeyActionId,omitempty"`

	// Identifier of the journey action map that triggered the action.
	JourneyActionMapID string `json:"journeyActionMapId,omitempty"`

	// Version of the journey action map that triggered the action.
	JourneyActionMapVersion int32 `json:"journeyActionMapVersion,omitempty"`

	// Primary identifier of the journey customer in the source where the activities originate from.
	JourneyCustomerID string `json:"journeyCustomerId,omitempty"`

	// Type of primary identifier of the journey customer (e.g. cookie).
	JourneyCustomerIDType string `json:"journeyCustomerIdType,omitempty"`

	// Unique identifier of the journey session.
	JourneyCustomerSessionID string `json:"journeyCustomerSessionId,omitempty"`

	// Type or category of journey sessions (e.g. web, ticket, delivery, atm).
	JourneyCustomerSessionIDType string `json:"journeyCustomerSessionIdType,omitempty"`

	// Media bridge ID for the conference session consistent across all participants
	MediaBridgeID string `json:"mediaBridgeId,omitempty"`

	// Count of any media (images, files, etc) included in this session
	MediaCount int32 `json:"mediaCount,omitempty"`

	// MediaEndpointStats associated with this session
	MediaEndpointStats []*AnalyticsMediaEndpointStat `json:"mediaEndpointStats"`

	// The session media type
	// Enum: [callback chat cobrowse email message screenshare video voice]
	MediaType string `json:"mediaType,omitempty"`

	// Message type for messaging services. E.g.: sms, facebook, twitter, line
	MessageType string `json:"messageType,omitempty"`

	// List of metrics for this session
	Metrics []*AnalyticsSessionMetric `json:"metrics"`

	// The participantId being monitored (if someone (e.g. an agent) is being monitored, this would be the ID of the participant that was monitored that would correspond to other participantIds present in the conversation)
	MonitoredParticipantID string `json:"monitoredParticipantId,omitempty"`

	// (Dialer) Unique identifier of the outbound campaign
	OutboundCampaignID string `json:"outboundCampaignId,omitempty"`

	// (Dialer) Unique identifier of the contact
	OutboundContactID string `json:"outboundContactId,omitempty"`

	// (Dialer) Unique identifier of the contact list that this contact belongs to
	OutboundContactListID string `json:"outboundContactListId,omitempty"`

	// This identifies pairs of related sessions on a conversation. E.g. an external session’s peerId will be the session that the call originally connected to, e.g. if an IVR was dialed, the IVR session, which will also have the external session’s ID as its peer. After that point, any transfers of that session to other internal components (acd, agent, etc.) will all spawn new sessions whose peerIds point back to that original external session.
	PeerID string `json:"peerId,omitempty"`

	// Proposed agents
	ProposedAgents []*AnalyticsProposedAgent `json:"proposedAgents"`

	// The original voice protocol call ID, e.g. a SIP call ID
	ProtocolCallID string `json:"protocolCallId,omitempty"`

	// The source provider for the communication.
	Provider string `json:"provider,omitempty"`

	// Flag determining if an audio recording was started or not
	Recording bool `json:"recording"`

	// Name, phone number, or email address of the remote party.
	Remote string `json:"remote,omitempty"`

	// Unique identifier for the remote party
	RemoteNameDisplayable string `json:"remoteNameDisplayable,omitempty"`

	// ID(s) of Skill(s) that have been removed by bullseye routing
	RemovedSkillIds []string `json:"removedSkillIds"`

	// Routing type(s) for requested/attempted routing methods.
	RequestedRoutings []string `json:"requestedRoutings"`

	// Unique identifier for the room
	RoomID string `json:"roomId,omitempty"`

	// Routing ring for bullseye or preferred agent routing
	RoutingRing int32 `json:"routingRing,omitempty"`

	// Direct ScreenShare address
	ScreenShareAddressSelf string `json:"screenShareAddressSelf,omitempty"`

	// A unique identifier for a PureCloud ScreenShare room
	ScreenShareRoomID string `json:"screenShareRoomId,omitempty"`

	// A unique identifier for a script
	ScriptID string `json:"scriptId,omitempty"`

	// List of segments for this session
	Segments []*AnalyticsConversationSegment `json:"segments"`

	// Selected agent ID
	SelectedAgentID string `json:"selectedAgentId,omitempty"`

	// Selected agent GPR rank
	SelectedAgentRank int32 `json:"selectedAgentRank,omitempty"`

	// Dialed number for the current session; this can be different from dnis, e.g. if the call was transferred
	SessionDnis string `json:"sessionDnis,omitempty"`

	// The unique identifier of this session
	SessionID string `json:"sessionId,omitempty"`

	// Flag determining if screenShare is started or not (true/false)
	SharingScreen bool `json:"sharingScreen"`

	// (Dialer) Whether the agent can skip the dialer contact
	SkipEnabled bool `json:"skipEnabled"`

	// The number of seconds before PureCloud begins the call for a call back (0 disables automatic calling)
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`

	// Complete routing method
	// Enum: [Bullseye Last Manual Predictive Preferred Standard]
	UsedRouting string `json:"usedRouting,omitempty"`

	// Direct Video address
	VideoAddressSelf string `json:"videoAddressSelf,omitempty"`

	// A unique identifier for a PureCloud video room
	VideoRoomID string `json:"videoRoomId,omitempty"`
}

// Validate validates this analytics session
func (m *AnalyticsSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallbackScheduledTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaEndpointStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProposedAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedRoutings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedRouting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsSession) validateCallbackScheduledTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CallbackScheduledTime) { // not required
		return nil
	}

	if err := validate.FormatOf("callbackScheduledTime", "body", "date-time", m.CallbackScheduledTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var analyticsSessionTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsSessionTypeDirectionPropEnum = append(analyticsSessionTypeDirectionPropEnum, v)
	}
}

const (

	// AnalyticsSessionDirectionInbound captures enum value "inbound"
	AnalyticsSessionDirectionInbound string = "inbound"

	// AnalyticsSessionDirectionOutbound captures enum value "outbound"
	AnalyticsSessionDirectionOutbound string = "outbound"
)

// prop value enum
func (m *AnalyticsSession) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsSessionTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsSession) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsSession) validateFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.Flow) { // not required
		return nil
	}

	if m.Flow != nil {
		if err := m.Flow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *AnalyticsSession) validateMediaEndpointStats(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaEndpointStats) { // not required
		return nil
	}

	for i := 0; i < len(m.MediaEndpointStats); i++ {
		if swag.IsZero(m.MediaEndpointStats[i]) { // not required
			continue
		}

		if m.MediaEndpointStats[i] != nil {
			if err := m.MediaEndpointStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mediaEndpointStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var analyticsSessionTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["callback","chat","cobrowse","email","message","screenshare","video","voice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsSessionTypeMediaTypePropEnum = append(analyticsSessionTypeMediaTypePropEnum, v)
	}
}

const (

	// AnalyticsSessionMediaTypeCallback captures enum value "callback"
	AnalyticsSessionMediaTypeCallback string = "callback"

	// AnalyticsSessionMediaTypeChat captures enum value "chat"
	AnalyticsSessionMediaTypeChat string = "chat"

	// AnalyticsSessionMediaTypeCobrowse captures enum value "cobrowse"
	AnalyticsSessionMediaTypeCobrowse string = "cobrowse"

	// AnalyticsSessionMediaTypeEmail captures enum value "email"
	AnalyticsSessionMediaTypeEmail string = "email"

	// AnalyticsSessionMediaTypeMessage captures enum value "message"
	AnalyticsSessionMediaTypeMessage string = "message"

	// AnalyticsSessionMediaTypeScreenshare captures enum value "screenshare"
	AnalyticsSessionMediaTypeScreenshare string = "screenshare"

	// AnalyticsSessionMediaTypeVideo captures enum value "video"
	AnalyticsSessionMediaTypeVideo string = "video"

	// AnalyticsSessionMediaTypeVoice captures enum value "voice"
	AnalyticsSessionMediaTypeVoice string = "voice"
)

// prop value enum
func (m *AnalyticsSession) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsSessionTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsSession) validateMediaType(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsSession) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {
		if swag.IsZero(m.Metrics[i]) { // not required
			continue
		}

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnalyticsSession) validateProposedAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.ProposedAgents) { // not required
		return nil
	}

	for i := 0; i < len(m.ProposedAgents); i++ {
		if swag.IsZero(m.ProposedAgents[i]) { // not required
			continue
		}

		if m.ProposedAgents[i] != nil {
			if err := m.ProposedAgents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proposedAgents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var analyticsSessionRequestedRoutingsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bullseye","Last","Manual","Predictive","Preferred","Standard"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsSessionRequestedRoutingsItemsEnum = append(analyticsSessionRequestedRoutingsItemsEnum, v)
	}
}

func (m *AnalyticsSession) validateRequestedRoutingsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsSessionRequestedRoutingsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsSession) validateRequestedRoutings(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedRoutings) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestedRoutings); i++ {

		// value enum
		if err := m.validateRequestedRoutingsItemsEnum("requestedRoutings"+"."+strconv.Itoa(i), "body", m.RequestedRoutings[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *AnalyticsSession) validateSegments(formats strfmt.Registry) error {

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

var analyticsSessionTypeUsedRoutingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bullseye","Last","Manual","Predictive","Preferred","Standard"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsSessionTypeUsedRoutingPropEnum = append(analyticsSessionTypeUsedRoutingPropEnum, v)
	}
}

const (

	// AnalyticsSessionUsedRoutingBullseye captures enum value "Bullseye"
	AnalyticsSessionUsedRoutingBullseye string = "Bullseye"

	// AnalyticsSessionUsedRoutingLast captures enum value "Last"
	AnalyticsSessionUsedRoutingLast string = "Last"

	// AnalyticsSessionUsedRoutingManual captures enum value "Manual"
	AnalyticsSessionUsedRoutingManual string = "Manual"

	// AnalyticsSessionUsedRoutingPredictive captures enum value "Predictive"
	AnalyticsSessionUsedRoutingPredictive string = "Predictive"

	// AnalyticsSessionUsedRoutingPreferred captures enum value "Preferred"
	AnalyticsSessionUsedRoutingPreferred string = "Preferred"

	// AnalyticsSessionUsedRoutingStandard captures enum value "Standard"
	AnalyticsSessionUsedRoutingStandard string = "Standard"
)

// prop value enum
func (m *AnalyticsSession) validateUsedRoutingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsSessionTypeUsedRoutingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsSession) validateUsedRouting(formats strfmt.Registry) error {

	if swag.IsZero(m.UsedRouting) { // not required
		return nil
	}

	// value enum
	if err := m.validateUsedRoutingEnum("usedRouting", "body", m.UsedRouting); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsSession) UnmarshalBinary(b []byte) error {
	var res AnalyticsSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
