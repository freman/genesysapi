// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Agent agent
//
// swagger:model Agent
type Agent struct {

	// The current stage for this agent
	Stage string `json:"stage,omitempty"`
}

// Validate validates this agent
func (m *Agent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this agent based on context it is used
func (m *Agent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Agent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Agent) UnmarshalBinary(b []byte) error {
	var res Agent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
