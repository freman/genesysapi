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

// BulkUpdateShiftTradeStateResponse bulk update shift trade state response
//
// swagger:model BulkUpdateShiftTradeStateResponse
type BulkUpdateShiftTradeStateResponse struct {

	// The ID for the operation
	OperationID string `json:"operationId,omitempty"`

	// The result of the operation.  Null unless status == Complete
	Result *BulkUpdateShiftTradeStateResult `json:"result,omitempty"`

	// The status of the operation
	// Enum: [Processing Complete Canceled Error]
	Status string `json:"status,omitempty"`
}

// Validate validates this bulk update shift trade state response
func (m *BulkUpdateShiftTradeStateResponse) Validate(formats strfmt.Registry) error {
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

func (m *BulkUpdateShiftTradeStateResponse) validateResult(formats strfmt.Registry) error {
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

var bulkUpdateShiftTradeStateResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Processing","Complete","Canceled","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bulkUpdateShiftTradeStateResponseTypeStatusPropEnum = append(bulkUpdateShiftTradeStateResponseTypeStatusPropEnum, v)
	}
}

const (

	// BulkUpdateShiftTradeStateResponseStatusProcessing captures enum value "Processing"
	BulkUpdateShiftTradeStateResponseStatusProcessing string = "Processing"

	// BulkUpdateShiftTradeStateResponseStatusComplete captures enum value "Complete"
	BulkUpdateShiftTradeStateResponseStatusComplete string = "Complete"

	// BulkUpdateShiftTradeStateResponseStatusCanceled captures enum value "Canceled"
	BulkUpdateShiftTradeStateResponseStatusCanceled string = "Canceled"

	// BulkUpdateShiftTradeStateResponseStatusError captures enum value "Error"
	BulkUpdateShiftTradeStateResponseStatusError string = "Error"
)

// prop value enum
func (m *BulkUpdateShiftTradeStateResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bulkUpdateShiftTradeStateResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BulkUpdateShiftTradeStateResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bulk update shift trade state response based on the context it is used
func (m *BulkUpdateShiftTradeStateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkUpdateShiftTradeStateResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BulkUpdateShiftTradeStateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkUpdateShiftTradeStateResponse) UnmarshalBinary(b []byte) error {
	var res BulkUpdateShiftTradeStateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
