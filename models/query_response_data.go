// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QueryResponseData query response data
//
// swagger:model QueryResponseData
type QueryResponseData struct {

	// Interval with start and end represented as ISO-8601 string. i.e: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'/yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	Interval string `json:"interval,omitempty"`

	// A list of aggregated metrics
	Metrics []*QueryResponseMetric `json:"metrics"`
}

// Validate validates this query response data
func (m *QueryResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseData) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {
		if swag.IsZero(m.Metrics[i]) { // not required
			continue
		}

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this query response data based on the context it is used
func (m *QueryResponseData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseData) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metrics); i++ {

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseData) UnmarshalBinary(b []byte) error {
	var res QueryResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
