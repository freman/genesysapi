// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AggregationResultEntry aggregation result entry
//
// swagger:model AggregationResultEntry
type AggregationResultEntry struct {

	// count
	Count int64 `json:"count,omitempty"`

	// For numericRange aggregations
	Gte float64 `json:"gte,omitempty"`

	// For numericRange aggregations
	Lt float64 `json:"lt,omitempty"`

	// For termFrequency aggregations
	Value string `json:"value,omitempty"`
}

// Validate validates this aggregation result entry
func (m *AggregationResultEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aggregation result entry based on context it is used
func (m *AggregationResultEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AggregationResultEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregationResultEntry) UnmarshalBinary(b []byte) error {
	var res AggregationResultEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
