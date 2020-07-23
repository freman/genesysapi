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

// KeywordSet keyword set
//
// swagger:model KeywordSet
type KeywordSet struct {

	// agents
	// Unique: true
	Agents []*User `json:"agents"`

	// description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The list of keywords to be used for keyword spotting.
	// Required: true
	Keywords []*Keyword `json:"keywords"`

	// Language code, such as 'en-US'
	// Required: true
	Language *string `json:"language"`

	// name
	Name string `json:"name,omitempty"`

	// The types of participants to use keyword spotting on.
	// Required: true
	// Unique: true
	ParticipantPurposes []string `json:"participantPurposes"`

	// queues
	// Unique: true
	Queues []*Queue `json:"queues"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this keyword set
func (m *KeywordSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipantPurposes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueues(formats); err != nil {
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

func (m *KeywordSet) validateAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.Agents) { // not required
		return nil
	}

	if err := validate.UniqueItems("agents", "body", m.Agents); err != nil {
		return err
	}

	for i := 0; i < len(m.Agents); i++ {
		if swag.IsZero(m.Agents[i]) { // not required
			continue
		}

		if m.Agents[i] != nil {
			if err := m.Agents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KeywordSet) validateKeywords(formats strfmt.Registry) error {

	if err := validate.Required("keywords", "body", m.Keywords); err != nil {
		return err
	}

	for i := 0; i < len(m.Keywords); i++ {
		if swag.IsZero(m.Keywords[i]) { // not required
			continue
		}

		if m.Keywords[i] != nil {
			if err := m.Keywords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keywords" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KeywordSet) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

var keywordSetParticipantPurposesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT","CUSTOMER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		keywordSetParticipantPurposesItemsEnum = append(keywordSetParticipantPurposesItemsEnum, v)
	}
}

func (m *KeywordSet) validateParticipantPurposesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, keywordSetParticipantPurposesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KeywordSet) validateParticipantPurposes(formats strfmt.Registry) error {

	if err := validate.Required("participantPurposes", "body", m.ParticipantPurposes); err != nil {
		return err
	}

	if err := validate.UniqueItems("participantPurposes", "body", m.ParticipantPurposes); err != nil {
		return err
	}

	for i := 0; i < len(m.ParticipantPurposes); i++ {

		// value enum
		if err := m.validateParticipantPurposesItemsEnum("participantPurposes"+"."+strconv.Itoa(i), "body", m.ParticipantPurposes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *KeywordSet) validateQueues(formats strfmt.Registry) error {

	if swag.IsZero(m.Queues) { // not required
		return nil
	}

	if err := validate.UniqueItems("queues", "body", m.Queues); err != nil {
		return err
	}

	for i := 0; i < len(m.Queues); i++ {
		if swag.IsZero(m.Queues[i]) { // not required
			continue
		}

		if m.Queues[i] != nil {
			if err := m.Queues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KeywordSet) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeywordSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeywordSet) UnmarshalBinary(b []byte) error {
	var res KeywordSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
