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

// EvaluationSource evaluation source
//
// swagger:model EvaluationSource
type EvaluationSource struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Type of the evaluation source.
	// Enum: [Policy User Unknown]
	Type string `json:"type,omitempty"`
}

// Validate validates this evaluation source
func (m *EvaluationSource) Validate(formats strfmt.Registry) error {
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

func (m *EvaluationSource) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var evaluationSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Policy","User","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		evaluationSourceTypeTypePropEnum = append(evaluationSourceTypeTypePropEnum, v)
	}
}

const (

	// EvaluationSourceTypePolicy captures enum value "Policy"
	EvaluationSourceTypePolicy string = "Policy"

	// EvaluationSourceTypeUser captures enum value "User"
	EvaluationSourceTypeUser string = "User"

	// EvaluationSourceTypeUnknown captures enum value "Unknown"
	EvaluationSourceTypeUnknown string = "Unknown"
)

// prop value enum
func (m *EvaluationSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, evaluationSourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EvaluationSource) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this evaluation source based on the context it is used
func (m *EvaluationSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *EvaluationSource) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *EvaluationSource) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationSource) UnmarshalBinary(b []byte) error {
	var res EvaluationSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}