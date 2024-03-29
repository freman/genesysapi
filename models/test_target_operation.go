// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TestTargetOperation Information about the Trigger test mode target validation step
//
// swagger:model TestTargetOperation
type TestTargetOperation struct {

	// Details about why the operation did or did not succeed
	Details []string `json:"details"`

	// Whether or not the operation matches expectations
	Matches bool `json:"matches"`

	// The name of the processing step
	Name string `json:"name,omitempty"`

	// The number of the processing step
	Step int32 `json:"step,omitempty"`
}

// Validate validates this test target operation
func (m *TestTargetOperation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this test target operation based on context it is used
func (m *TestTargetOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TestTargetOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestTargetOperation) UnmarshalBinary(b []byte) error {
	var res TestTargetOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
