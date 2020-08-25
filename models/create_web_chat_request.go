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

// CreateWebChatRequest create web chat request
//
// swagger:model CreateWebChatRequest
type CreateWebChatRequest struct {

	// The list of attributes to associate with the customer participant.
	Attributes map[string]string `json:"attributes,omitempty"`

	// The name of the customer participating in the web chat.
	CustomerName string `json:"customerName,omitempty"`

	// The ID of the langauge to use for routing.
	LanguageID string `json:"languageId,omitempty"`

	// The priority to assign to the conversation for routing.
	Priority int64 `json:"priority,omitempty"`

	// The name of the provider that is sourcing the web chat.
	// Required: true
	Provider *string `json:"provider"`

	// The ID of the queue to use for routing the chat conversation.
	// Required: true
	QueueID *string `json:"queueId"`

	// The list of skill ID's to use for routing.
	SkillIds []string `json:"skillIds"`
}

// Validate validates this create web chat request
func (m *CreateWebChatRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateWebChatRequest) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *CreateWebChatRequest) validateQueueID(formats strfmt.Registry) error {

	if err := validate.Required("queueId", "body", m.QueueID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateWebChatRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateWebChatRequest) UnmarshalBinary(b []byte) error {
	var res CreateWebChatRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}