// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JSONSearchResponse Json search response
//
// swagger:model JsonSearchResponse
type JSONSearchResponse struct {

	// aggregations
	Aggregations ArrayNode `json:"aggregations,omitempty"`

	// The total number of pages
	// Required: true
	PageCount *int32 `json:"pageCount"`

	// The current page number
	// Required: true
	PageNumber *int32 `json:"pageNumber"`

	// The current page size
	// Required: true
	PageSize *int32 `json:"pageSize"`

	// Search results
	// Required: true
	Results ArrayNode `json:"results"`

	// The total number of results found
	// Required: true
	Total *int64 `json:"total"`

	// Resource types the search was performed against
	// Required: true
	Types []string `json:"types"`
}

// Validate validates this Json search response
func (m *JSONSearchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONSearchResponse) validatePageCount(formats strfmt.Registry) error {

	if err := validate.Required("pageCount", "body", m.PageCount); err != nil {
		return err
	}

	return nil
}

func (m *JSONSearchResponse) validatePageNumber(formats strfmt.Registry) error {

	if err := validate.Required("pageNumber", "body", m.PageNumber); err != nil {
		return err
	}

	return nil
}

func (m *JSONSearchResponse) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

func (m *JSONSearchResponse) validateResults(formats strfmt.Registry) error {

	if m.Results == nil {
		return errors.Required("results", "body", nil)
	}

	return nil
}

func (m *JSONSearchResponse) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

func (m *JSONSearchResponse) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("types", "body", m.Types); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Json search response based on context it is used
func (m *JSONSearchResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JSONSearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONSearchResponse) UnmarshalBinary(b []byte) error {
	var res JSONSearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
