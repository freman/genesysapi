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

// AnalyticsParticipantWithoutAttributes analytics participant without attributes
//
// swagger:model AnalyticsParticipantWithoutAttributes
type AnalyticsParticipantWithoutAttributes struct {

	// External Contact Identifier
	ExternalContactID string `json:"externalContactId,omitempty"`

	// External Organization Identifier
	ExternalOrganizationID string `json:"externalOrganizationId,omitempty"`

	// Reason for which participant flagged conversation
	// Enum: [general]
	FlaggedReason string `json:"flaggedReason,omitempty"`

	// Unique identifier for the participant
	ParticipantID string `json:"participantId,omitempty"`

	// A human readable name identifying the participant
	ParticipantName string `json:"participantName,omitempty"`

	// The participant's purpose
	// Enum: [manual dialer inbound acd ivr voicemail outbound agent user station group customer external fax workflow campaign api]
	Purpose string `json:"purpose,omitempty"`

	// List of sessions associated to this participant
	Sessions []*AnalyticsSession `json:"sessions"`

	// The team id the user is a member of
	TeamID string `json:"teamId,omitempty"`

	// If a user, then this will be the unique identifier for the user
	UserID string `json:"userId,omitempty"`
}

// Validate validates this analytics participant without attributes
func (m *AnalyticsParticipantWithoutAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlaggedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurpose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var analyticsParticipantWithoutAttributesTypeFlaggedReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["general"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsParticipantWithoutAttributesTypeFlaggedReasonPropEnum = append(analyticsParticipantWithoutAttributesTypeFlaggedReasonPropEnum, v)
	}
}

const (

	// AnalyticsParticipantWithoutAttributesFlaggedReasonGeneral captures enum value "general"
	AnalyticsParticipantWithoutAttributesFlaggedReasonGeneral string = "general"
)

// prop value enum
func (m *AnalyticsParticipantWithoutAttributes) validateFlaggedReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsParticipantWithoutAttributesTypeFlaggedReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsParticipantWithoutAttributes) validateFlaggedReason(formats strfmt.Registry) error {

	if swag.IsZero(m.FlaggedReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateFlaggedReasonEnum("flaggedReason", "body", m.FlaggedReason); err != nil {
		return err
	}

	return nil
}

var analyticsParticipantWithoutAttributesTypePurposePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["manual","dialer","inbound","acd","ivr","voicemail","outbound","agent","user","station","group","customer","external","fax","workflow","campaign","api"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		analyticsParticipantWithoutAttributesTypePurposePropEnum = append(analyticsParticipantWithoutAttributesTypePurposePropEnum, v)
	}
}

const (

	// AnalyticsParticipantWithoutAttributesPurposeManual captures enum value "manual"
	AnalyticsParticipantWithoutAttributesPurposeManual string = "manual"

	// AnalyticsParticipantWithoutAttributesPurposeDialer captures enum value "dialer"
	AnalyticsParticipantWithoutAttributesPurposeDialer string = "dialer"

	// AnalyticsParticipantWithoutAttributesPurposeInbound captures enum value "inbound"
	AnalyticsParticipantWithoutAttributesPurposeInbound string = "inbound"

	// AnalyticsParticipantWithoutAttributesPurposeAcd captures enum value "acd"
	AnalyticsParticipantWithoutAttributesPurposeAcd string = "acd"

	// AnalyticsParticipantWithoutAttributesPurposeIvr captures enum value "ivr"
	AnalyticsParticipantWithoutAttributesPurposeIvr string = "ivr"

	// AnalyticsParticipantWithoutAttributesPurposeVoicemail captures enum value "voicemail"
	AnalyticsParticipantWithoutAttributesPurposeVoicemail string = "voicemail"

	// AnalyticsParticipantWithoutAttributesPurposeOutbound captures enum value "outbound"
	AnalyticsParticipantWithoutAttributesPurposeOutbound string = "outbound"

	// AnalyticsParticipantWithoutAttributesPurposeAgent captures enum value "agent"
	AnalyticsParticipantWithoutAttributesPurposeAgent string = "agent"

	// AnalyticsParticipantWithoutAttributesPurposeUser captures enum value "user"
	AnalyticsParticipantWithoutAttributesPurposeUser string = "user"

	// AnalyticsParticipantWithoutAttributesPurposeStation captures enum value "station"
	AnalyticsParticipantWithoutAttributesPurposeStation string = "station"

	// AnalyticsParticipantWithoutAttributesPurposeGroup captures enum value "group"
	AnalyticsParticipantWithoutAttributesPurposeGroup string = "group"

	// AnalyticsParticipantWithoutAttributesPurposeCustomer captures enum value "customer"
	AnalyticsParticipantWithoutAttributesPurposeCustomer string = "customer"

	// AnalyticsParticipantWithoutAttributesPurposeExternal captures enum value "external"
	AnalyticsParticipantWithoutAttributesPurposeExternal string = "external"

	// AnalyticsParticipantWithoutAttributesPurposeFax captures enum value "fax"
	AnalyticsParticipantWithoutAttributesPurposeFax string = "fax"

	// AnalyticsParticipantWithoutAttributesPurposeWorkflow captures enum value "workflow"
	AnalyticsParticipantWithoutAttributesPurposeWorkflow string = "workflow"

	// AnalyticsParticipantWithoutAttributesPurposeCampaign captures enum value "campaign"
	AnalyticsParticipantWithoutAttributesPurposeCampaign string = "campaign"

	// AnalyticsParticipantWithoutAttributesPurposeAPI captures enum value "api"
	AnalyticsParticipantWithoutAttributesPurposeAPI string = "api"
)

// prop value enum
func (m *AnalyticsParticipantWithoutAttributes) validatePurposeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, analyticsParticipantWithoutAttributesTypePurposePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AnalyticsParticipantWithoutAttributes) validatePurpose(formats strfmt.Registry) error {

	if swag.IsZero(m.Purpose) { // not required
		return nil
	}

	// value enum
	if err := m.validatePurposeEnum("purpose", "body", m.Purpose); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsParticipantWithoutAttributes) validateSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.Sessions) { // not required
		return nil
	}

	for i := 0; i < len(m.Sessions); i++ {
		if swag.IsZero(m.Sessions[i]) { // not required
			continue
		}

		if m.Sessions[i] != nil {
			if err := m.Sessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsParticipantWithoutAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsParticipantWithoutAttributes) UnmarshalBinary(b []byte) error {
	var res AnalyticsParticipantWithoutAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
