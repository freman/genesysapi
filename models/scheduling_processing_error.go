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

// SchedulingProcessingError scheduling processing error
//
// swagger:model SchedulingProcessingError
type SchedulingProcessingError struct {

	// A text description of the error
	// Read Only: true
	Description string `json:"description,omitempty"`

	// An internal code representing the type of error. BadJson for 'Unable to parse json.' NotFound for 'Resource not found.' Fail for 'An unexpected server error occured.'
	// Read Only: true
	// Enum: [BadJson NotFound Fail]
	InternalErrorCode string `json:"internalErrorCode,omitempty"`
}

// Validate validates this scheduling processing error
func (m *SchedulingProcessingError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInternalErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var schedulingProcessingErrorTypeInternalErrorCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BadJson","NotFound","Fail"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schedulingProcessingErrorTypeInternalErrorCodePropEnum = append(schedulingProcessingErrorTypeInternalErrorCodePropEnum, v)
	}
}

const (

	// SchedulingProcessingErrorInternalErrorCodeBadJSON captures enum value "BadJson"
	SchedulingProcessingErrorInternalErrorCodeBadJSON string = "BadJson"

	// SchedulingProcessingErrorInternalErrorCodeNotFound captures enum value "NotFound"
	SchedulingProcessingErrorInternalErrorCodeNotFound string = "NotFound"

	// SchedulingProcessingErrorInternalErrorCodeFail captures enum value "Fail"
	SchedulingProcessingErrorInternalErrorCodeFail string = "Fail"
)

// prop value enum
func (m *SchedulingProcessingError) validateInternalErrorCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, schedulingProcessingErrorTypeInternalErrorCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SchedulingProcessingError) validateInternalErrorCode(formats strfmt.Registry) error {

	if swag.IsZero(m.InternalErrorCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateInternalErrorCodeEnum("internalErrorCode", "body", m.InternalErrorCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulingProcessingError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulingProcessingError) UnmarshalBinary(b []byte) error {
	var res SchedulingProcessingError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
