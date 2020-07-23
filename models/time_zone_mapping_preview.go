// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeZoneMappingPreview time zone mapping preview
//
// swagger:model TimeZoneMappingPreview
type TimeZoneMappingPreview struct {

	// The associated ContactList
	ContactList *DomainEntityRef `json:"contactList,omitempty"`

	// The total number of contacts in the contact list
	ContactListSize int64 `json:"contactListSize,omitempty"`

	// The total number of contacts that will be dialed during the default window
	ContactsInDefaultWindow int64 `json:"contactsInDefaultWindow,omitempty"`

	// The total number of contacts that mapped to a single time zone
	ContactsMappedToASingleZone int64 `json:"contactsMappedToASingleZone,omitempty"`

	// The total number of contacts that mapped to a single time zone and were mapped using the zip code column
	ContactsMappedToASingleZoneUsingZipCode int64 `json:"contactsMappedToASingleZoneUsingZipCode,omitempty"`

	// The total number of contacts that mapped to multiple time zones
	ContactsMappedToMultipleZones int64 `json:"contactsMappedToMultipleZones,omitempty"`

	// The total number of contacts that mapped to multiple time zones and were mapped using the zip code column
	ContactsMappedToMultipleZonesUsingZipCode int64 `json:"contactsMappedToMultipleZonesUsingZipCode,omitempty"`

	// The number of contacts per time zone that mapped to only that time zone and were mapped using the zip code column
	ContactsMappedUsingZipCode map[string]int64 `json:"contactsMappedUsingZipCode,omitempty"`

	// The number of contacts per time zone that mapped to only that time zone
	ContactsPerTimeZone map[string]int64 `json:"contactsPerTimeZone,omitempty"`
}

// Validate validates this time zone mapping preview
func (m *TimeZoneMappingPreview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeZoneMappingPreview) validateContactList(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactList) { // not required
		return nil
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

// MarshalBinary interface implementation
func (m *TimeZoneMappingPreview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeZoneMappingPreview) UnmarshalBinary(b []byte) error {
	var res TimeZoneMappingPreview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
