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

// JourneyAggregateQueryPredicate journey aggregate query predicate
//
// swagger:model JourneyAggregateQueryPredicate
type JourneyAggregateQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [containsAllCondition containsAnyCondition endsWithCondition equalCondition journeyActionId journeyActionMapId journeyActionMapVersion journeyActionMediaType journeyActionTargetId journeyActionTemplateId journeyBlockingActionMapId journeyBlockingEmergencyScheduleGroupId journeyBlockingReason journeyBlockingScheduleGroupId journeyDeviceCategory journeyDeviceType journeyFrequencyCapReason journeyIpGeolocationCountry journeyOutcomeId journeySegmentId journeySegmentScope journeySessionId journeySessionSegmentId journeySessionType notContainsAllCondition notContainsAnyCondition notEqualCondition startsWithCondition touchpointActionMapId touchpointAgentId touchpointAttributionScope touchpointChannelMessageType touchpointChannelPlatform touchpointChannelType touchpointConversationId touchpointInteractionType touchpointQueueId touchpointWrapupCode]
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

// Validate validates this journey aggregate query predicate
func (m *JourneyAggregateQueryPredicate) Validate(formats strfmt.Registry) error {
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

var journeyAggregateQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["containsAllCondition","containsAnyCondition","endsWithCondition","equalCondition","journeyActionId","journeyActionMapId","journeyActionMapVersion","journeyActionMediaType","journeyActionTargetId","journeyActionTemplateId","journeyBlockingActionMapId","journeyBlockingEmergencyScheduleGroupId","journeyBlockingReason","journeyBlockingScheduleGroupId","journeyDeviceCategory","journeyDeviceType","journeyFrequencyCapReason","journeyIpGeolocationCountry","journeyOutcomeId","journeySegmentId","journeySegmentScope","journeySessionId","journeySessionSegmentId","journeySessionType","notContainsAllCondition","notContainsAnyCondition","notEqualCondition","startsWithCondition","touchpointActionMapId","touchpointAgentId","touchpointAttributionScope","touchpointChannelMessageType","touchpointChannelPlatform","touchpointChannelType","touchpointConversationId","touchpointInteractionType","touchpointQueueId","touchpointWrapupCode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journeyAggregateQueryPredicateTypeDimensionPropEnum = append(journeyAggregateQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// JourneyAggregateQueryPredicateDimensionContainsAllCondition captures enum value "containsAllCondition"
	JourneyAggregateQueryPredicateDimensionContainsAllCondition string = "containsAllCondition"

	// JourneyAggregateQueryPredicateDimensionContainsAnyCondition captures enum value "containsAnyCondition"
	JourneyAggregateQueryPredicateDimensionContainsAnyCondition string = "containsAnyCondition"

	// JourneyAggregateQueryPredicateDimensionEndsWithCondition captures enum value "endsWithCondition"
	JourneyAggregateQueryPredicateDimensionEndsWithCondition string = "endsWithCondition"

	// JourneyAggregateQueryPredicateDimensionEqualCondition captures enum value "equalCondition"
	JourneyAggregateQueryPredicateDimensionEqualCondition string = "equalCondition"

	// JourneyAggregateQueryPredicateDimensionJourneyActionID captures enum value "journeyActionId"
	JourneyAggregateQueryPredicateDimensionJourneyActionID string = "journeyActionId"

	// JourneyAggregateQueryPredicateDimensionJourneyActionMapID captures enum value "journeyActionMapId"
	JourneyAggregateQueryPredicateDimensionJourneyActionMapID string = "journeyActionMapId"

	// JourneyAggregateQueryPredicateDimensionJourneyActionMapVersion captures enum value "journeyActionMapVersion"
	JourneyAggregateQueryPredicateDimensionJourneyActionMapVersion string = "journeyActionMapVersion"

	// JourneyAggregateQueryPredicateDimensionJourneyActionMediaType captures enum value "journeyActionMediaType"
	JourneyAggregateQueryPredicateDimensionJourneyActionMediaType string = "journeyActionMediaType"

	// JourneyAggregateQueryPredicateDimensionJourneyActionTargetID captures enum value "journeyActionTargetId"
	JourneyAggregateQueryPredicateDimensionJourneyActionTargetID string = "journeyActionTargetId"

	// JourneyAggregateQueryPredicateDimensionJourneyActionTemplateID captures enum value "journeyActionTemplateId"
	JourneyAggregateQueryPredicateDimensionJourneyActionTemplateID string = "journeyActionTemplateId"

	// JourneyAggregateQueryPredicateDimensionJourneyBlockingActionMapID captures enum value "journeyBlockingActionMapId"
	JourneyAggregateQueryPredicateDimensionJourneyBlockingActionMapID string = "journeyBlockingActionMapId"

	// JourneyAggregateQueryPredicateDimensionJourneyBlockingEmergencyScheduleGroupID captures enum value "journeyBlockingEmergencyScheduleGroupId"
	JourneyAggregateQueryPredicateDimensionJourneyBlockingEmergencyScheduleGroupID string = "journeyBlockingEmergencyScheduleGroupId"

	// JourneyAggregateQueryPredicateDimensionJourneyBlockingReason captures enum value "journeyBlockingReason"
	JourneyAggregateQueryPredicateDimensionJourneyBlockingReason string = "journeyBlockingReason"

	// JourneyAggregateQueryPredicateDimensionJourneyBlockingScheduleGroupID captures enum value "journeyBlockingScheduleGroupId"
	JourneyAggregateQueryPredicateDimensionJourneyBlockingScheduleGroupID string = "journeyBlockingScheduleGroupId"

	// JourneyAggregateQueryPredicateDimensionJourneyDeviceCategory captures enum value "journeyDeviceCategory"
	JourneyAggregateQueryPredicateDimensionJourneyDeviceCategory string = "journeyDeviceCategory"

	// JourneyAggregateQueryPredicateDimensionJourneyDeviceType captures enum value "journeyDeviceType"
	JourneyAggregateQueryPredicateDimensionJourneyDeviceType string = "journeyDeviceType"

	// JourneyAggregateQueryPredicateDimensionJourneyFrequencyCapReason captures enum value "journeyFrequencyCapReason"
	JourneyAggregateQueryPredicateDimensionJourneyFrequencyCapReason string = "journeyFrequencyCapReason"

	// JourneyAggregateQueryPredicateDimensionJourneyIPGeolocationCountry captures enum value "journeyIpGeolocationCountry"
	JourneyAggregateQueryPredicateDimensionJourneyIPGeolocationCountry string = "journeyIpGeolocationCountry"

	// JourneyAggregateQueryPredicateDimensionJourneyOutcomeID captures enum value "journeyOutcomeId"
	JourneyAggregateQueryPredicateDimensionJourneyOutcomeID string = "journeyOutcomeId"

	// JourneyAggregateQueryPredicateDimensionJourneySegmentID captures enum value "journeySegmentId"
	JourneyAggregateQueryPredicateDimensionJourneySegmentID string = "journeySegmentId"

	// JourneyAggregateQueryPredicateDimensionJourneySegmentScope captures enum value "journeySegmentScope"
	JourneyAggregateQueryPredicateDimensionJourneySegmentScope string = "journeySegmentScope"

	// JourneyAggregateQueryPredicateDimensionJourneySessionID captures enum value "journeySessionId"
	JourneyAggregateQueryPredicateDimensionJourneySessionID string = "journeySessionId"

	// JourneyAggregateQueryPredicateDimensionJourneySessionSegmentID captures enum value "journeySessionSegmentId"
	JourneyAggregateQueryPredicateDimensionJourneySessionSegmentID string = "journeySessionSegmentId"

	// JourneyAggregateQueryPredicateDimensionJourneySessionType captures enum value "journeySessionType"
	JourneyAggregateQueryPredicateDimensionJourneySessionType string = "journeySessionType"

	// JourneyAggregateQueryPredicateDimensionNotContainsAllCondition captures enum value "notContainsAllCondition"
	JourneyAggregateQueryPredicateDimensionNotContainsAllCondition string = "notContainsAllCondition"

	// JourneyAggregateQueryPredicateDimensionNotContainsAnyCondition captures enum value "notContainsAnyCondition"
	JourneyAggregateQueryPredicateDimensionNotContainsAnyCondition string = "notContainsAnyCondition"

	// JourneyAggregateQueryPredicateDimensionNotEqualCondition captures enum value "notEqualCondition"
	JourneyAggregateQueryPredicateDimensionNotEqualCondition string = "notEqualCondition"

	// JourneyAggregateQueryPredicateDimensionStartsWithCondition captures enum value "startsWithCondition"
	JourneyAggregateQueryPredicateDimensionStartsWithCondition string = "startsWithCondition"

	// JourneyAggregateQueryPredicateDimensionTouchpointActionMapID captures enum value "touchpointActionMapId"
	JourneyAggregateQueryPredicateDimensionTouchpointActionMapID string = "touchpointActionMapId"

	// JourneyAggregateQueryPredicateDimensionTouchpointAgentID captures enum value "touchpointAgentId"
	JourneyAggregateQueryPredicateDimensionTouchpointAgentID string = "touchpointAgentId"

	// JourneyAggregateQueryPredicateDimensionTouchpointAttributionScope captures enum value "touchpointAttributionScope"
	JourneyAggregateQueryPredicateDimensionTouchpointAttributionScope string = "touchpointAttributionScope"

	// JourneyAggregateQueryPredicateDimensionTouchpointChannelMessageType captures enum value "touchpointChannelMessageType"
	JourneyAggregateQueryPredicateDimensionTouchpointChannelMessageType string = "touchpointChannelMessageType"

	// JourneyAggregateQueryPredicateDimensionTouchpointChannelPlatform captures enum value "touchpointChannelPlatform"
	JourneyAggregateQueryPredicateDimensionTouchpointChannelPlatform string = "touchpointChannelPlatform"

	// JourneyAggregateQueryPredicateDimensionTouchpointChannelType captures enum value "touchpointChannelType"
	JourneyAggregateQueryPredicateDimensionTouchpointChannelType string = "touchpointChannelType"

	// JourneyAggregateQueryPredicateDimensionTouchpointConversationID captures enum value "touchpointConversationId"
	JourneyAggregateQueryPredicateDimensionTouchpointConversationID string = "touchpointConversationId"

	// JourneyAggregateQueryPredicateDimensionTouchpointInteractionType captures enum value "touchpointInteractionType"
	JourneyAggregateQueryPredicateDimensionTouchpointInteractionType string = "touchpointInteractionType"

	// JourneyAggregateQueryPredicateDimensionTouchpointQueueID captures enum value "touchpointQueueId"
	JourneyAggregateQueryPredicateDimensionTouchpointQueueID string = "touchpointQueueId"

	// JourneyAggregateQueryPredicateDimensionTouchpointWrapupCode captures enum value "touchpointWrapupCode"
	JourneyAggregateQueryPredicateDimensionTouchpointWrapupCode string = "touchpointWrapupCode"
)

// prop value enum
func (m *JourneyAggregateQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journeyAggregateQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JourneyAggregateQueryPredicate) validateDimension(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var journeyAggregateQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journeyAggregateQueryPredicateTypeOperatorPropEnum = append(journeyAggregateQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// JourneyAggregateQueryPredicateOperatorMatches captures enum value "matches"
	JourneyAggregateQueryPredicateOperatorMatches string = "matches"

	// JourneyAggregateQueryPredicateOperatorExists captures enum value "exists"
	JourneyAggregateQueryPredicateOperatorExists string = "exists"

	// JourneyAggregateQueryPredicateOperatorNotExists captures enum value "notExists"
	JourneyAggregateQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *JourneyAggregateQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journeyAggregateQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JourneyAggregateQueryPredicate) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *JourneyAggregateQueryPredicate) validateRange(formats strfmt.Registry) error {
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

var journeyAggregateQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journeyAggregateQueryPredicateTypeTypePropEnum = append(journeyAggregateQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// JourneyAggregateQueryPredicateTypeDimension captures enum value "dimension"
	JourneyAggregateQueryPredicateTypeDimension string = "dimension"

	// JourneyAggregateQueryPredicateTypeProperty captures enum value "property"
	JourneyAggregateQueryPredicateTypeProperty string = "property"

	// JourneyAggregateQueryPredicateTypeMetric captures enum value "metric"
	JourneyAggregateQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *JourneyAggregateQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journeyAggregateQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JourneyAggregateQueryPredicate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this journey aggregate query predicate based on the context it is used
func (m *JourneyAggregateQueryPredicate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JourneyAggregateQueryPredicate) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

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
func (m *JourneyAggregateQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JourneyAggregateQueryPredicate) UnmarshalBinary(b []byte) error {
	var res JourneyAggregateQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
