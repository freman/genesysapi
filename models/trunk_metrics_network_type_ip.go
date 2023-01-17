// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrunkMetricsNetworkTypeIP trunk metrics network type Ip
//
// swagger:model TrunkMetricsNetworkTypeIp
type TrunkMetricsNetworkTypeIP struct {

	// Assigned IP Address for the interface
	Address string `json:"address,omitempty"`

	// Information about the error.
	ErrorInfo *TrunkErrorInfo `json:"errorInfo,omitempty"`
}

// Validate validates this trunk metrics network type Ip
func (m *TrunkMetricsNetworkTypeIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrunkMetricsNetworkTypeIP) validateErrorInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorInfo) { // not required
		return nil
	}

	if m.ErrorInfo != nil {
		if err := m.ErrorInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this trunk metrics network type Ip based on the context it is used
func (m *TrunkMetricsNetworkTypeIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrorInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrunkMetricsNetworkTypeIP) contextValidateErrorInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorInfo != nil {
		if err := m.ErrorInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrunkMetricsNetworkTypeIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrunkMetricsNetworkTypeIP) UnmarshalBinary(b []byte) error {
	var res TrunkMetricsNetworkTypeIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
