// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConversationReason Reasons for a failed message receipt.
//
// swagger:model ConversationReason
type ConversationReason struct {

	// The reason code for the failed message receipt.
	// Enum: [MessageExpired RateLimited MessageNotAllowed GeneralError UnsupportedMessage UnknownMessage InvalidMessageStructure InvalidDestination ServerError MediaTypeNotAllowed InvalidMediaContentLength RecipientOptedOut]
	Code string `json:"code,omitempty"`

	// Description of the reason for the failed message receipt.
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this conversation reason
func (m *ConversationReason) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversationReasonTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MessageExpired","RateLimited","MessageNotAllowed","GeneralError","UnsupportedMessage","UnknownMessage","InvalidMessageStructure","InvalidDestination","ServerError","MediaTypeNotAllowed","InvalidMediaContentLength","RecipientOptedOut"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationReasonTypeCodePropEnum = append(conversationReasonTypeCodePropEnum, v)
	}
}

const (

	// ConversationReasonCodeMessageExpired captures enum value "MessageExpired"
	ConversationReasonCodeMessageExpired string = "MessageExpired"

	// ConversationReasonCodeRateLimited captures enum value "RateLimited"
	ConversationReasonCodeRateLimited string = "RateLimited"

	// ConversationReasonCodeMessageNotAllowed captures enum value "MessageNotAllowed"
	ConversationReasonCodeMessageNotAllowed string = "MessageNotAllowed"

	// ConversationReasonCodeGeneralError captures enum value "GeneralError"
	ConversationReasonCodeGeneralError string = "GeneralError"

	// ConversationReasonCodeUnsupportedMessage captures enum value "UnsupportedMessage"
	ConversationReasonCodeUnsupportedMessage string = "UnsupportedMessage"

	// ConversationReasonCodeUnknownMessage captures enum value "UnknownMessage"
	ConversationReasonCodeUnknownMessage string = "UnknownMessage"

	// ConversationReasonCodeInvalidMessageStructure captures enum value "InvalidMessageStructure"
	ConversationReasonCodeInvalidMessageStructure string = "InvalidMessageStructure"

	// ConversationReasonCodeInvalidDestination captures enum value "InvalidDestination"
	ConversationReasonCodeInvalidDestination string = "InvalidDestination"

	// ConversationReasonCodeServerError captures enum value "ServerError"
	ConversationReasonCodeServerError string = "ServerError"

	// ConversationReasonCodeMediaTypeNotAllowed captures enum value "MediaTypeNotAllowed"
	ConversationReasonCodeMediaTypeNotAllowed string = "MediaTypeNotAllowed"

	// ConversationReasonCodeInvalidMediaContentLength captures enum value "InvalidMediaContentLength"
	ConversationReasonCodeInvalidMediaContentLength string = "InvalidMediaContentLength"

	// ConversationReasonCodeRecipientOptedOut captures enum value "RecipientOptedOut"
	ConversationReasonCodeRecipientOptedOut string = "RecipientOptedOut"
)

// prop value enum
func (m *ConversationReason) validateCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationReasonTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationReason) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *ConversationReason) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationReason) UnmarshalBinary(b []byte) error {
	var res ConversationReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
