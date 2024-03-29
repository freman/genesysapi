// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConversationAggregateQueryPredicate conversation aggregate query predicate
//
// swagger:model ConversationAggregateQueryPredicate
type ConversationAggregateQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [activeSkillId addressFrom addressTo agentAssistantId agentBullseyeRing agentOwned agentRank agentScore ani assignerId authenticated canonicalExternalContactId conversationId conversationInitiator convertedFrom convertedTo customerParticipation deliveryStatus destinationAddress direction disconnectType divisionId dnis edgeId eligibleAgentCount errorCode extendedDeliveryStatus externalContactId externalMediaCount externalOrganizationId externalTag firstQueue flaggedReason flowInType flowOutType groupId interactionType journeyActionId journeyActionMapId journeyActionMapVersion journeyCustomerId journeyCustomerIdType journeyCustomerSessionId journeyCustomerSessionIdType knowledgeBaseId mediaCount mediaType messageType originatingDirection outboundCampaignId outboundContactId outboundContactListId participantName peerId proposedAgentId provider purpose queueId remote removedSkillId reoffered requestedLanguageId requestedRouting requestedRoutingSkillId roomId routingPriority routingRing scoredAgentId selectedAgentId selectedAgentRank selfServed sessionDnis sessionId stationId teamId usedRouting userId waitingInteractionCount wrapUpCode]
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

// Validate validates this conversation aggregate query predicate
func (m *ConversationAggregateQueryPredicate) Validate(formats strfmt.Registry) error {
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

var conversationAggregateQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["activeSkillId","addressFrom","addressTo","agentAssistantId","agentBullseyeRing","agentOwned","agentRank","agentScore","ani","assignerId","authenticated","canonicalExternalContactId","conversationId","conversationInitiator","convertedFrom","convertedTo","customerParticipation","deliveryStatus","destinationAddress","direction","disconnectType","divisionId","dnis","edgeId","eligibleAgentCount","errorCode","extendedDeliveryStatus","externalContactId","externalMediaCount","externalOrganizationId","externalTag","firstQueue","flaggedReason","flowInType","flowOutType","groupId","interactionType","journeyActionId","journeyActionMapId","journeyActionMapVersion","journeyCustomerId","journeyCustomerIdType","journeyCustomerSessionId","journeyCustomerSessionIdType","knowledgeBaseId","mediaCount","mediaType","messageType","originatingDirection","outboundCampaignId","outboundContactId","outboundContactListId","participantName","peerId","proposedAgentId","provider","purpose","queueId","remote","removedSkillId","reoffered","requestedLanguageId","requestedRouting","requestedRoutingSkillId","roomId","routingPriority","routingRing","scoredAgentId","selectedAgentId","selectedAgentRank","selfServed","sessionDnis","sessionId","stationId","teamId","usedRouting","userId","waitingInteractionCount","wrapUpCode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationAggregateQueryPredicateTypeDimensionPropEnum = append(conversationAggregateQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// ConversationAggregateQueryPredicateDimensionActiveSkillID captures enum value "activeSkillId"
	ConversationAggregateQueryPredicateDimensionActiveSkillID string = "activeSkillId"

	// ConversationAggregateQueryPredicateDimensionAddressFrom captures enum value "addressFrom"
	ConversationAggregateQueryPredicateDimensionAddressFrom string = "addressFrom"

	// ConversationAggregateQueryPredicateDimensionAddressTo captures enum value "addressTo"
	ConversationAggregateQueryPredicateDimensionAddressTo string = "addressTo"

	// ConversationAggregateQueryPredicateDimensionAgentAssistantID captures enum value "agentAssistantId"
	ConversationAggregateQueryPredicateDimensionAgentAssistantID string = "agentAssistantId"

	// ConversationAggregateQueryPredicateDimensionAgentBullseyeRing captures enum value "agentBullseyeRing"
	ConversationAggregateQueryPredicateDimensionAgentBullseyeRing string = "agentBullseyeRing"

	// ConversationAggregateQueryPredicateDimensionAgentOwned captures enum value "agentOwned"
	ConversationAggregateQueryPredicateDimensionAgentOwned string = "agentOwned"

	// ConversationAggregateQueryPredicateDimensionAgentRank captures enum value "agentRank"
	ConversationAggregateQueryPredicateDimensionAgentRank string = "agentRank"

	// ConversationAggregateQueryPredicateDimensionAgentScore captures enum value "agentScore"
	ConversationAggregateQueryPredicateDimensionAgentScore string = "agentScore"

	// ConversationAggregateQueryPredicateDimensionAni captures enum value "ani"
	ConversationAggregateQueryPredicateDimensionAni string = "ani"

	// ConversationAggregateQueryPredicateDimensionAssignerID captures enum value "assignerId"
	ConversationAggregateQueryPredicateDimensionAssignerID string = "assignerId"

	// ConversationAggregateQueryPredicateDimensionAuthenticated captures enum value "authenticated"
	ConversationAggregateQueryPredicateDimensionAuthenticated string = "authenticated"

	// ConversationAggregateQueryPredicateDimensionCanonicalExternalContactID captures enum value "canonicalExternalContactId"
	ConversationAggregateQueryPredicateDimensionCanonicalExternalContactID string = "canonicalExternalContactId"

	// ConversationAggregateQueryPredicateDimensionConversationID captures enum value "conversationId"
	ConversationAggregateQueryPredicateDimensionConversationID string = "conversationId"

	// ConversationAggregateQueryPredicateDimensionConversationInitiator captures enum value "conversationInitiator"
	ConversationAggregateQueryPredicateDimensionConversationInitiator string = "conversationInitiator"

	// ConversationAggregateQueryPredicateDimensionConvertedFrom captures enum value "convertedFrom"
	ConversationAggregateQueryPredicateDimensionConvertedFrom string = "convertedFrom"

	// ConversationAggregateQueryPredicateDimensionConvertedTo captures enum value "convertedTo"
	ConversationAggregateQueryPredicateDimensionConvertedTo string = "convertedTo"

	// ConversationAggregateQueryPredicateDimensionCustomerParticipation captures enum value "customerParticipation"
	ConversationAggregateQueryPredicateDimensionCustomerParticipation string = "customerParticipation"

	// ConversationAggregateQueryPredicateDimensionDeliveryStatus captures enum value "deliveryStatus"
	ConversationAggregateQueryPredicateDimensionDeliveryStatus string = "deliveryStatus"

	// ConversationAggregateQueryPredicateDimensionDestinationAddress captures enum value "destinationAddress"
	ConversationAggregateQueryPredicateDimensionDestinationAddress string = "destinationAddress"

	// ConversationAggregateQueryPredicateDimensionDirection captures enum value "direction"
	ConversationAggregateQueryPredicateDimensionDirection string = "direction"

	// ConversationAggregateQueryPredicateDimensionDisconnectType captures enum value "disconnectType"
	ConversationAggregateQueryPredicateDimensionDisconnectType string = "disconnectType"

	// ConversationAggregateQueryPredicateDimensionDivisionID captures enum value "divisionId"
	ConversationAggregateQueryPredicateDimensionDivisionID string = "divisionId"

	// ConversationAggregateQueryPredicateDimensionDnis captures enum value "dnis"
	ConversationAggregateQueryPredicateDimensionDnis string = "dnis"

	// ConversationAggregateQueryPredicateDimensionEdgeID captures enum value "edgeId"
	ConversationAggregateQueryPredicateDimensionEdgeID string = "edgeId"

	// ConversationAggregateQueryPredicateDimensionEligibleAgentCount captures enum value "eligibleAgentCount"
	ConversationAggregateQueryPredicateDimensionEligibleAgentCount string = "eligibleAgentCount"

	// ConversationAggregateQueryPredicateDimensionErrorCode captures enum value "errorCode"
	ConversationAggregateQueryPredicateDimensionErrorCode string = "errorCode"

	// ConversationAggregateQueryPredicateDimensionExtendedDeliveryStatus captures enum value "extendedDeliveryStatus"
	ConversationAggregateQueryPredicateDimensionExtendedDeliveryStatus string = "extendedDeliveryStatus"

	// ConversationAggregateQueryPredicateDimensionExternalContactID captures enum value "externalContactId"
	ConversationAggregateQueryPredicateDimensionExternalContactID string = "externalContactId"

	// ConversationAggregateQueryPredicateDimensionExternalMediaCount captures enum value "externalMediaCount"
	ConversationAggregateQueryPredicateDimensionExternalMediaCount string = "externalMediaCount"

	// ConversationAggregateQueryPredicateDimensionExternalOrganizationID captures enum value "externalOrganizationId"
	ConversationAggregateQueryPredicateDimensionExternalOrganizationID string = "externalOrganizationId"

	// ConversationAggregateQueryPredicateDimensionExternalTag captures enum value "externalTag"
	ConversationAggregateQueryPredicateDimensionExternalTag string = "externalTag"

	// ConversationAggregateQueryPredicateDimensionFirstQueue captures enum value "firstQueue"
	ConversationAggregateQueryPredicateDimensionFirstQueue string = "firstQueue"

	// ConversationAggregateQueryPredicateDimensionFlaggedReason captures enum value "flaggedReason"
	ConversationAggregateQueryPredicateDimensionFlaggedReason string = "flaggedReason"

	// ConversationAggregateQueryPredicateDimensionFlowInType captures enum value "flowInType"
	ConversationAggregateQueryPredicateDimensionFlowInType string = "flowInType"

	// ConversationAggregateQueryPredicateDimensionFlowOutType captures enum value "flowOutType"
	ConversationAggregateQueryPredicateDimensionFlowOutType string = "flowOutType"

	// ConversationAggregateQueryPredicateDimensionGroupID captures enum value "groupId"
	ConversationAggregateQueryPredicateDimensionGroupID string = "groupId"

	// ConversationAggregateQueryPredicateDimensionInteractionType captures enum value "interactionType"
	ConversationAggregateQueryPredicateDimensionInteractionType string = "interactionType"

	// ConversationAggregateQueryPredicateDimensionJourneyActionID captures enum value "journeyActionId"
	ConversationAggregateQueryPredicateDimensionJourneyActionID string = "journeyActionId"

	// ConversationAggregateQueryPredicateDimensionJourneyActionMapID captures enum value "journeyActionMapId"
	ConversationAggregateQueryPredicateDimensionJourneyActionMapID string = "journeyActionMapId"

	// ConversationAggregateQueryPredicateDimensionJourneyActionMapVersion captures enum value "journeyActionMapVersion"
	ConversationAggregateQueryPredicateDimensionJourneyActionMapVersion string = "journeyActionMapVersion"

	// ConversationAggregateQueryPredicateDimensionJourneyCustomerID captures enum value "journeyCustomerId"
	ConversationAggregateQueryPredicateDimensionJourneyCustomerID string = "journeyCustomerId"

	// ConversationAggregateQueryPredicateDimensionJourneyCustomerIDType captures enum value "journeyCustomerIdType"
	ConversationAggregateQueryPredicateDimensionJourneyCustomerIDType string = "journeyCustomerIdType"

	// ConversationAggregateQueryPredicateDimensionJourneyCustomerSessionID captures enum value "journeyCustomerSessionId"
	ConversationAggregateQueryPredicateDimensionJourneyCustomerSessionID string = "journeyCustomerSessionId"

	// ConversationAggregateQueryPredicateDimensionJourneyCustomerSessionIDType captures enum value "journeyCustomerSessionIdType"
	ConversationAggregateQueryPredicateDimensionJourneyCustomerSessionIDType string = "journeyCustomerSessionIdType"

	// ConversationAggregateQueryPredicateDimensionKnowledgeBaseID captures enum value "knowledgeBaseId"
	ConversationAggregateQueryPredicateDimensionKnowledgeBaseID string = "knowledgeBaseId"

	// ConversationAggregateQueryPredicateDimensionMediaCount captures enum value "mediaCount"
	ConversationAggregateQueryPredicateDimensionMediaCount string = "mediaCount"

	// ConversationAggregateQueryPredicateDimensionMediaType captures enum value "mediaType"
	ConversationAggregateQueryPredicateDimensionMediaType string = "mediaType"

	// ConversationAggregateQueryPredicateDimensionMessageType captures enum value "messageType"
	ConversationAggregateQueryPredicateDimensionMessageType string = "messageType"

	// ConversationAggregateQueryPredicateDimensionOriginatingDirection captures enum value "originatingDirection"
	ConversationAggregateQueryPredicateDimensionOriginatingDirection string = "originatingDirection"

	// ConversationAggregateQueryPredicateDimensionOutboundCampaignID captures enum value "outboundCampaignId"
	ConversationAggregateQueryPredicateDimensionOutboundCampaignID string = "outboundCampaignId"

	// ConversationAggregateQueryPredicateDimensionOutboundContactID captures enum value "outboundContactId"
	ConversationAggregateQueryPredicateDimensionOutboundContactID string = "outboundContactId"

	// ConversationAggregateQueryPredicateDimensionOutboundContactListID captures enum value "outboundContactListId"
	ConversationAggregateQueryPredicateDimensionOutboundContactListID string = "outboundContactListId"

	// ConversationAggregateQueryPredicateDimensionParticipantName captures enum value "participantName"
	ConversationAggregateQueryPredicateDimensionParticipantName string = "participantName"

	// ConversationAggregateQueryPredicateDimensionPeerID captures enum value "peerId"
	ConversationAggregateQueryPredicateDimensionPeerID string = "peerId"

	// ConversationAggregateQueryPredicateDimensionProposedAgentID captures enum value "proposedAgentId"
	ConversationAggregateQueryPredicateDimensionProposedAgentID string = "proposedAgentId"

	// ConversationAggregateQueryPredicateDimensionProvider captures enum value "provider"
	ConversationAggregateQueryPredicateDimensionProvider string = "provider"

	// ConversationAggregateQueryPredicateDimensionPurpose captures enum value "purpose"
	ConversationAggregateQueryPredicateDimensionPurpose string = "purpose"

	// ConversationAggregateQueryPredicateDimensionQueueID captures enum value "queueId"
	ConversationAggregateQueryPredicateDimensionQueueID string = "queueId"

	// ConversationAggregateQueryPredicateDimensionRemote captures enum value "remote"
	ConversationAggregateQueryPredicateDimensionRemote string = "remote"

	// ConversationAggregateQueryPredicateDimensionRemovedSkillID captures enum value "removedSkillId"
	ConversationAggregateQueryPredicateDimensionRemovedSkillID string = "removedSkillId"

	// ConversationAggregateQueryPredicateDimensionReoffered captures enum value "reoffered"
	ConversationAggregateQueryPredicateDimensionReoffered string = "reoffered"

	// ConversationAggregateQueryPredicateDimensionRequestedLanguageID captures enum value "requestedLanguageId"
	ConversationAggregateQueryPredicateDimensionRequestedLanguageID string = "requestedLanguageId"

	// ConversationAggregateQueryPredicateDimensionRequestedRouting captures enum value "requestedRouting"
	ConversationAggregateQueryPredicateDimensionRequestedRouting string = "requestedRouting"

	// ConversationAggregateQueryPredicateDimensionRequestedRoutingSkillID captures enum value "requestedRoutingSkillId"
	ConversationAggregateQueryPredicateDimensionRequestedRoutingSkillID string = "requestedRoutingSkillId"

	// ConversationAggregateQueryPredicateDimensionRoomID captures enum value "roomId"
	ConversationAggregateQueryPredicateDimensionRoomID string = "roomId"

	// ConversationAggregateQueryPredicateDimensionRoutingPriority captures enum value "routingPriority"
	ConversationAggregateQueryPredicateDimensionRoutingPriority string = "routingPriority"

	// ConversationAggregateQueryPredicateDimensionRoutingRing captures enum value "routingRing"
	ConversationAggregateQueryPredicateDimensionRoutingRing string = "routingRing"

	// ConversationAggregateQueryPredicateDimensionScoredAgentID captures enum value "scoredAgentId"
	ConversationAggregateQueryPredicateDimensionScoredAgentID string = "scoredAgentId"

	// ConversationAggregateQueryPredicateDimensionSelectedAgentID captures enum value "selectedAgentId"
	ConversationAggregateQueryPredicateDimensionSelectedAgentID string = "selectedAgentId"

	// ConversationAggregateQueryPredicateDimensionSelectedAgentRank captures enum value "selectedAgentRank"
	ConversationAggregateQueryPredicateDimensionSelectedAgentRank string = "selectedAgentRank"

	// ConversationAggregateQueryPredicateDimensionSelfServed captures enum value "selfServed"
	ConversationAggregateQueryPredicateDimensionSelfServed string = "selfServed"

	// ConversationAggregateQueryPredicateDimensionSessionDnis captures enum value "sessionDnis"
	ConversationAggregateQueryPredicateDimensionSessionDnis string = "sessionDnis"

	// ConversationAggregateQueryPredicateDimensionSessionID captures enum value "sessionId"
	ConversationAggregateQueryPredicateDimensionSessionID string = "sessionId"

	// ConversationAggregateQueryPredicateDimensionStationID captures enum value "stationId"
	ConversationAggregateQueryPredicateDimensionStationID string = "stationId"

	// ConversationAggregateQueryPredicateDimensionTeamID captures enum value "teamId"
	ConversationAggregateQueryPredicateDimensionTeamID string = "teamId"

	// ConversationAggregateQueryPredicateDimensionUsedRouting captures enum value "usedRouting"
	ConversationAggregateQueryPredicateDimensionUsedRouting string = "usedRouting"

	// ConversationAggregateQueryPredicateDimensionUserID captures enum value "userId"
	ConversationAggregateQueryPredicateDimensionUserID string = "userId"

	// ConversationAggregateQueryPredicateDimensionWaitingInteractionCount captures enum value "waitingInteractionCount"
	ConversationAggregateQueryPredicateDimensionWaitingInteractionCount string = "waitingInteractionCount"

	// ConversationAggregateQueryPredicateDimensionWrapUpCode captures enum value "wrapUpCode"
	ConversationAggregateQueryPredicateDimensionWrapUpCode string = "wrapUpCode"
)

// prop value enum
func (m *ConversationAggregateQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationAggregateQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationAggregateQueryPredicate) validateDimension(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var conversationAggregateQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationAggregateQueryPredicateTypeOperatorPropEnum = append(conversationAggregateQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// ConversationAggregateQueryPredicateOperatorMatches captures enum value "matches"
	ConversationAggregateQueryPredicateOperatorMatches string = "matches"

	// ConversationAggregateQueryPredicateOperatorExists captures enum value "exists"
	ConversationAggregateQueryPredicateOperatorExists string = "exists"

	// ConversationAggregateQueryPredicateOperatorNotExists captures enum value "notExists"
	ConversationAggregateQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *ConversationAggregateQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationAggregateQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationAggregateQueryPredicate) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *ConversationAggregateQueryPredicate) validateRange(formats strfmt.Registry) error {
	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

var conversationAggregateQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationAggregateQueryPredicateTypeTypePropEnum = append(conversationAggregateQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// ConversationAggregateQueryPredicateTypeDimension captures enum value "dimension"
	ConversationAggregateQueryPredicateTypeDimension string = "dimension"

	// ConversationAggregateQueryPredicateTypeProperty captures enum value "property"
	ConversationAggregateQueryPredicateTypeProperty string = "property"

	// ConversationAggregateQueryPredicateTypeMetric captures enum value "metric"
	ConversationAggregateQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *ConversationAggregateQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationAggregateQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationAggregateQueryPredicate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this conversation aggregate query predicate based on the context it is used
func (m *ConversationAggregateQueryPredicate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationAggregateQueryPredicate) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

	if m.Range != nil {
		if err := m.Range.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("range")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationAggregateQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationAggregateQueryPredicate) UnmarshalBinary(b []byte) error {
	var res ConversationAggregateQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
