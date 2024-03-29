// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObservationValue observation value
//
// swagger:model ObservationValue
type ObservationValue struct {

	// The address that initiated an action
	AddressFrom string `json:"addressFrom,omitempty"`

	// The address receiving an action
	AddressTo string `json:"addressTo,omitempty"`

	// Automatic Number Identification (caller's number)
	Ani string `json:"ani,omitempty"`

	// Unique identifier for the conversation
	ConversationID string `json:"conversationId,omitempty"`

	// Session media type that was converted from in case of a media type conversion
	ConvertedFrom string `json:"convertedFrom,omitempty"`

	// Session media type that was converted to in case of a media type conversion
	ConvertedTo string `json:"convertedTo,omitempty"`

	// The direction of the communication
	// Enum: [inbound outbound]
	Direction string `json:"direction,omitempty"`

	// Dialed number identification service (number dialed by the calling party)
	Dnis string `json:"dnis,omitempty"`

	// The time at which the observation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	ObservationDate *strfmt.DateTime `json:"observationDate"`

	// A human readable name identifying the participant
	ParticipantName string `json:"participantName,omitempty"`

	// Unique identifier for the language requested for an interaction
	RequestedLanguageID string `json:"requestedLanguageId,omitempty"`

	// Unique identifier for a skill requested for an interaction
	// Unique: true
	RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`

	// All routing types for requested/attempted routing methods
	// Unique: true
	RequestedRoutings []string `json:"requestedRoutings"`

	// Routing priority for the current interaction
	RoutingPriority int64 `json:"routingPriority,omitempty"`

	// scored agents
	ScoredAgents []*AnalyticsScoredAgent `json:"scoredAgents"`

	// The unique identifier of this session
	SessionID string `json:"sessionId,omitempty"`

	// The team id the user is a member of
	TeamID string `json:"teamId,omitempty"`

	// Complete routing method
	// Enum: [Bullseye Conditional Direct Last Manual Predictive Preferred Standard Vip]
	UsedRouting string `json:"usedRouting,omitempty"`

	// Unique identifier for the user
	UserID string `json:"userId,omitempty"`
}

// Validate validates this observation value
func (m *ObservationValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObservationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedRoutingSkillIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedRoutings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoredAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedRouting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var observationValueTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inbound","outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		observationValueTypeDirectionPropEnum = append(observationValueTypeDirectionPropEnum, v)
	}
}

const (

	// ObservationValueDirectionInbound captures enum value "inbound"
	ObservationValueDirectionInbound string = "inbound"

	// ObservationValueDirectionOutbound captures enum value "outbound"
	ObservationValueDirectionOutbound string = "outbound"
)

// prop value enum
func (m *ObservationValue) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, observationValueTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObservationValue) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ObservationValue) validateObservationDate(formats strfmt.Registry) error {

	if err := validate.Required("observationDate", "body", m.ObservationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("observationDate", "body", "date-time", m.ObservationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObservationValue) validateRequestedRoutingSkillIds(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedRoutingSkillIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("requestedRoutingSkillIds", "body", m.RequestedRoutingSkillIds); err != nil {
		return err
	}

	return nil
}

var observationValueRequestedRoutingsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bullseye","Conditional","Direct","Last","Manual","Predictive","Preferred","Standard","Vip"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		observationValueRequestedRoutingsItemsEnum = append(observationValueRequestedRoutingsItemsEnum, v)
	}
}

func (m *ObservationValue) validateRequestedRoutingsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, observationValueRequestedRoutingsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObservationValue) validateRequestedRoutings(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedRoutings) { // not required
		return nil
	}

	if err := validate.UniqueItems("requestedRoutings", "body", m.RequestedRoutings); err != nil {
		return err
	}

	for i := 0; i < len(m.RequestedRoutings); i++ {

		// value enum
		if err := m.validateRequestedRoutingsItemsEnum("requestedRoutings"+"."+strconv.Itoa(i), "body", m.RequestedRoutings[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObservationValue) validateScoredAgents(formats strfmt.Registry) error {
	if swag.IsZero(m.ScoredAgents) { // not required
		return nil
	}

	for i := 0; i < len(m.ScoredAgents); i++ {
		if swag.IsZero(m.ScoredAgents[i]) { // not required
			continue
		}

		if m.ScoredAgents[i] != nil {
			if err := m.ScoredAgents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var observationValueTypeUsedRoutingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bullseye","Conditional","Direct","Last","Manual","Predictive","Preferred","Standard","Vip"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		observationValueTypeUsedRoutingPropEnum = append(observationValueTypeUsedRoutingPropEnum, v)
	}
}

const (

	// ObservationValueUsedRoutingBullseye captures enum value "Bullseye"
	ObservationValueUsedRoutingBullseye string = "Bullseye"

	// ObservationValueUsedRoutingConditional captures enum value "Conditional"
	ObservationValueUsedRoutingConditional string = "Conditional"

	// ObservationValueUsedRoutingDirect captures enum value "Direct"
	ObservationValueUsedRoutingDirect string = "Direct"

	// ObservationValueUsedRoutingLast captures enum value "Last"
	ObservationValueUsedRoutingLast string = "Last"

	// ObservationValueUsedRoutingManual captures enum value "Manual"
	ObservationValueUsedRoutingManual string = "Manual"

	// ObservationValueUsedRoutingPredictive captures enum value "Predictive"
	ObservationValueUsedRoutingPredictive string = "Predictive"

	// ObservationValueUsedRoutingPreferred captures enum value "Preferred"
	ObservationValueUsedRoutingPreferred string = "Preferred"

	// ObservationValueUsedRoutingStandard captures enum value "Standard"
	ObservationValueUsedRoutingStandard string = "Standard"

	// ObservationValueUsedRoutingVip captures enum value "Vip"
	ObservationValueUsedRoutingVip string = "Vip"
)

// prop value enum
func (m *ObservationValue) validateUsedRoutingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, observationValueTypeUsedRoutingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObservationValue) validateUsedRouting(formats strfmt.Registry) error {
	if swag.IsZero(m.UsedRouting) { // not required
		return nil
	}

	// value enum
	if err := m.validateUsedRoutingEnum("usedRouting", "body", m.UsedRouting); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this observation value based on the context it is used
func (m *ObservationValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScoredAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObservationValue) contextValidateScoredAgents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ScoredAgents); i++ {

		if m.ScoredAgents[i] != nil {
			if err := m.ScoredAgents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scoredAgents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObservationValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObservationValue) UnmarshalBinary(b []byte) error {
	var res ObservationValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
