// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QueryResults query results
//
// swagger:model QueryResults
type QueryResults struct {

	// facet info
	FacetInfo *QueryFacetInfo `json:"facetInfo,omitempty"`

	// results
	Results *DomainEntityListingQueryResult `json:"results,omitempty"`
}

// Validate validates this query results
func (m *QueryResults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacetInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResults) validateFacetInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.FacetInfo) { // not required
		return nil
	}

	if m.FacetInfo != nil {
		if err := m.FacetInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facetInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("facetInfo")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResults) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(m.Results) { // not required
		return nil
	}

	if m.Results != nil {
		if err := m.Results.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this query results based on the context it is used
func (m *QueryResults) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFacetInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResults) contextValidateFacetInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.FacetInfo != nil {
		if err := m.FacetInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facetInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("facetInfo")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResults) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	if m.Results != nil {
		if err := m.Results.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResults) UnmarshalBinary(b []byte) error {
	var res QueryResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
