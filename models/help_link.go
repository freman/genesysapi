// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HelpLink Link to a help or support resource
//
// swagger:model HelpLink
type HelpLink struct {

	// Description of the document or resource
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Link text of the resource
	// Read Only: true
	Title string `json:"title,omitempty"`

	// URI of the help resource
	// Read Only: true
	URI string `json:"uri,omitempty"`
}

// Validate validates this help link
func (m *HelpLink) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this help link based on the context it is used
func (m *HelpLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTitle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HelpLink) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *HelpLink) contextValidateTitle(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "title", "body", string(m.Title)); err != nil {
		return err
	}

	return nil
}

func (m *HelpLink) contextValidateURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uri", "body", string(m.URI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HelpLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HelpLink) UnmarshalBinary(b []byte) error {
	var res HelpLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
