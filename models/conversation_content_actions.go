// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConversationContentActions User actions available on the content. All actions are optional and all actions are executed simultaneously.
//
// swagger:model ConversationContentActions
type ConversationContentActions struct {

	// Text to be sent back in reply when the item is selected.
	Textback string `json:"textback,omitempty"`

	// A URL of a web page to direct the user to.
	URL string `json:"url,omitempty"`

	// The target window in which to open the URL. If empty will open a blank page or tab.
	URLTarget string `json:"urlTarget,omitempty"`
}

// Validate validates this conversation content actions
func (m *ConversationContentActions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConversationContentActions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationContentActions) UnmarshalBinary(b []byte) error {
	var res ConversationContentActions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}