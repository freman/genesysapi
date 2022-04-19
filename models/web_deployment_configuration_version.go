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

// WebDeploymentConfigurationVersion Details about the configuration version of a Web Deployment
//
// swagger:model WebDeploymentConfigurationVersion
type WebDeploymentConfigurationVersion struct {

	// The settings for authenticated deployments
	AuthenticationSettings *AuthenticationSettings `json:"authenticationSettings,omitempty"`

	// The settings for cobrowse
	Cobrowse *CobrowseSettings `json:"cobrowse,omitempty"`

	// A reference to the user who created the configuration version
	// Read Only: true
	CreatedUser *AddressableEntityRef `json:"createdUser,omitempty"`

	// The date the configuration version was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date the configuration version was most recently modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The date the configuration version was most recently published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DatePublished strfmt.DateTime `json:"datePublished,omitempty"`

	// The default language to use for the configuration required if the messenger is enabled
	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	// The description of the configuration
	Description string `json:"description,omitempty"`

	// The configuration version ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The settings for journey events
	JourneyEvents *JourneyEventsSettings `json:"journeyEvents,omitempty"`

	// A list of languages supported on the configuration required if the messenger is enabled
	Languages []string `json:"languages"`

	// A reference to the user who most recently modified the configuration version
	// Read Only: true
	LastModifiedUser *AddressableEntityRef `json:"lastModifiedUser,omitempty"`

	// The settings for messenger
	Messenger *MessengerSettings `json:"messenger,omitempty"`

	// The configuration version name
	// Required: true
	Name *string `json:"name"`

	// The settings for position
	Position *PositionSettings `json:"position,omitempty"`

	// A reference to the user who published the configuration version
	// Read Only: true
	PublishedUser *AddressableEntityRef `json:"publishedUser,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The current status of the configuration version
	// Enum: [Pending Active Inactive Error Deleting]
	Status string `json:"status,omitempty"`

	// The settings for support center
	SupportCenter *SupportCenterSettings `json:"supportCenter,omitempty"`

	// The version of the configuration
	// Read Only: true
	Version string `json:"version,omitempty"`
}

// Validate validates this web deployment configuration version
func (m *WebDeploymentConfigurationVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCobrowse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatePublished(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJourneyEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessenger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportCenter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebDeploymentConfigurationVersion) validateAuthenticationSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationSettings) { // not required
		return nil
	}

	if m.AuthenticationSettings != nil {
		if err := m.AuthenticationSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticationSettings")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateCobrowse(formats strfmt.Registry) error {

	if swag.IsZero(m.Cobrowse) { // not required
		return nil
	}

	if m.Cobrowse != nil {
		if err := m.Cobrowse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cobrowse")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateCreatedUser(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedUser) { // not required
		return nil
	}

	if m.CreatedUser != nil {
		if err := m.CreatedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdUser")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateDatePublished(formats strfmt.Registry) error {

	if swag.IsZero(m.DatePublished) { // not required
		return nil
	}

	if err := validate.FormatOf("datePublished", "body", "date-time", m.DatePublished.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateJourneyEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.JourneyEvents) { // not required
		return nil
	}

	if m.JourneyEvents != nil {
		if err := m.JourneyEvents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("journeyEvents")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateLastModifiedUser(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedUser) { // not required
		return nil
	}

	if m.LastModifiedUser != nil {
		if err := m.LastModifiedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastModifiedUser")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateMessenger(formats strfmt.Registry) error {

	if swag.IsZero(m.Messenger) { // not required
		return nil
	}

	if m.Messenger != nil {
		if err := m.Messenger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messenger")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validatePublishedUser(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishedUser) { // not required
		return nil
	}

	if m.PublishedUser != nil {
		if err := m.PublishedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publishedUser")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var webDeploymentConfigurationVersionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Active","Inactive","Error","Deleting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webDeploymentConfigurationVersionTypeStatusPropEnum = append(webDeploymentConfigurationVersionTypeStatusPropEnum, v)
	}
}

const (

	// WebDeploymentConfigurationVersionStatusPending captures enum value "Pending"
	WebDeploymentConfigurationVersionStatusPending string = "Pending"

	// WebDeploymentConfigurationVersionStatusActive captures enum value "Active"
	WebDeploymentConfigurationVersionStatusActive string = "Active"

	// WebDeploymentConfigurationVersionStatusInactive captures enum value "Inactive"
	WebDeploymentConfigurationVersionStatusInactive string = "Inactive"

	// WebDeploymentConfigurationVersionStatusError captures enum value "Error"
	WebDeploymentConfigurationVersionStatusError string = "Error"

	// WebDeploymentConfigurationVersionStatusDeleting captures enum value "Deleting"
	WebDeploymentConfigurationVersionStatusDeleting string = "Deleting"
)

// prop value enum
func (m *WebDeploymentConfigurationVersion) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webDeploymentConfigurationVersionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebDeploymentConfigurationVersion) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WebDeploymentConfigurationVersion) validateSupportCenter(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportCenter) { // not required
		return nil
	}

	if m.SupportCenter != nil {
		if err := m.SupportCenter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supportCenter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebDeploymentConfigurationVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebDeploymentConfigurationVersion) UnmarshalBinary(b []byte) error {
	var res WebDeploymentConfigurationVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
