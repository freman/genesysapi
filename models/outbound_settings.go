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

// OutboundSettings outbound settings
//
// swagger:model OutboundSettings
type OutboundSettings struct {

	// The number of seconds used to determine if a call is abandoned
	AbandonSeconds float64 `json:"abandonSeconds,omitempty"`

	// The settings for automatic time zone mapping. Note that changing these settings will change them for both voice and messaging campaigns.
	AutomaticTimeZoneMapping *AutomaticTimeZoneMappingSettings `json:"automaticTimeZoneMapping,omitempty"`

	// The denominator to be used in determining the compliance abandon rate
	// Enum: [ALL_CALLS CALLS_THAT_REACHED_QUEUE]
	ComplianceAbandonRateDenominator string `json:"complianceAbandonRateDenominator,omitempty"`

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The maximum number of calls that can be placed per agent on any campaign
	MaxCallsPerAgent int32 `json:"maxCallsPerAgent,omitempty"`

	// The maximum number of calls that can be configured to be placed per agent on any campaign
	// Read Only: true
	MaxConfigurableCallsPerAgent int32 `json:"maxConfigurableCallsPerAgent,omitempty"`

	// The maximum percentage of lines that should be used for Outbound, expressed as a decimal in the range [0.0, 1.0]
	MaxLineUtilization float64 `json:"maxLineUtilization,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this outbound settings
func (m *OutboundSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutomaticTimeZoneMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplianceAbandonRateDenominator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutboundSettings) validateAutomaticTimeZoneMapping(formats strfmt.Registry) error {
	if swag.IsZero(m.AutomaticTimeZoneMapping) { // not required
		return nil
	}

	if m.AutomaticTimeZoneMapping != nil {
		if err := m.AutomaticTimeZoneMapping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automaticTimeZoneMapping")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("automaticTimeZoneMapping")
			}
			return err
		}
	}

	return nil
}

var outboundSettingsTypeComplianceAbandonRateDenominatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL_CALLS","CALLS_THAT_REACHED_QUEUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		outboundSettingsTypeComplianceAbandonRateDenominatorPropEnum = append(outboundSettingsTypeComplianceAbandonRateDenominatorPropEnum, v)
	}
}

const (

	// OutboundSettingsComplianceAbandonRateDenominatorALLCALLS captures enum value "ALL_CALLS"
	OutboundSettingsComplianceAbandonRateDenominatorALLCALLS string = "ALL_CALLS"

	// OutboundSettingsComplianceAbandonRateDenominatorCALLSTHATREACHEDQUEUE captures enum value "CALLS_THAT_REACHED_QUEUE"
	OutboundSettingsComplianceAbandonRateDenominatorCALLSTHATREACHEDQUEUE string = "CALLS_THAT_REACHED_QUEUE"
)

// prop value enum
func (m *OutboundSettings) validateComplianceAbandonRateDenominatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, outboundSettingsTypeComplianceAbandonRateDenominatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OutboundSettings) validateComplianceAbandonRateDenominator(formats strfmt.Registry) error {
	if swag.IsZero(m.ComplianceAbandonRateDenominator) { // not required
		return nil
	}

	// value enum
	if err := m.validateComplianceAbandonRateDenominatorEnum("complianceAbandonRateDenominator", "body", m.ComplianceAbandonRateDenominator); err != nil {
		return err
	}

	return nil
}

func (m *OutboundSettings) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OutboundSettings) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OutboundSettings) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this outbound settings based on the context it is used
func (m *OutboundSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutomaticTimeZoneMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxConfigurableCallsPerAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutboundSettings) contextValidateAutomaticTimeZoneMapping(ctx context.Context, formats strfmt.Registry) error {

	if m.AutomaticTimeZoneMapping != nil {
		if err := m.AutomaticTimeZoneMapping.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automaticTimeZoneMapping")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("automaticTimeZoneMapping")
			}
			return err
		}
	}

	return nil
}

func (m *OutboundSettings) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *OutboundSettings) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *OutboundSettings) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *OutboundSettings) contextValidateMaxConfigurableCallsPerAgent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "maxConfigurableCallsPerAgent", "body", int32(m.MaxConfigurableCallsPerAgent)); err != nil {
		return err
	}

	return nil
}

func (m *OutboundSettings) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutboundSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutboundSettings) UnmarshalBinary(b []byte) error {
	var res OutboundSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
