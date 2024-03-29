// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EdgeMetricsDisk edge metrics disk
//
// swagger:model EdgeMetricsDisk
type EdgeMetricsDisk struct {

	// Available memory in bytes.
	AvailableBytes float64 `json:"availableBytes,omitempty"`

	// Disk partition name.
	PartitionName string `json:"partitionName,omitempty"`

	// Total memory in bytes.
	TotalBytes float64 `json:"totalBytes,omitempty"`
}

// Validate validates this edge metrics disk
func (m *EdgeMetricsDisk) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edge metrics disk based on context it is used
func (m *EdgeMetricsDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EdgeMetricsDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeMetricsDisk) UnmarshalBinary(b []byte) error {
	var res EdgeMetricsDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
