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

// AsyncQueryStatus async query status
//
// swagger:model AsyncQueryStatus
type AsyncQueryStatus struct {

	// The time at which the query completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	CompletionDate strfmt.DateTime `json:"completionDate,omitempty"`

	// The error associated with the current query, if the state is FAILED
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The time at which results for this query will expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

	// The current state of the asynchronous query
	// Enum: [QUEUED PENDING FAILED CANCELLED FULFILLED EXPIRED]
	State string `json:"state,omitempty"`

	// The time at which the query was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	SubmissionDate strfmt.DateTime `json:"submissionDate,omitempty"`
}

// Validate validates this async query status
func (m *AsyncQueryStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AsyncQueryStatus) validateCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completionDate", "body", "date-time", m.CompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AsyncQueryStatus) validateExpirationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var asyncQueryStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["QUEUED","PENDING","FAILED","CANCELLED","FULFILLED","EXPIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		asyncQueryStatusTypeStatePropEnum = append(asyncQueryStatusTypeStatePropEnum, v)
	}
}

const (

	// AsyncQueryStatusStateQUEUED captures enum value "QUEUED"
	AsyncQueryStatusStateQUEUED string = "QUEUED"

	// AsyncQueryStatusStatePENDING captures enum value "PENDING"
	AsyncQueryStatusStatePENDING string = "PENDING"

	// AsyncQueryStatusStateFAILED captures enum value "FAILED"
	AsyncQueryStatusStateFAILED string = "FAILED"

	// AsyncQueryStatusStateCANCELLED captures enum value "CANCELLED"
	AsyncQueryStatusStateCANCELLED string = "CANCELLED"

	// AsyncQueryStatusStateFULFILLED captures enum value "FULFILLED"
	AsyncQueryStatusStateFULFILLED string = "FULFILLED"

	// AsyncQueryStatusStateEXPIRED captures enum value "EXPIRED"
	AsyncQueryStatusStateEXPIRED string = "EXPIRED"
)

// prop value enum
func (m *AsyncQueryStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, asyncQueryStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AsyncQueryStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *AsyncQueryStatus) validateSubmissionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("submissionDate", "body", "date-time", m.SubmissionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AsyncQueryStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AsyncQueryStatus) UnmarshalBinary(b []byte) error {
	var res AsyncQueryStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
