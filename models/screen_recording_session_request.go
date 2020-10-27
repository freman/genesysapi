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

// ScreenRecordingSessionRequest screen recording session request
//
// swagger:model ScreenRecordingSessionRequest
type ScreenRecordingSessionRequest struct {

	// The screen recording session's archive date. Must be greater than 1 day from now if set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ArchiveDate strfmt.DateTime `json:"archiveDate,omitempty"`

	// The screen recording session's delete date. Must be greater than archiveDate if set, otherwise one day from now. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DeleteDate strfmt.DateTime `json:"deleteDate,omitempty"`

	// The screen recording session's state.  Values can be: 'stopped'
	// Enum: [STOPPED]
	State string `json:"state,omitempty"`
}

// Validate validates this screen recording session request
func (m *ScreenRecordingSessionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScreenRecordingSessionRequest) validateArchiveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ArchiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("archiveDate", "body", "date-time", m.ArchiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ScreenRecordingSessionRequest) validateDeleteDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeleteDate) { // not required
		return nil
	}

	if err := validate.FormatOf("deleteDate", "body", "date-time", m.DeleteDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var screenRecordingSessionRequestTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STOPPED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		screenRecordingSessionRequestTypeStatePropEnum = append(screenRecordingSessionRequestTypeStatePropEnum, v)
	}
}

const (

	// ScreenRecordingSessionRequestStateSTOPPED captures enum value "STOPPED"
	ScreenRecordingSessionRequestStateSTOPPED string = "STOPPED"
)

// prop value enum
func (m *ScreenRecordingSessionRequest) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, screenRecordingSessionRequestTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScreenRecordingSessionRequest) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScreenRecordingSessionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScreenRecordingSessionRequest) UnmarshalBinary(b []byte) error {
	var res ScreenRecordingSessionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
