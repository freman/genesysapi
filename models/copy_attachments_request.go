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

// CopyAttachmentsRequest copy attachments request
//
// swagger:model CopyAttachmentsRequest
type CopyAttachmentsRequest struct {

	// A list of attachments that will be copied from the source message to the current draft
	// Required: true
	Attachments []*Attachment `json:"attachments"`

	// A reference to the email message within the current conversation that owns the attachments to be copied
	// Required: true
	SourceMessage *DomainEntityRef `json:"sourceMessage"`
}

// Validate validates this copy attachments request
func (m *CopyAttachmentsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopyAttachmentsRequest) validateAttachments(formats strfmt.Registry) error {

	if err := validate.Required("attachments", "body", m.Attachments); err != nil {
		return err
	}

	for i := 0; i < len(m.Attachments); i++ {
		if swag.IsZero(m.Attachments[i]) { // not required
			continue
		}

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CopyAttachmentsRequest) validateSourceMessage(formats strfmt.Registry) error {

	if err := validate.Required("sourceMessage", "body", m.SourceMessage); err != nil {
		return err
	}

	if m.SourceMessage != nil {
		if err := m.SourceMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceMessage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this copy attachments request based on the context it is used
func (m *CopyAttachmentsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopyAttachmentsRequest) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attachments); i++ {

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CopyAttachmentsRequest) contextValidateSourceMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceMessage != nil {
		if err := m.SourceMessage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceMessage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CopyAttachmentsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyAttachmentsRequest) UnmarshalBinary(b []byte) error {
	var res CopyAttachmentsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
