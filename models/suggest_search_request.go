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

// SuggestSearchRequest suggest search request
//
// swagger:model SuggestSearchRequest
type SuggestSearchRequest struct {

	// Provides more details about a specified resource
	Expand []string `json:"expand"`

	// Suggest query
	// Required: true
	Query []*SuggestSearchCriteria `json:"query"`

	// Resource domain type to search
	// Required: true
	Types []string `json:"types"`
}

// Validate validates this suggest search request
func (m *SuggestSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
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

func (m *SuggestSearchRequest) validateQuery(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *SuggestSearchRequest) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("types", "body", m.Types); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SuggestSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuggestSearchRequest) UnmarshalBinary(b []byte) error {
	var res SuggestSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
