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

// TeamMemberEntityListing team member entity listing
//
// swagger:model TeamMemberEntityListing
type TeamMemberEntityListing struct {

	// entities
	Entities []*UserReferenceWithName `json:"entities"`

	// next Uri
	// Format: uri
	NextURI strfmt.URI `json:"nextUri,omitempty"`

	// previous Uri
	// Format: uri
	PreviousURI strfmt.URI `json:"previousUri,omitempty"`

	// self Uri
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this team member entity listing
func (m *TeamMemberEntityListing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
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

func (m *TeamMemberEntityListing) validateEntities(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TeamMemberEntityListing) validateNextURI(formats strfmt.Registry) error {
	if swag.IsZero(m.NextURI) { // not required
		return nil
	}

	if err := validate.FormatOf("nextUri", "body", "uri", m.NextURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TeamMemberEntityListing) validatePreviousURI(formats strfmt.Registry) error {
	if swag.IsZero(m.PreviousURI) { // not required
		return nil
	}

	if err := validate.FormatOf("previousUri", "body", "uri", m.PreviousURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TeamMemberEntityListing) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this team member entity listing based on the context it is used
func (m *TeamMemberEntityListing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamMemberEntityListing) contextValidateEntities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entities); i++ {

		if m.Entities[i] != nil {
			if err := m.Entities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamMemberEntityListing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamMemberEntityListing) UnmarshalBinary(b []byte) error {
	var res TeamMemberEntityListing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
