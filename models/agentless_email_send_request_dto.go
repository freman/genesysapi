// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AgentlessEmailSendRequestDto agentless email send request dto
//
// swagger:model AgentlessEmailSendRequestDto
type AgentlessEmailSendRequestDto struct {

	// The identifier of the conversation.
	ConversationID string `json:"conversationId,omitempty"`

	// The sender of the message.
	// Required: true
	FromAddress *EmailAddress `json:"fromAddress"`

	// The Content of the message, in HTML. Links, images and styles are allowed
	HTMLBody string `json:"htmlBody,omitempty"`

	// The address to use for reply.
	ReplyToAddress *EmailAddress `json:"replyToAddress,omitempty"`

	// The direction of the message.
	// Required: true
	// Enum: [Outbound Inbound Integration]
	SenderType *string `json:"senderType"`

	// The subject of the message.
	Subject string `json:"subject,omitempty"`

	// The Content of the message, in plain text.
	TextBody string `json:"textBody,omitempty"`

	// The recipient(s) of the message.
	// Required: true
	ToAddresses []*EmailAddress `json:"toAddresses"`
}

// Validate validates this agentless email send request dto
func (m *AgentlessEmailSendRequestDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplyToAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToAddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentlessEmailSendRequestDto) validateFromAddress(formats strfmt.Registry) error {

	if err := validate.Required("fromAddress", "body", m.FromAddress); err != nil {
		return err
	}

	if m.FromAddress != nil {
		if err := m.FromAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fromAddress")
			}
			return err
		}
	}

	return nil
}

func (m *AgentlessEmailSendRequestDto) validateReplyToAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplyToAddress) { // not required
		return nil
	}

	if m.ReplyToAddress != nil {
		if err := m.ReplyToAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replyToAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replyToAddress")
			}
			return err
		}
	}

	return nil
}

var agentlessEmailSendRequestDtoTypeSenderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Outbound","Inbound","Integration"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		agentlessEmailSendRequestDtoTypeSenderTypePropEnum = append(agentlessEmailSendRequestDtoTypeSenderTypePropEnum, v)
	}
}

const (

	// AgentlessEmailSendRequestDtoSenderTypeOutbound captures enum value "Outbound"
	AgentlessEmailSendRequestDtoSenderTypeOutbound string = "Outbound"

	// AgentlessEmailSendRequestDtoSenderTypeInbound captures enum value "Inbound"
	AgentlessEmailSendRequestDtoSenderTypeInbound string = "Inbound"

	// AgentlessEmailSendRequestDtoSenderTypeIntegration captures enum value "Integration"
	AgentlessEmailSendRequestDtoSenderTypeIntegration string = "Integration"
)

// prop value enum
func (m *AgentlessEmailSendRequestDto) validateSenderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, agentlessEmailSendRequestDtoTypeSenderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AgentlessEmailSendRequestDto) validateSenderType(formats strfmt.Registry) error {

	if err := validate.Required("senderType", "body", m.SenderType); err != nil {
		return err
	}

	// value enum
	if err := m.validateSenderTypeEnum("senderType", "body", *m.SenderType); err != nil {
		return err
	}

	return nil
}

func (m *AgentlessEmailSendRequestDto) validateToAddresses(formats strfmt.Registry) error {

	if err := validate.Required("toAddresses", "body", m.ToAddresses); err != nil {
		return err
	}

	for i := 0; i < len(m.ToAddresses); i++ {
		if swag.IsZero(m.ToAddresses[i]) { // not required
			continue
		}

		if m.ToAddresses[i] != nil {
			if err := m.ToAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("toAddresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("toAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this agentless email send request dto based on the context it is used
func (m *AgentlessEmailSendRequestDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFromAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplyToAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentlessEmailSendRequestDto) contextValidateFromAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.FromAddress != nil {
		if err := m.FromAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fromAddress")
			}
			return err
		}
	}

	return nil
}

func (m *AgentlessEmailSendRequestDto) contextValidateReplyToAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.ReplyToAddress != nil {
		if err := m.ReplyToAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replyToAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replyToAddress")
			}
			return err
		}
	}

	return nil
}

func (m *AgentlessEmailSendRequestDto) contextValidateToAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ToAddresses); i++ {

		if m.ToAddresses[i] != nil {
			if err := m.ToAddresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("toAddresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("toAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentlessEmailSendRequestDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentlessEmailSendRequestDto) UnmarshalBinary(b []byte) error {
	var res AgentlessEmailSendRequestDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}