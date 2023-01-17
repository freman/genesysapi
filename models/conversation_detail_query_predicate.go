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

// ConversationDetailQueryPredicate conversation detail query predicate
//
// swagger:model ConversationDetailQueryPredicate
type ConversationDetailQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [conversationEnd conversationId conversationInitiator conversationStart customerParticipation divisionId externalTag mediaStatsMinConversationMos originatingDirection]
	Dimension string `json:"dimension,omitempty"`

	// Left hand side for metric predicates
	// Enum: [nBlindTransferred nCobrowseSessions nConnected nConsult nConsultTransferred nError nFlow nFlowMilestone nFlowOutcome nFlowOutcomeFailed nOffered nOutbound nOutboundAbandoned nOutboundAttempted nOutboundConnected nOverSla nStateTransitionError nTransferred oExternalMediaCount oFlowMilestone oMediaCount oMessageTurn tAbandon tAcd tAcw tAgentResponseTime tAlert tAnswered tBarging tCallback tCallbackComplete tCoaching tCoachingComplete tConnected tContacting tConversationDuration tDialing tFirstConnect tFirstDial tFlow tFlowDisconnect tFlowExit tFlowOut tFlowOutcome tHandle tHeld tHeldComplete tIvr tMonitoring tMonitoringComplete tNotResponding tShortAbandon tTalk tTalkComplete tUserResponseTime tVoicemail]
	Metric string `json:"metric,omitempty"`

	// Optional operator, default is matches
	// Enum: [matches exists notExists]
	Operator string `json:"operator,omitempty"`

	// Right hand side for dimension or metric predicates
	Range *NumericRange `json:"range,omitempty"`

	// Optional type, can usually be inferred
	// Enum: [dimension property metric]
	Type string `json:"type,omitempty"`

	// Right hand side for dimension or metric predicates
	Value string `json:"value,omitempty"`
}

// Validate validates this conversation detail query predicate
func (m *ConversationDetailQueryPredicate) Validate(formats strfmt.Registry) error {
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

var conversationDetailQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["conversationEnd","conversationId","conversationInitiator","conversationStart","customerParticipation","divisionId","externalTag","mediaStatsMinConversationMos","originatingDirection"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationDetailQueryPredicateTypeDimensionPropEnum = append(conversationDetailQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// ConversationDetailQueryPredicateDimensionConversationEnd captures enum value "conversationEnd"
	ConversationDetailQueryPredicateDimensionConversationEnd string = "conversationEnd"

	// ConversationDetailQueryPredicateDimensionConversationID captures enum value "conversationId"
	ConversationDetailQueryPredicateDimensionConversationID string = "conversationId"

	// ConversationDetailQueryPredicateDimensionConversationInitiator captures enum value "conversationInitiator"
	ConversationDetailQueryPredicateDimensionConversationInitiator string = "conversationInitiator"

	// ConversationDetailQueryPredicateDimensionConversationStart captures enum value "conversationStart"
	ConversationDetailQueryPredicateDimensionConversationStart string = "conversationStart"

	// ConversationDetailQueryPredicateDimensionCustomerParticipation captures enum value "customerParticipation"
	ConversationDetailQueryPredicateDimensionCustomerParticipation string = "customerParticipation"

	// ConversationDetailQueryPredicateDimensionDivisionID captures enum value "divisionId"
	ConversationDetailQueryPredicateDimensionDivisionID string = "divisionId"

	// ConversationDetailQueryPredicateDimensionExternalTag captures enum value "externalTag"
	ConversationDetailQueryPredicateDimensionExternalTag string = "externalTag"

	// ConversationDetailQueryPredicateDimensionMediaStatsMinConversationMos captures enum value "mediaStatsMinConversationMos"
	ConversationDetailQueryPredicateDimensionMediaStatsMinConversationMos string = "mediaStatsMinConversationMos"

	// ConversationDetailQueryPredicateDimensionOriginatingDirection captures enum value "originatingDirection"
	ConversationDetailQueryPredicateDimensionOriginatingDirection string = "originatingDirection"
)

// prop value enum
func (m *ConversationDetailQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationDetailQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationDetailQueryPredicate) validateDimension(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var conversationDetailQueryPredicateTypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nBlindTransferred","nCobrowseSessions","nConnected","nConsult","nConsultTransferred","nError","nFlow","nFlowMilestone","nFlowOutcome","nFlowOutcomeFailed","nOffered","nOutbound","nOutboundAbandoned","nOutboundAttempted","nOutboundConnected","nOverSla","nStateTransitionError","nTransferred","oExternalMediaCount","oFlowMilestone","oMediaCount","oMessageTurn","tAbandon","tAcd","tAcw","tAgentResponseTime","tAlert","tAnswered","tBarging","tCallback","tCallbackComplete","tCoaching","tCoachingComplete","tConnected","tContacting","tConversationDuration","tDialing","tFirstConnect","tFirstDial","tFlow","tFlowDisconnect","tFlowExit","tFlowOut","tFlowOutcome","tHandle","tHeld","tHeldComplete","tIvr","tMonitoring","tMonitoringComplete","tNotResponding","tShortAbandon","tTalk","tTalkComplete","tUserResponseTime","tVoicemail"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationDetailQueryPredicateTypeMetricPropEnum = append(conversationDetailQueryPredicateTypeMetricPropEnum, v)
	}
}

const (

	// ConversationDetailQueryPredicateMetricNBlindTransferred captures enum value "nBlindTransferred"
	ConversationDetailQueryPredicateMetricNBlindTransferred string = "nBlindTransferred"

	// ConversationDetailQueryPredicateMetricNCobrowseSessions captures enum value "nCobrowseSessions"
	ConversationDetailQueryPredicateMetricNCobrowseSessions string = "nCobrowseSessions"

	// ConversationDetailQueryPredicateMetricNConnected captures enum value "nConnected"
	ConversationDetailQueryPredicateMetricNConnected string = "nConnected"

	// ConversationDetailQueryPredicateMetricNConsult captures enum value "nConsult"
	ConversationDetailQueryPredicateMetricNConsult string = "nConsult"

	// ConversationDetailQueryPredicateMetricNConsultTransferred captures enum value "nConsultTransferred"
	ConversationDetailQueryPredicateMetricNConsultTransferred string = "nConsultTransferred"

	// ConversationDetailQueryPredicateMetricNError captures enum value "nError"
	ConversationDetailQueryPredicateMetricNError string = "nError"

	// ConversationDetailQueryPredicateMetricNFlow captures enum value "nFlow"
	ConversationDetailQueryPredicateMetricNFlow string = "nFlow"

	// ConversationDetailQueryPredicateMetricNFlowMilestone captures enum value "nFlowMilestone"
	ConversationDetailQueryPredicateMetricNFlowMilestone string = "nFlowMilestone"

	// ConversationDetailQueryPredicateMetricNFlowOutcome captures enum value "nFlowOutcome"
	ConversationDetailQueryPredicateMetricNFlowOutcome string = "nFlowOutcome"

	// ConversationDetailQueryPredicateMetricNFlowOutcomeFailed captures enum value "nFlowOutcomeFailed"
	ConversationDetailQueryPredicateMetricNFlowOutcomeFailed string = "nFlowOutcomeFailed"

	// ConversationDetailQueryPredicateMetricNOffered captures enum value "nOffered"
	ConversationDetailQueryPredicateMetricNOffered string = "nOffered"

	// ConversationDetailQueryPredicateMetricNOutbound captures enum value "nOutbound"
	ConversationDetailQueryPredicateMetricNOutbound string = "nOutbound"

	// ConversationDetailQueryPredicateMetricNOutboundAbandoned captures enum value "nOutboundAbandoned"
	ConversationDetailQueryPredicateMetricNOutboundAbandoned string = "nOutboundAbandoned"

	// ConversationDetailQueryPredicateMetricNOutboundAttempted captures enum value "nOutboundAttempted"
	ConversationDetailQueryPredicateMetricNOutboundAttempted string = "nOutboundAttempted"

	// ConversationDetailQueryPredicateMetricNOutboundConnected captures enum value "nOutboundConnected"
	ConversationDetailQueryPredicateMetricNOutboundConnected string = "nOutboundConnected"

	// ConversationDetailQueryPredicateMetricNOverSLA captures enum value "nOverSla"
	ConversationDetailQueryPredicateMetricNOverSLA string = "nOverSla"

	// ConversationDetailQueryPredicateMetricNStateTransitionError captures enum value "nStateTransitionError"
	ConversationDetailQueryPredicateMetricNStateTransitionError string = "nStateTransitionError"

	// ConversationDetailQueryPredicateMetricNTransferred captures enum value "nTransferred"
	ConversationDetailQueryPredicateMetricNTransferred string = "nTransferred"

	// ConversationDetailQueryPredicateMetricOExternalMediaCount captures enum value "oExternalMediaCount"
	ConversationDetailQueryPredicateMetricOExternalMediaCount string = "oExternalMediaCount"

	// ConversationDetailQueryPredicateMetricOFlowMilestone captures enum value "oFlowMilestone"
	ConversationDetailQueryPredicateMetricOFlowMilestone string = "oFlowMilestone"

	// ConversationDetailQueryPredicateMetricOMediaCount captures enum value "oMediaCount"
	ConversationDetailQueryPredicateMetricOMediaCount string = "oMediaCount"

	// ConversationDetailQueryPredicateMetricOMessageTurn captures enum value "oMessageTurn"
	ConversationDetailQueryPredicateMetricOMessageTurn string = "oMessageTurn"

	// ConversationDetailQueryPredicateMetricTAbandon captures enum value "tAbandon"
	ConversationDetailQueryPredicateMetricTAbandon string = "tAbandon"

	// ConversationDetailQueryPredicateMetricTAcd captures enum value "tAcd"
	ConversationDetailQueryPredicateMetricTAcd string = "tAcd"

	// ConversationDetailQueryPredicateMetricTAcw captures enum value "tAcw"
	ConversationDetailQueryPredicateMetricTAcw string = "tAcw"

	// ConversationDetailQueryPredicateMetricTAgentResponseTime captures enum value "tAgentResponseTime"
	ConversationDetailQueryPredicateMetricTAgentResponseTime string = "tAgentResponseTime"

	// ConversationDetailQueryPredicateMetricTAlert captures enum value "tAlert"
	ConversationDetailQueryPredicateMetricTAlert string = "tAlert"

	// ConversationDetailQueryPredicateMetricTAnswered captures enum value "tAnswered"
	ConversationDetailQueryPredicateMetricTAnswered string = "tAnswered"

	// ConversationDetailQueryPredicateMetricTBarging captures enum value "tBarging"
	ConversationDetailQueryPredicateMetricTBarging string = "tBarging"

	// ConversationDetailQueryPredicateMetricTCallback captures enum value "tCallback"
	ConversationDetailQueryPredicateMetricTCallback string = "tCallback"

	// ConversationDetailQueryPredicateMetricTCallbackComplete captures enum value "tCallbackComplete"
	ConversationDetailQueryPredicateMetricTCallbackComplete string = "tCallbackComplete"

	// ConversationDetailQueryPredicateMetricTCoaching captures enum value "tCoaching"
	ConversationDetailQueryPredicateMetricTCoaching string = "tCoaching"

	// ConversationDetailQueryPredicateMetricTCoachingComplete captures enum value "tCoachingComplete"
	ConversationDetailQueryPredicateMetricTCoachingComplete string = "tCoachingComplete"

	// ConversationDetailQueryPredicateMetricTConnected captures enum value "tConnected"
	ConversationDetailQueryPredicateMetricTConnected string = "tConnected"

	// ConversationDetailQueryPredicateMetricTContacting captures enum value "tContacting"
	ConversationDetailQueryPredicateMetricTContacting string = "tContacting"

	// ConversationDetailQueryPredicateMetricTConversationDuration captures enum value "tConversationDuration"
	ConversationDetailQueryPredicateMetricTConversationDuration string = "tConversationDuration"

	// ConversationDetailQueryPredicateMetricTDialing captures enum value "tDialing"
	ConversationDetailQueryPredicateMetricTDialing string = "tDialing"

	// ConversationDetailQueryPredicateMetricTFirstConnect captures enum value "tFirstConnect"
	ConversationDetailQueryPredicateMetricTFirstConnect string = "tFirstConnect"

	// ConversationDetailQueryPredicateMetricTFirstDial captures enum value "tFirstDial"
	ConversationDetailQueryPredicateMetricTFirstDial string = "tFirstDial"

	// ConversationDetailQueryPredicateMetricTFlow captures enum value "tFlow"
	ConversationDetailQueryPredicateMetricTFlow string = "tFlow"

	// ConversationDetailQueryPredicateMetricTFlowDisconnect captures enum value "tFlowDisconnect"
	ConversationDetailQueryPredicateMetricTFlowDisconnect string = "tFlowDisconnect"

	// ConversationDetailQueryPredicateMetricTFlowExit captures enum value "tFlowExit"
	ConversationDetailQueryPredicateMetricTFlowExit string = "tFlowExit"

	// ConversationDetailQueryPredicateMetricTFlowOut captures enum value "tFlowOut"
	ConversationDetailQueryPredicateMetricTFlowOut string = "tFlowOut"

	// ConversationDetailQueryPredicateMetricTFlowOutcome captures enum value "tFlowOutcome"
	ConversationDetailQueryPredicateMetricTFlowOutcome string = "tFlowOutcome"

	// ConversationDetailQueryPredicateMetricTHandle captures enum value "tHandle"
	ConversationDetailQueryPredicateMetricTHandle string = "tHandle"

	// ConversationDetailQueryPredicateMetricTHeld captures enum value "tHeld"
	ConversationDetailQueryPredicateMetricTHeld string = "tHeld"

	// ConversationDetailQueryPredicateMetricTHeldComplete captures enum value "tHeldComplete"
	ConversationDetailQueryPredicateMetricTHeldComplete string = "tHeldComplete"

	// ConversationDetailQueryPredicateMetricTIvr captures enum value "tIvr"
	ConversationDetailQueryPredicateMetricTIvr string = "tIvr"

	// ConversationDetailQueryPredicateMetricTMonitoring captures enum value "tMonitoring"
	ConversationDetailQueryPredicateMetricTMonitoring string = "tMonitoring"

	// ConversationDetailQueryPredicateMetricTMonitoringComplete captures enum value "tMonitoringComplete"
	ConversationDetailQueryPredicateMetricTMonitoringComplete string = "tMonitoringComplete"

	// ConversationDetailQueryPredicateMetricTNotResponding captures enum value "tNotResponding"
	ConversationDetailQueryPredicateMetricTNotResponding string = "tNotResponding"

	// ConversationDetailQueryPredicateMetricTShortAbandon captures enum value "tShortAbandon"
	ConversationDetailQueryPredicateMetricTShortAbandon string = "tShortAbandon"

	// ConversationDetailQueryPredicateMetricTTalk captures enum value "tTalk"
	ConversationDetailQueryPredicateMetricTTalk string = "tTalk"

	// ConversationDetailQueryPredicateMetricTTalkComplete captures enum value "tTalkComplete"
	ConversationDetailQueryPredicateMetricTTalkComplete string = "tTalkComplete"

	// ConversationDetailQueryPredicateMetricTUserResponseTime captures enum value "tUserResponseTime"
	ConversationDetailQueryPredicateMetricTUserResponseTime string = "tUserResponseTime"

	// ConversationDetailQueryPredicateMetricTVoicemail captures enum value "tVoicemail"
	ConversationDetailQueryPredicateMetricTVoicemail string = "tVoicemail"
)

// prop value enum
func (m *ConversationDetailQueryPredicate) validateMetricEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationDetailQueryPredicateTypeMetricPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationDetailQueryPredicate) validateMetric(formats strfmt.Registry) error {
	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	// value enum
	if err := m.validateMetricEnum("metric", "body", m.Metric); err != nil {
		return err
	}

	return nil
}

var conversationDetailQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationDetailQueryPredicateTypeOperatorPropEnum = append(conversationDetailQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// ConversationDetailQueryPredicateOperatorMatches captures enum value "matches"
	ConversationDetailQueryPredicateOperatorMatches string = "matches"

	// ConversationDetailQueryPredicateOperatorExists captures enum value "exists"
	ConversationDetailQueryPredicateOperatorExists string = "exists"

	// ConversationDetailQueryPredicateOperatorNotExists captures enum value "notExists"
	ConversationDetailQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *ConversationDetailQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationDetailQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationDetailQueryPredicate) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *ConversationDetailQueryPredicate) validateRange(formats strfmt.Registry) error {
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

var conversationDetailQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationDetailQueryPredicateTypeTypePropEnum = append(conversationDetailQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// ConversationDetailQueryPredicateTypeDimension captures enum value "dimension"
	ConversationDetailQueryPredicateTypeDimension string = "dimension"

	// ConversationDetailQueryPredicateTypeProperty captures enum value "property"
	ConversationDetailQueryPredicateTypeProperty string = "property"

	// ConversationDetailQueryPredicateTypeMetric captures enum value "metric"
	ConversationDetailQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *ConversationDetailQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationDetailQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationDetailQueryPredicate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this conversation detail query predicate based on the context it is used
func (m *ConversationDetailQueryPredicate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationDetailQueryPredicate) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConversationDetailQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationDetailQueryPredicate) UnmarshalBinary(b []byte) error {
	var res ConversationDetailQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
