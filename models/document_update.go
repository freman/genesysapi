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

// DocumentUpdate document update
//
// swagger:model DocumentUpdate
type DocumentUpdate struct {

	// add tag ids
	AddTagIds []string `json:"addTagIds"`

	// add tags
	AddTags []string `json:"addTags"`

	// change number
	ChangeNumber int32 `json:"changeNumber,omitempty"`

	// The name of the document
	// Required: true
	Name *string `json:"name"`

	// read
	Read bool `json:"read"`

	// remove attributes
	RemoveAttributes []string `json:"removeAttributes"`

	// remove tag ids
	RemoveTagIds []string `json:"removeTagIds"`

	// remove tags
	RemoveTags []string `json:"removeTags"`

	// update attributes
	UpdateAttributes []*DocumentAttribute `json:"updateAttributes"`
}

// Validate validates this document update
func (m *DocumentUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocumentUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DocumentUpdate) validateUpdateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateAttributes); i++ {
		if swag.IsZero(m.UpdateAttributes[i]) { // not required
			continue
		}

		if m.UpdateAttributes[i] != nil {
			if err := m.UpdateAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateAttributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocumentUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentUpdate) UnmarshalBinary(b []byte) error {
	var res DocumentUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
