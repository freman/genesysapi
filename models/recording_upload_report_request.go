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

// RecordingUploadReportRequest recording upload report request
//
// swagger:model RecordingUploadReportRequest
type RecordingUploadReportRequest struct {

	// Report will include uploads since this date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	DateSince *strfmt.DateTime `json:"dateSince"`

	// Report will include uploads with this status
	// Enum: [Pending Success Failure]
	UploadStatus string `json:"uploadStatus,omitempty"`
}

// Validate validates this recording upload report request
func (m *RecordingUploadReportRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateSince(formats); err != nil {
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

func (m *RecordingUploadReportRequest) validateDateSince(formats strfmt.Registry) error {

	if err := validate.Required("dateSince", "body", m.DateSince); err != nil {
		return err
	}

	if err := validate.FormatOf("dateSince", "body", "date-time", m.DateSince.String(), formats); err != nil {
		return err
	}

	return nil
}

var recordingUploadReportRequestTypeUploadStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Success","Failure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recordingUploadReportRequestTypeUploadStatusPropEnum = append(recordingUploadReportRequestTypeUploadStatusPropEnum, v)
	}
}

const (

	// RecordingUploadReportRequestUploadStatusPending captures enum value "Pending"
	RecordingUploadReportRequestUploadStatusPending string = "Pending"

	// RecordingUploadReportRequestUploadStatusSuccess captures enum value "Success"
	RecordingUploadReportRequestUploadStatusSuccess string = "Success"

	// RecordingUploadReportRequestUploadStatusFailure captures enum value "Failure"
	RecordingUploadReportRequestUploadStatusFailure string = "Failure"
)

// prop value enum
func (m *RecordingUploadReportRequest) validateUploadStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recordingUploadReportRequestTypeUploadStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecordingUploadReportRequest) validateUploadStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateUploadStatusEnum("uploadStatus", "body", m.UploadStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recording upload report request based on context it is used
func (m *RecordingUploadReportRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecordingUploadReportRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecordingUploadReportRequest) UnmarshalBinary(b []byte) error {
	var res RecordingUploadReportRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}