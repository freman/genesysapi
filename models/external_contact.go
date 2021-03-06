// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExternalContact external contact
//
// swagger:model ExternalContact
type ExternalContact struct {

	// address
	Address *ContactAddress `json:"address,omitempty"`

	// cell phone
	CellPhone *PhoneNumber `json:"cellPhone,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreateDate strfmt.DateTime `json:"createDate,omitempty"`

	// Custom fields defined in the schema referenced by schemaId and schemaVersion.
	CustomFields map[string]interface{} `json:"customFields,omitempty"`

	// Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	// Read Only: true
	ExternalDataSources []*ExternalDataSource `json:"externalDataSources"`

	// external organization
	ExternalOrganization *ExternalOrganization `json:"externalOrganization,omitempty"`

	// A string that identifies an external system-of-record resource that may have more detailed information on the contact. It should be a valid URL (including the http/https protocol, port, and path [if any]). The value is automatically trimmed of any leading and trailing whitespace.
	ExternalSystemURL string `json:"externalSystemUrl,omitempty"`

	// facebook Id
	FacebookID *FacebookID `json:"facebookId,omitempty"`

	// The first name of the contact.
	// Required: true
	FirstName *string `json:"firstName"`

	// home phone
	HomePhone *PhoneNumber `json:"homePhone,omitempty"`

	// The globally unique identifier for the object.
	ID string `json:"id,omitempty"`

	// The last name of the contact.
	// Required: true
	LastName *string `json:"lastName"`

	// line Id
	LineID *LineID `json:"lineId,omitempty"`

	// middle name
	MiddleName string `json:"middleName,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ModifyDate strfmt.DateTime `json:"modifyDate,omitempty"`

	// other email
	OtherEmail string `json:"otherEmail,omitempty"`

	// other phone
	OtherPhone *PhoneNumber `json:"otherPhone,omitempty"`

	// personal email
	PersonalEmail string `json:"personalEmail,omitempty"`

	// salutation
	Salutation string `json:"salutation,omitempty"`

	// The schema defining custom fields for this contact
	Schema *DataSchema `json:"schema,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// survey opt out
	SurveyOptOut bool `json:"surveyOptOut"`

	// title
	Title string `json:"title,omitempty"`

	// twitter Id
	TwitterID *TwitterID `json:"twitterId,omitempty"`

	// whats app Id
	WhatsAppID *WhatsAppID `json:"whatsAppId,omitempty"`

	// work email
	WorkEmail string `json:"workEmail,omitempty"`

	// work phone
	WorkPhone *PhoneNumber `json:"workPhone,omitempty"`
}

// Validate validates this external contact
func (m *ExternalContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCellPhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDataSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacebookID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomePhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifyDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherPhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTwitterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhatsAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkPhone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalContact) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateCellPhone(formats strfmt.Registry) error {

	if swag.IsZero(m.CellPhone) { // not required
		return nil
	}

	if m.CellPhone != nil {
		if err := m.CellPhone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cellPhone")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateCreateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createDate", "body", "date-time", m.CreateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExternalContact) validateExternalDataSources(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalDataSources) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalDataSources); i++ {
		if swag.IsZero(m.ExternalDataSources[i]) { // not required
			continue
		}

		if m.ExternalDataSources[i] != nil {
			if err := m.ExternalDataSources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalDataSources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExternalContact) validateExternalOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalOrganization) { // not required
		return nil
	}

	if m.ExternalOrganization != nil {
		if err := m.ExternalOrganization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalOrganization")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateFacebookID(formats strfmt.Registry) error {

	if swag.IsZero(m.FacebookID) { // not required
		return nil
	}

	if m.FacebookID != nil {
		if err := m.FacebookID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facebookId")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *ExternalContact) validateHomePhone(formats strfmt.Registry) error {

	if swag.IsZero(m.HomePhone) { // not required
		return nil
	}

	if m.HomePhone != nil {
		if err := m.HomePhone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("homePhone")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *ExternalContact) validateLineID(formats strfmt.Registry) error {

	if swag.IsZero(m.LineID) { // not required
		return nil
	}

	if m.LineID != nil {
		if err := m.LineID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lineId")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateModifyDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifyDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifyDate", "body", "date-time", m.ModifyDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExternalContact) validateOtherPhone(formats strfmt.Registry) error {

	if swag.IsZero(m.OtherPhone) { // not required
		return nil
	}

	if m.OtherPhone != nil {
		if err := m.OtherPhone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("otherPhone")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExternalContact) validateTwitterID(formats strfmt.Registry) error {

	if swag.IsZero(m.TwitterID) { // not required
		return nil
	}

	if m.TwitterID != nil {
		if err := m.TwitterID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("twitterId")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateWhatsAppID(formats strfmt.Registry) error {

	if swag.IsZero(m.WhatsAppID) { // not required
		return nil
	}

	if m.WhatsAppID != nil {
		if err := m.WhatsAppID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("whatsAppId")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalContact) validateWorkPhone(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkPhone) { // not required
		return nil
	}

	if m.WorkPhone != nil {
		if err := m.WorkPhone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workPhone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalContact) UnmarshalBinary(b []byte) error {
	var res ExternalContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
