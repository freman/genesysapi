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

// Greeting greeting
//
// swagger:model Greeting
type Greeting struct {

	// audio file
	AudioFile *GreetingAudioFile `json:"audioFile,omitempty"`

	// audio t t s
	AudioTTS string `json:"audioTTS,omitempty"`

	// created by
	// Format: uri
	CreatedBy strfmt.URI `json:"createdBy,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// modified by
	// Format: uri
	ModifiedBy strfmt.URI `json:"modifiedBy,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Greeting owner
	// Required: true
	Owner *DomainEntity `json:"owner"`

	// Greeting owner type
	// Required: true
	// Enum: [USER ORGANIZATION GROUP]
	OwnerType *string `json:"ownerType"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Greeting type
	// Required: true
	// Enum: [STATION VOICEMAIL NAME]
	Type *string `json:"type"`
}

// Validate validates this greeting
func (m *Greeting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudioFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerType(formats); err != nil {
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

func (m *Greeting) validateAudioFile(formats strfmt.Registry) error {

	if swag.IsZero(m.AudioFile) { // not required
		return nil
	}

	if m.AudioFile != nil {
		if err := m.AudioFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audioFile")
			}
			return err
		}
	}

	return nil
}

func (m *Greeting) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("createdBy", "body", "uri", m.CreatedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Greeting) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Greeting) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedBy", "body", "uri", m.ModifiedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Greeting) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Greeting) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

var greetingTypeOwnerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USER","ORGANIZATION","GROUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		greetingTypeOwnerTypePropEnum = append(greetingTypeOwnerTypePropEnum, v)
	}
}

const (

	// GreetingOwnerTypeUSER captures enum value "USER"
	GreetingOwnerTypeUSER string = "USER"

	// GreetingOwnerTypeORGANIZATION captures enum value "ORGANIZATION"
	GreetingOwnerTypeORGANIZATION string = "ORGANIZATION"

	// GreetingOwnerTypeGROUP captures enum value "GROUP"
	GreetingOwnerTypeGROUP string = "GROUP"
)

// prop value enum
func (m *Greeting) validateOwnerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, greetingTypeOwnerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Greeting) validateOwnerType(formats strfmt.Registry) error {

	if err := validate.Required("ownerType", "body", m.OwnerType); err != nil {
		return err
	}

	// value enum
	if err := m.validateOwnerTypeEnum("ownerType", "body", *m.OwnerType); err != nil {
		return err
	}

	return nil
}

func (m *Greeting) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var greetingTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATION","VOICEMAIL","NAME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		greetingTypeTypePropEnum = append(greetingTypeTypePropEnum, v)
	}
}

const (

	// GreetingTypeSTATION captures enum value "STATION"
	GreetingTypeSTATION string = "STATION"

	// GreetingTypeVOICEMAIL captures enum value "VOICEMAIL"
	GreetingTypeVOICEMAIL string = "VOICEMAIL"

	// GreetingTypeNAME captures enum value "NAME"
	GreetingTypeNAME string = "NAME"
)

// prop value enum
func (m *Greeting) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, greetingTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Greeting) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Greeting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Greeting) UnmarshalBinary(b []byte) error {
	var res Greeting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
