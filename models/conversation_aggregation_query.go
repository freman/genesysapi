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

// ConversationAggregationQuery conversation aggregation query
//
// swagger:model ConversationAggregationQuery
type ConversationAggregationQuery struct {

	// Dimension to use as the alternative timestamp for data in the aggregate.  Choosing "eventTime" uses the actual time of the data event.
	// Enum: [eventTime]
	AlternateTimeDimension string `json:"alternateTimeDimension,omitempty"`

	// Behaves like a SQL WHERE clause. This is ANDed with the interval parameter. Expresses boolean logical predicates as well as dimensional filters
	Filter *ConversationAggregateQueryFilter `json:"filter,omitempty"`

	// Flattens any multivalued dimensions used in response groups (e.g. ['a','b','c']->'a,b,c')
	FlattenMultivaluedDimensions bool `json:"flattenMultivaluedDimensions"`

	// Granularity aggregates metrics into subpartitions within the time interval specified. The default granularity is the same duration as the interval. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	Granularity string `json:"granularity,omitempty"`

	// Behaves like a SQL GROUPBY. Allows for multiple levels of grouping as a list of dimensions. Partitions resulting aggregate computations into distinct named subgroups rather than across the entire result set as if it were one group.
	GroupBy []string `json:"groupBy"`

	// Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	// Required: true
	Interval *string `json:"interval"`

	// Behaves like a SQL SELECT clause. Enables retrieving only named metrics. If omitted, all metrics that are available will be returned (like SELECT *).
	Metrics []string `json:"metrics"`

	// Time zone context used to calculate response intervals (this allows resolving DST changes). The interval offset is used even when timeZone is specified. Default is UTC. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	TimeZone string `json:"timeZone,omitempty"`

	// Custom derived metric views
	Views []*ConversationAggregationView `json:"views"`
}

// Validate validates this conversation aggregation query
func (m *ConversationAggregationQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternateTimeDimension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViews(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversationAggregationQueryTypeAlternateTimeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eventTime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationAggregationQueryTypeAlternateTimeDimensionPropEnum = append(conversationAggregationQueryTypeAlternateTimeDimensionPropEnum, v)
	}
}

const (

	// ConversationAggregationQueryAlternateTimeDimensionEventTime captures enum value "eventTime"
	ConversationAggregationQueryAlternateTimeDimensionEventTime string = "eventTime"
)

// prop value enum
func (m *ConversationAggregationQuery) validateAlternateTimeDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationAggregationQueryTypeAlternateTimeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationAggregationQuery) validateAlternateTimeDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.AlternateTimeDimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlternateTimeDimensionEnum("alternateTimeDimension", "body", m.AlternateTimeDimension); err != nil {
		return err
	}

	return nil
}

func (m *ConversationAggregationQuery) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

var conversationAggregationQueryGroupByItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["addressFrom","addressTo","agentAssistantId","agentScore","ani","conversationId","convertedFrom","convertedTo","direction","disconnectType","divisionId","dnis","edgeId","externalContactId","externalMediaCount","externalOrganizationId","flaggedReason","flowOutType","groupId","interactionType","journeyActionId","journeyActionMapId","journeyActionMapVersion","journeyCustomerId","journeyCustomerIdType","journeyCustomerSessionId","journeyCustomerSessionIdType","mediaCount","mediaType","messageType","originatingDirection","outboundCampaignId","outboundContactId","outboundContactListId","participantName","peerId","provider","purpose","queueId","remote","requestedLanguageId","requestedRouting","requestedRoutingSkillId","roomId","routingPriority","scoredAgentId","selectedAgentId","selectedAgentRank","sessionDnis","sessionId","stationId","teamId","usedRouting","userId","wrapUpCode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationAggregationQueryGroupByItemsEnum = append(conversationAggregationQueryGroupByItemsEnum, v)
	}
}

func (m *ConversationAggregationQuery) validateGroupByItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationAggregationQueryGroupByItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationAggregationQuery) validateGroupBy(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupBy) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupBy); i++ {

		// value enum
		if err := m.validateGroupByItemsEnum("groupBy"+"."+strconv.Itoa(i), "body", m.GroupBy[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConversationAggregationQuery) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

var conversationAggregationQueryMetricsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nBlindTransferred","nConnected","nConsult","nConsultTransferred","nError","nOffered","nOutbound","nOutboundAbandoned","nOutboundAttempted","nOutboundConnected","nOverSla","nStateTransitionError","nTransferred","oExternalMediaCount","oInteracting","oMediaCount","oServiceLevel","oServiceTarget","oWaiting","tAbandon","tAcd","tAcw","tAgentResponseTime","tAlert","tAnswered","tContacting","tDialing","tFlowOut","tHandle","tHeld","tHeldComplete","tIvr","tMonitoring","tNotResponding","tShortAbandon","tTalk","tTalkComplete","tUserResponseTime","tVoicemail","tWait"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationAggregationQueryMetricsItemsEnum = append(conversationAggregationQueryMetricsItemsEnum, v)
	}
}

func (m *ConversationAggregationQuery) validateMetricsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationAggregationQueryMetricsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationAggregationQuery) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {

		// value enum
		if err := m.validateMetricsItemsEnum("metrics"+"."+strconv.Itoa(i), "body", m.Metrics[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConversationAggregationQuery) validateViews(formats strfmt.Registry) error {

	if swag.IsZero(m.Views) { // not required
		return nil
	}

	for i := 0; i < len(m.Views); i++ {
		if swag.IsZero(m.Views[i]) { // not required
			continue
		}

		if m.Views[i] != nil {
			if err := m.Views[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("views" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationAggregationQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationAggregationQuery) UnmarshalBinary(b []byte) error {
	var res ConversationAggregationQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
