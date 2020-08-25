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

// CursorContactListing cursor contact listing
//
// swagger:model CursorContactListing
type CursorContactListing struct {

	// entities
	Entities []*ExternalContact `json:"entities"`

	// next Uri
	NextURI string `json:"nextUri,omitempty"`

	// previous Uri
	PreviousURI string `json:"previousUri,omitempty"`

	// self Uri
	SelfURI string `json:"selfUri,omitempty"`
}

// Validate validates this cursor contact listing
func (m *CursorContactListing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CursorContactListing) validateEntities(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CursorContactListing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CursorContactListing) UnmarshalBinary(b []byte) error {
	var res CursorContactListing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
