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

// ConversationAssociation conversation association
//
// swagger:model ConversationAssociation
type ConversationAssociation struct {

	// Communication ID
	// Required: true
	CommunicationID *string `json:"communicationId"`

	// Conversation ID
	// Required: true
	ConversationID *string `json:"conversationId"`

	// An external contact ID.  If not supplied, implies the conversation should be disassociated with any external contact.
	ExternalContactID string `json:"externalContactId,omitempty"`

	// Media type
	// Required: true
	// Enum: [CALL CALLBACK CHAT COBROWSE EMAIL MESSAGE SOCIAL_EXPRESSION VIDEO SCREENSHARE]
	MediaType *string `json:"mediaType"`
}

// Validate validates this conversation association
func (m *ConversationAssociation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommunicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationAssociation) validateCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("communicationId", "body", m.CommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *ConversationAssociation) validateConversationID(formats strfmt.Registry) error {

	if err := validate.Required("conversationId", "body", m.ConversationID); err != nil {
		return err
	}

	return nil
}

var conversationAssociationTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CALL","CALLBACK","CHAT","COBROWSE","EMAIL","MESSAGE","SOCIAL_EXPRESSION","VIDEO","SCREENSHARE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationAssociationTypeMediaTypePropEnum = append(conversationAssociationTypeMediaTypePropEnum, v)
	}
}

const (

	// ConversationAssociationMediaTypeCALL captures enum value "CALL"
	ConversationAssociationMediaTypeCALL string = "CALL"

	// ConversationAssociationMediaTypeCALLBACK captures enum value "CALLBACK"
	ConversationAssociationMediaTypeCALLBACK string = "CALLBACK"

	// ConversationAssociationMediaTypeCHAT captures enum value "CHAT"
	ConversationAssociationMediaTypeCHAT string = "CHAT"

	// ConversationAssociationMediaTypeCOBROWSE captures enum value "COBROWSE"
	ConversationAssociationMediaTypeCOBROWSE string = "COBROWSE"

	// ConversationAssociationMediaTypeEMAIL captures enum value "EMAIL"
	ConversationAssociationMediaTypeEMAIL string = "EMAIL"

	// ConversationAssociationMediaTypeMESSAGE captures enum value "MESSAGE"
	ConversationAssociationMediaTypeMESSAGE string = "MESSAGE"

	// ConversationAssociationMediaTypeSOCIALEXPRESSION captures enum value "SOCIAL_EXPRESSION"
	ConversationAssociationMediaTypeSOCIALEXPRESSION string = "SOCIAL_EXPRESSION"

	// ConversationAssociationMediaTypeVIDEO captures enum value "VIDEO"
	ConversationAssociationMediaTypeVIDEO string = "VIDEO"

	// ConversationAssociationMediaTypeSCREENSHARE captures enum value "SCREENSHARE"
	ConversationAssociationMediaTypeSCREENSHARE string = "SCREENSHARE"
)

// prop value enum
func (m *ConversationAssociation) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationAssociationTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationAssociation) validateMediaType(formats strfmt.Registry) error {

	if err := validate.Required("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", *m.MediaType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationAssociation) UnmarshalBinary(b []byte) error {
	var res ConversationAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
