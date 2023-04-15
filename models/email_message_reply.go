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

// EmailMessageReply email message reply
//
// swagger:model EmailMessageReply
type EmailMessageReply struct {

	// The attachments of the email message.
	Attachments []*Attachment `json:"attachments"`

	// The recipients that were blind copied on the email message.
	Bcc []*EmailAddress `json:"bcc"`

	// The recipients that were copied on the email message.
	Cc []*EmailAddress `json:"cc"`

	// Indicates an estimation of the size of the current email as a whole, in its final, ready to be sent form.
	// Read Only: true
	EmailSizeBytes int32 `json:"emailSizeBytes,omitempty"`

	// The sender of the email message.
	// Required: true
	From *EmailAddress `json:"from"`

	// Indicates whether the history of previous emails of the conversation is included within the email bodies of this message.
	HistoryIncluded bool `json:"historyIncluded"`

	// The html body of the email message.
	HTMLBody string `json:"htmlBody,omitempty"`

	// Indicates the maximum allowed size for an email to be send via SMTP server, based on the email domain configuration
	// Read Only: true
	MaxEmailSizeBytes int32 `json:"maxEmailSizeBytes,omitempty"`

	// The receiver of the reply email message.
	ReplyTo *EmailAddress `json:"replyTo,omitempty"`

	// The subject of the email message.
	Subject string `json:"subject,omitempty"`

	// The text body of the email message.
	// Required: true
	TextBody *string `json:"textBody"`

	// The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// The recipients of the email message.
	// Required: true
	To []*EmailAddress `json:"to"`
}

// Validate validates this email message reply
func (m *EmailMessageReply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplyTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailMessageReply) validateAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.Attachments) { // not required
		return nil
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

func (m *EmailMessageReply) validateBcc(formats strfmt.Registry) error {
	if swag.IsZero(m.Bcc) { // not required
		return nil
	}

	for i := 0; i < len(m.Bcc); i++ {
		if swag.IsZero(m.Bcc[i]) { // not required
			continue
		}

		if m.Bcc[i] != nil {
			if err := m.Bcc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bcc" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bcc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMessageReply) validateCc(formats strfmt.Registry) error {
	if swag.IsZero(m.Cc) { // not required
		return nil
	}

	for i := 0; i < len(m.Cc); i++ {
		if swag.IsZero(m.Cc[i]) { // not required
			continue
		}

		if m.Cc[i] != nil {
			if err := m.Cc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cc" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMessageReply) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *EmailMessageReply) validateReplyTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplyTo) { // not required
		return nil
	}

	if m.ReplyTo != nil {
		if err := m.ReplyTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replyTo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replyTo")
			}
			return err
		}
	}

	return nil
}

func (m *EmailMessageReply) validateTextBody(formats strfmt.Registry) error {

	if err := validate.Required("textBody", "body", m.TextBody); err != nil {
		return err
	}

	return nil
}

func (m *EmailMessageReply) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EmailMessageReply) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	for i := 0; i < len(m.To); i++ {
		if swag.IsZero(m.To[i]) { // not required
			continue
		}

		if m.To[i] != nil {
			if err := m.To[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("to" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this email message reply based on the context it is used
func (m *EmailMessageReply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBcc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmailSizeBytes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxEmailSizeBytes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplyTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailMessageReply) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmailMessageReply) contextValidateBcc(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Bcc); i++ {

		if m.Bcc[i] != nil {
			if err := m.Bcc[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bcc" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bcc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMessageReply) contextValidateCc(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cc); i++ {

		if m.Cc[i] != nil {
			if err := m.Cc[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cc" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMessageReply) contextValidateEmailSizeBytes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "emailSizeBytes", "body", int32(m.EmailSizeBytes)); err != nil {
		return err
	}

	return nil
}

func (m *EmailMessageReply) contextValidateFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.From != nil {
		if err := m.From.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *EmailMessageReply) contextValidateMaxEmailSizeBytes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "maxEmailSizeBytes", "body", int32(m.MaxEmailSizeBytes)); err != nil {
		return err
	}

	return nil
}

func (m *EmailMessageReply) contextValidateReplyTo(ctx context.Context, formats strfmt.Registry) error {

	if m.ReplyTo != nil {
		if err := m.ReplyTo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replyTo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replyTo")
			}
			return err
		}
	}

	return nil
}

func (m *EmailMessageReply) contextValidateTo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.To); i++ {

		if m.To[i] != nil {
			if err := m.To[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("to" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailMessageReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailMessageReply) UnmarshalBinary(b []byte) error {
	var res EmailMessageReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}