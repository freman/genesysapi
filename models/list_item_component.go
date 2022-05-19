// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListItemComponent An entry in a List template.
//
// swagger:model ListItemComponent
type ListItemComponent struct {

	// The list item actions (Deprecated).
	Actions *ContentActions `json:"actions,omitempty"`

	// Text to show in the list item description.
	Description string `json:"description,omitempty"`

	// An ID assigned to this list item.
	ID string `json:"id,omitempty"`

	// URL of an image.
	Image string `json:"image,omitempty"`

	// An ID of the rich message instance.
	Rmid string `json:"rmid,omitempty"`

	// The main headline of the list item.
	Title string `json:"title,omitempty"`

	// The type of list item to render.
	// Enum: [ListItem ListItemBig]
	Type string `json:"type,omitempty"`
}

// Validate validates this list item component
func (m *ListItemComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListItemComponent) validateActions(formats strfmt.Registry) error {

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

var listItemComponentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ListItem","ListItemBig"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listItemComponentTypeTypePropEnum = append(listItemComponentTypeTypePropEnum, v)
	}
}

const (

	// ListItemComponentTypeListItem captures enum value "ListItem"
	ListItemComponentTypeListItem string = "ListItem"

	// ListItemComponentTypeListItemBig captures enum value "ListItemBig"
	ListItemComponentTypeListItemBig string = "ListItemBig"
)

// prop value enum
func (m *ListItemComponent) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listItemComponentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ListItemComponent) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListItemComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListItemComponent) UnmarshalBinary(b []byte) error {
	var res ListItemComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
