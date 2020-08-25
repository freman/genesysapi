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

// ScimUserListResponse Defines a response for a list of SCIM users.
//
// swagger:model ScimUserListResponse
type ScimUserListResponse struct {

	// The list of requested resources. If "count" is 0, then the list will be empty.
	// Read Only: true
	Resources []*ScimV2User `json:"Resources"`

	// The number of resources returned per page.
	// Read Only: true
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// The list of supported schemas.
	Schemas []string `json:"schemas"`

	// The 1-based index of the first result returned by this request. Add this to "itemsPerPage" when requesting the next page of results.
	// Read Only: true
	StartIndex int64 `json:"startIndex,omitempty"`

	// The total number of results.
	// Read Only: true
	TotalResults int64 `json:"totalResults,omitempty"`
}

// Validate validates this scim user list response
func (m *ScimUserListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimUserListResponse) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimUserListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimUserListResponse) UnmarshalBinary(b []byte) error {
	var res ScimUserListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}