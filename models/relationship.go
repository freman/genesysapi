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

// Relationship relationship
//
// swagger:model Relationship
type Relationship struct {

	// Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	// Read Only: true
	ExternalDataSources []*ExternalDataSource `json:"externalDataSources"`

	// The external organization this relationship is attached to
	// Required: true
	ExternalOrganization *ExternalOrganization `json:"externalOrganization"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The relationship or role of the user to this external organization.Examples: Account Manager, Sales Engineer, Implementation Consultant
	// Required: true
	Relationship *string `json:"relationship"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The user associated with the external organization
	// Required: true
	User *User `json:"user"`
}

// Validate validates this relationship
func (m *Relationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalDataSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Relationship) validateExternalDataSources(formats strfmt.Registry) error {

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

func (m *Relationship) validateExternalOrganization(formats strfmt.Registry) error {

	if err := validate.Required("externalOrganization", "body", m.ExternalOrganization); err != nil {
		return err
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

func (m *Relationship) validateRelationship(formats strfmt.Registry) error {

	if err := validate.Required("relationship", "body", m.Relationship); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Relationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Relationship) UnmarshalBinary(b []byte) error {
	var res Relationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
