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

// LearningAssignment Learning module assignment with user information
//
// swagger:model LearningAssignment
type LearningAssignment struct {

	// The assessment associated with this assignment
	Assessment *LearningAssessment `json:"assessment,omitempty"`

	// The assessment form associated with this assignment
	AssessmentForm *AssessmentForm `json:"assessmentForm,omitempty"`

	// The user who created the assignment
	// Read Only: true
	CreatedBy *UserReference `json:"createdBy,omitempty"`

	// The date when the assignment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date when the assignment was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The recommended completion date of the assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateRecommendedForCompletion strfmt.DateTime `json:"dateRecommendedForCompletion,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// True if this assignment was created manually
	// Read Only: true
	IsManual *bool `json:"isManual"`

	// True if the assignment is overdue
	// Read Only: true
	IsOverdue *bool `json:"isOverdue"`

	// True if the assessment was passed
	// Read Only: true
	IsPassed *bool `json:"isPassed"`

	// True if this assignment was created by a Rule
	// Read Only: true
	IsRule *bool `json:"isRule"`

	// The user who modified the assignment
	// Read Only: true
	ModifiedBy *UserReference `json:"modifiedBy,omitempty"`

	// The Learning module object associated with this assignment
	Module *LearningModule `json:"module,omitempty"`

	// The user's percentage score for this assignment
	// Read Only: true
	PercentageScore float32 `json:"percentageScore,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The Learning Assignment state
	// Enum: [Assigned InProgress Completed Deleted]
	State string `json:"state,omitempty"`

	// The user to whom the assignment is assigned
	User *UserReference `json:"user,omitempty"`

	// The version of Learning module assigned
	Version int32 `json:"version,omitempty"`
}

// Validate validates this learning assignment
func (m *LearningAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssessment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssessmentForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateRecommendedForCompletion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningAssignment) validateAssessment(formats strfmt.Registry) error {

	if swag.IsZero(m.Assessment) { // not required
		return nil
	}

	if m.Assessment != nil {
		if err := m.Assessment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessment")
			}
			return err
		}
	}

	return nil
}

func (m *LearningAssignment) validateAssessmentForm(formats strfmt.Registry) error {

	if swag.IsZero(m.AssessmentForm) { // not required
		return nil
	}

	if m.AssessmentForm != nil {
		if err := m.AssessmentForm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assessmentForm")
			}
			return err
		}
	}

	return nil
}

func (m *LearningAssignment) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *LearningAssignment) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LearningAssignment) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LearningAssignment) validateDateRecommendedForCompletion(formats strfmt.Registry) error {

	if swag.IsZero(m.DateRecommendedForCompletion) { // not required
		return nil
	}

	if err := validate.FormatOf("dateRecommendedForCompletion", "body", "date-time", m.DateRecommendedForCompletion.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LearningAssignment) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *LearningAssignment) validateModule(formats strfmt.Registry) error {

	if swag.IsZero(m.Module) { // not required
		return nil
	}

	if m.Module != nil {
		if err := m.Module.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module")
			}
			return err
		}
	}

	return nil
}

func (m *LearningAssignment) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var learningAssignmentTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Assigned","InProgress","Completed","Deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningAssignmentTypeStatePropEnum = append(learningAssignmentTypeStatePropEnum, v)
	}
}

const (

	// LearningAssignmentStateAssigned captures enum value "Assigned"
	LearningAssignmentStateAssigned string = "Assigned"

	// LearningAssignmentStateInProgress captures enum value "InProgress"
	LearningAssignmentStateInProgress string = "InProgress"

	// LearningAssignmentStateCompleted captures enum value "Completed"
	LearningAssignmentStateCompleted string = "Completed"

	// LearningAssignmentStateDeleted captures enum value "Deleted"
	LearningAssignmentStateDeleted string = "Deleted"
)

// prop value enum
func (m *LearningAssignment) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningAssignmentTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningAssignment) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *LearningAssignment) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningAssignment) UnmarshalBinary(b []byte) error {
	var res LearningAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
