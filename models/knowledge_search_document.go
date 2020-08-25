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

// KnowledgeSearchDocument knowledge search document
//
// swagger:model KnowledgeSearchDocument
type KnowledgeSearchDocument struct {

	// Document categories
	Categories []*KnowledgeCategory `json:"categories"`

	// The confidence associated with a document with respect to a search query
	// Read Only: true
	Confidence float64 `json:"confidence,omitempty"`

	// Document creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Document last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// External URL to the document
	ExternalURL string `json:"externalUrl,omitempty"`

	// FAQ document details
	Faq *DocumentFaq `json:"faq,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Knowledge base which document does belong to
	// Read Only: true
	KnowledgeBase *KnowledgeBase `json:"knowledgeBase,omitempty"`

	// Language of the document
	// Required: true
	// Enum: [en-US]
	LanguageCode *string `json:"languageCode"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Document type
	// Required: true
	// Enum: [Faq]
	Type *string `json:"type"`
}

// Validate validates this knowledge search document
func (m *KnowledgeSearchDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKnowledgeBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCode(formats); err != nil {
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

func (m *KnowledgeSearchDocument) validateCategories(formats strfmt.Registry) error {

	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	for i := 0; i < len(m.Categories); i++ {
		if swag.IsZero(m.Categories[i]) { // not required
			continue
		}

		if m.Categories[i] != nil {
			if err := m.Categories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KnowledgeSearchDocument) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeSearchDocument) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeSearchDocument) validateFaq(formats strfmt.Registry) error {

	if swag.IsZero(m.Faq) { // not required
		return nil
	}

	if m.Faq != nil {
		if err := m.Faq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("faq")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeSearchDocument) validateKnowledgeBase(formats strfmt.Registry) error {

	if swag.IsZero(m.KnowledgeBase) { // not required
		return nil
	}

	if m.KnowledgeBase != nil {
		if err := m.KnowledgeBase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("knowledgeBase")
			}
			return err
		}
	}

	return nil
}

var knowledgeSearchDocumentTypeLanguageCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-US"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeSearchDocumentTypeLanguageCodePropEnum = append(knowledgeSearchDocumentTypeLanguageCodePropEnum, v)
	}
}

const (

	// KnowledgeSearchDocumentLanguageCodeEnUS captures enum value "en-US"
	KnowledgeSearchDocumentLanguageCodeEnUS string = "en-US"
)

// prop value enum
func (m *KnowledgeSearchDocument) validateLanguageCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeSearchDocumentTypeLanguageCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeSearchDocument) validateLanguageCode(formats strfmt.Registry) error {

	if err := validate.Required("languageCode", "body", m.LanguageCode); err != nil {
		return err
	}

	// value enum
	if err := m.validateLanguageCodeEnum("languageCode", "body", *m.LanguageCode); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeSearchDocument) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var knowledgeSearchDocumentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Faq"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeSearchDocumentTypeTypePropEnum = append(knowledgeSearchDocumentTypeTypePropEnum, v)
	}
}

const (

	// KnowledgeSearchDocumentTypeFaq captures enum value "Faq"
	KnowledgeSearchDocumentTypeFaq string = "Faq"
)

// prop value enum
func (m *KnowledgeSearchDocument) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeSearchDocumentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeSearchDocument) validateType(formats strfmt.Registry) error {

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
func (m *KnowledgeSearchDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeSearchDocument) UnmarshalBinary(b []byte) error {
	var res KnowledgeSearchDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}