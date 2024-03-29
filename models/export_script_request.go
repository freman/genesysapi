// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExportScriptRequest Creating an exported script via Download Service
//
// swagger:model ExportScriptRequest
type ExportScriptRequest struct {

	// The final file name (no extension) of the script download: <fileName>.script
	FileName string `json:"fileName,omitempty"`

	// The UUID version of the script to be exported.  Defaults to the current editable version.
	VersionID string `json:"versionId,omitempty"`
}

// Validate validates this export script request
func (m *ExportScriptRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this export script request based on context it is used
func (m *ExportScriptRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExportScriptRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportScriptRequest) UnmarshalBinary(b []byte) error {
	var res ExportScriptRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
