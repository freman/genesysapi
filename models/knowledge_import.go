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

// KnowledgeImport knowledge import
//
// swagger:model KnowledgeImport
type KnowledgeImport struct {

	// Created date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// file type of the document
	// Required: true
	// Enum: [Csv JsonLines]
	FileType *string `json:"fileType"`

	// Id of the import operation
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Ignore headers for the specified file
	IgnoreHeaders bool `json:"ignoreHeaders"`

	// Knowledge base which document import does belong to
	// Read Only: true
	KnowledgeBase *KnowledgeBase `json:"knowledgeBase,omitempty"`

	// Language code
	// Read Only: true
	// Enum: [en-US en-UK en-AU de-DE]
	LanguageCode string `json:"languageCode,omitempty"`

	// Name of the import operation
	Name string `json:"name,omitempty"`

	// Report of the import operation
	// Read Only: true
	Report *ImportReport `json:"report,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Status of the operation
	// Read Only: true
	// Enum: [Created ValidationInProgress ValidationCompleted ValidationFailed Started InProgress Completed PartialCompleted Failed]
	Status string `json:"status,omitempty"`

	// Upload key
	// Required: true
	UploadKey *string `json:"uploadKey"`
}

// Validate validates this knowledge import
func (m *KnowledgeImport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKnowledgeBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeImport) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeImport) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

var knowledgeImportTypeFileTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Csv","JsonLines"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeImportTypeFileTypePropEnum = append(knowledgeImportTypeFileTypePropEnum, v)
	}
}

const (

	// KnowledgeImportFileTypeCsv captures enum value "Csv"
	KnowledgeImportFileTypeCsv string = "Csv"

	// KnowledgeImportFileTypeJSONLines captures enum value "JsonLines"
	KnowledgeImportFileTypeJSONLines string = "JsonLines"
)

// prop value enum
func (m *KnowledgeImport) validateFileTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeImportTypeFileTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeImport) validateFileType(formats strfmt.Registry) error {

	if err := validate.Required("fileType", "body", m.FileType); err != nil {
		return err
	}

	// value enum
	if err := m.validateFileTypeEnum("fileType", "body", *m.FileType); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeImport) validateKnowledgeBase(formats strfmt.Registry) error {

	if swag.IsZero(m.KnowledgeBase) { // not required
		return nil
	}

	if m.KnowledgeBase != nil {
		if err := m.KnowledgeBase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("knowledgeBase")
			}
			return err
		}
	}

	return nil
}

var knowledgeImportTypeLanguageCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-US","en-UK","en-AU","de-DE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeImportTypeLanguageCodePropEnum = append(knowledgeImportTypeLanguageCodePropEnum, v)
	}
}

const (

	// KnowledgeImportLanguageCodeEnUS captures enum value "en-US"
	KnowledgeImportLanguageCodeEnUS string = "en-US"

	// KnowledgeImportLanguageCodeEnUK captures enum value "en-UK"
	KnowledgeImportLanguageCodeEnUK string = "en-UK"

	// KnowledgeImportLanguageCodeEnAU captures enum value "en-AU"
	KnowledgeImportLanguageCodeEnAU string = "en-AU"

	// KnowledgeImportLanguageCodeDeDE captures enum value "de-DE"
	KnowledgeImportLanguageCodeDeDE string = "de-DE"
)

// prop value enum
func (m *KnowledgeImport) validateLanguageCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeImportTypeLanguageCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeImport) validateLanguageCode(formats strfmt.Registry) error {

	if swag.IsZero(m.LanguageCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateLanguageCodeEnum("languageCode", "body", m.LanguageCode); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeImport) validateReport(formats strfmt.Registry) error {

	if swag.IsZero(m.Report) { // not required
		return nil
	}

	if m.Report != nil {
		if err := m.Report.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeImport) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var knowledgeImportTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Created","ValidationInProgress","ValidationCompleted","ValidationFailed","Started","InProgress","Completed","PartialCompleted","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeImportTypeStatusPropEnum = append(knowledgeImportTypeStatusPropEnum, v)
	}
}

const (

	// KnowledgeImportStatusCreated captures enum value "Created"
	KnowledgeImportStatusCreated string = "Created"

	// KnowledgeImportStatusValidationInProgress captures enum value "ValidationInProgress"
	KnowledgeImportStatusValidationInProgress string = "ValidationInProgress"

	// KnowledgeImportStatusValidationCompleted captures enum value "ValidationCompleted"
	KnowledgeImportStatusValidationCompleted string = "ValidationCompleted"

	// KnowledgeImportStatusValidationFailed captures enum value "ValidationFailed"
	KnowledgeImportStatusValidationFailed string = "ValidationFailed"

	// KnowledgeImportStatusStarted captures enum value "Started"
	KnowledgeImportStatusStarted string = "Started"

	// KnowledgeImportStatusInProgress captures enum value "InProgress"
	KnowledgeImportStatusInProgress string = "InProgress"

	// KnowledgeImportStatusCompleted captures enum value "Completed"
	KnowledgeImportStatusCompleted string = "Completed"

	// KnowledgeImportStatusPartialCompleted captures enum value "PartialCompleted"
	KnowledgeImportStatusPartialCompleted string = "PartialCompleted"

	// KnowledgeImportStatusFailed captures enum value "Failed"
	KnowledgeImportStatusFailed string = "Failed"
)

// prop value enum
func (m *KnowledgeImport) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeImportTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeImport) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeImport) validateUploadKey(formats strfmt.Registry) error {

	if err := validate.Required("uploadKey", "body", m.UploadKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeImport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeImport) UnmarshalBinary(b []byte) error {
	var res KnowledgeImport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
