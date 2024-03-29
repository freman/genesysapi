// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResponseAssetSearchRequest response asset search request
//
// swagger:model ResponseAssetSearchRequest
type ResponseAssetSearchRequest struct {

	// The page of resources you want to retrieve
	PageNumber int32 `json:"pageNumber,omitempty"`

	// The number of results per page. Default: 25, Maximum: 100.
	PageSize int32 `json:"pageSize,omitempty"`

	// Filter the query results.
	// Required: true
	Query []*ResponseAssetFilter `json:"query"`

	// The field in the resource that you want to sort the results by
	SortBy string `json:"sortBy,omitempty"`

	// The sort order for results
	// Enum: [ASC DESC]
	SortOrder string `json:"sortOrder,omitempty"`
}

// Validate validates this response asset search request
func (m *ResponseAssetSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
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

func (m *ResponseAssetSearchRequest) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	for i := 0; i < len(m.Query); i++ {
		if swag.IsZero(m.Query[i]) { // not required
			continue
		}

		if m.Query[i] != nil {
			if err := m.Query[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("query" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("query" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var responseAssetSearchRequestTypeSortOrderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ASC","DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responseAssetSearchRequestTypeSortOrderPropEnum = append(responseAssetSearchRequestTypeSortOrderPropEnum, v)
	}
}

const (

	// ResponseAssetSearchRequestSortOrderASC captures enum value "ASC"
	ResponseAssetSearchRequestSortOrderASC string = "ASC"

	// ResponseAssetSearchRequestSortOrderDESC captures enum value "DESC"
	ResponseAssetSearchRequestSortOrderDESC string = "DESC"
)

// prop value enum
func (m *ResponseAssetSearchRequest) validateSortOrderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responseAssetSearchRequestTypeSortOrderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponseAssetSearchRequest) validateSortOrder(formats strfmt.Registry) error {
	if swag.IsZero(m.SortOrder) { // not required
		return nil
	}

	// value enum
	if err := m.validateSortOrderEnum("sortOrder", "body", m.SortOrder); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this response asset search request based on the context it is used
func (m *ResponseAssetSearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseAssetSearchRequest) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Query); i++ {

		if m.Query[i] != nil {
			if err := m.Query[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("query" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("query" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseAssetSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseAssetSearchRequest) UnmarshalBinary(b []byte) error {
	var res ResponseAssetSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
