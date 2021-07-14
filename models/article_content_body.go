// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArticleContentBody article content body
//
// swagger:model ArticleContentBody
type ArticleContentBody struct {

	// Presigned URL to retrieve the document content.
	// Read Only: true
	LocationURL string `json:"locationUrl,omitempty"`
}

// Validate validates this article content body
func (m *ArticleContentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArticleContentBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArticleContentBody) UnmarshalBinary(b []byte) error {
	var res ArticleContentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
