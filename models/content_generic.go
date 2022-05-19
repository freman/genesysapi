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

// ContentGeneric Deprecated, should use Card.
//
// swagger:model ContentGeneric
type ContentGeneric struct {

	// Actions to be taken (Deprecated).
	Actions *ContentActions `json:"actions,omitempty"`

	// An array of component objects.
	Components []*ButtonComponent `json:"components"`

	// Text to show in the description.
	Description string `json:"description,omitempty"`

	// URL of an image.
	Image string `json:"image,omitempty"`

	// Text to show in the title.
	Title string `json:"title,omitempty"`

	// URL of a video.
	Video string `json:"video,omitempty"`
}

// Validate validates this content generic
func (m *ContentGeneric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentGeneric) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {
		if err := m.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *ContentGeneric) validateComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentGeneric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentGeneric) UnmarshalBinary(b []byte) error {
	var res ContentGeneric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
