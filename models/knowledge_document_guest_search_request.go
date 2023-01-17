// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KnowledgeDocumentGuestSearchRequest knowledge document guest search request
//
// swagger:model KnowledgeDocumentGuestSearchRequest
type KnowledgeDocumentGuestSearchRequest struct {

	// The app where the session is started.
	// Read Only: true
	App *KnowledgeGuestSessionApp `json:"app,omitempty"`

	// Indicates whether the search results would also include draft documents.
	IncludeDraftDocuments bool `json:"includeDraftDocuments"`

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

	// The type of the query that initiates the search.
	// Enum: [AutoSearch ManualSearch Suggestion]
	QueryType string `json:"queryType,omitempty"`

	// The globally unique identifier for the search.
	// Read Only: true
	SearchID string `json:"searchId,omitempty"`

	// Session ID of the search.
	// Read Only: true
	SessionID string `json:"sessionId,omitempty"`

	// The total number of documents matching the query.
	// Read Only: true
	Total int32 `json:"total,omitempty"`
}

// Validate validates this knowledge document guest search request
func (m *KnowledgeDocumentGuestSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeDocumentGuestSearchRequest) validateApp(formats strfmt.Registry) error {
	if swag.IsZero(m.App) { // not required
		return nil
	}

	if m.App != nil {
		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentGuestSearchRequest) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	if err := validate.MinLength("query", "body", *m.Query, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("query", "body", *m.Query, 2147483647); err != nil {
		return err
	}

	return nil
}

var knowledgeDocumentGuestSearchRequestTypeQueryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AutoSearch","ManualSearch","Suggestion"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeDocumentGuestSearchRequestTypeQueryTypePropEnum = append(knowledgeDocumentGuestSearchRequestTypeQueryTypePropEnum, v)
	}
}

const (

	// KnowledgeDocumentGuestSearchRequestQueryTypeAutoSearch captures enum value "AutoSearch"
	KnowledgeDocumentGuestSearchRequestQueryTypeAutoSearch string = "AutoSearch"

	// KnowledgeDocumentGuestSearchRequestQueryTypeManualSearch captures enum value "ManualSearch"
	KnowledgeDocumentGuestSearchRequestQueryTypeManualSearch string = "ManualSearch"

	// KnowledgeDocumentGuestSearchRequestQueryTypeSuggestion captures enum value "Suggestion"
	KnowledgeDocumentGuestSearchRequestQueryTypeSuggestion string = "Suggestion"
)

// prop value enum
func (m *KnowledgeDocumentGuestSearchRequest) validateQueryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeDocumentGuestSearchRequestTypeQueryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeDocumentGuestSearchRequest) validateQueryType(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateQueryTypeEnum("queryType", "body", m.QueryType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this knowledge document guest search request based on the context it is used
func (m *KnowledgeDocumentGuestSearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePageCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearchID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeDocumentGuestSearchRequest) contextValidateApp(ctx context.Context, formats strfmt.Registry) error {

	if m.App != nil {
		if err := m.App.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentGuestSearchRequest) contextValidatePageCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "pageCount", "body", int32(m.PageCount)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocumentGuestSearchRequest) contextValidateSearchID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "searchId", "body", string(m.SearchID)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocumentGuestSearchRequest) contextValidateSessionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sessionId", "body", string(m.SessionID)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocumentGuestSearchRequest) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total", "body", int32(m.Total)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeDocumentGuestSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeDocumentGuestSearchRequest) UnmarshalBinary(b []byte) error {
	var res KnowledgeDocumentGuestSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
