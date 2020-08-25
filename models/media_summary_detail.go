// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MediaSummaryDetail media summary detail
//
// swagger:model MediaSummaryDetail
type MediaSummaryDetail struct {

	// active
	Active int32 `json:"active,omitempty"`

	// acw
	Acw int32 `json:"acw,omitempty"`
}

// Validate validates this media summary detail
func (m *MediaSummaryDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MediaSummaryDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaSummaryDetail) UnmarshalBinary(b []byte) error {
	var res MediaSummaryDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}