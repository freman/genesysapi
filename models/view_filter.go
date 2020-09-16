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

// ViewFilter view filter
//
// swagger:model ViewFilter
type ViewFilter struct {

	// The abandon durations in milliseconds used to filter the view
	AbandonDurationsMilliseconds []*NumericRange `json:"abandonDurationsMilliseconds"`

	// Indicates filtering for abandons
	Abandoned bool `json:"abandoned"`

	// The acd durations in milliseconds used to filter the view
	AcdDurationsMilliseconds []*NumericRange `json:"acdDurationsMilliseconds"`

	// The acw durations in milliseconds used to filter the view
	AcwDurationsMilliseconds []*NumericRange `json:"acwDurationsMilliseconds"`

	// The address from values are used to filter the view
	AddressFroms []string `json:"addressFroms"`

	// The address To values are used to filter the view
	AddressTos []string `json:"addressTos"`

	// The ani list ids are used to filter the view
	AniList []string `json:"aniList"`

	// Indicates filtering for answered interactions
	Answered bool `json:"answered"`

	// An interval of time to filter for scheduled callbacks. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	CallbackInterval string `json:"callbackInterval,omitempty"`

	// A list of callback numbers or substrings of numbers (ex: ["317", "13172222222"])
	CallbackNumberList []string `json:"callbackNumberList"`

	// The contact ids are used to filter the view
	ContactIds []string `json:"contactIds"`

	// The list of conversation ids used to filter the view
	ConversationIds []string `json:"conversationIds"`

	// A grouping of conversation level filters
	ConversationProperties *ConversationProperties `json:"conversationProperties,omitempty"`

	// The directions are used to filter the view
	Directions []string `json:"directions"`

	// The divison Ids used to filter the view
	DivisionIds []string `json:"divisionIds"`

	// The dnis list is used to filter the view
	DnisList []string `json:"dnisList"`

	// The durations in milliseconds used to filter the view
	DurationsMilliseconds []*NumericRange `json:"durationsMilliseconds"`

	// The evaluated agent ids are used to filter the view
	EvaluatedAgentIds []string `json:"evaluatedAgentIds"`

	// The evaluationCriticalScore is used to filter the view
	EvaluationCriticalScore *NumericRange `json:"evaluationCriticalScore,omitempty"`

	// The evaluation form ids are used to filter the view
	EvaluationFormIds []string `json:"evaluationFormIds"`

	// The evaluationScore is used to filter the view
	EvaluationScore *NumericRange `json:"evaluationScore,omitempty"`

	// The evaluator ids are used to filter the view
	EvaluatorIds []string `json:"evaluatorIds"`

	// The external contact ids are used to filter the view
	ExternalContactIds []string `json:"externalContactIds"`

	// The external org ids are used to filter the view
	ExternalOrgIds []string `json:"externalOrgIds"`

	// The user ids are used to fetch associated queues for the view
	FilterQueuesByUserIds []string `json:"filterQueuesByUserIds"`

	// The queue ids are used to fetch associated users for the view
	FilterUsersByQueueIds []string `json:"filterUsersByQueueIds"`

	// The list of destination types of the flow
	FlowDestinationTypes []string `json:"flowDestinationTypes"`

	// The list of reasons for the flow to disconnect
	FlowDisconnectReasons []string `json:"flowDisconnectReasons"`

	// A list of reasons of flow entry
	FlowEntryReasons []string `json:"flowEntryReasons"`

	// A list of types of the flow entry
	FlowEntryTypes []string `json:"flowEntryTypes"`

	// The list of flow Ids
	FlowIds []string `json:"flowIds"`

	// A list of flow out types
	FlowOutTypes []string `json:"flowOutTypes"`

	// A list of outcome ids of the flow
	FlowOutcomeIds []string `json:"flowOutcomeIds"`

	// A list of outcome values of the flow
	FlowOutcomeValues []string `json:"flowOutcomeValues"`

	// A list of types of the flow
	FlowTypes []string `json:"flowTypes"`

	// A list of versions of a flow
	FlowVersions []string `json:"flowVersions"`

	// A list of directory group ids
	GroupIds []string `json:"groupIds"`

	// The handle durations in milliseconds used to filter the view
	HandleDurationsMilliseconds []*NumericRange `json:"handleDurationsMilliseconds"`

	// Indicates filtering for Journey action map id
	HasJourneyActionMapID bool `json:"hasJourneyActionMapId"`

	// Indicates filtering for journey customer id
	HasJourneyCustomerID bool `json:"hasJourneyCustomerId"`

	// Indicates filtering for Journey visit id
	HasJourneyVisitID bool `json:"hasJourneyVisitId"`

	// Indicates filtering for presence of MMS media
	HasMedia bool `json:"hasMedia"`

	// The hold durations in milliseconds used to filter the view
	HoldDurationsMilliseconds []*NumericRange `json:"holdDurationsMilliseconds"`

	// Indicates filtering for blind transferred
	IsBlindTransferred bool `json:"isBlindTransferred"`

	// Indicates filtering for campaign
	IsCampaign bool `json:"isCampaign"`

	// Indicates filtering for consult transferred
	IsConsultTransferred bool `json:"isConsultTransferred"`

	// Indicates filtering for consulted
	IsConsulted bool `json:"isConsulted"`

	// Indicates filtering for ended
	IsEnded bool `json:"isEnded"`

	// Indicates filtering for survey
	IsSurveyed bool `json:"isSurveyed"`

	// The language groups used to filter the view
	LanguageGroups []string `json:"languageGroups"`

	// The language ids are used to filter the view
	LanguageIds []string `json:"languageIds"`

	// The location Ids used to filter the view
	LocationIds []string `json:"locationIds"`

	// The media types are used to filter the view
	MediaTypes []string `json:"mediaTypes"`

	// The message media types used to filter the view
	MessageTypes []string `json:"messageTypes"`

	// The desired range for mos values
	Mos *NumericRange `json:"mos,omitempty"`

	// The list of orginating directions used to filter the view
	OriginatingDirections []string `json:"originatingDirections"`

	// The outbound campaign ids are used to filter the view
	OutboundCampaignIds []string `json:"outboundCampaignIds"`

	// The outbound contact list ids are used to filter the view
	OutboundContactListIds []string `json:"outboundContactListIds"`

	// The list of promoter score ranges used to filter the view
	PromoterScores []*NumericRange `json:"promoterScores"`

	// A list of providers
	ProviderList []string `json:"providerList"`

	// The queue ids are used to filter the view
	QueueIds []string `json:"queueIds"`

	// The list of remote participants used to filter the view
	RemoteParticipants []string `json:"remoteParticipants"`

	// The report to user IDs used to filter the view
	ReportsTos []string `json:"reportsTos"`

	// A list of routing types requested
	RequestedRoutingTypes []string `json:"requestedRoutingTypes"`

	// The role Ids used to filter the view
	RoleIds []string `json:"roleIds"`

	// The list of session dnis used to filter the view
	SessionDnisList []string `json:"sessionDnisList"`

	// The list of SIP call ids used to filter the view
	SipCallIds []string `json:"sipCallIds"`

	// The skill groups used to filter the view
	SkillGroups []string `json:"skillGroups"`

	// The skill ids are used to filter the view
	SkillIds []string `json:"skillIds"`

	// The list of survey form context ids used to filter the view
	SurveyFormContextIds []string `json:"surveyFormContextIds"`

	// The survey form ids used to filter the view
	SurveyFormIds []string `json:"surveyFormIds"`

	// The survey NPS score used to filter the view
	SurveyNpsScore *NumericRange `json:"surveyNpsScore,omitempty"`

	// The survey promoter score used to filter the view
	SurveyPromoterScore *NumericRange `json:"surveyPromoterScore,omitempty"`

	// The survey question group score used to filter the view
	SurveyQuestionGroupScore *NumericRange `json:"surveyQuestionGroupScore,omitempty"`

	// The list of survey score ranges used to filter the view
	SurveyScores []*NumericRange `json:"surveyScores"`

	// The list of survey statuses used to filter the view
	SurveyStatuses []string `json:"surveyStatuses"`

	// The survey total score used to filter the view
	SurveyTotalScore *NumericRange `json:"surveyTotalScore,omitempty"`

	// The talk durations in milliseconds used to filter the view
	TalkDurationsMilliseconds []*NumericRange `json:"talkDurationsMilliseconds"`

	// Indicates filtering for transfers
	Transferred bool `json:"transferred"`

	// A list of routing types used
	UsedRoutingTypes []string `json:"usedRoutingTypes"`

	// The user ids are used to filter the view
	UserIds []string `json:"userIds"`

	// The wrap up codes are used to filter the view
	WrapUpCodes []string `json:"wrapUpCodes"`
}

// Validate validates this view filter
func (m *ViewFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbandonDurationsMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcdDurationsMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcwDurationsMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDurationsMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluationCriticalScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluationScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowDestinationTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowDisconnectReasons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowEntryTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowOutcomeValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandleDurationsMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldDurationsMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginatingDirections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePromoterScores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedRoutingTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveyNpsScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveyPromoterScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveyQuestionGroupScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveyScores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveyTotalScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTalkDurationsMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedRoutingTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewFilter) validateAbandonDurationsMilliseconds(formats strfmt.Registry) error {

	if swag.IsZero(m.AbandonDurationsMilliseconds) { // not required
		return nil
	}

	for i := 0; i < len(m.AbandonDurationsMilliseconds); i++ {
		if swag.IsZero(m.AbandonDurationsMilliseconds[i]) { // not required
			continue
		}

		if m.AbandonDurationsMilliseconds[i] != nil {
			if err := m.AbandonDurationsMilliseconds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("abandonDurationsMilliseconds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ViewFilter) validateAcdDurationsMilliseconds(formats strfmt.Registry) error {

	if swag.IsZero(m.AcdDurationsMilliseconds) { // not required
		return nil
	}

	for i := 0; i < len(m.AcdDurationsMilliseconds); i++ {
		if swag.IsZero(m.AcdDurationsMilliseconds[i]) { // not required
			continue
		}

		if m.AcdDurationsMilliseconds[i] != nil {
			if err := m.AcdDurationsMilliseconds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acdDurationsMilliseconds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ViewFilter) validateAcwDurationsMilliseconds(formats strfmt.Registry) error {

	if swag.IsZero(m.AcwDurationsMilliseconds) { // not required
		return nil
	}

	for i := 0; i < len(m.AcwDurationsMilliseconds); i++ {
		if swag.IsZero(m.AcwDurationsMilliseconds[i]) { // not required
			continue
		}

		if m.AcwDurationsMilliseconds[i] != nil {
			if err := m.AcwDurationsMilliseconds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acwDurationsMilliseconds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ViewFilter) validateConversationProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationProperties) { // not required
		return nil
	}

	if m.ConversationProperties != nil {
		if err := m.ConversationProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversationProperties")
			}
			return err
		}
	}

	return nil
}

var viewFilterDirectionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterDirectionsItemsEnum = append(viewFilterDirectionsItemsEnum, v)
	}
}

func (m *ViewFilter) validateDirectionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterDirectionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateDirections(formats strfmt.Registry) error {

	if swag.IsZero(m.Directions) { // not required
		return nil
	}

	for i := 0; i < len(m.Directions); i++ {

		// value enum
		if err := m.validateDirectionsItemsEnum("directions"+"."+strconv.Itoa(i), "body", m.Directions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ViewFilter) validateDurationsMilliseconds(formats strfmt.Registry) error {

	if swag.IsZero(m.DurationsMilliseconds) { // not required
		return nil
	}

	for i := 0; i < len(m.DurationsMilliseconds); i++ {
		if swag.IsZero(m.DurationsMilliseconds[i]) { // not required
			continue
		}

		if m.DurationsMilliseconds[i] != nil {
			if err := m.DurationsMilliseconds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("durationsMilliseconds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ViewFilter) validateEvaluationCriticalScore(formats strfmt.Registry) error {

	if swag.IsZero(m.EvaluationCriticalScore) { // not required
		return nil
	}

	if m.EvaluationCriticalScore != nil {
		if err := m.EvaluationCriticalScore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluationCriticalScore")
			}
			return err
		}
	}

	return nil
}

func (m *ViewFilter) validateEvaluationScore(formats strfmt.Registry) error {

	if swag.IsZero(m.EvaluationScore) { // not required
		return nil
	}

	if m.EvaluationScore != nil {
		if err := m.EvaluationScore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evaluationScore")
			}
			return err
		}
	}

	return nil
}

var viewFilterFlowDestinationTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACD","USER","GROUP","NUMBER","FLOW","SECURE_FLOW","ACD_VOICEMAIL","USER_VOICEMAIL","GROUP_VOICEMAIL","RETURN_TO_AGENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterFlowDestinationTypesItemsEnum = append(viewFilterFlowDestinationTypesItemsEnum, v)
	}
}

func (m *ViewFilter) validateFlowDestinationTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterFlowDestinationTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateFlowDestinationTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowDestinationTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.FlowDestinationTypes); i++ {

		// value enum
		if err := m.validateFlowDestinationTypesItemsEnum("flowDestinationTypes"+"."+strconv.Itoa(i), "body", m.FlowDestinationTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var viewFilterFlowDisconnectReasonsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FLOW_DISCONNECT","FLOW_ERROR_DISCONNECT","DISCONNECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterFlowDisconnectReasonsItemsEnum = append(viewFilterFlowDisconnectReasonsItemsEnum, v)
	}
}

func (m *ViewFilter) validateFlowDisconnectReasonsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterFlowDisconnectReasonsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateFlowDisconnectReasons(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowDisconnectReasons) { // not required
		return nil
	}

	for i := 0; i < len(m.FlowDisconnectReasons); i++ {

		// value enum
		if err := m.validateFlowDisconnectReasonsItemsEnum("flowDisconnectReasons"+"."+strconv.Itoa(i), "body", m.FlowDisconnectReasons[i]); err != nil {
			return err
		}

	}

	return nil
}

var viewFilterFlowEntryTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dnis","direct","flow","agent","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterFlowEntryTypesItemsEnum = append(viewFilterFlowEntryTypesItemsEnum, v)
	}
}

func (m *ViewFilter) validateFlowEntryTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterFlowEntryTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateFlowEntryTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowEntryTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.FlowEntryTypes); i++ {

		// value enum
		if err := m.validateFlowEntryTypesItemsEnum("flowEntryTypes"+"."+strconv.Itoa(i), "body", m.FlowEntryTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var viewFilterFlowOutcomeValuesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterFlowOutcomeValuesItemsEnum = append(viewFilterFlowOutcomeValuesItemsEnum, v)
	}
}

func (m *ViewFilter) validateFlowOutcomeValuesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterFlowOutcomeValuesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateFlowOutcomeValues(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowOutcomeValues) { // not required
		return nil
	}

	for i := 0; i < len(m.FlowOutcomeValues); i++ {

		// value enum
		if err := m.validateFlowOutcomeValuesItemsEnum("flowOutcomeValues"+"."+strconv.Itoa(i), "body", m.FlowOutcomeValues[i]); err != nil {
			return err
		}

	}

	return nil
}

var viewFilterFlowTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bot","commonmodule","inboundcall","inboundchat","inboundemail","inboundshortmessage","inqueuecall","outboundcall","securecall","surveyinvite","workflow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterFlowTypesItemsEnum = append(viewFilterFlowTypesItemsEnum, v)
	}
}

func (m *ViewFilter) validateFlowTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterFlowTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateFlowTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.FlowTypes); i++ {

		// value enum
		if err := m.validateFlowTypesItemsEnum("flowTypes"+"."+strconv.Itoa(i), "body", m.FlowTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ViewFilter) validateHandleDurationsMilliseconds(formats strfmt.Registry) error {

	if swag.IsZero(m.HandleDurationsMilliseconds) { // not required
		return nil
	}

	for i := 0; i < len(m.HandleDurationsMilliseconds); i++ {
		if swag.IsZero(m.HandleDurationsMilliseconds[i]) { // not required
			continue
		}

		if m.HandleDurationsMilliseconds[i] != nil {
			if err := m.HandleDurationsMilliseconds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("handleDurationsMilliseconds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ViewFilter) validateHoldDurationsMilliseconds(formats strfmt.Registry) error {

	if swag.IsZero(m.HoldDurationsMilliseconds) { // not required
		return nil
	}

	for i := 0; i < len(m.HoldDurationsMilliseconds); i++ {
		if swag.IsZero(m.HoldDurationsMilliseconds[i]) { // not required
			continue
		}

		if m.HoldDurationsMilliseconds[i] != nil {
			if err := m.HoldDurationsMilliseconds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("holdDurationsMilliseconds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var viewFilterMediaTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["voice","chat","email","callback","cobrowse","video","screenshare","message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterMediaTypesItemsEnum = append(viewFilterMediaTypesItemsEnum, v)
	}
}

func (m *ViewFilter) validateMediaTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterMediaTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateMediaTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.MediaTypes); i++ {

		// value enum
		if err := m.validateMediaTypesItemsEnum("mediaTypes"+"."+strconv.Itoa(i), "body", m.MediaTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var viewFilterMessageTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms","twitter","line","facebook","whatsapp","webmessaging"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterMessageTypesItemsEnum = append(viewFilterMessageTypesItemsEnum, v)
	}
}

func (m *ViewFilter) validateMessageTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterMessageTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateMessageTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.MessageTypes); i++ {

		// value enum
		if err := m.validateMessageTypesItemsEnum("messageTypes"+"."+strconv.Itoa(i), "body", m.MessageTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ViewFilter) validateMos(formats strfmt.Registry) error {

	if swag.IsZero(m.Mos) { // not required
		return nil
	}

	if m.Mos != nil {
		if err := m.Mos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mos")
			}
			return err
		}
	}

	return nil
}

var viewFilterOriginatingDirectionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterOriginatingDirectionsItemsEnum = append(viewFilterOriginatingDirectionsItemsEnum, v)
	}
}

func (m *ViewFilter) validateOriginatingDirectionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterOriginatingDirectionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateOriginatingDirections(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginatingDirections) { // not required
		return nil
	}

	for i := 0; i < len(m.OriginatingDirections); i++ {

		// value enum
		if err := m.validateOriginatingDirectionsItemsEnum("originatingDirections"+"."+strconv.Itoa(i), "body", m.OriginatingDirections[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ViewFilter) validatePromoterScores(formats strfmt.Registry) error {

	if swag.IsZero(m.PromoterScores) { // not required
		return nil
	}

	for i := 0; i < len(m.PromoterScores); i++ {
		if swag.IsZero(m.PromoterScores[i]) { // not required
			continue
		}

		if m.PromoterScores[i] != nil {
			if err := m.PromoterScores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("promoterScores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var viewFilterRequestedRoutingTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Predictive","Preferred","Manual","Last","Bullseye","Standard"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterRequestedRoutingTypesItemsEnum = append(viewFilterRequestedRoutingTypesItemsEnum, v)
	}
}

func (m *ViewFilter) validateRequestedRoutingTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterRequestedRoutingTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateRequestedRoutingTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedRoutingTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestedRoutingTypes); i++ {

		// value enum
		if err := m.validateRequestedRoutingTypesItemsEnum("requestedRoutingTypes"+"."+strconv.Itoa(i), "body", m.RequestedRoutingTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ViewFilter) validateSurveyNpsScore(formats strfmt.Registry) error {

	if swag.IsZero(m.SurveyNpsScore) { // not required
		return nil
	}

	if m.SurveyNpsScore != nil {
		if err := m.SurveyNpsScore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surveyNpsScore")
			}
			return err
		}
	}

	return nil
}

func (m *ViewFilter) validateSurveyPromoterScore(formats strfmt.Registry) error {

	if swag.IsZero(m.SurveyPromoterScore) { // not required
		return nil
	}

	if m.SurveyPromoterScore != nil {
		if err := m.SurveyPromoterScore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surveyPromoterScore")
			}
			return err
		}
	}

	return nil
}

func (m *ViewFilter) validateSurveyQuestionGroupScore(formats strfmt.Registry) error {

	if swag.IsZero(m.SurveyQuestionGroupScore) { // not required
		return nil
	}

	if m.SurveyQuestionGroupScore != nil {
		if err := m.SurveyQuestionGroupScore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surveyQuestionGroupScore")
			}
			return err
		}
	}

	return nil
}

func (m *ViewFilter) validateSurveyScores(formats strfmt.Registry) error {

	if swag.IsZero(m.SurveyScores) { // not required
		return nil
	}

	for i := 0; i < len(m.SurveyScores); i++ {
		if swag.IsZero(m.SurveyScores[i]) { // not required
			continue
		}

		if m.SurveyScores[i] != nil {
			if err := m.SurveyScores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("surveyScores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ViewFilter) validateSurveyTotalScore(formats strfmt.Registry) error {

	if swag.IsZero(m.SurveyTotalScore) { // not required
		return nil
	}

	if m.SurveyTotalScore != nil {
		if err := m.SurveyTotalScore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("surveyTotalScore")
			}
			return err
		}
	}

	return nil
}

func (m *ViewFilter) validateTalkDurationsMilliseconds(formats strfmt.Registry) error {

	if swag.IsZero(m.TalkDurationsMilliseconds) { // not required
		return nil
	}

	for i := 0; i < len(m.TalkDurationsMilliseconds); i++ {
		if swag.IsZero(m.TalkDurationsMilliseconds[i]) { // not required
			continue
		}

		if m.TalkDurationsMilliseconds[i] != nil {
			if err := m.TalkDurationsMilliseconds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("talkDurationsMilliseconds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var viewFilterUsedRoutingTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Predictive","Preferred","Manual","Last","Bullseye","Standard"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewFilterUsedRoutingTypesItemsEnum = append(viewFilterUsedRoutingTypesItemsEnum, v)
	}
}

func (m *ViewFilter) validateUsedRoutingTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewFilterUsedRoutingTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewFilter) validateUsedRoutingTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.UsedRoutingTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.UsedRoutingTypes); i++ {

		// value enum
		if err := m.validateUsedRoutingTypesItemsEnum("usedRoutingTypes"+"."+strconv.Itoa(i), "body", m.UsedRoutingTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ViewFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewFilter) UnmarshalBinary(b []byte) error {
	var res ViewFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
