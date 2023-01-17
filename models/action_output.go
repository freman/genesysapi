// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActionOutput Output definition of Action.
//
// swagger:model ActionOutput
type ActionOutput struct {

	// JSON schema that defines the body of response when request is not successful. If the 'flatten' query parameter is omitted or false, this field will be returned. Either errorSchema or errorSchemaFlattened will be returned, not both.
	ErrorSchema *JSONSchemaDocument `json:"errorSchema,omitempty"`

	// JSON schema that defines the body of response when request is not successful. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either errorSchema or errorSchemaFlattened will be returned, not both.
	ErrorSchemaFlattened interface{} `json:"errorSchemaFlattened,omitempty"`

	// URI to retrieve error schema
	ErrorSchemaURI string `json:"errorSchemaUri,omitempty"`

	// JSON schema that defines the transformed, successful result that will be sent back to the caller. If the 'flatten' query parameter is omitted or false, this field will be returned. Either successSchema or successSchemaFlattened will be returned, not both.
	SuccessSchema *JSONSchemaDocument `json:"successSchema,omitempty"`

	// JSON schema that defines the transformed, successful result that will be sent back to the caller. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either successSchema or successSchemaFlattened will be returned, not both.
	SuccessSchemaFlattened *JSONSchemaDocument `json:"successSchemaFlattened,omitempty"`

	// URI to retrieve success schema
	SuccessSchemaURI string `json:"successSchemaUri,omitempty"`
}

// Validate validates this action output
func (m *ActionOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccessSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccessSchemaFlattened(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionOutput) validateErrorSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorSchema) { // not required
		return nil
	}

	if m.ErrorSchema != nil {
		if err := m.ErrorSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorSchema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorSchema")
			}
			return err
		}
	}

	return nil
}

func (m *ActionOutput) validateSuccessSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.SuccessSchema) { // not required
		return nil
	}

	if m.SuccessSchema != nil {
		if err := m.SuccessSchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("successSchema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("successSchema")
			}
			return err
		}
	}

	return nil
}

func (m *ActionOutput) validateSuccessSchemaFlattened(formats strfmt.Registry) error {
	if swag.IsZero(m.SuccessSchemaFlattened) { // not required
		return nil
	}

	if m.SuccessSchemaFlattened != nil {
		if err := m.SuccessSchemaFlattened.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("successSchemaFlattened")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("successSchemaFlattened")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this action output based on the context it is used
func (m *ActionOutput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrorSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccessSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccessSchemaFlattened(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionOutput) contextValidateErrorSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorSchema != nil {
		if err := m.ErrorSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorSchema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorSchema")
			}
			return err
		}
	}

	return nil
}

func (m *ActionOutput) contextValidateSuccessSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.SuccessSchema != nil {
		if err := m.SuccessSchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("successSchema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("successSchema")
			}
			return err
		}
	}

	return nil
}

func (m *ActionOutput) contextValidateSuccessSchemaFlattened(ctx context.Context, formats strfmt.Registry) error {

	if m.SuccessSchemaFlattened != nil {
		if err := m.SuccessSchemaFlattened.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("successSchemaFlattened")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("successSchemaFlattened")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionOutput) UnmarshalBinary(b []byte) error {
	var res ActionOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
