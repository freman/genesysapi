// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KnowledgeExportJobDocumentsFilter knowledge export job documents filter
//
// swagger:model KnowledgeExportJobDocumentsFilter
type KnowledgeExportJobDocumentsFilter struct {

	// Retrieves the documents modified in specified date and time range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval string `json:"interval,omitempty"`
}

// Validate validates this knowledge export job documents filter
func (m *KnowledgeExportJobDocumentsFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeExportJobDocumentsFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeExportJobDocumentsFilter) UnmarshalBinary(b []byte) error {
	var res KnowledgeExportJobDocumentsFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
