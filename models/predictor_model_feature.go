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

// PredictorModelFeature predictor model feature
//
// swagger:model PredictorModelFeature
type PredictorModelFeature struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The percentage of how important a feature is in the model.
	// Read Only: true
	PercentageImportance float64 `json:"percentageImportance,omitempty"`

	// The type of feature.
	// Read Only: true
	// Enum: [User Customer Other]
	Type string `json:"type,omitempty"`
}

// Validate validates this predictor model feature
func (m *PredictorModelFeature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var predictorModelFeatureTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["User","Customer","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		predictorModelFeatureTypeTypePropEnum = append(predictorModelFeatureTypeTypePropEnum, v)
	}
}

const (

	// PredictorModelFeatureTypeUser captures enum value "User"
	PredictorModelFeatureTypeUser string = "User"

	// PredictorModelFeatureTypeCustomer captures enum value "Customer"
	PredictorModelFeatureTypeCustomer string = "Customer"

	// PredictorModelFeatureTypeOther captures enum value "Other"
	PredictorModelFeatureTypeOther string = "Other"
)

// prop value enum
func (m *PredictorModelFeature) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, predictorModelFeatureTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PredictorModelFeature) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this predictor model feature based on the context it is used
func (m *PredictorModelFeature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePercentageImportance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PredictorModelFeature) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PredictorModelFeature) contextValidatePercentageImportance(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "percentageImportance", "body", float64(m.PercentageImportance)); err != nil {
		return err
	}

	return nil
}

func (m *PredictorModelFeature) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PredictorModelFeature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PredictorModelFeature) UnmarshalBinary(b []byte) error {
	var res PredictorModelFeature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}