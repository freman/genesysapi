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

// TranscriptAggregateQueryPredicate transcript aggregate query predicate
//
// swagger:model TranscriptAggregateQueryPredicate
type TranscriptAggregateQueryPredicate struct {

	// Left hand side for dimension predicates
	// Enum: [addressFrom addressTo ani channel conversationId direction divisionId dnis flowId flowVersion mediaType messageType queueId resultsBy teamId topicId userId]
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

// Validate validates this transcript aggregate query predicate
func (m *TranscriptAggregateQueryPredicate) Validate(formats strfmt.Registry) error {
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

var transcriptAggregateQueryPredicateTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["addressFrom","addressTo","ani","channel","conversationId","direction","divisionId","dnis","flowId","flowVersion","mediaType","messageType","queueId","resultsBy","teamId","topicId","userId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptAggregateQueryPredicateTypeDimensionPropEnum = append(transcriptAggregateQueryPredicateTypeDimensionPropEnum, v)
	}
}

const (

	// TranscriptAggregateQueryPredicateDimensionAddressFrom captures enum value "addressFrom"
	TranscriptAggregateQueryPredicateDimensionAddressFrom string = "addressFrom"

	// TranscriptAggregateQueryPredicateDimensionAddressTo captures enum value "addressTo"
	TranscriptAggregateQueryPredicateDimensionAddressTo string = "addressTo"

	// TranscriptAggregateQueryPredicateDimensionAni captures enum value "ani"
	TranscriptAggregateQueryPredicateDimensionAni string = "ani"

	// TranscriptAggregateQueryPredicateDimensionChannel captures enum value "channel"
	TranscriptAggregateQueryPredicateDimensionChannel string = "channel"

	// TranscriptAggregateQueryPredicateDimensionConversationID captures enum value "conversationId"
	TranscriptAggregateQueryPredicateDimensionConversationID string = "conversationId"

	// TranscriptAggregateQueryPredicateDimensionDirection captures enum value "direction"
	TranscriptAggregateQueryPredicateDimensionDirection string = "direction"

	// TranscriptAggregateQueryPredicateDimensionDivisionID captures enum value "divisionId"
	TranscriptAggregateQueryPredicateDimensionDivisionID string = "divisionId"

	// TranscriptAggregateQueryPredicateDimensionDnis captures enum value "dnis"
	TranscriptAggregateQueryPredicateDimensionDnis string = "dnis"

	// TranscriptAggregateQueryPredicateDimensionFlowID captures enum value "flowId"
	TranscriptAggregateQueryPredicateDimensionFlowID string = "flowId"

	// TranscriptAggregateQueryPredicateDimensionFlowVersion captures enum value "flowVersion"
	TranscriptAggregateQueryPredicateDimensionFlowVersion string = "flowVersion"

	// TranscriptAggregateQueryPredicateDimensionMediaType captures enum value "mediaType"
	TranscriptAggregateQueryPredicateDimensionMediaType string = "mediaType"

	// TranscriptAggregateQueryPredicateDimensionMessageType captures enum value "messageType"
	TranscriptAggregateQueryPredicateDimensionMessageType string = "messageType"

	// TranscriptAggregateQueryPredicateDimensionQueueID captures enum value "queueId"
	TranscriptAggregateQueryPredicateDimensionQueueID string = "queueId"

	// TranscriptAggregateQueryPredicateDimensionResultsBy captures enum value "resultsBy"
	TranscriptAggregateQueryPredicateDimensionResultsBy string = "resultsBy"

	// TranscriptAggregateQueryPredicateDimensionTeamID captures enum value "teamId"
	TranscriptAggregateQueryPredicateDimensionTeamID string = "teamId"

	// TranscriptAggregateQueryPredicateDimensionTopicID captures enum value "topicId"
	TranscriptAggregateQueryPredicateDimensionTopicID string = "topicId"

	// TranscriptAggregateQueryPredicateDimensionUserID captures enum value "userId"
	TranscriptAggregateQueryPredicateDimensionUserID string = "userId"
)

// prop value enum
func (m *TranscriptAggregateQueryPredicate) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptAggregateQueryPredicateTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptAggregateQueryPredicate) validateDimension(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimension) { // not required
		return nil
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

var transcriptAggregateQueryPredicateTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["matches","exists","notExists"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptAggregateQueryPredicateTypeOperatorPropEnum = append(transcriptAggregateQueryPredicateTypeOperatorPropEnum, v)
	}
}

const (

	// TranscriptAggregateQueryPredicateOperatorMatches captures enum value "matches"
	TranscriptAggregateQueryPredicateOperatorMatches string = "matches"

	// TranscriptAggregateQueryPredicateOperatorExists captures enum value "exists"
	TranscriptAggregateQueryPredicateOperatorExists string = "exists"

	// TranscriptAggregateQueryPredicateOperatorNotExists captures enum value "notExists"
	TranscriptAggregateQueryPredicateOperatorNotExists string = "notExists"
)

// prop value enum
func (m *TranscriptAggregateQueryPredicate) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptAggregateQueryPredicateTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptAggregateQueryPredicate) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *TranscriptAggregateQueryPredicate) validateRange(formats strfmt.Registry) error {

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

var transcriptAggregateQueryPredicateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dimension","property","metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptAggregateQueryPredicateTypeTypePropEnum = append(transcriptAggregateQueryPredicateTypeTypePropEnum, v)
	}
}

const (

	// TranscriptAggregateQueryPredicateTypeDimension captures enum value "dimension"
	TranscriptAggregateQueryPredicateTypeDimension string = "dimension"

	// TranscriptAggregateQueryPredicateTypeProperty captures enum value "property"
	TranscriptAggregateQueryPredicateTypeProperty string = "property"

	// TranscriptAggregateQueryPredicateTypeMetric captures enum value "metric"
	TranscriptAggregateQueryPredicateTypeMetric string = "metric"
)

// prop value enum
func (m *TranscriptAggregateQueryPredicate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptAggregateQueryPredicateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptAggregateQueryPredicate) validateType(formats strfmt.Registry) error {

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
func (m *TranscriptAggregateQueryPredicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranscriptAggregateQueryPredicate) UnmarshalBinary(b []byte) error {
	var res TranscriptAggregateQueryPredicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
