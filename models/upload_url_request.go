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

// UploadURLRequest upload Url request
//
// swagger:model UploadUrlRequest
type UploadURLRequest struct {

	// Content MD5 of the file to upload
	ContentMd5 string `json:"contentMd5,omitempty"`

	// The content type of the file to upload. Allows all MIME types
	ContentType string `json:"contentType,omitempty"`

	// Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \{^}%`]">[~<#|
	FileName string `json:"fileName,omitempty"`

	// server side encryption
	// Enum: [AES256]
	ServerSideEncryption string `json:"serverSideEncryption,omitempty"`

	// The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
	SignedURLTimeoutSeconds int32 `json:"signedUrlTimeoutSeconds,omitempty"`
}

// Validate validates this upload Url request
func (m *UploadURLRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerSideEncryption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var uploadUrlRequestTypeServerSideEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AES256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		uploadUrlRequestTypeServerSideEncryptionPropEnum = append(uploadUrlRequestTypeServerSideEncryptionPropEnum, v)
	}
}

const (

	// UploadURLRequestServerSideEncryptionAES256 captures enum value "AES256"
	UploadURLRequestServerSideEncryptionAES256 string = "AES256"
)

// prop value enum
func (m *UploadURLRequest) validateServerSideEncryptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, uploadUrlRequestTypeServerSideEncryptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UploadURLRequest) validateServerSideEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerSideEncryption) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerSideEncryptionEnum("serverSideEncryption", "body", m.ServerSideEncryption); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this upload Url request based on context it is used
func (m *UploadURLRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UploadURLRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadURLRequest) UnmarshalBinary(b []byte) error {
	var res UploadURLRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
