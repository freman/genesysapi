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
)

// BulkFetchRelationshipsResponse bulk fetch relationships response
//
// swagger:model BulkFetchRelationshipsResponse
type BulkFetchRelationshipsResponse struct {

	// error count
	ErrorCount int32 `json:"errorCount,omitempty"`

	// error indexes
	ErrorIndexes []int32 `json:"errorIndexes"`

	// results
	Results []*BulkResponseResultRelationshipEntity `json:"results"`
}

// Validate validates this bulk fetch relationships response
func (m *BulkFetchRelationshipsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkFetchRelationshipsResponse) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bulk fetch relationships response based on the context it is used
func (m *BulkFetchRelationshipsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkFetchRelationshipsResponse) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Results); i++ {

		if m.Results[i] != nil {
			if err := m.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkFetchRelationshipsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkFetchRelationshipsResponse) UnmarshalBinary(b []byte) error {
	var res BulkFetchRelationshipsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
