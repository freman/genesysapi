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

// ContactListDivisionView contact list division view
//
// swagger:model ContactListDivisionView
type ContactListDivisionView struct {

	// The names of the contact data columns.
	// Required: true
	ColumnNames []string `json:"columnNames"`

	// The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The status of the import process.
	// Read Only: true
	ImportStatus *ImportStatus `json:"importStatus,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Indicates which columns are phone numbers.
	// Required: true
	PhoneColumns []*ContactPhoneNumberColumn `json:"phoneColumns"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The number of contacts in the ContactList.
	// Read Only: true
	Size int64 `json:"size,omitempty"`
}

// Validate validates this contact list division view
func (m *ContactListDivisionView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumnNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneColumns(formats); err != nil {
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

func (m *ContactListDivisionView) validateColumnNames(formats strfmt.Registry) error {

	if err := validate.Required("columnNames", "body", m.ColumnNames); err != nil {
		return err
	}

	return nil
}

func (m *ContactListDivisionView) validateDivision(formats strfmt.Registry) error {

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

func (m *ContactListDivisionView) validateImportStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ImportStatus) { // not required
		return nil
	}

	if m.ImportStatus != nil {
		if err := m.ImportStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("importStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ContactListDivisionView) validatePhoneColumns(formats strfmt.Registry) error {

	if err := validate.Required("phoneColumns", "body", m.PhoneColumns); err != nil {
		return err
	}

	for i := 0; i < len(m.PhoneColumns); i++ {
		if swag.IsZero(m.PhoneColumns[i]) { // not required
			continue
		}

		if m.PhoneColumns[i] != nil {
			if err := m.PhoneColumns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContactListDivisionView) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactListDivisionView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactListDivisionView) UnmarshalBinary(b []byte) error {
	var res ContactListDivisionView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
