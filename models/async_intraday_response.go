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

// AsyncIntradayResponse async intraday response
//
// swagger:model AsyncIntradayResponse
type AsyncIntradayResponse struct {

	// The ID for the operation
	OperationID string `json:"operationId,omitempty"`

	// The result of the operation.  Null unless status == Complete
	Result *BuIntradayResponse `json:"result,omitempty"`

	// The status of the operation
	// Enum: [Processing Complete Canceled Error]
	Status string `json:"status,omitempty"`
}

// Validate validates this async intraday response
func (m *AsyncIntradayResponse) Validate(formats strfmt.Registry) error {
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

func (m *AsyncIntradayResponse) validateResult(formats strfmt.Registry) error {

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

var asyncIntradayResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Processing","Complete","Canceled","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		asyncIntradayResponseTypeStatusPropEnum = append(asyncIntradayResponseTypeStatusPropEnum, v)
	}
}

const (

	// AsyncIntradayResponseStatusProcessing captures enum value "Processing"
	AsyncIntradayResponseStatusProcessing string = "Processing"

	// AsyncIntradayResponseStatusComplete captures enum value "Complete"
	AsyncIntradayResponseStatusComplete string = "Complete"

	// AsyncIntradayResponseStatusCanceled captures enum value "Canceled"
	AsyncIntradayResponseStatusCanceled string = "Canceled"

	// AsyncIntradayResponseStatusError captures enum value "Error"
	AsyncIntradayResponseStatusError string = "Error"
)

// prop value enum
func (m *AsyncIntradayResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, asyncIntradayResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AsyncIntradayResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *AsyncIntradayResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AsyncIntradayResponse) UnmarshalBinary(b []byte) error {
	var res AsyncIntradayResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
