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

// User user
//
// swagger:model User
type User struct {

	// acd auto answer
	AcdAutoAnswer bool `json:"acdAutoAnswer"`

	// Email addresses and phone numbers for this user
	Addresses []*Contact `json:"addresses"`

	// Roles and permissions assigned to the user
	// Read Only: true
	Authorization *UserAuthorization `json:"authorization,omitempty"`

	// biography
	Biography *Biography `json:"biography,omitempty"`

	// certifications
	Certifications []string `json:"certifications"`

	// chat
	Chat *Chat `json:"chat,omitempty"`

	// Summary of conversion statistics for conversation types.
	// Read Only: true
	ConversationSummary *UserConversationSummary `json:"conversationSummary,omitempty"`

	// The last time the user logged in using username and password. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateLastLogin strfmt.DateTime `json:"dateLastLogin,omitempty"`

	// department
	Department string `json:"department,omitempty"`

	// The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// employer info
	EmployerInfo *EmployerInfo `json:"employerInfo,omitempty"`

	// Current geolocation position
	// Read Only: true
	Geolocation *Geolocation `json:"geolocation,omitempty"`

	// The groups the user is a member of
	// Read Only: true
	Groups []*Group `json:"groups"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// images
	Images []*UserImage `json:"images"`

	// Integration presence
	// Read Only: true
	IntegrationPresence *UserPresence `json:"integrationPresence,omitempty"`

	// preferred language by the user
	// Read Only: true
	LanguagePreference string `json:"languagePreference,omitempty"`

	// Routing (ACD) languages possessed by the user
	// Read Only: true
	Languages []*UserRoutingLanguage `json:"languages"`

	// last token issued
	LastTokenIssued *OAuthLastTokenIssued `json:"lastTokenIssued,omitempty"`

	// The user placement at each site location.
	// Read Only: true
	Locations []*Location `json:"locations"`

	// manager
	Manager *User `json:"manager,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Determine if out of office is enabled
	// Read Only: true
	OutOfOffice *OutOfOffice `json:"outOfOffice,omitempty"`

	// Active presence
	// Read Only: true
	Presence *UserPresence `json:"presence,omitempty"`

	// Auto populated from addresses.
	// Read Only: true
	PrimaryContactInfo []*Contact `json:"primaryContactInfo"`

	// Profile skills possessed by the user
	// Read Only: true
	ProfileSkills []string `json:"profileSkills"`

	// ACD routing status
	// Read Only: true
	RoutingStatus *RoutingStatus `json:"routingStatus,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Routing (ACD) skills possessed by the user
	// Read Only: true
	Skills []*UserRoutingSkill `json:"skills"`

	// The current state for this user.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// Effective, default, and last station information
	// Read Only: true
	Station *UserStations `json:"station,omitempty"`

	// The team the user is a member of
	// Read Only: true
	Team *Team `json:"team,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// Required when updating a user, this value should be the current version of the user.  The current version can be obtained with a GET on the user before doing a PATCH.
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBiography(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateLastLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmployerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeolocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationPresence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTokenIssued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutOfOffice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryContactInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkills(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
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

func (m *User) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateAuthorization(formats strfmt.Registry) error {

	if swag.IsZero(m.Authorization) { // not required
		return nil
	}

	if m.Authorization != nil {
		if err := m.Authorization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorization")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateBiography(formats strfmt.Registry) error {

	if swag.IsZero(m.Biography) { // not required
		return nil
	}

	if m.Biography != nil {
		if err := m.Biography.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("biography")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateChat(formats strfmt.Registry) error {

	if swag.IsZero(m.Chat) { // not required
		return nil
	}

	if m.Chat != nil {
		if err := m.Chat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chat")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateConversationSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationSummary) { // not required
		return nil
	}

	if m.ConversationSummary != nil {
		if err := m.ConversationSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversationSummary")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateDateLastLogin(formats strfmt.Registry) error {

	if swag.IsZero(m.DateLastLogin) { // not required
		return nil
	}

	if err := validate.FormatOf("dateLastLogin", "body", "date-time", m.DateLastLogin.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateDivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateEmployerInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.EmployerInfo) { // not required
		return nil
	}

	if m.EmployerInfo != nil {
		if err := m.EmployerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("employerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateGeolocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Geolocation) { // not required
		return nil
	}

	if m.Geolocation != nil {
		if err := m.Geolocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geolocation")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateIntegrationPresence(formats strfmt.Registry) error {

	if swag.IsZero(m.IntegrationPresence) { // not required
		return nil
	}

	if m.IntegrationPresence != nil {
		if err := m.IntegrationPresence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrationPresence")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateLanguages(formats strfmt.Registry) error {

	if swag.IsZero(m.Languages) { // not required
		return nil
	}

	for i := 0; i < len(m.Languages); i++ {
		if swag.IsZero(m.Languages[i]) { // not required
			continue
		}

		if m.Languages[i] != nil {
			if err := m.Languages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateLastTokenIssued(formats strfmt.Registry) error {

	if swag.IsZero(m.LastTokenIssued) { // not required
		return nil
	}

	if m.LastTokenIssued != nil {
		if err := m.LastTokenIssued.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastTokenIssued")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	for i := 0; i < len(m.Locations); i++ {
		if swag.IsZero(m.Locations[i]) { // not required
			continue
		}

		if m.Locations[i] != nil {
			if err := m.Locations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateManager(formats strfmt.Registry) error {

	if swag.IsZero(m.Manager) { // not required
		return nil
	}

	if m.Manager != nil {
		if err := m.Manager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manager")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateOutOfOffice(formats strfmt.Registry) error {

	if swag.IsZero(m.OutOfOffice) { // not required
		return nil
	}

	if m.OutOfOffice != nil {
		if err := m.OutOfOffice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outOfOffice")
			}
			return err
		}
	}

	return nil
}

func (m *User) validatePresence(formats strfmt.Registry) error {

	if swag.IsZero(m.Presence) { // not required
		return nil
	}

	if m.Presence != nil {
		if err := m.Presence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("presence")
			}
			return err
		}
	}

	return nil
}

func (m *User) validatePrimaryContactInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryContactInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.PrimaryContactInfo); i++ {
		if swag.IsZero(m.PrimaryContactInfo[i]) { // not required
			continue
		}

		if m.PrimaryContactInfo[i] != nil {
			if err := m.PrimaryContactInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("primaryContactInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateRoutingStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingStatus) { // not required
		return nil
	}

	if m.RoutingStatus != nil {
		if err := m.RoutingStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routingStatus")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateSkills(formats strfmt.Registry) error {

	if swag.IsZero(m.Skills) { // not required
		return nil
	}

	for i := 0; i < len(m.Skills); i++ {
		if swag.IsZero(m.Skills[i]) { // not required
			continue
		}

		if m.Skills[i] != nil {
			if err := m.Skills[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var userTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userTypeStatePropEnum = append(userTypeStatePropEnum, v)
	}
}

const (

	// UserStateActive captures enum value "active"
	UserStateActive string = "active"

	// UserStateInactive captures enum value "inactive"
	UserStateInactive string = "inactive"

	// UserStateDeleted captures enum value "deleted"
	UserStateDeleted string = "deleted"
)

// prop value enum
func (m *User) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *User) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *User) validateStation(formats strfmt.Registry) error {

	if swag.IsZero(m.Station) { // not required
		return nil
	}

	if m.Station != nil {
		if err := m.Station.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("station")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateTeam(formats strfmt.Registry) error {

	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
