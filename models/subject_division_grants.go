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

// SubjectDivisionGrants subject division grants
//
// swagger:model SubjectDivisionGrants
type SubjectDivisionGrants struct {

	// divisions
	Divisions []*Division `json:"divisions"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// type
	// Enum: [PC_USER PC_GROUP PC_OAUTH_CLIENT PC_TRUSTEE_USER PC_TRUSTEE_GROUP UNKNOWN]
	Type string `json:"type,omitempty"`
}

// Validate validates this subject division grants
func (m *SubjectDivisionGrants) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectDivisionGrants) validateDivisions(formats strfmt.Registry) error {

	if swag.IsZero(m.Divisions) { // not required
		return nil
	}

	for i := 0; i < len(m.Divisions); i++ {
		if swag.IsZero(m.Divisions[i]) { // not required
			continue
		}

		if m.Divisions[i] != nil {
			if err := m.Divisions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("divisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubjectDivisionGrants) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var subjectDivisionGrantsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PC_USER","PC_GROUP","PC_OAUTH_CLIENT","PC_TRUSTEE_USER","PC_TRUSTEE_GROUP","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subjectDivisionGrantsTypeTypePropEnum = append(subjectDivisionGrantsTypeTypePropEnum, v)
	}
}

const (

	// SubjectDivisionGrantsTypePCUSER captures enum value "PC_USER"
	SubjectDivisionGrantsTypePCUSER string = "PC_USER"

	// SubjectDivisionGrantsTypePCGROUP captures enum value "PC_GROUP"
	SubjectDivisionGrantsTypePCGROUP string = "PC_GROUP"

	// SubjectDivisionGrantsTypePCOAUTHCLIENT captures enum value "PC_OAUTH_CLIENT"
	SubjectDivisionGrantsTypePCOAUTHCLIENT string = "PC_OAUTH_CLIENT"

	// SubjectDivisionGrantsTypePCTRUSTEEUSER captures enum value "PC_TRUSTEE_USER"
	SubjectDivisionGrantsTypePCTRUSTEEUSER string = "PC_TRUSTEE_USER"

	// SubjectDivisionGrantsTypePCTRUSTEEGROUP captures enum value "PC_TRUSTEE_GROUP"
	SubjectDivisionGrantsTypePCTRUSTEEGROUP string = "PC_TRUSTEE_GROUP"

	// SubjectDivisionGrantsTypeUNKNOWN captures enum value "UNKNOWN"
	SubjectDivisionGrantsTypeUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *SubjectDivisionGrants) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, subjectDivisionGrantsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SubjectDivisionGrants) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubjectDivisionGrants) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectDivisionGrants) UnmarshalBinary(b []byte) error {
	var res SubjectDivisionGrants
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}