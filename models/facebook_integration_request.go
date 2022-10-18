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

// FacebookIntegrationRequest facebook integration request
//
// swagger:model FacebookIntegrationRequest
type FacebookIntegrationRequest struct {

	// The app Id of Facebook app. The appId is required when a customer wants to use their own approved Facebook app.
	AppID string `json:"appId,omitempty"`

	// The app Secret of Facebook app. The appSecret is required when appId is provided.
	AppSecret string `json:"appSecret,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Defines the message settings to be applied for this integration
	MessagingSetting *MessagingSettingRequestReference `json:"messagingSetting,omitempty"`

	// The name of the Facebook Integration
	// Required: true
	Name *string `json:"name"`

	// The long-lived Page Access Token of Facebook page.
	// See https://developers.facebook.com/docs/facebook-login/access-tokens.
	// When a pageAccessToken is provided, pageId and userAccessToken are not required.
	PageAccessToken string `json:"pageAccessToken,omitempty"`

	// The page Id of Facebook page. The pageId is required when userAccessToken is provided.
	PageID string `json:"pageId,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Defines the SupportedContent profile configured for an integration
	SupportedContent *SupportedContentReference `json:"supportedContent,omitempty"`

	// The short-lived User Access Token of the Facebook user logged into the Facebook app.
	// See https://developers.facebook.com/docs/facebook-login/access-tokens.
	// When userAccessToken is provided, pageId is mandatory.
	// When userAccessToken/pageId combination is provided, pageAccessToken is not required.
	UserAccessToken string `json:"userAccessToken,omitempty"`
}

// Validate validates this facebook integration request
func (m *FacebookIntegrationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessagingSetting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *FacebookIntegrationRequest) validateMessagingSetting(formats strfmt.Registry) error {

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

func (m *FacebookIntegrationRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegrationRequest) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegrationRequest) validateSupportedContent(formats strfmt.Registry) error {

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
func (m *FacebookIntegrationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacebookIntegrationRequest) UnmarshalBinary(b []byte) error {
	var res FacebookIntegrationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
