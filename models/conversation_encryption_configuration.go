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

// ConversationEncryptionConfiguration conversation encryption configuration
//
// swagger:model ConversationEncryptionConfiguration
type ConversationEncryptionConfiguration struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Type should be 'KmsSymmetric' when create or update Key configurations, 'None' to disable configuration.
	// Required: true
	// Enum: [KmsSymmetric LocalKeyManager Native None]
	KeyConfigurationType *string `json:"keyConfigurationType"`

	// The error message related to the configuration
	LastError *ErrorBody `json:"lastError,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// keyConfigurationType is always KmsSymmetric, and should be the arn to the key alias for the master key
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this conversation encryption configuration
func (m *ConversationEncryptionConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyConfigurationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversationEncryptionConfigurationTypeKeyConfigurationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["KmsSymmetric","LocalKeyManager","Native","None"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationEncryptionConfigurationTypeKeyConfigurationTypePropEnum = append(conversationEncryptionConfigurationTypeKeyConfigurationTypePropEnum, v)
	}
}

const (

	// ConversationEncryptionConfigurationKeyConfigurationTypeKmsSymmetric captures enum value "KmsSymmetric"
	ConversationEncryptionConfigurationKeyConfigurationTypeKmsSymmetric string = "KmsSymmetric"

	// ConversationEncryptionConfigurationKeyConfigurationTypeLocalKeyManager captures enum value "LocalKeyManager"
	ConversationEncryptionConfigurationKeyConfigurationTypeLocalKeyManager string = "LocalKeyManager"

	// ConversationEncryptionConfigurationKeyConfigurationTypeNative captures enum value "Native"
	ConversationEncryptionConfigurationKeyConfigurationTypeNative string = "Native"

	// ConversationEncryptionConfigurationKeyConfigurationTypeNone captures enum value "None"
	ConversationEncryptionConfigurationKeyConfigurationTypeNone string = "None"
)

// prop value enum
func (m *ConversationEncryptionConfiguration) validateKeyConfigurationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationEncryptionConfigurationTypeKeyConfigurationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationEncryptionConfiguration) validateKeyConfigurationType(formats strfmt.Registry) error {

	if err := validate.Required("keyConfigurationType", "body", m.KeyConfigurationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateKeyConfigurationTypeEnum("keyConfigurationType", "body", *m.KeyConfigurationType); err != nil {
		return err
	}

	return nil
}

func (m *ConversationEncryptionConfiguration) validateLastError(formats strfmt.Registry) error {

	if swag.IsZero(m.LastError) { // not required
		return nil
	}

	if m.LastError != nil {
		if err := m.LastError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastError")
			}
			return err
		}
	}

	return nil
}

func (m *ConversationEncryptionConfiguration) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConversationEncryptionConfiguration) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationEncryptionConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationEncryptionConfiguration) UnmarshalBinary(b []byte) error {
	var res ConversationEncryptionConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}