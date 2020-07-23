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

// ModelingStatusResponse modeling status response
//
// swagger:model ModelingStatusResponse
type ModelingStatusResponse struct {

	// If the request could not be properly processed, error details will be given here.
	// Read Only: true
	ErrorDetails []*ModelingProcessingError `json:"errorDetails"`

	// The ID generated for the modeling job.  Use to GET result when job is completed.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The uri of the modeling result. It has a value if the status is either 'Success', 'PartialFailure', or 'Failed'.
	// Read Only: true
	ModelingResultURI string `json:"modelingResultUri,omitempty"`

	// The status of the modeling job.
	// Read Only: true
	// Enum: [Pending Success Failed Ongoing PartialFailure]
	Status string `json:"status,omitempty"`
}

// Validate validates this modeling status response
func (m *ModelingStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelingStatusResponse) validateErrorDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ErrorDetails); i++ {
		if swag.IsZero(m.ErrorDetails[i]) { // not required
			continue
		}

		if m.ErrorDetails[i] != nil {
			if err := m.ErrorDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errorDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var modelingStatusResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Success","Failed","Ongoing","PartialFailure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelingStatusResponseTypeStatusPropEnum = append(modelingStatusResponseTypeStatusPropEnum, v)
	}
}

const (

	// ModelingStatusResponseStatusPending captures enum value "Pending"
	ModelingStatusResponseStatusPending string = "Pending"

	// ModelingStatusResponseStatusSuccess captures enum value "Success"
	ModelingStatusResponseStatusSuccess string = "Success"

	// ModelingStatusResponseStatusFailed captures enum value "Failed"
	ModelingStatusResponseStatusFailed string = "Failed"

	// ModelingStatusResponseStatusOngoing captures enum value "Ongoing"
	ModelingStatusResponseStatusOngoing string = "Ongoing"

	// ModelingStatusResponseStatusPartialFailure captures enum value "PartialFailure"
	ModelingStatusResponseStatusPartialFailure string = "PartialFailure"
)

// prop value enum
func (m *ModelingStatusResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modelingStatusResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModelingStatusResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelingStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelingStatusResponse) UnmarshalBinary(b []byte) error {
	var res ModelingStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
