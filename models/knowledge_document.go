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

// KnowledgeDocument knowledge document
//
// swagger:model KnowledgeDocument
type KnowledgeDocument struct {

	// Article
	Article *DocumentArticle `json:"article,omitempty"`

	// Document categories
	Categories []*KnowledgeCategory `json:"categories"`

	// Document creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Document last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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
	// Enum: [en-US en-UK en-AU de-DE]
	LanguageCode *string `json:"languageCode"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Document type
	// Required: true
	// Enum: [Faq Article]
	Type *string `json:"type"`
}

// Validate validates this knowledge document
func (m *KnowledgeDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArticle(formats); err != nil {
		res = append(res, err)
	}

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

func (m *KnowledgeDocument) validateArticle(formats strfmt.Registry) error {

	if swag.IsZero(m.Article) { // not required
		return nil
	}

	if m.Article != nil {
		if err := m.Article.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("article")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocument) validateCategories(formats strfmt.Registry) error {

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

func (m *KnowledgeDocument) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocument) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocument) validateFaq(formats strfmt.Registry) error {

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

func (m *KnowledgeDocument) validateKnowledgeBase(formats strfmt.Registry) error {

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

var knowledgeDocumentTypeLanguageCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-US","en-UK","en-AU","de-DE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeDocumentTypeLanguageCodePropEnum = append(knowledgeDocumentTypeLanguageCodePropEnum, v)
	}
}

const (

	// KnowledgeDocumentLanguageCodeEnUS captures enum value "en-US"
	KnowledgeDocumentLanguageCodeEnUS string = "en-US"

	// KnowledgeDocumentLanguageCodeEnUK captures enum value "en-UK"
	KnowledgeDocumentLanguageCodeEnUK string = "en-UK"

	// KnowledgeDocumentLanguageCodeEnAU captures enum value "en-AU"
	KnowledgeDocumentLanguageCodeEnAU string = "en-AU"

	// KnowledgeDocumentLanguageCodeDeDE captures enum value "de-DE"
	KnowledgeDocumentLanguageCodeDeDE string = "de-DE"
)

// prop value enum
func (m *KnowledgeDocument) validateLanguageCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeDocumentTypeLanguageCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeDocument) validateLanguageCode(formats strfmt.Registry) error {

	if err := validate.Required("languageCode", "body", m.LanguageCode); err != nil {
		return err
	}

	// value enum
	if err := m.validateLanguageCodeEnum("languageCode", "body", *m.LanguageCode); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocument) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var knowledgeDocumentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Faq","Article"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeDocumentTypeTypePropEnum = append(knowledgeDocumentTypeTypePropEnum, v)
	}
}

const (

	// KnowledgeDocumentTypeFaq captures enum value "Faq"
	KnowledgeDocumentTypeFaq string = "Faq"

	// KnowledgeDocumentTypeArticle captures enum value "Article"
	KnowledgeDocumentTypeArticle string = "Article"
)

// prop value enum
func (m *KnowledgeDocument) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeDocumentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeDocument) validateType(formats strfmt.Registry) error {

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
func (m *KnowledgeDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeDocument) UnmarshalBinary(b []byte) error {
	var res KnowledgeDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
