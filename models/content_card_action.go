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

// ContentCardAction A card action that a user can take.
//
// swagger:model ContentCardAction
type ContentCardAction struct {

	// Text to be returned as the payload from a ButtonResponse when a button is clicked. The payload and text are a combination which will have to be unique across each card and carousel in order to determine which button was clicked in that card or carousel.
	Payload string `json:"payload,omitempty"`

	// The response text from the button click.
	Text string `json:"text,omitempty"`

	// Describes the type of action.
	// Enum: [Link Postback]
	Type string `json:"type,omitempty"`

	// A URL of a web page to direct the user to.
	URL string `json:"url,omitempty"`
}

// Validate validates this content card action
func (m *ContentCardAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var contentCardActionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Link","Postback"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentCardActionTypeTypePropEnum = append(contentCardActionTypeTypePropEnum, v)
	}
}

const (

	// ContentCardActionTypeLink captures enum value "Link"
	ContentCardActionTypeLink string = "Link"

	// ContentCardActionTypePostback captures enum value "Postback"
	ContentCardActionTypePostback string = "Postback"
)

// prop value enum
func (m *ContentCardAction) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contentCardActionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContentCardAction) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this content card action based on context it is used
func (m *ContentCardAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContentCardAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentCardAction) UnmarshalBinary(b []byte) error {
	var res ContentCardAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
