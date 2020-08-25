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

// KnowledgeBase knowledge base
//
// swagger:model KnowledgeBase
type KnowledgeBase struct {

	// Core language for knowledge base in which initial content must be created first
	// Required: true
	// Enum: [en-US]
	CoreLanguage *string `json:"coreLanguage"`

	// Knowledge base creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Knowledge base last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Knowledge base description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this knowledge base
func (m *KnowledgeBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoreLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
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

var knowledgeBaseTypeCoreLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-US"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeBaseTypeCoreLanguagePropEnum = append(knowledgeBaseTypeCoreLanguagePropEnum, v)
	}
}

const (

	// KnowledgeBaseCoreLanguageEnUS captures enum value "en-US"
	KnowledgeBaseCoreLanguageEnUS string = "en-US"
)

// prop value enum
func (m *KnowledgeBase) validateCoreLanguageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeBaseTypeCoreLanguagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeBase) validateCoreLanguage(formats strfmt.Registry) error {

	if err := validate.Required("coreLanguage", "body", m.CoreLanguage); err != nil {
		return err
	}

	// value enum
	if err := m.validateCoreLanguageEnum("coreLanguage", "body", *m.CoreLanguage); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeBase) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeBase) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeBase) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeBase) UnmarshalBinary(b []byte) error {
	var res KnowledgeBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}