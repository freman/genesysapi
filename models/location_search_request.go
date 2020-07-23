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

// LocationSearchRequest location search request
//
// swagger:model LocationSearchRequest
type LocationSearchRequest struct {

	// Provides more details about a specified resource
	Expand []string `json:"expand"`

	// The page of resources you want to retrieve
	PageNumber int32 `json:"pageNumber,omitempty"`

	// The number of results per page
	PageSize int32 `json:"pageSize,omitempty"`

	// query
	Query []*LocationSearchCriteria `json:"query"`

	// Multi-value sort order, list of multiple sort values
	Sort []*SearchSort `json:"sort"`

	// The field in the resource that you want to sort the results by
	SortBy string `json:"sortBy,omitempty"`

	// The sort order for results
	// Enum: [ASC DESC SCORE]
	SortOrder string `json:"sortOrder,omitempty"`
}

// Validate validates this location search request
func (m *LocationSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationSearchRequest) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.Query) { // not required
		return nil
	}

	for i := 0; i < len(m.Query); i++ {
		if swag.IsZero(m.Query[i]) { // not required
			continue
		}

		if m.Query[i] != nil {
			if err := m.Query[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("query" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LocationSearchRequest) validateSort(formats strfmt.Registry) error {

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

var locationSearchRequestTypeSortOrderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ASC","DESC","SCORE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		locationSearchRequestTypeSortOrderPropEnum = append(locationSearchRequestTypeSortOrderPropEnum, v)
	}
}

const (

	// LocationSearchRequestSortOrderASC captures enum value "ASC"
	LocationSearchRequestSortOrderASC string = "ASC"

	// LocationSearchRequestSortOrderDESC captures enum value "DESC"
	LocationSearchRequestSortOrderDESC string = "DESC"

	// LocationSearchRequestSortOrderSCORE captures enum value "SCORE"
	LocationSearchRequestSortOrderSCORE string = "SCORE"
)

// prop value enum
func (m *LocationSearchRequest) validateSortOrderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, locationSearchRequestTypeSortOrderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LocationSearchRequest) validateSortOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.SortOrder) { // not required
		return nil
	}

	// value enum
	if err := m.validateSortOrderEnum("sortOrder", "body", m.SortOrder); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocationSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationSearchRequest) UnmarshalBinary(b []byte) error {
	var res LocationSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
