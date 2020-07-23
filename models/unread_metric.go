// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UnreadMetric unread metric
//
// swagger:model UnreadMetric
type UnreadMetric struct {

	// The count of unread alerts for a specific rule type.
	Count int32 `json:"count,omitempty"`
}

// Validate validates this unread metric
func (m *UnreadMetric) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UnreadMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnreadMetric) UnmarshalBinary(b []byte) error {
	var res UnreadMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
