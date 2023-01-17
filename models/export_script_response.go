// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExportScriptResponse export script response
//
// swagger:model ExportScriptResponse
type ExportScriptResponse struct {

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this export script response
func (m *ExportScriptResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this export script response based on context it is used
func (m *ExportScriptResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExportScriptResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportScriptResponse) UnmarshalBinary(b []byte) error {
	var res ExportScriptResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
