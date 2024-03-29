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

// NluDetectionResponse nlu detection response
//
// swagger:model NluDetectionResponse
type NluDetectionResponse struct {

	// input
	Input *NluDetectionInput `json:"input,omitempty"`

	// output
	Output *NluDetectionOutput `json:"output,omitempty"`

	// The NLU domain version which performed the detection.
	// Read Only: true
	Version *NluDomainVersion `json:"version,omitempty"`
}

// Validate validates this nlu detection response
func (m *NluDetectionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NluDetectionResponse) validateInput(formats strfmt.Registry) error {
	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

func (m *NluDetectionResponse) validateOutput(formats strfmt.Registry) error {
	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

func (m *NluDetectionResponse) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nlu detection response based on the context it is used
func (m *NluDetectionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NluDetectionResponse) contextValidateInput(ctx context.Context, formats strfmt.Registry) error {

	if m.Input != nil {
		if err := m.Input.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

func (m *NluDetectionResponse) contextValidateOutput(ctx context.Context, formats strfmt.Registry) error {

	if m.Output != nil {
		if err := m.Output.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

func (m *NluDetectionResponse) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NluDetectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NluDetectionResponse) UnmarshalBinary(b []byte) error {
	var res NluDetectionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
