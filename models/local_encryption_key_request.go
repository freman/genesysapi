// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LocalEncryptionKeyRequest local encryption key request
//
// swagger:model LocalEncryptionKeyRequest
type LocalEncryptionKeyRequest struct {

	// The local configuration id that contains metadata on private local service
	// Required: true
	ConfigID *string `json:"configId"`

	// The key pair id from the local service.
	// Required: true
	KeypairID *string `json:"keypairId"`

	// Base 64 encoded public key, generated by the local service.
	// Required: true
	PublicKey *string `json:"publicKey"`
}

// Validate validates this local encryption key request
func (m *LocalEncryptionKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeypairID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalEncryptionKeyRequest) validateConfigID(formats strfmt.Registry) error {

	if err := validate.Required("configId", "body", m.ConfigID); err != nil {
		return err
	}

	return nil
}

func (m *LocalEncryptionKeyRequest) validateKeypairID(formats strfmt.Registry) error {

	if err := validate.Required("keypairId", "body", m.KeypairID); err != nil {
		return err
	}

	return nil
}

func (m *LocalEncryptionKeyRequest) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this local encryption key request based on context it is used
func (m *LocalEncryptionKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocalEncryptionKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalEncryptionKeyRequest) UnmarshalBinary(b []byte) error {
	var res LocalEncryptionKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
