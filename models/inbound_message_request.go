// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InboundMessageRequest inbound message request
//
// swagger:model InboundMessageRequest
type InboundMessageRequest struct {

	// The list of attributes to associate with the customer participant.
	Attributes map[string]string `json:"attributes,omitempty"`

	// The ID of the flow to use for routing email conversation. This field is mutually exclusive with queueId
	FlowID string `json:"flowId,omitempty"`

	// The email address of the sender of the email.
	FromAddress string `json:"fromAddress,omitempty"`

	// The name of the sender of the email.
	FromName string `json:"fromName,omitempty"`

	// The ID of the language to use for routing.
	LanguageID string `json:"languageId,omitempty"`

	// The priority to assign to the conversation for routing.
	Priority int32 `json:"priority,omitempty"`

	// The name of the provider that is sourcing the email such as Oracle, Salesforce, etc.
	// Required: true
	Provider *string `json:"provider"`

	// The ID of the queue to use for routing the email conversation. This field is mutually exclusive with flowId
	QueueID string `json:"queueId,omitempty"`

	// The list of skill ID's to use for routing.
	SkillIds []string `json:"skillIds"`

	// The subject of the email
	Subject string `json:"subject,omitempty"`

	// The email address of the recipient of the email.
	ToAddress string `json:"toAddress,omitempty"`

	// The name of the recipient of the email.
	ToName string `json:"toName,omitempty"`
}

// Validate validates this inbound message request
func (m *InboundMessageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundMessageRequest) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InboundMessageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InboundMessageRequest) UnmarshalBinary(b []byte) error {
	var res InboundMessageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}