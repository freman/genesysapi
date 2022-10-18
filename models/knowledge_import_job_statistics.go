// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KnowledgeImportJobStatistics knowledge import job statistics
//
// swagger:model KnowledgeImportJobStatistics
type KnowledgeImportJobStatistics struct {

	// Number of categories failed to import.
	CountCategoryImportFailure int32 `json:"countCategoryImportFailure,omitempty"`

	// Number of imported categories.
	CountCategoryImportSuccess int32 `json:"countCategoryImportSuccess,omitempty"`

	// Number of categories that failed validation for import.
	CountCategoryValidationFailure int32 `json:"countCategoryValidationFailure,omitempty"`

	// Number of categories that validated successfully for import.
	CountCategoryValidationSuccess int32 `json:"countCategoryValidationSuccess,omitempty"`

	// Number of documents will be created by the import.
	CountDocumentImportActivityCreate int32 `json:"countDocumentImportActivityCreate,omitempty"`

	// Number of documents will be updated by the import.
	CountDocumentImportActivityUpdate int32 `json:"countDocumentImportActivityUpdate,omitempty"`

	// Number of documents failed to import.
	CountDocumentImportFailure int32 `json:"countDocumentImportFailure,omitempty"`

	// Number of imported documents.
	CountDocumentImportSuccess int32 `json:"countDocumentImportSuccess,omitempty"`

	// Number of documents will be imported as draft.
	CountDocumentStateDraft int32 `json:"countDocumentStateDraft,omitempty"`

	// Number of documents will be imported as published.
	CountDocumentStatePublished int32 `json:"countDocumentStatePublished,omitempty"`

	// Number of documents that failed validation for import.
	CountDocumentValidationFailure int32 `json:"countDocumentValidationFailure,omitempty"`

	// Number of documents that validated successfully for import.
	CountDocumentValidationSuccess int32 `json:"countDocumentValidationSuccess,omitempty"`

	// Number of labels failed to import.
	CountLabelImportFailure int32 `json:"countLabelImportFailure,omitempty"`

	// Number of imported labels.
	CountLabelImportSuccess int32 `json:"countLabelImportSuccess,omitempty"`

	// Number of labels that failed validation for import.
	CountLabelValidationFailure int32 `json:"countLabelValidationFailure,omitempty"`

	// Number of labels that validated successfully for import.
	CountLabelValidationSuccess int32 `json:"countLabelValidationSuccess,omitempty"`

	// Shows whether the import treated as migration or not.
	MigrationDetected bool `json:"migrationDetected"`
}

// Validate validates this knowledge import job statistics
func (m *KnowledgeImportJobStatistics) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeImportJobStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeImportJobStatistics) UnmarshalBinary(b []byte) error {
	var res KnowledgeImportJobStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
