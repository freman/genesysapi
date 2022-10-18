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

// MessagingCampaign messaging campaign
//
// swagger:model MessagingCampaign
type MessagingCampaign struct {

	// Whether this messaging campaign is always running
	AlwaysRunning bool `json:"alwaysRunning"`

	// The callable time set for this messaging campaign.
	CallableTimeSet *DomainEntityRef `json:"callableTimeSet,omitempty"`

	// The current status of the messaging campaign. A messaging campaign may be turned 'on' or 'off'.
	// Enum: [on stopping off complete invalid]
	CampaignStatus string `json:"campaignStatus,omitempty"`

	// The contact list that this messaging campaign will send messages for.
	// Required: true
	ContactList *DomainEntityRef `json:"contactList"`

	// The contact list filter to check before sending a message for this messaging campaign.
	ContactListFilters []*DomainEntityRef `json:"contactListFilters"`

	// The order in which to sort contacts for dialing, based on up to four columns.
	ContactSorts []*ContactSort `json:"contactSorts"`

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The division this entity belongs to.
	Division *DomainEntityRef `json:"division,omitempty"`

	// The dnc lists to check before sending a message for this messaging campaign.
	DncLists []*DomainEntityRef `json:"dncLists"`

	// Configuration for this messaging campaign to send Email messages.
	EmailConfig *EmailConfig `json:"emailConfig,omitempty"`

	// A list of current error conditions associated with this messaging campaign.
	Errors []*RestErrorDetail `json:"errors"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// How many messages this messaging campaign will send per minute.
	// Required: true
	MessagesPerMinute *int32 `json:"messagesPerMinute"`

	// name
	Name string `json:"name,omitempty"`

	// Rule Sets to be applied while this campaign is sending messages
	RuleSets []*DomainEntityRef `json:"ruleSets"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Configuration for this messaging campaign to send SMS messages.
	SmsConfig *SmsConfig `json:"smsConfig,omitempty"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this messaging campaign
func (m *MessagingCampaign) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallableTimeSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCampaignStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactListFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactSorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDncLists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessagesPerMinute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmsConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessagingCampaign) validateCallableTimeSet(formats strfmt.Registry) error {

	if swag.IsZero(m.CallableTimeSet) { // not required
		return nil
	}

	if m.CallableTimeSet != nil {
		if err := m.CallableTimeSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callableTimeSet")
			}
			return err
		}
	}

	return nil
}

var messagingCampaignTypeCampaignStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","stopping","off","complete","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messagingCampaignTypeCampaignStatusPropEnum = append(messagingCampaignTypeCampaignStatusPropEnum, v)
	}
}

const (

	// MessagingCampaignCampaignStatusOn captures enum value "on"
	MessagingCampaignCampaignStatusOn string = "on"

	// MessagingCampaignCampaignStatusStopping captures enum value "stopping"
	MessagingCampaignCampaignStatusStopping string = "stopping"

	// MessagingCampaignCampaignStatusOff captures enum value "off"
	MessagingCampaignCampaignStatusOff string = "off"

	// MessagingCampaignCampaignStatusComplete captures enum value "complete"
	MessagingCampaignCampaignStatusComplete string = "complete"

	// MessagingCampaignCampaignStatusInvalid captures enum value "invalid"
	MessagingCampaignCampaignStatusInvalid string = "invalid"
)

// prop value enum
func (m *MessagingCampaign) validateCampaignStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messagingCampaignTypeCampaignStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessagingCampaign) validateCampaignStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CampaignStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCampaignStatusEnum("campaignStatus", "body", m.CampaignStatus); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaign) validateContactList(formats strfmt.Registry) error {

	if err := validate.Required("contactList", "body", m.ContactList); err != nil {
		return err
	}

	if m.ContactList != nil {
		if err := m.ContactList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactList")
			}
			return err
		}
	}

	return nil
}

func (m *MessagingCampaign) validateContactListFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactListFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.ContactListFilters); i++ {
		if swag.IsZero(m.ContactListFilters[i]) { // not required
			continue
		}

		if m.ContactListFilters[i] != nil {
			if err := m.ContactListFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contactListFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagingCampaign) validateContactSorts(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactSorts) { // not required
		return nil
	}

	for i := 0; i < len(m.ContactSorts); i++ {
		if swag.IsZero(m.ContactSorts[i]) { // not required
			continue
		}

		if m.ContactSorts[i] != nil {
			if err := m.ContactSorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contactSorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagingCampaign) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaign) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaign) validateDivision(formats strfmt.Registry) error {

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

func (m *MessagingCampaign) validateDncLists(formats strfmt.Registry) error {

	if swag.IsZero(m.DncLists) { // not required
		return nil
	}

	for i := 0; i < len(m.DncLists); i++ {
		if swag.IsZero(m.DncLists[i]) { // not required
			continue
		}

		if m.DncLists[i] != nil {
			if err := m.DncLists[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dncLists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagingCampaign) validateEmailConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.EmailConfig) { // not required
		return nil
	}

	if m.EmailConfig != nil {
		if err := m.EmailConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailConfig")
			}
			return err
		}
	}

	return nil
}

func (m *MessagingCampaign) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagingCampaign) validateMessagesPerMinute(formats strfmt.Registry) error {

	if err := validate.Required("messagesPerMinute", "body", m.MessagesPerMinute); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaign) validateRuleSets(formats strfmt.Registry) error {

	if swag.IsZero(m.RuleSets) { // not required
		return nil
	}

	for i := 0; i < len(m.RuleSets); i++ {
		if swag.IsZero(m.RuleSets[i]) { // not required
			continue
		}

		if m.RuleSets[i] != nil {
			if err := m.RuleSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ruleSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagingCampaign) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaign) validateSmsConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SmsConfig) { // not required
		return nil
	}

	if m.SmsConfig != nil {
		if err := m.SmsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smsConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessagingCampaign) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessagingCampaign) UnmarshalBinary(b []byte) error {
	var res MessagingCampaign
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
