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

// ContentCard Card content object.
//
// swagger:model ContentCard
type ContentCard struct {

	// An array of action objects.
	// Required: true
	Actions []*ContentCardAction `json:"actions"`

	// The default button action.
	DefaultAction *ContentCardAction `json:"defaultAction,omitempty"`

	// Text to show in the description.
	Description string `json:"description,omitempty"`

	// URL of an image.
	Image string `json:"image,omitempty"`

	// Text to show in the title.
	// Required: true
	Title *string `json:"title"`

	// URL of a video.
	Video string `json:"video,omitempty"`
}

// Validate validates this content card
func (m *ContentCard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentCard) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentCard) validateDefaultAction(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultAction) { // not required
		return nil
	}

	if m.DefaultAction != nil {
		if err := m.DefaultAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultAction")
			}
			return err
		}
	}

	return nil
}

func (m *ContentCard) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentCard) UnmarshalBinary(b []byte) error {
	var res ContentCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}