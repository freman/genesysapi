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

// KnowledgeDocumentSearch knowledge document search
//
// swagger:model KnowledgeDocumentSearch
type KnowledgeDocumentSearch struct {

	// Number of pages returned in the result calculated according to the pageSize and the total
	// Read Only: true
	PageCount int32 `json:"pageCount,omitempty"`

	// Page number of the returned results.
	PageNumber int32 `json:"pageNumber,omitempty"`

	// Page size of the returned results.
	PageSize int32 `json:"pageSize,omitempty"`

	// Query to search content in the knowledge base. Maximum of 30 records per query can be fetched.
	// Required: true
	// Max Length: 2147483647
	// Min Length: 3
	Query *string `json:"query"`

	// Documents matching the search query.
	// Read Only: true
	Results []*KnowledgeDocumentSearchResult `json:"results"`

	// The globally unique identifier for the search.
	// Read Only: true
	SearchID string `json:"searchId,omitempty"`

	// The total number of documents matching the query.
	// Read Only: true
	Total int32 `json:"total,omitempty"`
}

// Validate validates this knowledge document search
func (m *KnowledgeDocumentSearch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeDocumentSearch) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	if err := validate.MinLength("query", "body", string(*m.Query), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("query", "body", string(*m.Query), 2147483647); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocumentSearch) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeDocumentSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeDocumentSearch) UnmarshalBinary(b []byte) error {
	var res KnowledgeDocumentSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
