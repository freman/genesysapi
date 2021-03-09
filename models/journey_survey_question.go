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

// JourneySurveyQuestion journey survey question
//
// swagger:model JourneySurveyQuestion
type JourneySurveyQuestion struct {

	// Choices available to user.
	Choices []string `json:"choices"`

	// The customer property that the answer maps to.
	// Enum: [givenName familyName email phone gender companyName]
	CustomerProperty string `json:"customerProperty,omitempty"`

	// Whether answering this question is mandatory.
	IsMandatory bool `json:"isMandatory"`

	// Label of question.
	// Required: true
	Label *string `json:"label"`

	// Type of survey question.
	// Enum: [text hidden select checkbox textarea]
	Type string `json:"type,omitempty"`
}

// Validate validates this journey survey question
func (m *JourneySurveyQuestion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerProperty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
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

var journeySurveyQuestionTypeCustomerPropertyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["givenName","familyName","email","phone","gender","companyName"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journeySurveyQuestionTypeCustomerPropertyPropEnum = append(journeySurveyQuestionTypeCustomerPropertyPropEnum, v)
	}
}

const (

	// JourneySurveyQuestionCustomerPropertyGivenName captures enum value "givenName"
	JourneySurveyQuestionCustomerPropertyGivenName string = "givenName"

	// JourneySurveyQuestionCustomerPropertyFamilyName captures enum value "familyName"
	JourneySurveyQuestionCustomerPropertyFamilyName string = "familyName"

	// JourneySurveyQuestionCustomerPropertyEmail captures enum value "email"
	JourneySurveyQuestionCustomerPropertyEmail string = "email"

	// JourneySurveyQuestionCustomerPropertyPhone captures enum value "phone"
	JourneySurveyQuestionCustomerPropertyPhone string = "phone"

	// JourneySurveyQuestionCustomerPropertyGender captures enum value "gender"
	JourneySurveyQuestionCustomerPropertyGender string = "gender"

	// JourneySurveyQuestionCustomerPropertyCompanyName captures enum value "companyName"
	JourneySurveyQuestionCustomerPropertyCompanyName string = "companyName"
)

// prop value enum
func (m *JourneySurveyQuestion) validateCustomerPropertyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journeySurveyQuestionTypeCustomerPropertyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JourneySurveyQuestion) validateCustomerProperty(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerProperty) { // not required
		return nil
	}

	// value enum
	if err := m.validateCustomerPropertyEnum("customerProperty", "body", m.CustomerProperty); err != nil {
		return err
	}

	return nil
}

func (m *JourneySurveyQuestion) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var journeySurveyQuestionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text","hidden","select","checkbox","textarea"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		journeySurveyQuestionTypeTypePropEnum = append(journeySurveyQuestionTypeTypePropEnum, v)
	}
}

const (

	// JourneySurveyQuestionTypeText captures enum value "text"
	JourneySurveyQuestionTypeText string = "text"

	// JourneySurveyQuestionTypeHidden captures enum value "hidden"
	JourneySurveyQuestionTypeHidden string = "hidden"

	// JourneySurveyQuestionTypeSelect captures enum value "select"
	JourneySurveyQuestionTypeSelect string = "select"

	// JourneySurveyQuestionTypeCheckbox captures enum value "checkbox"
	JourneySurveyQuestionTypeCheckbox string = "checkbox"

	// JourneySurveyQuestionTypeTextarea captures enum value "textarea"
	JourneySurveyQuestionTypeTextarea string = "textarea"
)

// prop value enum
func (m *JourneySurveyQuestion) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, journeySurveyQuestionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JourneySurveyQuestion) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JourneySurveyQuestion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JourneySurveyQuestion) UnmarshalBinary(b []byte) error {
	var res JourneySurveyQuestion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
