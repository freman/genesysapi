// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SegmentEstimateCount segment estimate count
//
// swagger:model SegmentEstimateCount
type SegmentEstimateCount struct {

	// Estimate count per segment.
	Count int32 `json:"count,omitempty"`

	// ID of Segment.
	SegmentID string `json:"segmentId,omitempty"`
}

// Validate validates this segment estimate count
func (m *SegmentEstimateCount) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SegmentEstimateCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentEstimateCount) UnmarshalBinary(b []byte) error {
	var res SegmentEstimateCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
