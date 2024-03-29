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

// AuditSearchResult audit search result
//
// swagger:model AuditSearchResult
type AuditSearchResult struct {

	// audit messages
	AuditMessages []*AuditMessage `json:"auditMessages"`

	// facet info
	FacetInfo []*FacetInfo `json:"facetInfo"`

	// The number of pages of results.
	PageCount int32 `json:"pageCount,omitempty"`

	// Which page was returned.
	PageNumber int32 `json:"pageNumber,omitempty"`

	// The number of results in a page.
	PageSize int32 `json:"pageSize,omitempty"`

	// The total number of results.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this audit search result
func (m *AuditSearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacetInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditSearchResult) validateAuditMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditMessages) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditMessages); i++ {
		if swag.IsZero(m.AuditMessages[i]) { // not required
			continue
		}

		if m.AuditMessages[i] != nil {
			if err := m.AuditMessages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditMessages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuditSearchResult) validateFacetInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.FacetInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.FacetInfo); i++ {
		if swag.IsZero(m.FacetInfo[i]) { // not required
			continue
		}

		if m.FacetInfo[i] != nil {
			if err := m.FacetInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("facetInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("facetInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this audit search result based on the context it is used
func (m *AuditSearchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFacetInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditSearchResult) contextValidateAuditMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditMessages); i++ {

		if m.AuditMessages[i] != nil {
			if err := m.AuditMessages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditMessages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuditSearchResult) contextValidateFacetInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FacetInfo); i++ {

		if m.FacetInfo[i] != nil {
			if err := m.FacetInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("facetInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("facetInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditSearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditSearchResult) UnmarshalBinary(b []byte) error {
	var res AuditSearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
