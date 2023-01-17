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
	"github.com/go-openapi/validate"
)

// ProgramMappings program mappings
//
// swagger:model ProgramMappings
type ProgramMappings struct {

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// flows
	Flows []*AddressableEntityRef `json:"flows"`

	// modified by
	ModifiedBy *AddressableEntityRef `json:"modifiedBy,omitempty"`

	// program
	Program *BaseProgramEntity `json:"program,omitempty"`

	// queues
	Queues []*AddressableEntityRef `json:"queues"`
}

// Validate validates this program mappings
func (m *ProgramMappings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlows(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgram(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProgramMappings) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProgramMappings) validateFlows(formats strfmt.Registry) error {
	if swag.IsZero(m.Flows) { // not required
		return nil
	}

	for i := 0; i < len(m.Flows); i++ {
		if swag.IsZero(m.Flows[i]) { // not required
			continue
		}

		if m.Flows[i] != nil {
			if err := m.Flows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProgramMappings) validateModifiedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ProgramMappings) validateProgram(formats strfmt.Registry) error {
	if swag.IsZero(m.Program) { // not required
		return nil
	}

	if m.Program != nil {
		if err := m.Program.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("program")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("program")
			}
			return err
		}
	}

	return nil
}

func (m *ProgramMappings) validateQueues(formats strfmt.Registry) error {
	if swag.IsZero(m.Queues) { // not required
		return nil
	}

	for i := 0; i < len(m.Queues); i++ {
		if swag.IsZero(m.Queues[i]) { // not required
			continue
		}

		if m.Queues[i] != nil {
			if err := m.Queues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this program mappings based on the context it is used
func (m *ProgramMappings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProgram(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProgramMappings) contextValidateFlows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Flows); i++ {

		if m.Flows[i] != nil {
			if err := m.Flows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProgramMappings) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ProgramMappings) contextValidateProgram(ctx context.Context, formats strfmt.Registry) error {

	if m.Program != nil {
		if err := m.Program.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("program")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("program")
			}
			return err
		}
	}

	return nil
}

func (m *ProgramMappings) contextValidateQueues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Queues); i++ {

		if m.Queues[i] != nil {
			if err := m.Queues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProgramMappings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgramMappings) UnmarshalBinary(b []byte) error {
	var res ProgramMappings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
