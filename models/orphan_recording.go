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

// OrphanRecording orphan recording
//
// swagger:model OrphanRecording
type OrphanRecording struct {

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"createdTime,omitempty"`

	// file state
	// Enum: [ARCHIVED AVAILABLE DELETED RESTORED RESTORING UPLOADING]
	FileState string `json:"fileState,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// media size bytes
	MediaSizeBytes int64 `json:"mediaSizeBytes,omitempty"`

	// media type
	// Enum: [CALL CHAT EMAIL SCREEN]
	MediaType string `json:"mediaType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The status of the orphaned recording's conversation.
	// Enum: [NO_CONVERSATION UNKNOWN_CONVERSATION CONVERSATION_NOT_COMPLETE CONVERSATION_NOT_EVALUATED EVALUATED]
	OrphanStatus string `json:"orphanStatus,omitempty"`

	// provider endpoint
	ProviderEndpoint *Endpoint `json:"providerEndpoint,omitempty"`

	// provider type
	// Enum: [EDGE CHAT EMAIL SCREEN_RECORDING]
	ProviderType string `json:"providerType,omitempty"`

	// recording
	Recording *Recording `json:"recording,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	RecoveredTime strfmt.DateTime `json:"recoveredTime,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// An identifier used during recovery operations by the supplying hybrid platform to track back and determine which interaction this recording is associated with
	SourceOrphaningID string `json:"sourceOrphaningId,omitempty"`
}

// Validate validates this orphan recording
func (m *OrphanRecording) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrphanStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecording(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveredTime(formats); err != nil {
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

func (m *OrphanRecording) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createdTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var orphanRecordingTypeFileStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ARCHIVED","AVAILABLE","DELETED","RESTORED","RESTORING","UPLOADING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orphanRecordingTypeFileStatePropEnum = append(orphanRecordingTypeFileStatePropEnum, v)
	}
}

const (

	// OrphanRecordingFileStateARCHIVED captures enum value "ARCHIVED"
	OrphanRecordingFileStateARCHIVED string = "ARCHIVED"

	// OrphanRecordingFileStateAVAILABLE captures enum value "AVAILABLE"
	OrphanRecordingFileStateAVAILABLE string = "AVAILABLE"

	// OrphanRecordingFileStateDELETED captures enum value "DELETED"
	OrphanRecordingFileStateDELETED string = "DELETED"

	// OrphanRecordingFileStateRESTORED captures enum value "RESTORED"
	OrphanRecordingFileStateRESTORED string = "RESTORED"

	// OrphanRecordingFileStateRESTORING captures enum value "RESTORING"
	OrphanRecordingFileStateRESTORING string = "RESTORING"

	// OrphanRecordingFileStateUPLOADING captures enum value "UPLOADING"
	OrphanRecordingFileStateUPLOADING string = "UPLOADING"
)

// prop value enum
func (m *OrphanRecording) validateFileStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orphanRecordingTypeFileStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrphanRecording) validateFileState(formats strfmt.Registry) error {

	if swag.IsZero(m.FileState) { // not required
		return nil
	}

	// value enum
	if err := m.validateFileStateEnum("fileState", "body", m.FileState); err != nil {
		return err
	}

	return nil
}

var orphanRecordingTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CALL","CHAT","EMAIL","SCREEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orphanRecordingTypeMediaTypePropEnum = append(orphanRecordingTypeMediaTypePropEnum, v)
	}
}

const (

	// OrphanRecordingMediaTypeCALL captures enum value "CALL"
	OrphanRecordingMediaTypeCALL string = "CALL"

	// OrphanRecordingMediaTypeCHAT captures enum value "CHAT"
	OrphanRecordingMediaTypeCHAT string = "CHAT"

	// OrphanRecordingMediaTypeEMAIL captures enum value "EMAIL"
	OrphanRecordingMediaTypeEMAIL string = "EMAIL"

	// OrphanRecordingMediaTypeSCREEN captures enum value "SCREEN"
	OrphanRecordingMediaTypeSCREEN string = "SCREEN"
)

// prop value enum
func (m *OrphanRecording) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orphanRecordingTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrphanRecording) validateMediaType(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

var orphanRecordingTypeOrphanStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NO_CONVERSATION","UNKNOWN_CONVERSATION","CONVERSATION_NOT_COMPLETE","CONVERSATION_NOT_EVALUATED","EVALUATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orphanRecordingTypeOrphanStatusPropEnum = append(orphanRecordingTypeOrphanStatusPropEnum, v)
	}
}

const (

	// OrphanRecordingOrphanStatusNOCONVERSATION captures enum value "NO_CONVERSATION"
	OrphanRecordingOrphanStatusNOCONVERSATION string = "NO_CONVERSATION"

	// OrphanRecordingOrphanStatusUNKNOWNCONVERSATION captures enum value "UNKNOWN_CONVERSATION"
	OrphanRecordingOrphanStatusUNKNOWNCONVERSATION string = "UNKNOWN_CONVERSATION"

	// OrphanRecordingOrphanStatusCONVERSATIONNOTCOMPLETE captures enum value "CONVERSATION_NOT_COMPLETE"
	OrphanRecordingOrphanStatusCONVERSATIONNOTCOMPLETE string = "CONVERSATION_NOT_COMPLETE"

	// OrphanRecordingOrphanStatusCONVERSATIONNOTEVALUATED captures enum value "CONVERSATION_NOT_EVALUATED"
	OrphanRecordingOrphanStatusCONVERSATIONNOTEVALUATED string = "CONVERSATION_NOT_EVALUATED"

	// OrphanRecordingOrphanStatusEVALUATED captures enum value "EVALUATED"
	OrphanRecordingOrphanStatusEVALUATED string = "EVALUATED"
)

// prop value enum
func (m *OrphanRecording) validateOrphanStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orphanRecordingTypeOrphanStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrphanRecording) validateOrphanStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.OrphanStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrphanStatusEnum("orphanStatus", "body", m.OrphanStatus); err != nil {
		return err
	}

	return nil
}

func (m *OrphanRecording) validateProviderEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderEndpoint) { // not required
		return nil
	}

	if m.ProviderEndpoint != nil {
		if err := m.ProviderEndpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("providerEndpoint")
			}
			return err
		}
	}

	return nil
}

var orphanRecordingTypeProviderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EDGE","CHAT","EMAIL","SCREEN_RECORDING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orphanRecordingTypeProviderTypePropEnum = append(orphanRecordingTypeProviderTypePropEnum, v)
	}
}

const (

	// OrphanRecordingProviderTypeEDGE captures enum value "EDGE"
	OrphanRecordingProviderTypeEDGE string = "EDGE"

	// OrphanRecordingProviderTypeCHAT captures enum value "CHAT"
	OrphanRecordingProviderTypeCHAT string = "CHAT"

	// OrphanRecordingProviderTypeEMAIL captures enum value "EMAIL"
	OrphanRecordingProviderTypeEMAIL string = "EMAIL"

	// OrphanRecordingProviderTypeSCREENRECORDING captures enum value "SCREEN_RECORDING"
	OrphanRecordingProviderTypeSCREENRECORDING string = "SCREEN_RECORDING"
)

// prop value enum
func (m *OrphanRecording) validateProviderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orphanRecordingTypeProviderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrphanRecording) validateProviderType(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderTypeEnum("providerType", "body", m.ProviderType); err != nil {
		return err
	}

	return nil
}

func (m *OrphanRecording) validateRecording(formats strfmt.Registry) error {

	if swag.IsZero(m.Recording) { // not required
		return nil
	}

	if m.Recording != nil {
		if err := m.Recording.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recording")
			}
			return err
		}
	}

	return nil
}

func (m *OrphanRecording) validateRecoveredTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RecoveredTime) { // not required
		return nil
	}

	if err := validate.FormatOf("recoveredTime", "body", "date-time", m.RecoveredTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrphanRecording) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrphanRecording) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrphanRecording) UnmarshalBinary(b []byte) error {
	var res OrphanRecording
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
