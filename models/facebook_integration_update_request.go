// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FacebookIntegrationUpdateRequest facebook integration update request
//
// swagger:model FacebookIntegrationUpdateRequest
type FacebookIntegrationUpdateRequest struct {

	// The long-lived Page Access Token of a facebook page.
	// See https://developers.facebook.com/docs/facebook-login/access-tokens.
	// Either pageAccessToken or userAccessToken should be provided.
	PageAccessToken string `json:"pageAccessToken,omitempty"`

	// The short-lived User Access Token of the facebook user logged into the facebook app.
	// See https://developers.facebook.com/docs/facebook-login/access-tokens.
	// Either pageAccessToken or userAccessToken should be provided.
	UserAccessToken string `json:"userAccessToken,omitempty"`
}

// Validate validates this facebook integration update request
func (m *FacebookIntegrationUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FacebookIntegrationUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacebookIntegrationUpdateRequest) UnmarshalBinary(b []byte) error {
	var res FacebookIntegrationUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
