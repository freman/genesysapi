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

// GroupsSearchResponse groups search response
//
// swagger:model GroupsSearchResponse
type GroupsSearchResponse struct {

	// Q64 value for the current page of results
	CurrentPage string `json:"currentPage,omitempty"`

	// Q64 value for the next page of results
	NextPage string `json:"nextPage,omitempty"`

	// The total number of pages
	// Required: true
	PageCount *int32 `json:"pageCount"`

	// The current page number
	// Required: true
	PageNumber *int32 `json:"pageNumber"`

	// The current page size
	// Required: true
	PageSize *int32 `json:"pageSize"`

	// Q64 value for the previous page of results
	PreviousPage string `json:"previousPage,omitempty"`

	// Search results
	// Required: true
	Results []*Group `json:"results"`

	// The total number of results found
	// Required: true
	Total *int64 `json:"total"`

	// Resource types the search was performed against
	// Required: true
	Types []string `json:"types"`
}

// Validate validates this groups search response
func (m *GroupsSearchResponse) Validate(formats strfmt.Registry) error {
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

func (m *GroupsSearchResponse) validatePageCount(formats strfmt.Registry) error {

	if err := validate.Required("pageCount", "body", m.PageCount); err != nil {
		return err
	}

	return nil
}

func (m *GroupsSearchResponse) validatePageNumber(formats strfmt.Registry) error {

	if err := validate.Required("pageNumber", "body", m.PageNumber); err != nil {
		return err
	}

	return nil
}

func (m *GroupsSearchResponse) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

func (m *GroupsSearchResponse) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("results", "body", m.Results); err != nil {
		return err
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

func (m *GroupsSearchResponse) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

func (m *GroupsSearchResponse) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("types", "body", m.Types); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupsSearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsSearchResponse) UnmarshalBinary(b []byte) error {
	var res GroupsSearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}