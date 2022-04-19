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

// OpenIntegrationUpdateRequest open integration update request
//
// swagger:model OpenIntegrationUpdateRequest
type OpenIntegrationUpdateRequest struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// messaging setting
	MessagingSetting *MessagingSettingReference `json:"messagingSetting,omitempty"`

	// The name of the Open messaging integration.
	// Required: true
	Name *string `json:"name"`

	// The outbound notification webhook signature secret token.
	OutboundNotificationWebhookSignatureSecretToken string `json:"outboundNotificationWebhookSignatureSecretToken,omitempty"`

	// The outbound notification webhook URL for the Open messaging integration.
	OutboundNotificationWebhookURL string `json:"outboundNotificationWebhookUrl,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Defines the SupportedContent profile configured for an integration
	SupportedContent *SupportedContentReference `json:"supportedContent,omitempty"`

	// The user specified headers for the Open messaging integration.
	WebhookHeaders map[string]string `json:"webhookHeaders,omitempty"`
}

// Validate validates this open integration update request
func (m *OpenIntegrationUpdateRequest) Validate(formats strfmt.Registry) error {
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

func (m *OpenIntegrationUpdateRequest) validateMessagingSetting(formats strfmt.Registry) error {

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

func (m *OpenIntegrationUpdateRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OpenIntegrationUpdateRequest) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OpenIntegrationUpdateRequest) validateSupportedContent(formats strfmt.Registry) error {

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
func (m *OpenIntegrationUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenIntegrationUpdateRequest) UnmarshalBinary(b []byte) error {
	var res OpenIntegrationUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
