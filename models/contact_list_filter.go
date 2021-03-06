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

// ContactListFilter contact list filter
//
// swagger:model ContactListFilter
type ContactListFilter struct {

	// Groups of conditions to filter the contacts by.
	Clauses []*ContactListFilterClause `json:"clauses"`

	// The contact list the filter is based on.
	// Required: true
	ContactList *DomainEntityRef `json:"contactList"`

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// How to join clauses together.
	// Enum: [AND OR]
	FilterType string `json:"filterType,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the list.
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this contact list filter
func (m *ContactListFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClauses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *ContactListFilter) validateClauses(formats strfmt.Registry) error {

	if swag.IsZero(m.Clauses) { // not required
		return nil
	}

	for i := 0; i < len(m.Clauses); i++ {
		if swag.IsZero(m.Clauses[i]) { // not required
			continue
		}

		if m.Clauses[i] != nil {
			if err := m.Clauses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clauses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactListFilter) validateContactList(formats strfmt.Registry) error {

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

func (m *ContactListFilter) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactListFilter) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

var contactListFilterTypeFilterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AND","OR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactListFilterTypeFilterTypePropEnum = append(contactListFilterTypeFilterTypePropEnum, v)
	}
}

const (

	// ContactListFilterFilterTypeAND captures enum value "AND"
	ContactListFilterFilterTypeAND string = "AND"

	// ContactListFilterFilterTypeOR captures enum value "OR"
	ContactListFilterFilterTypeOR string = "OR"
)

// prop value enum
func (m *ContactListFilter) validateFilterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contactListFilterTypeFilterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContactListFilter) validateFilterType(formats strfmt.Registry) error {

	if swag.IsZero(m.FilterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFilterTypeEnum("filterType", "body", m.FilterType); err != nil {
		return err
	}

	return nil
}

func (m *ContactListFilter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ContactListFilter) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactListFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactListFilter) UnmarshalBinary(b []byte) error {
	var res ContactListFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
