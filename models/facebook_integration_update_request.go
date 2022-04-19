// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FacebookIntegrationUpdateRequest facebook integration update request
//
// swagger:model FacebookIntegrationUpdateRequest
type FacebookIntegrationUpdateRequest struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// messaging setting
	MessagingSetting *MessagingSettingReference `json:"messagingSetting,omitempty"`

	// The name of the Facebook Integration
	Name string `json:"name,omitempty"`

	// The long-lived Page Access Token of Facebook page.
	// See https://developers.facebook.com/docs/facebook-login/access-tokens.
	// Either pageAccessToken or userAccessToken should be provided.
	PageAccessToken string `json:"pageAccessToken,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Defines the SupportedContent profile configured for an integration
	SupportedContent *SupportedContentReference `json:"supportedContent,omitempty"`

	// The short-lived User Access Token of the Facebook user logged into the Facebook app.
	// See https://developers.facebook.com/docs/facebook-login/access-tokens.
	// Either pageAccessToken or userAccessToken should be provided.
	UserAccessToken string `json:"userAccessToken,omitempty"`
}

// Validate validates this facebook integration update request
func (m *FacebookIntegrationUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessagingSetting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacebookIntegrationUpdateRequest) validateMessagingSetting(formats strfmt.Registry) error {

	if swag.IsZero(m.MessagingSetting) { // not required
		return nil
	}

	if m.MessagingSetting != nil {
		if err := m.MessagingSetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagingSetting")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegrationUpdateRequest) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegrationUpdateRequest) validateSupportedContent(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedContent) { // not required
		return nil
	}

	if m.SupportedContent != nil {
		if err := m.SupportedContent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supportedContent")
			}
			return err
		}
	}

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
