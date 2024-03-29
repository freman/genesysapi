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

// BuAsyncScheduleResponse bu async schedule response
//
// swagger:model BuAsyncScheduleResponse
type BuAsyncScheduleResponse struct {

	// The ID for the operation
	OperationID string `json:"operationId,omitempty"`

	// The result of the operation.  Null unless status == Complete
	Result *BuScheduleMetadata `json:"result,omitempty"`

	// The status of the operation
	// Enum: [Processing Complete Canceled Error]
	Status string `json:"status,omitempty"`
}

// Validate validates this bu async schedule response
func (m *BuAsyncScheduleResponse) Validate(formats strfmt.Registry) error {
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

func (m *BuAsyncScheduleResponse) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

var buAsyncScheduleResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Processing","Complete","Canceled","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buAsyncScheduleResponseTypeStatusPropEnum = append(buAsyncScheduleResponseTypeStatusPropEnum, v)
	}
}

const (

	// BuAsyncScheduleResponseStatusProcessing captures enum value "Processing"
	BuAsyncScheduleResponseStatusProcessing string = "Processing"

	// BuAsyncScheduleResponseStatusComplete captures enum value "Complete"
	BuAsyncScheduleResponseStatusComplete string = "Complete"

	// BuAsyncScheduleResponseStatusCanceled captures enum value "Canceled"
	BuAsyncScheduleResponseStatusCanceled string = "Canceled"

	// BuAsyncScheduleResponseStatusError captures enum value "Error"
	BuAsyncScheduleResponseStatusError string = "Error"
)

// prop value enum
func (m *BuAsyncScheduleResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buAsyncScheduleResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuAsyncScheduleResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bu async schedule response based on the context it is used
func (m *BuAsyncScheduleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAsyncScheduleResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {
		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuAsyncScheduleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuAsyncScheduleResponse) UnmarshalBinary(b []byte) error {
	var res BuAsyncScheduleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
