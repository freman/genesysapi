// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecordingMetadata recording metadata
//
// swagger:model RecordingMetadata
type RecordingMetadata struct {

	// Annotations that belong to the recording. Populated when recording filestate is AVAILABLE.
	Annotations []*Annotation `json:"annotations"`

	// The date the recording will be archived. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ArchiveDate strfmt.DateTime `json:"archiveDate,omitempty"`

	// The type of archive medium used. Example: CloudArchive
	// Enum: [CLOUDARCHIVE]
	ArchiveMedium string `json:"archiveMedium,omitempty"`

	// conversation Id
	ConversationID string `json:"conversationId,omitempty"`

	// The date the recording will be deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DeleteDate strfmt.DateTime `json:"deleteDate,omitempty"`

	// end time
	EndTime string `json:"endTime,omitempty"`

	// The date the recording will be exported. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ExportDate strfmt.DateTime `json:"exportDate,omitempty"`

	// The date the recording was exported. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ExportedDate strfmt.DateTime `json:"exportedDate,omitempty"`

	// Represents the current file state for a recording. Examples: Uploading, Archived, etc
	// Enum: [ARCHIVED AVAILABLE DELETED RESTORED RESTORING UPLOADING ERROR]
	FileState string `json:"fileState,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// How many archive restorations the organization is allowed to have.
	MaxAllowedRestorationsForOrg int32 `json:"maxAllowedRestorationsForOrg,omitempty"`

	// The type of media that the recording is. At the moment that could be audio, chat, email, or message.
	Media string `json:"media,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// The remaining archive restorations the organization has.
	RemainingRestorationsAllowedForOrg int32 `json:"remainingRestorationsAllowedForOrg,omitempty"`

	// The amount of time a restored recording will remain restored before being archived again. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	RestoreExpirationTime strfmt.DateTime `json:"restoreExpirationTime,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The session id represents an external resource id, such as email, call, chat, etc
	SessionID string `json:"sessionId,omitempty"`

	// The start time of the recording for screen recordings. Null for other types.
	StartTime string `json:"startTime,omitempty"`
}

// Validate validates this recording metadata
func (m *RecordingMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchiveMedium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreExpirationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecordingMetadata) validateAnnotations(formats strfmt.Registry) error {

	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	for i := 0; i < len(m.Annotations); i++ {
		if swag.IsZero(m.Annotations[i]) { // not required
			continue
		}

		if m.Annotations[i] != nil {
			if err := m.Annotations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("annotations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecordingMetadata) validateArchiveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ArchiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("archiveDate", "body", "date-time", m.ArchiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var recordingMetadataTypeArchiveMediumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLOUDARCHIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recordingMetadataTypeArchiveMediumPropEnum = append(recordingMetadataTypeArchiveMediumPropEnum, v)
	}
}

const (

	// RecordingMetadataArchiveMediumCLOUDARCHIVE captures enum value "CLOUDARCHIVE"
	RecordingMetadataArchiveMediumCLOUDARCHIVE string = "CLOUDARCHIVE"
)

// prop value enum
func (m *RecordingMetadata) validateArchiveMediumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recordingMetadataTypeArchiveMediumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecordingMetadata) validateArchiveMedium(formats strfmt.Registry) error {

	if swag.IsZero(m.ArchiveMedium) { // not required
		return nil
	}

	// value enum
	if err := m.validateArchiveMediumEnum("archiveMedium", "body", m.ArchiveMedium); err != nil {
		return err
	}

	return nil
}

func (m *RecordingMetadata) validateDeleteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeleteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("deleteDate", "body", "date-time", m.DeleteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecordingMetadata) validateExportDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExportDate) { // not required
		return nil
	}

	if err := validate.FormatOf("exportDate", "body", "date-time", m.ExportDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecordingMetadata) validateExportedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExportedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("exportedDate", "body", "date-time", m.ExportedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var recordingMetadataTypeFileStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ARCHIVED","AVAILABLE","DELETED","RESTORED","RESTORING","UPLOADING","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recordingMetadataTypeFileStatePropEnum = append(recordingMetadataTypeFileStatePropEnum, v)
	}
}

const (

	// RecordingMetadataFileStateARCHIVED captures enum value "ARCHIVED"
	RecordingMetadataFileStateARCHIVED string = "ARCHIVED"

	// RecordingMetadataFileStateAVAILABLE captures enum value "AVAILABLE"
	RecordingMetadataFileStateAVAILABLE string = "AVAILABLE"

	// RecordingMetadataFileStateDELETED captures enum value "DELETED"
	RecordingMetadataFileStateDELETED string = "DELETED"

	// RecordingMetadataFileStateRESTORED captures enum value "RESTORED"
	RecordingMetadataFileStateRESTORED string = "RESTORED"

	// RecordingMetadataFileStateRESTORING captures enum value "RESTORING"
	RecordingMetadataFileStateRESTORING string = "RESTORING"

	// RecordingMetadataFileStateUPLOADING captures enum value "UPLOADING"
	RecordingMetadataFileStateUPLOADING string = "UPLOADING"

	// RecordingMetadataFileStateERROR captures enum value "ERROR"
	RecordingMetadataFileStateERROR string = "ERROR"
)

// prop value enum
func (m *RecordingMetadata) validateFileStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recordingMetadataTypeFileStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecordingMetadata) validateFileState(formats strfmt.Registry) error {

	if swag.IsZero(m.FileState) { // not required
		return nil
	}

	// value enum
	if err := m.validateFileStateEnum("fileState", "body", m.FileState); err != nil {
		return err
	}

	return nil
}

func (m *RecordingMetadata) validateRestoreExpirationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RestoreExpirationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("restoreExpirationTime", "body", "date-time", m.RestoreExpirationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecordingMetadata) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecordingMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecordingMetadata) UnmarshalBinary(b []byte) error {
	var res RecordingMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
