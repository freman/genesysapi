// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AggregationRange aggregation range
//
// swagger:model AggregationRange
type AggregationRange struct {

	// Greater than or equal to
	Gte float64 `json:"gte,omitempty"`

	// Less than
	Lt float64 `json:"lt,omitempty"`
}

// Validate validates this aggregation range
func (m *AggregationRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aggregation range based on context it is used
func (m *AggregationRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AggregationRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregationRange) UnmarshalBinary(b []byte) error {
	var res AggregationRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
