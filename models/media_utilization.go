// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MediaUtilization media utilization
//
// swagger:model MediaUtilization
type MediaUtilization struct {

	// If true, then track non-ACD conversations against utilization
	IncludeNonAcd bool `json:"includeNonAcd"`

	// Defines the list of other media types that can interrupt a conversation of this media type.  Values include call, chat, email, callback, and message.
	InterruptableMediaTypes []string `json:"interruptableMediaTypes"`

	// Defines the maximum number of conversations of this type that an agent can handle at one time.
	MaximumCapacity int32 `json:"maximumCapacity,omitempty"`
}

// Validate validates this media utilization
func (m *MediaUtilization) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this media utilization based on context it is used
func (m *MediaUtilization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MediaUtilization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaUtilization) UnmarshalBinary(b []byte) error {
	var res MediaUtilization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
