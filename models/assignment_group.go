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

// AssignmentGroup assignment group
//
// swagger:model AssignmentGroup
type AssignmentGroup struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// self Uri
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// type
	// Enum: [TEAM]
	Type string `json:"type,omitempty"`
}

// Validate validates this assignment group
func (m *AssignmentGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *AssignmentGroup) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var assignmentGroupTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TEAM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assignmentGroupTypeTypePropEnum = append(assignmentGroupTypeTypePropEnum, v)
	}
}

const (

	// AssignmentGroupTypeTEAM captures enum value "TEAM"
	AssignmentGroupTypeTEAM string = "TEAM"
)

// prop value enum
func (m *AssignmentGroup) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assignmentGroupTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AssignmentGroup) validateType(formats strfmt.Registry) error {

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
func (m *AssignmentGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignmentGroup) UnmarshalBinary(b []byte) error {
	var res AssignmentGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
