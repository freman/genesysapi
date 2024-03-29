// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScimConfigResourceTypesListResponse Defines a response for a list of SCIM resource types.
//
// swagger:model ScimConfigResourceTypesListResponse
type ScimConfigResourceTypesListResponse struct {

	// The list of requested resources.
	// Read Only: true
	Resources []*ScimConfigResourceType `json:"Resources"`

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

// Validate validates this scim config resource types list response
func (m *ScimConfigResourceTypesListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimConfigResourceTypesListResponse) validateResources(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this scim config resource types list response based on the context it is used
func (m *ScimConfigResourceTypesListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemsPerPage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimConfigResourceTypesListResponse) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "Resources", "body", []*ScimConfigResourceType(m.Resources)); err != nil {
		return err
	}

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {
			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScimConfigResourceTypesListResponse) contextValidateItemsPerPage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "itemsPerPage", "body", int64(m.ItemsPerPage)); err != nil {
		return err
	}

	return nil
}

func (m *ScimConfigResourceTypesListResponse) contextValidateStartIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "startIndex", "body", int64(m.StartIndex)); err != nil {
		return err
	}

	return nil
}

func (m *ScimConfigResourceTypesListResponse) contextValidateTotalResults(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "totalResults", "body", int64(m.TotalResults)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimConfigResourceTypesListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimConfigResourceTypesListResponse) UnmarshalBinary(b []byte) error {
	var res ScimConfigResourceTypesListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
