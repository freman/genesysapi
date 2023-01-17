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

// BotAggregateQueryPredicate bot aggregate query predicate
//
// swagger:model BotAggregateQueryPredicate
type BotAggregateQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [askActionId askActionResult botFinalIntent botId botIntent botProduct botProvider botRecognitionFailureReason botResult botSessionId botSlot botVersion conversationId externalContactId knowledgeBaseId lastActionId lastInputActionId mediaType messageType selfServed]
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

// Validate validates this bot aggregate query predicate
func (m *BotAggregateQueryPredicate) Validate(formats strfmt.Registry) error {
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

var botAggregateQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["askActionId","askActionResult","botFinalIntent","botId","botIntent","botProduct","botProvider","botRecognitionFailureReason","botResult","botSessionId","botSlot","botVersion","conversationId","externalContactId","knowledgeBaseId","lastActionId","lastInputActionId","mediaType","messageType","selfServed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		botAggregateQueryPredicateTypeDimensionPropEnum = append(botAggregateQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// BotAggregateQueryPredicateDimensionAskActionID captures enum value "askActionId"
	BotAggregateQueryPredicateDimensionAskActionID string = "askActionId"

	// BotAggregateQueryPredicateDimensionAskActionResult captures enum value "askActionResult"
	BotAggregateQueryPredicateDimensionAskActionResult string = "askActionResult"

	// BotAggregateQueryPredicateDimensionBotFinalIntent captures enum value "botFinalIntent"
	BotAggregateQueryPredicateDimensionBotFinalIntent string = "botFinalIntent"

	// BotAggregateQueryPredicateDimensionBotID captures enum value "botId"
	BotAggregateQueryPredicateDimensionBotID string = "botId"

	// BotAggregateQueryPredicateDimensionBotIntent captures enum value "botIntent"
	BotAggregateQueryPredicateDimensionBotIntent string = "botIntent"

	// BotAggregateQueryPredicateDimensionBotProduct captures enum value "botProduct"
	BotAggregateQueryPredicateDimensionBotProduct string = "botProduct"

	// BotAggregateQueryPredicateDimensionBotProvider captures enum value "botProvider"
	BotAggregateQueryPredicateDimensionBotProvider string = "botProvider"

	// BotAggregateQueryPredicateDimensionBotRecognitionFailureReason captures enum value "botRecognitionFailureReason"
	BotAggregateQueryPredicateDimensionBotRecognitionFailureReason string = "botRecognitionFailureReason"

	// BotAggregateQueryPredicateDimensionBotResult captures enum value "botResult"
	BotAggregateQueryPredicateDimensionBotResult string = "botResult"

	// BotAggregateQueryPredicateDimensionBotSessionID captures enum value "botSessionId"
	BotAggregateQueryPredicateDimensionBotSessionID string = "botSessionId"

	// BotAggregateQueryPredicateDimensionBotSlot captures enum value "botSlot"
	BotAggregateQueryPredicateDimensionBotSlot string = "botSlot"

	// BotAggregateQueryPredicateDimensionBotVersion captures enum value "botVersion"
	BotAggregateQueryPredicateDimensionBotVersion string = "botVersion"

	// BotAggregateQueryPredicateDimensionConversationID captures enum value "conversationId"
	BotAggregateQueryPredicateDimensionConversationID string = "conversationId"

	// BotAggregateQueryPredicateDimensionExternalContactID captures enum value "externalContactId"
	BotAggregateQueryPredicateDimensionExternalContactID string = "externalContactId"

	// BotAggregateQueryPredicateDimensionKnowledgeBaseID captures enum value "knowledgeBaseId"
	BotAggregateQueryPredicateDimensionKnowledgeBaseID string = "knowledgeBaseId"

	// BotAggregateQueryPredicateDimensionLastActionID captures enum value "lastActionId"
	BotAggregateQueryPredicateDimensionLastActionID string = "lastActionId"

	// BotAggregateQueryPredicateDimensionLastInputActionID captures enum value "lastInputActionId"
	BotAggregateQueryPredicateDimensionLastInputActionID string = "lastInputActionId"

	// BotAggregateQueryPredicateDimensionMediaType captures enum value "mediaType"
	BotAggregateQueryPredicateDimensionMediaType string = "mediaType"

	// BotAggregateQueryPredicateDimensionMessageType captures enum value "messageType"
	BotAggregateQueryPredicateDimensionMessageType string = "messageType"

	// BotAggregateQueryPredicateDimensionSelfServed captures enum value "selfServed"
	BotAggregateQueryPredicateDimensionSelfServed string = "selfServed"
)

// prop value enum
func (m *BotAggregateQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, botAggregateQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BotAggregateQueryPredicate) validateDimension(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var botAggregateQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		botAggregateQueryPredicateTypeOperatorPropEnum = append(botAggregateQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// BotAggregateQueryPredicateOperatorMatches captures enum value "matches"
	BotAggregateQueryPredicateOperatorMatches string = "matches"

	// BotAggregateQueryPredicateOperatorExists captures enum value "exists"
	BotAggregateQueryPredicateOperatorExists string = "exists"

	// BotAggregateQueryPredicateOperatorNotExists captures enum value "notExists"
	BotAggregateQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *BotAggregateQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, botAggregateQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BotAggregateQueryPredicate) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *BotAggregateQueryPredicate) validateRange(formats strfmt.Registry) error {
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

var botAggregateQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		botAggregateQueryPredicateTypeTypePropEnum = append(botAggregateQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// BotAggregateQueryPredicateTypeDimension captures enum value "dimension"
	BotAggregateQueryPredicateTypeDimension string = "dimension"

	// BotAggregateQueryPredicateTypeProperty captures enum value "property"
	BotAggregateQueryPredicateTypeProperty string = "property"

	// BotAggregateQueryPredicateTypeMetric captures enum value "metric"
	BotAggregateQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *BotAggregateQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, botAggregateQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BotAggregateQueryPredicate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bot aggregate query predicate based on the context it is used
func (m *BotAggregateQueryPredicate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BotAggregateQueryPredicate) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BotAggregateQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BotAggregateQueryPredicate) UnmarshalBinary(b []byte) error {
	var res BotAggregateQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
