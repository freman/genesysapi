// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WhatsAppIntegrationUpdateRequest whats app integration update request
//
// swagger:model WhatsAppIntegrationUpdateRequest
type WhatsAppIntegrationUpdateRequest struct {

	// The action used to activate and then confirm a WhatsApp Integration.
	// Enum: [Activate Confirm]
	Action string `json:"action,omitempty"`

	// The authentication method used to confirm a WhatsApp Integration activation. If action is set to Activate, then authenticationMethod is a required field.
	// Enum: [Sms Voice]
	AuthenticationMethod string `json:"authenticationMethod,omitempty"`

	// The confirmation code sent by Whatsapp to you during the activation step. If action is set to Confirm, then confirmationCode is a required field.
	ConfirmationCode string `json:"confirmationCode,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// messaging setting
	MessagingSetting *MessagingSettingReference `json:"messagingSetting,omitempty"`

	// WhatsApp Integration name
	Name string `json:"name,omitempty"`

	// Phone number to associate with the WhatsApp integration
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Defines the SupportedContent profile configured for an integration
	SupportedContent *SupportedContentReference `json:"supportedContent,omitempty"`
}

// Validate validates this whats app integration update request
func (m *WhatsAppIntegrationUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

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

var whatsAppIntegrationUpdateRequestTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Activate","Confirm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		whatsAppIntegrationUpdateRequestTypeActionPropEnum = append(whatsAppIntegrationUpdateRequestTypeActionPropEnum, v)
	}
}

const (

	// WhatsAppIntegrationUpdateRequestActionActivate captures enum value "Activate"
	WhatsAppIntegrationUpdateRequestActionActivate string = "Activate"

	// WhatsAppIntegrationUpdateRequestActionConfirm captures enum value "Confirm"
	WhatsAppIntegrationUpdateRequestActionConfirm string = "Confirm"
)

// prop value enum
func (m *WhatsAppIntegrationUpdateRequest) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, whatsAppIntegrationUpdateRequestTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WhatsAppIntegrationUpdateRequest) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var whatsAppIntegrationUpdateRequestTypeAuthenticationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sms","Voice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		whatsAppIntegrationUpdateRequestTypeAuthenticationMethodPropEnum = append(whatsAppIntegrationUpdateRequestTypeAuthenticationMethodPropEnum, v)
	}
}

const (

	// WhatsAppIntegrationUpdateRequestAuthenticationMethodSms captures enum value "Sms"
	WhatsAppIntegrationUpdateRequestAuthenticationMethodSms string = "Sms"

	// WhatsAppIntegrationUpdateRequestAuthenticationMethodVoice captures enum value "Voice"
	WhatsAppIntegrationUpdateRequestAuthenticationMethodVoice string = "Voice"
)

// prop value enum
func (m *WhatsAppIntegrationUpdateRequest) validateAuthenticationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, whatsAppIntegrationUpdateRequestTypeAuthenticationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WhatsAppIntegrationUpdateRequest) validateAuthenticationMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationMethodEnum("authenticationMethod", "body", m.AuthenticationMethod); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegrationUpdateRequest) validateMessagingSetting(formats strfmt.Registry) error {

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

func (m *WhatsAppIntegrationUpdateRequest) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegrationUpdateRequest) validateSupportedContent(formats strfmt.Registry) error {

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
func (m *WhatsAppIntegrationUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WhatsAppIntegrationUpdateRequest) UnmarshalBinary(b []byte) error {
	var res WhatsAppIntegrationUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
