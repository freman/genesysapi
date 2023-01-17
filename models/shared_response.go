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

// SharedResponse shared response
//
// swagger:model SharedResponse
type SharedResponse struct {

	// document
	Document *Document `json:"document,omitempty"`

	// download Uri
	// Format: uri
	DownloadURI strfmt.URI `json:"downloadUri,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// share
	Share *Share `json:"share,omitempty"`

	// view Uri
	// Format: uri
	ViewURI strfmt.URI `json:"viewUri,omitempty"`
}

// Validate validates this shared response
func (m *SharedResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShare(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SharedResponse) validateDocument(formats strfmt.Registry) error {
	if swag.IsZero(m.Document) { // not required
		return nil
	}

	if m.Document != nil {
		if err := m.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

func (m *SharedResponse) validateDownloadURI(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadURI) { // not required
		return nil
	}

	if err := validate.FormatOf("downloadUri", "body", "uri", m.DownloadURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SharedResponse) validateShare(formats strfmt.Registry) error {
	if swag.IsZero(m.Share) { // not required
		return nil
	}

	if m.Share != nil {
		if err := m.Share.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("share")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("share")
			}
			return err
		}
	}

	return nil
}

func (m *SharedResponse) validateViewURI(formats strfmt.Registry) error {
	if swag.IsZero(m.ViewURI) { // not required
		return nil
	}

	if err := validate.FormatOf("viewUri", "body", "uri", m.ViewURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this shared response based on the context it is used
func (m *SharedResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShare(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SharedResponse) contextValidateDocument(ctx context.Context, formats strfmt.Registry) error {

	if m.Document != nil {
		if err := m.Document.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

func (m *SharedResponse) contextValidateShare(ctx context.Context, formats strfmt.Registry) error {

	if m.Share != nil {
		if err := m.Share.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("share")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("share")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SharedResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharedResponse) UnmarshalBinary(b []byte) error {
	var res SharedResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
