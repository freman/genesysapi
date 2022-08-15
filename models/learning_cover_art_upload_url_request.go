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

// LearningCoverArtUploadURLRequest learning cover art upload Url request
//
// swagger:model LearningCoverArtUploadUrlRequest
type LearningCoverArtUploadURLRequest struct {

	// Content MD5 of the file to upload
	ContentMd5 string `json:"contentMd5,omitempty"`

	// The content type of the file to upload.
	// Required: true
	// Enum: [image/bmp image/gif image/jpeg image/jpg image/png]
	ContentType *string `json:"contentType"`

	// Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \{^}%`]">[~<#|
	FileName string `json:"fileName,omitempty"`

	// server side encryption
	// Enum: [AES256]
	ServerSideEncryption string `json:"serverSideEncryption,omitempty"`

	// The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
	SignedURLTimeoutSeconds int32 `json:"signedUrlTimeoutSeconds,omitempty"`
}

// Validate validates this learning cover art upload Url request
func (m *LearningCoverArtUploadURLRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerSideEncryption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var learningCoverArtUploadUrlRequestTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["image/bmp","image/gif","image/jpeg","image/jpg","image/png"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningCoverArtUploadUrlRequestTypeContentTypePropEnum = append(learningCoverArtUploadUrlRequestTypeContentTypePropEnum, v)
	}
}

const (

	// LearningCoverArtUploadURLRequestContentTypeImageBmp captures enum value "image/bmp"
	LearningCoverArtUploadURLRequestContentTypeImageBmp string = "image/bmp"

	// LearningCoverArtUploadURLRequestContentTypeImageGif captures enum value "image/gif"
	LearningCoverArtUploadURLRequestContentTypeImageGif string = "image/gif"

	// LearningCoverArtUploadURLRequestContentTypeImageJpeg captures enum value "image/jpeg"
	LearningCoverArtUploadURLRequestContentTypeImageJpeg string = "image/jpeg"

	// LearningCoverArtUploadURLRequestContentTypeImageJpg captures enum value "image/jpg"
	LearningCoverArtUploadURLRequestContentTypeImageJpg string = "image/jpg"

	// LearningCoverArtUploadURLRequestContentTypeImagePng captures enum value "image/png"
	LearningCoverArtUploadURLRequestContentTypeImagePng string = "image/png"
)

// prop value enum
func (m *LearningCoverArtUploadURLRequest) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningCoverArtUploadUrlRequestTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningCoverArtUploadURLRequest) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("contentType", "body", m.ContentType); err != nil {
		return err
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", *m.ContentType); err != nil {
		return err
	}

	return nil
}

var learningCoverArtUploadUrlRequestTypeServerSideEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AES256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningCoverArtUploadUrlRequestTypeServerSideEncryptionPropEnum = append(learningCoverArtUploadUrlRequestTypeServerSideEncryptionPropEnum, v)
	}
}

const (

	// LearningCoverArtUploadURLRequestServerSideEncryptionAES256 captures enum value "AES256"
	LearningCoverArtUploadURLRequestServerSideEncryptionAES256 string = "AES256"
)

// prop value enum
func (m *LearningCoverArtUploadURLRequest) validateServerSideEncryptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningCoverArtUploadUrlRequestTypeServerSideEncryptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningCoverArtUploadURLRequest) validateServerSideEncryption(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerSideEncryption) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerSideEncryptionEnum("serverSideEncryption", "body", m.ServerSideEncryption); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningCoverArtUploadURLRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningCoverArtUploadURLRequest) UnmarshalBinary(b []byte) error {
	var res LearningCoverArtUploadURLRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
