// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DocumentContentBlock document content block
//
// swagger:model DocumentContentBlock
type DocumentContentBlock struct {

	// Image. It must contain a value if the type of the block is Image.
	// Required: true
	Image *DocumentBodyImage `json:"image"`

	// Text. It must contain a value if the type of the block is Text.
	// Required: true
	Text *DocumentText `json:"text"`

	// The type of the paragraph block.
	// Required: true
	// Enum: [Text Image]
	Type *string `json:"type"`
}

// Validate validates this document content block
func (m *DocumentContentBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
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

func (m *DocumentContentBlock) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *DocumentContentBlock) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	if m.Text != nil {
		if err := m.Text.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("text")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("text")
			}
			return err
		}
	}

	return nil
}

var documentContentBlockTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Text","Image"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentContentBlockTypeTypePropEnum = append(documentContentBlockTypeTypePropEnum, v)
	}
}

const (

	// DocumentContentBlockTypeText captures enum value "Text"
	DocumentContentBlockTypeText string = "Text"

	// DocumentContentBlockTypeImage captures enum value "Image"
	DocumentContentBlockTypeImage string = "Image"
)

// prop value enum
func (m *DocumentContentBlock) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentContentBlockTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DocumentContentBlock) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this document content block based on the context it is used
func (m *DocumentContentBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateText(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocumentContentBlock) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {
		if err := m.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *DocumentContentBlock) contextValidateText(ctx context.Context, formats strfmt.Registry) error {

	if m.Text != nil {
		if err := m.Text.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("text")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("text")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocumentContentBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentContentBlock) UnmarshalBinary(b []byte) error {
	var res DocumentContentBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
