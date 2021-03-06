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

	// Marker for an agent that skipped after call work
	AcwSkipped bool `json:"acwSkipped"`

	// address from
	AddressFrom string `json:"addressFrom,omitempty"`

	// address other
	AddressOther string `json:"addressOther,omitempty"`

	// address self
	AddressSelf string `json:"addressSelf,omitempty"`

	// address to
	AddressTo string `json:"addressTo,omitempty"`

	// Unique identifier of the active virtual agent assistant
	AgentAssistantID string `json:"agentAssistantId,omitempty"`

	// Automatic Number Identification (caller's number)
	Ani string `json:"ani,omitempty"`

	// ID of the user that manually assigned a conversation
	AssignerID string `json:"assignerId,omitempty"`

	// List of numbers to callback
	CallbackNumbers []string `json:"callbackNumbers"`

	// Scheduled callback date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CallbackScheduledTime strfmt.DateTime `json:"callbackScheduledTime,omitempty"`

	// The name of the user requesting a call back
	CallbackUserName string `json:"callbackUserName,omitempty"`

	// Describe side of the cobrowse (sharer or viewer)
	CobrowseRole string `json:"cobrowseRole,omitempty"`

	// A unique identifier for a PureCloud Cobrowse room.
	CobrowseRoomID string `json:"cobrowseRoomId,omitempty"`

	// Direction
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// (Dialer) Unique identifier of the contact list that this contact belongs to
	DispositionAnalyzer string `json:"dispositionAnalyzer,omitempty"`

	// (Dialer) Result of the analysis
	DispositionName string `json:"dispositionName,omitempty"`

	// Dialed number identification service (number dialed by the calling party)
	Dnis string `json:"dnis,omitempty"`

	// Unique identifier of the edge device
	EdgeID string `json:"edgeId,omitempty"`

	// IVR flow execution associated with this session
	Flow *AnalyticsFlow `json:"flow,omitempty"`

	// Type of flow in that occurred, e.g. acd, ivr, etc.
	FlowInType string `json:"flowInType,omitempty"`

	// Type of flow out that occurred, e.g. voicemail, callback, or acd
	FlowOutType string `json:"flowOutType,omitempty"`

	// Journey action ID
	JourneyActionID string `json:"journeyActionId,omitempty"`

	// Journey action map ID
	JourneyActionMapID string `json:"journeyActionMapId,omitempty"`

	// Journey action map version
	JourneyActionMapVersion string `json:"journeyActionMapVersion,omitempty"`

	// ID of the journey customer
	JourneyCustomerID string `json:"journeyCustomerId,omitempty"`

	// Type of the journey customer ID
	JourneyCustomerIDType string `json:"journeyCustomerIdType,omitempty"`

	// ID of the journey customer session
	JourneyCustomerSessionID string `json:"journeyCustomerSessionId,omitempty"`

	// Type of the journey customer session ID
	JourneyCustomerSessionIDType string `json:"journeyCustomerSessionIdType,omitempty"`

	// media bridge Id
	MediaBridgeID string `json:"mediaBridgeId,omitempty"`

	// Count of any media (images, files, etc) included in this session
	MediaCount int32 `json:"mediaCount,omitempty"`

	// Media endpoint stats associated with this session
	MediaEndpointStats []*AnalyticsMediaEndpointStat `json:"mediaEndpointStats"`

	// The session media type
	// Enum: [voice chat email callback cobrowse video screenshare message]
	MediaType string `json:"mediaType,omitempty"`

	// Message type for messaging services such as sms
	// Enum: [sms facebook twitter line]
	MessageType string `json:"messageType,omitempty"`

	// List of metrics for this session
	Metrics []*AnalyticsSessionMetric `json:"metrics"`

	// monitored participant Id
	MonitoredParticipantID string `json:"monitoredParticipantId,omitempty"`

	// The sessionID being monitored
	MonitoredSessionID string `json:"monitoredSessionId,omitempty"`

	// (Dialer) Unique identifier of the outbound campaign
	OutboundCampaignID string `json:"outboundCampaignId,omitempty"`

	// (Dialer) Unique identifier of the contact
	OutboundContactID string `json:"outboundContactId,omitempty"`

	// (Dialer) Unique identifier of the contact list that this contact belongs to
	OutboundContactListID string `json:"outboundContactListId,omitempty"`

	// A unique identifier for a peer
	PeerID string `json:"peerId,omitempty"`

	// Proposed agents
	ProposedAgents []*AnalyticsProposedAgent `json:"proposedAgents"`

	// The original voice protocol call ID, e.g. a SIP call ID
	ProtocolCallID string `json:"protocolCallId,omitempty"`

	// The source provider for the communication
	Provider string `json:"provider,omitempty"`

	// Flag determining if an audio recording was started or not
	Recording bool `json:"recording"`

	// Name, phone number, or email address of the remote party.
	Remote string `json:"remote,omitempty"`

	// remote name displayable
	RemoteNameDisplayable string `json:"remoteNameDisplayable,omitempty"`

	// All routing types for requested/attempted routing methods.
	RequestedRoutings []string `json:"requestedRoutings"`

	// Unique identifier for the room
	RoomID string `json:"roomId,omitempty"`

	// Direct ScreenShare address
	ScreenShareAddressSelf string `json:"screenShareAddressSelf,omitempty"`

	// A unique identifier for a PureCloud ScreenShare room.
	ScreenShareRoomID string `json:"screenShareRoomId,omitempty"`

	// A unique identifier for a script
	ScriptID string `json:"scriptId,omitempty"`

	// List of segments for this session
	Segments []*AnalyticsConversationSegment `json:"segments"`

	// Selected agent id
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

	// The number of seconds before PureCloud begins the call for a call back. 0 disables automatic calling
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`

	// Complete routing method
	// Enum: [Predictive Preferred Manual Last Bullseye Standard]
	UsedRouting string `json:"usedRouting,omitempty"`

	// Direct Video address
	VideoAddressSelf string `json:"videoAddressSelf,omitempty"`

	// A unique identifier for a PureCloud video room.
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

	if err := m.validateMessageType(formats); err != nil {
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
	if err := json.Unmarshal([]byte(`["voice","chat","email","callback","cobrowse","video","screenshare","message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsSessionTypeMediaTypePropEnum = append(analyticsSessionTypeMediaTypePropEnum, v)
	}
}

const (

	// AnalyticsSessionMediaTypeVoice captures enum value "voice"
	AnalyticsSessionMediaTypeVoice string = "voice"

	// AnalyticsSessionMediaTypeChat captures enum value "chat"
	AnalyticsSessionMediaTypeChat string = "chat"

	// AnalyticsSessionMediaTypeEmail captures enum value "email"
	AnalyticsSessionMediaTypeEmail string = "email"

	// AnalyticsSessionMediaTypeCallback captures enum value "callback"
	AnalyticsSessionMediaTypeCallback string = "callback"

	// AnalyticsSessionMediaTypeCobrowse captures enum value "cobrowse"
	AnalyticsSessionMediaTypeCobrowse string = "cobrowse"

	// AnalyticsSessionMediaTypeVideo captures enum value "video"
	AnalyticsSessionMediaTypeVideo string = "video"

	// AnalyticsSessionMediaTypeScreenshare captures enum value "screenshare"
	AnalyticsSessionMediaTypeScreenshare string = "screenshare"

	// AnalyticsSessionMediaTypeMessage captures enum value "message"
	AnalyticsSessionMediaTypeMessage string = "message"
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

var analyticsSessionTypeMessageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms","facebook","twitter","line"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsSessionTypeMessageTypePropEnum = append(analyticsSessionTypeMessageTypePropEnum, v)
	}
}

const (

	// AnalyticsSessionMessageTypeSms captures enum value "sms"
	AnalyticsSessionMessageTypeSms string = "sms"

	// AnalyticsSessionMessageTypeFacebook captures enum value "facebook"
	AnalyticsSessionMessageTypeFacebook string = "facebook"

	// AnalyticsSessionMessageTypeTwitter captures enum value "twitter"
	AnalyticsSessionMessageTypeTwitter string = "twitter"

	// AnalyticsSessionMessageTypeLine captures enum value "line"
	AnalyticsSessionMessageTypeLine string = "line"
)

// prop value enum
func (m *AnalyticsSession) validateMessageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsSessionTypeMessageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsSession) validateMessageType(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageTypeEnum("messageType", "body", m.MessageType); err != nil {
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
	if err := json.Unmarshal([]byte(`["Predictive","Preferred","Manual","Last","Bullseye","Standard"]`), &res); err != nil {
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
	if err := json.Unmarshal([]byte(`["Predictive","Preferred","Manual","Last","Bullseye","Standard"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsSessionTypeUsedRoutingPropEnum = append(analyticsSessionTypeUsedRoutingPropEnum, v)
	}
}

const (

	// AnalyticsSessionUsedRoutingPredictive captures enum value "Predictive"
	AnalyticsSessionUsedRoutingPredictive string = "Predictive"

	// AnalyticsSessionUsedRoutingPreferred captures enum value "Preferred"
	AnalyticsSessionUsedRoutingPreferred string = "Preferred"

	// AnalyticsSessionUsedRoutingManual captures enum value "Manual"
	AnalyticsSessionUsedRoutingManual string = "Manual"

	// AnalyticsSessionUsedRoutingLast captures enum value "Last"
	AnalyticsSessionUsedRoutingLast string = "Last"

	// AnalyticsSessionUsedRoutingBullseye captures enum value "Bullseye"
	AnalyticsSessionUsedRoutingBullseye string = "Bullseye"

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
