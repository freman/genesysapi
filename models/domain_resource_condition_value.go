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

// DomainResourceConditionValue domain resource condition value
//
// swagger:model DomainResourceConditionValue
type DomainResourceConditionValue struct {

	// queue
	Queue *Queue `json:"queue,omitempty"`

	// type
	// Enum: [SCALAR VARIABLE USER QUEUE]
	Type string `json:"type,omitempty"`

	// user
	User *User `json:"user,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this domain resource condition value
func (m *DomainResourceConditionValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainResourceConditionValue) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

var domainResourceConditionValueTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SCALAR","VARIABLE","USER","QUEUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainResourceConditionValueTypeTypePropEnum = append(domainResourceConditionValueTypeTypePropEnum, v)
	}
}

const (

	// DomainResourceConditionValueTypeSCALAR captures enum value "SCALAR"
	DomainResourceConditionValueTypeSCALAR string = "SCALAR"

	// DomainResourceConditionValueTypeVARIABLE captures enum value "VARIABLE"
	DomainResourceConditionValueTypeVARIABLE string = "VARIABLE"

	// DomainResourceConditionValueTypeUSER captures enum value "USER"
	DomainResourceConditionValueTypeUSER string = "USER"

	// DomainResourceConditionValueTypeQUEUE captures enum value "QUEUE"
	DomainResourceConditionValueTypeQUEUE string = "QUEUE"
)

// prop value enum
func (m *DomainResourceConditionValue) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, domainResourceConditionValueTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainResourceConditionValue) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DomainResourceConditionValue) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain resource condition value based on the context it is used
func (m *DomainResourceConditionValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainResourceConditionValue) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {
		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *DomainResourceConditionValue) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainResourceConditionValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainResourceConditionValue) UnmarshalBinary(b []byte) error {
	var res DomainResourceConditionValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
