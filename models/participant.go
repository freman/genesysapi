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

// Participant participant
//
// swagger:model Participant
type Participant struct {

	// The address for the this participant. For a phone call this will be the ANI.
	Address string `json:"address,omitempty"`

	// Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs int32 `json:"alertingTimeoutMs,omitempty"`

	// The address for the this participant. For a phone call this will be the ANI.
	Ani string `json:"ani,omitempty"`

	// The ani-based name for this participant.
	AniName string `json:"aniName,omitempty"`

	// Additional participant attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// If this participant barged in a participant's call, then this will be the id of the targeted participant.
	BargedParticipantID string `json:"bargedParticipantId,omitempty"`

	// callbacks
	Callbacks []*Callback `json:"callbacks"`

	// calls
	Calls []*Call `json:"calls"`

	// chats
	Chats []*ConversationChat `json:"chats"`

	// If this participant is a coach, then this will be the id of the participant that is being coached.
	CoachedParticipantID string `json:"coachedParticipantId,omitempty"`

	// cobrowsesessions
	Cobrowsesessions []*Cobrowsesession `json:"cobrowsesessions"`

	// The timestamp when this participant was connected to the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConnectedTime strfmt.DateTime `json:"connectedTime,omitempty"`

	// If this participant is part of a consult transfer, then this will be the participant id of the participant being transferred.
	ConsultParticipantID string `json:"consultParticipantId,omitempty"`

	// Information on how a communication should be routed to an agent.
	ConversationRoutingData *ConversationRoutingData `json:"conversationRoutingData,omitempty"`

	// The address for the this participant. For a phone call this will be the ANI.
	Dnis string `json:"dnis,omitempty"`

	// emails
	Emails []*Email `json:"emails"`

	// The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndAcwTime strfmt.DateTime `json:"endAcwTime,omitempty"`

	// The timestamp when this participant disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// evaluations
	Evaluations []*Evaluation `json:"evaluations"`

	// If this participant represents an external contact, then this will be the globally unique identifier for the external contact.
	ExternalContactID string `json:"externalContactId,omitempty"`

	// If this participant represents an external org, then this will be the globally unique identifier for the external org.
	ExternalOrganizationID string `json:"externalOrganizationId,omitempty"`

	// The reason specifying why participant flagged the conversation.
	// Enum: [general]
	FlaggedReason string `json:"flaggedReason,omitempty"`

	// If present, group of users the participant represents.
	GroupID string `json:"groupId,omitempty"`

	// A globally unique identifier for this conversation.
	ID string `json:"id,omitempty"`

	// An ISO 639 language code specifying the locale for this participant
	Locale string `json:"locale,omitempty"`

	// messages
	Messages []*Message `json:"messages"`

	// If this participant is a monitor, then this will be the id of the participant that is being monitored.
	MonitoredParticipantID string `json:"monitoredParticipantId,omitempty"`

	// A human readable name identifying the participant.
	Name string `json:"name,omitempty"`

	// A well known string that specifies the type of this participant.
	ParticipantType string `json:"participantType,omitempty"`

	// A well known string that specifies the purpose of this participant.
	Purpose string `json:"purpose,omitempty"`

	// If present, the queue id that the communication channel came in on.
	QueueID string `json:"queueId,omitempty"`

	// If present, the queue name that the communication channel came in on.
	QueueName string `json:"queueName,omitempty"`

	// The current screen recording state for this participant.
	// Enum: [requested active paused stopped error timeout]
	ScreenRecordingState string `json:"screenRecordingState,omitempty"`

	// screenshares
	Screenshares []*Screenshare `json:"screenshares"`

	// social expressions
	SocialExpressions []*SocialExpression `json:"socialExpressions"`

	// The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartAcwTime strfmt.DateTime `json:"startAcwTime,omitempty"`

	// The timestamp when this participant joined the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// The team id that this participant is a member of when added to the conversation.
	TeamID string `json:"teamId,omitempty"`

	// If this participant represents a user, then this will be the globally unique identifier for the user.
	UserID string `json:"userId,omitempty"`

	// If this participant represents a user, then this will be an URI that can be used to fetch the user.
	UserURI string `json:"userUri,omitempty"`

	// videos
	Videos []*Video `json:"videos"`

	// Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`

	// This field controls how the UI prompts the agent for a wrapup.
	// Enum: [mandatory optional agentRequested timeout forcedTimeout]
	WrapupPrompt string `json:"wrapupPrompt,omitempty"`

	// True iff this participant is required to enter wrapup for this conversation.
	WrapupRequired bool `json:"wrapupRequired"`

	// The UI sets this field when the agent chooses to skip entering a wrapup for this participant.
	WrapupSkipped bool `json:"wrapupSkipped"`

	// Specifies how long a timed ACW session will last.
	WrapupTimeoutMs int32 `json:"wrapupTimeoutMs,omitempty"`
}

// Validate validates this participant
func (m *Participant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallbacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCobrowsesessions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationRoutingData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndAcwTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlaggedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScreenRecordingState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScreenshares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSocialExpressions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAcwTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVideos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrapup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrapupPrompt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Participant) validateCallbacks(formats strfmt.Registry) error {
	if swag.IsZero(m.Callbacks) { // not required
		return nil
	}

	for i := 0; i < len(m.Callbacks); i++ {
		if swag.IsZero(m.Callbacks[i]) { // not required
			continue
		}

		if m.Callbacks[i] != nil {
			if err := m.Callbacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("callbacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("callbacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) validateCalls(formats strfmt.Registry) error {
	if swag.IsZero(m.Calls) { // not required
		return nil
	}

	for i := 0; i < len(m.Calls); i++ {
		if swag.IsZero(m.Calls[i]) { // not required
			continue
		}

		if m.Calls[i] != nil {
			if err := m.Calls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("calls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("calls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) validateChats(formats strfmt.Registry) error {
	if swag.IsZero(m.Chats) { // not required
		return nil
	}

	for i := 0; i < len(m.Chats); i++ {
		if swag.IsZero(m.Chats[i]) { // not required
			continue
		}

		if m.Chats[i] != nil {
			if err := m.Chats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) validateCobrowsesessions(formats strfmt.Registry) error {
	if swag.IsZero(m.Cobrowsesessions) { // not required
		return nil
	}

	for i := 0; i < len(m.Cobrowsesessions); i++ {
		if swag.IsZero(m.Cobrowsesessions[i]) { // not required
			continue
		}

		if m.Cobrowsesessions[i] != nil {
			if err := m.Cobrowsesessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cobrowsesessions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cobrowsesessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) validateConnectedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedTime", "body", "date-time", m.ConnectedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Participant) validateConversationRoutingData(formats strfmt.Registry) error {
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

func (m *Participant) validateEmails(formats strfmt.Registry) error {
	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) validateEndAcwTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endAcwTime", "body", "date-time", m.EndAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Participant) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Participant) validateEvaluations(formats strfmt.Registry) error {
	if swag.IsZero(m.Evaluations) { // not required
		return nil
	}

	for i := 0; i < len(m.Evaluations); i++ {
		if swag.IsZero(m.Evaluations[i]) { // not required
			continue
		}

		if m.Evaluations[i] != nil {
			if err := m.Evaluations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evaluations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("evaluations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var participantTypeFlaggedReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		participantTypeFlaggedReasonPropEnum = append(participantTypeFlaggedReasonPropEnum, v)
	}
}

const (

	// ParticipantFlaggedReasonGeneral captures enum value "general"
	ParticipantFlaggedReasonGeneral string = "general"
)

// prop value enum
func (m *Participant) validateFlaggedReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, participantTypeFlaggedReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Participant) validateFlaggedReason(formats strfmt.Registry) error {
	if swag.IsZero(m.FlaggedReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlaggedReasonEnum("flaggedReason", "body", m.FlaggedReason); err != nil {
		return err
	}

	return nil
}

func (m *Participant) validateMessages(formats strfmt.Registry) error {
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

var participantTypeScreenRecordingStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["requested","active","paused","stopped","error","timeout"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		participantTypeScreenRecordingStatePropEnum = append(participantTypeScreenRecordingStatePropEnum, v)
	}
}

const (

	// ParticipantScreenRecordingStateRequested captures enum value "requested"
	ParticipantScreenRecordingStateRequested string = "requested"

	// ParticipantScreenRecordingStateActive captures enum value "active"
	ParticipantScreenRecordingStateActive string = "active"

	// ParticipantScreenRecordingStatePaused captures enum value "paused"
	ParticipantScreenRecordingStatePaused string = "paused"

	// ParticipantScreenRecordingStateStopped captures enum value "stopped"
	ParticipantScreenRecordingStateStopped string = "stopped"

	// ParticipantScreenRecordingStateError captures enum value "error"
	ParticipantScreenRecordingStateError string = "error"

	// ParticipantScreenRecordingStateTimeout captures enum value "timeout"
	ParticipantScreenRecordingStateTimeout string = "timeout"
)

// prop value enum
func (m *Participant) validateScreenRecordingStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, participantTypeScreenRecordingStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Participant) validateScreenRecordingState(formats strfmt.Registry) error {
	if swag.IsZero(m.ScreenRecordingState) { // not required
		return nil
	}

	// value enum
	if err := m.validateScreenRecordingStateEnum("screenRecordingState", "body", m.ScreenRecordingState); err != nil {
		return err
	}

	return nil
}

func (m *Participant) validateScreenshares(formats strfmt.Registry) error {
	if swag.IsZero(m.Screenshares) { // not required
		return nil
	}

	for i := 0; i < len(m.Screenshares); i++ {
		if swag.IsZero(m.Screenshares[i]) { // not required
			continue
		}

		if m.Screenshares[i] != nil {
			if err := m.Screenshares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("screenshares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("screenshares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) validateSocialExpressions(formats strfmt.Registry) error {
	if swag.IsZero(m.SocialExpressions) { // not required
		return nil
	}

	for i := 0; i < len(m.SocialExpressions); i++ {
		if swag.IsZero(m.SocialExpressions[i]) { // not required
			continue
		}

		if m.SocialExpressions[i] != nil {
			if err := m.SocialExpressions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("socialExpressions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("socialExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) validateStartAcwTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartAcwTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startAcwTime", "body", "date-time", m.StartAcwTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Participant) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Participant) validateVideos(formats strfmt.Registry) error {
	if swag.IsZero(m.Videos) { // not required
		return nil
	}

	for i := 0; i < len(m.Videos); i++ {
		if swag.IsZero(m.Videos[i]) { // not required
			continue
		}

		if m.Videos[i] != nil {
			if err := m.Videos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("videos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("videos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) validateWrapup(formats strfmt.Registry) error {
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

var participantTypeWrapupPromptPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mandatory","optional","agentRequested","timeout","forcedTimeout"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		participantTypeWrapupPromptPropEnum = append(participantTypeWrapupPromptPropEnum, v)
	}
}

const (

	// ParticipantWrapupPromptMandatory captures enum value "mandatory"
	ParticipantWrapupPromptMandatory string = "mandatory"

	// ParticipantWrapupPromptOptional captures enum value "optional"
	ParticipantWrapupPromptOptional string = "optional"

	// ParticipantWrapupPromptAgentRequested captures enum value "agentRequested"
	ParticipantWrapupPromptAgentRequested string = "agentRequested"

	// ParticipantWrapupPromptTimeout captures enum value "timeout"
	ParticipantWrapupPromptTimeout string = "timeout"

	// ParticipantWrapupPromptForcedTimeout captures enum value "forcedTimeout"
	ParticipantWrapupPromptForcedTimeout string = "forcedTimeout"
)

// prop value enum
func (m *Participant) validateWrapupPromptEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, participantTypeWrapupPromptPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Participant) validateWrapupPrompt(formats strfmt.Registry) error {
	if swag.IsZero(m.WrapupPrompt) { // not required
		return nil
	}

	// value enum
	if err := m.validateWrapupPromptEnum("wrapupPrompt", "body", m.WrapupPrompt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this participant based on the context it is used
func (m *Participant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCallbacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCalls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCobrowsesessions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversationRoutingData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvaluations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScreenshares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSocialExpressions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVideos(ctx, formats); err != nil {
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

func (m *Participant) contextValidateCallbacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Callbacks); i++ {

		if m.Callbacks[i] != nil {
			if err := m.Callbacks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("callbacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("callbacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateCalls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Calls); i++ {

		if m.Calls[i] != nil {
			if err := m.Calls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("calls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("calls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateChats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Chats); i++ {

		if m.Chats[i] != nil {
			if err := m.Chats[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateCobrowsesessions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cobrowsesessions); i++ {

		if m.Cobrowsesessions[i] != nil {
			if err := m.Cobrowsesessions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cobrowsesessions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cobrowsesessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateConversationRoutingData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Participant) contextValidateEmails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Emails); i++ {

		if m.Emails[i] != nil {
			if err := m.Emails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateEvaluations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Evaluations); i++ {

		if m.Evaluations[i] != nil {
			if err := m.Evaluations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evaluations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("evaluations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Participant) contextValidateScreenshares(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Screenshares); i++ {

		if m.Screenshares[i] != nil {
			if err := m.Screenshares[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("screenshares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("screenshares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateSocialExpressions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SocialExpressions); i++ {

		if m.SocialExpressions[i] != nil {
			if err := m.SocialExpressions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("socialExpressions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("socialExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateVideos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Videos); i++ {

		if m.Videos[i] != nil {
			if err := m.Videos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("videos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("videos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Participant) contextValidateWrapup(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Participant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Participant) UnmarshalBinary(b []byte) error {
	var res Participant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
