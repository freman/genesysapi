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

// KnowledgeDocumentRequest knowledge document request
//
// swagger:model KnowledgeDocumentRequest
type KnowledgeDocumentRequest struct {

	// Article details
	Article *DocumentArticle `json:"article,omitempty"`

	// Document categories
	Categories []*DocumentCategoryInput `json:"categories"`

	// External Url to the document
	ExternalURL string `json:"externalUrl,omitempty"`

	// Faq document details
	Faq *DocumentFaq `json:"faq,omitempty"`

	// Document type according to assigned template
	// Required: true
	// Enum: [Faq Article]
	Type *string `json:"type"`
}

// Validate validates this knowledge document request
func (m *KnowledgeDocumentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArticle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaq(formats); err != nil {
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

func (m *KnowledgeDocumentRequest) validateArticle(formats strfmt.Registry) error {

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

func (m *KnowledgeDocumentRequest) validateCategories(formats strfmt.Registry) error {

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

func (m *KnowledgeDocumentRequest) validateFaq(formats strfmt.Registry) error {

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

var knowledgeDocumentRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Faq","Article"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeDocumentRequestTypeTypePropEnum = append(knowledgeDocumentRequestTypeTypePropEnum, v)
	}
}

const (

	// KnowledgeDocumentRequestTypeFaq captures enum value "Faq"
	KnowledgeDocumentRequestTypeFaq string = "Faq"

	// KnowledgeDocumentRequestTypeArticle captures enum value "Article"
	KnowledgeDocumentRequestTypeArticle string = "Article"
)

// prop value enum
func (m *KnowledgeDocumentRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeDocumentRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeDocumentRequest) validateType(formats strfmt.Registry) error {

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
func (m *KnowledgeDocumentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeDocumentRequest) UnmarshalBinary(b []byte) error {
	var res KnowledgeDocumentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
