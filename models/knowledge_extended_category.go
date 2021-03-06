// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KnowledgeExtendedCategory knowledge extended category
//
// swagger:model KnowledgeExtendedCategory
type KnowledgeExtendedCategory struct {

	// Category children
	// Read Only: true
	Children []*KnowledgeCategory `json:"children"`

	// Category creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Category last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Category description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Knowledge base which category does belong to
	// Read Only: true
	KnowledgeBase *KnowledgeBase `json:"knowledgeBase,omitempty"`

	// Actual language of the category
	// Read Only: true
	// Enum: [en-US de-DE]
	LanguageCode string `json:"languageCode,omitempty"`

	// Category name
	// Required: true
	Name *string `json:"name"`

	// Category parent
	// Read Only: true
	Parent *KnowledgeCategory `json:"parent,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this knowledge extended category
func (m *KnowledgeExtendedCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKnowledgeBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
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

func (m *KnowledgeExtendedCategory) validateChildren(formats strfmt.Registry) error {

	if swag.IsZero(m.Children) { // not required
		return nil
	}

	for i := 0; i < len(m.Children); i++ {
		if swag.IsZero(m.Children[i]) { // not required
			continue
		}

		if m.Children[i] != nil {
			if err := m.Children[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KnowledgeExtendedCategory) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeExtendedCategory) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeExtendedCategory) validateKnowledgeBase(formats strfmt.Registry) error {

	if swag.IsZero(m.KnowledgeBase) { // not required
		return nil
	}

	if m.KnowledgeBase != nil {
		if err := m.KnowledgeBase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("knowledgeBase")
			}
			return err
		}
	}

	return nil
}

var knowledgeExtendedCategoryTypeLanguageCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-US","de-DE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeExtendedCategoryTypeLanguageCodePropEnum = append(knowledgeExtendedCategoryTypeLanguageCodePropEnum, v)
	}
}

const (

	// KnowledgeExtendedCategoryLanguageCodeEnUS captures enum value "en-US"
	KnowledgeExtendedCategoryLanguageCodeEnUS string = "en-US"

	// KnowledgeExtendedCategoryLanguageCodeDeDE captures enum value "de-DE"
	KnowledgeExtendedCategoryLanguageCodeDeDE string = "de-DE"
)

// prop value enum
func (m *KnowledgeExtendedCategory) validateLanguageCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeExtendedCategoryTypeLanguageCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeExtendedCategory) validateLanguageCode(formats strfmt.Registry) error {

	if swag.IsZero(m.LanguageCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateLanguageCodeEnum("languageCode", "body", m.LanguageCode); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeExtendedCategory) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeExtendedCategory) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeExtendedCategory) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeExtendedCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeExtendedCategory) UnmarshalBinary(b []byte) error {
	var res KnowledgeExtendedCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
