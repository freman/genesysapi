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

// Organization organization
//
// swagger:model Organization
type Organization struct {

	// The default country code for this organization. Example: 'US'
	DefaultCountryCode string `json:"defaultCountryCode,omitempty"`

	// The default language for this organization. Example: 'en'
	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	// default site Id
	DefaultSiteID string `json:"defaultSiteId,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// The state of features available for the organization.
	// Read Only: true
	Features map[string]bool `json:"features,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Organizations Originating Platform.
	// Read Only: true
	// Enum: [PureCloud PureEngage PureEngageCloud PureConnect PureConnectCloud Unknown]
	ProductPlatform string `json:"productPlatform,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The current state. Examples are active, inactive, deleted.
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// Email address where support tickets are sent to.
	SupportURI string `json:"supportURI,omitempty"`

	// The short name for the organization. This field is globally unique and cannot be changed.
	// Read Only: true
	ThirdPartyOrgName string `json:"thirdPartyOrgName,omitempty"`

	// third party URI
	// Format: uri
	ThirdPartyURI strfmt.URI `json:"thirdPartyURI,omitempty"`

	// The current version of the organization.
	// Required: true
	Version *int32 `json:"version"`

	// voicemail enabled
	VoicemailEnabled bool `json:"voicemailEnabled"`
}

// Validate validates this organization
func (m *Organization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProductPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThirdPartyURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var organizationTypeProductPlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PureCloud","PureEngage","PureEngageCloud","PureConnect","PureConnectCloud","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		organizationTypeProductPlatformPropEnum = append(organizationTypeProductPlatformPropEnum, v)
	}
}

const (

	// OrganizationProductPlatformPureCloud captures enum value "PureCloud"
	OrganizationProductPlatformPureCloud string = "PureCloud"

	// OrganizationProductPlatformPureEngage captures enum value "PureEngage"
	OrganizationProductPlatformPureEngage string = "PureEngage"

	// OrganizationProductPlatformPureEngageCloud captures enum value "PureEngageCloud"
	OrganizationProductPlatformPureEngageCloud string = "PureEngageCloud"

	// OrganizationProductPlatformPureConnect captures enum value "PureConnect"
	OrganizationProductPlatformPureConnect string = "PureConnect"

	// OrganizationProductPlatformPureConnectCloud captures enum value "PureConnectCloud"
	OrganizationProductPlatformPureConnectCloud string = "PureConnectCloud"

	// OrganizationProductPlatformUnknown captures enum value "Unknown"
	OrganizationProductPlatformUnknown string = "Unknown"
)

// prop value enum
func (m *Organization) validateProductPlatformEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, organizationTypeProductPlatformPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Organization) validateProductPlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductPlatform) { // not required
		return nil
	}

	// value enum
	if err := m.validateProductPlatformEnum("productPlatform", "body", m.ProductPlatform); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var organizationTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		organizationTypeStatePropEnum = append(organizationTypeStatePropEnum, v)
	}
}

const (

	// OrganizationStateActive captures enum value "active"
	OrganizationStateActive string = "active"

	// OrganizationStateInactive captures enum value "inactive"
	OrganizationStateInactive string = "inactive"

	// OrganizationStateDeleted captures enum value "deleted"
	OrganizationStateDeleted string = "deleted"
)

// prop value enum
func (m *Organization) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, organizationTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Organization) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateThirdPartyURI(formats strfmt.Registry) error {

	if swag.IsZero(m.ThirdPartyURI) { // not required
		return nil
	}

	if err := validate.FormatOf("thirdPartyURI", "body", "uri", m.ThirdPartyURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Organization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Organization) UnmarshalBinary(b []byte) error {
	var res Organization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
