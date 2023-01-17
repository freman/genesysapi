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

// CreateEmailRequest create email request
//
// swagger:model CreateEmailRequest
type CreateEmailRequest struct {

	// The list of attributes to associate with the customer participant.
	Attributes map[string]string `json:"attributes,omitempty"`

	// Specify OUTBOUND to send an email on behalf of a queue, or INBOUND to create an external conversation. An external conversation is one where the provider is not PureCloud based.
	// Enum: [OUTBOUND INBOUND]
	Direction string `json:"direction,omitempty"`

	// The external contact with which the email should be associated. This field is only valid for OUTBOUND email.
	ExternalContactID string `json:"externalContactId,omitempty"`

	// The ID of the flow to use for routing email conversation. This field is mutually exclusive with queueId
	FlowID string `json:"flowId,omitempty"`

	// The email address of the sender of the email.
	FromAddress string `json:"fromAddress,omitempty"`

	// The name of the sender of the email.
	FromName string `json:"fromName,omitempty"`

	// An HTML body content of the email.
	HTMLBody string `json:"htmlBody,omitempty"`

	// The ID of the language to use for routing.
	LanguageID string `json:"languageId,omitempty"`

	// The priority to assign to the conversation for routing.
	Priority int64 `json:"priority,omitempty"`

	// The name of the provider that is sourcing the emails. The Provider "PureCloud Email" is reserved for native emails.
	// Required: true
	Provider *string `json:"provider"`

	// The ID of the queue to use for routing the email conversation. This field is mutually exclusive with flowId
	QueueID string `json:"queueId,omitempty"`

	// The list of skill ID's to use for routing.
	SkillIds []string `json:"skillIds"`

	// The subject of the email
	Subject string `json:"subject,omitempty"`

	// A text body content of the email.
	TextBody string `json:"textBody,omitempty"`

	// The email address of the recipient of the email.
	ToAddress string `json:"toAddress,omitempty"`

	// The name of the recipient of the email.
	ToName string `json:"toName,omitempty"`
}

// Validate validates this create email request
func (m *CreateEmailRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createEmailRequestTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OUTBOUND","INBOUND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createEmailRequestTypeDirectionPropEnum = append(createEmailRequestTypeDirectionPropEnum, v)
	}
}

const (

	// CreateEmailRequestDirectionOUTBOUND captures enum value "OUTBOUND"
	CreateEmailRequestDirectionOUTBOUND string = "OUTBOUND"

	// CreateEmailRequestDirectionINBOUND captures enum value "INBOUND"
	CreateEmailRequestDirectionINBOUND string = "INBOUND"
)

// prop value enum
func (m *CreateEmailRequest) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createEmailRequestTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateEmailRequest) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmailRequest) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create email request based on context it is used
func (m *CreateEmailRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateEmailRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateEmailRequest) UnmarshalBinary(b []byte) error {
	var res CreateEmailRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
