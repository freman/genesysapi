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

// SegmentDetailQueryPredicate segment detail query predicate
//
// swagger:model SegmentDetailQueryPredicate
type SegmentDetailQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [addressFrom addressTo agentAssistantId agentOwned ani authenticated bargedParticipantId callbackNumber callbackScheduledTime coachedParticipantId conference deliveryStatus destinationAddress destinationConversationId direction disconnectType dnis edgeId errorCode exitReason extendedDeliveryStatus externalContactId externalOrganizationId flaggedReason flowId flowName flowOutType flowOutcome flowOutcomeId flowOutcomeValue flowVersion groupId journeyActionId journeyActionMapId journeyCustomerId journeyCustomerIdType journeyCustomerSessionId mediaCount mediaType messageType monitoredParticipantId outboundCampaignId outboundContactId outboundContactListId participantName protocolCallId provider purpose queueId recording remote remoteNameDisplayable requestedLanguageId requestedRouting requestedRoutingSkillId scoredAgentId scriptId segmentEnd segmentType sessionDnis sipResponseCode subject teamId transferTargetAddress transferTargetName transferType usedRouting userId wrapUpCode wrapUpNote]
	Dimension string `json:"dimension,omitempty"`

	// Left hand side for metric predicates
	// Enum: [tSegmentDuration]
	Metric string `json:"metric,omitempty"`

	// Optional operator, default is matches
	// Enum: [matches exists notExists]
	Operator string `json:"operator,omitempty"`

	// Left hand side for property predicates
	Property string `json:"property,omitempty"`

	// Left hand side for property predicates
	// Enum: [bool integer real date string uuid]
	PropertyType string `json:"propertyType,omitempty"`

	// Right hand side for dimension, metric, or property predicates
	Range *NumericRange `json:"range,omitempty"`

	// Optional type, can usually be inferred
	// Enum: [dimension property metric]
	Type string `json:"type,omitempty"`

	// Right hand side for dimension, metric, or property predicates
	Value string `json:"value,omitempty"`
}

// Validate validates this segment detail query predicate
func (m *SegmentDetailQueryPredicate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDimension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyType(formats); err != nil {
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

var segmentDetailQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["addressFrom","addressTo","agentAssistantId","agentOwned","ani","authenticated","bargedParticipantId","callbackNumber","callbackScheduledTime","coachedParticipantId","conference","deliveryStatus","destinationAddress","destinationConversationId","direction","disconnectType","dnis","edgeId","errorCode","exitReason","extendedDeliveryStatus","externalContactId","externalOrganizationId","flaggedReason","flowId","flowName","flowOutType","flowOutcome","flowOutcomeId","flowOutcomeValue","flowVersion","groupId","journeyActionId","journeyActionMapId","journeyCustomerId","journeyCustomerIdType","journeyCustomerSessionId","mediaCount","mediaType","messageType","monitoredParticipantId","outboundCampaignId","outboundContactId","outboundContactListId","participantName","protocolCallId","provider","purpose","queueId","recording","remote","remoteNameDisplayable","requestedLanguageId","requestedRouting","requestedRoutingSkillId","scoredAgentId","scriptId","segmentEnd","segmentType","sessionDnis","sipResponseCode","subject","teamId","transferTargetAddress","transferTargetName","transferType","usedRouting","userId","wrapUpCode","wrapUpNote"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		segmentDetailQueryPredicateTypeDimensionPropEnum = append(segmentDetailQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// SegmentDetailQueryPredicateDimensionAddressFrom captures enum value "addressFrom"
	SegmentDetailQueryPredicateDimensionAddressFrom string = "addressFrom"

	// SegmentDetailQueryPredicateDimensionAddressTo captures enum value "addressTo"
	SegmentDetailQueryPredicateDimensionAddressTo string = "addressTo"

	// SegmentDetailQueryPredicateDimensionAgentAssistantID captures enum value "agentAssistantId"
	SegmentDetailQueryPredicateDimensionAgentAssistantID string = "agentAssistantId"

	// SegmentDetailQueryPredicateDimensionAgentOwned captures enum value "agentOwned"
	SegmentDetailQueryPredicateDimensionAgentOwned string = "agentOwned"

	// SegmentDetailQueryPredicateDimensionAni captures enum value "ani"
	SegmentDetailQueryPredicateDimensionAni string = "ani"

	// SegmentDetailQueryPredicateDimensionAuthenticated captures enum value "authenticated"
	SegmentDetailQueryPredicateDimensionAuthenticated string = "authenticated"

	// SegmentDetailQueryPredicateDimensionBargedParticipantID captures enum value "bargedParticipantId"
	SegmentDetailQueryPredicateDimensionBargedParticipantID string = "bargedParticipantId"

	// SegmentDetailQueryPredicateDimensionCallbackNumber captures enum value "callbackNumber"
	SegmentDetailQueryPredicateDimensionCallbackNumber string = "callbackNumber"

	// SegmentDetailQueryPredicateDimensionCallbackScheduledTime captures enum value "callbackScheduledTime"
	SegmentDetailQueryPredicateDimensionCallbackScheduledTime string = "callbackScheduledTime"

	// SegmentDetailQueryPredicateDimensionCoachedParticipantID captures enum value "coachedParticipantId"
	SegmentDetailQueryPredicateDimensionCoachedParticipantID string = "coachedParticipantId"

	// SegmentDetailQueryPredicateDimensionConference captures enum value "conference"
	SegmentDetailQueryPredicateDimensionConference string = "conference"

	// SegmentDetailQueryPredicateDimensionDeliveryStatus captures enum value "deliveryStatus"
	SegmentDetailQueryPredicateDimensionDeliveryStatus string = "deliveryStatus"

	// SegmentDetailQueryPredicateDimensionDestinationAddress captures enum value "destinationAddress"
	SegmentDetailQueryPredicateDimensionDestinationAddress string = "destinationAddress"

	// SegmentDetailQueryPredicateDimensionDestinationConversationID captures enum value "destinationConversationId"
	SegmentDetailQueryPredicateDimensionDestinationConversationID string = "destinationConversationId"

	// SegmentDetailQueryPredicateDimensionDirection captures enum value "direction"
	SegmentDetailQueryPredicateDimensionDirection string = "direction"

	// SegmentDetailQueryPredicateDimensionDisconnectType captures enum value "disconnectType"
	SegmentDetailQueryPredicateDimensionDisconnectType string = "disconnectType"

	// SegmentDetailQueryPredicateDimensionDnis captures enum value "dnis"
	SegmentDetailQueryPredicateDimensionDnis string = "dnis"

	// SegmentDetailQueryPredicateDimensionEdgeID captures enum value "edgeId"
	SegmentDetailQueryPredicateDimensionEdgeID string = "edgeId"

	// SegmentDetailQueryPredicateDimensionErrorCode captures enum value "errorCode"
	SegmentDetailQueryPredicateDimensionErrorCode string = "errorCode"

	// SegmentDetailQueryPredicateDimensionExitReason captures enum value "exitReason"
	SegmentDetailQueryPredicateDimensionExitReason string = "exitReason"

	// SegmentDetailQueryPredicateDimensionExtendedDeliveryStatus captures enum value "extendedDeliveryStatus"
	SegmentDetailQueryPredicateDimensionExtendedDeliveryStatus string = "extendedDeliveryStatus"

	// SegmentDetailQueryPredicateDimensionExternalContactID captures enum value "externalContactId"
	SegmentDetailQueryPredicateDimensionExternalContactID string = "externalContactId"

	// SegmentDetailQueryPredicateDimensionExternalOrganizationID captures enum value "externalOrganizationId"
	SegmentDetailQueryPredicateDimensionExternalOrganizationID string = "externalOrganizationId"

	// SegmentDetailQueryPredicateDimensionFlaggedReason captures enum value "flaggedReason"
	SegmentDetailQueryPredicateDimensionFlaggedReason string = "flaggedReason"

	// SegmentDetailQueryPredicateDimensionFlowID captures enum value "flowId"
	SegmentDetailQueryPredicateDimensionFlowID string = "flowId"

	// SegmentDetailQueryPredicateDimensionFlowName captures enum value "flowName"
	SegmentDetailQueryPredicateDimensionFlowName string = "flowName"

	// SegmentDetailQueryPredicateDimensionFlowOutType captures enum value "flowOutType"
	SegmentDetailQueryPredicateDimensionFlowOutType string = "flowOutType"

	// SegmentDetailQueryPredicateDimensionFlowOutcome captures enum value "flowOutcome"
	SegmentDetailQueryPredicateDimensionFlowOutcome string = "flowOutcome"

	// SegmentDetailQueryPredicateDimensionFlowOutcomeID captures enum value "flowOutcomeId"
	SegmentDetailQueryPredicateDimensionFlowOutcomeID string = "flowOutcomeId"

	// SegmentDetailQueryPredicateDimensionFlowOutcomeValue captures enum value "flowOutcomeValue"
	SegmentDetailQueryPredicateDimensionFlowOutcomeValue string = "flowOutcomeValue"

	// SegmentDetailQueryPredicateDimensionFlowVersion captures enum value "flowVersion"
	SegmentDetailQueryPredicateDimensionFlowVersion string = "flowVersion"

	// SegmentDetailQueryPredicateDimensionGroupID captures enum value "groupId"
	SegmentDetailQueryPredicateDimensionGroupID string = "groupId"

	// SegmentDetailQueryPredicateDimensionJourneyActionID captures enum value "journeyActionId"
	SegmentDetailQueryPredicateDimensionJourneyActionID string = "journeyActionId"

	// SegmentDetailQueryPredicateDimensionJourneyActionMapID captures enum value "journeyActionMapId"
	SegmentDetailQueryPredicateDimensionJourneyActionMapID string = "journeyActionMapId"

	// SegmentDetailQueryPredicateDimensionJourneyCustomerID captures enum value "journeyCustomerId"
	SegmentDetailQueryPredicateDimensionJourneyCustomerID string = "journeyCustomerId"

	// SegmentDetailQueryPredicateDimensionJourneyCustomerIDType captures enum value "journeyCustomerIdType"
	SegmentDetailQueryPredicateDimensionJourneyCustomerIDType string = "journeyCustomerIdType"

	// SegmentDetailQueryPredicateDimensionJourneyCustomerSessionID captures enum value "journeyCustomerSessionId"
	SegmentDetailQueryPredicateDimensionJourneyCustomerSessionID string = "journeyCustomerSessionId"

	// SegmentDetailQueryPredicateDimensionMediaCount captures enum value "mediaCount"
	SegmentDetailQueryPredicateDimensionMediaCount string = "mediaCount"

	// SegmentDetailQueryPredicateDimensionMediaType captures enum value "mediaType"
	SegmentDetailQueryPredicateDimensionMediaType string = "mediaType"

	// SegmentDetailQueryPredicateDimensionMessageType captures enum value "messageType"
	SegmentDetailQueryPredicateDimensionMessageType string = "messageType"

	// SegmentDetailQueryPredicateDimensionMonitoredParticipantID captures enum value "monitoredParticipantId"
	SegmentDetailQueryPredicateDimensionMonitoredParticipantID string = "monitoredParticipantId"

	// SegmentDetailQueryPredicateDimensionOutboundCampaignID captures enum value "outboundCampaignId"
	SegmentDetailQueryPredicateDimensionOutboundCampaignID string = "outboundCampaignId"

	// SegmentDetailQueryPredicateDimensionOutboundContactID captures enum value "outboundContactId"
	SegmentDetailQueryPredicateDimensionOutboundContactID string = "outboundContactId"

	// SegmentDetailQueryPredicateDimensionOutboundContactListID captures enum value "outboundContactListId"
	SegmentDetailQueryPredicateDimensionOutboundContactListID string = "outboundContactListId"

	// SegmentDetailQueryPredicateDimensionParticipantName captures enum value "participantName"
	SegmentDetailQueryPredicateDimensionParticipantName string = "participantName"

	// SegmentDetailQueryPredicateDimensionProtocolCallID captures enum value "protocolCallId"
	SegmentDetailQueryPredicateDimensionProtocolCallID string = "protocolCallId"

	// SegmentDetailQueryPredicateDimensionProvider captures enum value "provider"
	SegmentDetailQueryPredicateDimensionProvider string = "provider"

	// SegmentDetailQueryPredicateDimensionPurpose captures enum value "purpose"
	SegmentDetailQueryPredicateDimensionPurpose string = "purpose"

	// SegmentDetailQueryPredicateDimensionQueueID captures enum value "queueId"
	SegmentDetailQueryPredicateDimensionQueueID string = "queueId"

	// SegmentDetailQueryPredicateDimensionRecording captures enum value "recording"
	SegmentDetailQueryPredicateDimensionRecording string = "recording"

	// SegmentDetailQueryPredicateDimensionRemote captures enum value "remote"
	SegmentDetailQueryPredicateDimensionRemote string = "remote"

	// SegmentDetailQueryPredicateDimensionRemoteNameDisplayable captures enum value "remoteNameDisplayable"
	SegmentDetailQueryPredicateDimensionRemoteNameDisplayable string = "remoteNameDisplayable"

	// SegmentDetailQueryPredicateDimensionRequestedLanguageID captures enum value "requestedLanguageId"
	SegmentDetailQueryPredicateDimensionRequestedLanguageID string = "requestedLanguageId"

	// SegmentDetailQueryPredicateDimensionRequestedRouting captures enum value "requestedRouting"
	SegmentDetailQueryPredicateDimensionRequestedRouting string = "requestedRouting"

	// SegmentDetailQueryPredicateDimensionRequestedRoutingSkillID captures enum value "requestedRoutingSkillId"
	SegmentDetailQueryPredicateDimensionRequestedRoutingSkillID string = "requestedRoutingSkillId"

	// SegmentDetailQueryPredicateDimensionScoredAgentID captures enum value "scoredAgentId"
	SegmentDetailQueryPredicateDimensionScoredAgentID string = "scoredAgentId"

	// SegmentDetailQueryPredicateDimensionScriptID captures enum value "scriptId"
	SegmentDetailQueryPredicateDimensionScriptID string = "scriptId"

	// SegmentDetailQueryPredicateDimensionSegmentEnd captures enum value "segmentEnd"
	SegmentDetailQueryPredicateDimensionSegmentEnd string = "segmentEnd"

	// SegmentDetailQueryPredicateDimensionSegmentType captures enum value "segmentType"
	SegmentDetailQueryPredicateDimensionSegmentType string = "segmentType"

	// SegmentDetailQueryPredicateDimensionSessionDnis captures enum value "sessionDnis"
	SegmentDetailQueryPredicateDimensionSessionDnis string = "sessionDnis"

	// SegmentDetailQueryPredicateDimensionSipResponseCode captures enum value "sipResponseCode"
	SegmentDetailQueryPredicateDimensionSipResponseCode string = "sipResponseCode"

	// SegmentDetailQueryPredicateDimensionSubject captures enum value "subject"
	SegmentDetailQueryPredicateDimensionSubject string = "subject"

	// SegmentDetailQueryPredicateDimensionTeamID captures enum value "teamId"
	SegmentDetailQueryPredicateDimensionTeamID string = "teamId"

	// SegmentDetailQueryPredicateDimensionTransferTargetAddress captures enum value "transferTargetAddress"
	SegmentDetailQueryPredicateDimensionTransferTargetAddress string = "transferTargetAddress"

	// SegmentDetailQueryPredicateDimensionTransferTargetName captures enum value "transferTargetName"
	SegmentDetailQueryPredicateDimensionTransferTargetName string = "transferTargetName"

	// SegmentDetailQueryPredicateDimensionTransferType captures enum value "transferType"
	SegmentDetailQueryPredicateDimensionTransferType string = "transferType"

	// SegmentDetailQueryPredicateDimensionUsedRouting captures enum value "usedRouting"
	SegmentDetailQueryPredicateDimensionUsedRouting string = "usedRouting"

	// SegmentDetailQueryPredicateDimensionUserID captures enum value "userId"
	SegmentDetailQueryPredicateDimensionUserID string = "userId"

	// SegmentDetailQueryPredicateDimensionWrapUpCode captures enum value "wrapUpCode"
	SegmentDetailQueryPredicateDimensionWrapUpCode string = "wrapUpCode"

	// SegmentDetailQueryPredicateDimensionWrapUpNote captures enum value "wrapUpNote"
	SegmentDetailQueryPredicateDimensionWrapUpNote string = "wrapUpNote"
)

// prop value enum
func (m *SegmentDetailQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, segmentDetailQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SegmentDetailQueryPredicate) validateDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var segmentDetailQueryPredicateTypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tSegmentDuration"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		segmentDetailQueryPredicateTypeMetricPropEnum = append(segmentDetailQueryPredicateTypeMetricPropEnum, v)
	}
}

const (

	// SegmentDetailQueryPredicateMetricTSegmentDuration captures enum value "tSegmentDuration"
	SegmentDetailQueryPredicateMetricTSegmentDuration string = "tSegmentDuration"
)

// prop value enum
func (m *SegmentDetailQueryPredicate) validateMetricEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, segmentDetailQueryPredicateTypeMetricPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SegmentDetailQueryPredicate) validateMetric(formats strfmt.Registry) error {

	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	// value enum
	if err := m.validateMetricEnum("metric", "body", m.Metric); err != nil {
		return err
	}

	return nil
}

var segmentDetailQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		segmentDetailQueryPredicateTypeOperatorPropEnum = append(segmentDetailQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// SegmentDetailQueryPredicateOperatorMatches captures enum value "matches"
	SegmentDetailQueryPredicateOperatorMatches string = "matches"

	// SegmentDetailQueryPredicateOperatorExists captures enum value "exists"
	SegmentDetailQueryPredicateOperatorExists string = "exists"

	// SegmentDetailQueryPredicateOperatorNotExists captures enum value "notExists"
	SegmentDetailQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *SegmentDetailQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, segmentDetailQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SegmentDetailQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

var segmentDetailQueryPredicateTypePropertyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bool","integer","real","date","string","uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		segmentDetailQueryPredicateTypePropertyTypePropEnum = append(segmentDetailQueryPredicateTypePropertyTypePropEnum, v)
	}
}

const (

	// SegmentDetailQueryPredicatePropertyTypeBool captures enum value "bool"
	SegmentDetailQueryPredicatePropertyTypeBool string = "bool"

	// SegmentDetailQueryPredicatePropertyTypeInteger captures enum value "integer"
	SegmentDetailQueryPredicatePropertyTypeInteger string = "integer"

	// SegmentDetailQueryPredicatePropertyTypeReal captures enum value "real"
	SegmentDetailQueryPredicatePropertyTypeReal string = "real"

	// SegmentDetailQueryPredicatePropertyTypeDate captures enum value "date"
	SegmentDetailQueryPredicatePropertyTypeDate string = "date"

	// SegmentDetailQueryPredicatePropertyTypeString captures enum value "string"
	SegmentDetailQueryPredicatePropertyTypeString string = "string"

	// SegmentDetailQueryPredicatePropertyTypeUUID captures enum value "uuid"
	SegmentDetailQueryPredicatePropertyTypeUUID string = "uuid"
)

// prop value enum
func (m *SegmentDetailQueryPredicate) validatePropertyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, segmentDetailQueryPredicateTypePropertyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SegmentDetailQueryPredicate) validatePropertyType(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePropertyTypeEnum("propertyType", "body", m.PropertyType); err != nil {
		return err
	}

	return nil
}

func (m *SegmentDetailQueryPredicate) validateRange(formats strfmt.Registry) error {

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

var segmentDetailQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		segmentDetailQueryPredicateTypeTypePropEnum = append(segmentDetailQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// SegmentDetailQueryPredicateTypeDimension captures enum value "dimension"
	SegmentDetailQueryPredicateTypeDimension string = "dimension"

	// SegmentDetailQueryPredicateTypeProperty captures enum value "property"
	SegmentDetailQueryPredicateTypeProperty string = "property"

	// SegmentDetailQueryPredicateTypeMetric captures enum value "metric"
	SegmentDetailQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *SegmentDetailQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, segmentDetailQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SegmentDetailQueryPredicate) validateType(formats strfmt.Registry) error {

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
func (m *SegmentDetailQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentDetailQueryPredicate) UnmarshalBinary(b []byte) error {
	var res SegmentDetailQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
