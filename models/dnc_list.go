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

// DncList dnc list
//
// swagger:model DncList
type DncList struct {

	// The contact method. Required if dncSourceType is rds.
	// Enum: [Email Phone]
	ContactMethod string `json:"contactMethod,omitempty"`

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The division this DncList belongs to.
	Division *DomainEntityRef `json:"division,omitempty"`

	// The list of dnc.com codes to be treated as DNC. Required if the dncSourceType is dnc.com.
	// Unique: true
	DncCodes []string `json:"dncCodes"`

	// The type of the DncList.
	// Required: true
	// Read Only: true
	// Enum: [rds dnc.com gryphon]
	DncSourceType string `json:"dncSourceType"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The status of the import process
	// Read Only: true
	ImportStatus *ImportStatus `json:"importStatus,omitempty"`

	// A gryphon license number. Required if the dncSourceType is gryphon.
	LicenseID string `json:"licenseId,omitempty"`

	// A dnc.com loginId. Required if the dncSourceType is dnc.com.
	LoginID string `json:"loginId,omitempty"`

	// The name of the DncList.
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The total number of phone numbers in the DncList.
	// Read Only: true
	Size int64 `json:"size,omitempty"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this dnc list
func (m *DncList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactMethod(formats); err != nil {
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

	if err := m.validateDncCodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDncSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportStatus(formats); err != nil {
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

var dncListTypeContactMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Email","Phone"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dncListTypeContactMethodPropEnum = append(dncListTypeContactMethodPropEnum, v)
	}
}

const (

	// DncListContactMethodEmail captures enum value "Email"
	DncListContactMethodEmail string = "Email"

	// DncListContactMethodPhone captures enum value "Phone"
	DncListContactMethodPhone string = "Phone"
)

// prop value enum
func (m *DncList) validateContactMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dncListTypeContactMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DncList) validateContactMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateContactMethodEnum("contactMethod", "body", m.ContactMethod); err != nil {
		return err
	}

	return nil
}

func (m *DncList) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DncList) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DncList) validateDivision(formats strfmt.Registry) error {

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

func (m *DncList) validateDncCodes(formats strfmt.Registry) error {

	if swag.IsZero(m.DncCodes) { // not required
		return nil
	}

	if err := validate.UniqueItems("dncCodes", "body", m.DncCodes); err != nil {
		return err
	}

	return nil
}

var dncListTypeDncSourceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rds","dnc.com","gryphon"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dncListTypeDncSourceTypePropEnum = append(dncListTypeDncSourceTypePropEnum, v)
	}
}

const (

	// DncListDncSourceTypeRds captures enum value "rds"
	DncListDncSourceTypeRds string = "rds"

	// DncListDncSourceTypeDncCom captures enum value "dnc.com"
	DncListDncSourceTypeDncCom string = "dnc.com"

	// DncListDncSourceTypeGryphon captures enum value "gryphon"
	DncListDncSourceTypeGryphon string = "gryphon"
)

// prop value enum
func (m *DncList) validateDncSourceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dncListTypeDncSourceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DncList) validateDncSourceType(formats strfmt.Registry) error {

	if err := validate.RequiredString("dncSourceType", "body", string(m.DncSourceType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateDncSourceTypeEnum("dncSourceType", "body", m.DncSourceType); err != nil {
		return err
	}

	return nil
}

func (m *DncList) validateImportStatus(formats strfmt.Registry) error {

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

func (m *DncList) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DncList) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DncList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DncList) UnmarshalBinary(b []byte) error {
	var res DncList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
