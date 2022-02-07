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

// AnalyticsConversation analytics conversation
//
// swagger:model AnalyticsConversation
type AnalyticsConversation struct {

	// The end time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConversationEnd strfmt.DateTime `json:"conversationEnd,omitempty"`

	// Unique identifier for the conversation
	ConversationID string `json:"conversationId,omitempty"`

	// Indicates the participant purpose of the participant initiating a message conversation
	// Enum: [acd agent api botflow campaign customer dialer external fax group inbound ivr manual outbound station user voicemail workflow]
	ConversationInitiator string `json:"conversationInitiator,omitempty"`

	// The start time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ConversationStart strfmt.DateTime `json:"conversationStart,omitempty"`

	// Identifier(s) of division(s) associated with a conversation
	DivisionIds []string `json:"divisionIds"`

	// Evaluations associated with this conversation
	Evaluations []*AnalyticsEvaluation `json:"evaluations"`

	// External tag for the conversation
	ExternalTag string `json:"externalTag,omitempty"`

	// The unique identifier(s) of the knowledge base(s) used
	KnowledgeBaseIds []string `json:"knowledgeBaseIds"`

	// The lowest estimated average MOS among all the audio streams belonging to this conversation
	MediaStatsMinConversationMos float64 `json:"mediaStatsMinConversationMos,omitempty"`

	// The lowest R-factor value among all of the audio streams belonging to this conversation
	MediaStatsMinConversationRFactor float64 `json:"mediaStatsMinConversationRFactor,omitempty"`

	// The original direction of the conversation
	// Enum: [inbound outbound]
	OriginatingDirection string `json:"originatingDirection,omitempty"`

	// Participants in the conversation
	Participants []*AnalyticsParticipant `json:"participants"`

	// Resolutions associated with this conversation
	Resolutions []*AnalyticsResolution `json:"resolutions"`

	// Indicates whether all flow sessions were self serviced
	SelfServed bool `json:"selfServed"`

	// Surveys associated with this conversation
	Surveys []*AnalyticsSurvey `json:"surveys"`
}

// Validate validates this analytics conversation
func (m *AnalyticsConversation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversationEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationInitiator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginatingDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolutions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsConversation) validateConversationEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("conversationEnd", "body", "date-time", m.ConversationEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

var analyticsConversationTypeConversationInitiatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["acd","agent","api","botflow","campaign","customer","dialer","external","fax","group","inbound","ivr","manual","outbound","station","user","voicemail","workflow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsConversationTypeConversationInitiatorPropEnum = append(analyticsConversationTypeConversationInitiatorPropEnum, v)
	}
}

const (

	// AnalyticsConversationConversationInitiatorAcd captures enum value "acd"
	AnalyticsConversationConversationInitiatorAcd string = "acd"

	// AnalyticsConversationConversationInitiatorAgent captures enum value "agent"
	AnalyticsConversationConversationInitiatorAgent string = "agent"

	// AnalyticsConversationConversationInitiatorAPI captures enum value "api"
	AnalyticsConversationConversationInitiatorAPI string = "api"

	// AnalyticsConversationConversationInitiatorBotflow captures enum value "botflow"
	AnalyticsConversationConversationInitiatorBotflow string = "botflow"

	// AnalyticsConversationConversationInitiatorCampaign captures enum value "campaign"
	AnalyticsConversationConversationInitiatorCampaign string = "campaign"

	// AnalyticsConversationConversationInitiatorCustomer captures enum value "customer"
	AnalyticsConversationConversationInitiatorCustomer string = "customer"

	// AnalyticsConversationConversationInitiatorDialer captures enum value "dialer"
	AnalyticsConversationConversationInitiatorDialer string = "dialer"

	// AnalyticsConversationConversationInitiatorExternal captures enum value "external"
	AnalyticsConversationConversationInitiatorExternal string = "external"

	// AnalyticsConversationConversationInitiatorFax captures enum value "fax"
	AnalyticsConversationConversationInitiatorFax string = "fax"

	// AnalyticsConversationConversationInitiatorGroup captures enum value "group"
	AnalyticsConversationConversationInitiatorGroup string = "group"

	// AnalyticsConversationConversationInitiatorInbound captures enum value "inbound"
	AnalyticsConversationConversationInitiatorInbound string = "inbound"

	// AnalyticsConversationConversationInitiatorIvr captures enum value "ivr"
	AnalyticsConversationConversationInitiatorIvr string = "ivr"

	// AnalyticsConversationConversationInitiatorManual captures enum value "manual"
	AnalyticsConversationConversationInitiatorManual string = "manual"

	// AnalyticsConversationConversationInitiatorOutbound captures enum value "outbound"
	AnalyticsConversationConversationInitiatorOutbound string = "outbound"

	// AnalyticsConversationConversationInitiatorStation captures enum value "station"
	AnalyticsConversationConversationInitiatorStation string = "station"

	// AnalyticsConversationConversationInitiatorUser captures enum value "user"
	AnalyticsConversationConversationInitiatorUser string = "user"

	// AnalyticsConversationConversationInitiatorVoicemail captures enum value "voicemail"
	AnalyticsConversationConversationInitiatorVoicemail string = "voicemail"

	// AnalyticsConversationConversationInitiatorWorkflow captures enum value "workflow"
	AnalyticsConversationConversationInitiatorWorkflow string = "workflow"
)

// prop value enum
func (m *AnalyticsConversation) validateConversationInitiatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsConversationTypeConversationInitiatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsConversation) validateConversationInitiator(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationInitiator) { // not required
		return nil
	}

	// value enum
	if err := m.validateConversationInitiatorEnum("conversationInitiator", "body", m.ConversationInitiator); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsConversation) validateConversationStart(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationStart) { // not required
		return nil
	}

	if err := validate.FormatOf("conversationStart", "body", "date-time", m.ConversationStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsConversation) validateEvaluations(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

var analyticsConversationTypeOriginatingDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsConversationTypeOriginatingDirectionPropEnum = append(analyticsConversationTypeOriginatingDirectionPropEnum, v)
	}
}

const (

	// AnalyticsConversationOriginatingDirectionInbound captures enum value "inbound"
	AnalyticsConversationOriginatingDirectionInbound string = "inbound"

	// AnalyticsConversationOriginatingDirectionOutbound captures enum value "outbound"
	AnalyticsConversationOriginatingDirectionOutbound string = "outbound"
)

// prop value enum
func (m *AnalyticsConversation) validateOriginatingDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsConversationTypeOriginatingDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsConversation) validateOriginatingDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginatingDirection) { // not required
		return nil
	}

	// value enum
	if err := m.validateOriginatingDirectionEnum("originatingDirection", "body", m.OriginatingDirection); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsConversation) validateParticipants(formats strfmt.Registry) error {

	if swag.IsZero(m.Participants) { // not required
		return nil
	}

	for i := 0; i < len(m.Participants); i++ {
		if swag.IsZero(m.Participants[i]) { // not required
			continue
		}

		if m.Participants[i] != nil {
			if err := m.Participants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("participants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnalyticsConversation) validateResolutions(formats strfmt.Registry) error {

	if swag.IsZero(m.Resolutions) { // not required
		return nil
	}

	for i := 0; i < len(m.Resolutions); i++ {
		if swag.IsZero(m.Resolutions[i]) { // not required
			continue
		}

		if m.Resolutions[i] != nil {
			if err := m.Resolutions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resolutions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnalyticsConversation) validateSurveys(formats strfmt.Registry) error {

	if swag.IsZero(m.Surveys) { // not required
		return nil
	}

	for i := 0; i < len(m.Surveys); i++ {
		if swag.IsZero(m.Surveys[i]) { // not required
			continue
		}

		if m.Surveys[i] != nil {
			if err := m.Surveys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("surveys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsConversation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsConversation) UnmarshalBinary(b []byte) error {
	var res AnalyticsConversation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
