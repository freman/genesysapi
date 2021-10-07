// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReportingTurnIntentSlot reporting turn intent slot
//
// swagger:model ReportingTurnIntentSlot
type ReportingTurnIntentSlot struct {

	// The confidence score this slot received during detection.
	Confidence float64 `json:"confidence,omitempty"`

	// The name of the slot.
	Name string `json:"name,omitempty"`

	// The NLU entity type of the slot (either builtin or user defined)
	Type string `json:"type,omitempty"`

	// The value of the slot.
	Value string `json:"value,omitempty"`
}

// Validate validates this reporting turn intent slot
func (m *ReportingTurnIntentSlot) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportingTurnIntentSlot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingTurnIntentSlot) UnmarshalBinary(b []byte) error {
	var res ReportingTurnIntentSlot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
