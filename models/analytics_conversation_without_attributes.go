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

// AnalyticsConversationWithoutAttributes analytics conversation without attributes
//
// swagger:model AnalyticsConversationWithoutAttributes
type AnalyticsConversationWithoutAttributes struct {

	// Date/time the conversation ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ConversationEnd strfmt.DateTime `json:"conversationEnd,omitempty"`

	// Unique identifier for the conversation
	ConversationID string `json:"conversationId,omitempty"`

	// Date/time the conversation started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ConversationStart strfmt.DateTime `json:"conversationStart,omitempty"`

	// Identifiers of divisions associated with this conversation
	DivisionIds []string `json:"divisionIds"`

	// Evaluations tied to this conversation
	Evaluations []*AnalyticsEvaluation `json:"evaluations"`

	// The lowest estimated average MOS among all the audio streams belonging to this conversation
	MediaStatsMinConversationMos float64 `json:"mediaStatsMinConversationMos,omitempty"`

	// The lowest R-factor value among all of the audio streams belonging to this conversation
	MediaStatsMinConversationRFactor float64 `json:"mediaStatsMinConversationRFactor,omitempty"`

	// The original direction of the conversation
	// Enum: [inbound outbound]
	OriginatingDirection string `json:"originatingDirection,omitempty"`

	// Participants in the conversation
	Participants []*AnalyticsParticipantWithoutAttributes `json:"participants"`

	// Surveys tied to this conversation
	Surveys []*AnalyticsSurvey `json:"surveys"`
}

// Validate validates this analytics conversation without attributes
func (m *AnalyticsConversationWithoutAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversationEnd(formats); err != nil {
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

	if err := m.validateSurveys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsConversationWithoutAttributes) validateConversationEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("conversationEnd", "body", "date-time", m.ConversationEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsConversationWithoutAttributes) validateConversationStart(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationStart) { // not required
		return nil
	}

	if err := validate.FormatOf("conversationStart", "body", "date-time", m.ConversationStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsConversationWithoutAttributes) validateEvaluations(formats strfmt.Registry) error {

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

var analyticsConversationWithoutAttributesTypeOriginatingDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsConversationWithoutAttributesTypeOriginatingDirectionPropEnum = append(analyticsConversationWithoutAttributesTypeOriginatingDirectionPropEnum, v)
	}
}

const (

	// AnalyticsConversationWithoutAttributesOriginatingDirectionInbound captures enum value "inbound"
	AnalyticsConversationWithoutAttributesOriginatingDirectionInbound string = "inbound"

	// AnalyticsConversationWithoutAttributesOriginatingDirectionOutbound captures enum value "outbound"
	AnalyticsConversationWithoutAttributesOriginatingDirectionOutbound string = "outbound"
)

// prop value enum
func (m *AnalyticsConversationWithoutAttributes) validateOriginatingDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsConversationWithoutAttributesTypeOriginatingDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsConversationWithoutAttributes) validateOriginatingDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginatingDirection) { // not required
		return nil
	}

	// value enum
	if err := m.validateOriginatingDirectionEnum("originatingDirection", "body", m.OriginatingDirection); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsConversationWithoutAttributes) validateParticipants(formats strfmt.Registry) error {

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

func (m *AnalyticsConversationWithoutAttributes) validateSurveys(formats strfmt.Registry) error {

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
func (m *AnalyticsConversationWithoutAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsConversationWithoutAttributes) UnmarshalBinary(b []byte) error {
	var res AnalyticsConversationWithoutAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
