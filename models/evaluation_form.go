// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EvaluationForm evaluation form
//
// swagger:model EvaluationForm
type EvaluationForm struct {

	// context Id
	ContextID string `json:"contextId,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The evaluation form name
	// Required: true
	Name *string `json:"name"`

	// published
	Published bool `json:"published"`

	// published versions
	PublishedVersions *DomainEntityListingEvaluationForm `json:"publishedVersions,omitempty"`

	// A list of question groups
	// Required: true
	QuestionGroups []*EvaluationQuestionGroup `json:"questionGroups"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this evaluation form
func (m *EvaluationForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuestionGroups(formats); err != nil {
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

func (m *EvaluationForm) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EvaluationForm) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EvaluationForm) validatePublishedVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishedVersions) { // not required
		return nil
	}

	if m.PublishedVersions != nil {
		if err := m.PublishedVersions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publishedVersions")
			}
			return err
		}
	}

	return nil
}

func (m *EvaluationForm) validateQuestionGroups(formats strfmt.Registry) error {

	if err := validate.Required("questionGroups", "body", m.QuestionGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.QuestionGroups); i++ {
		if swag.IsZero(m.QuestionGroups[i]) { // not required
			continue
		}

		if m.QuestionGroups[i] != nil {
			if err := m.QuestionGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("questionGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EvaluationForm) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationForm) UnmarshalBinary(b []byte) error {
	var res EvaluationForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
