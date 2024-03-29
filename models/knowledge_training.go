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

// KnowledgeTraining knowledge training
//
// swagger:model KnowledgeTraining
type KnowledgeTraining struct {

	// Training completed date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCompleted strfmt.DateTime `json:"dateCompleted,omitempty"`

	// Trained Documents Promoted date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DatePromoted strfmt.DateTime `json:"datePromoted,omitempty"`

	// Trigger date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateTriggered strfmt.DateTime `json:"dateTriggered,omitempty"`

	// Any error message during the Training or Promote action.
	// Read Only: true
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Knowledge Base that the training belongs to.
	// Read Only: true
	KnowledgeBase *KnowledgeBase `json:"knowledgeBase,omitempty"`

	// State of the Trained Documents, which can be one of these Draft, Active, Discarded, Archived.
	// Read Only: true
	// Enum: [Draft Active Discarded Archived]
	KnowledgeDocumentsState string `json:"knowledgeDocumentsState,omitempty"`

	// Language of the documents that are trained.
	// Read Only: true
	LanguageCode string `json:"languageCode,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Training status.
	// Read Only: true
	// Enum: [Queued InProgress Succeeded Failed]
	Status string `json:"status,omitempty"`
}

// Validate validates this knowledge training
func (m *KnowledgeTraining) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatePromoted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTriggered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKnowledgeBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKnowledgeDocumentsState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *KnowledgeTraining) validateDateCompleted(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCompleted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCompleted", "body", "date-time", m.DateCompleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) validateDatePromoted(formats strfmt.Registry) error {
	if swag.IsZero(m.DatePromoted) { // not required
		return nil
	}

	if err := validate.FormatOf("datePromoted", "body", "date-time", m.DatePromoted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) validateDateTriggered(formats strfmt.Registry) error {
	if swag.IsZero(m.DateTriggered) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTriggered", "body", "date-time", m.DateTriggered.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) validateKnowledgeBase(formats strfmt.Registry) error {
	if swag.IsZero(m.KnowledgeBase) { // not required
		return nil
	}

	if m.KnowledgeBase != nil {
		if err := m.KnowledgeBase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("knowledgeBase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("knowledgeBase")
			}
			return err
		}
	}

	return nil
}

var knowledgeTrainingTypeKnowledgeDocumentsStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Draft","Active","Discarded","Archived"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeTrainingTypeKnowledgeDocumentsStatePropEnum = append(knowledgeTrainingTypeKnowledgeDocumentsStatePropEnum, v)
	}
}

const (

	// KnowledgeTrainingKnowledgeDocumentsStateDraft captures enum value "Draft"
	KnowledgeTrainingKnowledgeDocumentsStateDraft string = "Draft"

	// KnowledgeTrainingKnowledgeDocumentsStateActive captures enum value "Active"
	KnowledgeTrainingKnowledgeDocumentsStateActive string = "Active"

	// KnowledgeTrainingKnowledgeDocumentsStateDiscarded captures enum value "Discarded"
	KnowledgeTrainingKnowledgeDocumentsStateDiscarded string = "Discarded"

	// KnowledgeTrainingKnowledgeDocumentsStateArchived captures enum value "Archived"
	KnowledgeTrainingKnowledgeDocumentsStateArchived string = "Archived"
)

// prop value enum
func (m *KnowledgeTraining) validateKnowledgeDocumentsStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeTrainingTypeKnowledgeDocumentsStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeTraining) validateKnowledgeDocumentsState(formats strfmt.Registry) error {
	if swag.IsZero(m.KnowledgeDocumentsState) { // not required
		return nil
	}

	// value enum
	if err := m.validateKnowledgeDocumentsStateEnum("knowledgeDocumentsState", "body", m.KnowledgeDocumentsState); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var knowledgeTrainingTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Queued","InProgress","Succeeded","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeTrainingTypeStatusPropEnum = append(knowledgeTrainingTypeStatusPropEnum, v)
	}
}

const (

	// KnowledgeTrainingStatusQueued captures enum value "Queued"
	KnowledgeTrainingStatusQueued string = "Queued"

	// KnowledgeTrainingStatusInProgress captures enum value "InProgress"
	KnowledgeTrainingStatusInProgress string = "InProgress"

	// KnowledgeTrainingStatusSucceeded captures enum value "Succeeded"
	KnowledgeTrainingStatusSucceeded string = "Succeeded"

	// KnowledgeTrainingStatusFailed captures enum value "Failed"
	KnowledgeTrainingStatusFailed string = "Failed"
)

// prop value enum
func (m *KnowledgeTraining) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeTrainingTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeTraining) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this knowledge training based on the context it is used
func (m *KnowledgeTraining) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateCompleted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatePromoted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateTriggered(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKnowledgeBase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKnowledgeDocumentsState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeTraining) contextValidateDateCompleted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCompleted", "body", strfmt.DateTime(m.DateCompleted)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateDatePromoted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "datePromoted", "body", strfmt.DateTime(m.DatePromoted)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateDateTriggered(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateTriggered", "body", strfmt.DateTime(m.DateTriggered)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateErrorMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "errorMessage", "body", string(m.ErrorMessage)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateKnowledgeBase(ctx context.Context, formats strfmt.Registry) error {

	if m.KnowledgeBase != nil {
		if err := m.KnowledgeBase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("knowledgeBase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("knowledgeBase")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateKnowledgeDocumentsState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "knowledgeDocumentsState", "body", string(m.KnowledgeDocumentsState)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateLanguageCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "languageCode", "body", string(m.LanguageCode)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeTraining) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeTraining) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeTraining) UnmarshalBinary(b []byte) error {
	var res KnowledgeTraining
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
