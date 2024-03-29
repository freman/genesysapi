// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EncryptionKey encryption key
//
// swagger:model EncryptionKey
type EncryptionKey struct {

	// create date of the key pair. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreateDate strfmt.DateTime `json:"createDate,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Key type used in this configuration
	// Enum: [KmsSymmetric LocalKeyManager Native None]
	KeyConfigurationType string `json:"keyConfigurationType,omitempty"`

	// key data summary (base 64 encoded public key)
	KeydataSummary string `json:"keydataSummary,omitempty"`

	// ARN of internal key to be wrapped by AWS KMS Symmetric key
	KmsKeyArn string `json:"kmsKeyArn,omitempty"`

	// Local configuration
	LocalEncryptionConfiguration *LocalEncryptionConfiguration `json:"localEncryptionConfiguration,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// user that requested generation of public key
	User *User `json:"user,omitempty"`
}

// Validate validates this encryption key
func (m *EncryptionKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyConfigurationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalEncryptionConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionKey) validateCreateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createDate", "body", "date-time", m.CreateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var encryptionKeyTypeKeyConfigurationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["KmsSymmetric","LocalKeyManager","Native","None"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		encryptionKeyTypeKeyConfigurationTypePropEnum = append(encryptionKeyTypeKeyConfigurationTypePropEnum, v)
	}
}

const (

	// EncryptionKeyKeyConfigurationTypeKmsSymmetric captures enum value "KmsSymmetric"
	EncryptionKeyKeyConfigurationTypeKmsSymmetric string = "KmsSymmetric"

	// EncryptionKeyKeyConfigurationTypeLocalKeyManager captures enum value "LocalKeyManager"
	EncryptionKeyKeyConfigurationTypeLocalKeyManager string = "LocalKeyManager"

	// EncryptionKeyKeyConfigurationTypeNative captures enum value "Native"
	EncryptionKeyKeyConfigurationTypeNative string = "Native"

	// EncryptionKeyKeyConfigurationTypeNone captures enum value "None"
	EncryptionKeyKeyConfigurationTypeNone string = "None"
)

// prop value enum
func (m *EncryptionKey) validateKeyConfigurationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, encryptionKeyTypeKeyConfigurationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EncryptionKey) validateKeyConfigurationType(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyConfigurationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateKeyConfigurationTypeEnum("keyConfigurationType", "body", m.KeyConfigurationType); err != nil {
		return err
	}

	return nil
}

func (m *EncryptionKey) validateLocalEncryptionConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalEncryptionConfiguration) { // not required
		return nil
	}

	if m.LocalEncryptionConfiguration != nil {
		if err := m.LocalEncryptionConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localEncryptionConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localEncryptionConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionKey) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EncryptionKey) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this encryption key based on the context it is used
func (m *EncryptionKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalEncryptionConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionKey) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *EncryptionKey) contextValidateLocalEncryptionConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalEncryptionConfiguration != nil {
		if err := m.LocalEncryptionConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localEncryptionConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localEncryptionConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionKey) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *EncryptionKey) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionKey) UnmarshalBinary(b []byte) error {
	var res EncryptionKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
