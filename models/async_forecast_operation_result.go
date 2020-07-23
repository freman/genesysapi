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

// AsyncForecastOperationResult async forecast operation result
//
// swagger:model AsyncForecastOperationResult
type AsyncForecastOperationResult struct {

	// The ID for the operation
	OperationID string `json:"operationId,omitempty"`

	// Percent progress for the operation
	Progress int32 `json:"progress,omitempty"`

	// The result of the operation.  Null unless status == Complete
	Result *BuShortTermForecast `json:"result,omitempty"`

	// The status of the operation
	// Enum: [Processing Complete Canceled Error]
	Status string `json:"status,omitempty"`
}

// Validate validates this async forecast operation result
func (m *AsyncForecastOperationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
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

func (m *AsyncForecastOperationResult) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

var asyncForecastOperationResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Processing","Complete","Canceled","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		asyncForecastOperationResultTypeStatusPropEnum = append(asyncForecastOperationResultTypeStatusPropEnum, v)
	}
}

const (

	// AsyncForecastOperationResultStatusProcessing captures enum value "Processing"
	AsyncForecastOperationResultStatusProcessing string = "Processing"

	// AsyncForecastOperationResultStatusComplete captures enum value "Complete"
	AsyncForecastOperationResultStatusComplete string = "Complete"

	// AsyncForecastOperationResultStatusCanceled captures enum value "Canceled"
	AsyncForecastOperationResultStatusCanceled string = "Canceled"

	// AsyncForecastOperationResultStatusError captures enum value "Error"
	AsyncForecastOperationResultStatusError string = "Error"
)

// prop value enum
func (m *AsyncForecastOperationResult) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, asyncForecastOperationResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AsyncForecastOperationResult) validateStatus(formats strfmt.Registry) error {

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
func (m *AsyncForecastOperationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AsyncForecastOperationResult) UnmarshalBinary(b []byte) error {
	var res AsyncForecastOperationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
