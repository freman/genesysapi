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

// UserSearchRequest user search request
//
// swagger:model UserSearchRequest
type UserSearchRequest struct {

	// This property only applies to api/v2/user/search; when set to true add additional search criteria to filter users by: directory:user:view
	EnforcePermissions bool `json:"enforcePermissions"`

	// Provides more details about a specified resource
	Expand []string `json:"expand"`

	// Gets an integration presence for users instead of their defaults. This parameter will only be used when presence is provided as an "expand". When using this parameter the maximum number of users that can be returned is 100.
	// Enum: [MicrosoftTeams ZoomPhone EightByEight]
	IntegrationPresenceSource string `json:"integrationPresenceSource,omitempty"`

	// The page of resources you want to retrieve
	PageNumber int32 `json:"pageNumber,omitempty"`

	// The number of results per page
	PageSize int32 `json:"pageSize,omitempty"`

	// query
	Query []*UserSearchCriteria `json:"query"`

	// Multi-value sort order, list of multiple sort values
	Sort []*SearchSort `json:"sort"`

	// The field in the resource that you want to sort the results by
	SortBy string `json:"sortBy,omitempty"`

	// The sort order for results
	// Enum: [ASC DESC SCORE]
	SortOrder string `json:"sortOrder,omitempty"`
}

// Validate validates this user search request
func (m *UserSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrationPresenceSource(formats); err != nil {
		res = append(res, err)
	}

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

var userSearchRequestTypeIntegrationPresenceSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MicrosoftTeams","ZoomPhone","EightByEight"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userSearchRequestTypeIntegrationPresenceSourcePropEnum = append(userSearchRequestTypeIntegrationPresenceSourcePropEnum, v)
	}
}

const (

	// UserSearchRequestIntegrationPresenceSourceMicrosoftTeams captures enum value "MicrosoftTeams"
	UserSearchRequestIntegrationPresenceSourceMicrosoftTeams string = "MicrosoftTeams"

	// UserSearchRequestIntegrationPresenceSourceZoomPhone captures enum value "ZoomPhone"
	UserSearchRequestIntegrationPresenceSourceZoomPhone string = "ZoomPhone"

	// UserSearchRequestIntegrationPresenceSourceEightByEight captures enum value "EightByEight"
	UserSearchRequestIntegrationPresenceSourceEightByEight string = "EightByEight"
)

// prop value enum
func (m *UserSearchRequest) validateIntegrationPresenceSourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userSearchRequestTypeIntegrationPresenceSourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserSearchRequest) validateIntegrationPresenceSource(formats strfmt.Registry) error {
	if swag.IsZero(m.IntegrationPresenceSource) { // not required
		return nil
	}

	// value enum
	if err := m.validateIntegrationPresenceSourceEnum("integrationPresenceSource", "body", m.IntegrationPresenceSource); err != nil {
		return err
	}

	return nil
}

func (m *UserSearchRequest) validateQuery(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("query" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserSearchRequest) validateSort(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var userSearchRequestTypeSortOrderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ASC","DESC","SCORE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userSearchRequestTypeSortOrderPropEnum = append(userSearchRequestTypeSortOrderPropEnum, v)
	}
}

const (

	// UserSearchRequestSortOrderASC captures enum value "ASC"
	UserSearchRequestSortOrderASC string = "ASC"

	// UserSearchRequestSortOrderDESC captures enum value "DESC"
	UserSearchRequestSortOrderDESC string = "DESC"

	// UserSearchRequestSortOrderSCORE captures enum value "SCORE"
	UserSearchRequestSortOrderSCORE string = "SCORE"
)

// prop value enum
func (m *UserSearchRequest) validateSortOrderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userSearchRequestTypeSortOrderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserSearchRequest) validateSortOrder(formats strfmt.Registry) error {
	if swag.IsZero(m.SortOrder) { // not required
		return nil
	}

	// value enum
	if err := m.validateSortOrderEnum("sortOrder", "body", m.SortOrder); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user search request based on the context it is used
func (m *UserSearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSearchRequest) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserSearchRequest) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sort); i++ {

		if m.Sort[i] != nil {
			if err := m.Sort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSearchRequest) UnmarshalBinary(b []byte) error {
	var res UserSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
