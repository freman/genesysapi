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

// PolicyEntityListing policy entity listing
//
// swagger:model PolicyEntityListing
type PolicyEntityListing struct {

	// entities
	Entities []*Policy `json:"entities"`

	// first Uri
	// Format: uri
	FirstURI strfmt.URI `json:"firstUri,omitempty"`

	// last Uri
	// Format: uri
	LastURI strfmt.URI `json:"lastUri,omitempty"`

	// next Uri
	// Format: uri
	NextURI strfmt.URI `json:"nextUri,omitempty"`

	// page count
	PageCount int32 `json:"pageCount,omitempty"`

	// page number
	PageNumber int32 `json:"pageNumber,omitempty"`

	// page size
	PageSize int32 `json:"pageSize,omitempty"`

	// previous Uri
	// Format: uri
	PreviousURI strfmt.URI `json:"previousUri,omitempty"`

	// self Uri
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this policy entity listing
func (m *PolicyEntityListing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyEntityListing) validateEntities(formats strfmt.Registry) error {

	if swag.IsZero(m.Entities) { // not required
		return nil
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyEntityListing) validateFirstURI(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstURI) { // not required
		return nil
	}

	if err := validate.FormatOf("firstUri", "body", "uri", m.FirstURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PolicyEntityListing) validateLastURI(formats strfmt.Registry) error {

	if swag.IsZero(m.LastURI) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUri", "body", "uri", m.LastURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PolicyEntityListing) validateNextURI(formats strfmt.Registry) error {

	if swag.IsZero(m.NextURI) { // not required
		return nil
	}

	if err := validate.FormatOf("nextUri", "body", "uri", m.NextURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PolicyEntityListing) validatePreviousURI(formats strfmt.Registry) error {

	if swag.IsZero(m.PreviousURI) { // not required
		return nil
	}

	if err := validate.FormatOf("previousUri", "body", "uri", m.PreviousURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PolicyEntityListing) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyEntityListing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyEntityListing) UnmarshalBinary(b []byte) error {
	var res PolicyEntityListing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
