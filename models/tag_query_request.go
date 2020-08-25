// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TagQueryRequest tag query request
//
// swagger:model TagQueryRequest
type TagQueryRequest struct {

	// page number
	PageNumber int32 `json:"pageNumber,omitempty"`

	// page size
	PageSize int32 `json:"pageSize,omitempty"`

	// query
	Query string `json:"query,omitempty"`
}

// Validate validates this tag query request
func (m *TagQueryRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagQueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagQueryRequest) UnmarshalBinary(b []byte) error {
	var res TagQueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}