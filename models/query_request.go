// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QueryRequest query request
//
// swagger:model QueryRequest
type QueryRequest struct {

	// attribute filters
	AttributeFilters []*AttributeFilterItem `json:"attributeFilters"`

	// facet name requests
	FacetNameRequests []string `json:"facetNameRequests"`

	// filters
	Filters []*ContentFilterItem `json:"filters"`

	// include shares
	IncludeShares bool `json:"includeShares"`

	// page number
	PageNumber int32 `json:"pageNumber,omitempty"`

	// page size
	PageSize int32 `json:"pageSize,omitempty"`

	// query phrase
	QueryPhrase string `json:"queryPhrase,omitempty"`

	// sort
	Sort []*SortItem `json:"sort"`
}

// Validate validates this query request
func (m *QueryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryRequest) validateAttributeFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributeFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.AttributeFilters); i++ {
		if swag.IsZero(m.AttributeFilters[i]) { // not required
			continue
		}

		if m.AttributeFilters[i] != nil {
			if err := m.AttributeFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryRequest) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueryRequest) validateSort(formats strfmt.Registry) error {

	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	for i := 0; i < len(m.Sort); i++ {
		if swag.IsZero(m.Sort[i]) { // not required
			continue
		}

		if m.Sort[i] != nil {
			if err := m.Sort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryRequest) UnmarshalBinary(b []byte) error {
	var res QueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
