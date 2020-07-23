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

// SystemPromptAsset system prompt asset
//
// swagger:model SystemPromptAsset
type SystemPromptAsset struct {

	// duration seconds
	DurationSeconds float64 `json:"durationSeconds,omitempty"`

	// has default
	HasDefault bool `json:"hasDefault,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The asset resource language
	// Required: true
	Language *string `json:"language"`

	// language default
	LanguageDefault bool `json:"languageDefault,omitempty"`

	// media Uri
	MediaURI string `json:"mediaUri,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// prompt Id
	PromptID string `json:"promptId,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// tags
	Tags map[string][]string `json:"tags,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// tts string
	TtsString string `json:"ttsString,omitempty"`

	// upload status
	// Enum: [created uploaded transcoded transcodeFailed]
	UploadStatus string `json:"uploadStatus,omitempty"`

	// upload Uri
	UploadURI string `json:"uploadUri,omitempty"`
}

// Validate validates this system prompt asset
func (m *SystemPromptAsset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemPromptAsset) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *SystemPromptAsset) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var systemPromptAssetTypeUploadStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["created","uploaded","transcoded","transcodeFailed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		systemPromptAssetTypeUploadStatusPropEnum = append(systemPromptAssetTypeUploadStatusPropEnum, v)
	}
}

const (

	// SystemPromptAssetUploadStatusCreated captures enum value "created"
	SystemPromptAssetUploadStatusCreated string = "created"

	// SystemPromptAssetUploadStatusUploaded captures enum value "uploaded"
	SystemPromptAssetUploadStatusUploaded string = "uploaded"

	// SystemPromptAssetUploadStatusTranscoded captures enum value "transcoded"
	SystemPromptAssetUploadStatusTranscoded string = "transcoded"

	// SystemPromptAssetUploadStatusTranscodeFailed captures enum value "transcodeFailed"
	SystemPromptAssetUploadStatusTranscodeFailed string = "transcodeFailed"
)

// prop value enum
func (m *SystemPromptAsset) validateUploadStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, systemPromptAssetTypeUploadStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SystemPromptAsset) validateUploadStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.UploadStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateUploadStatusEnum("uploadStatus", "body", m.UploadStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemPromptAsset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemPromptAsset) UnmarshalBinary(b []byte) error {
	var res SystemPromptAsset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
