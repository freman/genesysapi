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

// LearningAssessment learning assessment
//
// swagger:model LearningAssessment
type LearningAssessment struct {

	// Answers for the assessment
	Answers *AssessmentScoringSet `json:"answers,omitempty"`

	// The Id of the related assessment form
	// Read Only: true
	AssessmentFormID string `json:"assessmentFormId,omitempty"`

	// The Id of the assessment
	// Read Only: true
	AssessmentID string `json:"assessmentId,omitempty"`

	// The context Id of the related assessment form
	// Read Only: true
	ContextID string `json:"contextId,omitempty"`

	// Date the assessment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date the assessment was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Date the assessment was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateSubmitted strfmt.DateTime `json:"dateSubmitted,omitempty"`

	// Status of the assessment
	// Read Only: true
	// Enum: [Pending InProgress Finished]
	Status string `json:"status,omitempty"`
}

// Validate validates this learning assessment
func (m *LearningAssessment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateSubmitted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningAssessment) validateAnswers(formats strfmt.Registry) error {

	if swag.IsZero(m.Answers) { // not required
		return nil
	}

	if m.Answers != nil {
		if err := m.Answers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("answers")
			}
			return err
		}
	}

	return nil
}

func (m *LearningAssessment) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LearningAssessment) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LearningAssessment) validateDateSubmitted(formats strfmt.Registry) error {

	if swag.IsZero(m.DateSubmitted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateSubmitted", "body", "date-time", m.DateSubmitted.String(), formats); err != nil {
		return err
	}

	return nil
}

var learningAssessmentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","InProgress","Finished"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningAssessmentTypeStatusPropEnum = append(learningAssessmentTypeStatusPropEnum, v)
	}
}

const (

	// LearningAssessmentStatusPending captures enum value "Pending"
	LearningAssessmentStatusPending string = "Pending"

	// LearningAssessmentStatusInProgress captures enum value "InProgress"
	LearningAssessmentStatusInProgress string = "InProgress"

	// LearningAssessmentStatusFinished captures enum value "Finished"
	LearningAssessmentStatusFinished string = "Finished"
)

// prop value enum
func (m *LearningAssessment) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningAssessmentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningAssessment) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningAssessment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningAssessment) UnmarshalBinary(b []byte) error {
	var res LearningAssessment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
