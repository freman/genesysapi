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

// ProgramsEntityListing programs entity listing
//
// swagger:model ProgramsEntityListing
type ProgramsEntityListing struct {

	// entities
	Entities []*ListedProgram `json:"entities"`

	// next Uri
	// Format: uri
	NextURI strfmt.URI `json:"nextUri,omitempty"`

	// page count
	PageCount int32 `json:"pageCount,omitempty"`

	// page size
	PageSize int32 `json:"pageSize,omitempty"`

	// self Uri
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this programs entity listing
func (m *ProgramsEntityListing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextURI(formats); err != nil {
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

func (m *ProgramsEntityListing) validateEntities(formats strfmt.Registry) error {

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

func (m *ProgramsEntityListing) validateNextURI(formats strfmt.Registry) error {

	if swag.IsZero(m.NextURI) { // not required
		return nil
	}

	if err := validate.FormatOf("nextUri", "body", "uri", m.NextURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProgramsEntityListing) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProgramsEntityListing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgramsEntityListing) UnmarshalBinary(b []byte) error {
	var res ProgramsEntityListing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
