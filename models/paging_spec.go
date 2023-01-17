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

// PagingSpec paging spec
//
// swagger:model PagingSpec
type PagingSpec struct {

	// How many pages in
	// Required: true
	PageNumber *int32 `json:"pageNumber"`

	// How many results per page
	// Required: true
	PageSize *int32 `json:"pageSize"`
}

// Validate validates this paging spec
func (m *PagingSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagingSpec) validatePageNumber(formats strfmt.Registry) error {

	if err := validate.Required("pageNumber", "body", m.PageNumber); err != nil {
		return err
	}

	return nil
}

func (m *PagingSpec) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this paging spec based on context it is used
func (m *PagingSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PagingSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagingSpec) UnmarshalBinary(b []byte) error {
	var res PagingSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
