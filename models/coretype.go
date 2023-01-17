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

// Coretype coretype
//
// swagger:model Coretype
type Coretype struct {

	// A boolean indicating if the core type's version is the current one in use by the system
	Current bool `json:"current"`

	// The date the core type was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Specific to the "tag" core type, this is an array of strings naming the tag item fields of the core type subject to validation
	ItemValidationFields []string `json:"itemValidationFields"`

	// A structure denoting the system-imposed minimum and maximum string length for string-array based core types such as "tag" and "enum".  Forexample, the validationLimits for a schema field using a tag core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schema author on individual tags.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field's tags.
	// Example: \"validationLimits\": {\n     \"minLength\": {\"min\": 1, \"max\": 100},\n     \"maxLength\": {\"min\": 1, \"max\": 100}\n}
	ItemValidationLimits *ItemValidationLimits `json:"itemValidationLimits,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The core type's built-in schema
	Schema *Schema `json:"schema,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// An array of strings naming the fields of the core type subject to validation.  Validation constraints are specified by a schema author using the core type.
	ValidationFields []string `json:"validationFields"`

	// A structure denoting the system-imposed minimum and maximum string length (for text-based core types) or numeric values (for number-based) core types.  For example, the validationLimits for a text-based core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schemaauthor on a text field.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field.
	// Example: \"validationLimits\": {\n\"minLength\": {\"min\": 0, \"max\": 100},\n\"maxLength\": {\"min\": 1, \"max\": 100}\n}
	ValidationLimits *ValidationLimits `json:"validationLimits,omitempty"`

	// A positive integer denoting the core type's version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this coretype
func (m *Coretype) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemValidationLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Coretype) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Coretype) validateItemValidationLimits(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemValidationLimits) { // not required
		return nil
	}

	if m.ItemValidationLimits != nil {
		if err := m.ItemValidationLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemValidationLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("itemValidationLimits")
			}
			return err
		}
	}

	return nil
}

func (m *Coretype) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *Coretype) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Coretype) validateValidationLimits(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationLimits) { // not required
		return nil
	}

	if m.ValidationLimits != nil {
		if err := m.ValidationLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validationLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validationLimits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this coretype based on the context it is used
func (m *Coretype) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemValidationLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidationLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Coretype) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Coretype) contextValidateItemValidationLimits(ctx context.Context, formats strfmt.Registry) error {

	if m.ItemValidationLimits != nil {
		if err := m.ItemValidationLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemValidationLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("itemValidationLimits")
			}
			return err
		}
	}

	return nil
}

func (m *Coretype) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *Coretype) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Coretype) contextValidateValidationLimits(ctx context.Context, formats strfmt.Registry) error {

	if m.ValidationLimits != nil {
		if err := m.ValidationLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validationLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validationLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Coretype) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Coretype) UnmarshalBinary(b []byte) error {
	var res Coretype
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
