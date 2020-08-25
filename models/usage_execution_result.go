// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UsageExecutionResult usage execution result
//
// swagger:model UsageExecutionResult
type UsageExecutionResult struct {

	// The id of the query execution
	ExecutionID string `json:"executionId,omitempty"`

	// URI where the query results can be retrieved
	ResultsURI string `json:"resultsUri,omitempty"`
}

// Validate validates this usage execution result
func (m *UsageExecutionResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsageExecutionResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsageExecutionResult) UnmarshalBinary(b []byte) error {
	var res UsageExecutionResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}