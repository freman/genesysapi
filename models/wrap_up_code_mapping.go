// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WrapUpCodeMapping wrap up code mapping
//
// swagger:model WrapUpCodeMapping
type WrapUpCodeMapping struct {

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The default set of wrap-up flags. These will be used if there is no entry for a given wrap-up code in the mapping.
	// Unique: true
	DefaultSet []string `json:"defaultSet"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// A map from wrap-up code identifiers to a set of wrap-up flags.
	Mapping map[string][]string `json:"mapping,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this wrap up code mapping
func (m *WrapUpCodeMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapping(formats); err != nil {
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

func (m *WrapUpCodeMapping) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WrapUpCodeMapping) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

var wrapUpCodeMappingDefaultSetItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONTACT_UNCALLABLE","NUMBER_UNCALLABLE","RIGHT_PARTY_CONTACT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wrapUpCodeMappingDefaultSetItemsEnum = append(wrapUpCodeMappingDefaultSetItemsEnum, v)
	}
}

func (m *WrapUpCodeMapping) validateDefaultSetItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wrapUpCodeMappingDefaultSetItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WrapUpCodeMapping) validateDefaultSet(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultSet) { // not required
		return nil
	}

	if err := validate.UniqueItems("defaultSet", "body", m.DefaultSet); err != nil {
		return err
	}

	for i := 0; i < len(m.DefaultSet); i++ {

		// value enum
		if err := m.validateDefaultSetItemsEnum("defaultSet"+"."+strconv.Itoa(i), "body", m.DefaultSet[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *WrapUpCodeMapping) validateMapping(formats strfmt.Registry) error {
	if swag.IsZero(m.Mapping) { // not required
		return nil
	}

	for k := range m.Mapping {

		if err := validate.UniqueItems("mapping"+"."+k, "body", m.Mapping[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Mapping[k]); i++ {

			// value enum
			if err := m.validateMappingItemsEnum("mapping"+"."+k+"."+strconv.Itoa(i), "body", m.Mapping[k][i]); err != nil {
				return err
			}

		}

	}

	return nil
}

func (m *WrapUpCodeMapping) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wrap up code mapping based on the context it is used
func (m *WrapUpCodeMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WrapUpCodeMapping) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *WrapUpCodeMapping) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *WrapUpCodeMapping) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WrapUpCodeMapping) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WrapUpCodeMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WrapUpCodeMapping) UnmarshalBinary(b []byte) error {
	var res WrapUpCodeMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
