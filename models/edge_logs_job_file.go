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

// EdgeLogsJobFile edge logs job file
//
// swagger:model EdgeLogsJobFile
type EdgeLogsJobFile struct {

	// The ID of the user that created the resource.
	CreatedBy string `json:"createdBy,omitempty"`

	// The application that created the resource.
	CreatedByApp string `json:"createdByApp,omitempty"`

	// The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The resource's description.
	Description string `json:"description,omitempty"`

	// The download ID to use with the downloads API.
	DownloadID string `json:"downloadId,omitempty"`

	// The path of this file on the Edge.
	// Format: uri
	EdgePath strfmt.URI `json:"edgePath,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The ID of the user that last modified the resource.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The application that last modified the resource.
	ModifiedByApp string `json:"modifiedByApp,omitempty"`

	// The name of the entity.
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The size of this file in bytes.
	SizeBytes float64 `json:"sizeBytes,omitempty"`

	// Indicates if the resource is active, inactive, or deleted.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`

	// The time this log file was last modified on the Edge. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	TimeModified strfmt.DateTime `json:"timeModified,omitempty"`

	// The status of the upload of this file from the Edge to the cloud.  Use /upload to start an upload.
	// Enum: [UPLOADING NOT_UPLOADED UPLOADED ERROR_ON_UPLOAD]
	UploadStatus string `json:"uploadStatus,omitempty"`

	// The current version of the resource.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this edge logs job file
func (m *EdgeLogsJobFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeModified(formats); err != nil {
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

func (m *EdgeLogsJobFile) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLogsJobFile) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLogsJobFile) validateEdgePath(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgePath) { // not required
		return nil
	}

	if err := validate.FormatOf("edgePath", "body", "uri", m.EdgePath.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLogsJobFile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLogsJobFile) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var edgeLogsJobFileTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgeLogsJobFileTypeStatePropEnum = append(edgeLogsJobFileTypeStatePropEnum, v)
	}
}

const (

	// EdgeLogsJobFileStateActive captures enum value "active"
	EdgeLogsJobFileStateActive string = "active"

	// EdgeLogsJobFileStateInactive captures enum value "inactive"
	EdgeLogsJobFileStateInactive string = "inactive"

	// EdgeLogsJobFileStateDeleted captures enum value "deleted"
	EdgeLogsJobFileStateDeleted string = "deleted"
)

// prop value enum
func (m *EdgeLogsJobFile) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, edgeLogsJobFileTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgeLogsJobFile) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLogsJobFile) validateTimeCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLogsJobFile) validateTimeModified(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeModified) { // not required
		return nil
	}

	if err := validate.FormatOf("timeModified", "body", "date-time", m.TimeModified.String(), formats); err != nil {
		return err
	}

	return nil
}

var edgeLogsJobFileTypeUploadStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UPLOADING","NOT_UPLOADED","UPLOADED","ERROR_ON_UPLOAD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgeLogsJobFileTypeUploadStatusPropEnum = append(edgeLogsJobFileTypeUploadStatusPropEnum, v)
	}
}

const (

	// EdgeLogsJobFileUploadStatusUPLOADING captures enum value "UPLOADING"
	EdgeLogsJobFileUploadStatusUPLOADING string = "UPLOADING"

	// EdgeLogsJobFileUploadStatusNOTUPLOADED captures enum value "NOT_UPLOADED"
	EdgeLogsJobFileUploadStatusNOTUPLOADED string = "NOT_UPLOADED"

	// EdgeLogsJobFileUploadStatusUPLOADED captures enum value "UPLOADED"
	EdgeLogsJobFileUploadStatusUPLOADED string = "UPLOADED"

	// EdgeLogsJobFileUploadStatusERRORONUPLOAD captures enum value "ERROR_ON_UPLOAD"
	EdgeLogsJobFileUploadStatusERRORONUPLOAD string = "ERROR_ON_UPLOAD"
)

// prop value enum
func (m *EdgeLogsJobFile) validateUploadStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, edgeLogsJobFileTypeUploadStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgeLogsJobFile) validateUploadStatus(formats strfmt.Registry) error {

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
func (m *EdgeLogsJobFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeLogsJobFile) UnmarshalBinary(b []byte) error {
	var res EdgeLogsJobFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}