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

// KnowledgeSearchRequest knowledge search request
//
// swagger:model KnowledgeSearchRequest
type KnowledgeSearchRequest struct {

	// Document type to be used while searching
	// Required: true
	// Enum: [Faq]
	DocumentType *string `json:"documentType"`

	// query search for specific languageCode
	// Required: true
	LanguageCode *string `json:"languageCode"`

	// Page number of the returned results
	// Required: true
	PageNumber *int32 `json:"pageNumber"`

	// Page size of the returned results
	// Required: true
	PageSize *int32 `json:"pageSize"`

	// Input query to search content in the knowledge base
	// Required: true
	Query *string `json:"query"`

	// If true the search query will be executed on draft documents, else it will be on active documents
	SearchOnDraftDocuments bool `json:"searchOnDraftDocuments,omitempty"`
}

// Validate validates this knowledge search request
func (m *KnowledgeSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var knowledgeSearchRequestTypeDocumentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Faq"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeSearchRequestTypeDocumentTypePropEnum = append(knowledgeSearchRequestTypeDocumentTypePropEnum, v)
	}
}

const (

	// KnowledgeSearchRequestDocumentTypeFaq captures enum value "Faq"
	KnowledgeSearchRequestDocumentTypeFaq string = "Faq"
)

// prop value enum
func (m *KnowledgeSearchRequest) validateDocumentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeSearchRequestTypeDocumentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeSearchRequest) validateDocumentType(formats strfmt.Registry) error {

	if err := validate.Required("documentType", "body", m.DocumentType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDocumentTypeEnum("documentType", "body", *m.DocumentType); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeSearchRequest) validateLanguageCode(formats strfmt.Registry) error {

	if err := validate.Required("languageCode", "body", m.LanguageCode); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeSearchRequest) validatePageNumber(formats strfmt.Registry) error {

	if err := validate.Required("pageNumber", "body", m.PageNumber); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeSearchRequest) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeSearchRequest) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeSearchRequest) UnmarshalBinary(b []byte) error {
	var res KnowledgeSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
