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

// FlowAggregateQueryPredicate flow aggregate query predicate
//
// swagger:model FlowAggregateQueryPredicate
type FlowAggregateQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [activeSkillId addressFrom addressTo agentAssistantId agentBullseyeRing agentOwned agentRank agentScore ani assignerId authenticated conversationId convertedFrom convertedTo deliveryStatus destinationAddress direction disconnectType divisionId dnis edgeId eligibleAgentCount endingLanguage entryReason entryType exitReason externalContactId externalMediaCount externalOrganizationId externalTag firstQueue flaggedReason flowId flowInType flowMilestoneId flowName flowOutType flowOutcome flowOutcomeId flowOutcomeValue flowType flowVersion groupId interactionType journeyActionId journeyActionMapId journeyActionMapVersion journeyCustomerId journeyCustomerIdType journeyCustomerSessionId journeyCustomerSessionIdType knowledgeBaseId mediaCount mediaType messageType originatingDirection outboundCampaignId outboundContactId outboundContactListId participantName peerId proposedAgentId provider purpose queueId recognitionFailureReason remote removedSkillId reoffered requestedLanguageId requestedRouting requestedRoutingSkillId roomId routingPriority routingRing scoredAgentId selectedAgentId selectedAgentRank selfServed sessionDnis sessionId startingLanguage stationId teamId transferTargetAddress transferTargetName transferType usedRouting userId waitingInteractionCount wrapUpCode]
	Dimension string `json:"dimension,omitempty"`

	// Optional operator, default is matches
	// Enum: [matches exists notExists]
	Operator string `json:"operator,omitempty"`

	// Right hand side for dimension predicates
	Range *NumericRange `json:"range,omitempty"`

	// Optional type, can usually be inferred
	// Enum: [dimension property metric]
	Type string `json:"type,omitempty"`

	// Right hand side for dimension predicates
	Value string `json:"value,omitempty"`
}

// Validate validates this flow aggregate query predicate
func (m *FlowAggregateQueryPredicate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDimension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var flowAggregateQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["activeSkillId","addressFrom","addressTo","agentAssistantId","agentBullseyeRing","agentOwned","agentRank","agentScore","ani","assignerId","authenticated","conversationId","convertedFrom","convertedTo","deliveryStatus","destinationAddress","direction","disconnectType","divisionId","dnis","edgeId","eligibleAgentCount","endingLanguage","entryReason","entryType","exitReason","externalContactId","externalMediaCount","externalOrganizationId","externalTag","firstQueue","flaggedReason","flowId","flowInType","flowMilestoneId","flowName","flowOutType","flowOutcome","flowOutcomeId","flowOutcomeValue","flowType","flowVersion","groupId","interactionType","journeyActionId","journeyActionMapId","journeyActionMapVersion","journeyCustomerId","journeyCustomerIdType","journeyCustomerSessionId","journeyCustomerSessionIdType","knowledgeBaseId","mediaCount","mediaType","messageType","originatingDirection","outboundCampaignId","outboundContactId","outboundContactListId","participantName","peerId","proposedAgentId","provider","purpose","queueId","recognitionFailureReason","remote","removedSkillId","reoffered","requestedLanguageId","requestedRouting","requestedRoutingSkillId","roomId","routingPriority","routingRing","scoredAgentId","selectedAgentId","selectedAgentRank","selfServed","sessionDnis","sessionId","startingLanguage","stationId","teamId","transferTargetAddress","transferTargetName","transferType","usedRouting","userId","waitingInteractionCount","wrapUpCode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowAggregateQueryPredicateTypeDimensionPropEnum = append(flowAggregateQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// FlowAggregateQueryPredicateDimensionActiveSkillID captures enum value "activeSkillId"
	FlowAggregateQueryPredicateDimensionActiveSkillID string = "activeSkillId"

	// FlowAggregateQueryPredicateDimensionAddressFrom captures enum value "addressFrom"
	FlowAggregateQueryPredicateDimensionAddressFrom string = "addressFrom"

	// FlowAggregateQueryPredicateDimensionAddressTo captures enum value "addressTo"
	FlowAggregateQueryPredicateDimensionAddressTo string = "addressTo"

	// FlowAggregateQueryPredicateDimensionAgentAssistantID captures enum value "agentAssistantId"
	FlowAggregateQueryPredicateDimensionAgentAssistantID string = "agentAssistantId"

	// FlowAggregateQueryPredicateDimensionAgentBullseyeRing captures enum value "agentBullseyeRing"
	FlowAggregateQueryPredicateDimensionAgentBullseyeRing string = "agentBullseyeRing"

	// FlowAggregateQueryPredicateDimensionAgentOwned captures enum value "agentOwned"
	FlowAggregateQueryPredicateDimensionAgentOwned string = "agentOwned"

	// FlowAggregateQueryPredicateDimensionAgentRank captures enum value "agentRank"
	FlowAggregateQueryPredicateDimensionAgentRank string = "agentRank"

	// FlowAggregateQueryPredicateDimensionAgentScore captures enum value "agentScore"
	FlowAggregateQueryPredicateDimensionAgentScore string = "agentScore"

	// FlowAggregateQueryPredicateDimensionAni captures enum value "ani"
	FlowAggregateQueryPredicateDimensionAni string = "ani"

	// FlowAggregateQueryPredicateDimensionAssignerID captures enum value "assignerId"
	FlowAggregateQueryPredicateDimensionAssignerID string = "assignerId"

	// FlowAggregateQueryPredicateDimensionAuthenticated captures enum value "authenticated"
	FlowAggregateQueryPredicateDimensionAuthenticated string = "authenticated"

	// FlowAggregateQueryPredicateDimensionConversationID captures enum value "conversationId"
	FlowAggregateQueryPredicateDimensionConversationID string = "conversationId"

	// FlowAggregateQueryPredicateDimensionConvertedFrom captures enum value "convertedFrom"
	FlowAggregateQueryPredicateDimensionConvertedFrom string = "convertedFrom"

	// FlowAggregateQueryPredicateDimensionConvertedTo captures enum value "convertedTo"
	FlowAggregateQueryPredicateDimensionConvertedTo string = "convertedTo"

	// FlowAggregateQueryPredicateDimensionDeliveryStatus captures enum value "deliveryStatus"
	FlowAggregateQueryPredicateDimensionDeliveryStatus string = "deliveryStatus"

	// FlowAggregateQueryPredicateDimensionDestinationAddress captures enum value "destinationAddress"
	FlowAggregateQueryPredicateDimensionDestinationAddress string = "destinationAddress"

	// FlowAggregateQueryPredicateDimensionDirection captures enum value "direction"
	FlowAggregateQueryPredicateDimensionDirection string = "direction"

	// FlowAggregateQueryPredicateDimensionDisconnectType captures enum value "disconnectType"
	FlowAggregateQueryPredicateDimensionDisconnectType string = "disconnectType"

	// FlowAggregateQueryPredicateDimensionDivisionID captures enum value "divisionId"
	FlowAggregateQueryPredicateDimensionDivisionID string = "divisionId"

	// FlowAggregateQueryPredicateDimensionDnis captures enum value "dnis"
	FlowAggregateQueryPredicateDimensionDnis string = "dnis"

	// FlowAggregateQueryPredicateDimensionEdgeID captures enum value "edgeId"
	FlowAggregateQueryPredicateDimensionEdgeID string = "edgeId"

	// FlowAggregateQueryPredicateDimensionEligibleAgentCount captures enum value "eligibleAgentCount"
	FlowAggregateQueryPredicateDimensionEligibleAgentCount string = "eligibleAgentCount"

	// FlowAggregateQueryPredicateDimensionEndingLanguage captures enum value "endingLanguage"
	FlowAggregateQueryPredicateDimensionEndingLanguage string = "endingLanguage"

	// FlowAggregateQueryPredicateDimensionEntryReason captures enum value "entryReason"
	FlowAggregateQueryPredicateDimensionEntryReason string = "entryReason"

	// FlowAggregateQueryPredicateDimensionEntryType captures enum value "entryType"
	FlowAggregateQueryPredicateDimensionEntryType string = "entryType"

	// FlowAggregateQueryPredicateDimensionExitReason captures enum value "exitReason"
	FlowAggregateQueryPredicateDimensionExitReason string = "exitReason"

	// FlowAggregateQueryPredicateDimensionExternalContactID captures enum value "externalContactId"
	FlowAggregateQueryPredicateDimensionExternalContactID string = "externalContactId"

	// FlowAggregateQueryPredicateDimensionExternalMediaCount captures enum value "externalMediaCount"
	FlowAggregateQueryPredicateDimensionExternalMediaCount string = "externalMediaCount"

	// FlowAggregateQueryPredicateDimensionExternalOrganizationID captures enum value "externalOrganizationId"
	FlowAggregateQueryPredicateDimensionExternalOrganizationID string = "externalOrganizationId"

	// FlowAggregateQueryPredicateDimensionExternalTag captures enum value "externalTag"
	FlowAggregateQueryPredicateDimensionExternalTag string = "externalTag"

	// FlowAggregateQueryPredicateDimensionFirstQueue captures enum value "firstQueue"
	FlowAggregateQueryPredicateDimensionFirstQueue string = "firstQueue"

	// FlowAggregateQueryPredicateDimensionFlaggedReason captures enum value "flaggedReason"
	FlowAggregateQueryPredicateDimensionFlaggedReason string = "flaggedReason"

	// FlowAggregateQueryPredicateDimensionFlowID captures enum value "flowId"
	FlowAggregateQueryPredicateDimensionFlowID string = "flowId"

	// FlowAggregateQueryPredicateDimensionFlowInType captures enum value "flowInType"
	FlowAggregateQueryPredicateDimensionFlowInType string = "flowInType"

	// FlowAggregateQueryPredicateDimensionFlowMilestoneID captures enum value "flowMilestoneId"
	FlowAggregateQueryPredicateDimensionFlowMilestoneID string = "flowMilestoneId"

	// FlowAggregateQueryPredicateDimensionFlowName captures enum value "flowName"
	FlowAggregateQueryPredicateDimensionFlowName string = "flowName"

	// FlowAggregateQueryPredicateDimensionFlowOutType captures enum value "flowOutType"
	FlowAggregateQueryPredicateDimensionFlowOutType string = "flowOutType"

	// FlowAggregateQueryPredicateDimensionFlowOutcome captures enum value "flowOutcome"
	FlowAggregateQueryPredicateDimensionFlowOutcome string = "flowOutcome"

	// FlowAggregateQueryPredicateDimensionFlowOutcomeID captures enum value "flowOutcomeId"
	FlowAggregateQueryPredicateDimensionFlowOutcomeID string = "flowOutcomeId"

	// FlowAggregateQueryPredicateDimensionFlowOutcomeValue captures enum value "flowOutcomeValue"
	FlowAggregateQueryPredicateDimensionFlowOutcomeValue string = "flowOutcomeValue"

	// FlowAggregateQueryPredicateDimensionFlowType captures enum value "flowType"
	FlowAggregateQueryPredicateDimensionFlowType string = "flowType"

	// FlowAggregateQueryPredicateDimensionFlowVersion captures enum value "flowVersion"
	FlowAggregateQueryPredicateDimensionFlowVersion string = "flowVersion"

	// FlowAggregateQueryPredicateDimensionGroupID captures enum value "groupId"
	FlowAggregateQueryPredicateDimensionGroupID string = "groupId"

	// FlowAggregateQueryPredicateDimensionInteractionType captures enum value "interactionType"
	FlowAggregateQueryPredicateDimensionInteractionType string = "interactionType"

	// FlowAggregateQueryPredicateDimensionJourneyActionID captures enum value "journeyActionId"
	FlowAggregateQueryPredicateDimensionJourneyActionID string = "journeyActionId"

	// FlowAggregateQueryPredicateDimensionJourneyActionMapID captures enum value "journeyActionMapId"
	FlowAggregateQueryPredicateDimensionJourneyActionMapID string = "journeyActionMapId"

	// FlowAggregateQueryPredicateDimensionJourneyActionMapVersion captures enum value "journeyActionMapVersion"
	FlowAggregateQueryPredicateDimensionJourneyActionMapVersion string = "journeyActionMapVersion"

	// FlowAggregateQueryPredicateDimensionJourneyCustomerID captures enum value "journeyCustomerId"
	FlowAggregateQueryPredicateDimensionJourneyCustomerID string = "journeyCustomerId"

	// FlowAggregateQueryPredicateDimensionJourneyCustomerIDType captures enum value "journeyCustomerIdType"
	FlowAggregateQueryPredicateDimensionJourneyCustomerIDType string = "journeyCustomerIdType"

	// FlowAggregateQueryPredicateDimensionJourneyCustomerSessionID captures enum value "journeyCustomerSessionId"
	FlowAggregateQueryPredicateDimensionJourneyCustomerSessionID string = "journeyCustomerSessionId"

	// FlowAggregateQueryPredicateDimensionJourneyCustomerSessionIDType captures enum value "journeyCustomerSessionIdType"
	FlowAggregateQueryPredicateDimensionJourneyCustomerSessionIDType string = "journeyCustomerSessionIdType"

	// FlowAggregateQueryPredicateDimensionKnowledgeBaseID captures enum value "knowledgeBaseId"
	FlowAggregateQueryPredicateDimensionKnowledgeBaseID string = "knowledgeBaseId"

	// FlowAggregateQueryPredicateDimensionMediaCount captures enum value "mediaCount"
	FlowAggregateQueryPredicateDimensionMediaCount string = "mediaCount"

	// FlowAggregateQueryPredicateDimensionMediaType captures enum value "mediaType"
	FlowAggregateQueryPredicateDimensionMediaType string = "mediaType"

	// FlowAggregateQueryPredicateDimensionMessageType captures enum value "messageType"
	FlowAggregateQueryPredicateDimensionMessageType string = "messageType"

	// FlowAggregateQueryPredicateDimensionOriginatingDirection captures enum value "originatingDirection"
	FlowAggregateQueryPredicateDimensionOriginatingDirection string = "originatingDirection"

	// FlowAggregateQueryPredicateDimensionOutboundCampaignID captures enum value "outboundCampaignId"
	FlowAggregateQueryPredicateDimensionOutboundCampaignID string = "outboundCampaignId"

	// FlowAggregateQueryPredicateDimensionOutboundContactID captures enum value "outboundContactId"
	FlowAggregateQueryPredicateDimensionOutboundContactID string = "outboundContactId"

	// FlowAggregateQueryPredicateDimensionOutboundContactListID captures enum value "outboundContactListId"
	FlowAggregateQueryPredicateDimensionOutboundContactListID string = "outboundContactListId"

	// FlowAggregateQueryPredicateDimensionParticipantName captures enum value "participantName"
	FlowAggregateQueryPredicateDimensionParticipantName string = "participantName"

	// FlowAggregateQueryPredicateDimensionPeerID captures enum value "peerId"
	FlowAggregateQueryPredicateDimensionPeerID string = "peerId"

	// FlowAggregateQueryPredicateDimensionProposedAgentID captures enum value "proposedAgentId"
	FlowAggregateQueryPredicateDimensionProposedAgentID string = "proposedAgentId"

	// FlowAggregateQueryPredicateDimensionProvider captures enum value "provider"
	FlowAggregateQueryPredicateDimensionProvider string = "provider"

	// FlowAggregateQueryPredicateDimensionPurpose captures enum value "purpose"
	FlowAggregateQueryPredicateDimensionPurpose string = "purpose"

	// FlowAggregateQueryPredicateDimensionQueueID captures enum value "queueId"
	FlowAggregateQueryPredicateDimensionQueueID string = "queueId"

	// FlowAggregateQueryPredicateDimensionRecognitionFailureReason captures enum value "recognitionFailureReason"
	FlowAggregateQueryPredicateDimensionRecognitionFailureReason string = "recognitionFailureReason"

	// FlowAggregateQueryPredicateDimensionRemote captures enum value "remote"
	FlowAggregateQueryPredicateDimensionRemote string = "remote"

	// FlowAggregateQueryPredicateDimensionRemovedSkillID captures enum value "removedSkillId"
	FlowAggregateQueryPredicateDimensionRemovedSkillID string = "removedSkillId"

	// FlowAggregateQueryPredicateDimensionReoffered captures enum value "reoffered"
	FlowAggregateQueryPredicateDimensionReoffered string = "reoffered"

	// FlowAggregateQueryPredicateDimensionRequestedLanguageID captures enum value "requestedLanguageId"
	FlowAggregateQueryPredicateDimensionRequestedLanguageID string = "requestedLanguageId"

	// FlowAggregateQueryPredicateDimensionRequestedRouting captures enum value "requestedRouting"
	FlowAggregateQueryPredicateDimensionRequestedRouting string = "requestedRouting"

	// FlowAggregateQueryPredicateDimensionRequestedRoutingSkillID captures enum value "requestedRoutingSkillId"
	FlowAggregateQueryPredicateDimensionRequestedRoutingSkillID string = "requestedRoutingSkillId"

	// FlowAggregateQueryPredicateDimensionRoomID captures enum value "roomId"
	FlowAggregateQueryPredicateDimensionRoomID string = "roomId"

	// FlowAggregateQueryPredicateDimensionRoutingPriority captures enum value "routingPriority"
	FlowAggregateQueryPredicateDimensionRoutingPriority string = "routingPriority"

	// FlowAggregateQueryPredicateDimensionRoutingRing captures enum value "routingRing"
	FlowAggregateQueryPredicateDimensionRoutingRing string = "routingRing"

	// FlowAggregateQueryPredicateDimensionScoredAgentID captures enum value "scoredAgentId"
	FlowAggregateQueryPredicateDimensionScoredAgentID string = "scoredAgentId"

	// FlowAggregateQueryPredicateDimensionSelectedAgentID captures enum value "selectedAgentId"
	FlowAggregateQueryPredicateDimensionSelectedAgentID string = "selectedAgentId"

	// FlowAggregateQueryPredicateDimensionSelectedAgentRank captures enum value "selectedAgentRank"
	FlowAggregateQueryPredicateDimensionSelectedAgentRank string = "selectedAgentRank"

	// FlowAggregateQueryPredicateDimensionSelfServed captures enum value "selfServed"
	FlowAggregateQueryPredicateDimensionSelfServed string = "selfServed"

	// FlowAggregateQueryPredicateDimensionSessionDnis captures enum value "sessionDnis"
	FlowAggregateQueryPredicateDimensionSessionDnis string = "sessionDnis"

	// FlowAggregateQueryPredicateDimensionSessionID captures enum value "sessionId"
	FlowAggregateQueryPredicateDimensionSessionID string = "sessionId"

	// FlowAggregateQueryPredicateDimensionStartingLanguage captures enum value "startingLanguage"
	FlowAggregateQueryPredicateDimensionStartingLanguage string = "startingLanguage"

	// FlowAggregateQueryPredicateDimensionStationID captures enum value "stationId"
	FlowAggregateQueryPredicateDimensionStationID string = "stationId"

	// FlowAggregateQueryPredicateDimensionTeamID captures enum value "teamId"
	FlowAggregateQueryPredicateDimensionTeamID string = "teamId"

	// FlowAggregateQueryPredicateDimensionTransferTargetAddress captures enum value "transferTargetAddress"
	FlowAggregateQueryPredicateDimensionTransferTargetAddress string = "transferTargetAddress"

	// FlowAggregateQueryPredicateDimensionTransferTargetName captures enum value "transferTargetName"
	FlowAggregateQueryPredicateDimensionTransferTargetName string = "transferTargetName"

	// FlowAggregateQueryPredicateDimensionTransferType captures enum value "transferType"
	FlowAggregateQueryPredicateDimensionTransferType string = "transferType"

	// FlowAggregateQueryPredicateDimensionUsedRouting captures enum value "usedRouting"
	FlowAggregateQueryPredicateDimensionUsedRouting string = "usedRouting"

	// FlowAggregateQueryPredicateDimensionUserID captures enum value "userId"
	FlowAggregateQueryPredicateDimensionUserID string = "userId"

	// FlowAggregateQueryPredicateDimensionWaitingInteractionCount captures enum value "waitingInteractionCount"
	FlowAggregateQueryPredicateDimensionWaitingInteractionCount string = "waitingInteractionCount"

	// FlowAggregateQueryPredicateDimensionWrapUpCode captures enum value "wrapUpCode"
	FlowAggregateQueryPredicateDimensionWrapUpCode string = "wrapUpCode"
)

// prop value enum
func (m *FlowAggregateQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, flowAggregateQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FlowAggregateQueryPredicate) validateDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var flowAggregateQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowAggregateQueryPredicateTypeOperatorPropEnum = append(flowAggregateQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// FlowAggregateQueryPredicateOperatorMatches captures enum value "matches"
	FlowAggregateQueryPredicateOperatorMatches string = "matches"

	// FlowAggregateQueryPredicateOperatorExists captures enum value "exists"
	FlowAggregateQueryPredicateOperatorExists string = "exists"

	// FlowAggregateQueryPredicateOperatorNotExists captures enum value "notExists"
	FlowAggregateQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *FlowAggregateQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, flowAggregateQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FlowAggregateQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *FlowAggregateQueryPredicate) validateRange(formats strfmt.Registry) error {

	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

var flowAggregateQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowAggregateQueryPredicateTypeTypePropEnum = append(flowAggregateQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// FlowAggregateQueryPredicateTypeDimension captures enum value "dimension"
	FlowAggregateQueryPredicateTypeDimension string = "dimension"

	// FlowAggregateQueryPredicateTypeProperty captures enum value "property"
	FlowAggregateQueryPredicateTypeProperty string = "property"

	// FlowAggregateQueryPredicateTypeMetric captures enum value "metric"
	FlowAggregateQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *FlowAggregateQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, flowAggregateQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FlowAggregateQueryPredicate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowAggregateQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowAggregateQueryPredicate) UnmarshalBinary(b []byte) error {
	var res FlowAggregateQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
