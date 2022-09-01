// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AssignedLearningModule Learning module response
//
// swagger:model AssignedLearningModule
type AssignedLearningModule struct {

	// The assessment form for learning module
	AssessmentForm *AssessmentForm `json:"assessmentForm,omitempty"`

	// The completion time of learning module in days
	// Required: true
	CompletionTimeInDays *int32 `json:"completionTimeInDays"`

	// The cover art for the learning module
	CoverArt *LearningModuleCoverArtResponse `json:"coverArt,omitempty"`

	// The user who created learning module
	// Read Only: true
	CreatedBy *UserReference `json:"createdBy,omitempty"`

	// The current assignments for the requested users
	CurrentAssignments []*LearningAssignment `json:"currentAssignments"`

	// The date/time learning module was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date/time learning module was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The description of learning module
	Description string `json:"description,omitempty"`

	// The external ID of the learning module
	// Read Only: true
	ExternalID string `json:"externalId,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The list of inform steps in a learning module
	InformSteps []*LearningModuleInformStep `json:"informSteps"`

	// If true, learning module is archived
	// Read Only: true
	IsArchived *bool `json:"isArchived"`

	// If true, learning module is published
	// Read Only: true
	IsPublished *bool `json:"isPublished"`

	// The user who modified learning module
	// Read Only: true
	ModifiedBy *UserReference `json:"modifiedBy,omitempty"`

	// The name of learning module
	// Required: true
	Name *string `json:"name"`

	// The rule for learning module; read-only, and only populated when requested via expand param.
	// Read Only: true
	Rule *LearningModuleRule `json:"rule,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The source of the learning module
	// Read Only: true
	// Enum: [UserCreated GenesysBeyond]
	Source string `json:"source,omitempty"`

	// The learning module summary data
	SummaryData *LearningModuleSummary `json:"summaryData,omitempty"`

	// The type for the learning module
	// Enum: [Informational AssessedContent Assessment]
	Type string `json:"type,omitempty"`

	// The version of published learning module
	// Read Only: true
	Version int32 `json:"version,omitempty"`
}

// Validate validates this assigned learning module
func (m *AssignedLearningModule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssessmentForm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletionTimeInDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoverArt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInformSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryData(formats); err != nil {
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

func (m *AssignedLearningModule) validateAssessmentForm(formats strfmt.Registry) error {

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

func (m *AssignedLearningModule) validateCompletionTimeInDays(formats strfmt.Registry) error {

	if err := validate.Required("completionTimeInDays", "body", m.CompletionTimeInDays); err != nil {
		return err
	}

	return nil
}

func (m *AssignedLearningModule) validateCoverArt(formats strfmt.Registry) error {

	if swag.IsZero(m.CoverArt) { // not required
		return nil
	}

	if m.CoverArt != nil {
		if err := m.CoverArt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coverArt")
			}
			return err
		}
	}

	return nil
}

func (m *AssignedLearningModule) validateCreatedBy(formats strfmt.Registry) error {

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

func (m *AssignedLearningModule) validateCurrentAssignments(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentAssignments) { // not required
		return nil
	}

	for i := 0; i < len(m.CurrentAssignments); i++ {
		if swag.IsZero(m.CurrentAssignments[i]) { // not required
			continue
		}

		if m.CurrentAssignments[i] != nil {
			if err := m.CurrentAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("currentAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssignedLearningModule) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AssignedLearningModule) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AssignedLearningModule) validateInformSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.InformSteps) { // not required
		return nil
	}

	for i := 0; i < len(m.InformSteps); i++ {
		if swag.IsZero(m.InformSteps[i]) { // not required
			continue
		}

		if m.InformSteps[i] != nil {
			if err := m.InformSteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("informSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AssignedLearningModule) validateModifiedBy(formats strfmt.Registry) error {

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

func (m *AssignedLearningModule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AssignedLearningModule) validateRule(formats strfmt.Registry) error {

	if swag.IsZero(m.Rule) { // not required
		return nil
	}

	if m.Rule != nil {
		if err := m.Rule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule")
			}
			return err
		}
	}

	return nil
}

func (m *AssignedLearningModule) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var assignedLearningModuleTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UserCreated","GenesysBeyond"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assignedLearningModuleTypeSourcePropEnum = append(assignedLearningModuleTypeSourcePropEnum, v)
	}
}

const (

	// AssignedLearningModuleSourceUserCreated captures enum value "UserCreated"
	AssignedLearningModuleSourceUserCreated string = "UserCreated"

	// AssignedLearningModuleSourceGenesysBeyond captures enum value "GenesysBeyond"
	AssignedLearningModuleSourceGenesysBeyond string = "GenesysBeyond"
)

// prop value enum
func (m *AssignedLearningModule) validateSourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assignedLearningModuleTypeSourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AssignedLearningModule) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *AssignedLearningModule) validateSummaryData(formats strfmt.Registry) error {

	if swag.IsZero(m.SummaryData) { // not required
		return nil
	}

	if m.SummaryData != nil {
		if err := m.SummaryData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryData")
			}
			return err
		}
	}

	return nil
}

var assignedLearningModuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Informational","AssessedContent","Assessment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assignedLearningModuleTypeTypePropEnum = append(assignedLearningModuleTypeTypePropEnum, v)
	}
}

const (

	// AssignedLearningModuleTypeInformational captures enum value "Informational"
	AssignedLearningModuleTypeInformational string = "Informational"

	// AssignedLearningModuleTypeAssessedContent captures enum value "AssessedContent"
	AssignedLearningModuleTypeAssessedContent string = "AssessedContent"

	// AssignedLearningModuleTypeAssessment captures enum value "Assessment"
	AssignedLearningModuleTypeAssessment string = "Assessment"
)

// prop value enum
func (m *AssignedLearningModule) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assignedLearningModuleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AssignedLearningModule) validateType(formats strfmt.Registry) error {

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
func (m *AssignedLearningModule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignedLearningModule) UnmarshalBinary(b []byte) error {
	var res AssignedLearningModule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}