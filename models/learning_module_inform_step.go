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

// LearningModuleInformStep learning module inform step
//
// swagger:model LearningModuleInformStep
type LearningModuleInformStep struct {

	// The document type for Content type Inform step
	ContentType string `json:"contentType,omitempty"`

	// The name of the inform step or content
	Name string `json:"name,omitempty"`

	// The order of inform step in a learning module
	// Required: true
	Order *int32 `json:"order"`

	// The sharing uri for Content type inform step
	SharingURI string `json:"sharingUri,omitempty"`

	// The learning module inform step type
	// Required: true
	// Enum: [Url Content GenesysBuiltInCourse]
	Type *string `json:"type"`

	// The value for inform step
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this learning module inform step
func (m *LearningModuleInformStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningModuleInformStep) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("order", "body", m.Order); err != nil {
		return err
	}

	return nil
}

var learningModuleInformStepTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Url","Content","GenesysBuiltInCourse"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningModuleInformStepTypeTypePropEnum = append(learningModuleInformStepTypeTypePropEnum, v)
	}
}

const (

	// LearningModuleInformStepTypeURL captures enum value "Url"
	LearningModuleInformStepTypeURL string = "Url"

	// LearningModuleInformStepTypeContent captures enum value "Content"
	LearningModuleInformStepTypeContent string = "Content"

	// LearningModuleInformStepTypeGenesysBuiltInCourse captures enum value "GenesysBuiltInCourse"
	LearningModuleInformStepTypeGenesysBuiltInCourse string = "GenesysBuiltInCourse"
)

// prop value enum
func (m *LearningModuleInformStep) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningModuleInformStepTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningModuleInformStep) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *LearningModuleInformStep) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningModuleInformStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningModuleInformStep) UnmarshalBinary(b []byte) error {
	var res LearningModuleInformStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
